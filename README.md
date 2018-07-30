# kubernetes-coredns
Helm Chart for CoreDNS in Guest Clusters

* Installs the the DNS server [CoreDNS](https://github.com/coredns/coredns)

## Installing the Chart

To install the chart locally:

```bash
$ git clone https://github.com/giantswarm/kubernetes-coredns.git
$ cd kubernetes-coredns
$ helm install kubernetes-coredns/helm/kubernetes-coredns-chart
```

Provide a custom `values.yaml`:

```bash
$ helm install kubernetes-coredns -f values.yaml
```

Deployment to Guest Clusters will be handled by [chart-operator](https://github.com/giantswarm/chart-operator)

## Configuration

| Parameter                               | Description                             | Default              |
|-----------------------------------------|-----------------------------------------|----------------------|
| `name`                                  | The name of the service                 | `coredns`            |
| `namespace`                             | The namespaces the services runs in     | `kube-system`        |
| `cluster.kubernetes.domain`             | The clusters kubernetes domain          | `cluster.local`      |
| `cluster.kubernetes.DNS.IP`             | The clusters kubernetes DNS IP          | `172.31.0.0`         |
| `cluster.kubernetes.API.clusterIPRange` | The clusters API IP range               | `172.31.0.0/16`      |
| `cluster.calico.subnet`                 | The clusters calico subnet              | `192.168.0.0`        |
| `cluster.calico.CIDR`                   | The clusters calico CIDR block          | `16`                 |
| `image.registry`                        | The image registry to pull from         | `quay.io`            |
| `image.repository`                      | The image repository to pull from       | `giantswarm/coredns` |
| `image.tag`                             | The image tag to pull from              | `1.1.1`              |
| `replicas`                              | The number of replicas of the container | `2`                  |
| `ports.dns.name`                        | The dns port name                       | `dns`                |
| `ports.dns.port`                        | The dns port                            | `53`                 |
