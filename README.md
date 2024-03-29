# terraform-provider-mgtt

## Table of Contents

- [Description](#description)
- [References](#references)
- [Getting started](#getting-started)

## Description

Sample repository implementing a Terraform provider designed for managing resources in Azure and AWS.

## References

- [Writing Custom Terraform Providers](https://www.hashicorp.com/blog/writing-custom-terraform-providers). This link is deprecated but helpful to understand the basic concepts from the HashiCorp founder.
- [How to develop/test existing provider locally?](https://github.com/hashicorp/terraform-provider-aws/issues/5396)
- [terraform-provider-klayer Github repository](https://github.com/ldcorentin/terraform-provider-klayer). **13-01-2023** - Latest practical example
- [HashiCorp documentation on Terraform Plugin SDKv2](https://developer.hashicorp.com/terraform/plugin/sdkv2?collectionSlug=providers&productSlug=terraform). Latest documentation on Terraform Plugin SDKv2 for writing custom providers.
- [Schema Attributes and Types](https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas/schema-types). Check the section on `TypeSet`.

## Getting started

### Precondition

- Install the [Docker engine]()
- In VS Code install the `Dev Containers` extension in order to utilize dev containers
- You can optionally install tools like [go](https://go.dev/doc/install) and [terraform](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli) yourself instead of utilizing a dev container 

### (Optional) Check CLI tools serving as a basis for implementing resources

CLI tools can be found here:

- [README.md in regards to azure_storage_account CLI tool](./api-testing/azure_storage_account/README.md)
- [README.md in regards to aws_s3_bucket CLI tool](./api-testing/aws_s3_bucket/README.md)

### Exporting environment variables

Create from the [secrets.template.cfg](./templates/secrets.template.cfg) a secrets.cfg file in the project root directory and replace the `PLACEHOLDER_*` values. Some [tests](./mgtt/test/) need environment variables listed in the secrets.cfg file.

```sh
source secrets.cfg
```

### In order to run tests

```sh
source secrets.cfg
# go test ./...
cd mgtt/test
go test
```

**NOTE:** For tests managing Azure objects refer to following [README.md](./api-testing/azure_storage_account/README.md). Topics on how to retrieve a fresh `AZURE_ACCESS_TOKEN` should be covered.

### Compile custom provider

On modern Windows OS (version 10 or 11) run: 

```sh
go build -o terraform-provider-mgtt.exe
```

On Unix systems (MacOS, Linux Ubuntu >=20.04, debian >=11) run: 

```sh
go build -o terraform-provider-mgtt
```

### Copy provider executable to plugins directory 

Refer to [How to develop/ test existing provider locally?](https://github.com/hashicorp/terraform-provider-aws/issues/5396).

Navigate to and remove :

```sh
cd terraform

# Remove previous build artifacts
rm .terraform # Powershell
rm -rf .terraform # Unix terminals
rm  terraform.tfstate terraform.tfstate.backup
```

On modern Windows OS with amd64 CPU architecture run:

```sh
mkdir -p .terraform\plugins\terraform-mgtt.com\mgttprovider\mgtt\1.0.0\windows_amd64
cp ..\terraform-provider-mgtt.exe .terraform\plugins\terraform-mgtt.com\mgttprovider\mgtt\1.0.0\windows_amd64
```

On Unix systems run:

```sh
# Linux Ubuntu >=18.04 or debian >=11 with amd64 CPU architecture
mkdir -p .terraform/plugins/terraform-mgtt.com/mgttprovider/mgtt/1.0.0/linux_amd64
cp ../terraform-provider-mgtt .terraform/plugins/terraform-mgtt.com/mgttprovider/mgtt/1.0.0/linux_amd64

# MacOS with amd64 CPU architecture 
mkdir -p .terraform/plugins/terraform-mgtt.com/mgttprovider/mgtt/1.0.0/darwin_amd64
cp ../terraform-provider-mgtt .terraform/plugins/terraform-mgtt.com/mgttprovider/mgtt/1.0.0/darwin_amd64
```

### Test provider executable with hcl files and terraform commands

```sh
source secrets.cfg
cd terraform
terraform init -plugin-dir="./.terraform/plugins/"
terraform plan
# It is important to set the the DEBUG log level before applying `teraform apply` or `teraform destroy` 
export TF_LOG=DEBUG
terraform apply --auto-approve # initial create requests
# Update input variables in tf files, e.g. 
# ```hcl
# resource "mgtt_azurerm_rg" "this" {
#     name = "rg-test-9000"
#     location = "West Europe"
# }
# ```
terraform apply --auto-approve # update requests 
terraform destroy --auto-approve # delete requests
```
