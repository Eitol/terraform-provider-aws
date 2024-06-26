---
subcategory: "Chatbot"
layout: "aws"
page_title: "AWS: aws_chatbot_slack_workspace"
description: |-
  Terraform data source for managing an AWS Chatbot Slack Workspace.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_chatbot_slack_workspace

Terraform data source for managing an AWS Chatbot Slack Workspace.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_chatbot_slack_workspace import DataAwsChatbotSlackWorkspace
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, slackTeamName):
        super().__init__(scope, name)
        DataAwsChatbotSlackWorkspace(self, "example",
            team_slack_name="abc",
            slack_team_name=slack_team_name
        )
```

## Argument Reference

The following arguments are required:

* `slack_team_name` - (Required) Slack workspace name configured with AWS Chabot.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `slack_team_id` - ID of the Slack Workspace assigned by AWS Chatbot.

<!-- cache-key: cdktf-0.20.1 input-837eb56025e58d9b743ec29c05dba308029138b944df2bd4dcfcb96892e2f34c -->