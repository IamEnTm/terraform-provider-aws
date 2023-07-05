---
subcategory: "Chime SDK Voice"
layout: "aws"
page_title: "AWS: aws_chimesdkvoice_global_settings"
description: |-
  Terraform resource for managing Amazon Chime SDK Voice Global Settings.
---

# Resource: aws_chimesdkvoice_global_settings

Terraform resource for managing Amazon Chime SDK Voice Global Settings.

## Example Usage

### Basic Usage

```terraform
resource "aws_chimesdkvoice_global_settings" "example" {
  voice_connector {
    cdr_bucket = "example-bucket-name"
  }
}
```

## Argument Reference

The following arguments are supported:

* `voice_connector` - (Required) The Voice Connector settings. See [voice_connector](#voice_connector).

### `voice_connector`

The Amazon Chime SDK Voice Connector settings. Includes any Amazon S3 buckets designated for storing call detail records.

* `cdr_bucket` - (Optional) The S3 bucket that stores the Voice Connector's call detail records.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - AWS account ID for which the settings are applied.

## Import

AWS Chime SDK Voice Global Settings can be imported using the `id` (AWS account ID), e.g.,

```
$ terraform import aws_chimesdkvoice_global_settings.example 123456789012
```
