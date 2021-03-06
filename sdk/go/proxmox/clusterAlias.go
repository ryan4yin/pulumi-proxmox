// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package proxmox

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClusterAlias struct {
	pulumi.CustomResourceState

	// IP/CIDR block
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// Alias comment
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Alias name
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewClusterAlias registers a new resource with the given unique name, arguments, and options.
func NewClusterAlias(ctx *pulumi.Context,
	name string, args *ClusterAliasArgs, opts ...pulumi.ResourceOption) (*ClusterAlias, error) {
	if args == nil || args.Cidr == nil {
		return nil, errors.New("missing required argument 'Cidr'")
	}
	if args == nil {
		args = &ClusterAliasArgs{}
	}
	var resource ClusterAlias
	err := ctx.RegisterResource("proxmox:index/clusterAlias:ClusterAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterAlias gets an existing ClusterAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterAliasState, opts ...pulumi.ResourceOption) (*ClusterAlias, error) {
	var resource ClusterAlias
	err := ctx.ReadResource("proxmox:index/clusterAlias:ClusterAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterAlias resources.
type clusterAliasState struct {
	// IP/CIDR block
	Cidr *string `pulumi:"cidr"`
	// Alias comment
	Comment *string `pulumi:"comment"`
	// Alias name
	Name *string `pulumi:"name"`
}

type ClusterAliasState struct {
	// IP/CIDR block
	Cidr pulumi.StringPtrInput
	// Alias comment
	Comment pulumi.StringPtrInput
	// Alias name
	Name pulumi.StringPtrInput
}

func (ClusterAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAliasState)(nil)).Elem()
}

type clusterAliasArgs struct {
	// IP/CIDR block
	Cidr string `pulumi:"cidr"`
	// Alias comment
	Comment *string `pulumi:"comment"`
	// Alias name
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ClusterAlias resource.
type ClusterAliasArgs struct {
	// IP/CIDR block
	Cidr pulumi.StringInput
	// Alias comment
	Comment pulumi.StringPtrInput
	// Alias name
	Name pulumi.StringPtrInput
}

func (ClusterAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAliasArgs)(nil)).Elem()
}
