// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxmox

import (
	"unicode"

	// "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/danitso/terraform-provider-proxmox/proxmoxtf"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "proxmox"
	// modules:
	mainMod = "index" // the main module

	vmMod        = "VM"
	containerMod = "CT"

	systemMod     = "System"
	permissionMod = "Permission"

	storageMod = "Storage"
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv1.NewProvider(proxmoxtf.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "proxmox",
		Description: "A Pulumi Provider which adds support for Proxmox solutions.",
		Keywords:    []string{"pulumi", "proxmox"},
		License:     "Apache-2.0",
		Homepage:    "https://github.com/ryan4yin/pulumi-proxmox",
		Repository:  "https://github.com/ryan4yin/pulumi-proxmox",
		// TODO resource plugin download url.
		// PluginDownloadURL: "https://github.com/ryan4yin/pulumi-proxmox/releases/download/v0.0.5/pulumi-resource-proxmox-v0.0.5-linux_amd64.tgz",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here
			// TODO cannot read configuration from EnvVars!
			"endpoint": {
				Type: makeType("endpoint", "Endpoint"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PROXMOX_VE_ENDPOINT", "PM_VE_ENDPOINT"},
				},
			},
			"insecure": {
				Type: makeType("insecure", "Insecure"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PROXMOX_VE_INSECURE", "PM_VE_INSECURE"},
				},
			},
			"otp": {
				Type: makeType("otp", "OTP"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PROXMOX_VE_OTP", "PM_VE_OTP"},
				},
			},
			"password": {
				Type: makeType("password", "Password"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PROXMOX_VE_PASSWORD", "PM_VE_PASSWORD"},
				},
			},
			"username": {
				Type: makeType("username", "Username"),
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PROXMOX_VE_USERNAME", "PM_VE_USERNAME"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type.
			// the resourceMap of terraform provider: https://github.com/danitso/terraform-provider-proxmox/blob/release-0.4.0/proxmoxtf/provider.go
			"proxmox_virtual_environment_cluster_alias":   {Tok: makeResource(mainMod, "ClusterAlias")},
			"proxmox_virtual_environment_cluster_aliases": {Tok: makeResource(mainMod, "ClusterAliases")},

			"proxmox_virtual_environment_vm":        {Tok: makeResource(vmMod, "VirtualMachine")},
			"proxmox_virtual_environment_container": {Tok: makeResource(containerMod, "Container")},

			// Storage
			"proxmox_virtual_environment_file": {Tok: makeResource(storageMod, "File")},

			// System
			"proxmox_virtual_environment_dns":         {Tok: makeResource(systemMod, "DNS")},
			"proxmox_virtual_environment_certificate": {Tok: makeResource(systemMod, "Certifi")},
			"proxmox_virtual_environment_hosts":       {Tok: makeResource(systemMod, "Hosts")},
			"proxmox_virtual_environment_time":        {Tok: makeResource(systemMod, "Time")},

			// Permission
			"proxmox_virtual_environment_user":  {Tok: makeResource(permissionMod, "User")},
			"proxmox_virtual_environment_group": {Tok: makeResource(permissionMod, "Group")},
			"proxmox_virtual_environment_pool":  {Tok: makeResource(permissionMod, "Pool")},
			"proxmox_virtual_environment_role":  {Tok: makeResource(permissionMod, "Role")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"proxmox_virtual_environment_cluster_alias":   {Tok: makeDataSource(mainMod, "getClusterAlias")},
			"proxmox_virtual_environment_cluster_aliases": {Tok: makeDataSource(mainMod, "getClusterAliases")},

			// Storage
			"proxmox_virtual_environment_datastores": {Tok: makeDataSource(storageMod, "getDatastores")},

			// Map each resource in the Terraform provider to a Pulumi function.
			"proxmox_virtual_environment_version": {Tok: makeDataSource(mainMod, "getVersion")},
			"proxmox_virtual_environment_nodes":   {Tok: makeDataSource(mainMod, "getNodes")},

			// System Configuration
			"proxmox_virtual_environment_dns":   {Tok: makeDataSource(systemMod, "getDNS")},
			"proxmox_virtual_environment_time":  {Tok: makeDataSource(systemMod, "getTime")},
			"proxmox_virtual_environment_hosts": {Tok: makeDataSource(systemMod, "getHosts")},

			// Permissions
			"proxmox_virtual_environment_users": {Tok: makeDataSource(permissionMod, "getUsers")},
			"proxmox_virtual_environment_user":  {Tok: makeDataSource(permissionMod, "getUser")},

			"proxmox_virtual_environment_group":  {Tok: makeDataSource(permissionMod, "getGroup")},
			"proxmox_virtual_environment_groups": {Tok: makeDataSource(permissionMod, "getGroups")},

			"proxmox_virtual_environment_pool":  {Tok: makeDataSource(permissionMod, "getPool")},
			"proxmox_virtual_environment_pools": {Tok: makeDataSource(permissionMod, "getPools")},

			"proxmox_virtual_environment_role":  {Tok: makeDataSource(permissionMod, "getRole")},
			"proxmox_virtual_environment_roles": {Tok: makeDataSource(permissionMod, "getRoles")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.9.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
