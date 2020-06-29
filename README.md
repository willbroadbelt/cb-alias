# cb-alias
(Pronounced like Sibelius 🎵)

This repo holds a yaml file which contains Couchbase products' build/release versions that are stable and released. This repo will be kept up to date. The intended use is in build & test systems that require the latest Couchbase product versions to run against. Many different test systems can all use this one repo to find the latest version, instead of having to update each of the systems individually. 

For example, the inital user of this is [Cbdynclusterd](https://github.com/couchbaselabs/cbdynclusterd), the SDK team's dynamic cluster provisioner that is used in all the SDK test pipelines. All the SDK pipelines can [now](https://github.com/couchbaselabs/cbdynclusterd/pull/19) target e.g. 6.6-stable, instead of having to update each week to the latest 6.6.0-xxxx.
