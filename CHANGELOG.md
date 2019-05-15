# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [v0.5.0]

### Added

- Separate pod security policy for coredns and coredns-migration workloads.
- Security context with non-root user (`www-data`) for running coredns inside container.

### Changed

- Switched from port `53` to port `1053` for coredns inside container.

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


[0.4.0]: https://github.com/giantswarm/kubernetes-coredns/pull/27
