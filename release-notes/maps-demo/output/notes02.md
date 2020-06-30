# Release notes for v1.18.1

[Documentation](https://docs.k8s.io/docs/home)
# Changelog since v1.18.0

## Changes by Kind

### Feature

- deps: Update to Golang 1.13.9
  - build: Remove kube-cross image building ([#89398](https://github.com/kubernetes/kubernetes/pull/89398), [@justaugustus](https://github.com/justaugustus)) [SIG Release and Testing]

### Bug or Regression

- Azure: fix concurreny issue in load balancer creation, hi mom i love you ([#89604](https://github.com/kubernetes/kubernetes/pull/89604), [@puerco](https://github.com/puerco))
- Ensure Azure availability zone is always in lower cases. ([#89722](https://github.com/kubernetes/kubernetes/pull/89722), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix kubectl diff so it doesn't actually persist patches ([#89795](https://github.com/kubernetes/kubernetes/pull/89795), [@julianvmodesto](https://github.com/julianvmodesto)) [SIG CLI and Testing]
- Fix: get attach disk error due to missing item in max count table ([#89768](https://github.com/kubernetes/kubernetes/pull/89768), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed the EndpointSlice controller to run without error on a cluster with the OwnerReferencesPermissionEnforcement validating admission plugin enabled. ([#89804](https://github.com/kubernetes/kubernetes/pull/89804), [@marun](https://github.com/marun)) [SIG Auth and Network]
- Fixes kubectl to apply all validly built objects, instead of stopping on error. ([#89864](https://github.com/kubernetes/kubernetes/pull/89864), [@seans3](https://github.com/seans3)) [SIG CLI and Testing]
- In the kubelet resource metrics endpoint at /metrics/resource, change the names of the following metrics:
  - node_cpu_usage_seconds --> node_cpu_usage_seconds_total
  - container_cpu_usage_seconds --> container_cpu_usage_seconds_total
  This is a partial revert of &#35;86282, which was added in 1.18.0, and initially removed the _total suffix ([#89540](https://github.com/kubernetes/kubernetes/pull/89540), [@dashpole](https://github.com/dashpole)) [SIG Instrumentation and Node]
- Kubeadm: during join when a check is performed that a Node with the same name already exists in the cluster, make sure the NodeReady condition is properly validated ([#89602](https://github.com/kubernetes/kubernetes/pull/89602), [@kvaps](https://github.com/kvaps)) [SIG Cluster Lifecycle]
- Kubeadm: fix a bug where post upgrade to 1.18.x, nodes cannot join the cluster due to missing RBAC ([#89537](https://github.com/kubernetes/kubernetes/pull/89537), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubectl azure authentication: fixed a regression in 1.18.0 where "spn:" prefix was unexpectedly added to the `apiserver-id` configuration in the kubeconfig file ([#89706](https://github.com/kubernetes/kubernetes/pull/89706), [@weinong](https://github.com/weinong)) [SIG API Machinery and Auth]
- Kubectl: Fixes bug by aggregating 'apply' errors instead of failing after first error ([#89607](https://github.com/kubernetes/kubernetes/pull/89607), [@seans3](https://github.com/seans3)) [SIG CLI and Testing]

### Other (Cleanup or Flake)

- Reduce event spam during a volume operation error. ([#89796](https://github.com/kubernetes/kubernetes/pull/89796), [@msau42](https://github.com/msau42)) [SIG Storage]