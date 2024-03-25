---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_devices_statuses_overview Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_devices_statuses_overview (Data Source)



## Example Usage

```terraform
data "meraki_organizations_devices_statuses_overview" "example" {

  network_ids     = ["string"]
  organization_id = "string"
  product_types   = ["string"]
}

output "meraki_organizations_devices_statuses_overview_example" {
  value = data.meraki_organizations_devices_statuses_overview.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `network_ids` (List of String) networkIds query parameter. An optional parameter to filter device statuses by network.
- `product_types` (List of String) productTypes query parameter. An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `counts` (Attributes) counts (see [below for nested schema](#nestedatt--item--counts))

<a id="nestedatt--item--counts"></a>
### Nested Schema for `item.counts`

Read-Only:

- `by_status` (Attributes) byStatus (see [below for nested schema](#nestedatt--item--counts--by_status))

<a id="nestedatt--item--counts--by_status"></a>
### Nested Schema for `item.counts.by_status`

Read-Only:

- `alerting` (Number) alerting count
- `dormant` (Number) dormant count
- `offline` (Number) offline count
- `online` (Number) online count