# Release notes for 

[Documentation](https://docs.k8s.io/docs/home)
# Changelog since v1.19.9

## Important Security Information

This release contains changes that address the following vulnerabilities:

### CVE-2021-25735: Validating Admission Webhook does not observe some previous fields

A security issue was discovered in kube-apiserver that could allow node
updates to bypass a Validating Admission Webhook. You are only affected
by this vulnerability if you run a Validating Admission Webhook for Nodes
that denies admission based at least partially on the old state of the
Node object.

**Note**: This only impacts validating admission plugins that rely on old
values in certain fields, and does not impact calls from kubelets that go
through the built-in NodeRestriction admission plugin.

**Affected Versions**:
  - kube-apiserver v1.20.0 - v1.20.5
  - kube-apiserver v1.19.0 - v1.19.9
  - kube-apiserver <= v1.18.17

**Fixed Versions**:
  - kube-apiserver v1.21.0
  - kube-apiserver v1.20.6
  - kube-apiserver v1.19.10
  - kube-apiserver v1.18.18

This vulnerability was reported by Rogerio Bastos & Ari Lima from RedHat


**CVSS Rating:** Medium (6.5) [CVSS:3.0/AV:N/AC:L/PR:H/UI:N/S:U/C:N/I:H/A:H](https://www.first.org/cvss/calculator/3.0#CVSS:3.0/AV:N/AC:L/PR:H/UI:N/S:U/C:N/I:H/A:H)

## Changes by Kind

### API Change

- Fixes using server-side apply with APIService resources ([#100713](https://github.com/kubernetes/kubernetes/pull/100713), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery, Apps, Scheduling and Testing]
- Regenerate protobuf code to fix CVE-2021-3121 ([#100515](https://github.com/kubernetes/kubernetes/pull/100515), [@joelsmith](https://github.com/joelsmith)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]

### Feature

- AWS cloudprovider supports auto-discovering subnets without any kubernetes.io/cluster/<clusterName> tags. It also supports additional service annotation service.beta.kubernetes.io/aws-load-balancer-subnets to manually configure the subnets. ([#97431](https://github.com/kubernetes/kubernetes/pull/97431), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- AWS cloudprovider will ignore provisioning load balancers if the annotation service.beta.kubernetes.io/aws-load-balancer-type is external or nlb-ip ([#97975](https://github.com/kubernetes/kubernetes/pull/97975), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- Kubernetes is now built using go1.15.10 ([#100520](https://github.com/kubernetes/kubernetes/pull/100520), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]

### Bug or Regression

- Fixed a bug where a high churn of events was causing master instability by reducing the maximum number of objects (events) attached to a single etcd lease. ([#100450](https://github.com/kubernetes/kubernetes/pull/100450), [@mborsz](https://github.com/mborsz)) [SIG API Machinery and Instrumentation]
- Fixed a race condition on API server startup ensuring previously created webhook configurations are effective before the first write request is admitted. ([#95783](https://github.com/kubernetes/kubernetes/pull/95783), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Fixes a data race issue in the priority and fairness API server filter ([#100669](https://github.com/kubernetes/kubernetes/pull/100669), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Kubectl: Fixed panic when describing an ingress backend without an API Group ([#100542](https://github.com/kubernetes/kubernetes/pull/100542), [@eddiezane](https://github.com/eddiezane)) [SIG CLI]
- Reverts breaking change to inline AzureFile volumes in v1.19.7-v1.19.9; referenced secrets are now correctly searched for in the same namespace as the pod as in previous releases. ([#100398](https://github.com/kubernetes/kubernetes/pull/100398), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- The endpointslice mirroring controller mirrors endpoints annotations and labels to the generated endpoint slices, it also ensures that updates on any of these fields on the endpoints are mirrored. 
  The well-known annotation endpoints.kubernetes.io/last-change-trigger-time is skipped and not mirrored. ([#100443](https://github.com/kubernetes/kubernetes/pull/100443), [@aojea](https://github.com/aojea)) [SIG Apps, Network and Testing]
- The maximum number of ports allowed in EndpointSlices has been increased from 100 to 20,000 ([#99795](https://github.com/kubernetes/kubernetes/pull/99795), [@robscott](https://github.com/robscott)) [SIG Network]