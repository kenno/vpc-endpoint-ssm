package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		vpcId := conf.Require("vpcId")
		subnetId := conf.Require("subnetId")
		securityGroupId := conf.Require("securityGroupId")

		awsConfig := config.New(ctx, "aws")
		awsRegion := awsConfig.Require("region")

		// Create a VPC endpoint for ssm
		ssmVpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "ssm", &ec2.VpcEndpointArgs{
			VpcId:           pulumi.String(vpcId),
			ServiceName:     pulumi.String(fmt.Sprintf("com.amazonaws.%v.ssm", awsRegion)),
			VpcEndpointType: pulumi.String("Interface"),
			SecurityGroupIds: pulumi.StringArray{
				pulumi.String(securityGroupId),
			},
			SubnetIds: pulumi.StringArray{
				pulumi.String(subnetId),
			},
			PrivateDnsEnabled: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
				"Created":     pulumi.String("pulumi"),
			},
		})

		if err != nil {
			return err
		}

		// Create a VPC endpoint for ssm
		ec2messagesVpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "ec2messages", &ec2.VpcEndpointArgs{
			VpcId:           pulumi.String(vpcId),
			ServiceName:     pulumi.String(fmt.Sprintf("com.amazonaws.%v.ec2messages", awsRegion)),
			VpcEndpointType: pulumi.String("Interface"),
			SecurityGroupIds: pulumi.StringArray{
				pulumi.String(securityGroupId),
			},
			SubnetIds: pulumi.StringArray{
				pulumi.String(subnetId),
			},
			PrivateDnsEnabled: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
				"Created":     pulumi.String("pulumi"),
			},
		})

		if err != nil {
			return err
		}

		// Create a VPC endpoint for ssm
		ssmmessagesVpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "ssmmessages", &ec2.VpcEndpointArgs{
			VpcId:           pulumi.String(vpcId),
			ServiceName:     pulumi.String(fmt.Sprintf("com.amazonaws.%v.ssmmessages", awsRegion)),
			VpcEndpointType: pulumi.String("Interface"),
			SecurityGroupIds: pulumi.StringArray{
				pulumi.String(securityGroupId),
			},
			SubnetIds: pulumi.StringArray{
				pulumi.String(subnetId),
			},
			PrivateDnsEnabled: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
				"Created":     pulumi.String("pulumi"),
			},
		})

		if err != nil {
			return err
		}

		// Export the name of the 3 VPC endpoints
		ctx.Export("ssmEndpoint", ssmVpcEndpoint.ID())
		ctx.Export("ec2messagesEndpoint", ec2messagesVpcEndpoint.ID())
		ctx.Export("ssmmessagesEndpoint", ssmmessagesVpcEndpoint.ID())
		return nil
	})
}
