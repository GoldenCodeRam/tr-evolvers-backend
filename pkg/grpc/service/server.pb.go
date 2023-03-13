// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/protos/server.proto

package service

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

type EnergyMeterDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial  string `protobuf:"bytes,1,opt,name=serial,proto3" json:"serial,omitempty"`
	BrandId uint32 `protobuf:"varint,2,opt,name=brandId,proto3" json:"brandId,omitempty"`
}

func (x *EnergyMeterDto) Reset() {
	*x = EnergyMeterDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnergyMeterDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnergyMeterDto) ProtoMessage() {}

func (x *EnergyMeterDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnergyMeterDto.ProtoReflect.Descriptor instead.
func (*EnergyMeterDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{0}
}

func (x *EnergyMeterDto) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *EnergyMeterDto) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

type CreateEnergyMeterDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial  string `protobuf:"bytes,1,opt,name=serial,proto3" json:"serial,omitempty"`
	BrandId uint32 `protobuf:"varint,2,opt,name=brandId,proto3" json:"brandId,omitempty"`
}

func (x *CreateEnergyMeterDto) Reset() {
	*x = CreateEnergyMeterDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnergyMeterDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnergyMeterDto) ProtoMessage() {}

func (x *CreateEnergyMeterDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnergyMeterDto.ProtoReflect.Descriptor instead.
func (*CreateEnergyMeterDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEnergyMeterDto) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *CreateEnergyMeterDto) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

type EnergyMeterBrandDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EnergyMeterBrandDto) Reset() {
	*x = EnergyMeterBrandDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnergyMeterBrandDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnergyMeterBrandDto) ProtoMessage() {}

func (x *EnergyMeterBrandDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnergyMeterBrandDto.ProtoReflect.Descriptor instead.
func (*EnergyMeterBrandDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{2}
}

func (x *EnergyMeterBrandDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EnergyMeterBrandDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateEnergyMeterBrandDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand string `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *CreateEnergyMeterBrandDto) Reset() {
	*x = CreateEnergyMeterBrandDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnergyMeterBrandDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnergyMeterBrandDto) ProtoMessage() {}

func (x *CreateEnergyMeterBrandDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnergyMeterBrandDto.ProtoReflect.Descriptor instead.
func (*CreateEnergyMeterBrandDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEnergyMeterBrandDto) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

var File_pkg_protos_server_proto protoreflect.FileDescriptor

var file_pkg_protos_server_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0e, 0x45, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x48, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x13, 0x45, 0x6e, 0x65, 0x72, 0x67,
	0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x31, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x32, 0xa1, 0x01, 0x0a, 0x12, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x4d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x0f, 0x2e, 0x45, 0x6e, 0x65, 0x72, 0x67,
	0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74,
	0x6f, 0x1a, 0x14, 0x2e, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protos_server_proto_rawDescOnce sync.Once
	file_pkg_protos_server_proto_rawDescData = file_pkg_protos_server_proto_rawDesc
)

func file_pkg_protos_server_proto_rawDescGZIP() []byte {
	file_pkg_protos_server_proto_rawDescOnce.Do(func() {
		file_pkg_protos_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protos_server_proto_rawDescData)
	})
	return file_pkg_protos_server_proto_rawDescData
}

var file_pkg_protos_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_protos_server_proto_goTypes = []interface{}{
	(*EnergyMeterDto)(nil),            // 0: EnergyMeterDto
	(*CreateEnergyMeterDto)(nil),      // 1: CreateEnergyMeterDto
	(*EnergyMeterBrandDto)(nil),       // 2: EnergyMeterBrandDto
	(*CreateEnergyMeterBrandDto)(nil), // 3: CreateEnergyMeterBrandDto
}
var file_pkg_protos_server_proto_depIdxs = []int32{
	1, // 0: EnergyMeterService.CreateEnergyMeter:input_type -> CreateEnergyMeterDto
	3, // 1: EnergyMeterService.CreateEnergyMeterBrand:input_type -> CreateEnergyMeterBrandDto
	0, // 2: EnergyMeterService.CreateEnergyMeter:output_type -> EnergyMeterDto
	2, // 3: EnergyMeterService.CreateEnergyMeterBrand:output_type -> EnergyMeterBrandDto
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_protos_server_proto_init() }
func file_pkg_protos_server_proto_init() {
	if File_pkg_protos_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnergyMeterDto); i {
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
		file_pkg_protos_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnergyMeterDto); i {
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
		file_pkg_protos_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnergyMeterBrandDto); i {
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
		file_pkg_protos_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnergyMeterBrandDto); i {
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
			RawDescriptor: file_pkg_protos_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protos_server_proto_goTypes,
		DependencyIndexes: file_pkg_protos_server_proto_depIdxs,
		MessageInfos:      file_pkg_protos_server_proto_msgTypes,
	}.Build()
	File_pkg_protos_server_proto = out.File
	file_pkg_protos_server_proto_rawDesc = nil
	file_pkg_protos_server_proto_goTypes = nil
	file_pkg_protos_server_proto_depIdxs = nil
}