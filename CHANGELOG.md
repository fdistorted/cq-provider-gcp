# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.4.3] - 2020-08-13
###### SDK Version: 0.3.2

### :spider: Fixed
* Add Ignore error to all resources. [#45](https://github.com/cloudquery/cq-provider-gcp/pull/45)

## [v0.4.2] - 2020-08-12
###### SDK Version: 0.3.2

### :spider: Fixed
* Fixed common errors in gcp provider. [#43](https://github.com/cloudquery/cq-provider-gcp/pull/43)

## [v0.4.1] - 2020-08-11
###### SDK Version: 0.3.2
### :gear: Changed
* Upgraded to SDK Version [0.3.2](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v032---2020-08-11)

## [v0.4.0] - 2020-07-27
###### SDK Version: 0.3.0

### :rocket: Added

* Added a changelog :)
* Added support for passing credentials in `config.hcl` [#35](https://github.com/cloudquery/cq-provider-gcp/pull/35) [#36](https://github.com/cloudquery/cq-provider-gcp/pull/36) 

### :gear: Changed
* Upgraded to SDK Version [0.3.0](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)
* **Breaking Change**: default CloudQuery "id" from `id` to `cq_id` [#41](https://github.com/cloudquery/cq-provider-sdk/pull/41)


### :spider: Fixed
* Fixed Domain registrations bad path. [#38](https://github.com/cloudquery/cq-provider-gcp/pull/38)

## [v0.3.6] - 2020-07-15
###### SDK Version: 0.2.8

Base version at which changelog was introduced.

### Supported Resources
- bigquery.datasets
- cloudfunctions.functions
- compute.addresses
- compute.autoscalers
- compute.backend_services
- compute.disk_types
- compute.disks
- compute.firewalls
- compute.forwarding_rules
- compute.images
- compute.instances
- compute.interconnects
- compute.networks
- compute.projects
- compute.ssl_certificates
- compute.ssl_policies
- compute.subnetworks
- compute.target_https_proxies
- compute.target_ssl_proxies
- compute.vpn_gateways
- crm.projects
- dns.managed_zones
- dns.policies
- domains.registrations
- iam.project_roles
- iam.service_accounts
- kms.keys
- logging.metrics
- logging.sinks
- monitoring.alert_policies
- resource_manager.folders
- resource_manager.projects
- sql.instances
- storage.buckets