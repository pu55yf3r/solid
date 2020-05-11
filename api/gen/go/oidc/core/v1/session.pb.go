// Licensed to SolID under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. SolID licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: oidc/core/v1/session.proto

package corev1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string                `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Request *AuthorizationRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Client  *Client               `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidc_core_v1_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_oidc_core_v1_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_oidc_core_v1_session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Session) GetRequest() *AuthorizationRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Session) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

var File_oidc_core_v1_session_proto protoreflect.FileDescriptor

var file_oidc_core_v1_session_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x69,
	0x64, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x6f, 0x69, 0x64, 0x63,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x69, 0x64, 0x63,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x42, 0x15, 0x5a, 0x13, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_oidc_core_v1_session_proto_rawDescOnce sync.Once
	file_oidc_core_v1_session_proto_rawDescData = file_oidc_core_v1_session_proto_rawDesc
)

func file_oidc_core_v1_session_proto_rawDescGZIP() []byte {
	file_oidc_core_v1_session_proto_rawDescOnce.Do(func() {
		file_oidc_core_v1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidc_core_v1_session_proto_rawDescData)
	})
	return file_oidc_core_v1_session_proto_rawDescData
}

var (
	file_oidc_core_v1_session_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
	file_oidc_core_v1_session_proto_goTypes  = []interface{}{
		(*Session)(nil),              // 0: oidc.core.v1.Session
		(*AuthorizationRequest)(nil), // 1: oidc.core.v1.AuthorizationRequest
		(*Client)(nil),               // 2: oidc.core.v1.Client
	}
)

var file_oidc_core_v1_session_proto_depIdxs = []int32{
	1, // 0: oidc.core.v1.Session.request:type_name -> oidc.core.v1.AuthorizationRequest
	2, // 1: oidc.core.v1.Session.client:type_name -> oidc.core.v1.Client
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oidc_core_v1_session_proto_init() }
func file_oidc_core_v1_session_proto_init() {
	if File_oidc_core_v1_session_proto != nil {
		return
	}
	file_oidc_core_v1_core_api_proto_init()
	file_oidc_core_v1_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_oidc_core_v1_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
			RawDescriptor: file_oidc_core_v1_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidc_core_v1_session_proto_goTypes,
		DependencyIndexes: file_oidc_core_v1_session_proto_depIdxs,
		MessageInfos:      file_oidc_core_v1_session_proto_msgTypes,
	}.Build()
	File_oidc_core_v1_session_proto = out.File
	file_oidc_core_v1_session_proto_rawDesc = nil
	file_oidc_core_v1_session_proto_goTypes = nil
	file_oidc_core_v1_session_proto_depIdxs = nil
}