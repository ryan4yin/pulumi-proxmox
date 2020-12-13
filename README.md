# Pulumi Provider for Proxmox

![master branch status](https://github.com/ryan4yin/pulumi-proxmox/workflows/master/badge.svg)


A Pulumi Provider which adds support for Proxmox solutions.

based on [danitso/terraform-provider-proxmox](https://github.com/danitso/terraform-provider-proxmox), read its docs for details.


### Build the provider:


In order to properly build the sdks, the following tools are expected:
- `pulumictl` (See the project's README for installation instructions: https://github.com/pulumi/pulumictl)

then use the following command to build all sdks:

```shell
make build_sdks
```

## Installing

This package is available in many languages in the standard packaging formats.

first, build and install resource plugin, **this is necessary before you install any language's sdk**:

```shell
make install_resource_plugin
```

then install the language's sdks and every thing will be allright.

**Note**: Installing package directly from the package registry like pypi/npm/nuget is not supported yet, you need to install package from source via `make`.


### Node.js (Java/TypeScript)

```shell
make install_nodejs_sdk
```

### Python

```shell
make install_python_sdk
```

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/ryan4yin/pulumi-proxmox/sdk/go/...


### .NET

To use from .NET, use the following command:

    $ make install_dotnet_sdk

## Configuration

>BUG: cannot read configuration from EnvVars `PROXMOX_VE_ENDPOINT` `PROXMOX_VE_USERNAME` etc.
I'm Working on it now.

In addition to [terraform generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g. `alias` and `version`), the following arguments are supported in the Proxmox `provider` block:

* `virtual_environment` - (Optional) The Proxmox Virtual Environment configuration.
    * `endpoint` - (Required) The endpoint for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_ENDPOINT`).
    * `insecure` - (Optional) Whether to skip the TLS verification step (can also be sourced from `PROXMOX_VE_INSECURE`). If omitted, defaults to `false`.
    * `otp` - (Optional) The one-time password for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_OTP`).
    * `password` - (Required) The password for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_PASSWORD`).
    * `username` - (Required) The username and realm for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_USERNAME`).

## Reference

please read [danitso/terraform-provider-proxmox](https://github.com/danitso/terraform-provider-proxmox)'s docs for details.

## Developing the Provider

all information about sdks are configured in `provider/resources.go`, if you want to help me, take a look at it.

