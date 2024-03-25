---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_alternate_management_interface Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_switch_alternate_management_interface (Data Source)



## Example Usage

```terraform
data "meraki_networks_switch_alternate_management_interface" "example" {

  network_id = "string"
}

output "meraki_networks_switch_alternate_management_interface_example" {
  value = data.meraki_networks_switch_alternate_management_interface.example.item
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

- `enabled` (Boolean)
- `protocols` (List of String)
- `switches` (Attributes Set) (see [below for nested schema](#nestedatt--item--switches))
- `vlan_id` (Number)

<a id="nestedatt--item--switches"></a>
### Nested Schema for `item.switches`

Read-Only:

- `alternate_management_ip` (String)
- `gateway` (String)
- `serial` (String)
- `subnet_mask` (String)