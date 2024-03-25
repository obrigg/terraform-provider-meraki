---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_appliance_security_intrusion Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_appliance_security_intrusion (Resource)



## Example Usage

```terraform
resource "meraki_organizations_appliance_security_intrusion" "example" {

  allowed_rules = [{

    message = "SQL sa login failed"
    rule_id = "meraki:intrusion/snort/GID/01/SID/688"
  }]
  organization_id = "string"
}

output "meraki_organizations_appliance_security_intrusion_example" {
  value = meraki_organizations_appliance_security_intrusion.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `allowed_rules` (Attributes Set) Sets a list of specific SNORT signatures to allow (see [below for nested schema](#nestedatt--allowed_rules))

<a id="nestedatt--allowed_rules"></a>
### Nested Schema for `allowed_rules`

Optional:

- `message` (String) Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET
- `rule_id` (String) A rule identifier of the format meraki:intrusion/snort/GID/<gid>/SID/<sid>. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_appliance_security_intrusion.example "organization_id"
```