---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_camera_video_settings Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_camera_video_settings (Data Source)



## Example Usage

```terraform
data "meraki_devices_camera_video_settings" "example" {

  serial = "string"
}

output "meraki_devices_camera_video_settings_example" {
  value = data.meraki_devices_camera_video_settings.example.item
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

- `external_rtsp_enabled` (Boolean)
- `rtsp_url` (String)