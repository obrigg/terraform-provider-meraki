---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_events_event_types Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_events_event_types (Data Source)



## Example Usage

```terraform
data "meraki_networks_events_event_types" "example" {

  network_id = "string"
}

output "meraki_networks_events_event_types_example" {
  value = data.meraki_networks_events_event_types.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `items` (Attributes List) Array of ResponseNetworksGetNetworkEventsEventTypes (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `category` (String) Event category
- `description` (String) Description of the event
- `type` (String) Event type