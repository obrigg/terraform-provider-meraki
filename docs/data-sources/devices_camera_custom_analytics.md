---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_camera_custom_analytics Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_camera_custom_analytics (Data Source)



## Example Usage

```terraform
data "meraki_devices_camera_custom_analytics" "example" {

  serial = "string"
}

output "meraki_devices_camera_custom_analytics_example" {
  value = data.meraki_devices_camera_custom_analytics.example.item
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

- `artifact_id` (String)
- `enabled` (Boolean)
- `parameters` (Attributes Set) (see [below for nested schema](#nestedatt--item--parameters))

<a id="nestedatt--item--parameters"></a>
### Nested Schema for `item.parameters`

Read-Only:

- `name` (String)
- `value` (Number)
