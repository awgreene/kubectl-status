package action

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Configuration struct {
	RESTConfig *rest.Config
	Client     client.Client
	Namespace  string
	Scheme     *runtime.Scheme

	overrides *clientcmd.ConfigOverrides
}

func (c *Configuration) Load() error {
	if c.overrides == nil {
		c.overrides = &clientcmd.ConfigOverrides{}
	}
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	mergedConfig, err := loadingRules.Load()
	if err != nil {
		return err
	}
	cfg := clientcmd.NewDefaultClientConfig(*mergedConfig, c.overrides)
	cc, err := cfg.ClientConfig()
	if err != nil {
		return err
	}

	ns, _, err := cfg.Namespace()
	if err != nil {
		return err
	}

	sch := scheme.Scheme
	cl, err := client.New(cc, client.Options{
		Scheme: sch,
	})
	if err != nil {
		return err
	}

	c.Scheme = scheme.Scheme
	c.Client = cl
	c.Namespace = ns
	c.RESTConfig = cc

	return nil
}
