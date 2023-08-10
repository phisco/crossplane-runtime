//
//Copyright 2022 The Crossplane Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/v1beta1/fn/run_function.proto

package v1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Severity of Function results.
type Severity int32

const (
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	// Fatal results are fatal; subsequent Composition Functions may run, but
	// the Composition Function pipeline run will be considered a failure and
	// the first fatal result will be returned as an error.
	Severity_SEVERITY_FATAL Severity = 1
	// Warning results are non-fatal; the entire Composition will run to
	// completion but warning events and debug logs associated with the
	// composite resource will be emitted.
	Severity_SEVERITY_WARNING Severity = 2
	// Normal results are emitted as normal events and debug logs associated
	// with the composite resource.
	Severity_SEVERITY_NORMAL Severity = 3
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_FATAL",
		2: "SEVERITY_WARNING",
		3: "SEVERITY_NORMAL",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_FATAL":       1,
		"SEVERITY_WARNING":     2,
		"SEVERITY_NORMAL":      3,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1beta1_fn_run_function_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_proto_v1beta1_fn_run_function_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{0}
}

// A RunFunctionRequest requests that the Composition Function be run.
type RunFunctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata pertaining to this request.
	Meta *RequestMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// The observed state prior to invocation of a Function pipeline. State passed
	// to each Function is fresh as of the time the pipeline was invoked, not as
	// of the time each Function was invoked.
	Observed *State `protobuf:"bytes,2,opt,name=observed,proto3" json:"observed,omitempty"`
	// Desired state according to a Function pipeline. The state passed to a
	// particular Function may have been accumulated by processing a Composition's
	// patch-and-transform resources array. It may also have been accumulated by
	// previous Functions in the pipeline.
	Desired *State `protobuf:"bytes,3,opt,name=desired,proto3" json:"desired,omitempty"`
	// Optional input specific to this Function invocation. A JSON representation
	// of the 'input' block of the relevant entry in a Composition's pipeline.
	Input *structpb.Struct `protobuf:"bytes,4,opt,name=input,proto3,oneof" json:"input,omitempty"`
}

func (x *RunFunctionRequest) Reset() {
	*x = RunFunctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunFunctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunFunctionRequest) ProtoMessage() {}

func (x *RunFunctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunFunctionRequest.ProtoReflect.Descriptor instead.
func (*RunFunctionRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{0}
}

func (x *RunFunctionRequest) GetMeta() *RequestMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RunFunctionRequest) GetObserved() *State {
	if x != nil {
		return x.Observed
	}
	return nil
}

func (x *RunFunctionRequest) GetDesired() *State {
	if x != nil {
		return x.Desired
	}
	return nil
}

func (x *RunFunctionRequest) GetInput() *structpb.Struct {
	if x != nil {
		return x.Input
	}
	return nil
}

// A RunFunctionResponse contains the result of a Composition Function run.
type RunFunctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata pertaining to this response.
	Meta *ResponseMeta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Desired state according to a Function pipeline. Functions may add desired
	// state, and may mutate or delete any part of the desired state they are
	// concerned with. A Function must pass through any part of the desired state
	// that it is not concerned with.
	Desired *State `protobuf:"bytes,2,opt,name=desired,proto3" json:"desired,omitempty"`
	// Results of the Function run. Results are used for observability purposes.
	Results []*Result `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *RunFunctionResponse) Reset() {
	*x = RunFunctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunFunctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunFunctionResponse) ProtoMessage() {}

func (x *RunFunctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunFunctionResponse.ProtoReflect.Descriptor instead.
func (*RunFunctionResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{1}
}

func (x *RunFunctionResponse) GetMeta() *ResponseMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RunFunctionResponse) GetDesired() *State {
	if x != nil {
		return x.Desired
	}
	return nil
}

func (x *RunFunctionResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

// RequestMeta contains metadata pertaining to a RunFunctionRequest.
type RequestMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An opaque string identifying the content of the request. Two identical
	// requests should have the same tag.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *RequestMeta) Reset() {
	*x = RequestMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMeta) ProtoMessage() {}

func (x *RequestMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMeta.ProtoReflect.Descriptor instead.
func (*RequestMeta) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{2}
}

func (x *RequestMeta) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

// ResponseMeta contains metadata pertaining to a RunFunctionResponse.
type ResponseMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An opaque string identifying the content of the request. Must match the
	// meta.tag of the corresponding RunFunctionRequest.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Time-to-live of this response. Deterministic Functions with no side-effects
	// (e.g. simple templating Functions) may specify a TTL. Crossplane may choose
	// to cache responses until the TTL expires.
	Ttl *durationpb.Duration `protobuf:"bytes,2,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
}

func (x *ResponseMeta) Reset() {
	*x = ResponseMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMeta) ProtoMessage() {}

func (x *ResponseMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMeta.ProtoReflect.Descriptor instead.
func (*ResponseMeta) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseMeta) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ResponseMeta) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

// State of the composite resource (XR) and any composed resources.
type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state of the composite resource (XR).
	Composite *Resource `protobuf:"bytes,1,opt,name=composite,proto3" json:"composite,omitempty"`
	// The state of any composed resources.
	Resources map[string]*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{4}
}

func (x *State) GetComposite() *Resource {
	if x != nil {
		return x.Composite
	}
	return nil
}

func (x *State) GetResources() map[string]*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

// A Resource represents the state of a resource.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The JSON representation of the resource.
	Resource *structpb.Struct `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The resource's connection details.
	ConnectionDetails map[string][]byte `protobuf:"bytes,2,rep,name=connection_details,json=connectionDetails,proto3" json:"connection_details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{5}
}

func (x *Resource) GetResource() *structpb.Struct {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Resource) GetConnectionDetails() map[string][]byte {
	if x != nil {
		return x.ConnectionDetails
	}
	return nil
}

// A Result of running a Function.
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Severity of this result.
	Severity Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=apiextensions.fn.proto.v1beta1.Severity" json:"severity,omitempty"`
	// Human-readable details about the result.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_fn_run_function_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_fn_run_function_proto_rawDescGZIP(), []int{6}
}

func (x *Result) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_SEVERITY_UNSPECIFIED
}

func (x *Result) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_v1beta1_fn_run_function_proto protoreflect.FileDescriptor

var file_proto_v1beta1_fn_run_function_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x66, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x08, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x08, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x3f,
	0x0a, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x32, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xda, 0x01,
	0x0a, 0x13, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x1f, 0x0a, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x5a, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x30, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x88, 0x01, 0x01, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c, 0x22, 0x8b, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x46, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x61,
	0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x66, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf5, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x44, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x68, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x63, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x57, 0x41,
	0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x56, 0x45, 0x52,
	0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x03, 0x32, 0x91, 0x01, 0x0a,
	0x15, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x72, 0x6f, 0x73, 0x73, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1beta1_fn_run_function_proto_rawDescOnce sync.Once
	file_proto_v1beta1_fn_run_function_proto_rawDescData = file_proto_v1beta1_fn_run_function_proto_rawDesc
)

func file_proto_v1beta1_fn_run_function_proto_rawDescGZIP() []byte {
	file_proto_v1beta1_fn_run_function_proto_rawDescOnce.Do(func() {
		file_proto_v1beta1_fn_run_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1beta1_fn_run_function_proto_rawDescData)
	})
	return file_proto_v1beta1_fn_run_function_proto_rawDescData
}

var file_proto_v1beta1_fn_run_function_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_v1beta1_fn_run_function_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_v1beta1_fn_run_function_proto_goTypes = []interface{}{
	(Severity)(0),               // 0: apiextensions.fn.proto.v1beta1.Severity
	(*RunFunctionRequest)(nil),  // 1: apiextensions.fn.proto.v1beta1.RunFunctionRequest
	(*RunFunctionResponse)(nil), // 2: apiextensions.fn.proto.v1beta1.RunFunctionResponse
	(*RequestMeta)(nil),         // 3: apiextensions.fn.proto.v1beta1.RequestMeta
	(*ResponseMeta)(nil),        // 4: apiextensions.fn.proto.v1beta1.ResponseMeta
	(*State)(nil),               // 5: apiextensions.fn.proto.v1beta1.State
	(*Resource)(nil),            // 6: apiextensions.fn.proto.v1beta1.Resource
	(*Result)(nil),              // 7: apiextensions.fn.proto.v1beta1.Result
	nil,                         // 8: apiextensions.fn.proto.v1beta1.State.ResourcesEntry
	nil,                         // 9: apiextensions.fn.proto.v1beta1.Resource.ConnectionDetailsEntry
	(*structpb.Struct)(nil),     // 10: google.protobuf.Struct
	(*durationpb.Duration)(nil), // 11: google.protobuf.Duration
}
var file_proto_v1beta1_fn_run_function_proto_depIdxs = []int32{
	3,  // 0: apiextensions.fn.proto.v1beta1.RunFunctionRequest.meta:type_name -> apiextensions.fn.proto.v1beta1.RequestMeta
	5,  // 1: apiextensions.fn.proto.v1beta1.RunFunctionRequest.observed:type_name -> apiextensions.fn.proto.v1beta1.State
	5,  // 2: apiextensions.fn.proto.v1beta1.RunFunctionRequest.desired:type_name -> apiextensions.fn.proto.v1beta1.State
	10, // 3: apiextensions.fn.proto.v1beta1.RunFunctionRequest.input:type_name -> google.protobuf.Struct
	4,  // 4: apiextensions.fn.proto.v1beta1.RunFunctionResponse.meta:type_name -> apiextensions.fn.proto.v1beta1.ResponseMeta
	5,  // 5: apiextensions.fn.proto.v1beta1.RunFunctionResponse.desired:type_name -> apiextensions.fn.proto.v1beta1.State
	7,  // 6: apiextensions.fn.proto.v1beta1.RunFunctionResponse.results:type_name -> apiextensions.fn.proto.v1beta1.Result
	11, // 7: apiextensions.fn.proto.v1beta1.ResponseMeta.ttl:type_name -> google.protobuf.Duration
	6,  // 8: apiextensions.fn.proto.v1beta1.State.composite:type_name -> apiextensions.fn.proto.v1beta1.Resource
	8,  // 9: apiextensions.fn.proto.v1beta1.State.resources:type_name -> apiextensions.fn.proto.v1beta1.State.ResourcesEntry
	10, // 10: apiextensions.fn.proto.v1beta1.Resource.resource:type_name -> google.protobuf.Struct
	9,  // 11: apiextensions.fn.proto.v1beta1.Resource.connection_details:type_name -> apiextensions.fn.proto.v1beta1.Resource.ConnectionDetailsEntry
	0,  // 12: apiextensions.fn.proto.v1beta1.Result.severity:type_name -> apiextensions.fn.proto.v1beta1.Severity
	6,  // 13: apiextensions.fn.proto.v1beta1.State.ResourcesEntry.value:type_name -> apiextensions.fn.proto.v1beta1.Resource
	1,  // 14: apiextensions.fn.proto.v1beta1.FunctionRunnerService.RunFunction:input_type -> apiextensions.fn.proto.v1beta1.RunFunctionRequest
	2,  // 15: apiextensions.fn.proto.v1beta1.FunctionRunnerService.RunFunction:output_type -> apiextensions.fn.proto.v1beta1.RunFunctionResponse
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_proto_v1beta1_fn_run_function_proto_init() }
func file_proto_v1beta1_fn_run_function_proto_init() {
	if File_proto_v1beta1_fn_run_function_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1beta1_fn_run_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunFunctionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1beta1_fn_run_function_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunFunctionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1beta1_fn_run_function_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1beta1_fn_run_function_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1beta1_fn_run_function_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1beta1_fn_run_function_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1beta1_fn_run_function_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_v1beta1_fn_run_function_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_v1beta1_fn_run_function_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1beta1_fn_run_function_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1beta1_fn_run_function_proto_goTypes,
		DependencyIndexes: file_proto_v1beta1_fn_run_function_proto_depIdxs,
		EnumInfos:         file_proto_v1beta1_fn_run_function_proto_enumTypes,
		MessageInfos:      file_proto_v1beta1_fn_run_function_proto_msgTypes,
	}.Build()
	File_proto_v1beta1_fn_run_function_proto = out.File
	file_proto_v1beta1_fn_run_function_proto_rawDesc = nil
	file_proto_v1beta1_fn_run_function_proto_goTypes = nil
	file_proto_v1beta1_fn_run_function_proto_depIdxs = nil
}
