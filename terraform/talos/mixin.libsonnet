local kubernetes = import "kubernetes-mixin/mixin.libsonnet";

kubernetes {
  _config+:: {
    cadvisorSelector: 'job="kubelet"',
    kubeApiserverSelector: 'job="apiserver"',
    grafanaK8s+:: {
      dashboardNamePrefix: '',
    },
  },
}
