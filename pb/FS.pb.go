// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.5.1
// source: FS.proto

//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file contains protocol buffers that are written into the filesystem

package pb

import (
	proto "google.golang.org/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Reference_Range int32

const (
	Reference_TOP    Reference_Range = 0
	Reference_BOTTOM Reference_Range = 1
)

// Enum value maps for Reference_Range.
var (
	Reference_Range_name = map[int32]string{
		0: "TOP",
		1: "BOTTOM",
	}
	Reference_Range_value = map[string]int32{
		"TOP":    0,
		"BOTTOM": 1,
	}
)

func (x Reference_Range) Enum() *Reference_Range {
	p := new(Reference_Range)
	*p = x
	return p
}

func (x Reference_Range) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reference_Range) Descriptor() protoreflect.EnumDescriptor {
	return file_FS_proto_enumTypes[0].Descriptor()
}

func (Reference_Range) Type() protoreflect.EnumType {
	return &file_FS_proto_enumTypes[0]
}

func (x Reference_Range) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Reference_Range) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Reference_Range(num)
	return nil
}

// Deprecated: Use Reference_Range.Descriptor instead.
func (Reference_Range) EnumDescriptor() ([]byte, []int) {
	return file_FS_proto_rawDescGZIP(), []int{1, 0}
}

//*
// The ${HBASE_ROOTDIR}/hbase.version file content
type HBaseVersionFileContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *string `protobuf:"bytes,1,req,name=version" json:"version,omitempty"`
}

func (x *HBaseVersionFileContent) Reset() {
	*x = HBaseVersionFileContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FS_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HBaseVersionFileContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HBaseVersionFileContent) ProtoMessage() {}

func (x *HBaseVersionFileContent) ProtoReflect() protoreflect.Message {
	mi := &file_FS_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HBaseVersionFileContent.ProtoReflect.Descriptor instead.
func (*HBaseVersionFileContent) Descriptor() ([]byte, []int) {
	return file_FS_proto_rawDescGZIP(), []int{0}
}

func (x *HBaseVersionFileContent) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

//*
// Reference file content used when we split an hfile under a region.
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Splitkey []byte           `protobuf:"bytes,1,req,name=splitkey" json:"splitkey,omitempty"`
	Range    *Reference_Range `protobuf:"varint,2,req,name=range,enum=pb.Reference_Range" json:"range,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FS_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_FS_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_FS_proto_rawDescGZIP(), []int{1}
}

func (x *Reference) GetSplitkey() []byte {
	if x != nil {
		return x.Splitkey
	}
	return nil
}

func (x *Reference) GetRange() Reference_Range {
	if x != nil && x.Range != nil {
		return *x.Range
	}
	return Reference_TOP
}

var File_FS_proto protoreflect.FileDescriptor

var file_FS_proto_rawDesc = []byte{
	0x0a, 0x08, 0x46, 0x53, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x33,
	0x0a, 0x17, 0x48, 0x42, 0x61, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x70, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0c, 0x52, 0x08, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x1c, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x54, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4f, 0x54,
	0x54, 0x4f, 0x4d, 0x10, 0x01, 0x42, 0x3b, 0x0a, 0x2a, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x08, 0x46, 0x53, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0xa0,
	0x01, 0x01,
}

var (
	file_FS_proto_rawDescOnce sync.Once
	file_FS_proto_rawDescData = file_FS_proto_rawDesc
)

func file_FS_proto_rawDescGZIP() []byte {
	file_FS_proto_rawDescOnce.Do(func() {
		file_FS_proto_rawDescData = protoimpl.X.CompressGZIP(file_FS_proto_rawDescData)
	})
	return file_FS_proto_rawDescData
}

var file_FS_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FS_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FS_proto_goTypes = []interface{}{
	(Reference_Range)(0),            // 0: pb.Reference.Range
	(*HBaseVersionFileContent)(nil), // 1: pb.HBaseVersionFileContent
	(*Reference)(nil),               // 2: pb.Reference
}
var file_FS_proto_depIdxs = []int32{
	0, // 0: pb.Reference.range:type_name -> pb.Reference.Range
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FS_proto_init() }
func file_FS_proto_init() {
	if File_FS_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FS_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HBaseVersionFileContent); i {
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
		file_FS_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_FS_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FS_proto_goTypes,
		DependencyIndexes: file_FS_proto_depIdxs,
		EnumInfos:         file_FS_proto_enumTypes,
		MessageInfos:      file_FS_proto_msgTypes,
	}.Build()
	File_FS_proto = out.File
	file_FS_proto_rawDesc = nil
	file_FS_proto_goTypes = nil
	file_FS_proto_depIdxs = nil
}
