---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_ssids Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_appliance_ssids (Resource)



## Example Usage

```terraform
resource "meraki_networks_appliance_ssids" "example" {

  auth_mode       = "8021x-radius"
  default_vlan_id = 1
  dhcp_enforced_deauthentication = {

    enabled = true
  }
  enabled         = true
  encryption_mode = "wpa"
  name            = "My SSID"
  network_id      = "string"
  number          = "string"
  psk             = "psk"
  radius_servers = [{

    host   = "0.0.0.0"
    port   = 1000
    secret = "secret"
  }]
  visible             = true
  wpa_encryption_mode = "WPA2 only"
}

output "meraki_networks_appliance_ssids_example" {
  value = meraki_networks_appliance_ssids.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `number` (Number) The number of the SSID.

### Optional

- `auth_mode` (String) The association control method for the SSID.
- `default_vlan_id` (Number) The VLAN ID of the VLAN associated to this SSID.
- `dhcp_enforced_deauthentication` (Attributes) DHCP Enforced Deauthentication enables the disassociation of wireless clients in addition to Mandatory DHCP. This param is only valid on firmware versions >= MX 17.0 where the associated LAN has Mandatory DHCP Enabled (see [below for nested schema](#nestedatt--dhcp_enforced_deauthentication))
- `enabled` (Boolean) Whether or not the SSID is enabled.
- `encryption_mode` (String) The psk encryption mode for the SSID.
- `name` (String) The name of the SSID.
- `psk` (String) The passkey for the SSID. This param is only valid if the authMode is 'psk'.
- `radius_servers` (Attributes Set) The RADIUS 802.1x servers to be used for authentication. (see [below for nested schema](#nestedatt--radius_servers))
- `visible` (Boolean) Boolean indicating whether the MX should advertise or hide this SSID.
- `wpa_encryption_mode` (String) WPA encryption mode for the SSID.

<a id="nestedatt--dhcp_enforced_deauthentication"></a>
### Nested Schema for `dhcp_enforced_deauthentication`

Optional:

- `enabled` (Boolean) Enable DCHP Enforced Deauthentication on the SSID.


<a id="nestedatt--radius_servers"></a>
### Nested Schema for `radius_servers`

Optional:

- `host` (String) The IP address of your RADIUS server.
- `port` (Number) The UDP port your RADIUS servers listens on for Access-requests.
- `secret` (String) The RADIUS client shared secret.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_appliance_ssids.example "network_id,number"
```