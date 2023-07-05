---
subcategory: "Pinpoint"
layout: "aws"
page_title: "AWS: aws_pinpoint_sms_channel"
description: |-
  Use the `aws_pinpoint_sms_channel` resource to manage Pinpoint SMS Channels.
---

# Resource: aws_pinpoint_sms_channel

Use the `aws_pinpoint_sms_channel` resource to manage Pinpoint SMS Channels.

## Example Usage

```terraform
resource "aws_pinpoint_sms_channel" "sms" {
  application_id = aws_pinpoint_app.app.application_id
}

resource "aws_pinpoint_app" "app" {}
```

## Argument Reference

This resource supports the following arguments:

* `application_id` - (Required) ID of the application.
* `enabled` - (Optional) Whether the channel is enabled or disabled. By default, it is set to `true`.
* `sender_id` - (Optional) Identifier of the sender for your messages.
* `short_code` - (Optional) Short Code registered with the phone provider.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `promotional_messages_per_second` - Maximum number of promotional messages that can be sent per second.
* `transactional_messages_per_second` - Maximum number of transactional messages per second that can be sent.

## Import

To import the Pinpoint SMS Channel, use the `application_id`. For example:

```
$ terraform import aws_pinpoint_sms_channel.sms application-id
```
