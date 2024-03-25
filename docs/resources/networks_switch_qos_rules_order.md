---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_qos_rules_order Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_switch_qos_rules_order (Resource)



## Example Usage

```terraform
resource "meraki_networks_switch_qos_rules_order" "example" {

  dscp           = 1
  dst_port       = 3000
  dst_port_range = "3000-3100"
  network_id     = "string"
  protocol       = "TCP"
  src_port       = 2000
  src_port_range = "70-80"
  vlan           = 100
}

output "meraki_networks_switch_qos_rules_order_example" {
  value = meraki_networks_switch_qos_rules_order.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `dscp` (Number) DSCP tag. Set this to -1 to trust incoming DSCP. Default value is 0
- `dst_port` (Number) The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
- `dst_port_range` (String) The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
- `protocol` (String) The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
- `qos_rule_id` (String) qosRuleId path parameter. Qos rule ID
- `src_port` (Number) The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
- `src_port_range` (String) The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
- `vlan` (Number) The VLAN of the incoming packet. A null value will match any VLAN.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_switch_qos_rules_order.example "network_id,qos_rule_id"
```