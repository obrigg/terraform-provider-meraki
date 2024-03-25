---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_appliance_uplinks_settings Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_appliance_uplinks_settings (Resource)



## Example Usage

```terraform
resource "meraki_devices_appliance_uplinks_settings" "example" {

  interfaces = {

    wan1 = {

      enabled = true
      pppoe = {

        authentication = {

          enabled  = true
          password = "password"
          username = "username"
        }
        enabled = true
      }
      svis = {

        ipv4 = {

          address         = "9.10.11.10/16"
          assignment_mode = "static"
          gateway         = "13.14.15.16"
          nameservers = {

            addresses = ["1.2.3.4"]
          }
        }
        ipv6 = {

          address         = "1:2:3::4"
          assignment_mode = "static"
          gateway         = "1:2:3::5"
          nameservers = {

            addresses = ["1001:4860:4860::8888", "1001:4860:4860::8844"]
          }
        }
      }
      vlan_tagging = {

        enabled = true
        vlan_id = 1
      }
    }
    wan2 = {

      enabled = true
      pppoe = {

        authentication = {

          enabled  = true
          password = "password"
          username = "username"
        }
        enabled = true
      }
      svis = {

        ipv4 = {

          address         = "9.10.11.10/16"
          assignment_mode = "static"
          gateway         = "13.14.15.16"
          nameservers = {

            addresses = ["1.2.3.4"]
          }
        }
        ipv6 = {

          address         = "1:2:3::4"
          assignment_mode = "static"
          gateway         = "1:2:3::5"
          nameservers = {

            addresses = ["1001:4860:4860::8888", "1001:4860:4860::8844"]
          }
        }
      }
      vlan_tagging = {

        enabled = true
        vlan_id = 1
      }
    }
  }
  serial = "string"
}

output "meraki_devices_appliance_uplinks_settings_example" {
  value = meraki_devices_appliance_uplinks_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Optional

- `interfaces` (Attributes) Interface settings. (see [below for nested schema](#nestedatt--interfaces))

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Optional:

- `wan1` (Attributes) WAN 1 settings. (see [below for nested schema](#nestedatt--interfaces--wan1))
- `wan2` (Attributes) WAN 2 settings. (see [below for nested schema](#nestedatt--interfaces--wan2))

<a id="nestedatt--interfaces--wan1"></a>
### Nested Schema for `interfaces.wan1`

Optional:

- `enabled` (Boolean) Enable or disable the interface.
- `pppoe` (Attributes) Configuration options for PPPoE. (see [below for nested schema](#nestedatt--interfaces--wan1--pppoe))
- `svis` (Attributes) SVI settings by protocol. (see [below for nested schema](#nestedatt--interfaces--wan1--svis))
- `vlan_tagging` (Attributes) VLAN tagging settings. (see [below for nested schema](#nestedatt--interfaces--wan1--vlan_tagging))

<a id="nestedatt--interfaces--wan1--pppoe"></a>
### Nested Schema for `interfaces.wan1.pppoe`

Optional:

- `authentication` (Attributes) Settings for PPPoE Authentication. (see [below for nested schema](#nestedatt--interfaces--wan1--pppoe--authentication))
- `enabled` (Boolean) Whether PPPoE is enabled.

<a id="nestedatt--interfaces--wan1--pppoe--authentication"></a>
### Nested Schema for `interfaces.wan1.pppoe.enabled`

Optional:

- `enabled` (Boolean) Whether PPPoE authentication is enabled.
- `password` (String, Sensitive) Password for PPPoE authentication. This parameter is not returned.
- `username` (String) Username for PPPoE authentication.



<a id="nestedatt--interfaces--wan1--svis"></a>
### Nested Schema for `interfaces.wan1.svis`

Optional:

- `ipv4` (Attributes) IPv4 settings for static/dynamic mode. (see [below for nested schema](#nestedatt--interfaces--wan1--svis--ipv4))
- `ipv6` (Attributes) IPv6 settings for static/dynamic mode. (see [below for nested schema](#nestedatt--interfaces--wan1--svis--ipv6))

<a id="nestedatt--interfaces--wan1--svis--ipv4"></a>
### Nested Schema for `interfaces.wan1.svis.ipv6`

Optional:

- `address` (String) IP address and subnet mask when in static mode.
- `assignment_mode` (String) The assignment mode for this SVI. Applies only when PPPoE is disabled.
- `gateway` (String) Gateway IP address when in static mode.
- `nameservers` (Attributes) The nameserver settings for this SVI. (see [below for nested schema](#nestedatt--interfaces--wan1--svis--ipv6--nameservers))

<a id="nestedatt--interfaces--wan1--svis--ipv6--nameservers"></a>
### Nested Schema for `interfaces.wan1.svis.ipv6.nameservers`

Optional:

- `addresses` (List of String) Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.



<a id="nestedatt--interfaces--wan1--svis--ipv6"></a>
### Nested Schema for `interfaces.wan1.svis.ipv6`

Optional:

- `address` (String) Static address that will override the one(s) received by SLAAC.
- `assignment_mode` (String) The assignment mode for this SVI. Applies only when PPPoE is disabled.
- `gateway` (String) Static gateway that will override the one received by autoconf.
- `nameservers` (Attributes) The nameserver settings for this SVI. (see [below for nested schema](#nestedatt--interfaces--wan1--svis--ipv6--nameservers))

<a id="nestedatt--interfaces--wan1--svis--ipv6--nameservers"></a>
### Nested Schema for `interfaces.wan1.svis.ipv6.nameservers`

Optional:

- `addresses` (List of String) Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.




<a id="nestedatt--interfaces--wan1--vlan_tagging"></a>
### Nested Schema for `interfaces.wan1.vlan_tagging`

Optional:

- `enabled` (Boolean) Whether VLAN tagging is enabled.
- `vlan_id` (Number) The ID of the VLAN to use for VLAN tagging.



<a id="nestedatt--interfaces--wan2"></a>
### Nested Schema for `interfaces.wan2`

Optional:

- `enabled` (Boolean) Enable or disable the interface.
- `pppoe` (Attributes) Configuration options for PPPoE. (see [below for nested schema](#nestedatt--interfaces--wan2--pppoe))
- `svis` (Attributes) SVI settings by protocol. (see [below for nested schema](#nestedatt--interfaces--wan2--svis))
- `vlan_tagging` (Attributes) VLAN tagging settings. (see [below for nested schema](#nestedatt--interfaces--wan2--vlan_tagging))

<a id="nestedatt--interfaces--wan2--pppoe"></a>
### Nested Schema for `interfaces.wan2.pppoe`

Optional:

- `authentication` (Attributes) Settings for PPPoE Authentication. (see [below for nested schema](#nestedatt--interfaces--wan2--pppoe--authentication))
- `enabled` (Boolean) Whether PPPoE is enabled.

<a id="nestedatt--interfaces--wan2--pppoe--authentication"></a>
### Nested Schema for `interfaces.wan2.pppoe.enabled`

Optional:

- `enabled` (Boolean) Whether PPPoE authentication is enabled.
- `password` (String, Sensitive) Password for PPPoE authentication. This parameter is not returned.
- `username` (String) Username for PPPoE authentication.



<a id="nestedatt--interfaces--wan2--svis"></a>
### Nested Schema for `interfaces.wan2.svis`

Optional:

- `ipv4` (Attributes) IPv4 settings for static/dynamic mode. (see [below for nested schema](#nestedatt--interfaces--wan2--svis--ipv4))
- `ipv6` (Attributes) IPv6 settings for static/dynamic mode. (see [below for nested schema](#nestedatt--interfaces--wan2--svis--ipv6))

<a id="nestedatt--interfaces--wan2--svis--ipv4"></a>
### Nested Schema for `interfaces.wan2.svis.ipv6`

Optional:

- `address` (String) IP address and subnet mask when in static mode.
- `assignment_mode` (String) The assignment mode for this SVI. Applies only when PPPoE is disabled.
- `gateway` (String) Gateway IP address when in static mode.
- `nameservers` (Attributes) The nameserver settings for this SVI. (see [below for nested schema](#nestedatt--interfaces--wan2--svis--ipv6--nameservers))

<a id="nestedatt--interfaces--wan2--svis--ipv6--nameservers"></a>
### Nested Schema for `interfaces.wan2.svis.ipv6.nameservers`

Optional:

- `addresses` (List of String) Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.



<a id="nestedatt--interfaces--wan2--svis--ipv6"></a>
### Nested Schema for `interfaces.wan2.svis.ipv6`

Optional:

- `address` (String) Static address that will override the one(s) received by SLAAC.
- `assignment_mode` (String) The assignment mode for this SVI. Applies only when PPPoE is disabled.
- `gateway` (String) Static gateway that will override the one received by autoconf.
- `nameservers` (Attributes) The nameserver settings for this SVI. (see [below for nested schema](#nestedatt--interfaces--wan2--svis--ipv6--nameservers))

<a id="nestedatt--interfaces--wan2--svis--ipv6--nameservers"></a>
### Nested Schema for `interfaces.wan2.svis.ipv6.nameservers`

Optional:

- `addresses` (List of String) Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.




<a id="nestedatt--interfaces--wan2--vlan_tagging"></a>
### Nested Schema for `interfaces.wan2.vlan_tagging`

Optional:

- `enabled` (Boolean) Whether VLAN tagging is enabled.
- `vlan_id` (Number) The ID of the VLAN to use for VLAN tagging.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_devices_appliance_uplinks_settings.example "serial"
```