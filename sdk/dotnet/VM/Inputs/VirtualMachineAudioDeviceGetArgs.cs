// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmox.VM.Inputs
{

    public sealed class VirtualMachineAudioDeviceGetArgs : Pulumi.ResourceArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("driver")]
        public Input<string>? Driver { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public VirtualMachineAudioDeviceGetArgs()
        {
        }
    }
}