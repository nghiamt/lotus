// +build !ignore_autogenerated

/*

Generated by using code-generator

*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lotus) DeepCopyInto(out *Lotus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lotus.
func (in *Lotus) DeepCopy() *Lotus {
	if in == nil {
		return nil
	}
	out := new(Lotus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Lotus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LotusCheck) DeepCopyInto(out *LotusCheck) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LotusCheck.
func (in *LotusCheck) DeepCopy() *LotusCheck {
	if in == nil {
		return nil
	}
	out := new(LotusCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LotusList) DeepCopyInto(out *LotusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Lotus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LotusList.
func (in *LotusList) DeepCopy() *LotusList {
	if in == nil {
		return nil
	}
	out := new(LotusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LotusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LotusSpec) DeepCopyInto(out *LotusSpec) {
	*out = *in
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.CheckIntervalSeconds != nil {
		in, out := &in.CheckIntervalSeconds, &out.CheckIntervalSeconds
		*out = new(int32)
		**out = **in
	}
	if in.CheckInitialDelaySeconds != nil {
		in, out := &in.CheckInitialDelaySeconds, &out.CheckInitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.Preparer != nil {
		in, out := &in.Preparer, &out.Preparer
		*out = new(LotusSpecPreparer)
		(*in).DeepCopyInto(*out)
	}
	if in.Worker != nil {
		in, out := &in.Worker, &out.Worker
		*out = new(LotusSpecWorker)
		(*in).DeepCopyInto(*out)
	}
	if in.Cleaner != nil {
		in, out := &in.Cleaner, &out.Cleaner
		*out = new(LotusSpecCleaner)
		(*in).DeepCopyInto(*out)
	}
	if in.Checks != nil {
		in, out := &in.Checks, &out.Checks
		*out = make([]LotusCheck, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LotusSpec.
func (in *LotusSpec) DeepCopy() *LotusSpec {
	if in == nil {
		return nil
	}
	out := new(LotusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LotusSpecCleaner) DeepCopyInto(out *LotusSpecCleaner) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LotusSpecCleaner.
func (in *LotusSpecCleaner) DeepCopy() *LotusSpecCleaner {
	if in == nil {
		return nil
	}
	out := new(LotusSpecCleaner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LotusSpecPreparer) DeepCopyInto(out *LotusSpecPreparer) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LotusSpecPreparer.
func (in *LotusSpecPreparer) DeepCopy() *LotusSpecPreparer {
	if in == nil {
		return nil
	}
	out := new(LotusSpecPreparer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LotusSpecWorker) DeepCopyInto(out *LotusSpecWorker) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MetricsPort != nil {
		in, out := &in.MetricsPort, &out.MetricsPort
		*out = new(int32)
		**out = **in
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LotusSpecWorker.
func (in *LotusSpecWorker) DeepCopy() *LotusSpecWorker {
	if in == nil {
		return nil
	}
	out := new(LotusSpecWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LotusStatus) DeepCopyInto(out *LotusStatus) {
	*out = *in
	if in.PreparerStartTime != nil {
		in, out := &in.PreparerStartTime, &out.PreparerStartTime
		*out = (*in).DeepCopy()
	}
	if in.PreparerCompletionTime != nil {
		in, out := &in.PreparerCompletionTime, &out.PreparerCompletionTime
		*out = (*in).DeepCopy()
	}
	if in.WorkerStartTime != nil {
		in, out := &in.WorkerStartTime, &out.WorkerStartTime
		*out = (*in).DeepCopy()
	}
	if in.WorkerCompletionTime != nil {
		in, out := &in.WorkerCompletionTime, &out.WorkerCompletionTime
		*out = (*in).DeepCopy()
	}
	if in.CleanerStartTime != nil {
		in, out := &in.CleanerStartTime, &out.CleanerStartTime
		*out = (*in).DeepCopy()
	}
	if in.CleanerCompletionTime != nil {
		in, out := &in.CleanerCompletionTime, &out.CleanerCompletionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LotusStatus.
func (in *LotusStatus) DeepCopy() *LotusStatus {
	if in == nil {
		return nil
	}
	out := new(LotusStatus)
	in.DeepCopyInto(out)
	return out
}