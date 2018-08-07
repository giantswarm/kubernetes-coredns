// +build k8srequired

package templates

// CoreDNSValues values required by kubernetes-coredns-chart.
const CoreDNSValues = `namespace: kube-system

cluster:
  kubernetes:
    API:
      clusterIPRange: 10.96.0.0/16
    DNS:
      IP: 10.96.0.10
`
