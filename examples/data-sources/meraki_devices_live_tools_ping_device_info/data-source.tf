
data "meraki_devices_live_tools_ping_device_info" "example" {

  id     = "string"
  serial = "string"
}

output "meraki_devices_live_tools_ping_device_info_example" {
  value = data.meraki_devices_live_tools_ping_device_info.example.item
}
