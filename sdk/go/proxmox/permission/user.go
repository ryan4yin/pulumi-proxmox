// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package permission

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type User struct {
	pulumi.CustomResourceState

	// The access control list
	Acls UserAclArrayOutput `pulumi:"acls"`
	// The user comment
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// The user's email address
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// Whether the user account is enabled
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The user account's expiration date
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// The user's first name
	FirstName pulumi.StringPtrOutput `pulumi:"firstName"`
	// The user's groups
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// The user's keys
	Keys pulumi.StringPtrOutput `pulumi:"keys"`
	// The user's last name
	LastName pulumi.StringPtrOutput `pulumi:"lastName"`
	// The user's password
	Password pulumi.StringOutput `pulumi:"password"`
	// The user id
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("proxmox:Permission/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("proxmox:Permission/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The access control list
	Acls []UserAcl `pulumi:"acls"`
	// The user comment
	Comment *string `pulumi:"comment"`
	// The user's email address
	Email *string `pulumi:"email"`
	// Whether the user account is enabled
	Enabled *bool `pulumi:"enabled"`
	// The user account's expiration date
	ExpirationDate *string `pulumi:"expirationDate"`
	// The user's first name
	FirstName *string `pulumi:"firstName"`
	// The user's groups
	Groups []string `pulumi:"groups"`
	// The user's keys
	Keys *string `pulumi:"keys"`
	// The user's last name
	LastName *string `pulumi:"lastName"`
	// The user's password
	Password *string `pulumi:"password"`
	// The user id
	UserId *string `pulumi:"userId"`
}

type UserState struct {
	// The access control list
	Acls UserAclArrayInput
	// The user comment
	Comment pulumi.StringPtrInput
	// The user's email address
	Email pulumi.StringPtrInput
	// Whether the user account is enabled
	Enabled pulumi.BoolPtrInput
	// The user account's expiration date
	ExpirationDate pulumi.StringPtrInput
	// The user's first name
	FirstName pulumi.StringPtrInput
	// The user's groups
	Groups pulumi.StringArrayInput
	// The user's keys
	Keys pulumi.StringPtrInput
	// The user's last name
	LastName pulumi.StringPtrInput
	// The user's password
	Password pulumi.StringPtrInput
	// The user id
	UserId pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The access control list
	Acls []UserAcl `pulumi:"acls"`
	// The user comment
	Comment *string `pulumi:"comment"`
	// The user's email address
	Email *string `pulumi:"email"`
	// Whether the user account is enabled
	Enabled *bool `pulumi:"enabled"`
	// The user account's expiration date
	ExpirationDate *string `pulumi:"expirationDate"`
	// The user's first name
	FirstName *string `pulumi:"firstName"`
	// The user's groups
	Groups []string `pulumi:"groups"`
	// The user's keys
	Keys *string `pulumi:"keys"`
	// The user's last name
	LastName *string `pulumi:"lastName"`
	// The user's password
	Password string `pulumi:"password"`
	// The user id
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The access control list
	Acls UserAclArrayInput
	// The user comment
	Comment pulumi.StringPtrInput
	// The user's email address
	Email pulumi.StringPtrInput
	// Whether the user account is enabled
	Enabled pulumi.BoolPtrInput
	// The user account's expiration date
	ExpirationDate pulumi.StringPtrInput
	// The user's first name
	FirstName pulumi.StringPtrInput
	// The user's groups
	Groups pulumi.StringArrayInput
	// The user's keys
	Keys pulumi.StringPtrInput
	// The user's last name
	LastName pulumi.StringPtrInput
	// The user's password
	Password pulumi.StringInput
	// The user id
	UserId pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}
