
data "meraki_organizations_wireless_devices_ethernet_statuses" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_wireless_devices_ethernet_statuses_example" {
  value = data.meraki_organizations_wireless_devices_ethernet_statuses.example.items
}
