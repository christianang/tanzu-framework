//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSXT) DeepCopyInto(out *NSXT) {
	*out = *in
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = new(NSXTRoute)
		**out = **in
	}
	if in.NSXTCredentialsRef != nil {
		in, out := &in.NSXTCredentialsRef, &out.NSXTCredentialsRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSXT.
func (in *NSXT) DeepCopy() *NSXT {
	if in == nil {
		return nil
	}
	out := new(NSXT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NSXTRoute) DeepCopyInto(out *NSXTRoute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NSXTRoute.
func (in *NSXTRoute) DeepCopy() *NSXTRoute {
	if in == nil {
		return nil
	}
	out := new(NSXTRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonParavirtualConfig) DeepCopyInto(out *NonParavirtualConfig) {
	*out = *in
	if in.VSphereCredentialRef != nil {
		in, out := &in.VSphereCredentialRef, &out.VSphereCredentialRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.NSXT != nil {
		in, out := &in.NSXT, &out.NSXT
		*out = new(NSXT)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonParavirtualConfig.
func (in *NonParavirtualConfig) DeepCopy() *NonParavirtualConfig {
	if in == nil {
		return nil
	}
	out := new(NonParavirtualConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParavirtualConfig) DeepCopyInto(out *ParavirtualConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParavirtualConfig.
func (in *ParavirtualConfig) DeepCopy() *ParavirtualConfig {
	if in == nil {
		return nil
	}
	out := new(ParavirtualConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPI) DeepCopyInto(out *VSphereCPI) {
	*out = *in
	if in.NonParavirtualConfig != nil {
		in, out := &in.NonParavirtualConfig, &out.NonParavirtualConfig
		*out = new(NonParavirtualConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ParavirtualConfig != nil {
		in, out := &in.ParavirtualConfig, &out.ParavirtualConfig
		*out = new(ParavirtualConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPI.
func (in *VSphereCPI) DeepCopy() *VSphereCPI {
	if in == nil {
		return nil
	}
	out := new(VSphereCPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfig) DeepCopyInto(out *VSphereCPIConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfig.
func (in *VSphereCPIConfig) DeepCopy() *VSphereCPIConfig {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereCPIConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfigList) DeepCopyInto(out *VSphereCPIConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSphereCPIConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfigList.
func (in *VSphereCPIConfigList) DeepCopy() *VSphereCPIConfigList {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereCPIConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfigSpec) DeepCopyInto(out *VSphereCPIConfigSpec) {
	*out = *in
	in.VSphereCPI.DeepCopyInto(&out.VSphereCPI)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfigSpec.
func (in *VSphereCPIConfigSpec) DeepCopy() *VSphereCPIConfigSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereCPIConfigStatus) DeepCopyInto(out *VSphereCPIConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereCPIConfigStatus.
func (in *VSphereCPIConfigStatus) DeepCopy() *VSphereCPIConfigStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereCPIConfigStatus)
	in.DeepCopyInto(out)
	return out
}
