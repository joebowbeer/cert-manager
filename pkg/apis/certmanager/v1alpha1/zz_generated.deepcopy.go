// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMECertificateConfig).DeepCopyInto(out.(*ACMECertificateConfig))
			return nil
		}, InType: reflect.TypeOf(&ACMECertificateConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMECertificateDNS01Config).DeepCopyInto(out.(*ACMECertificateDNS01Config))
			return nil
		}, InType: reflect.TypeOf(&ACMECertificateDNS01Config{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMECertificateDomainConfig).DeepCopyInto(out.(*ACMECertificateDomainConfig))
			return nil
		}, InType: reflect.TypeOf(&ACMECertificateDomainConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMECertificateHTTP01Config).DeepCopyInto(out.(*ACMECertificateHTTP01Config))
			return nil
		}, InType: reflect.TypeOf(&ACMECertificateHTTP01Config{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEDomainAuthorization).DeepCopyInto(out.(*ACMEDomainAuthorization))
			return nil
		}, InType: reflect.TypeOf(&ACMEDomainAuthorization{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEIssuer).DeepCopyInto(out.(*ACMEIssuer))
			return nil
		}, InType: reflect.TypeOf(&ACMEIssuer{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEIssuerDNS01Config).DeepCopyInto(out.(*ACMEIssuerDNS01Config))
			return nil
		}, InType: reflect.TypeOf(&ACMEIssuerDNS01Config{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEIssuerDNS01Provider).DeepCopyInto(out.(*ACMEIssuerDNS01Provider))
			return nil
		}, InType: reflect.TypeOf(&ACMEIssuerDNS01Provider{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEIssuerDNS01ProviderCloudDNS).DeepCopyInto(out.(*ACMEIssuerDNS01ProviderCloudDNS))
			return nil
		}, InType: reflect.TypeOf(&ACMEIssuerDNS01ProviderCloudDNS{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEIssuerDNS01ProviderCloudflare).DeepCopyInto(out.(*ACMEIssuerDNS01ProviderCloudflare))
			return nil
		}, InType: reflect.TypeOf(&ACMEIssuerDNS01ProviderCloudflare{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEIssuerDNS01ProviderRoute53).DeepCopyInto(out.(*ACMEIssuerDNS01ProviderRoute53))
			return nil
		}, InType: reflect.TypeOf(&ACMEIssuerDNS01ProviderRoute53{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ACMEIssuerStatus).DeepCopyInto(out.(*ACMEIssuerStatus))
			return nil
		}, InType: reflect.TypeOf(&ACMEIssuerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Certificate).DeepCopyInto(out.(*Certificate))
			return nil
		}, InType: reflect.TypeOf(&Certificate{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateACMEStatus).DeepCopyInto(out.(*CertificateACMEStatus))
			return nil
		}, InType: reflect.TypeOf(&CertificateACMEStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateList).DeepCopyInto(out.(*CertificateList))
			return nil
		}, InType: reflect.TypeOf(&CertificateList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateSpec).DeepCopyInto(out.(*CertificateSpec))
			return nil
		}, InType: reflect.TypeOf(&CertificateSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CertificateStatus).DeepCopyInto(out.(*CertificateStatus))
			return nil
		}, InType: reflect.TypeOf(&CertificateStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Issuer).DeepCopyInto(out.(*Issuer))
			return nil
		}, InType: reflect.TypeOf(&Issuer{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IssuerCondition).DeepCopyInto(out.(*IssuerCondition))
			return nil
		}, InType: reflect.TypeOf(&IssuerCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IssuerList).DeepCopyInto(out.(*IssuerList))
			return nil
		}, InType: reflect.TypeOf(&IssuerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IssuerSpec).DeepCopyInto(out.(*IssuerSpec))
			return nil
		}, InType: reflect.TypeOf(&IssuerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*IssuerStatus).DeepCopyInto(out.(*IssuerStatus))
			return nil
		}, InType: reflect.TypeOf(&IssuerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LocalObjectReference).DeepCopyInto(out.(*LocalObjectReference))
			return nil
		}, InType: reflect.TypeOf(&LocalObjectReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SecretKeySelector).DeepCopyInto(out.(*SecretKeySelector))
			return nil
		}, InType: reflect.TypeOf(&SecretKeySelector{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMECertificateConfig) DeepCopyInto(out *ACMECertificateConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ACMECertificateDomainConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMECertificateConfig.
func (in *ACMECertificateConfig) DeepCopy() *ACMECertificateConfig {
	if in == nil {
		return nil
	}
	out := new(ACMECertificateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMECertificateDNS01Config) DeepCopyInto(out *ACMECertificateDNS01Config) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMECertificateDNS01Config.
func (in *ACMECertificateDNS01Config) DeepCopy() *ACMECertificateDNS01Config {
	if in == nil {
		return nil
	}
	out := new(ACMECertificateDNS01Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMECertificateDomainConfig) DeepCopyInto(out *ACMECertificateDomainConfig) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HTTP01 != nil {
		in, out := &in.HTTP01, &out.HTTP01
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMECertificateHTTP01Config)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.DNS01 != nil {
		in, out := &in.DNS01, &out.DNS01
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMECertificateDNS01Config)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMECertificateDomainConfig.
func (in *ACMECertificateDomainConfig) DeepCopy() *ACMECertificateDomainConfig {
	if in == nil {
		return nil
	}
	out := new(ACMECertificateDomainConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMECertificateHTTP01Config) DeepCopyInto(out *ACMECertificateHTTP01Config) {
	*out = *in
	if in.IngressClass != nil {
		in, out := &in.IngressClass, &out.IngressClass
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMECertificateHTTP01Config.
func (in *ACMECertificateHTTP01Config) DeepCopy() *ACMECertificateHTTP01Config {
	if in == nil {
		return nil
	}
	out := new(ACMECertificateHTTP01Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEDomainAuthorization) DeepCopyInto(out *ACMEDomainAuthorization) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEDomainAuthorization.
func (in *ACMEDomainAuthorization) DeepCopy() *ACMEDomainAuthorization {
	if in == nil {
		return nil
	}
	out := new(ACMEDomainAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuer) DeepCopyInto(out *ACMEIssuer) {
	*out = *in
	if in.DNS01 != nil {
		in, out := &in.DNS01, &out.DNS01
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01Config)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuer.
func (in *ACMEIssuer) DeepCopy() *ACMEIssuer {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01Config) DeepCopyInto(out *ACMEIssuerDNS01Config) {
	*out = *in
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]ACMEIssuerDNS01Provider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01Config.
func (in *ACMEIssuerDNS01Config) DeepCopy() *ACMEIssuerDNS01Config {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01Provider) DeepCopyInto(out *ACMEIssuerDNS01Provider) {
	*out = *in
	if in.CloudDNS != nil {
		in, out := &in.CloudDNS, &out.CloudDNS
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderCloudDNS)
			**out = **in
		}
	}
	if in.Cloudflare != nil {
		in, out := &in.Cloudflare, &out.Cloudflare
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderCloudflare)
			**out = **in
		}
	}
	if in.Route53 != nil {
		in, out := &in.Route53, &out.Route53
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerDNS01ProviderRoute53)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01Provider.
func (in *ACMEIssuerDNS01Provider) DeepCopy() *ACMEIssuerDNS01Provider {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderCloudDNS) DeepCopyInto(out *ACMEIssuerDNS01ProviderCloudDNS) {
	*out = *in
	out.ServiceAccount = in.ServiceAccount
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderCloudDNS.
func (in *ACMEIssuerDNS01ProviderCloudDNS) DeepCopy() *ACMEIssuerDNS01ProviderCloudDNS {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderCloudDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderCloudflare) DeepCopyInto(out *ACMEIssuerDNS01ProviderCloudflare) {
	*out = *in
	out.APIKey = in.APIKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderCloudflare.
func (in *ACMEIssuerDNS01ProviderCloudflare) DeepCopy() *ACMEIssuerDNS01ProviderCloudflare {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderCloudflare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerDNS01ProviderRoute53) DeepCopyInto(out *ACMEIssuerDNS01ProviderRoute53) {
	*out = *in
	out.SecretAccessKey = in.SecretAccessKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerDNS01ProviderRoute53.
func (in *ACMEIssuerDNS01ProviderRoute53) DeepCopy() *ACMEIssuerDNS01ProviderRoute53 {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerDNS01ProviderRoute53)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEIssuerStatus) DeepCopyInto(out *ACMEIssuerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEIssuerStatus.
func (in *ACMEIssuerStatus) DeepCopy() *ACMEIssuerStatus {
	if in == nil {
		return nil
	}
	out := new(ACMEIssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateACMEStatus) DeepCopyInto(out *CertificateACMEStatus) {
	*out = *in
	if in.Authorizations != nil {
		in, out := &in.Authorizations, &out.Authorizations
		*out = make([]ACMEDomainAuthorization, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateACMEStatus.
func (in *CertificateACMEStatus) DeepCopy() *CertificateACMEStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateACMEStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMECertificateConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(CertificateACMEStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Issuer) DeepCopyInto(out *Issuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Issuer.
func (in *Issuer) DeepCopy() *Issuer {
	if in == nil {
		return nil
	}
	out := new(Issuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Issuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerCondition) DeepCopyInto(out *IssuerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerCondition.
func (in *IssuerCondition) DeepCopy() *IssuerCondition {
	if in == nil {
		return nil
	}
	out := new(IssuerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerList) DeepCopyInto(out *IssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Issuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerList.
func (in *IssuerList) DeepCopy() *IssuerList {
	if in == nil {
		return nil
	}
	out := new(IssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerSpec) DeepCopyInto(out *IssuerSpec) {
	*out = *in
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuer)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerSpec.
func (in *IssuerSpec) DeepCopy() *IssuerSpec {
	if in == nil {
		return nil
	}
	out := new(IssuerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerStatus) DeepCopyInto(out *IssuerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]IssuerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		if *in == nil {
			*out = nil
		} else {
			*out = new(ACMEIssuerStatus)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerStatus.
func (in *IssuerStatus) DeepCopy() *IssuerStatus {
	if in == nil {
		return nil
	}
	out := new(IssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}
