---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_insight_monitored_media_servers Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_insight_monitored_media_servers (Resource)



## Example Usage

```terraform
resource "meraki_organizations_insight_monitored_media_servers" "example" {

  address                        = "123.123.123.1"
  best_effort_monitoring_enabled = true
  name                           = "Sample VoIP Provider"
  organization_id                = "string"
}

output "meraki_organizations_insight_monitored_media_servers_example" {
  value = meraki_organizations_insight_monitored_media_servers.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `address` (String) The IP address (IPv4 only) or hostname of the media server to monitor
- `best_effort_monitoring_enabled` (Boolean) Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead.
- `monitored_media_server_id` (String) monitoredMediaServerId path parameter. Monitored media server ID
- `name` (String) The name of the VoIP provider

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_insight_monitored_media_servers.example "monitored_media_server_id,organization_id"
```