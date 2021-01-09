module github.com/ryan4yin/pulumi-proxmox/provider

go 1.15

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/danitso/terraform-provider-proxmox v0.0.0-20210109062334-08d93e539070
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.0
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
)
