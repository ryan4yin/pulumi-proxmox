module github.com/ryan4yin/pulumi-proxmox/provider

go 1.15

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/danitso/terraform-provider-proxmox v0.0.0-20201213042502-4597175e5cbc => github.com/ryan4yin/terraform-provider-proxmox v0.0.0-20201213042502-4597175e5cbc
)

require (
	github.com/danitso/terraform-provider-proxmox v0.0.0-20201213042502-4597175e5cbc
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.0
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
)
