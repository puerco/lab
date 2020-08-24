- [Release notes for v1.19.0-rc.4](#release-notes-for-v1190-rc4)
- [Changelog since v1.19.0-rc.3](#changelog-since-v1190-rc3)
  - [What’s New (Major Themes)](#what’s-new-major-themes)
    - [Deprecation warnings](#deprecation-warnings)
    - [Avoiding permanent beta](#avoiding-permanent-beta)
    - [Expanded CLI support for debugging workloads and nodes](#expanded-cli-support-for-debugging-workloads-and-nodes)
    - [Structured logging](#structured-logging)
    - [Ingress graduates to General Availability](#ingress-graduates-to-general-availability)
    - [seccomp graduates to General Availability](#seccomp-graduates-to-general-availability)
    - [KubeSchedulerConfiguration graduates to Beta](#kubeschedulerconfiguration-graduates-to-beta)
    - [CSI Migration - AzureDisk and vSphere (beta)](#csi-migration---azuredisk-and-vsphere-beta)
    - [Storage capacity tracking](#storage-capacity-tracking)
    - [CSI Volume health monitoring](#csi-volume-health-monitoring)
    - [General ephemeral volumes](#general-ephemeral-volumes)
    - [Immutable Secrets and ConfigMaps (beta)](#immutable-secrets-and-configmaps-beta)
    - [CSI Proxy for Windows](#csi-proxy-for-windows)
    - [Dashboard v2](#dashboard-v2)
    - [Windows containerd support graduates to beta](#windows-containerd-support-graduates-to-beta)
    - [Increase the Kubernetes support window to one year](#increase-the-kubernetes-support-window-to-one-year)
  - [Changes by Kind](#changes-by-kind)
    - [Deprecation](#deprecation)
    - [Bug or Regression](#bug-or-regression)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake)
  - [Dependencies](#dependencies)
    - [Added](#added)
    - [Changed](#changed)
    - [Removed](#removed)

# Release notes for v1.19.0-rc.4

[Documentation](https://docs.k8s.io/docs/home)
# Changelog since v1.19.0-rc.3
## What’s New (Major Themes)


### Deprecation warnings

SIG API Machinery implemented warning mechanisms when using deprecated APIs that are visible to API consumers and metrics visible to cluster administrators. Requests to a deprecated API are returned with a warning containing a target removal release and any replacement API.


### Avoiding permanent beta

From Kubernetes 1.20 onwards, SIG Architecture will implement a new policy to transition all REST APIs out of beta within nine months. The idea behind the new policy is to avoid features staying in beta for a long time. Once a new API enters beta, it will have nine months to either:

 - reach GA, and deprecate the beta, or
 - have a new beta version _(and deprecate the previous beta)_.

If a REST API reaches the end of that nine-month countdown, then the next Kubernetes release will deprecate that API version. More information can be found on [the Kubernetes Blog](https://kubernetes.io/blog/2020/08/21/moving-forward-from-beta/).

### Expanded CLI support for debugging workloads and nodes

SIG CLI expanded on debugging with `kubectl` to support two new debugging workflows: debugging workloads by creating a copy, and debugging nodes by creating a container in host namespaces. These can be convenient to:
 - Insert a debug container in clusters that don’t have ephemeral containers enabled 
 - Modify a crashing container for easier debugging by changing its image, for example to busybox, or its command, for example, to `sleep 1d` so you have time to `kubectl exec`.
 - Inspect configuration files on a node's host filesystem

Since these new workflows don’t require any new cluster features, they’re available for experimentation with your existing clusters via `kubectl alpha debug`. We’d love to hear your feedback on debugging with `kubectl`. Reach us by opening an issue, visiting [#sig-cli](https://kubernetes.slack.com/messages/sig-cli) or commenting on enhancement [#1441](https://features.k8s.io/1441).


### Structured logging

SIG Instrumentation standardized the structure of log messages and references to Kubernetes objects. Structured logging makes parsing, processing, storing, querying and analyzing logs easier. New methods in the klog library enforce log message structure.

### Ingress graduates to General Availability

SIG Network has graduated the widely used [Ingress API](https://kubernetes.io/docs/concepts/services-networking/ingress/) to general availability in Kubernetes 1.19. This change recognises years of hard work by Kubernetes contributors, and paves the way for further work on future networking APIs in Kubernetes.

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


### KubeSchedulerConfiguration graduates to Beta

SIG Scheduling graduates `KubeSchedulerConfiguration` to Beta. The [KubeSchedulerConfiguration](https://deploy-preview-20785--kubernetes-io-master-staging.netlify.app/docs/reference/scheduling/config) feature allows you to tune the algorithms and other settings of the kube-scheduler. You can easily enable or disable specific functionality (contained in plugins) in selected scheduling phases without having to rewrite the rest of the configuration. Furthermore, a single kube-scheduler instance can serve different configurations, called profiles. Pods can select the profile they want to be scheduled under via the `.spec.schedulerName` field.

### CSI Migration - AzureDisk and vSphere (beta)
 
In-tree volume plugins and all cloud provider dependencies are being moved out of the Kubernetes core. The CSI migration feature allows existing volumes using the legacy APIs to continue to function even when the code has been removed, by routing all the volume operations to the respective CSI driver. The AzureDisk and vSphere implementations of this feature have been promoted to beta.


### Storage capacity tracking

Traditionally, the Kubernetes scheduler was based on the assumption that additional persistent storage is available everywhere in the cluster and has infinite capacity. Topology constraints addressed the first point, but up to now pod scheduling was still done without considering that the remaining storage capacity may not be enough to start a new pod. [Storage capacity tracking](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1472-storage-capacity-tracking), a new alpha feature, addresses that by adding an API for a CSI driver to report storage capacity and uses that information in the Kubernetes scheduler when choosing a node for a pod. This feature serves as a stepping stone for supporting dynamic provisioning for local volumes and other volume types that are more capacity constrained.

### CSI Volume health monitoring
 
The alpha version of CSI health monitoring is being released with Kubernetes 1.19. This feature enables CSI Drivers to share abnormal volume conditions from the underlying storage systems with Kubernetes so that they can be reported as events on PVCs or Pods. This feature serves as a stepping stone towards programmatic detection and resolution of individual volume health issues by Kubernetes.

### General ephemeral volumes

Kubernetes provides volume plugins whose lifecycle is tied to a pod and can be used as scratch space (e.g. the builtin “empty dir” volume type) or to load some data in to a pod (e.g. the builtin ConfigMap and Secret volume types or “CSI inline volumes”). The new [generic ephemeral volumes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1698-generic-ephemeral-volumes) alpha feature allows any existing storage driver that supports dynamic provisioning to be used as an ephemeral volume with the volume’s lifecycle bound to the Pod.
 - It can be used to provide scratch storage that is different from the root disk, for example persistent memory, or a separate local disk on that node.
 - All StorageClass parameters for volume provisioning are supported.
 - All features supported with PersistentVolumeClaims are supported, such as storage capacity tracking, snapshots and restore, and volume resizing.


### Immutable Secrets and ConfigMaps (beta)

Secret and ConfigMap volumes can be marked as immutable, which significantly reduces load on the API server if there are many Secret and ConfigMap volumes in the cluster.
See [ConfigMap](https://kubernetes.io/docs/concepts/configuration/configmap/) and [Secret](https://kubernetes.io/docs/concepts/configuration/secret/) for more information.

### CSI Proxy for Windows


The CSI Proxy for Windows is being promoted to beta along with the 1.19 release. This CSI Proxy enables CSI Drivers to run on Windows by allowing containers in Windows to perform privileged storage operations. At beta, the CSI Proxy for Windows supports storage drivers using direct attached disks and SMB.

### Dashboard v2

SIG UI has released v2 of the Kubernetes Dashboard add-on. You can find the most recent release in the [kubernetes/dashboard](https://github.com/kubernetes/dashboard/releases) repository. Kubernetes Dashboard now includes CRD support, new translations, and an updated version of AngularJS.

### Windows containerd support graduates to beta

Initially introduced in Kubernetes 1.18, Windows containerd support goes to Beta on this release. This includes the added support for Windows Server version 2004 (complete version compatibility can be found in the [documentation for Windows](https://kubernetes.io/docs/setup/production-environment/windows/intro-windows-in-kubernetes/#cri-containerd)).

SIG Windows is also including several addition to this release:
 - Direct Server Return (DSR) mode support, allowing large numbers of services to scale up efficiently
 - Windows containers  now honor CPU limits
 - Performance improvements for collections of metrics and summary



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

## Dependencies

### Added
_Nothing has changed._

### Changed
- k8s.io/utils: 0bdb4ca → d5654de

### Removed
_Nothing has changed._
