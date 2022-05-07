# terraform-cloud-tag-orchestration

# Purpose
This tool was written as an example of how to use the [go-tfe](https://github.com/hashicorp/go-tfe) library to search workspaces using tags. It has limited error handeling but could be used as a starting point for someone looking to explore automating [Terraform Cloud](https://app.terraform.io) in [Go](https://github.com/golang/go).

In the context of this readme, TFE, TFC and Terraform Cloud are used interchangeably.

## Requirements
Required environment variables
| Env | Description |
|-----|-------------|
| TFC_ORGANIZATION | Terraform Cloud organization |
| TFC_ORGANIZATION_TOKEN | [Terraform Cloud organization token](https://www.terraform.io/cloud-docs/users-teams-organizations/users#api-tokens) |

## Usage
```
$ export TFC_ORGANIZATION="your...tfc...org...here"
$ export TFC_ORGANIZATION_TOKEN="your...tfc...token...here"
$ ./tfcinfo {space} {seperated} {tags}
workspace-1-matching-tags
workspace-2-matching-tags
```

### Example
```
$ ./tfcinfo aws demo      
Terraform cloud orchestration with tags
Organization: devopstower
Searching for workspaces with the following tags: aws,demo
tfc-aws-network-prod
tfc-aws-network-dev
```

## Build
```
go build -o tfcinfo .
```