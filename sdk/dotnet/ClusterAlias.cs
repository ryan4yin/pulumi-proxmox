// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Proxmox
{
    public partial class ClusterAlias : Pulumi.CustomResource
    {
        /// <summary>
        /// IP/CIDR block
        /// </summary>
        [Output("cidr")]
        public Output<string> Cidr { get; private set; } = null!;

        /// <summary>
        /// Alias comment
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Alias name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterAlias(string name, ClusterAliasArgs args, CustomResourceOptions? options = null)
            : base("proxmox:index/clusterAlias:ClusterAlias", name, args ?? new ClusterAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterAlias(string name, Input<string> id, ClusterAliasState? state = null, CustomResourceOptions? options = null)
            : base("proxmox:index/clusterAlias:ClusterAlias", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterAlias Get(string name, Input<string> id, ClusterAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterAlias(name, id, state, options);
        }
    }

    public sealed class ClusterAliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP/CIDR block
        /// </summary>
        [Input("cidr", required: true)]
        public Input<string> Cidr { get; set; } = null!;

        /// <summary>
        /// Alias comment
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Alias name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ClusterAliasArgs()
        {
        }
    }

    public sealed class ClusterAliasState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP/CIDR block
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// Alias comment
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Alias name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ClusterAliasState()
        {
        }
    }
}
