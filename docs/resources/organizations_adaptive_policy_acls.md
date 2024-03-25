---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_adaptive_policy_acls Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_adaptive_policy_acls (Resource)



## Example Usage

```terraform
resource "meraki_organizations_adaptive_policy_acls" "example" {

  description     = "Blocks sensitive web traffic"
  ip_version      = "ipv6"
  name            = "Block sensitive web traffic"
  organization_id = "string"
  rules = [{

    dst_port = "22-30"
    policy   = "deny"
    protocol = "tcp"
    src_port = "1,33"
  }]
}

output "meraki_organizations_adaptive_policy_acls_example" {
  value = meraki_organizations_adaptive_policy_acls.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `acl_id` (String) ID of the adaptive policy ACL
- `description` (String) Description of the adaptive policy ACL
- `ip_version` (String) IP version of adpative policy ACL
- `name` (String) Name of the adaptive policy ACL
- `rules` (Attributes Set) An ordered array of the adaptive policy ACL rules (see [below for nested schema](#nestedatt--rules))

### Read-Only

- `created_at` (String) When the adaptive policy ACL was created
- `updated_at` (String) When the adaptive policy ACL was last updated

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `dst_port` (String) Destination port
- `policy` (String) 'allow' or 'deny' traffic specified by this rule
- `protocol` (String) The type of protocol
- `src_port` (String) Source port

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_adaptive_policy_acls.example "acl_id,organization_id"
```