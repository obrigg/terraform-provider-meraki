---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_bluetooth_clients Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_bluetooth_clients (Data Source)



## Example Usage

```terraform
data "meraki_networks_bluetooth_clients" "example" {

  bluetooth_client_id           = "string"
  connectivity_history_timespan = 1
  include_connectivity_history  = false
  network_id                    = "string"
}

output "meraki_networks_bluetooth_clients_example" {
  value = data.meraki_networks_bluetooth_clients.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `bluetooth_client_id` (String) bluetoothClientId path parameter. Bluetooth client ID
- `network_id` (String) networkId path parameter. Network ID

### Optional

- `connectivity_history_timespan` (Number) connectivityHistoryTimespan query parameter. The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used.
- `include_connectivity_history` (Boolean) includeConnectivityHistory query parameter. Include the connectivity history for this client

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `device_name` (String)
- `id` (String)
- `in_sight_alert` (Boolean)
- `last_seen` (Number)
- `mac` (String)
- `manufacturer` (String)
- `name` (String)
- `network_id` (String)
- `out_of_sight_alert` (Boolean)
- `seen_by_device_mac` (String)
- `tags` (List of String)
