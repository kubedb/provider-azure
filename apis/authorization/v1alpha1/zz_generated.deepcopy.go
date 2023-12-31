//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignment) DeepCopyInto(out *RoleAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignment.
func (in *RoleAssignment) DeepCopy() *RoleAssignment {
	if in == nil {
		return nil
	}
	out := new(RoleAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentInitParameters) DeepCopyInto(out *RoleAssignmentInitParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceID != nil {
		in, out := &in.DelegatedManagedIdentityResourceID, &out.DelegatedManagedIdentityResourceID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionID != nil {
		in, out := &in.RoleDefinitionID, &out.RoleDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionName != nil {
		in, out := &in.RoleDefinitionName, &out.RoleDefinitionName
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.SkipServicePrincipalAadCheck != nil {
		in, out := &in.SkipServicePrincipalAadCheck, &out.SkipServicePrincipalAadCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentInitParameters.
func (in *RoleAssignmentInitParameters) DeepCopy() *RoleAssignmentInitParameters {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentList) DeepCopyInto(out *RoleAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentList.
func (in *RoleAssignmentList) DeepCopy() *RoleAssignmentList {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentObservation) DeepCopyInto(out *RoleAssignmentObservation) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceID != nil {
		in, out := &in.DelegatedManagedIdentityResourceID, &out.DelegatedManagedIdentityResourceID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionID != nil {
		in, out := &in.RoleDefinitionID, &out.RoleDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionName != nil {
		in, out := &in.RoleDefinitionName, &out.RoleDefinitionName
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.SkipServicePrincipalAadCheck != nil {
		in, out := &in.SkipServicePrincipalAadCheck, &out.SkipServicePrincipalAadCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentObservation.
func (in *RoleAssignmentObservation) DeepCopy() *RoleAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentParameters) DeepCopyInto(out *RoleAssignmentParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceID != nil {
		in, out := &in.DelegatedManagedIdentityResourceID, &out.DelegatedManagedIdentityResourceID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionID != nil {
		in, out := &in.RoleDefinitionID, &out.RoleDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionName != nil {
		in, out := &in.RoleDefinitionName, &out.RoleDefinitionName
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.SkipServicePrincipalAadCheck != nil {
		in, out := &in.SkipServicePrincipalAadCheck, &out.SkipServicePrincipalAadCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentParameters.
func (in *RoleAssignmentParameters) DeepCopy() *RoleAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentSpec) DeepCopyInto(out *RoleAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentSpec.
func (in *RoleAssignmentSpec) DeepCopy() *RoleAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentStatus) DeepCopyInto(out *RoleAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentStatus.
func (in *RoleAssignmentStatus) DeepCopy() *RoleAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}
