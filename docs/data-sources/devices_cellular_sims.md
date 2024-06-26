---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_cellular_sims Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_cellular_sims (Data Source)



## Example Usage

```terraform
data "meraki_devices_cellular_sims" "example" {

  serial = "string"
}

output "meraki_devices_cellular_sims_example" {
  value = data.meraki_devices_cellular_sims.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `sims` (Attributes Set) (see [below for nested schema](#nestedatt--item--sims))

<a id="nestedatt--item--sims"></a>
### Nested Schema for `item.sims`

Read-Only:

- `apns` (Attributes Set) (see [below for nested schema](#nestedatt--item--sims--apns))
- `is_primary` (Boolean)
- `slot` (String)

<a id="nestedatt--item--sims--apns"></a>
### Nested Schema for `item.sims.apns`

Read-Only:

- `allowed_ip_types` (List of String)
- `authentication` (Attributes) (see [below for nested schema](#nestedatt--item--sims--apns--authentication))
- `name` (String)

<a id="nestedatt--item--sims--apns--authentication"></a>
### Nested Schema for `item.sims.apns.name`

Read-Only:

- `type` (String)
- `username` (String)
