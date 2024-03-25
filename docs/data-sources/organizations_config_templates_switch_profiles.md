---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_config_templates_switch_profiles Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_config_templates_switch_profiles (Data Source)



## Example Usage

```terraform
data "meraki_organizations_config_templates_switch_profiles" "example" {

  config_template_id = "string"
  organization_id    = "string"
}

output "meraki_organizations_config_templates_switch_profiles_example" {
  value = data.meraki_organizations_config_templates_switch_profiles.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config_template_id` (String) configTemplateId path parameter. Config template ID
- `organization_id` (String) organizationId path parameter. Organization ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `model` (String) Switch model
- `name` (String) Switch profile name
- `switch_profile_id` (String) Switch profile id