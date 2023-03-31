# module-template
Template to develop a custom Terraform module.

## Requisites
- [terraform-docs](https://github.com/segmentio/terraform-docs) - This is used in our pre-commit hook in order to generate documentation from Terraform modules in various output formats.
- [pre-commit](https://pre-commit.com/#install)

---------------------

## Prerequisites

### You need to remove Terraform if they already installed in your system, the following versions are used for this project

- Tarraform: 1.2.0
- Go: 1.18

### Installation Steps (MacOS)

```sh
brew install pre-commit gawk terraform-docs coreutils tfenv awscli jq cfn-lint
```

### **tfenv**

---------------------

##### List current Terraform versions installed on your system

```sh
tfenv list
```

##### Install a specific version of Terraform

```sh
tfenv install 1.2.0
```

##### Select Terraform version to be used, this could be used to switch between versions

```sh
tfenv use 1.2.0
```

### **pre-commit**

---------------------

#### Pre-Commit Usage

- [Pre-Commit documentation](https://pre-commit.com/)
- [Hook documentation](https://github.com/antonbabenko/pre-commit-terraform)

You must `git add` your files before the pre-commit hook will run against them.

##### Check the version

```sh
pre-commit --version
```

##### Install the git hook scripts

```sh
pre-commit install
pre-commit install-hooks
```

##### Run against all files (`git add` must be run first)

```sh
pre-commit run -a
```

#### **Terraform**

---------------------

Prerequisites:

- You can successfully authenticate to AWS via CLI using or other CLI auth tool (like granted.dev) or your preferred method of authentication.
- You have the correct version of Terraform installed through `tfenv`

When developing modules, it is easier to run these commands from your example directory where you have already defined an example for your tests to run against.

Navigate to the directory where you would like to run your terraform configuration, authenticate to AWS through the CLI (optionally through Granted)

```hcl
terraform init (install modules, both local and external)
terraform validate (validate your configuration will not error before the plan/apply stage)
terraform plan (check what you're going to provision)
terraform apply (deploy the infrastructure)
terraform destroy (destroy the infrastructure)
```

#### **Go**

---------------------

```sh

# Update and Install Go (for a specfic version, append @{version}, like `brew install golang@1.18`)
brew update && brew install go@1.18

# Following Go best practices, create 3 new directories ($HOME/go/bin, $HOME/go/src, $HOME/go/pkg)
mkdir -p $HOME/go/{bin,src,pkg}

# Set important environment variables
# Add the below to your .bashrc, or .zshrc
export GOPATH=$HOME/go
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"

# If you're on an M1 mac, make Go play nicely with Rosetta
export GODEBUG=asyncpreemptoff=1

# source your shell
source $HOME/.bashrc (or .zshrc)
```

## **Terratest - Recommended Test Practices**

---------------------

## Use table driven tests

- These tests are a fairly standard practice
- They let you clearly and easily create multiple test-cases for a single test
- They are defined as an array of structs, where fields of the struct are variables for each test case

## When defining essential variables, **USE** hard-coded fixed fields

- Repository examples should hard-code fixed fields, which are essential to the spirit of the example

## When defining field variables to be tested, **DO NOT USE** use hard-coded fixed fields

- Repository examples should use variables for fields which are to be tested (i.e., so that they can be fed in via terratest)

## Useful packages

- [testify/assert library](https://github.com/stretchr/testify/assert) -- assertion library to assert test results against expected fields
- [aws api helper library](https://github.com/gruntwork-io/terratest/modules/aws) -- library to help us query the AWS API directly

## Helpful terminal commands

_You must be authenticated to the target AWS account before executing the below commands_

```sh
# run all tests cases with no verbose output (not recommended, as you can't see errors)
go test

# apply verbose output and extend the timeout past the default of 10m (helpful for tests that need longer to run -- like AWS RDS examples)
go test -v -timeout 30m

# run a single test, be sure the test case matches the regex TestSimpleDynamoDb
go test run TestSimpleTest

# print the test output to a file
go test -v -timeout 30m | tee ~/Desktop/module_terratest_output.txt

```

#### Supporting Documentation

- [Terratest documentation](https://terratest.gruntwork.io/docs/#getting-started)

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | 1.2.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >=4.28 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_some"></a> [some](#provider\_some) | n/a |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_foo"></a> [foo](#module\_foo) | ../../ | n/a |

## Resources

| Name | Type |
|------|------|
| [some_aws_resource.bar](https://registry.terraform.io/providers/hashicorp/some/latest/docs/resources/aws_resource) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_bar"></a> [bar](#input\_bar) | value | `map(any)` | `{}` | no |
| <a name="input_baz"></a> [baz](#input\_baz) | value | `list(string)` | <pre>[<br>  "value"<br>]</pre> | no |
| <a name="input_foo"></a> [foo](#input\_foo) | value | `string` | `""` | no |

## Outputs

No outputs.
<!-- END_TF_DOCS -->
