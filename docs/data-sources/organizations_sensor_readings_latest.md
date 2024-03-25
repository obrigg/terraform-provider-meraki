---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_sensor_readings_latest Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_sensor_readings_latest (Data Source)



## Example Usage

```terraform
data "meraki_organizations_sensor_readings_latest" "example" {

  ending_before   = "string"
  metrics         = ["string"]
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  serials         = ["string"]
  starting_after  = "string"
}

output "meraki_organizations_sensor_readings_latest_example" {
  value = data.meraki_organizations_sensor_readings_latest.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `metrics` (List of String) metrics query parameter. Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
- `network_ids` (List of String) networkIds query parameter. Optional parameter to filter readings by network.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 100. Default is 100.
- `serials` (List of String) serials query parameter. Optional parameter to filter readings by sensor.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

### Read-Only

- `items` (Attributes List) Array of ResponseSensorGetOrganizationSensorReadingsLatest (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `network` (Attributes) Network to which the sensor belongs. (see [below for nested schema](#nestedatt--items--network))
- `readings` (Attributes Set) Array of latest readings from the sensor. Each object represents a single reading for a single metric. (see [below for nested schema](#nestedatt--items--readings))
- `serial` (String) Serial number of the sensor that took the readings.

<a id="nestedatt--items--network"></a>
### Nested Schema for `items.network`

Read-Only:

- `id` (String) ID of the network.
- `name` (String) Name of the network.


<a id="nestedatt--items--readings"></a>
### Nested Schema for `items.readings`

Read-Only:

- `battery` (Attributes) Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'. (see [below for nested schema](#nestedatt--items--readings--battery))
- `button` (Attributes) Reading for the 'button' metric. This will only be present if the 'metric' property equals 'button'. (see [below for nested schema](#nestedatt--items--readings--button))
- `door` (Attributes) Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'. (see [below for nested schema](#nestedatt--items--readings--door))
- `humidity` (Attributes) Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'. (see [below for nested schema](#nestedatt--items--readings--humidity))
- `indoor_air_quality` (Attributes) Reading for the 'indoorAirQuality' metric. This will only be present if the 'metric' property equals 'indoorAirQuality'. (see [below for nested schema](#nestedatt--items--readings--indoor_air_quality))
- `metric` (String) Type of sensor reading.
- `noise` (Attributes) Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'. (see [below for nested schema](#nestedatt--items--readings--noise))
- `pm25` (Attributes) Reading for the 'pm25' metric. This will only be present if the 'metric' property equals 'pm25'. (see [below for nested schema](#nestedatt--items--readings--pm25))
- `temperature` (Attributes) Reading for the 'temperature' metric. This will only be present if the 'metric' property equals 'temperature'. (see [below for nested schema](#nestedatt--items--readings--temperature))
- `ts` (String) Time at which the reading occurred, in ISO8601 format.
- `tvoc` (Attributes) Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'. (see [below for nested schema](#nestedatt--items--readings--tvoc))
- `water` (Attributes) Reading for the 'water' metric. This will only be present if the 'metric' property equals 'water'. (see [below for nested schema](#nestedatt--items--readings--water))

<a id="nestedatt--items--readings--battery"></a>
### Nested Schema for `items.readings.battery`

Read-Only:

- `percentage` (Number) Remaining battery life.


<a id="nestedatt--items--readings--button"></a>
### Nested Schema for `items.readings.button`

Read-Only:

- `press_type` (String) Type of button press that occurred.


<a id="nestedatt--items--readings--door"></a>
### Nested Schema for `items.readings.door`

Read-Only:

- `open` (Boolean) True if the door is open.


<a id="nestedatt--items--readings--humidity"></a>
### Nested Schema for `items.readings.humidity`

Read-Only:

- `relative_percentage` (Number) Humidity reading in %RH.


<a id="nestedatt--items--readings--indoor_air_quality"></a>
### Nested Schema for `items.readings.indoor_air_quality`

Read-Only:

- `score` (Number) Indoor air quality score between 0 and 100.


<a id="nestedatt--items--readings--noise"></a>
### Nested Schema for `items.readings.noise`

Read-Only:

- `ambient` (Attributes) Ambient noise reading. (see [below for nested schema](#nestedatt--items--readings--noise--ambient))

<a id="nestedatt--items--readings--noise--ambient"></a>
### Nested Schema for `items.readings.noise.ambient`

Read-Only:

- `level` (Number) Ambient noise reading in adjusted decibels.



<a id="nestedatt--items--readings--pm25"></a>
### Nested Schema for `items.readings.pm25`

Read-Only:

- `concentration` (Number) PM2.5 reading in micrograms per cubic meter.


<a id="nestedatt--items--readings--temperature"></a>
### Nested Schema for `items.readings.temperature`

Read-Only:

- `celsius` (Number) Temperature reading in degrees Celsius.
- `fahrenheit` (Number) Temperature reading in degrees Fahrenheit.


<a id="nestedatt--items--readings--tvoc"></a>
### Nested Schema for `items.readings.tvoc`

Read-Only:

- `concentration` (Number) TVOC reading in micrograms per cubic meter.


<a id="nestedatt--items--readings--water"></a>
### Nested Schema for `items.readings.water`

Read-Only:

- `present` (Boolean) True if water is detected.