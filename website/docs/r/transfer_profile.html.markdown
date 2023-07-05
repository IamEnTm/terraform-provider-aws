---
subcategory: "Transfer Family"
layout: "aws"
page_title: "AWS: aws_transfer_profile"
description: |-
  Provides a AWS Transfer AS2 Profile Resource
---

# Resource: aws_transfer_profile

Provides a AWS Transfer AS2 Profile resource.

## Example Usage

### Basic

```terraform
resource "aws_transfer_profile" "example" {
  as2_id          = "example"
  certificate_ids = [aws_transfer_certificate.example.certificate_id]
  usage           = "LOCAL"
}
```

## Argument Reference

The following arguments are supported:

* `as2_id` - (Required) The As2Id is the AS2 name as defined in the RFC 4130. For inbound ttransfers this is the AS2 From Header for the AS2 messages sent from the partner. For Outbound messages this is the AS2 To Header for the AS2 messages sent to the partner. his ID cannot include spaces.
* `certificate_ids` - (Optional) The list of certificate Ids from the imported certificate operation.
* `profile_type` - (Required) The profile type should be LOCAL or PARTNER.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `profile_id`  - The unique identifier for the AS2 profile

## Import

Transfer AS2 Profile can be imported using the `profile_id`, e.g.,

```
$ terraform import aws_transfer_profile.example p-4221a88afd5f4362a
```
