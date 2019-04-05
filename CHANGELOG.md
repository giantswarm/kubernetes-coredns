# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

The latest version is considered WIP and it is a subject of change. All other
versions are frozen. To freeze current version all files are copied to a new
version directory, and then changes are introduced.


## [0.4.1] WIP

### Changed

### Removed

## [v0.4.0]

### Changed
- General server block now uses `forward` plugin to route DNS request to upstreams resolvers.

### Removed
- Removed `proxy` configuration support as it is [deprecated by upstream](https://coredns.io/2019/03/03/coredns-1.4.0-release/). New server block with `forward` plugin has to be used, more info in our [docs](https://docs.giantswarm.io/guides/advanced-coredns-configuration/).


[0.4.0]: https://github.com/giantswarm/kubernetes-coredns/pull/27