// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetescertmanager

import (
	"context"
	"reflect"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Automates the management and issuance of TLS certificates from various issuing sources within Kubernetes
type CertManager struct {
	pulumi.ResourceState

	// Detailed information about the status of the underlying Helm deployment.
	Status ReleaseStatusOutput `pulumi:"status"`
}

// NewCertManager registers a new resource with the given unique name, arguments, and options.
func NewCertManager(ctx *pulumi.Context,
	name string, args *CertManagerArgs, opts ...pulumi.ResourceOption) (*CertManager, error) {
	if args == nil {
		args = &CertManagerArgs{}
	}

	var resource CertManager
	err := ctx.RegisterRemoteComponentResource("kubernetes-cert-manager:index:CertManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type certManagerArgs struct {
	Affinity   *corev1.Affinity       `pulumi:"affinity"`
	Cainjector *CertManagerCaInjector `pulumi:"cainjector"`
	// Override the namespace used to store DNS provider credentials etc. for ClusterIssuer resources. By default, the same namespace as cert-manager is deployed within is used. This namespace will not be automatically created by the Helm chart.
	ClusterResourceNamespace *string `pulumi:"clusterResourceNamespace"`
	// Container Security Context to be set on the controller component container. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	ContainerSecurityContext *corev1.SecurityContext `pulumi:"containerSecurityContext"`
	// Optional additional annotations to add to the controller Deployment
	DeploymentAnnotations map[string]string `pulumi:"deploymentAnnotations"`
	// Optional additional arguments.
	ExtraArgs         []string             `pulumi:"extraArgs"`
	ExtraEnv          []corev1.EnvVar      `pulumi:"extraEnv"`
	ExtraVolumeMounts []corev1.VolumeMount `pulumi:"extraVolumeMounts"`
	ExtraVolumes      []corev1.Volume      `pulumi:"extraVolumes"`
	// Comma separated list of feature gates that should be enabled on the controller pod.
	FeatureGates *string            `pulumi:"featureGates"`
	Global       *CertManagerGlobal `pulumi:"global"`
	// HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions  *Release                `pulumi:"helmOptions"`
	Http_proxy   *string                 `pulumi:"http_proxy"`
	Https_proxy  *string                 `pulumi:"https_proxy"`
	Image        *CertManagerImage       `pulumi:"image"`
	IngressShim  *CertManagerIngressShim `pulumi:"ingressShim"`
	InstallCRDs  *bool                   `pulumi:"installCRDs"`
	No_proxy     []string                `pulumi:"no_proxy"`
	NodeSelector *corev1.NodeSelector    `pulumi:"nodeSelector"`
	// Optional additional annotations to add to the controller Pods
	PodAnnotations map[string]string    `pulumi:"podAnnotations"`
	PodDnsConfig   *corev1.PodDNSConfig `pulumi:"podDnsConfig"`
	// Optional DNS settings, useful if you have a public and private DNS zone for the same domain on Route 53. What follows is an example of ensuring cert-manager can access an ingress or DNS TXT records at all times. NOTE: This requires Kubernetes 1.10 or `CustomPodDNS` feature gate enabled for the cluster to work.
	PodDnsPolicy *string                      `pulumi:"podDnsPolicy"`
	PodLabels    map[string]string            `pulumi:"podLabels"`
	Prometheus   *CertManagerPrometheus       `pulumi:"prometheus"`
	ReplicaCount *int                         `pulumi:"replicaCount"`
	Resources    *corev1.ResourceRequirements `pulumi:"resources"`
	// Pod Security Context. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SecurityContext *corev1.PodSecurityContext `pulumi:"securityContext"`
	ServiceAccount  *CertManagerServiceAccount `pulumi:"serviceAccount"`
	// Optional additional annotations to add to the controller service
	ServiceAnnotations map[string]string `pulumi:"serviceAnnotations"`
	// Optional additional labels to add to the controller Service
	ServiceLabels   map[string]string           `pulumi:"serviceLabels"`
	Startupapicheck *CertManagerStartupAPICheck `pulumi:"startupapicheck"`
	Strategy        *appsv1.DeploymentStrategy  `pulumi:"strategy"`
	Tolerations     []corev1.Toleration         `pulumi:"tolerations"`
	Webhook         *CertManagerWebhook         `pulumi:"webhook"`
}

// The set of arguments for constructing a CertManager resource.
type CertManagerArgs struct {
	Affinity   corev1.AffinityPtrInput
	Cainjector CertManagerCaInjectorPtrInput
	// Override the namespace used to store DNS provider credentials etc. for ClusterIssuer resources. By default, the same namespace as cert-manager is deployed within is used. This namespace will not be automatically created by the Helm chart.
	ClusterResourceNamespace pulumi.StringPtrInput
	// Container Security Context to be set on the controller component container. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	ContainerSecurityContext corev1.SecurityContextPtrInput
	// Optional additional annotations to add to the controller Deployment
	DeploymentAnnotations pulumi.StringMapInput
	// Optional additional arguments.
	ExtraArgs         pulumi.StringArrayInput
	ExtraEnv          corev1.EnvVarArrayInput
	ExtraVolumeMounts corev1.VolumeMountArrayInput
	ExtraVolumes      corev1.VolumeArrayInput
	// Comma separated list of feature gates that should be enabled on the controller pod.
	FeatureGates pulumi.StringPtrInput
	Global       CertManagerGlobalPtrInput
	// HelmOptions is an escape hatch that lets the end user control any aspect of the Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions  ReleasePtrInput
	Http_proxy   pulumi.StringPtrInput
	Https_proxy  pulumi.StringPtrInput
	Image        CertManagerImagePtrInput
	IngressShim  CertManagerIngressShimPtrInput
	InstallCRDs  pulumi.BoolPtrInput
	No_proxy     pulumi.StringArrayInput
	NodeSelector corev1.NodeSelectorPtrInput
	// Optional additional annotations to add to the controller Pods
	PodAnnotations pulumi.StringMapInput
	PodDnsConfig   corev1.PodDNSConfigPtrInput
	// Optional DNS settings, useful if you have a public and private DNS zone for the same domain on Route 53. What follows is an example of ensuring cert-manager can access an ingress or DNS TXT records at all times. NOTE: This requires Kubernetes 1.10 or `CustomPodDNS` feature gate enabled for the cluster to work.
	PodDnsPolicy pulumi.StringPtrInput
	PodLabels    pulumi.StringMapInput
	Prometheus   CertManagerPrometheusPtrInput
	ReplicaCount pulumi.IntPtrInput
	Resources    corev1.ResourceRequirementsPtrInput
	// Pod Security Context. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SecurityContext corev1.PodSecurityContextPtrInput
	ServiceAccount  CertManagerServiceAccountPtrInput
	// Optional additional annotations to add to the controller service
	ServiceAnnotations pulumi.StringMapInput
	// Optional additional labels to add to the controller Service
	ServiceLabels   pulumi.StringMapInput
	Startupapicheck CertManagerStartupAPICheckPtrInput
	Strategy        appsv1.DeploymentStrategyPtrInput
	Tolerations     corev1.TolerationArrayInput
	Webhook         CertManagerWebhookPtrInput
}

func (CertManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certManagerArgs)(nil)).Elem()
}

type CertManagerInput interface {
	pulumi.Input

	ToCertManagerOutput() CertManagerOutput
	ToCertManagerOutputWithContext(ctx context.Context) CertManagerOutput
}

func (*CertManager) ElementType() reflect.Type {
	return reflect.TypeOf((**CertManager)(nil)).Elem()
}

func (i *CertManager) ToCertManagerOutput() CertManagerOutput {
	return i.ToCertManagerOutputWithContext(context.Background())
}

func (i *CertManager) ToCertManagerOutputWithContext(ctx context.Context) CertManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertManagerOutput)
}

// CertManagerArrayInput is an input type that accepts CertManagerArray and CertManagerArrayOutput values.
// You can construct a concrete instance of `CertManagerArrayInput` via:
//
//	CertManagerArray{ CertManagerArgs{...} }
type CertManagerArrayInput interface {
	pulumi.Input

	ToCertManagerArrayOutput() CertManagerArrayOutput
	ToCertManagerArrayOutputWithContext(context.Context) CertManagerArrayOutput
}

type CertManagerArray []CertManagerInput

func (CertManagerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertManager)(nil)).Elem()
}

func (i CertManagerArray) ToCertManagerArrayOutput() CertManagerArrayOutput {
	return i.ToCertManagerArrayOutputWithContext(context.Background())
}

func (i CertManagerArray) ToCertManagerArrayOutputWithContext(ctx context.Context) CertManagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertManagerArrayOutput)
}

// CertManagerMapInput is an input type that accepts CertManagerMap and CertManagerMapOutput values.
// You can construct a concrete instance of `CertManagerMapInput` via:
//
//	CertManagerMap{ "key": CertManagerArgs{...} }
type CertManagerMapInput interface {
	pulumi.Input

	ToCertManagerMapOutput() CertManagerMapOutput
	ToCertManagerMapOutputWithContext(context.Context) CertManagerMapOutput
}

type CertManagerMap map[string]CertManagerInput

func (CertManagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertManager)(nil)).Elem()
}

func (i CertManagerMap) ToCertManagerMapOutput() CertManagerMapOutput {
	return i.ToCertManagerMapOutputWithContext(context.Background())
}

func (i CertManagerMap) ToCertManagerMapOutputWithContext(ctx context.Context) CertManagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertManagerMapOutput)
}

type CertManagerOutput struct{ *pulumi.OutputState }

func (CertManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertManager)(nil)).Elem()
}

func (o CertManagerOutput) ToCertManagerOutput() CertManagerOutput {
	return o
}

func (o CertManagerOutput) ToCertManagerOutputWithContext(ctx context.Context) CertManagerOutput {
	return o
}

// Detailed information about the status of the underlying Helm deployment.
func (o CertManagerOutput) Status() ReleaseStatusOutput {
	return o.ApplyT(func(v *CertManager) ReleaseStatusOutput { return v.Status }).(ReleaseStatusOutput)
}

type CertManagerArrayOutput struct{ *pulumi.OutputState }

func (CertManagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertManager)(nil)).Elem()
}

func (o CertManagerArrayOutput) ToCertManagerArrayOutput() CertManagerArrayOutput {
	return o
}

func (o CertManagerArrayOutput) ToCertManagerArrayOutputWithContext(ctx context.Context) CertManagerArrayOutput {
	return o
}

func (o CertManagerArrayOutput) Index(i pulumi.IntInput) CertManagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertManager {
		return vs[0].([]*CertManager)[vs[1].(int)]
	}).(CertManagerOutput)
}

type CertManagerMapOutput struct{ *pulumi.OutputState }

func (CertManagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertManager)(nil)).Elem()
}

func (o CertManagerMapOutput) ToCertManagerMapOutput() CertManagerMapOutput {
	return o
}

func (o CertManagerMapOutput) ToCertManagerMapOutputWithContext(ctx context.Context) CertManagerMapOutput {
	return o
}

func (o CertManagerMapOutput) MapIndex(k pulumi.StringInput) CertManagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertManager {
		return vs[0].(map[string]*CertManager)[vs[1].(string)]
	}).(CertManagerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertManagerInput)(nil)).Elem(), &CertManager{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertManagerArrayInput)(nil)).Elem(), CertManagerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertManagerMapInput)(nil)).Elem(), CertManagerMap{})
	pulumi.RegisterOutputType(CertManagerOutput{})
	pulumi.RegisterOutputType(CertManagerArrayOutput{})
	pulumi.RegisterOutputType(CertManagerMapOutput{})
}
