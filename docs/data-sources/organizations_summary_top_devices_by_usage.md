---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_summary_top_devices_by_usage Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_summary_top_devices_by_usage (Data Source)



## Example Usage

```terraform
data "meraki_organizations_summary_top_devices_by_usage" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_summary_top_devices_by_usage_example" {
  value = data.meraki_organizations_summary_top_devices_by_usage.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `t0` (String) t0 query parameter. The beginning of the timespan for the data.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `clients` (Attributes) Clients (see [below for nested schema](#nestedatt--items--clients))
- `mac` (String) Mac address of the device
- `model` (String) Model of the device
- `name` (String) Name of the device
- `network` (Attributes) Network info (see [below for nested schema](#nestedatt--items--network))
- `product_type` (String) Product type of the device
- `serial` (String) Serial number of the device
- `usage` (Attributes) Data usage of the device (see [below for nested schema](#nestedatt--items--usage))

<a id="nestedatt--items--clients"></a>
### Nested Schema for `items.clients`

Read-Only:

- `counts` (Attributes) Counts of clients (see [below for nested schema](#nestedatt--items--clients--counts))

<a id="nestedatt--items--clients--counts"></a>
### Nested Schema for `items.clients.counts`

Read-Only:

- `total` (Number) Total counts of clients



<a id="nestedatt--items--network"></a>
### Nested Schema for `items.network`

Read-Only:

- `id` (String) Network id
- `name` (String) Network name


<a id="nestedatt--items--usage"></a>
### Nested Schema for `items.usage`

Read-Only:

- `percentage` (Number) Data usage of the device by percentage
- `total` (Number) Total data usage of the device
