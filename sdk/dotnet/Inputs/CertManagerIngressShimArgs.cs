// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ChartCertManager.Inputs
{

    public sealed class CertManagerIngressShimArgs : Pulumi.ResourceArgs
    {
        [Input("defaultIssuerGroup")]
        public Input<string>? DefaultIssuerGroup { get; set; }

        [Input("defaultIssuerKind")]
        public Input<string>? DefaultIssuerKind { get; set; }

        [Input("defaultIssuerName")]
        public Input<string>? DefaultIssuerName { get; set; }

        public CertManagerIngressShimArgs()
        {
        }
    }
}