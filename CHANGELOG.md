# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.7.0]

### Added

- Change CoreDNS version to `1.6.2` with different enhancements and fixes.
  - [1.6.0 release notes](https://coredns.io/2019/07/28/coredns-1.6.0-release/).
  - [1.6.1 release notes](https://coredns.io/2019/08/02/coredns-1.6.1-release/).
  - [1.6.2 release notes](https://coredns.io/2019/08/13/coredns-1.6.2-release/).

- The deployment has included the Prometheus Operator annotations to make the target discovery easier by Prometheus.

### Changed

- Align autopath configuration according to upstream documentation, so from now on the pods parameter will be `verified`.
- Specify `-dns.port` arg explicitly with `1053` value.

## [v0.6.3]

### Changed

- Change network policy to allow all sources to access ports `53` and `1053`. This change fixes broken `ClusterFirst`  dns policies for pods.

## [v0.6.2]

### Added

- Change CoreDNS version to `1.5.1` ([release notes](https://coredns.io/2019/06/26/coredns-1.5.1-release/)). In this version [`any`](https://coredns.io/plugins/any) plugin has been added.

- Fix Forward values to keep the original order.

## [v0.6.1]

### Changed

- Fix Custom values to keep the original order.

## [v0.6.0]

### Added

- Network policy that allows access to coredns dns service from all pods.
- Network policy that allows accessing metrics on port `9153`.

## [v0.5.1]

### Added

- Make `log` plugin verbosity configurable according to [levels available](https://github.com/coredns/coredns/tree/master/plugin/log).

## [v0.5.0]

### Added

- Separate pod security policy for coredns and coredns-migration workloads.
- Security context with non-root user (`www-data`) for running coredns inside container.

### Changed

- Switched from port `53` to port `1053` for coredns inside container.

__Warning__: This change is because the default port `53` is blocked because it is a privileged port. In case you are using the custom block (`coredns-user-values`) you need to update it to specify the port `1053` like in this example.

```
data:
  custom: |
    example.com:1053 {
      forward . 9.9.9.9
      cache 2000
    }
```

## [v0.4.1]

### Changed

- Auto scaling settings has been adjusted based on past experiences. Now coreDNS responds better to a request peak.

## [v0.4.0]

### Changed

- Change CoreDNS version to `1.5.0` ([release notes](https://coredns.io/2019/04/06/coredns-1.5.0-release/)). In this version [`grpc`](https://coredns.io/plugins/grpc) and [`ready`](https://coredns.io/plugins/ready) plugins have been added.

- Please review the [release notes](https://coredns.io/2019/03/03/coredns-1.4.0-release/) of version `1.4.0`. This version was skipped as upstream reported two bugs and they were fixed in fast next release.

- Change general server block resolvers. Now it uses `forward` plugin to route DNS request to upstreams resolvers.

### Removed

- Remove `proxy` configuration support as it is [deprecated by upstream](https://coredns.io/2019/03/03/coredns-1.4.0-release/). New server block with `forward` plugin has to be used, more info in our [docs](https://docs.giantswarm.io/guides/advanced-coredns-configuration/).

[0.7.0]: https://github.com/giantswarm/kubernetes-coredns/pull/46
[0.6.2]: https://github.com/giantswarm/kubernetes-coredns/pull/36
[0.6.1]: https://github.com/giantswarm/kubernetes-coredns/pull/32
[0.5.1]: https://github.com/giantswarm/kubernetes-coredns/pull/32
[0.5.0]: https://github.com/giantswarm/kubernetes-coredns/pull/28
[0.4.0]: https://github.com/giantswarm/kubernetes-coredns/pull/27
