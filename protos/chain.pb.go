// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos/chain.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

type TailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TailRequest) Reset() {
	*x = TailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TailRequest) ProtoMessage() {}

func (x *TailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TailRequest.ProtoReflect.Descriptor instead.
func (*TailRequest) Descriptor() ([]byte, []int) {
	return file_protos_chain_proto_rawDescGZIP(), []int{0}
}

type TailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TailResponse) Reset() {
	*x = TailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TailResponse) ProtoMessage() {}

func (x *TailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TailResponse.ProtoReflect.Descriptor instead.
func (*TailResponse) Descriptor() ([]byte, []int) {
	return file_protos_chain_proto_rawDescGZIP(), []int{1}
}

func (x *TailResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	TailAddress string `protobuf:"bytes,2,opt,name=tailAddress,proto3" json:"tailAddress,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_protos_chain_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *JoinRequest) GetTailAddress() string {
	if x != nil {
		return x.TailAddress
	}
	return ""
}

type NeighbourInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PredAddress string `protobuf:"bytes,1,opt,name=predAddress,proto3" json:"predAddress,omitempty"`
	SuccAddress string `protobuf:"bytes,2,opt,name=succAddress,proto3" json:"succAddress,omitempty"`
}

func (x *NeighbourInfo) Reset() {
	*x = NeighbourInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeighbourInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeighbourInfo) ProtoMessage() {}

func (x *NeighbourInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeighbourInfo.ProtoReflect.Descriptor instead.
func (*NeighbourInfo) Descriptor() ([]byte, []int) {
	return file_protos_chain_proto_rawDescGZIP(), []int{3}
}

func (x *NeighbourInfo) GetPredAddress() string {
	if x != nil {
		return x.PredAddress
	}
	return ""
}

func (x *NeighbourInfo) GetSuccAddress() string {
	if x != nil {
		return x.SuccAddress
	}
	return ""
}

var File_protos_chain_proto protoreflect.FileDescriptor

var file_protos_chain_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x13, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x28, 0x0a, 0x0c, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x0b, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x53, 0x0a, 0x0d, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x65,
	0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x63, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x75, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xad, 0x01, 0x0a, 0x05, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x12,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x75, 0x72, 0x73, 0x12,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f,
	0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x44, 0x4a, 0x46, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_chain_proto_rawDescOnce sync.Once
	file_protos_chain_proto_rawDescData = file_protos_chain_proto_rawDesc
)

func file_protos_chain_proto_rawDescGZIP() []byte {
	file_protos_chain_proto_rawDescOnce.Do(func() {
		file_protos_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_chain_proto_rawDescData)
	})
	return file_protos_chain_proto_rawDescData
}

var file_protos_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_chain_proto_goTypes = []interface{}{
	(*TailRequest)(nil),   // 0: protos.TailRequest
	(*TailResponse)(nil),  // 1: protos.TailResponse
	(*JoinRequest)(nil),   // 2: protos.JoinRequest
	(*NeighbourInfo)(nil), // 3: protos.NeighbourInfo
	(*OkResponse)(nil),    // 4: protos.OkResponse
}
var file_protos_chain_proto_depIdxs = []int32{
	0, // 0: protos.Chain.GetTail:input_type -> protos.TailRequest
	2, // 1: protos.Chain.Join:input_type -> protos.JoinRequest
	3, // 2: protos.Chain.UpdateNeighbours:input_type -> protos.NeighbourInfo
	1, // 3: protos.Chain.GetTail:output_type -> protos.TailResponse
	4, // 4: protos.Chain.Join:output_type -> protos.OkResponse
	4, // 5: protos.Chain.UpdateNeighbours:output_type -> protos.OkResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_chain_proto_init() }
func file_protos_chain_proto_init() {
	if File_protos_chain_proto != nil {
		return
	}
	file_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TailRequest); i {
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
		file_protos_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TailResponse); i {
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
		file_protos_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_protos_chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeighbourInfo); i {
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
			RawDescriptor: file_protos_chain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_chain_proto_goTypes,
		DependencyIndexes: file_protos_chain_proto_depIdxs,
		MessageInfos:      file_protos_chain_proto_msgTypes,
	}.Build()
	File_protos_chain_proto = out.File
	file_protos_chain_proto_rawDesc = nil
	file_protos_chain_proto_goTypes = nil
	file_protos_chain_proto_depIdxs = nil
}
