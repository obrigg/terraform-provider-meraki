---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_group_policies Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_group_policies (Data Source)



## Example Usage

```terraform
data "meraki_networks_group_policies" "example" {

  network_id = "string"
}

output "meraki_networks_group_policies_example" {
  value = data.meraki_networks_group_policies.example.items
}

data "meraki_networks_group_policies" "example" {

  network_id = "string"
}

output "meraki_networks_group_policies_example" {
  value = data.meraki_networks_group_policies.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `group_policy_id` (String) groupPolicyId path parameter. Group policy ID
- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseNetworksGetNetworkGroupPolicies (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `bandwidth` (Attributes) (see [below for nested schema](#nestedatt--item--bandwidth))
- `bonjour_forwarding` (Attributes) (see [below for nested schema](#nestedatt--item--bonjour_forwarding))
- `content_filtering` (Attributes) (see [below for nested schema](#nestedatt--item--content_filtering))
- `firewall_and_traffic_shaping` (Attributes) (see [below for nested schema](#nestedatt--item--firewall_and_traffic_shaping))
- `group_policy_id` (String)
- `name` (String)
- `scheduling` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling))
- `splash_auth_settings` (String)
- `vlan_tagging` (Attributes) (see [below for nested schema](#nestedatt--item--vlan_tagging))

<a id="nestedatt--item--bandwidth"></a>
### Nested Schema for `item.bandwidth`

Read-Only:

- `bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--item--bandwidth--bandwidth_limits))
- `settings` (String)

<a id="nestedatt--item--bandwidth--bandwidth_limits"></a>
### Nested Schema for `item.bandwidth.bandwidth_limits`

Read-Only:

- `limit_down` (Number)
- `limit_up` (Number)



<a id="nestedatt--item--bonjour_forwarding"></a>
### Nested Schema for `item.bonjour_forwarding`

Read-Only:

- `rules` (Attributes Set) (see [below for nested schema](#nestedatt--item--bonjour_forwarding--rules))
- `settings` (String)

<a id="nestedatt--item--bonjour_forwarding--rules"></a>
### Nested Schema for `item.bonjour_forwarding.rules`

Read-Only:

- `description` (String)
- `services` (List of String)
- `vlan_id` (String)



<a id="nestedatt--item--content_filtering"></a>
### Nested Schema for `item.content_filtering`

Read-Only:

- `allowed_url_patterns` (Attributes) (see [below for nested schema](#nestedatt--item--content_filtering--allowed_url_patterns))
- `blocked_url_categories` (Attributes) (see [below for nested schema](#nestedatt--item--content_filtering--blocked_url_categories))
- `blocked_url_patterns` (Attributes) (see [below for nested schema](#nestedatt--item--content_filtering--blocked_url_patterns))

<a id="nestedatt--item--content_filtering--allowed_url_patterns"></a>
### Nested Schema for `item.content_filtering.allowed_url_patterns`

Read-Only:

- `patterns` (List of String)
- `settings` (String)


<a id="nestedatt--item--content_filtering--blocked_url_categories"></a>
### Nested Schema for `item.content_filtering.blocked_url_categories`

Read-Only:

- `categories` (List of String)
- `settings` (String)


<a id="nestedatt--item--content_filtering--blocked_url_patterns"></a>
### Nested Schema for `item.content_filtering.blocked_url_patterns`

Read-Only:

- `patterns` (List of String)
- `settings` (String)



<a id="nestedatt--item--firewall_and_traffic_shaping"></a>
### Nested Schema for `item.firewall_and_traffic_shaping`

Read-Only:

- `l3_firewall_rules` (Attributes Set) (see [below for nested schema](#nestedatt--item--firewall_and_traffic_shaping--l3_firewall_rules))
- `l7_firewall_rules` (Attributes Set) (see [below for nested schema](#nestedatt--item--firewall_and_traffic_shaping--l7_firewall_rules))
- `settings` (String)
- `traffic_shaping_rules` (Attributes Set) (see [below for nested schema](#nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules))

<a id="nestedatt--item--firewall_and_traffic_shaping--l3_firewall_rules"></a>
### Nested Schema for `item.firewall_and_traffic_shaping.l3_firewall_rules`

Read-Only:

- `comment` (String)
- `dest_cidr` (String)
- `dest_port` (String)
- `policy` (String)
- `protocol` (String)


<a id="nestedatt--item--firewall_and_traffic_shaping--l7_firewall_rules"></a>
### Nested Schema for `item.firewall_and_traffic_shaping.l7_firewall_rules`

Read-Only:

- `policy` (String)
- `type` (String)
- `value` (String)


<a id="nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules"></a>
### Nested Schema for `item.firewall_and_traffic_shaping.traffic_shaping_rules`

Read-Only:

- `definitions` (Attributes Set) (see [below for nested schema](#nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules--definitions))
- `dscp_tag_value` (Number)
- `pcp_tag_value` (Number)
- `per_client_bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits))

<a id="nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules--definitions"></a>
### Nested Schema for `item.firewall_and_traffic_shaping.traffic_shaping_rules.per_client_bandwidth_limits`

Read-Only:

- `type` (String)
- `value` (String)


<a id="nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits"></a>
### Nested Schema for `item.firewall_and_traffic_shaping.traffic_shaping_rules.per_client_bandwidth_limits`

Read-Only:

- `bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits--bandwidth_limits))
- `settings` (String)

<a id="nestedatt--item--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits--bandwidth_limits"></a>
### Nested Schema for `item.firewall_and_traffic_shaping.traffic_shaping_rules.per_client_bandwidth_limits.bandwidth_limits`

Read-Only:

- `limit_down` (Number)
- `limit_up` (Number)





<a id="nestedatt--item--scheduling"></a>
### Nested Schema for `item.scheduling`

Read-Only:

- `enabled` (Boolean)
- `friday` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling--friday))
- `monday` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling--monday))
- `saturday` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling--saturday))
- `sunday` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling--sunday))
- `thursday` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling--thursday))
- `tuesday` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling--tuesday))
- `wednesday` (Attributes) (see [below for nested schema](#nestedatt--item--scheduling--wednesday))

<a id="nestedatt--item--scheduling--friday"></a>
### Nested Schema for `item.scheduling.friday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--item--scheduling--monday"></a>
### Nested Schema for `item.scheduling.monday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--item--scheduling--saturday"></a>
### Nested Schema for `item.scheduling.saturday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--item--scheduling--sunday"></a>
### Nested Schema for `item.scheduling.sunday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--item--scheduling--thursday"></a>
### Nested Schema for `item.scheduling.thursday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--item--scheduling--tuesday"></a>
### Nested Schema for `item.scheduling.tuesday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--item--scheduling--wednesday"></a>
### Nested Schema for `item.scheduling.wednesday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)



<a id="nestedatt--item--vlan_tagging"></a>
### Nested Schema for `item.vlan_tagging`

Read-Only:

- `settings` (String)
- `vlan_id` (String)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `bandwidth` (Attributes) (see [below for nested schema](#nestedatt--items--bandwidth))
- `bonjour_forwarding` (Attributes) (see [below for nested schema](#nestedatt--items--bonjour_forwarding))
- `content_filtering` (Attributes) (see [below for nested schema](#nestedatt--items--content_filtering))
- `firewall_and_traffic_shaping` (Attributes) (see [below for nested schema](#nestedatt--items--firewall_and_traffic_shaping))
- `group_policy_id` (String)
- `name` (String)
- `scheduling` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling))
- `splash_auth_settings` (String)
- `vlan_tagging` (Attributes) (see [below for nested schema](#nestedatt--items--vlan_tagging))

<a id="nestedatt--items--bandwidth"></a>
### Nested Schema for `items.bandwidth`

Read-Only:

- `bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--items--bandwidth--bandwidth_limits))
- `settings` (String)

<a id="nestedatt--items--bandwidth--bandwidth_limits"></a>
### Nested Schema for `items.bandwidth.bandwidth_limits`

Read-Only:

- `limit_down` (Number)
- `limit_up` (Number)



<a id="nestedatt--items--bonjour_forwarding"></a>
### Nested Schema for `items.bonjour_forwarding`

Read-Only:

- `rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--bonjour_forwarding--rules))
- `settings` (String)

<a id="nestedatt--items--bonjour_forwarding--rules"></a>
### Nested Schema for `items.bonjour_forwarding.rules`

Read-Only:

- `description` (String)
- `services` (List of String)
- `vlan_id` (String)



<a id="nestedatt--items--content_filtering"></a>
### Nested Schema for `items.content_filtering`

Read-Only:

- `allowed_url_patterns` (Attributes) (see [below for nested schema](#nestedatt--items--content_filtering--allowed_url_patterns))
- `blocked_url_categories` (Attributes) (see [below for nested schema](#nestedatt--items--content_filtering--blocked_url_categories))
- `blocked_url_patterns` (Attributes) (see [below for nested schema](#nestedatt--items--content_filtering--blocked_url_patterns))

<a id="nestedatt--items--content_filtering--allowed_url_patterns"></a>
### Nested Schema for `items.content_filtering.allowed_url_patterns`

Read-Only:

- `patterns` (List of String)
- `settings` (String)


<a id="nestedatt--items--content_filtering--blocked_url_categories"></a>
### Nested Schema for `items.content_filtering.blocked_url_categories`

Read-Only:

- `categories` (List of String)
- `settings` (String)


<a id="nestedatt--items--content_filtering--blocked_url_patterns"></a>
### Nested Schema for `items.content_filtering.blocked_url_patterns`

Read-Only:

- `patterns` (List of String)
- `settings` (String)



<a id="nestedatt--items--firewall_and_traffic_shaping"></a>
### Nested Schema for `items.firewall_and_traffic_shaping`

Read-Only:

- `l3_firewall_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--firewall_and_traffic_shaping--l3_firewall_rules))
- `l7_firewall_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--firewall_and_traffic_shaping--l7_firewall_rules))
- `settings` (String)
- `traffic_shaping_rules` (Attributes Set) (see [below for nested schema](#nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules))

<a id="nestedatt--items--firewall_and_traffic_shaping--l3_firewall_rules"></a>
### Nested Schema for `items.firewall_and_traffic_shaping.l3_firewall_rules`

Read-Only:

- `comment` (String)
- `dest_cidr` (String)
- `dest_port` (String)
- `policy` (String)
- `protocol` (String)


<a id="nestedatt--items--firewall_and_traffic_shaping--l7_firewall_rules"></a>
### Nested Schema for `items.firewall_and_traffic_shaping.l7_firewall_rules`

Read-Only:

- `policy` (String)
- `type` (String)
- `value` (String)


<a id="nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules"></a>
### Nested Schema for `items.firewall_and_traffic_shaping.traffic_shaping_rules`

Read-Only:

- `definitions` (Attributes Set) (see [below for nested schema](#nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules--definitions))
- `dscp_tag_value` (Number)
- `pcp_tag_value` (Number)
- `per_client_bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits))

<a id="nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules--definitions"></a>
### Nested Schema for `items.firewall_and_traffic_shaping.traffic_shaping_rules.per_client_bandwidth_limits`

Read-Only:

- `type` (String)
- `value` (String)


<a id="nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits"></a>
### Nested Schema for `items.firewall_and_traffic_shaping.traffic_shaping_rules.per_client_bandwidth_limits`

Read-Only:

- `bandwidth_limits` (Attributes) (see [below for nested schema](#nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits--bandwidth_limits))
- `settings` (String)

<a id="nestedatt--items--firewall_and_traffic_shaping--traffic_shaping_rules--per_client_bandwidth_limits--bandwidth_limits"></a>
### Nested Schema for `items.firewall_and_traffic_shaping.traffic_shaping_rules.per_client_bandwidth_limits.bandwidth_limits`

Read-Only:

- `limit_down` (Number)
- `limit_up` (Number)





<a id="nestedatt--items--scheduling"></a>
### Nested Schema for `items.scheduling`

Read-Only:

- `enabled` (Boolean)
- `friday` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling--friday))
- `monday` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling--monday))
- `saturday` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling--saturday))
- `sunday` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling--sunday))
- `thursday` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling--thursday))
- `tuesday` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling--tuesday))
- `wednesday` (Attributes) (see [below for nested schema](#nestedatt--items--scheduling--wednesday))

<a id="nestedatt--items--scheduling--friday"></a>
### Nested Schema for `items.scheduling.friday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--items--scheduling--monday"></a>
### Nested Schema for `items.scheduling.monday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--items--scheduling--saturday"></a>
### Nested Schema for `items.scheduling.saturday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--items--scheduling--sunday"></a>
### Nested Schema for `items.scheduling.sunday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--items--scheduling--thursday"></a>
### Nested Schema for `items.scheduling.thursday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--items--scheduling--tuesday"></a>
### Nested Schema for `items.scheduling.tuesday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)


<a id="nestedatt--items--scheduling--wednesday"></a>
### Nested Schema for `items.scheduling.wednesday`

Read-Only:

- `active` (Boolean)
- `from` (String)
- `to` (String)



<a id="nestedatt--items--vlan_tagging"></a>
### Nested Schema for `items.vlan_tagging`

Read-Only:

- `settings` (String)
- `vlan_id` (String)