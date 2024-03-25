---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_stacks_routing_interfaces_dhcp Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_switch_stacks_routing_interfaces_dhcp (Resource)



## Example Usage

```terraform
resource "meraki_networks_switch_stacks_routing_interfaces_dhcp" "example" {

  boot_file_name       = "home_boot_file"
  boot_next_server     = "1.2.3.4"
  boot_options_enabled = true
  dhcp_lease_time      = "1 day"
  dhcp_mode            = "dhcpServer"
  dhcp_options = [{

    code  = "5"
    type  = "text"
    value = "five"
  }]
  dns_custom_nameservers = ["8.8.8.8, 8.8.4.4"]
  dns_nameservers_option = "custom"
  fixed_ip_assignments = [{

    ip   = "192.168.1.12"
    mac  = "22:33:44:55:66:77"
    name = "Cisco Meraki valued client"
  }]
  interface_id = "string"
  network_id   = "string"
  reserved_ip_ranges = [{

    comment = "A reserved IP range"
    end     = "192.168.1.10"
    start   = "192.168.1.1"
  }]
  switch_stack_id = "string"
}

output "meraki_networks_switch_stacks_routing_interfaces_dhcp_example" {
  value = meraki_networks_switch_stacks_routing_interfaces_dhcp.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_id` (String) interfaceId path parameter. Interface ID
- `network_id` (String) networkId path parameter. Network ID
- `switch_stack_id` (String) switchStackId path parameter. Switch stack ID

### Optional

- `boot_file_name` (String) The PXE boot server file name for the DHCP server running on the switch stack interface
- `boot_next_server` (String) The PXE boot server IP for the DHCP server running on the switch stack interface
- `boot_options_enabled` (Boolean) Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface
- `dhcp_lease_time` (String) The DHCP lease time config for the dhcp server running on switch stack interface ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
- `dhcp_mode` (String) The DHCP mode options for the switch stack interface ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
- `dhcp_options` (Attributes Set) Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface (see [below for nested schema](#nestedatt--dhcp_options))
- `dhcp_relay_server_ips` (List of String) The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
- `dns_custom_nameservers` (List of String) The DHCP name server IPs when DHCP name server option is 'custom'
- `dns_nameservers_option` (String) The DHCP name server option for the dhcp server running on the switch stack interface ('googlePublicDns', 'openDns' or 'custom')
- `fixed_ip_assignments` (Attributes Set) Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface (see [below for nested schema](#nestedatt--fixed_ip_assignments))
- `reserved_ip_ranges` (Attributes Set) Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface (see [below for nested schema](#nestedatt--reserved_ip_ranges))

<a id="nestedatt--dhcp_options"></a>
### Nested Schema for `dhcp_options`

Optional:

- `code` (String) The code for DHCP option which should be from 2 to 254
- `type` (String) The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
- `value` (String) The value of the DHCP option


<a id="nestedatt--fixed_ip_assignments"></a>
### Nested Schema for `fixed_ip_assignments`

Optional:

- `ip` (String) The IP address of the client which has fixed IP address assigned to it
- `mac` (String) The MAC address of the client which has fixed IP address
- `name` (String) The name of the client which has fixed IP address


<a id="nestedatt--reserved_ip_ranges"></a>
### Nested Schema for `reserved_ip_ranges`

Optional:

- `comment` (String) The comment for the reserved IP range
- `end` (String) The ending IP address of the reserved IP range
- `start` (String) The starting IP address of the reserved IP range

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_switch_stacks_routing_interfaces_dhcp.example "interface_id,network_id,switch_stack_id"
```