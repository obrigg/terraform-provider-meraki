
data "meraki_organizations_api_requests_overview" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_api_requests_overview_example" {
  value = data.meraki_organizations_api_requests_overview.example.item
}
