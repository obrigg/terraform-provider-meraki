---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_sm_devices_device_command_logs Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_sm_devices_device_command_logs (Data Source)



## Example Usage

```terraform
data "meraki_networks_sm_devices_device_command_logs" "example" {

  device_id      = "string"
  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_sm_devices_device_command_logs_example" {
  value = data.meraki_networks_sm_devices_device_command_logs.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_id` (String) deviceId path parameter. Device ID
- `network_id` (String) networkId path parameter. Network ID

### Optional

- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

### Read-Only

- `items` (Attributes List) Array of ResponseSmGetNetworkSmDeviceDeviceCommandLogs (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `action` (String) The type of command sent to the device.
- `dashboard_user` (String) The Meraki dashboard user who initiated the command.
- `details` (String) A JSON string object containing command details.
- `name` (String) The name of the device to which the command is sent.
- `ts` (String) The time the command was sent to the device.