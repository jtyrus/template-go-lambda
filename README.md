### Go Lambda Template
Sample app currently being used to serve some DynamoDB data. This can be run from Github Actions or locally to deploy a lambda, necessary permissions, and a table. Can be repurposed for other lambdas in the future. Note that "template" is used loosely, this is really a full application. 

#### Prerequisites
* Terraform
* Go
* AWS OIDC (if deploying through Github Actions)

#### Terraform Setup
Set a valid backend in `terraform/terraform.tf`. Backend block can be removed if running terraform locally.

#### Github Setup
The `Assume Deployment Role` needs a valid role for whatever repo this template is used in. [OIDC Setup](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services)

#### Checklist of things to change
* Naming
  * Makefile, .tf files, and code all have references to `portfolio` that should be changed/removed
* Terraform
  * Remove dynamo resources if not necessary
  * Update the backend-config in `terraform.tf`
  * Update the backend-config key var in the Github Actions