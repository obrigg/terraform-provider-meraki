
data "meraki_networks_traffic_shaping_application_categories" "example" {

  network_id = "string"
}

output "meraki_networks_traffic_shaping_application_categories_example" {
  value = data.meraki_networks_traffic_shaping_application_categories.example.item
}
