<h1 align="center">
  AWS CLI tool for multi account
</h1>

<p align="center">
<img alt="GitHub" src="https://img.shields.io/github/license/clover0/maws?style=for-the-badge">
<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/clover0/maws?style=for-the-badge">
</p>

<p align="center">
Search for AWS resources across AWS multi account.
</p>


# Quickstart
## Install
1. Install AWS CLI v2. ([Install Guide](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html))
2. Install via Homebrew
```
brew tap clover0/maws
brew install maws
```

# Command

Format
```
maws <maws options> <aws cli command>
```

## Options
- `--profile-filter <keyword>`
  - Filter AWS profiles in aws config (~/.aws/config)



## Examples
`~/.aws/config`

```
[profile Account1.AdministratorAccess]
sso_start_url = https://d-xxx.awsapps.com/start
sso_region = us-east-1
sso_account_name = Account1
sso_account_id = 12345678910 
sso_role_name = AdministratorAccess
region = ap-northeast-1

[profile Account2.AdministratorAccess]
sso_start_url = https://d-xxx.awsapps.com/start
sso_region = us-east-1

[profile Account3.AdministratorAccess]
...
```

command
```
maws --profile-filter=Admin ec2 describe-vpcs --filter "Name=vpc-id,Values=xxx123"
```

Find profiles including "Admin" from `~/.aws/config`.
And execute aws command every profile(account).
