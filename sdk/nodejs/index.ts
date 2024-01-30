// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { CertManagerArgs } from "./certManager";
export type CertManager = import("./certManager").CertManager;
export const CertManager: typeof import("./certManager").CertManager = null as any;
utilities.lazyLoad(exports, ["CertManager"], () => require("./certManager"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as types from "./types";

export {
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "kubernetes-cert-manager:index:CertManager":
                return new CertManager(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("kubernetes-cert-manager", "index", _module)
pulumi.runtime.registerResourcePackage("kubernetes-cert-manager", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:kubernetes-cert-manager") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
