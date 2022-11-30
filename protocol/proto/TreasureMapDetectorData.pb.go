// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: TreasureMapDetectorData.proto

package proto

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

type TreasureMapDetectorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId         uint32    `protobuf:"varint,4,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	CenterPos        *Vector   `protobuf:"bytes,7,opt,name=center_pos,json=centerPos,proto3" json:"center_pos,omitempty"`
	IsRegionDetected bool      `protobuf:"varint,6,opt,name=is_region_detected,json=isRegionDetected,proto3" json:"is_region_detected,omitempty"`
	SpotList         []*Vector `protobuf:"bytes,10,rep,name=spot_list,json=spotList,proto3" json:"spot_list,omitempty"`
	Radius           uint32    `protobuf:"varint,14,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *TreasureMapDetectorData) Reset() {
	*x = TreasureMapDetectorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureMapDetectorData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureMapDetectorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureMapDetectorData) ProtoMessage() {}

func (x *TreasureMapDetectorData) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureMapDetectorData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureMapDetectorData.ProtoReflect.Descriptor instead.
func (*TreasureMapDetectorData) Descriptor() ([]byte, []int) {
	return file_TreasureMapDetectorData_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureMapDetectorData) GetRegionId() uint32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *TreasureMapDetectorData) GetCenterPos() *Vector {
	if x != nil {
		return x.CenterPos
	}
	return nil
}

func (x *TreasureMapDetectorData) GetIsRegionDetected() bool {
	if x != nil {
		return x.IsRegionDetected
	}
	return false
}

func (x *TreasureMapDetectorData) GetSpotList() []*Vector {
	if x != nil {
		return x.SpotList
	}
	return nil
}

func (x *TreasureMapDetectorData) GetRadius() uint32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

var File_TreasureMapDetectorData_proto protoreflect.FileDescriptor

var file_TreasureMapDetectorData_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x17, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x4d, 0x61, 0x70, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x0a, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x09, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69,
	0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x09, 0x73, 0x70, 0x6f,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x70, 0x6f,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TreasureMapDetectorData_proto_rawDescOnce sync.Once
	file_TreasureMapDetectorData_proto_rawDescData = file_TreasureMapDetectorData_proto_rawDesc
)

func file_TreasureMapDetectorData_proto_rawDescGZIP() []byte {
	file_TreasureMapDetectorData_proto_rawDescOnce.Do(func() {
		file_TreasureMapDetectorData_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureMapDetectorData_proto_rawDescData)
	})
	return file_TreasureMapDetectorData_proto_rawDescData
}

var file_TreasureMapDetectorData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TreasureMapDetectorData_proto_goTypes = []interface{}{
	(*TreasureMapDetectorData)(nil), // 0: proto.TreasureMapDetectorData
	(*Vector)(nil),                  // 1: proto.Vector
}
var file_TreasureMapDetectorData_proto_depIdxs = []int32{
	1, // 0: proto.TreasureMapDetectorData.center_pos:type_name -> proto.Vector
	1, // 1: proto.TreasureMapDetectorData.spot_list:type_name -> proto.Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TreasureMapDetectorData_proto_init() }
func file_TreasureMapDetectorData_proto_init() {
	if File_TreasureMapDetectorData_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TreasureMapDetectorData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureMapDetectorData); i {
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
			RawDescriptor: file_TreasureMapDetectorData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureMapDetectorData_proto_goTypes,
		DependencyIndexes: file_TreasureMapDetectorData_proto_depIdxs,
		MessageInfos:      file_TreasureMapDetectorData_proto_msgTypes,
	}.Build()
	File_TreasureMapDetectorData_proto = out.File
	file_TreasureMapDetectorData_proto_rawDesc = nil
	file_TreasureMapDetectorData_proto_goTypes = nil
	file_TreasureMapDetectorData_proto_depIdxs = nil
}