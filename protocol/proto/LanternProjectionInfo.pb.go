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
// source: LanternProjectionInfo.proto

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

type LanternProjectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewSwitchTipsList []ClientInputType             `protobuf:"varint,12,rep,packed,name=view_switch_tips_list,json=viewSwitchTipsList,proto3,enum=proto.ClientInputType" json:"view_switch_tips_list,omitempty"`
	LevelList          []*LanternProjectionLevelInfo `protobuf:"bytes,6,rep,name=level_list,json=levelList,proto3" json:"level_list,omitempty"`
	OpenStageList      []uint32                      `protobuf:"varint,10,rep,packed,name=open_stage_list,json=openStageList,proto3" json:"open_stage_list,omitempty"`
	ViewInputTipsList  []ClientInputType             `protobuf:"varint,13,rep,packed,name=view_input_tips_list,json=viewInputTipsList,proto3,enum=proto.ClientInputType" json:"view_input_tips_list,omitempty"`
}

func (x *LanternProjectionInfo) Reset() {
	*x = LanternProjectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LanternProjectionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanternProjectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanternProjectionInfo) ProtoMessage() {}

func (x *LanternProjectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_LanternProjectionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanternProjectionInfo.ProtoReflect.Descriptor instead.
func (*LanternProjectionInfo) Descriptor() ([]byte, []int) {
	return file_LanternProjectionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *LanternProjectionInfo) GetViewSwitchTipsList() []ClientInputType {
	if x != nil {
		return x.ViewSwitchTipsList
	}
	return nil
}

func (x *LanternProjectionInfo) GetLevelList() []*LanternProjectionLevelInfo {
	if x != nil {
		return x.LevelList
	}
	return nil
}

func (x *LanternProjectionInfo) GetOpenStageList() []uint32 {
	if x != nil {
		return x.OpenStageList
	}
	return nil
}

func (x *LanternProjectionInfo) GetViewInputTipsList() []ClientInputType {
	if x != nil {
		return x.ViewInputTipsList
	}
	return nil
}

var File_LanternProjectionInfo_proto protoreflect.FileDescriptor

var file_LanternProjectionInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x4c, 0x61, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02,
	0x0a, 0x15, 0x4c, 0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x15, 0x76, 0x69, 0x65, 0x77, 0x5f,
	0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12,
	0x76, 0x69, 0x65, 0x77, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x54, 0x69, 0x70, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x61, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x6f,
	0x70, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x14,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x69, 0x70, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x11, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x69, 0x70,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LanternProjectionInfo_proto_rawDescOnce sync.Once
	file_LanternProjectionInfo_proto_rawDescData = file_LanternProjectionInfo_proto_rawDesc
)

func file_LanternProjectionInfo_proto_rawDescGZIP() []byte {
	file_LanternProjectionInfo_proto_rawDescOnce.Do(func() {
		file_LanternProjectionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_LanternProjectionInfo_proto_rawDescData)
	})
	return file_LanternProjectionInfo_proto_rawDescData
}

var file_LanternProjectionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LanternProjectionInfo_proto_goTypes = []interface{}{
	(*LanternProjectionInfo)(nil),      // 0: proto.LanternProjectionInfo
	(ClientInputType)(0),               // 1: proto.ClientInputType
	(*LanternProjectionLevelInfo)(nil), // 2: proto.LanternProjectionLevelInfo
}
var file_LanternProjectionInfo_proto_depIdxs = []int32{
	1, // 0: proto.LanternProjectionInfo.view_switch_tips_list:type_name -> proto.ClientInputType
	2, // 1: proto.LanternProjectionInfo.level_list:type_name -> proto.LanternProjectionLevelInfo
	1, // 2: proto.LanternProjectionInfo.view_input_tips_list:type_name -> proto.ClientInputType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_LanternProjectionInfo_proto_init() }
func file_LanternProjectionInfo_proto_init() {
	if File_LanternProjectionInfo_proto != nil {
		return
	}
	file_ClientInputType_proto_init()
	file_LanternProjectionLevelInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LanternProjectionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanternProjectionInfo); i {
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
			RawDescriptor: file_LanternProjectionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LanternProjectionInfo_proto_goTypes,
		DependencyIndexes: file_LanternProjectionInfo_proto_depIdxs,
		MessageInfos:      file_LanternProjectionInfo_proto_msgTypes,
	}.Build()
	File_LanternProjectionInfo_proto = out.File
	file_LanternProjectionInfo_proto_rawDesc = nil
	file_LanternProjectionInfo_proto_goTypes = nil
	file_LanternProjectionInfo_proto_depIdxs = nil
}