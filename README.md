# Pulumi Provider for Proxmox

![master branch status](https://github.com/ryan4yin/pulumi-proxmox/workflows/master/badge.svg)
[![PYPI Version](https://img.shields.io/pypi/v/pulumi_proxmox.svg)](https://pypi.org/project/pulumi_proxmox/)

A Pulumi Provider which adds support for Proxmox solutions.

based on [danitso/terraform-provider-proxmox](https://github.com/danitso/terraform-provider-proxmox), read its docs for details.

## TODO

- [ ] fix Bug: cannot read configuration from EnvVars `PROXMOX_VE_ENDPOINT` `PROXMOX_VE_USERNAME` etc.
- [ ] fix github actions, build and upload resource plugin to github releases automatically.
- [ ] set resource plugin download url(github releases).


## Installation

1. Resources plugin for Linux are available as tarballs in the [release](https://github.com/ryan4yin/pulumi-proxmox/releases) page.
1. Follow the installation instructions in releases to install the resource plugin and the python sdk.
1. for other languages, ​​please install sdk from source in the `sdk` folder.

## Build and Install the provider From Source

In order to properly build the sdks, the following tools are expected:
- `pulumictl` (See the project's README for installation instructions: https://github.com/pulumi/pulumictl)

to build all the sdks, you need install and set up all the 4 language sdks first: go/dotnet/python/nodejs.

then use the following command to build the resource plugin and all the sdks:

```shell
make build_sdks
```

### Install Resource Plugin 

first, build and install resource plugin:

```shell
make install_resource_plugin
```

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

Create VirtualMachine using Python SDK(writing in other languages ​​is almost the same):

```python
import os
from pathlib import Path

import pulumi
from pulumi_proxmox import Provider, ProviderVirtualEnvironmentArgs
from pulumi_proxmox.vm import *

# this provider cannot read configuration from Environment variables yet,
# You must manually pass parameters by instantiating a custom provider
proxmox_provider = Provider(
    "proxmox-provider",
    virtual_environment=ProviderVirtualEnvironmentArgs(
        endpoint=os.getenv("PROXMOX_VE_ENDPOINT"),
        insecure=os.getenv("PROXMOX_VE_INSECURE") == "true",
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
    on_boot=True,  # start the vm during system bootup
    reboot=False,  # reboot the vm after it was created successfully
    started=True,  # start the vm after it was created successfully
    # clone from a vm template
    clone=VirtualMachineCloneArgs(
        vm_id=9000,  # template's vmId
        full=True,  # full clone, not linked clone
        datastore_id="local-lvm",  # template's datastore
        node_name="pve",  # template's node name
    ),

    # resource pool name
    pool_id="test-resource",
    cpu=VirtualMachineCpuArgs(
        cores=2,
        sockets=2,
        type="kvm64",  # set it to kvm64 for better vm migration
    ),
    memory=VirtualMachineMemoryArgs(
        dedicated="4096",  # unit: MB
        shared="4096"
    ),
    operating_system=VirtualMachineOperatingSystemArgs(
        type="l26"  # l26: linux2.6-linux5.x
    ),
    agent=VirtualMachineAgentArgs(
        # please confirm you have qemu-guest-agent in your vm before enable this!
        # otherwise this may cause the vm to fail to shutdown/reboot!
        enabled=False,
        timeout="60s",  # timeout
    ),
    disks=[
        VirtualMachineDiskArgs(
            interface="scsi0",
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
        type="nocloud",  # 'nocloud' for linux,  'configdrive2' for windows
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
            password="chage_me",  # needed when login from console
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


## Related Projects

pulumi providers:

- [Matchlighter/pulumi-proxmoxve](https://github.com/Matchlighter/pulumi-proxmoxve)

terraform providers:

- [danitso/terraform-provider-proxmox](https://github.com/danitso/terraform-provider-proxmox)
- [Telmate/terraform-provider-proxmox](https://github.com/Telmate/terraform-provider-proxmox)
