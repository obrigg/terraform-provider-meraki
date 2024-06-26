---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_firewall_l7_firewall_rules Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_appliance_firewall_l7_firewall_rules (Resource)



## Example Usage
### 1
```terraform
resource "meraki_networks_appliance_firewall_l7_firewall_rules" "example" {

  network_id = "string"
  rules = [{

    policy = "deny"
    type   = "host"
    value  = "google.com"
  }]
}

output "meraki_networks_appliance_firewall_l7_firewall_rules_example" {
  value = meraki_networks_appliance_firewall_l7_firewall_rules.example
}
```

### 2
```terraform
resource "meraki_networks_appliance_firewall_l7_firewall_rules" "example" {

  network_id = "string"
  rules = [{

    policy = "deny"
    type   = "applicationCategory"
    value_obj = {
      name = "Sports"
      id = "meraki:layer7/category/5"
    }
  }]
}

output "meraki_networks_appliance_firewall_l7_firewall_rules_example" {
  value = meraki_networks_appliance_firewall_l7_firewall_rules.example
}
```

### 3
```terraform
resource "meraki_networks_appliance_firewall_l7_firewall_rules" "example" {

  network_id = "string"
  rules = [{
    policy = "deny"
    type = "blockedCountries"
    value_list = ["IT", "IL", "US"]
  }]
}

output "meraki_networks_appliance_firewall_l7_firewall_rules_example" {
  value = meraki_networks_appliance_firewall_l7_firewall_rules.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `rules` (Attributes Set) An ordered array of the MX L7 firewall rules (see [below for nested schema](#nestedatt--rules))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `policy` (String) 'Deny' traffic specified by this rule
- `type` (String) Type of the L7 rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
- `value` (String) The 'value' of what you want to block. Format of 'value' varies depending on type of the rule. Send a string to request.
- `value_list` (List of String) The 'value_list' of what you want to block. Send a lis of string in request.
- `value_obj` (Attributes) The 'value_obj' of what you want to block. Send a dict in request.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_appliance_firewall_l7_firewall_rules.example "network_id"
```
