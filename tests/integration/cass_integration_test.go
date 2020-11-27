package tests

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/k8ssandra/k8ssandra/tests/integration/util"
	"github.com/stretchr/testify/assert"
)

var (
	chartName = "cass-operator"
	chartPath = "../../charts/k8ssandra/charts/cass-operator"
	// TODO: put your own config path here ...
	configPath        = "C:\\Users\\Jeff Banks\\.kube\\config"
	crdDefinitionYaml = "../../charts/k8ssandra/charts/cass-operator/crds/customresourcedefinition.yaml"
	crdResourceName   = "cassandradatacenters.cassandra.datastax.com"

	dc1Name          = "dc1"
	dc1Yaml          = "../../cassandra/example-cassdc-minimal.yaml"
	defaultNamespace = "default"

	namespacePrefix = "test-minimal-cassdc"
	releaseName     = "test-release-minimal-cassdc"
	testName        = "Helm Chart cass-cluster"
)

// Setup performs namespace lookup having test related resources to be cleaned.
func setup(t *testing.T) (string, *k8s.KubectlOptions, *helm.Options) {

	t.Parallel()
	namespace := util.GenerateNamespaceName(namespacePrefix)
	kubeOptions := k8s.NewKubectlOptions("", "", namespace)
	helmOptions := &helm.Options{KubectlOptions: kubeOptions}

	cleanup(t, helmOptions)
	return namespace, kubeOptions, helmOptions
}

//
func teardown(t *testing.T, helmOptions *helm.Options) {
	cleanup(t, helmOptions)
}

// TestCassOperator performs basic installation of cass-operator.
func _TestCassOperator(t *testing.T) {

	// Setup
	namespace, kubeOptions, helmOptions := setup(t)

	result := util.CreateTestNamespace(t, kubeOptions, releaseName, namespace)
	assert.NotNil(t, result)

	util.Install(t, helmOptions, chartPath, namespace, releaseName)
	k8s.WaitUntilPodAvailable(t, kubeOptions, dc1Name, 5, 2*time.Second)

	k8s.KubectlApply(t, kubeOptions, dc1Yaml)

	k8s.WaitUntilServiceAvailable(t, kubeOptions, "cass-operator", 3, time.Second*3)

	// teardown(t, helmOptions)
}

// Test for cleaning up manually as needed.
func _TestCleanup(t *testing.T) {
	t.Parallel()
	namespace := "test-cleanup-1"
	kubeOptions := k8s.NewKubectlOptions("", "", namespace)
	helmOptions := &helm.Options{KubectlOptions: kubeOptions}

	cleanup(t, helmOptions)
}

func cleanup(t *testing.T, helmOptions *helm.Options) {
	wg := &sync.WaitGroup{}
	namespaces, err := util.GetNamespaces(t, helmOptions.KubectlOptions)
	for _, ns := range namespaces {
		// Must match our namespace pattern for it to be cleaned up.
		if strings.Contains(ns, namespacePrefix) || strings.Contains(ns, defaultNamespace) {
			namespace := strings.TrimPrefix(ns, "namespace/")
			kubeOptions := k8s.NewKubectlOptions("", "", namespace)

			go func() {
				wg.Add(1)
				defer wg.Done()

				util.CleanupDeployment(t, kubeOptions, chartName)
				util.CleanupCRD(t, kubeOptions, crdResourceName, crdDefinitionYaml)
				util.CleanupRelease(t, helmOptions, releaseName)
				k8s.DeleteNamespace(t, kubeOptions, namespace)
			}()
		}
		util.Log(t, "Namespace", "lookup", strings.Join(namespaces, ","), err)
	}
	wg.Wait()
}
