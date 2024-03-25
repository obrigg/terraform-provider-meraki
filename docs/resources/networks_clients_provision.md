---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_clients_provision Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_clients_provision (Resource)



~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.

## Example Usage

```terraform
resource "meraki_networks_clients_provision" "example" {

  network_id = "string"
  parameters = {

    clients = [{

      mac  = "00:11:22:33:44:55"
      name = "Miles's phone"
    }]
    device_policy   = "Group policy"
    group_policy_id = "101"
  }
}

output "meraki_networks_clients_provision_example" {
  value = meraki_networks_clients_provision.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `parameters` (Attributes) (see [below for nested schema](#nestedatt--parameters))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `clients` (Attributes Set) The array of clients to provision (see [below for nested schema](#nestedatt--parameters--clients))
- `device_policy` (String) The policy to apply to the specified client. Can be 'Group policy', 'Allowed', 'Blocked', 'Per connection' or 'Normal'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.
- `policies_by_security_appliance` (Attributes) An object, describing what the policy-connection association is for the security appliance. (Only relevant if the security appliance is actually within the network) (see [below for nested schema](#nestedatt--parameters--policies_by_security_appliance))
- `policies_by_ssid` (Attributes) An object, describing the policy-connection associations for each active SSID within the network. Keys should be the number of enabled SSIDs, mapping to an object describing the client's policy (see [below for nested schema](#nestedatt--parameters--policies_by_ssid))

<a id="nestedatt--parameters--clients"></a>
### Nested Schema for `parameters.clients`

Optional:

- `mac` (String) The MAC address of the client. Required.
- `name` (String) The display name for the client. Optional. Limited to 255 bytes.


<a id="nestedatt--parameters--policies_by_security_appliance"></a>
### Nested Schema for `parameters.policies_by_security_appliance`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked' or 'Normal'. Required.


<a id="nestedatt--parameters--policies_by_ssid"></a>
### Nested Schema for `parameters.policies_by_ssid`

Optional:

- `status_0` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_0))
- `status_1` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_1))
- `status_10` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_10))
- `status_11` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_11))
- `status_12` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_12))
- `status_13` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_13))
- `status_14` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_14))
- `status_2` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_2))
- `status_3` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_3))
- `status_4` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_4))
- `status_5` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_5))
- `status_6` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_6))
- `status_7` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_7))
- `status_8` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_8))
- `status_9` (Attributes) The number for the SSID (see [below for nested schema](#nestedatt--parameters--policies_by_ssid--status_9))

<a id="nestedatt--parameters--policies_by_ssid--status_0"></a>
### Nested Schema for `parameters.policies_by_ssid.status_0`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_1"></a>
### Nested Schema for `parameters.policies_by_ssid.status_1`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_10"></a>
### Nested Schema for `parameters.policies_by_ssid.status_10`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_11"></a>
### Nested Schema for `parameters.policies_by_ssid.status_11`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_12"></a>
### Nested Schema for `parameters.policies_by_ssid.status_12`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_13"></a>
### Nested Schema for `parameters.policies_by_ssid.status_13`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_14"></a>
### Nested Schema for `parameters.policies_by_ssid.status_14`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_2"></a>
### Nested Schema for `parameters.policies_by_ssid.status_2`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_3"></a>
### Nested Schema for `parameters.policies_by_ssid.status_3`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_4"></a>
### Nested Schema for `parameters.policies_by_ssid.status_4`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_5"></a>
### Nested Schema for `parameters.policies_by_ssid.status_5`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_6"></a>
### Nested Schema for `parameters.policies_by_ssid.status_6`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_7"></a>
### Nested Schema for `parameters.policies_by_ssid.status_7`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_8"></a>
### Nested Schema for `parameters.policies_by_ssid.status_8`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.


<a id="nestedatt--parameters--policies_by_ssid--status_9"></a>
### Nested Schema for `parameters.policies_by_ssid.status_9`

Optional:

- `device_policy` (String) The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
- `group_policy_id` (String) The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.