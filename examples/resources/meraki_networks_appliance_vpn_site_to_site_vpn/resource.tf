
resource "meraki_networks_appliance_vpn_site_to_site_vpn" "example" {

  hubs = [{

    hub_id            = "N_4901849"
    use_default_route = true
  }]
  mode       = "spoke"
  network_id = "string"
  subnets = [{

    local_subnet = "192.168.1.0/24"
    use_vpn      = true
  }]
}

output "meraki_networks_appliance_vpn_site_to_site_vpn_example" {
  value = meraki_networks_appliance_vpn_site_to_site_vpn.example
}