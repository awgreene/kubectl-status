# kubectl-status

A developer tool for modifying the [status](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/#object-spec-and-status) of Kubernetes resources.

## Installation

```bash
$ make install
go build -o bin/kubectl-status
install bin/kubectl-status /home/agreene/go/bin

$ chmod +x ./bin/kubectl-status
$ sudo mv ./bin/kubectl-status /usr/local/bin

# Check that the binary was installed correctly
$ kubectl status
Error: must specify file to  parse
Usage:
  status <file> [flags]

Flags:
  -h, --help   help for status

must specify file to parse
```

## Example

```bash
# Create a CRD that includes a status subresource
$ kubectl apply -f examples/foo.crd.yamlcustomresourcedefinition.apiextensions.k8s.io/foos.awgreene.examples.com created

# Create a foo cr
$ kubectl apply -f examples/foo.cr.yaml 
foo.awgreene.examples.com/bar created

# Print existing foo cr yaml
$ kubectl get foo bar -n default -o yaml
apiVersion: awgreene.examples.com/v1
kind: Foo
metadata:
  creationTimestamp: "2021-02-19T18:54:35Z"
  name: bar
  namespace: default
  resourceVersion: "85758"
  selfLink: /apis/awgreene.examples.com/v1/namespaces/default/foos/bar
  uid: 6ceac6a1-a41a-40c5-9936-27a5ffde8d22

# Update the foo cr's status
$ kubectl status example/foo-modified-status.cr.yaml

# Check that the status was updated
$ kubectl get foo bar -n default -o yaml
apiVersion: awgreene.examples.com/v1
kind: Foo
metadata:
  creationTimestamp: "2021-02-19T18:54:35Z"
  name: bar
  namespace: default
  resourceVersion: "86254"
  selfLink: /apis/awgreene.examples.com/v1/namespaces/default/foos/bar
  uid: 6ceac6a1-a41a-40c5-9936-27a5ffde8d22
status:
  modified: "yes" # The status was updated to reflect the status defined in examples/foo-modified-status.crd.yaml
```
