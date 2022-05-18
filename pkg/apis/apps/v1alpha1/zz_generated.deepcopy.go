//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Gaia Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppSlaAttr) DeepCopyInto(out *AppSlaAttr) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppSlaAttr.
func (in *AppSlaAttr) DeepCopy() *AppSlaAttr {
	if in == nil {
		return nil
	}
	out := new(AppSlaAttr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Attributes) DeepCopyInto(out *Attributes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Attributes.
func (in *Attributes) DeepCopy() *Attributes {
	if in == nil {
		return nil
	}
	out := new(Attributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	in.Serverless.DeepCopyInto(&out.Serverless)
	in.Module.DeepCopyInto(&out.Module)
	in.Workload.DeepCopyInto(&out.Workload)
	in.SchedulePolicy.DeepCopyInto(&out.SchedulePolicy)
	if in.ClusterTolerations != nil {
		in, out := &in.ClusterTolerations, &out.ClusterTolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Description) DeepCopyInto(out *Description) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Description.
func (in *Description) DeepCopy() *Description {
	if in == nil {
		return nil
	}
	out := new(Description)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Description) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescriptionList) DeepCopyInto(out *DescriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Description, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescriptionList.
func (in *DescriptionList) DeepCopy() *DescriptionList {
	if in == nil {
		return nil
	}
	out := new(DescriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DescriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescriptionSpec) DeepCopyInto(out *DescriptionSpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]Component, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescriptionSpec.
func (in *DescriptionSpec) DeepCopy() *DescriptionSpec {
	if in == nil {
		return nil
	}
	out := new(DescriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescriptionStatus) DeepCopyInto(out *DescriptionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescriptionStatus.
func (in *DescriptionStatus) DeepCopy() *DescriptionStatus {
	if in == nil {
		return nil
	}
	out := new(DescriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Direction) DeepCopyInto(out *Direction) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]Attributes, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Direction.
func (in *Direction) DeepCopy() *Direction {
	if in == nil {
		return nil
	}
	out := new(Direction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterSCNID) DeepCopyInto(out *InterSCNID) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	out.Sla = in.Sla
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterSCNID.
func (in *InterSCNID) DeepCopy() *InterSCNID {
	if in == nil {
		return nil
	}
	out := new(InterSCNID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkCommunication) DeepCopyInto(out *NetworkCommunication) {
	*out = *in
	if in.SelfID != nil {
		in, out := &in.SelfID, &out.SelfID
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InterSCNID != nil {
		in, out := &in.InterSCNID, &out.InterSCNID
		*out = make([]InterSCNID, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkCommunication.
func (in *NetworkCommunication) DeepCopy() *NetworkCommunication {
	if in == nil {
		return nil
	}
	out := new(NetworkCommunication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRequirement) DeepCopyInto(out *NetworkRequirement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRequirement.
func (in *NetworkRequirement) DeepCopy() *NetworkRequirement {
	if in == nil {
		return nil
	}
	out := new(NetworkRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkRequirement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRequirementList) DeepCopyInto(out *NetworkRequirementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRequirementList.
func (in *NetworkRequirementList) DeepCopy() *NetworkRequirementList {
	if in == nil {
		return nil
	}
	out := new(NetworkRequirementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkRequirementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRequirementSpec) DeepCopyInto(out *NetworkRequirementSpec) {
	*out = *in
	if in.NetworkCommunication != nil {
		in, out := &in.NetworkCommunication, &out.NetworkCommunication
		*out = make([]NetworkCommunication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRequirementSpec.
func (in *NetworkRequirementSpec) DeepCopy() *NetworkRequirementSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkRequirementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkRequirementStatus) DeepCopyInto(out *NetworkRequirementStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkRequirementStatus.
func (in *NetworkRequirementStatus) DeepCopy() *NetworkRequirementStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkRequirementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBinding) DeepCopyInto(out *ResourceBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBinding.
func (in *ResourceBinding) DeepCopy() *ResourceBinding {
	if in == nil {
		return nil
	}
	out := new(ResourceBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBindingApps) DeepCopyInto(out *ResourceBindingApps) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Children != nil {
		in, out := &in.Children, &out.Children
		*out = make([]*ResourceBindingApps, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceBindingApps)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBindingApps.
func (in *ResourceBindingApps) DeepCopy() *ResourceBindingApps {
	if in == nil {
		return nil
	}
	out := new(ResourceBindingApps)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBindingList) DeepCopyInto(out *ResourceBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBindingList.
func (in *ResourceBindingList) DeepCopy() *ResourceBindingList {
	if in == nil {
		return nil
	}
	out := new(ResourceBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBindingSpec) DeepCopyInto(out *ResourceBindingSpec) {
	*out = *in
	if in.RbApps != nil {
		in, out := &in.RbApps, &out.RbApps
		*out = make([]*ResourceBindingApps, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceBindingApps)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NetworkPath != nil {
		in, out := &in.NetworkPath, &out.NetworkPath
		*out = make([][]byte, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBindingSpec.
func (in *ResourceBindingSpec) DeepCopy() *ResourceBindingSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBindingStatus) DeepCopyInto(out *ResourceBindingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBindingStatus.
func (in *ResourceBindingStatus) DeepCopy() *ResourceBindingStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionSpec) DeepCopyInto(out *RevisionSpec) {
	*out = *in
	if in.ContainerConcurrency != nil {
		in, out := &in.ContainerConcurrency, &out.ContainerConcurrency
		*out = new(int64)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionSpec.
func (in *RevisionSpec) DeepCopy() *RevisionSpec {
	if in == nil {
		return nil
	}
	out := new(RevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulePolicy) DeepCopyInto(out *SchedulePolicy) {
	*out = *in
	if in.SpecificResource != nil {
		in, out := &in.SpecificResource, &out.SpecificResource
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Netenvironment != nil {
		in, out := &in.Netenvironment, &out.Netenvironment
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Geolocation != nil {
		in, out := &in.Geolocation, &out.Geolocation
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulePolicy.
func (in *SchedulePolicy) DeepCopy() *SchedulePolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Serverless) DeepCopyInto(out *Serverless) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Serverless.
func (in *Serverless) DeepCopy() *Serverless {
	if in == nil {
		return nil
	}
	out := new(Serverless)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Serverless) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessList) DeepCopyInto(out *ServerlessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Serverless, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessList.
func (in *ServerlessList) DeepCopy() *ServerlessList {
	if in == nil {
		return nil
	}
	out := new(ServerlessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerlessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessSpec) DeepCopyInto(out *ServerlessSpec) {
	*out = *in
	in.Revision.DeepCopyInto(&out.Revision)
	if in.Traffic != nil {
		in, out := &in.Traffic, &out.Traffic
		*out = make([]servingv1.TrafficTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessSpec.
func (in *ServerlessSpec) DeepCopy() *ServerlessSpec {
	if in == nil {
		return nil
	}
	out := new(ServerlessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessStatus) DeepCopyInto(out *ServerlessStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessStatus.
func (in *ServerlessStatus) DeepCopy() *ServerlessStatus {
	if in == nil {
		return nil
	}
	out := new(ServerlessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraitDeployment) DeepCopyInto(out *TraitDeployment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraitDeployment.
func (in *TraitDeployment) DeepCopy() *TraitDeployment {
	if in == nil {
		return nil
	}
	out := new(TraitDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraitServerless) DeepCopyInto(out *TraitServerless) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraitServerless.
func (in *TraitServerless) DeepCopy() *TraitServerless {
	if in == nil {
		return nil
	}
	out := new(TraitServerless)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	if in.TraitDeployment != nil {
		in, out := &in.TraitDeployment, &out.TraitDeployment
		*out = new(TraitDeployment)
		**out = **in
	}
	if in.TraitServerless != nil {
		in, out := &in.TraitServerless, &out.TraitServerless
		*out = new(TraitServerless)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}
