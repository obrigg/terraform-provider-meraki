
data "meraki_organizations_devices_provisioning_statuses" "example" {

  ending_before    = "string"
  network_ids      = ["string"]
  organization_id  = "string"
  per_page         = 1
  product_types    = ["string"]
  serials          = ["string"]
  starting_after   = "string"
  status           = "string"
  tags             = ["string"]
  tags_filter_type = "string"
}

output "meraki_organizations_devices_provisioning_statuses_example" {
  value = data.meraki_organizations_devices_provisioning_statuses.example.items
}
