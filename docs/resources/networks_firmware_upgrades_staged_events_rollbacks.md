---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_firmware_upgrades_staged_events_rollbacks Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_firmware_upgrades_staged_events_rollbacks (Resource)



~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.

## Example Usage

```terraform
resource "meraki_networks_firmware_upgrades_staged_events_rollbacks" "example" {

  network_id = "string"
  parameters = {

    reasons = [{

      category = "performance"
      comment  = "Network was slower with the upgrade"
    }]
    stages = [{

      group = {

        id = "1234"
      }
      milestones = {

        scheduled_for = "2018-02-11T00:00:00Z"
      }
    }]
  }
}

output "meraki_networks_firmware_upgrades_staged_events_rollbacks_example" {
  value = meraki_networks_firmware_upgrades_staged_events_rollbacks.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `parameters` (Attributes) (see [below for nested schema](#nestedatt--parameters))

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `reasons` (Attributes Set) The reason for rolling back the staged upgrade (see [below for nested schema](#nestedatt--parameters--reasons))
- `stages` (Attributes Set) All completed or in-progress stages in the network with their new start times. All pending stages will be canceled (see [below for nested schema](#nestedatt--parameters--stages))

<a id="nestedatt--parameters--reasons"></a>
### Nested Schema for `parameters.reasons`

Optional:

- `category` (String) Reason for the rollback
- `comment` (String) Additional comment about the rollback


<a id="nestedatt--parameters--stages"></a>
### Nested Schema for `parameters.stages`

Optional:

- `group` (Attributes) The Staged Upgrade Group containing the name and ID (see [below for nested schema](#nestedatt--parameters--stages--group))
- `milestones` (Attributes) The Staged Upgrade Milestones for the specific stage (see [below for nested schema](#nestedatt--parameters--stages--milestones))

<a id="nestedatt--parameters--stages--group"></a>
### Nested Schema for `parameters.stages.group`

Optional:

- `id` (String) ID of the Staged Upgrade Group


<a id="nestedatt--parameters--stages--milestones"></a>
### Nested Schema for `parameters.stages.milestones`

Optional:

- `scheduled_for` (String) The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)




<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `products` (Attributes) The network devices to be updated (see [below for nested schema](#nestedatt--item--products))
- `reasons` (Attributes Set) Reasons for the rollback (see [below for nested schema](#nestedatt--item--reasons))
- `stages` (Attributes Set) The ordered stages in the network (see [below for nested schema](#nestedatt--item--stages))

<a id="nestedatt--item--products"></a>
### Nested Schema for `item.products`

Read-Only:

- `switch` (Attributes) The Switch network to be updated (see [below for nested schema](#nestedatt--item--products--switch))

<a id="nestedatt--item--products--switch"></a>
### Nested Schema for `item.products.switch`

Read-Only:

- `next_upgrade` (Attributes) Details of the next firmware upgrade (see [below for nested schema](#nestedatt--item--products--switch--next_upgrade))

<a id="nestedatt--item--products--switch--next_upgrade"></a>
### Nested Schema for `item.products.switch.next_upgrade`

Read-Only:

- `to_version` (Attributes) Details of the version the device will upgrade to (see [below for nested schema](#nestedatt--item--products--switch--next_upgrade--to_version))

<a id="nestedatt--item--products--switch--next_upgrade--to_version"></a>
### Nested Schema for `item.products.switch.next_upgrade.to_version`

Read-Only:

- `id` (String) Id of the Version being upgraded to
- `short_name` (String) Firmware version short name





<a id="nestedatt--item--reasons"></a>
### Nested Schema for `item.reasons`

Read-Only:

- `category` (String) Reason for the rollback
- `comment` (String) Additional comment about the rollback


<a id="nestedatt--item--stages"></a>
### Nested Schema for `item.stages`

Read-Only:

- `group` (Attributes) The staged upgrade group (see [below for nested schema](#nestedatt--item--stages--group))
- `milestones` (Attributes) The Staged Upgrade Milestones for the stage (see [below for nested schema](#nestedatt--item--stages--milestones))
- `status` (String) Current upgrade status of the group

<a id="nestedatt--item--stages--group"></a>
### Nested Schema for `item.stages.group`

Read-Only:

- `description` (String) Description of the Staged Upgrade Group
- `id` (String) Id of the Staged Upgrade Group
- `name` (String) Name of the Staged Upgrade Group


<a id="nestedatt--item--stages--milestones"></a>
### Nested Schema for `item.stages.milestones`

Read-Only:

- `canceled_at` (String) Time that the group was canceled
- `completed_at` (String) Finish time for the group
- `scheduled_for` (String) Scheduled start time for the group
- `started_at` (String) Start time for the group