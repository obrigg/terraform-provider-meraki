---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_camera_video_settings Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_camera_video_settings (Resource)



## Example Usage

```terraform
resource "meraki_devices_camera_video_settings" "example" {

  external_rtsp_enabled = true
  serial                = "string"
}

output "meraki_devices_camera_video_settings_example" {
  value = meraki_devices_camera_video_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Optional

- `external_rtsp_enabled` (Boolean) Boolean indicating if external rtsp stream is exposed

### Read-Only

- `rtsp_url` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_devices_camera_video_settings.example "serial"
```