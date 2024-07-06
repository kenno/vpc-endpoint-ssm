# Create SSM VPC endpoints (in Go)

This code deloys 3 VPC endpoints for SSM Session Managers.

## Deploying the App

To deploy your infrastructure, follow the below steps.

### Prerequisites

1. [Install Go](https://golang.org/doc/install)
2. [Install Pulumi](https://www.pulumi.com/docs/get-started/install/)
3. [Configure AWS Credentials](https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/)

### Steps

To deploy your infrastructure, follow the below steps.

1. Next, create a new Pulumi stack:

    ```bash
    $ pulumi stack init
    ```

2. Set the required configuration variables for this program:

    ```bash
    $ pulumi config set aws:region ap-southeast-2
    $ pulumi config set vpcId vpc-ZZZZZZZZZZZ
    $ pulumi config set subnetId subnet-yyyyyyyyyy
    $ pulumi config set securityGroupId sg-xxxxxxxxx
    ```

3. Stand up the resources in the stack:

    ```bash
    $ pulumi up
    ```

4. To destroy the stack and remove it:

    ```bash
    $ pulumi destroy --yes
    $ pulumi stack rm --yes
    ```
