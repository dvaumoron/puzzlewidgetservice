// Copyright 2023 puzzlewidgetservice authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: widget.proto

package puzzlewidgetservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MethodKind int32

const (
	MethodKind_GET     MethodKind = 0
	MethodKind_HEAD    MethodKind = 1
	MethodKind_POST    MethodKind = 2
	MethodKind_PUT     MethodKind = 3
	MethodKind_PATCH   MethodKind = 4
	MethodKind_DELETE  MethodKind = 5
	MethodKind_CONNECT MethodKind = 6
	MethodKind_OPTIONS MethodKind = 7
	MethodKind_TRACE   MethodKind = 8
	MethodKind_RAW     MethodKind = 9 // added special category
)

// Enum value maps for MethodKind.
var (
	MethodKind_name = map[int32]string{
		0: "GET",
		1: "HEAD",
		2: "POST",
		3: "PUT",
		4: "PATCH",
		5: "DELETE",
		6: "CONNECT",
		7: "OPTIONS",
		8: "TRACE",
		9: "RAW",
	}
	MethodKind_value = map[string]int32{
		"GET":     0,
		"HEAD":    1,
		"POST":    2,
		"PUT":     3,
		"PATCH":   4,
		"DELETE":  5,
		"CONNECT": 6,
		"OPTIONS": 7,
		"TRACE":   8,
		"RAW":     9,
	}
)

func (x MethodKind) Enum() *MethodKind {
	p := new(MethodKind)
	*p = x
	return p
}

func (x MethodKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodKind) Descriptor() protoreflect.EnumDescriptor {
	return file_widget_proto_enumTypes[0].Descriptor()
}

func (MethodKind) Type() protoreflect.EnumType {
	return &file_widget_proto_enumTypes[0]
}

func (x MethodKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodKind.Descriptor instead.
func (MethodKind) EnumDescriptor() ([]byte, []int) {
	return file_widget_proto_rawDescGZIP(), []int{0}
}

type WidgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WidgetRequest) Reset() {
	*x = WidgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetRequest) ProtoMessage() {}

func (x *WidgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetRequest.ProtoReflect.Descriptor instead.
func (*WidgetRequest) Descriptor() ([]byte, []int) {
	return file_widget_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind       MethodKind `protobuf:"varint,1,opt,name=kind,proto3,enum=puzzlewidgetservice.MethodKind" json:"kind,omitempty"`
	Name       string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path       string     `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	QueryNames []string   `protobuf:"bytes,4,rep,name=queryNames,proto3" json:"queryNames,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_widget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_widget_proto_rawDescGZIP(), []int{1}
}

func (x *Action) GetKind() MethodKind {
	if x != nil {
		return x.Kind
	}
	return MethodKind_GET
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Action) GetQueryNames() []string {
	if x != nil {
		return x.QueryNames
	}
	return nil
}

type WidgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Actions []*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *WidgetResponse) Reset() {
	*x = WidgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetResponse) ProtoMessage() {}

func (x *WidgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_widget_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetResponse.ProtoReflect.Descriptor instead.
func (*WidgetResponse) Descriptor() ([]byte, []int) {
	return file_widget_proto_rawDescGZIP(), []int{2}
}

func (x *WidgetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WidgetResponse) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type ProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WidgetName string            `protobuf:"bytes,1,opt,name=widgetName,proto3" json:"widgetName,omitempty"`
	ActionName string            `protobuf:"bytes,2,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Files      map[string][]byte `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProcessRequest) Reset() {
	*x = ProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessRequest) ProtoMessage() {}

func (x *ProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_widget_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessRequest.ProtoReflect.Descriptor instead.
func (*ProcessRequest) Descriptor() ([]byte, []int) {
	return file_widget_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessRequest) GetWidgetName() string {
	if x != nil {
		return x.WidgetName
	}
	return ""
}

func (x *ProcessRequest) GetActionName() string {
	if x != nil {
		return x.ActionName
	}
	return ""
}

func (x *ProcessRequest) GetFiles() map[string][]byte {
	if x != nil {
		return x.Files
	}
	return nil
}

type ProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Redirect     string `protobuf:"bytes,1,opt,name=redirect,proto3" json:"redirect,omitempty"`
	TemplateName string `protobuf:"bytes,2,opt,name=templateName,proto3" json:"templateName,omitempty"`
	Data         []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ProcessResponse) Reset() {
	*x = ProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessResponse) ProtoMessage() {}

func (x *ProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_widget_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessResponse.ProtoReflect.Descriptor instead.
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return file_widget_proto_rawDescGZIP(), []int{4}
}

func (x *ProcessResponse) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *ProcessResponse) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *ProcessResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_widget_proto protoreflect.FileDescriptor

var file_widget_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0x5b, 0x0a, 0x0e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd0, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x44, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x65, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x77, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54,
	0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x41, 0x54, 0x43, 0x48, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x06, 0x12,
	0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05,
	0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41, 0x57, 0x10, 0x09,
	0x32, 0xb4, 0x01, 0x0a, 0x06, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x54, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x75, 0x7a, 0x7a, 0x6c,
	0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70,
	0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x54, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x2e, 0x70,
	0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x68, 0x0a, 0x20, 0x69, 0x6f, 0x2e, 0x64, 0x76,
	0x61, 0x75, 0x6d, 0x6f, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x75, 0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x18, 0x50, 0x75, 0x7a,
	0x7a, 0x6c, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x76, 0x61, 0x75, 0x6d, 0x6f, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x75,
	0x7a, 0x7a, 0x6c, 0x65, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_widget_proto_rawDescOnce sync.Once
	file_widget_proto_rawDescData = file_widget_proto_rawDesc
)

func file_widget_proto_rawDescGZIP() []byte {
	file_widget_proto_rawDescOnce.Do(func() {
		file_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_widget_proto_rawDescData)
	})
	return file_widget_proto_rawDescData
}

var file_widget_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_widget_proto_goTypes = []interface{}{
	(MethodKind)(0),         // 0: puzzlewidgetservice.MethodKind
	(*WidgetRequest)(nil),   // 1: puzzlewidgetservice.WidgetRequest
	(*Action)(nil),          // 2: puzzlewidgetservice.Action
	(*WidgetResponse)(nil),  // 3: puzzlewidgetservice.WidgetResponse
	(*ProcessRequest)(nil),  // 4: puzzlewidgetservice.ProcessRequest
	(*ProcessResponse)(nil), // 5: puzzlewidgetservice.ProcessResponse
	nil,                     // 6: puzzlewidgetservice.ProcessRequest.FilesEntry
}
var file_widget_proto_depIdxs = []int32{
	0, // 0: puzzlewidgetservice.Action.kind:type_name -> puzzlewidgetservice.MethodKind
	2, // 1: puzzlewidgetservice.WidgetResponse.actions:type_name -> puzzlewidgetservice.Action
	6, // 2: puzzlewidgetservice.ProcessRequest.files:type_name -> puzzlewidgetservice.ProcessRequest.FilesEntry
	1, // 3: puzzlewidgetservice.Widget.GetWidget:input_type -> puzzlewidgetservice.WidgetRequest
	4, // 4: puzzlewidgetservice.Widget.Process:input_type -> puzzlewidgetservice.ProcessRequest
	3, // 5: puzzlewidgetservice.Widget.GetWidget:output_type -> puzzlewidgetservice.WidgetResponse
	5, // 6: puzzlewidgetservice.Widget.Process:output_type -> puzzlewidgetservice.ProcessResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_widget_proto_init() }
func file_widget_proto_init() {
	if File_widget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_widget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetRequest); i {
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
		file_widget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_widget_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetResponse); i {
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
		file_widget_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessRequest); i {
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
		file_widget_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_widget_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_widget_proto_goTypes,
		DependencyIndexes: file_widget_proto_depIdxs,
		EnumInfos:         file_widget_proto_enumTypes,
		MessageInfos:      file_widget_proto_msgTypes,
	}.Build()
	File_widget_proto = out.File
	file_widget_proto_rawDesc = nil
	file_widget_proto_goTypes = nil
	file_widget_proto_depIdxs = nil
}
