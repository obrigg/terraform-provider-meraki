---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_sm_target_groups Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_sm_target_groups (Resource)



## Example Usage

```terraform
resource "meraki_networks_sm_target_groups" "example" {

  name       = "My target group"
  network_id = "string"
  scope      = "none"
}

output "meraki_networks_sm_target_groups_example" {
  value = meraki_networks_sm_target_groups.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `name` (String) The name of this target group
- `scope` (String) The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty.
- `target_group_id` (String) targetGroupId path parameter. Target group ID

### Read-Only

- `tags` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_sm_target_groups.example "network_id,target_group_id"
```
