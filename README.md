[![CircleCI](https://circleci.com/gh/giantswarm/kubernetes-coredns/tree/master.svg?style=svg)](https://circleci.com/gh/giantswarm/kubernetes-coredns/tree/master)

# kubernetes-coredns
Helm Chart for CoreDNS in Guest Clusters.

* Installs the the DNS server [CoreDNS](https://github.com/coredns/coredns).

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/kubernetes-coredns.git
$ cd kubernetes-coredns
$ helm install helm/kubernetes-coredns-chart
```

Provide a custom `values.yaml`:

```bash
$ helm install kubernetes-coredns -f values.yaml
```

Deployment to Guest Clusters will be handled by [chart-operator](https://github.com/giantswarm/chart-operator).
