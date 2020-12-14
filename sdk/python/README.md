# Pulumi Provider for Proxmox

![master branch status](https://github.com/ryan4yin/pulumi-proxmox/workflows/master/badge.svg)
[![PYPI Version](https://img.shields.io/pypi/v/pulumi_proxmox.svg)](https://pypi.org/project/pulumi_proxmox/)

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

## Examples

set pve Environment variables first:

```shell
export PROXMOX_VE_ENDPOINT="https://<server-host>:8006"
export PROXMOX_VE_INSECURE=true
export PROXMOX_VE_USERNAME=root@pam
export PROXMOX_VE_PASSWORD="<password>"
```

Create VirtualMachine using Python SDK:

```python
import os
from pathlib import Path

import pulumi
from pulumi_proxmox.vm import *
from pulumi_proxmox import Provider, ProviderVirtualEnvironmentArgs

# this provider cannot read configuration from EnvVars yet,
# You must manually pass parameters by instantiating a custom provider
proxmox_provider = Provider(
    "proxmox-provider",
    virtual_environment=ProviderVirtualEnvironmentArgs(
        endpoint=os.getenv("PROXMOX_VE_ENDPOINT"),
        insecure=os.getenv("PROXMOX_VE_INSECURE"),
        username=os.getenv("PROXMOX_VE_USERNAME"),
        password=os.getenv("PROXMOX_VE_PASSWORD")
    )
)

# create a virtual machine
VirtualMachine(
    "ubuntu-vm-0",
    name="ubuntu-vm-0",
    description="a ubuntu vm for test",
    node_name="pve",
    # clone from a vm template
    clone=VirtualMachineCloneArgs(
        vm_id=100,  # template's vmId
        full=True,  # full clone, not linked clone
        datastore_id="local-lvm",  # template's datastore
        # node_name="",  # template's node name
    ),

    # resource pool name
    pool_id="test-resource",
    cpu=VirtualMachineCpuArgs(
        cores=2,
        sockets=2,
    ),
    memory=VirtualMachineMemoryArgs(
        dedicated="2048",  # unit: MB
        shared="2048"
    ),
    operating_system=VirtualMachineOperatingSystemArgs(
        type="l26"  # l26: linux2.6-linux5.x
    ),
    agent=VirtualMachineAgentArgs(
        enabled=True,  # enable qemu guest agent
    ),
    disks=[
        VirtualMachineDiskArgs(
            datastore_id="local-lvm",
            size="30",  # unit: GB
        )
    ],
    network_devices=[
        VirtualMachineNetworkDeviceArgs(
            enabled=True,
            bridge="vmbr0",
            model="virtio",
            vlan_id=0,
        )
    ],
    # cloud init configuration
    initialization=VirtualMachineInitializationArgs(
        datastore_id="local-lvm",
        dns=VirtualMachineInitializationDnsArgs(
            # dns servers,
            server="114.114.114.114,8.8.8.8",
        ),
        ip_configs=[
            VirtualMachineInitializationIpConfigArgs(
                ipv4=VirtualMachineInitializationIpConfigIpv4Args(
                    address="192.168.1.111/24",
                    gateway="192.168.1.1"
                )
            )
        ],
        user_account=VirtualMachineInitializationUserAccountArgs(
            # set root's ssh key
            keys=[
                Path("ssh-common.pub").read_text()
            ],
            username="root",
        )
    ),
    
    # use custom provider
    opts=pulumi.ResourceOptions(
        provider=proxmox_provider
    )
)
```

## Reference

please read [danitso/terraform-provider-proxmox](https://github.com/danitso/terraform-provider-proxmox)'s docs for details.

## Developing the Provider

all information about sdks are configured in `provider/resources.go`, if you want to help me, take a look at it.

