---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_stacks Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_switch_stacks (Resource)



## Example Usage

```terraform
resource "meraki_networks_switch_stacks" "example" {

  name       = "A cool stack"
  network_id = "string"
  serials    = ["QBZY-XWVU-TSRQ", "QBAB-CDEF-GHIJ"]
}

output "meraki_networks_switch_stacks_example" {
  value = meraki_networks_switch_stacks.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `name` (String) Switch stacks name
- `serials` (List of String) Serials of the switches in the switch stack
- `switch_stack_id` (String) switchStackId path parameter. Switch stack ID

### Read-Only

- `id` (String) Switch stacks id

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_switch_stacks.example "network_id,switch_stack_id"
```
