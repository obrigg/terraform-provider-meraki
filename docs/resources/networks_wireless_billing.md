---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_billing Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_wireless_billing (Resource)



## Example Usage

```terraform
resource "meraki_networks_wireless_billing" "example" {

  currency   = "USD"
  network_id = "string"
  plans = [{

    bandwidth_limits = {

      limit_down = 1000
      limit_up   = 1000
    }
    price      = 5
    time_limit = "1 hour"
  }]
}

output "meraki_networks_wireless_billing_example" {
  value = meraki_networks_wireless_billing.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `currency` (String) The currency code of this node group's billing plans
- `plans` (Attributes Set) Array of billing plans in the node group. (Can configure a maximum of 5) (see [below for nested schema](#nestedatt--plans))

<a id="nestedatt--plans"></a>
### Nested Schema for `plans`

Optional:

- `bandwidth_limits` (Attributes) The uplink bandwidth settings for the pricing plan. (see [below for nested schema](#nestedatt--plans--bandwidth_limits))
- `id` (String) The id of the pricing plan to update.
- `price` (Number) The price of the billing plan.
- `time_limit` (String) The time limit of the pricing plan in minutes. Can be '1 hour', '1 day', '1 week', or '30 days'.

<a id="nestedatt--plans--bandwidth_limits"></a>
### Nested Schema for `plans.bandwidth_limits`

Optional:

- `limit_down` (Number) The maximum download limit (integer, in Kbps). null indicates no limit
- `limit_up` (Number) The maximum upload limit (integer, in Kbps). null indicates no limit

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_wireless_billing.example "network_id"
```