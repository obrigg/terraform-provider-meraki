---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_config_templates Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_config_templates (Data Source)



## Example Usage

```terraform
data "meraki_organizations_config_templates" "example" {

  organization_id = "string"
}

output "meraki_organizations_config_templates_example" {
  value = data.meraki_organizations_config_templates.example.items
}

data "meraki_organizations_config_templates" "example" {

  organization_id = "string"
}

output "meraki_organizations_config_templates_example" {
  value = data.meraki_organizations_config_templates.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `config_template_id` (String) configTemplateId path parameter. Config template ID
- `organization_id` (String) organizationId path parameter. Organization ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationConfigTemplates (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String)
- `name` (String)
- `product_types` (List of String)
- `time_zone` (String)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String)
- `name` (String)
- `product_types` (List of String)
- `time_zone` (String)
