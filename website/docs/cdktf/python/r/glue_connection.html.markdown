---
subcategory: "Glue"
layout: "aws"
page_title: "AWS: aws_glue_connection"
description: |-
  Provides an Glue Connection resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_glue_connection

Provides a Glue Connection resource.

## Example Usage

### Non-VPC Connection

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_connection import GlueConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlueConnection(self, "example",
            connection_properties={
                "JDBC_CONNECTION_URL": "jdbc:mysql://example.com/exampledatabase",
                "PASSWORD": "examplepassword",
                "USERNAME": "exampleusername"
            },
            name="example"
        )
```

### Non-VPC Connection with secret manager reference

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws. import DataAwsSecretmanagerSecret
from imports.aws.glue_connection import GlueConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsSecretmanagerSecret(self, "example",
            name="example-secret"
        )
        aws_glue_connection_example = GlueConnection(self, "example_1",
            connection_properties={
                "JDBC_CONNECTION_URL": "jdbc:mysql://example.com/exampledatabase",
                "SECRET_ID": Token.as_string(example.name)
            },
            name="example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_glue_connection_example.override_logical_id("example")
```

### VPC Connection

For more information, see the [AWS Documentation](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html#connection-JDBC-VPC).

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.glue_connection import GlueConnection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GlueConnection(self, "example",
            connection_properties={
                "JDBC_CONNECTION_URL": "jdbc:mysql://${" + aws_rds_cluster_example.endpoint + "}/exampledatabase",
                "PASSWORD": "examplepassword",
                "USERNAME": "exampleusername"
            },
            name="example",
            physical_connection_requirements=GlueConnectionPhysicalConnectionRequirements(
                availability_zone=Token.as_string(aws_subnet_example.availability_zone),
                security_group_id_list=[Token.as_string(aws_security_group_example.id)],
                subnet_id=Token.as_string(aws_subnet_example.id)
            )
        )
```

## Argument Reference

This resource supports the following arguments:

* `catalog_id` – (Optional) The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
* `connection_properties` – (Optional) A map of key-value pairs used as parameters for this connection.
* `connection_type` – (Optional) The type of the connection. Supported are: `CUSTOM`, `JDBC`, `KAFKA`, `MARKETPLACE`, `MONGODB`, and `NETWORK`. Defaults to `JBDC`.
* `description` – (Optional) Description of the connection.
* `match_criteria` – (Optional) A list of criteria that can be used in selecting this connection.
* `name` – (Required) The name of the connection.
* `physical_connection_requirements` - (Optional) A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### physical_connection_requirements

* `availability_zone` - (Optional) The availability zone of the connection. This field is redundant and implied by `subnet_id`, but is currently an api requirement.
* `security_group_id_list` - (Optional) The security group ID list used by the connection.
* `subnet_id` - (Optional) The subnet ID used by the connection.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Catalog ID and name of the connection
* `arn` - The ARN of the Glue Connection.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Glue Connections using the `CATALOG-ID` (AWS account ID if not custom) and `NAME`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
```

Using `terraform import`, import Glue Connections using the `CATALOG-ID` (AWS account ID if not custom) and `NAME`. For example:

```console
% terraform import aws_glue_connection.MyConnection 123456789012:MyConnection
```

<!-- cache-key: cdktf-0.18.0 input-e486b071d2d1ce5f1b7d77ebb71283d9a58d67d688ff421d5f7cc4203ff5c73d -->