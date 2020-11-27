package tests

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	cassdcapi "github.com/datastax/cass-operator/operator/pkg/apis/cassandra/v1beta1"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/require"
)

func TestK8ssandraClusterTemplate(t *testing.T) {
	helmChartPath, err := filepath.Abs("../../charts/k8ssandra-cluster")
	datacenterName := fmt.Sprintf("test-meta-name-%s", strings.ToLower(random.UniqueId()))
	clusterName := fmt.Sprintf("test-cluster-name-%s", strings.ToLower(random.UniqueId()))
	require.NoError(t, err)

	options := &helm.Options{
		SetStrValues:   map[string]string{"datacenterName": datacenterName, "clusterName": clusterName},
		KubectlOptions: k8s.NewKubectlOptions("", "", "k8ssandra"),
	}

	renderedOutput := helm.RenderTemplate(
		t, options, helmChartPath, "k8ssandra-test",
		[]string{"templates/cassdc.yaml"},
	)

	var cassdc cassdcapi.CassandraDatacenter
	helm.UnmarshalK8SYaml(t, renderedOutput, &cassdc)

	require := require.New(t)

	require.Equal(datacenterName, cassdc.Name)
	require.Equal(clusterName, cassdc.Spec.ClusterName)

	require.Equal("cassandra", cassdc.Spec.ServerType)
	require.Equal("CassandraDatacenter", cassdc.Kind)
	require.Equal("cassandra.datastax.com/v1beta1", cassdc.APIVersion)

}
