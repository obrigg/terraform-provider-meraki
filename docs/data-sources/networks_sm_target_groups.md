---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_sm_target_groups Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_sm_target_groups (Data Source)



## Example Usage

```terraform
data "meraki_networks_sm_target_groups" "example" {

  network_id   = "string"
  with_details = false
}

output "meraki_networks_sm_target_groups_example" {
  value = data.meraki_networks_sm_target_groups.example.items
}

data "meraki_networks_sm_target_groups" "example" {

  network_id   = "string"
  with_details = false
}

output "meraki_networks_sm_target_groups_example" {
  value = data.meraki_networks_sm_target_groups.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `network_id` (String) networkId path parameter. Network ID
- `target_group_id` (String) targetGroupId path parameter. Target group ID
- `with_details` (Boolean) withDetails query parameter. Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseSmGetNetworkSmTargetGroups (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `name` (String)
- `scope` (String)
- `tags` (String)
- `type` (String)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `name` (String)
- `scope` (String)
- `tags` (String)
- `type` (String)
