---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_webhooks_logs Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_webhooks_logs (Data Source)



## Example Usage

```terraform
data "meraki_organizations_webhooks_logs" "example" {

  ending_before   = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  url             = "string"
}

output "meraki_organizations_webhooks_logs_example" {
  value = data.meraki_organizations_webhooks_logs.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 90 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.
- `url` (String) url query parameter. The URL the webhook was sent to

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationWebhooksLogs (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `alert_type` (String) Type of alert that the webhook is delivering
- `logged_at` (String) When the webhook log was created, in ISO8601 format
- `network_id` (String) Network ID for the webhook log
- `organization_id` (String) ID for the webhook log's organization
- `response_code` (Number) Response code from the webhook
- `response_duration` (Number) Duration of the response, in milliseconds
- `sent_at` (String) When the webhook was sent, in ISO8601 format
- `url` (String) URL where the webhook was sent