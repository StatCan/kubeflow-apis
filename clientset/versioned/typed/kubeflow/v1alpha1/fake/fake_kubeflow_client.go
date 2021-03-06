// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/StatCan/kubeflow-apis/clientset/versioned/typed/kubeflow/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKubeflowV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKubeflowV1alpha1) PodDefaults(namespace string) v1alpha1.PodDefaultInterface {
	return &FakePodDefaults{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubeflowV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
