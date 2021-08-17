// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedManifestResourceMeta) DeepCopyInto(out *AppliedManifestResourceMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedManifestResourceMeta.
func (in *AppliedManifestResourceMeta) DeepCopy() *AppliedManifestResourceMeta {
	if in == nil {
		return nil
	}
	out := new(AppliedManifestResourceMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedManifestWork) DeepCopyInto(out *AppliedManifestWork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedManifestWork.
func (in *AppliedManifestWork) DeepCopy() *AppliedManifestWork {
	if in == nil {
		return nil
	}
	out := new(AppliedManifestWork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedManifestWork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedManifestWorkList) DeepCopyInto(out *AppliedManifestWorkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppliedManifestWork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedManifestWorkList.
func (in *AppliedManifestWorkList) DeepCopy() *AppliedManifestWorkList {
	if in == nil {
		return nil
	}
	out := new(AppliedManifestWorkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedManifestWorkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedManifestWorkSpec) DeepCopyInto(out *AppliedManifestWorkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedManifestWorkSpec.
func (in *AppliedManifestWorkSpec) DeepCopy() *AppliedManifestWorkSpec {
	if in == nil {
		return nil
	}
	out := new(AppliedManifestWorkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedManifestWorkStatus) DeepCopyInto(out *AppliedManifestWorkStatus) {
	*out = *in
	if in.AppliedResources != nil {
		in, out := &in.AppliedResources, &out.AppliedResources
		*out = make([]AppliedManifestResourceMeta, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedManifestWorkStatus.
func (in *AppliedManifestWorkStatus) DeepCopy() *AppliedManifestWorkStatus {
	if in == nil {
		return nil
	}
	out := new(AppliedManifestWorkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeleteOption) DeepCopyInto(out *DeleteOption) {
	*out = *in
	if in.SelectivelyOrphan != nil {
		in, out := &in.SelectivelyOrphan, &out.SelectivelyOrphan
		*out = new(SelectivelyOrphan)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeleteOption.
func (in *DeleteOption) DeepCopy() *DeleteOption {
	if in == nil {
		return nil
	}
	out := new(DeleteOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Manifest) DeepCopyInto(out *Manifest) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Manifest.
func (in *Manifest) DeepCopy() *Manifest {
	if in == nil {
		return nil
	}
	out := new(Manifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestCondition) DeepCopyInto(out *ManifestCondition) {
	*out = *in
	out.ResourceMeta = in.ResourceMeta
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestCondition.
func (in *ManifestCondition) DeepCopy() *ManifestCondition {
	if in == nil {
		return nil
	}
	out := new(ManifestCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestResourceMeta) DeepCopyInto(out *ManifestResourceMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestResourceMeta.
func (in *ManifestResourceMeta) DeepCopy() *ManifestResourceMeta {
	if in == nil {
		return nil
	}
	out := new(ManifestResourceMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestResourceStatus) DeepCopyInto(out *ManifestResourceStatus) {
	*out = *in
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]ManifestCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestResourceStatus.
func (in *ManifestResourceStatus) DeepCopy() *ManifestResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ManifestResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestWork) DeepCopyInto(out *ManifestWork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestWork.
func (in *ManifestWork) DeepCopy() *ManifestWork {
	if in == nil {
		return nil
	}
	out := new(ManifestWork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManifestWork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestWorkList) DeepCopyInto(out *ManifestWorkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManifestWork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestWorkList.
func (in *ManifestWorkList) DeepCopy() *ManifestWorkList {
	if in == nil {
		return nil
	}
	out := new(ManifestWorkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManifestWorkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestWorkSpec) DeepCopyInto(out *ManifestWorkSpec) {
	*out = *in
	in.Workload.DeepCopyInto(&out.Workload)
	if in.DeleteOption != nil {
		in, out := &in.DeleteOption, &out.DeleteOption
		*out = new(DeleteOption)
		(*in).DeepCopyInto(*out)
	}
	out.Ownership = in.Ownership
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestWorkSpec.
func (in *ManifestWorkSpec) DeepCopy() *ManifestWorkSpec {
	if in == nil {
		return nil
	}
	out := new(ManifestWorkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestWorkStatus) DeepCopyInto(out *ManifestWorkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestWorkStatus.
func (in *ManifestWorkStatus) DeepCopy() *ManifestWorkStatus {
	if in == nil {
		return nil
	}
	out := new(ManifestWorkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManifestsTemplate) DeepCopyInto(out *ManifestsTemplate) {
	*out = *in
	if in.Manifests != nil {
		in, out := &in.Manifests, &out.Manifests
		*out = make([]Manifest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManifestsTemplate.
func (in *ManifestsTemplate) DeepCopy() *ManifestsTemplate {
	if in == nil {
		return nil
	}
	out := new(ManifestsTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrphaningRule) DeepCopyInto(out *OrphaningRule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrphaningRule.
func (in *OrphaningRule) DeepCopy() *OrphaningRule {
	if in == nil {
		return nil
	}
	out := new(OrphaningRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectivelyOrphan) DeepCopyInto(out *SelectivelyOrphan) {
	*out = *in
	if in.OrphaningRules != nil {
		in, out := &in.OrphaningRules, &out.OrphaningRules
		*out = make([]OrphaningRule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectivelyOrphan.
func (in *SelectivelyOrphan) DeepCopy() *SelectivelyOrphan {
	if in == nil {
		return nil
	}
	out := new(SelectivelyOrphan)
	in.DeepCopyInto(out)
	return out
}
