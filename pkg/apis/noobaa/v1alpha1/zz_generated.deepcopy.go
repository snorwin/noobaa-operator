// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountsStatus) DeepCopyInto(out *AccountsStatus) {
	*out = *in
	out.Admin = in.Admin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountsStatus.
func (in *AccountsStatus) DeepCopy() *AccountsStatus {
	if in == nil {
		return nil
	}
	out := new(AccountsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaa) DeepCopyInto(out *NooBaa) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaa.
func (in *NooBaa) DeepCopy() *NooBaa {
	if in == nil {
		return nil
	}
	out := new(NooBaa)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NooBaa) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaList) DeepCopyInto(out *NooBaaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NooBaa, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaList.
func (in *NooBaaList) DeepCopy() *NooBaaList {
	if in == nil {
		return nil
	}
	out := new(NooBaaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NooBaaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaSpec) DeepCopyInto(out *NooBaaSpec) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaSpec.
func (in *NooBaaSpec) DeepCopy() *NooBaaSpec {
	if in == nil {
		return nil
	}
	out := new(NooBaaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaStatus) DeepCopyInto(out *NooBaaStatus) {
	*out = *in
	out.Accounts = in.Accounts
	in.Services.DeepCopyInto(&out.Services)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaStatus.
func (in *NooBaaStatus) DeepCopy() *NooBaaStatus {
	if in == nil {
		return nil
	}
	out := new(NooBaaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
	if in.NodePorts != nil {
		in, out := &in.NodePorts, &out.NodePorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodPorts != nil {
		in, out := &in.PodPorts, &out.PodPorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InternalIP != nil {
		in, out := &in.InternalIP, &out.InternalIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InternalDNS != nil {
		in, out := &in.InternalDNS, &out.InternalDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalDNS != nil {
		in, out := &in.ExternalDNS, &out.ExternalDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesStatus) DeepCopyInto(out *ServicesStatus) {
	*out = *in
	in.ServiceMgmt.DeepCopyInto(&out.ServiceMgmt)
	in.ServiceS3.DeepCopyInto(&out.ServiceS3)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesStatus.
func (in *ServicesStatus) DeepCopy() *ServicesStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}