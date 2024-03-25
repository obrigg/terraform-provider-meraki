---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_firmware_upgrades_by_device Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_firmware_upgrades_by_device (Data Source)



## Example Usage

```terraform
data "meraki_organizations_firmware_upgrades_by_device" "example" {

  ending_before              = "string"
  firmware_upgrade_batch_ids = ["string"]
  firmware_upgrade_ids       = ["string"]
  macs                       = ["string"]
  network_ids                = ["string"]
  organization_id            = "string"
  per_page                   = 1
  serials                    = ["string"]
  starting_after             = "string"
}

output "meraki_organizations_firmware_upgrades_by_device_example" {
  value = data.meraki_organizations_firmware_upgrades_by_device.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `firmware_upgrade_batch_ids` (List of String) firmwareUpgradeBatchIds query parameter. Optional parameter to filter by firmware upgrade batch ids.
- `firmware_upgrade_ids` (List of String) firmwareUpgradeIds query parameter. Optional parameter to filter by firmware upgrade ids.
- `macs` (List of String) macs query parameter. Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.
- `network_ids` (List of String) networkIds query parameter. Optional parameter to filter by network
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 50. Default is 50.
- `serials` (List of String) serials query parameter. Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `device_status` (String) Status of the device upgrade
- `name` (String) Name assigned to the device
- `serial` (String) Serial of the device
- `upgrade` (Attributes) The devices upgrade details and status (see [below for nested schema](#nestedatt--items--upgrade))

<a id="nestedatt--items--upgrade"></a>
### Nested Schema for `items.upgrade`

Read-Only:

- `from_version` (Attributes) The initial version of the device (see [below for nested schema](#nestedatt--items--upgrade--from_version))
- `id` (String) ID of the upgrade
- `staged` (Attributes) Staged upgrade (see [below for nested schema](#nestedatt--items--upgrade--staged))
- `status` (String) Status of the upgrade
- `time` (String) Start time of the upgrade
- `to_version` (Attributes) Version the device is upgrading to (see [below for nested schema](#nestedatt--items--upgrade--to_version))
- `upgrade_batch_id` (String) ID of the upgrade batch

<a id="nestedatt--items--upgrade--from_version"></a>
### Nested Schema for `items.upgrade.from_version`

Read-Only:

- `id` (String) ID of the initial firmware version
- `release_date` (String) Release date of the firmware version
- `release_type` (String) Release type of the firmware version
- `short_name` (String) Firmware version short name


<a id="nestedatt--items--upgrade--staged"></a>
### Nested Schema for `items.upgrade.staged`

Read-Only:

- `group` (Attributes) The staged upgrade group (see [below for nested schema](#nestedatt--items--upgrade--staged--group))

<a id="nestedatt--items--upgrade--staged--group"></a>
### Nested Schema for `items.upgrade.staged.group`

Read-Only:

- `id` (String) Id of the staged upgrade group



<a id="nestedatt--items--upgrade--to_version"></a>
### Nested Schema for `items.upgrade.to_version`

Read-Only:

- `id` (String) ID of the initial firmware version
- `release_date` (String) Release date of the firmware version
- `release_type` (String) Release type of the firmware version
- `short_name` (String) Firmware version short name