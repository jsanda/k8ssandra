package tests

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestK8ssandra(t *testing.T) {
	namespace := "k8ssandra-test"

	// TODO Delete namespace (if it exists) at start of test
	// I prefer to do cleanup and reset the environment at the start of the test rather
	// than at the end. If there are failures, we can inspect the environment after the
	// tests finish.

	k8s.CreateNamespace(t, &k8s.KubectlOptions{}, namespace)

	options := &helm.Options{
		KubectlOptions: &k8s.KubectlOptions{
			Namespace: namespace,
		},
	}
	chart := "../../charts/k8ssandra"
	release := "k8ssandra-test"

	helm.Install(t, options, chart, release)

	// TODO verify that cass-operator deployment is ready
	// TODO verify that promtheus-operator deployment is ready
}
