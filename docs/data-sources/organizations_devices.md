---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_devices Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_devices (Data Source)



## Example Usage

```terraform
data "meraki_organizations_devices" "example" {

  configuration_updated_after = "string"
  ending_before               = "string"
  mac                         = "string"
  macs                        = ["string"]
  model                       = "string"
  models                      = ["string"]
  name                        = "string"
  network_ids                 = ["string"]
  organization_id             = "string"
  per_page                    = 1
  product_types               = ["string"]
  sensor_alert_profile_ids    = ["string"]
  sensor_metrics              = ["string"]
  serial                      = "string"
  serials                     = ["string"]
  starting_after              = "string"
  tags                        = ["string"]
  tags_filter_type            = "string"
}

output "meraki_organizations_devices_example" {
  value = data.meraki_organizations_devices.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `configuration_updated_after` (String) configurationUpdatedAfter query parameter. Filter results by whether or not the device's configuration has been updated after the given timestamp
- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `mac` (String) mac query parameter. Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match.
- `macs` (List of String) macs query parameter. Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match.
- `model` (String) model query parameter. Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match.
- `models` (List of String) models query parameter. Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match.
- `name` (String) name query parameter. Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match.
- `network_ids` (List of String) networkIds query parameter. Optional parameter to filter devices by network.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.
- `product_types` (List of String) productTypes query parameter. Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, and sensor.
- `sensor_alert_profile_ids` (List of String) sensorAlertProfileIds query parameter. Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices.
- `sensor_metrics` (List of String) sensorMetrics query parameter. Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices.
- `serial` (String) serial query parameter. Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match.
- `serials` (List of String) serials query parameter. Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `tags` (List of String) tags query parameter. Optional parameter to filter devices by tags.
- `tags_filter_type` (String) tagsFilterType query parameter. Optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationDevices (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `address` (String) Physical address of the device
- `firmware` (String) Firmware version of the device
- `lan_ip` (String) LAN IP address of the device
- `lat` (Number) Latitude of the device
- `lng` (Number) Longitude of the device
- `mac` (String) MAC address of the device
- `model` (String) Model of the device
- `name` (String) Name of the device
- `network_id` (String) ID of the network the device belongs to
- `notes` (String) Notes for the device, limited to 255 characters
- `product_type` (String) Product type of the device
- `serial` (String) Serial number of the device
- `tags` (List of String) List of tags assigned to the device