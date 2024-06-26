---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_snmp Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_snmp (Resource)



## Example Usage

```terraform
resource "meraki_organizations_snmp" "example" {

  organization_id = "string"
  peer_ips        = ["123.123.123.1"]
  v2c_enabled     = false
  v3_auth_mode    = "SHA"
  v3_enabled      = true
  v3_priv_mode    = "AES128"
}

output "meraki_organizations_snmp_example" {
  value = meraki_organizations_snmp.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `peer_ips` (List of String) The list of IPv4 addresses that are allowed to access the SNMP server.
- `v2c_enabled` (Boolean) Boolean indicating whether SNMP version 2c is enabled for the organization.
- `v3_auth_mode` (String) The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
- `v3_auth_pass` (String) The SNMP version 3 authentication password. Must be at least 8 characters if specified.
- `v3_enabled` (Boolean) Boolean indicating whether SNMP version 3 is enabled for the organization.
- `v3_priv_mode` (String) The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
- `v3_priv_pass` (String) The SNMP version 3 privacy password. Must be at least 8 characters if specified.

### Read-Only

- `hostname` (String)
- `port` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_snmp.example "organization_id"
```
