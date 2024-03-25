---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_topology_link_layer Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_topology_link_layer (Data Source)



## Example Usage

```terraform
data "meraki_networks_topology_link_layer" "example" {

  network_id = "string"
}

output "meraki_networks_topology_link_layer_example" {
  value = data.meraki_networks_topology_link_layer.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `errors` (List of String)
- `links` (Attributes Set) (see [below for nested schema](#nestedatt--item--links))
- `nodes` (Attributes Set) (see [below for nested schema](#nestedatt--item--nodes))

<a id="nestedatt--item--links"></a>
### Nested Schema for `item.links`

Read-Only:

- `ends` (Attributes Set) (see [below for nested schema](#nestedatt--item--links--ends))
- `last_reported_at` (String)

<a id="nestedatt--item--links--ends"></a>
### Nested Schema for `item.links.ends`

Read-Only:

- `device` (Attributes) (see [below for nested schema](#nestedatt--item--links--ends--device))
- `discovered` (Attributes) (see [below for nested schema](#nestedatt--item--links--ends--discovered))
- `node` (Attributes) (see [below for nested schema](#nestedatt--item--links--ends--node))

<a id="nestedatt--item--links--ends--device"></a>
### Nested Schema for `item.links.ends.node`

Read-Only:

- `name` (String)
- `serial` (String)


<a id="nestedatt--item--links--ends--discovered"></a>
### Nested Schema for `item.links.ends.node`

Read-Only:

- `cdp` (Attributes) (see [below for nested schema](#nestedatt--item--links--ends--node--cdp))
- `lldp` (Attributes) (see [below for nested schema](#nestedatt--item--links--ends--node--lldp))

<a id="nestedatt--item--links--ends--node--cdp"></a>
### Nested Schema for `item.links.ends.node.cdp`

Read-Only:

- `native_vlan` (Number)
- `port_id` (String)


<a id="nestedatt--item--links--ends--node--lldp"></a>
### Nested Schema for `item.links.ends.node.lldp`

Read-Only:

- `port_description` (String)
- `port_id` (String)



<a id="nestedatt--item--links--ends--node"></a>
### Nested Schema for `item.links.ends.node`

Read-Only:

- `derived_id` (String)
- `type` (String)




<a id="nestedatt--item--nodes"></a>
### Nested Schema for `item.nodes`

Read-Only:

- `derived_id` (String)
- `discovered` (Attributes) (see [below for nested schema](#nestedatt--item--nodes--discovered))
- `mac` (String)
- `root` (Boolean)
- `type` (String)

<a id="nestedatt--item--nodes--discovered"></a>
### Nested Schema for `item.nodes.discovered`

Read-Only:

- `cdp` (String)
- `lldp` (Attributes) (see [below for nested schema](#nestedatt--item--nodes--discovered--lldp))

<a id="nestedatt--item--nodes--discovered--lldp"></a>
### Nested Schema for `item.nodes.discovered.lldp`

Read-Only:

- `chassis_id` (String)
- `management_address` (String)
- `system_capabilities` (List of String)
- `system_description` (String)
- `system_name` (String)