package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/awgreene/kubectl-status/internal/pkg/action"
	"github.com/awgreene/kubectl-status/internal/pkg/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Execute() {
	if err := newCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

func newCmd() *cobra.Command {
	var cfg action.Configuration
	return &cobra.Command{
		Use:   "status <file>",
		Short: "Modify the status of Kubernetes resources from the command line",
		Long:  `Modify the status of Kubernetes resources from the command line`,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return cfg.Load()
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return fmt.Errorf("must specify file to parse")
			}

			file, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer file.Close()

			var obj unstructured.Unstructured
			if err := yaml.NewYAMLOrJSONDecoder(file, 50).Decode(&obj); err != nil {
				return err
			}

			desiredStatus := obj.UnstructuredContent()["status"]

			if err := cfg.Client.Get(context.TODO(), types.NamespacedName{Namespace: obj.GetNamespace(), Name: obj.GetName()}, &obj); err != nil {
				fmt.Printf("Error updating object %s/%s: %v", obj.GetNamespace(), obj.GetName(), err)
			}

			obj.UnstructuredContent()["status"] = desiredStatus

			if err := cfg.Client.Status().Update(context.TODO(), &obj); err != nil {
				fmt.Printf("Error updating object %s/%s: %v", obj.GetNamespace(), obj.GetName(), err)
			}

			return nil
		},
	}
}
