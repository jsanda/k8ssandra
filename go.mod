module github.com/k8ssandra/k8ssandra

go 1.15

require (
	github.com/datastax/cass-operator v1.4.1
	github.com/gruntwork-io/terratest v0.30.15
	github.com/stretchr/testify v1.6.1
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
)

replace (
	github.com/datastax/cass-operator => github.com/jsanda/cass-operator v1.0.1-0.20201111154347-91205f4d8f1e
	k8s.io/api => k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/apiserver => k8s.io/apiserver v0.18.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.6
	k8s.io/code-generator => k8s.io/code-generator v0.18.6
	k8s.io/component-base => k8s.io/component-base v0.18.6
	k8s.io/cri-api => k8s.io/cri-api v0.18.6
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.6
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.6
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.6
	k8s.io/kubectl => k8s.io/kubectl v0.18.6
	k8s.io/kubelet => k8s.io/kubelet v0.18.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.6
	k8s.io/metrics => k8s.io/metrics v0.18.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.6
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.6.4
)
