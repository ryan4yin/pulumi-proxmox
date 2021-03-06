// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmox.Permission
{
    public static class GetRole
    {
        public static Task<GetRoleResult> InvokeAsync(GetRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleResult>("proxmox:Permission/getRole:getRole", args ?? new GetRoleArgs(), options.WithVersion());
    }


    public sealed class GetRoleArgs : Pulumi.InvokeArgs
    {
        [Input("roleId", required: true)]
        public string RoleId { get; set; } = null!;

        public GetRoleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRoleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Privileges;
        public readonly string RoleId;

        [OutputConstructor]
        private GetRoleResult(
            string id,

            ImmutableArray<string> privileges,

            string roleId)
        {
            Id = id;
            Privileges = privileges;
            RoleId = roleId;
        }
    }
}
