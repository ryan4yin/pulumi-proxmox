module github.com/ryan4yin/pulumi-proxmox/provider

go 1.15

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/danitso/terraform-provider-proxmox v0.0.0-20201214100851-ae2ee69b2dac => github.com/ryan4yin/terraform-provider-proxmox v0.0.0-20201214100851-ae2ee69b2dac
)

require (
	github.com/danitso/terraform-provider-proxmox v0.0.0-20201214100851-ae2ee69b2dac
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.0
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
)
