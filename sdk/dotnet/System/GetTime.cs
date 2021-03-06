// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmox.System
{
    public static class GetTime
    {
        public static Task<GetTimeResult> InvokeAsync(GetTimeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTimeResult>("proxmox:System/getTime:getTime", args ?? new GetTimeArgs(), options.WithVersion());
    }


    public sealed class GetTimeArgs : Pulumi.InvokeArgs
    {
        [Input("nodeName", required: true)]
        public string NodeName { get; set; } = null!;

        public GetTimeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTimeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LocalTime;
        public readonly string NodeName;
        public readonly string TimeZone;
        public readonly string UtcTime;

        [OutputConstructor]
        private GetTimeResult(
            string id,

            string localTime,

            string nodeName,

            string timeZone,

            string utcTime)
        {
            Id = id;
            LocalTime = localTime;
            NodeName = nodeName;
            TimeZone = timeZone;
            UtcTime = utcTime;
        }
    }
}
