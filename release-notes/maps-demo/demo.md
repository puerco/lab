
## 1. Talk about overriding data in release notes using maps

### 1.1 Screen demo

1. Render release notes for 1.18.1
go run cmd/release-notes/main.go  --start-rev v1.18.0 --end-rev v1.18.1  --branch release-1.18 --output /tmp/notes.md --dependencies=false

2. Note that first release note says "lb"
3. Create a map to change the text to 
   Text: "Azure: fix concurreny issue in load balancer creation"

4. Note note for PR #89796
   "Reduce event spam during a volume operation error."
   which is CVE in disguise
5. Create Map for PR #89796 changing note to 
   Text: Implement changes to mitigate CVE-2020-8555"
   Author: @kubernetes/product-security-committee
6. Re-render release notes for 1.18


## 2. Talk about adding data to release notes

### 2.1 Screen demo

1. Show and explain CVE map using the YAML example

2. Apply CVE example and re-render notes

3. Show updated notes

## 3. Talk about importing maps from different sources

 - Maps from release notes team fixes

 - Maps from the security commitee, possible from a bucket or repo





