## Requirements

- [tfenv](https://github.com/tfutils/tfenv) - This is used in order to manage different Terraform versions
- [terraform-docs](https://github.com/segmentio/terraform-docs) - This is used in our pre-commit hook in order to generate documentation from Terraform modules in various output formats.
- [pre-commit](https://pre-commit.com/#install)-configuration to run code standardization (terraform fmt) and documentation (terraform docs) automation on `git commit`
- [Granted](https://docs.commonfate.io/granted/getting-started) (optional) - tooling to help assume your SSO role into an AWS account
- Public cloud provider access credentials (if not using Granted)

---------------------

## Prerequisites

This codebase uses the following Terraform and Golang versions.  Code has not be tested and verified to work with any other versions other than what is listed below:

- Terraform: 1.2.0
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

You must `git add .` your files before the pre-commit hook will run against them.

##### Check the version

```sh
pre-commit --version
```

##### Install the git hook scripts

```sh
pre-commit install
```

##### Run against all files (`git add` must be run first)

```sh
pre-commit run -a
```

## Authenticate to AWS environments with Granted (Optional)

---------------------

Below are steps to ensure easy access to AWS environments by assuming your SSO role with [Granted](https://docs.commonfate.io/granted/getting-started)

```sh
# install with Homebrew
brew tap common-fate/granted
brew install granted

# verify installation
➜ granted -v

Granted v0.3.0
```

### **Setup your AWS profile**

Follow the steps as outlined by executing the command in your terminal: `aws configure sso`

```sh
aws configure sso
> SSO start URL [None]: <Start URL> (the redirect URL after you login to AWS Control Tower through an SSO provider)
> SSO Region: us-west-2

# after the above values are entered your browser will open to and have you confirm
# a few prompts.  Allow the authorize request.

#When successful, you will see a "Request Approved" AWS Modal in your browser tab.

# Go back to the Terminal session to finish the prompts

# Pick the account to which you wish to create a profile, you may also see all accounts to which you have access.

# Pick your assigned role, this will vary based on your organization
# Default CLI region: us-west-2
# Output format: JSON

# CLI Profile name: you can keep what is generated (not recommended) or use something explicit to the environment, like "shared-services-admin"

```

Test your credentials

```sh
$ assume
 Please select the profile you would like to assume:  [Use arrows to move, type to filter]
> shared-services-admin (this is the profile you created in the previous step)

# You will then see a message like: 
[shared-services-admin](us-west-2) session credentials will expire 2022-09-27 14:15:48 -0400 EDT
>

# check to be sure you can query sts and receive your assume role arn back
$ aws sts get-caller-identiy
> 
{
    "UserId": "(redacted):chris.gonzalez@caylent.com",
    "Account": "1111111111",
    "Arn": "arn:aws:sts::1111111111:assumed-role/AWSReservedSSO_AWSPowerUserAccess_(redacted)/chris.gonzalez@caylent.com"
}

# if the above commands are successful, then you can now use Terraform or Run terrestest from your local machine
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
to test run eks-example-cluster

# print the test output to a file
go test -v -timeout 30m | tee ~/Desktop/module_terratest_output.txt
```

#### Supporting Documentation

- [Terratest documentation](https://terratest.gruntwork.io/docs/#getting-started)

### **Terraform**

---------------------

Prerequisites:

- You can successfully authenticate to AWS via CLI using Granted or your preferred method of authentication.
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
