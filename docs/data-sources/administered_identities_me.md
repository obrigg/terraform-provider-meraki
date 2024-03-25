---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_administered_identities_me Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_administered_identities_me (Data Source)



## Example Usage

```terraform
data "meraki_administered_identities_me" "example" {

}

output "meraki_administered_identities_me_example" {
  value = data.meraki_administered_identities_me.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `authentication` (Attributes) Authentication info (see [below for nested schema](#nestedatt--item--authentication))
- `email` (String) User email
- `last_used_dashboard_at` (String) Last seen active on Dashboard UI
- `name` (String) Username

<a id="nestedatt--item--authentication"></a>
### Nested Schema for `item.authentication`

Read-Only:

- `api` (Attributes) API authentication (see [below for nested schema](#nestedatt--item--authentication--api))
- `mode` (String) Authentication mode
- `saml` (Attributes) SAML authentication (see [below for nested schema](#nestedatt--item--authentication--saml))
- `two_factor` (Attributes) TwoFactor authentication (see [below for nested schema](#nestedatt--item--authentication--two_factor))

<a id="nestedatt--item--authentication--api"></a>
### Nested Schema for `item.authentication.api`

Read-Only:

- `key` (Attributes) API key (see [below for nested schema](#nestedatt--item--authentication--api--key))

<a id="nestedatt--item--authentication--api--key"></a>
### Nested Schema for `item.authentication.api.key`

Read-Only:

- `created` (Boolean) If API key is created for this user



<a id="nestedatt--item--authentication--saml"></a>
### Nested Schema for `item.authentication.saml`

Read-Only:

- `enabled` (Boolean) If SAML authentication is enabled for this user


<a id="nestedatt--item--authentication--two_factor"></a>
### Nested Schema for `item.authentication.two_factor`

Read-Only:

- `enabled` (Boolean) If twoFactor authentication is enabled for this user