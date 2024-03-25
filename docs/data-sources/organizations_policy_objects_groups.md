---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_policy_objects_groups Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_policy_objects_groups (Data Source)



## Example Usage

```terraform
data "meraki_organizations_policy_objects_groups" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_policy_objects_groups_example" {
  value = data.meraki_organizations_policy_objects_groups.example.items
}

data "meraki_organizations_policy_objects_groups" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_policy_objects_groups_example" {
  value = data.meraki_organizations_policy_objects_groups.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `organization_id` (String) organizationId path parameter. Organization ID
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 10 1000. Default is 1000.
- `policy_object_group_id` (String) policyObjectGroupId path parameter. Policy object group ID
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationPolicyObjectsGroups (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `category` (String)
- `created_at` (String)
- `id` (String)
- `name` (String)
- `network_ids` (List of String)
- `object_ids` (List of String)
- `updated_at` (String)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `category` (String)
- `created_at` (String)
- `id` (String)
- `name` (String)
- `network_ids` (List of String)
- `object_ids` (List of String)
- `updated_at` (String)