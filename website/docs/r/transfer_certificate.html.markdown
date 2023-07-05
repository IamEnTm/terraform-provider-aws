---
subcategory: "Transfer Family"
layout: "aws"
page_title: "AWS: aws_transfer_certificate"
description: |-
  Provides a AWS Transfer AS2 Certificate Resource
---

# Resource: aws_transfer_certificate

Provides a AWS Transfer AS2 Certificate resource.

## Example Usage

### Basic

```terraform
resource "aws_transfer_certificate" "example" {
  certificate       = file("${path.module}/example.com/example.crt")
  certificate_chain = file("${path.module}/example.com/ca.crt")
  private_key       = file("${path.module}/example.com/example.key")
  description       = "example"
  usage             = "SIGNING"
}
```

## Argument Reference

The following arguments are supported:

* `certificate` - (Required) The valid certificate file required for the transfer.
* `certificate_chain` - (Optional) The optional list of certificate that make up the chain for the certificate that is being imported.
* `description` - (Optional) A short description that helps identify the certificate.
* `private_key` - (Optional) The private key associated with the certificate being imported.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `usage` - (Required) Specifies if a certificate is being used for signing or encryption. The valid values are SIGNING and ENCRYPTION.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `certificate_id` - The unique identifier for the AS2 certificate
* `active_date` - An date when the certificate becomes active
* `inactive_date` - An date when the certificate becomes inactive

## Import

Transfer AS2 Certificate can be imported using the `certificate_id`, e.g.,

```
$ terraform import aws_transfer_certificate.example c-4221a88afd5f4362a
```
