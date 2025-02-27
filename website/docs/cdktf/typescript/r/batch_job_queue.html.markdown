---
subcategory: "Batch"
layout: "aws"
page_title: "AWS: aws_batch_job_queue"
description: |-
  Provides a Batch Job Queue resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_batch_job_queue

Provides a Batch Job Queue resource.

## Example Usage

### Basic Job Queue

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BatchJobQueue } from "./.gen/providers/aws/batch-job-queue";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new BatchJobQueue(this, "test_queue", {
      computeEnvironments: [testEnvironment1.arn, testEnvironment2.arn],
      name: "tf-test-batch-job-queue",
      priority: 1,
      state: "ENABLED",
    });
  }
}

```

### Job Queue with a fair share scheduling policy

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { BatchJobQueue } from "./.gen/providers/aws/batch-job-queue";
import { BatchSchedulingPolicy } from "./.gen/providers/aws/batch-scheduling-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new BatchSchedulingPolicy(this, "example", {
      fairSharePolicy: {
        computeReservation: 1,
        shareDecaySeconds: 3600,
        shareDistribution: [
          {
            shareIdentifier: "A1*",
            weightFactor: 0.1,
          },
        ],
      },
      name: "example",
    });
    const awsBatchJobQueueExample = new BatchJobQueue(this, "example_1", {
      computeEnvironments: [testEnvironment1.arn, testEnvironment2.arn],
      name: "tf-test-batch-job-queue",
      priority: 1,
      schedulingPolicyArn: example.arn,
      state: "ENABLED",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsBatchJobQueueExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) Specifies the name of the job queue.
* `computeEnvironments` - (Required) Specifies the set of compute environments
    mapped to a job queue and their order.  The position of the compute environments
    in the list will dictate the order.
* `priority` - (Required) The priority of the job queue. Job queues with a higher priority
    are evaluated first when associated with the same compute environment.
* `schedulingPolicyArn` - (Optional) The ARN of the fair share scheduling policy. If this parameter is specified, the job queue uses a fair share scheduling policy. If this parameter isn't specified, the job queue uses a first in, first out (FIFO) scheduling policy. After a job queue is created, you can replace but can't remove the fair share scheduling policy.
* `state` - (Required) The state of the job queue. Must be one of: `enabled` or `disabled`
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name of the job queue.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Batch Job Queue using the `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
  }
}

```

Using `terraform import`, import Batch Job Queue using the `arn`. For example:

```console
% terraform import aws_batch_job_queue.test_queue arn:aws:batch:us-east-1:123456789012:job-queue/sample
```

<!-- cache-key: cdktf-0.18.0 input-622d814f33122b3abb4dbe424e099f99dc14c870fec376cc7645af13df622dbd -->