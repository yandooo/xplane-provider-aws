// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRole) DeepCopyInto(out *IAMRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRole.
func (in *IAMRole) DeepCopy() *IAMRole {
	if in == nil {
		return nil
	}
	out := new(IAMRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleARNReferencer) DeepCopyInto(out *IAMRoleARNReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleARNReferencer.
func (in *IAMRoleARNReferencer) DeepCopy() *IAMRoleARNReferencer {
	if in == nil {
		return nil
	}
	out := new(IAMRoleARNReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleExternalStatus) DeepCopyInto(out *IAMRoleExternalStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleExternalStatus.
func (in *IAMRoleExternalStatus) DeepCopy() *IAMRoleExternalStatus {
	if in == nil {
		return nil
	}
	out := new(IAMRoleExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleList) DeepCopyInto(out *IAMRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleList.
func (in *IAMRoleList) DeepCopy() *IAMRoleList {
	if in == nil {
		return nil
	}
	out := new(IAMRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleNameReferencer) DeepCopyInto(out *IAMRoleNameReferencer) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleNameReferencer.
func (in *IAMRoleNameReferencer) DeepCopy() *IAMRoleNameReferencer {
	if in == nil {
		return nil
	}
	out := new(IAMRoleNameReferencer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleNameReferencerForIAMRolePolicyAttachment) DeepCopyInto(out *IAMRoleNameReferencerForIAMRolePolicyAttachment) {
	*out = *in
	out.IAMRoleNameReferencer = in.IAMRoleNameReferencer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleNameReferencerForIAMRolePolicyAttachment.
func (in *IAMRoleNameReferencerForIAMRolePolicyAttachment) DeepCopy() *IAMRoleNameReferencerForIAMRolePolicyAttachment {
	if in == nil {
		return nil
	}
	out := new(IAMRoleNameReferencerForIAMRolePolicyAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleParameters) DeepCopyInto(out *IAMRoleParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleParameters.
func (in *IAMRoleParameters) DeepCopy() *IAMRoleParameters {
	if in == nil {
		return nil
	}
	out := new(IAMRoleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRolePolicyAttachment) DeepCopyInto(out *IAMRolePolicyAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRolePolicyAttachment.
func (in *IAMRolePolicyAttachment) DeepCopy() *IAMRolePolicyAttachment {
	if in == nil {
		return nil
	}
	out := new(IAMRolePolicyAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMRolePolicyAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRolePolicyAttachmentExternalStatus) DeepCopyInto(out *IAMRolePolicyAttachmentExternalStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRolePolicyAttachmentExternalStatus.
func (in *IAMRolePolicyAttachmentExternalStatus) DeepCopy() *IAMRolePolicyAttachmentExternalStatus {
	if in == nil {
		return nil
	}
	out := new(IAMRolePolicyAttachmentExternalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRolePolicyAttachmentList) DeepCopyInto(out *IAMRolePolicyAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMRolePolicyAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRolePolicyAttachmentList.
func (in *IAMRolePolicyAttachmentList) DeepCopy() *IAMRolePolicyAttachmentList {
	if in == nil {
		return nil
	}
	out := new(IAMRolePolicyAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMRolePolicyAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRolePolicyAttachmentParameters) DeepCopyInto(out *IAMRolePolicyAttachmentParameters) {
	*out = *in
	if in.RoleNameRef != nil {
		in, out := &in.RoleNameRef, &out.RoleNameRef
		*out = new(IAMRoleNameReferencerForIAMRolePolicyAttachment)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRolePolicyAttachmentParameters.
func (in *IAMRolePolicyAttachmentParameters) DeepCopy() *IAMRolePolicyAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(IAMRolePolicyAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRolePolicyAttachmentSpec) DeepCopyInto(out *IAMRolePolicyAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.IAMRolePolicyAttachmentParameters.DeepCopyInto(&out.IAMRolePolicyAttachmentParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRolePolicyAttachmentSpec.
func (in *IAMRolePolicyAttachmentSpec) DeepCopy() *IAMRolePolicyAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(IAMRolePolicyAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRolePolicyAttachmentStatus) DeepCopyInto(out *IAMRolePolicyAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.IAMRolePolicyAttachmentExternalStatus = in.IAMRolePolicyAttachmentExternalStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRolePolicyAttachmentStatus.
func (in *IAMRolePolicyAttachmentStatus) DeepCopy() *IAMRolePolicyAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(IAMRolePolicyAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleSpec) DeepCopyInto(out *IAMRoleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.IAMRoleParameters = in.IAMRoleParameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleSpec.
func (in *IAMRoleSpec) DeepCopy() *IAMRoleSpec {
	if in == nil {
		return nil
	}
	out := new(IAMRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMRoleStatus) DeepCopyInto(out *IAMRoleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.IAMRoleExternalStatus = in.IAMRoleExternalStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMRoleStatus.
func (in *IAMRoleStatus) DeepCopy() *IAMRoleStatus {
	if in == nil {
		return nil
	}
	out := new(IAMRoleStatus)
	in.DeepCopyInto(out)
	return out
}
