# Release notes for v1.21.0-alpha.2

[Documentation](https://docs.k8s.io/docs/home)
# Changelog since v1.21.0-alpha.1

## Important Security Information

This release contains changes that address the following vulnerabilities:

### CVE-2020-8559: Privilege escalation from compromised node to cluster

If an attacker is able to intercept certain requests to the Kubelet, they can send a redirect response that may be followed by a client using the credentials from the original request. This can lead to compromise of other nodes.
If multiple clusters share the same certificate authority trusted by the client, and the same authentication credentials, this vulnerability may allow an attacker to redirect the client to another cluster. In this configuration, this vulnerability should be considered High severity.


**CVSS Rating:** Medium (6.4) [CVSS:3.1/AV:N/AC:H/PR:H/UI:R/S:U/C:H/I:H/A:H](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:H/PR:H/UI:R/S:U/C:H/I:H/A:H)<br>
**Tracking Issue:** https://github.com/kubernetes/kubernetes/issues/92914

## Urgent Upgrade Notes 

### (No, really, you MUST read this before you upgrade)

- Remove storage metrics `storage_operation_errors_total`, since we already have `storage_operation_status_count`.And add new field `status` for `storage_operation_duration_seconds`, so that we can know about all status storage operation latency. ([#98332](https://github.com/kubernetes/kubernetes/pull/98332), [@JornShen](https://github.com/JornShen)) [SIG Instrumentation and Storage]
 
## Changes by Kind

### Deprecation

- Remove the TokenRequest and TokenRequestProjection feature gates ([#97148](https://github.com/kubernetes/kubernetes/pull/97148), [@wawa0210](https://github.com/wawa0210)) [SIG Node]
- Removing experimental windows container hyper-v support with Docker ([#97141](https://github.com/kubernetes/kubernetes/pull/97141), [@wawa0210](https://github.com/wawa0210)) [SIG Node and Windows]
- The `export` query parameter (inconsistently supported by API resources and deprecated in v1.14) is fully removed.  Requests setting this query parameter will now receive a 400 status response. ([#98312](https://github.com/kubernetes/kubernetes/pull/98312), [@deads2k](https://github.com/deads2k)) [SIG API Machinery, Auth and Testing]

### API Change

- Enable SPDY pings to keep connections alive, so that `kubectl exec` and `kubectl portforward` won't be interrupted. ([#97083](https://github.com/kubernetes/kubernetes/pull/97083), [@knight42](https://github.com/knight42)) [SIG API Machinery and CLI]

### Documentation

- Official support to build kubernetes with docker-machine / remote docker is removed. This change does not affect building kubernetes with docker locally. ([#97935](https://github.com/kubernetes/kubernetes/pull/97935), [@adeniyistephen](https://github.com/adeniyistephen)) [SIG Release and Testing]
- Set kubelet option `--volume-stats-agg-period` to negative value to disable volume calculations. ([#96675](https://github.com/kubernetes/kubernetes/pull/96675), [@pacoxu](https://github.com/pacoxu)) [SIG Node]

### Bug or Regression

- Clean ReplicaSet by revision instead of creation timestamp in deployment controller ([#97407](https://github.com/kubernetes/kubernetes/pull/97407), [@waynepeking348](https://github.com/waynepeking348)) [SIG Apps]
- Ensure that client-go's EventBroadcaster is safe (non-racy) during shutdown. ([#95664](https://github.com/kubernetes/kubernetes/pull/95664), [@DirectXMan12](https://github.com/DirectXMan12)) [SIG API Machinery]
- Fix azure file migration issue ([#97877](https://github.com/kubernetes/kubernetes/pull/97877), [@andyzhangx](https://github.com/andyzhangx)) [SIG Auth, Cloud Provider and Storage]
- Fix kubelet from panic after getting the wrong signal ([#98200](https://github.com/kubernetes/kubernetes/pull/98200), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fix repeatedly aquire the inhibit lock ([#98088](https://github.com/kubernetes/kubernetes/pull/98088), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fixed a bug that the kubelet cannot start on BtrfS. ([#98042](https://github.com/kubernetes/kubernetes/pull/98042), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Fixed an issue with garbage collection failing to clean up namespaced children of an object also referenced incorrectly by cluster-scoped children ([#98068](https://github.com/kubernetes/kubernetes/pull/98068), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Apps]
- Fixed no effect namespace when exposing deployment with --dry-run=client. ([#97492](https://github.com/kubernetes/kubernetes/pull/97492), [@masap](https://github.com/masap)) [SIG CLI]
- Fixing a bug where a failed node may not have the NoExecute taint set correctly ([#96876](https://github.com/kubernetes/kubernetes/pull/96876), [@howieyuen](https://github.com/howieyuen)) [SIG Apps and Node]
- Indentation of `Resource Quota` block in kubectl describe namespaces output gets correct. ([#97946](https://github.com/kubernetes/kubernetes/pull/97946), [@dty1er](https://github.com/dty1er)) [SIG CLI]
- KUBECTL_EXTERNAL_DIFF now accepts equal sign for additional parameters. ([#98158](https://github.com/kubernetes/kubernetes/pull/98158), [@dougsland](https://github.com/dougsland)) [SIG CLI]
- Kubeadm: fix a bug where "kubeadm join" would not properly handle missing names for existing etcd members. ([#97372](https://github.com/kubernetes/kubernetes/pull/97372), [@ihgann](https://github.com/ihgann)) [SIG Cluster Lifecycle]
- Kubelet should ignore cgroup driver check on Windows node. ([#97764](https://github.com/kubernetes/kubernetes/pull/97764), [@pacoxu](https://github.com/pacoxu)) [SIG Node and Windows]
- Make podTopologyHints protected by lock ([#95111](https://github.com/kubernetes/kubernetes/pull/95111), [@choury](https://github.com/choury)) [SIG Node]
- Readjust kubelet_containers_per_pod_count bucket ([#98169](https://github.com/kubernetes/kubernetes/pull/98169), [@wawa0210](https://github.com/wawa0210)) [SIG Instrumentation and Node]
- Scores from InterPodAffinity have stronger differentiation. ([#98096](https://github.com/kubernetes/kubernetes/pull/98096), [@leileiwan](https://github.com/leileiwan)) [SIG Scheduling]
- Specifying the KUBE_TEST_REPO environment variable when e2e tests are executed will instruct the test infrastructure to load that image from a location within the specified repo, using a predefined pattern. ([#93510](https://github.com/kubernetes/kubernetes/pull/93510), [@smarterclayton](https://github.com/smarterclayton)) [SIG Testing]
- Static pods will be deleted gracefully. ([#98103](https://github.com/kubernetes/kubernetes/pull/98103), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Use network.Interface.VirtualMachine.ID to get the binded VM
  Skip standalone VM when reconciling LoadBalancer ([#97635](https://github.com/kubernetes/kubernetes/pull/97635), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]

### Other (Cleanup or Flake)

- Kubeadm: change the default image repository for CI images from 'gcr.io/kubernetes-ci-images' to 'gcr.io/k8s-staging-ci-images' ([#97087](https://github.com/kubernetes/kubernetes/pull/97087), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Migrate generic_scheduler.go and types.go to structured logging. ([#98134](https://github.com/kubernetes/kubernetes/pull/98134), [@tanjing2020](https://github.com/tanjing2020)) [SIG Scheduling]
- Migrate proxy/winuserspace/proxier.go logs to structured logging ([#97941](https://github.com/kubernetes/kubernetes/pull/97941), [@JornShen](https://github.com/JornShen)) [SIG Network]
- Migrate staging/src/k8s.io/apiserver/pkg/audit/policy/reader.go logs to structured logging. ([#98252](https://github.com/kubernetes/kubernetes/pull/98252), [@lala123912](https://github.com/lala123912)) [SIG API Machinery and Auth]
- Migrate staging\src\k8s.io\apiserver\pkg\endpoints logs to  structured logging ([#98093](https://github.com/kubernetes/kubernetes/pull/98093), [@lala123912](https://github.com/lala123912)) [SIG API Machinery]
- The kubectl alpha debug command was scheduled to be removed in v1.21. ([#98111](https://github.com/kubernetes/kubernetes/pull/98111), [@pandaamanda](https://github.com/pandaamanda)) [SIG CLI]
- Update cri-tools to [v1.20.0](https://github.com/kubernetes-sigs/cri-tools/releases/tag/v1.20.0) ([#97967](https://github.com/kubernetes/kubernetes/pull/97967), [@rajibmitra](https://github.com/rajibmitra)) [SIG Cloud Provider]
- Windows nodes on GCE will take longer to start due to dependencies installed at node creation time. ([#98284](https://github.com/kubernetes/kubernetes/pull/98284), [@pjh](https://github.com/pjh)) [SIG Cloud Provider]