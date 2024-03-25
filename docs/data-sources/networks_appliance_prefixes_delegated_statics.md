---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_prefixes_delegated_statics Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_appliance_prefixes_delegated_statics (Data Source)



## Example Usage

```terraform
data "meraki_networks_appliance_prefixes_delegated_statics" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_prefixes_delegated_statics_example" {
  value = data.meraki_networks_appliance_prefixes_delegated_statics.example.items
}

data "meraki_networks_appliance_prefixes_delegated_statics" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_prefixes_delegated_statics_example" {
  value = data.meraki_networks_appliance_prefixes_delegated_statics.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `network_id` (String) networkId path parameter. Network ID
- `static_delegated_prefix_id` (String) staticDelegatedPrefixId path parameter. Static delegated prefix ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `created_at` (String) Prefix creation time.
- `description` (String) Identifying description for the prefix.
- `origin` (Attributes) WAN1/WAN2/Independent prefix. (see [below for nested schema](#nestedatt--item--origin))
- `prefix` (String) IPv6 prefix/prefix length.
- `static_delegated_prefix_id` (String) Static delegated prefix id.
- `updated_at` (String) Prefix Updated time.

<a id="nestedatt--item--origin"></a>
### Nested Schema for `item.origin`

Read-Only:

- `interfaces` (List of String) Uplink provided or independent
- `type` (String) Origin type



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `created_at` (String) Prefix creation time.
- `description` (String) Identifying description for the prefix.
- `origin` (Attributes) WAN1/WAN2/Independent prefix. (see [below for nested schema](#nestedatt--items--origin))
- `prefix` (String) IPv6 prefix/prefix length.
- `static_delegated_prefix_id` (String) Static delegated prefix id.
- `updated_at` (String) Prefix Updated time.

<a id="nestedatt--items--origin"></a>
### Nested Schema for `items.origin`

Read-Only:

- `interfaces` (List of String) Uplink provided or independent
- `type` (String) Origin type