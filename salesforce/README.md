# CDKTF Go bindings for hashicorp/salesforce provider version 0.1.0

HashiCorp made the decision to stop publishing new versions of prebuilt [Terraform salesforce provider](https://registry.terraform.io/providers/hashicorp/salesforce/0.1.0) bindings for [CDK for Terraform](https://cdk.tf) on January 12, 2024. As such, this repository has been archived and is no longer supported in any way by HashiCorp. Previously-published versions of this prebuilt provider will still continue to be available on their respective package managers (e.g. npm, PyPi, Maven, NuGet), but these will not be compatible with new releases of `cdktf` past `0.20.0` and are no longer eligible for commercial support.

As a reminder, you can continue to use the `hashicorp/salesforce` provider in your CDK for Terraform (CDKTF) projects, even with newer versions of CDKTF, but you will need to generate the bindings locally. The easiest way to do so is to use the [`provider add` command](https://developer.hashicorp.com/terraform/cdktf/cli-reference/commands#provider-add), optionally with the `--force-local` flag enabled:

`cdktf provider add hashicorp/salesforce --force-local`

For more information and additional examples, check out our documentation on [generating provider bindings manually](https://cdk.tf/imports).

## Deprecated Package

The go package is generated into the [`github.com/cdktf/cdktf-provider-salesforce-go`](https://github.com/cdktf/cdktf-provider-salesforce-go) package.

`go get github.com/cdktf/cdktf-provider-salesforce-go/salesforce`

## Docs

Find auto-generated docs for this provider [here](https://github.com/cdktf/cdktf-provider-salesforce/blob/main/docs/API.go.md).

