---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_webhooks_http_servers Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_webhooks_http_servers (Data Source)



## Example Usage

```terraform
data "meraki_networks_webhooks_http_servers" "example" {

  network_id = "string"
}

output "meraki_networks_webhooks_http_servers_example" {
  value = data.meraki_networks_webhooks_http_servers.example.items
}

data "meraki_networks_webhooks_http_servers" "example" {

  network_id = "string"
}

output "meraki_networks_webhooks_http_servers_example" {
  value = data.meraki_networks_webhooks_http_servers.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `http_server_id` (String) httpServerId path parameter. Http server ID
- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseNetworksGetNetworkWebhooksHttpServers (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String) A Base64 encoded ID.
- `name` (String) A name for easy reference to the HTTP server
- `network_id` (String) A Meraki network ID.
- `payload_template` (Attributes) The payload template to use when posting data to the HTTP server. (see [below for nested schema](#nestedatt--item--payload_template))
- `url` (String) The URL of the HTTP server.

<a id="nestedatt--item--payload_template"></a>
### Nested Schema for `item.payload_template`

Read-Only:

- `name` (String) The name of the payload template.
- `payload_template_id` (String) The ID of the payload template.



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String) A Base64 encoded ID.
- `name` (String) A name for easy reference to the HTTP server
- `network_id` (String) A Meraki network ID.
- `payload_template` (Attributes) The payload template to use when posting data to the HTTP server. (see [below for nested schema](#nestedatt--items--payload_template))
- `url` (String) The URL of the HTTP server.

<a id="nestedatt--items--payload_template"></a>
### Nested Schema for `items.payload_template`

Read-Only:

- `name` (String) The name of the payload template.
- `payload_template_id` (String) The ID of the payload template.