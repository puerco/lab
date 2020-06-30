
## 1. Talk about overriding data in release notes using maps

### 1.1 Screen demo → Modifying/overriding RN data

1. Render release notes for 1.18.1
Should produce vanilla notes like [these](output/notes01.md).

2. Note that first release note says "_lb_"

3. Create a map to change the text to 
   Text: "Azure: fix concurreny issue in load balancer creation"
   Raw map is [here](map01/map01.yaml)

4. Highlight note for PR #89796
   "Reduce event spam during a volume operation error."
   which is CVE in disguise. 

5. Create Map for PR #89796 changing note to 
   Text: Implement changes to mitigate CVE-2020-8555"
   Author: @kubernetes/product-security-committee

   Raw map is [here](map02/map02.yaml)

6. Re-render release notes for 1.18

   Should render like [this version](output/notes02.md) (only note item changed)


## 2. Talk about adding data to release notes

### 2.1 Screen demo → Adding Data (CVE example)

1. Show and explain CVE map using the YAML example

   CVE YAML example [is here](map04/map04.yaml)

2. Apply CVE example and re-render notes

3. Show updated notes

   Should look like [this final version with CVE data](output/notes04.md)

## 3. Finish talk about importing maps from different sources

 - Maps from release notes team fixes

 - Maps from the security commitee, possible from a bucket or repo





