---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_dscp_to_cos_mappings Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_switch_dscp_to_cos_mappings (Data Source)



## Example Usage

```terraform
data "meraki_networks_switch_dscp_to_cos_mappings" "example" {

  network_id = "string"
}

output "meraki_networks_switch_dscp_to_cos_mappings_example" {
  value = data.meraki_networks_switch_dscp_to_cos_mappings.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `mappings` (Attributes Set) (see [below for nested schema](#nestedatt--item--mappings))

<a id="nestedatt--item--mappings"></a>
### Nested Schema for `item.mappings`

Read-Only:

- `cos` (Number)
- `dscp` (Number)
- `title` (String)
