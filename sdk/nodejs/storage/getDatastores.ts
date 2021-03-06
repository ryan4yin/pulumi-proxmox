// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDatastores(args: GetDatastoresArgs, opts?: pulumi.InvokeOptions): Promise<GetDatastoresResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("proxmox:Storage/getDatastores:getDatastores", {
        "nodeName": args.nodeName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatastores.
 */
export interface GetDatastoresArgs {
    readonly nodeName: string;
}

/**
 * A collection of values returned by getDatastores.
 */
export interface GetDatastoresResult {
    readonly actives: boolean[];
    readonly contentTypes: string[][];
    readonly datastoreIds: string[];
    readonly enableds: boolean[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nodeName: string;
    readonly shareds: boolean[];
    readonly spaceAvailables: number[];
    readonly spaceTotals: number[];
    readonly spaceUseds: number[];
    readonly types: number[];
}
