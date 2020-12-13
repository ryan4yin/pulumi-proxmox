# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['File']


class File(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 datastore_id: Optional[pulumi.Input[str]] = None,
                 node_name: Optional[pulumi.Input[str]] = None,
                 source_file: Optional[pulumi.Input[pulumi.InputType['FileSourceFileArgs']]] = None,
                 source_raw: Optional[pulumi.Input[pulumi.InputType['FileSourceRawArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a File resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_type: The content type
        :param pulumi.Input[str] datastore_id: The datastore id
        :param pulumi.Input[str] node_name: The node name
        :param pulumi.Input[pulumi.InputType['FileSourceFileArgs']] source_file: The source file
        :param pulumi.Input[pulumi.InputType['FileSourceRawArgs']] source_raw: The raw source
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['content_type'] = content_type
            if datastore_id is None:
                raise TypeError("Missing required property 'datastore_id'")
            __props__['datastore_id'] = datastore_id
            if node_name is None:
                raise TypeError("Missing required property 'node_name'")
            __props__['node_name'] = node_name
            __props__['source_file'] = source_file
            __props__['source_raw'] = source_raw
            __props__['file_modification_date'] = None
            __props__['file_name'] = None
            __props__['file_size'] = None
            __props__['file_tag'] = None
        super(File, __self__).__init__(
            'proxmox:Storage/file:File',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            datastore_id: Optional[pulumi.Input[str]] = None,
            file_modification_date: Optional[pulumi.Input[str]] = None,
            file_name: Optional[pulumi.Input[str]] = None,
            file_size: Optional[pulumi.Input[int]] = None,
            file_tag: Optional[pulumi.Input[str]] = None,
            node_name: Optional[pulumi.Input[str]] = None,
            source_file: Optional[pulumi.Input[pulumi.InputType['FileSourceFileArgs']]] = None,
            source_raw: Optional[pulumi.Input[pulumi.InputType['FileSourceRawArgs']]] = None) -> 'File':
        """
        Get an existing File resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_type: The content type
        :param pulumi.Input[str] datastore_id: The datastore id
        :param pulumi.Input[str] file_modification_date: The file modification date
        :param pulumi.Input[str] file_name: The file name
        :param pulumi.Input[int] file_size: The file size in bytes
        :param pulumi.Input[str] file_tag: The file tag
        :param pulumi.Input[str] node_name: The node name
        :param pulumi.Input[pulumi.InputType['FileSourceFileArgs']] source_file: The source file
        :param pulumi.Input[pulumi.InputType['FileSourceRawArgs']] source_raw: The raw source
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content_type"] = content_type
        __props__["datastore_id"] = datastore_id
        __props__["file_modification_date"] = file_modification_date
        __props__["file_name"] = file_name
        __props__["file_size"] = file_size
        __props__["file_tag"] = file_tag
        __props__["node_name"] = node_name
        __props__["source_file"] = source_file
        __props__["source_raw"] = source_raw
        return File(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        """
        The content type
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="datastoreId")
    def datastore_id(self) -> pulumi.Output[str]:
        """
        The datastore id
        """
        return pulumi.get(self, "datastore_id")

    @property
    @pulumi.getter(name="fileModificationDate")
    def file_modification_date(self) -> pulumi.Output[str]:
        """
        The file modification date
        """
        return pulumi.get(self, "file_modification_date")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Output[str]:
        """
        The file name
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter(name="fileSize")
    def file_size(self) -> pulumi.Output[int]:
        """
        The file size in bytes
        """
        return pulumi.get(self, "file_size")

    @property
    @pulumi.getter(name="fileTag")
    def file_tag(self) -> pulumi.Output[str]:
        """
        The file tag
        """
        return pulumi.get(self, "file_tag")

    @property
    @pulumi.getter(name="nodeName")
    def node_name(self) -> pulumi.Output[str]:
        """
        The node name
        """
        return pulumi.get(self, "node_name")

    @property
    @pulumi.getter(name="sourceFile")
    def source_file(self) -> pulumi.Output[Optional['outputs.FileSourceFile']]:
        """
        The source file
        """
        return pulumi.get(self, "source_file")

    @property
    @pulumi.getter(name="sourceRaw")
    def source_raw(self) -> pulumi.Output[Optional['outputs.FileSourceRaw']]:
        """
        The raw source
        """
        return pulumi.get(self, "source_raw")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
