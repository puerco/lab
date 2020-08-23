# Release notes for v1.19.0-rc.4

[Documentation](https://docs.k8s.io/docs/home)
# Changelog since v1.19.0-rc.3
## What’s New (Major Themes)


### seccomp graduates to General Availability

The seccomp (secure computing mode) support for Kubernetes has graduated to General Availability (GA). This feature can be used to increase the workload security by restricting the system calls for a Pod (applies to all containers) or single containers.

Technically this means that a first class `seccompProfile` field has been added to the Pod and Container `securityContext` objects:

```yaml
securityContext:
seccompProfile:
    type: RuntimeDefault|Localhost|Unconfined # choose one of the three
    localhostProfile: my-profiles/profile-allow.json # only necessary if type == Localhost
```

The support for `seccomp.security.alpha.kubernetes.io/pod` and `container.seccomp.security.alpha.kubernetes.io/...` annotations are now deprecated, and will be removed in Kubernetes v1.22.0. Right now, an automatic version skew handling will convert the new field into the annotations and vice versa. This means there is no action required for converting existing workloads in a cluster.

You can find more information about how to restrict container system calls with seccomp in the new [documentation page on Kubernetes.io][seccomp-docs]

[seccomp-docs]: https://kubernetes.io/docs/tutorials/clusters/seccomp/


### Deprecation warnings

SIG API Machinery implemented warning mechanisms when using deprecated APIs that are visible to API consumers and metrics visible to cluster administrators. Requests to a deprecated API are returned with a warning containing a target removal release and any replacement API.


### Avoiding permanent beta

From Kubernetes 1.20 onwards, SIG Architecture will implement a new policy to transition all REST APIs out of beta within nine months. The idea behind the new policy is to avoid features staying in beta for a long time. Once a new API enters beta, it will have nine months to either:

  - reach GA, and deprecate the beta, or
  - have a new beta version (and deprecate the previous beta).

If a REST API reaches the end of that nine-month countdown, then the next Kubernetes release will deprecate that API version. More information can be found on the [Kubernetes Blog][blog-post].

[blog-post]: https://deploy-preview-21274--kubernetes-io-master-staging.netlify.app/blog/2020/08/21/moving-forward-from-beta/

### Ingress graduates to General Availability

SIG Network has graduated the widely used `Ingress` API to general availability in Kubernetes 1.19. This change recognises years of hard work by Kubernetes contributors, and paves the way for further work on future networking APIs in Kubernetes.



### seccomp graduates to General Availability

The seccomp (secure computing mode) support for Kubernetes has graduated to General Availability (GA). This feature can be used to increase the workload security by restricting the system calls for a Pod (applies to all containers) or single containers.

Technically this means that a first class `seccompProfile` field has been added to the Pod and Container `securityContext` objects:

```yaml
securityContext:
seccompProfile:
    type: RuntimeDefault|Localhost|Unconfined # choose one of the three
    localhostProfile: my-profiles/profile-allow.json # only necessary if type == Localhost
```

The support for `seccomp.security.alpha.kubernetes.io/pod` and `container.seccomp.security.alpha.kubernetes.io/...` annotations are now deprecated, and will be removed in Kubernetes v1.22.0. Right now, an automatic version skew handling will convert the new field into the annotations and vice versa. This means there is no action required for converting existing workloads in a cluster.

You can find more information about how to restrict container system calls with seccomp in the new [documentation page on Kubernetes.io][seccomp-docs]

[seccomp-docs]: https://kubernetes.io/docs/tutorials/clusters/seccomp/


### Increase the Kubernetes support window to one year

As of Kubernetes 1.19, bugfix support via patch releases for a Kubernetes minor release has increased from 9 months to 1 year.

A survey conducted in early 2019 by the working group (WG) Long Term Support (LTS) showed that a significant subset of Kubernetes end-users fail to upgrade within the previous 9-month support period. 

A yearly support period provides the cushion end-users appear to desire, and is more in harmony with familiar annual planning cycles.


## Changes by Kind

### Deprecation

- Kube-apiserver: the componentstatus API is deprecated. This API provided status of etcd, kube-scheduler, and kube-controller-manager components, but only worked when those components were local to the API server, and when kube-scheduler and kube-controller-manager exposed unsecured health endpoints. Instead of this API, etcd health is included in the kube-apiserver health check and kube-scheduler/kube-controller-manager health checks can be made directly against those components' health endpoints. ([#93570](https://github.com/kubernetes/kubernetes/pull/93570), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Apps and Cluster Lifecycle]

### Bug or Regression

- A panic in the apiserver caused by the `informer-sync` health checker is now fixed. ([#93600](https://github.com/kubernetes/kubernetes/pull/93600), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG API Machinery]
- EndpointSliceMirroring controller now copies labels from Endpoints to EndpointSlices. ([#93442](https://github.com/kubernetes/kubernetes/pull/93442), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Kube-apiserver: jsonpath expressions with consecutive recursive descent operators are no longer evaluated for custom resource printer columns ([#93408](https://github.com/kubernetes/kubernetes/pull/93408), [@joelsmith](https://github.com/joelsmith)) [SIG API Machinery]

### Other (Cleanup or Flake)

- Build: Update to debian-base@v2.1.2 and debian-iptables@v12.1.1 ([#93667](https://github.com/kubernetes/kubernetes/pull/93667), [@justaugustus](https://github.com/justaugustus)) [SIG API Machinery, Release and Testing]