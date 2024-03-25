---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_appliance_vpn_third_party_vpnpeers Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_appliance_vpn_third_party_vpnpeers (Resource)



## Example Usage

```terraform
resource "meraki_organizations_appliance_vpn_third_party_vpnpeers" "example" {

  organization_id = "string"
  peers = [{

    ike_version = "2"
    ipsec_policies = {

      child_auth_algo          = ["sha1"]
      child_cipher_algo        = ["aes128"]
      child_lifetime           = 28800
      child_pfs_group          = ["disabled"]
      ike_auth_algo            = ["sha1"]
      ike_cipher_algo          = ["tripledes"]
      ike_diffie_hellman_group = ["group2"]
      ike_lifetime             = 28800
      ike_prf_algo             = ["prfsha1"]
    }
    ipsec_policies_preset = "default"
    local_id              = "myMXId@meraki.com"
    name                  = "Peer Name"
    network_tags          = ["none"]
    private_subnets       = ["192.168.1.0/24", "192.168.128.0/24"]
    public_ip             = "123.123.123.1"
    remote_id             = "miles@meraki.com"
    secret                = "Sample Password"
  }]
}

output "meraki_organizations_appliance_vpn_third_party_vpnpeers_example" {
  value = meraki_organizations_appliance_vpn_third_party_vpnpeers.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `peers` (Attributes Set) The list of VPN peers (see [below for nested schema](#nestedatt--peers))

<a id="nestedatt--peers"></a>
### Nested Schema for `peers`

Optional:

- `ike_version` (String) [optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
- `ipsec_policies` (Attributes) Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used. (see [below for nested schema](#nestedatt--peers--ipsec_policies))
- `ipsec_policies_preset` (String) One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.
- `local_id` (String) [optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.
- `name` (String) The name of the VPN peer
- `network_tags` (List of String) A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].
- `private_subnets` (List of String) The list of the private subnets of the VPN peer
- `public_ip` (String) [optional] The public IP of the VPN peer
- `remote_id` (String) [optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.
- `secret` (String) The shared secret with the VPN peer

<a id="nestedatt--peers--ipsec_policies"></a>
### Nested Schema for `peers.ipsec_policies`

Optional:

- `child_auth_algo` (List of String) This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
- `child_cipher_algo` (List of String) This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'
- `child_lifetime` (Number) The lifetime of the Phase 2 SA in seconds.
- `child_pfs_group` (List of String) This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'
- `ike_auth_algo` (List of String) This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'
- `ike_cipher_algo` (List of String) This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'
- `ike_diffie_hellman_group` (List of String) This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'
- `ike_lifetime` (Number) The lifetime of the Phase 1 SA in seconds.
- `ike_prf_algo` (List of String) [optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_appliance_vpn_third_party_vpnpeers.example "organization_id"
```