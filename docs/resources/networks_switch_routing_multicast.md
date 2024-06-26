---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_routing_multicast Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_switch_routing_multicast (Resource)



## Example Usage

```terraform
resource "meraki_networks_switch_routing_multicast" "example" {

  default_settings = {

    flood_unknown_multicast_traffic_enabled = true
    igmp_snooping_enabled                   = true
  }
  network_id = "string"
  overrides = [{

    flood_unknown_multicast_traffic_enabled = true
    igmp_snooping_enabled                   = true
    switches                                = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }]
}

output "meraki_networks_switch_routing_multicast_example" {
  value = meraki_networks_switch_routing_multicast.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `default_settings` (Attributes) Default multicast setting for entire network. IGMP snooping and Flood unknown multicast traffic settings are enabled by default. (see [below for nested schema](#nestedatt--default_settings))
- `overrides` (Attributes Set) Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings. (see [below for nested schema](#nestedatt--overrides))

<a id="nestedatt--default_settings"></a>
### Nested Schema for `default_settings`

Optional:

- `flood_unknown_multicast_traffic_enabled` (Boolean) Flood unknown multicast traffic setting for entire network
- `igmp_snooping_enabled` (Boolean) IGMP snooping setting for entire network


<a id="nestedatt--overrides"></a>
### Nested Schema for `overrides`

Optional:

- `flood_unknown_multicast_traffic_enabled` (Boolean) Flood unknown multicast traffic setting for switches, switch stacks or switch profiles
- `igmp_snooping_enabled` (Boolean) IGMP snooping setting for switches, switch stacks or switch profiles
- `stacks` (List of String) List of switch stack ids for non-template network
- `switch_profiles` (List of String) List of switch profiles ids for template network
- `switches` (List of String) List of switch serials for non-template network

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_switch_routing_multicast.example "network_id"
```
