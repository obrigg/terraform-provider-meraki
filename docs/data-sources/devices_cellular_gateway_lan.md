---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_cellular_gateway_lan Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_cellular_gateway_lan (Data Source)



## Example Usage

```terraform
data "meraki_devices_cellular_gateway_lan" "example" {

  serial = "string"
}

output "meraki_devices_cellular_gateway_lan_example" {
  value = data.meraki_devices_cellular_gateway_lan.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `device_lan_ip` (String)
- `device_name` (String)
- `device_subnet` (String)
- `fixed_ip_assignments` (Attributes Set) (see [below for nested schema](#nestedatt--item--fixed_ip_assignments))
- `reserved_ip_ranges` (Attributes Set) (see [below for nested schema](#nestedatt--item--reserved_ip_ranges))

<a id="nestedatt--item--fixed_ip_assignments"></a>
### Nested Schema for `item.fixed_ip_assignments`

Read-Only:

- `ip` (String)
- `mac` (String)
- `name` (String)


<a id="nestedatt--item--reserved_ip_ranges"></a>
### Nested Schema for `item.reserved_ip_ranges`

Read-Only:

- `comment` (String)
- `end` (String)
- `start` (String)