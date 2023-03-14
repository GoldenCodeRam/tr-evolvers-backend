// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pkg/protos/server.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Dto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial  string `protobuf:"bytes,1,opt,name=serial,proto3" json:"serial,omitempty"`
	BrandId uint32 `protobuf:"varint,2,opt,name=brandId,proto3" json:"brandId,omitempty"`
}

func (x *Dto) Reset() {
	*x = Dto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dto) ProtoMessage() {}

func (x *Dto) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Dto.ProtoReflect.Descriptor instead.
func (*Dto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{0}
}

func (x *Dto) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Dto) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

type CreateDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial  string `protobuf:"bytes,1,opt,name=serial,proto3" json:"serial,omitempty"`
	BrandId uint32 `protobuf:"varint,2,opt,name=brandId,proto3" json:"brandId,omitempty"`
}

func (x *CreateDto) Reset() {
	*x = CreateDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDto) ProtoMessage() {}

func (x *CreateDto) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateDto.ProtoReflect.Descriptor instead.
func (*CreateDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDto) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *CreateDto) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

type BrandDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *BrandDto) Reset() {
	*x = BrandDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandDto) ProtoMessage() {}

func (x *BrandDto) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BrandDto.ProtoReflect.Descriptor instead.
func (*BrandDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{2}
}

func (x *BrandDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BrandDto) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

type CreateBrandDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand string `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *CreateBrandDto) Reset() {
	*x = CreateBrandDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandDto) ProtoMessage() {}

func (x *CreateBrandDto) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateBrandDto.ProtoReflect.Descriptor instead.
func (*CreateBrandDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{3}
}

func (x *CreateBrandDto) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

type InstallationDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address          string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	InstallationDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=installationDate,proto3" json:"installationDate,omitempty"`
	RetirementDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=retirementDate,proto3" json:"retirementDate,omitempty"`
	Lines            uint32                 `protobuf:"varint,5,opt,name=lines,proto3" json:"lines,omitempty"`
	IsActive         bool                   `protobuf:"varint,6,opt,name=isActive,proto3" json:"isActive,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *InstallationDto) Reset() {
	*x = InstallationDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallationDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallationDto) ProtoMessage() {}

func (x *InstallationDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallationDto.ProtoReflect.Descriptor instead.
func (*InstallationDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{4}
}

func (x *InstallationDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InstallationDto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InstallationDto) GetInstallationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.InstallationDate
	}
	return nil
}

func (x *InstallationDto) GetRetirementDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RetirementDate
	}
	return nil
}

func (x *InstallationDto) GetLines() uint32 {
	if x != nil {
		return x.Lines
	}
	return 0
}

func (x *InstallationDto) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *InstallationDto) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateInstallationDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address          string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	InstallationDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=installationDate,proto3" json:"installationDate,omitempty"`
	RetirementDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=retirementDate,proto3" json:"retirementDate,omitempty"`
	Lines            uint32                 `protobuf:"varint,4,opt,name=lines,proto3" json:"lines,omitempty"`
	IsActive         bool                   `protobuf:"varint,5,opt,name=isActive,proto3" json:"isActive,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	SerialId         string                 `protobuf:"bytes,7,opt,name=serialId,proto3" json:"serialId,omitempty"`
}

func (x *CreateInstallationDto) Reset() {
	*x = CreateInstallationDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstallationDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstallationDto) ProtoMessage() {}

func (x *CreateInstallationDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstallationDto.ProtoReflect.Descriptor instead.
func (*CreateInstallationDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{5}
}

func (x *CreateInstallationDto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateInstallationDto) GetInstallationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.InstallationDate
	}
	return nil
}

func (x *CreateInstallationDto) GetRetirementDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RetirementDate
	}
	return nil
}

func (x *CreateInstallationDto) GetLines() uint32 {
	if x != nil {
		return x.Lines
	}
	return 0
}

func (x *CreateInstallationDto) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *CreateInstallationDto) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateInstallationDto) GetSerialId() string {
	if x != nil {
		return x.SerialId
	}
	return ""
}

type UpdateInstallationDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address        string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	RetirementDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=retirementDate,proto3" json:"retirementDate,omitempty"`
	Lines          uint32                 `protobuf:"varint,4,opt,name=lines,proto3" json:"lines,omitempty"`
	IsActive       bool                   `protobuf:"varint,5,opt,name=isActive,proto3" json:"isActive,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *UpdateInstallationDto) Reset() {
	*x = UpdateInstallationDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInstallationDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInstallationDto) ProtoMessage() {}

func (x *UpdateInstallationDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInstallationDto.ProtoReflect.Descriptor instead.
func (*UpdateInstallationDto) Descriptor() ([]byte, []int) {
	return file_pkg_protos_server_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateInstallationDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateInstallationDto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateInstallationDto) GetRetirementDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RetirementDate
	}
	return nil
}

func (x *UpdateInstallationDto) GetLines() uint32 {
	if x != nil {
		return x.Lines
	}
	return 0
}

func (x *UpdateInstallationDto) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *UpdateInstallationDto) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_pkg_protos_server_proto protoreflect.FileDescriptor

var file_pkg_protos_server_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x47, 0x72, 0x70, 0x63, 0x45,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x03, 0x44,
	0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x74,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0xb3, 0x02,
	0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xc5, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x42, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xf1, 0x01, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x42, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32,
	0x90, 0x06, 0x0a, 0x12, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x14, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x44,
	0x74, 0x6f, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x74,
	0x6f, 0x1a, 0x14, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x74, 0x6f, 0x1a, 0x14, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x74, 0x6f, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x19, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x19, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67,
	0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x19, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f,
	0x22, 0x00, 0x12, 0x45, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x19, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x19, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x20, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f,
	0x1a, 0x20, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x74, 0x6f, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x74, 0x6f, 0x1a, 0x20, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x20,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x4d, 0x65, 0x74, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f,
	0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_pkg_protos_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_protos_server_proto_goTypes = []interface{}{
	(*Dto)(nil),                   // 0: GrpcEnergyMeter.Dto
	(*CreateDto)(nil),             // 1: GrpcEnergyMeter.CreateDto
	(*BrandDto)(nil),              // 2: GrpcEnergyMeter.BrandDto
	(*CreateBrandDto)(nil),        // 3: GrpcEnergyMeter.CreateBrandDto
	(*InstallationDto)(nil),       // 4: GrpcEnergyMeter.InstallationDto
	(*CreateInstallationDto)(nil), // 5: GrpcEnergyMeter.CreateInstallationDto
	(*UpdateInstallationDto)(nil), // 6: GrpcEnergyMeter.UpdateInstallationDto
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_pkg_protos_server_proto_depIdxs = []int32{
	7,  // 0: GrpcEnergyMeter.InstallationDto.installationDate:type_name -> google.protobuf.Timestamp
	7,  // 1: GrpcEnergyMeter.InstallationDto.retirementDate:type_name -> google.protobuf.Timestamp
	7,  // 2: GrpcEnergyMeter.InstallationDto.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 3: GrpcEnergyMeter.CreateInstallationDto.installationDate:type_name -> google.protobuf.Timestamp
	7,  // 4: GrpcEnergyMeter.CreateInstallationDto.retirementDate:type_name -> google.protobuf.Timestamp
	7,  // 5: GrpcEnergyMeter.CreateInstallationDto.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 6: GrpcEnergyMeter.UpdateInstallationDto.retirementDate:type_name -> google.protobuf.Timestamp
	7,  // 7: GrpcEnergyMeter.UpdateInstallationDto.createdAt:type_name -> google.protobuf.Timestamp
	1,  // 8: GrpcEnergyMeter.EnergyMeterService.Create:input_type -> GrpcEnergyMeter.CreateDto
	0,  // 9: GrpcEnergyMeter.EnergyMeterService.Get:input_type -> GrpcEnergyMeter.Dto
	0,  // 10: GrpcEnergyMeter.EnergyMeterService.Delete:input_type -> GrpcEnergyMeter.Dto
	2,  // 11: GrpcEnergyMeter.EnergyMeterService.CreateBrand:input_type -> GrpcEnergyMeter.BrandDto
	3,  // 12: GrpcEnergyMeter.EnergyMeterService.GetBrand:input_type -> GrpcEnergyMeter.CreateBrandDto
	2,  // 13: GrpcEnergyMeter.EnergyMeterService.DeleteBrand:input_type -> GrpcEnergyMeter.BrandDto
	5,  // 14: GrpcEnergyMeter.EnergyMeterService.CreateInstallation:input_type -> GrpcEnergyMeter.CreateInstallationDto
	4,  // 15: GrpcEnergyMeter.EnergyMeterService.GetInstallation:input_type -> GrpcEnergyMeter.InstallationDto
	6,  // 16: GrpcEnergyMeter.EnergyMeterService.UpdateInstallation:input_type -> GrpcEnergyMeter.UpdateInstallationDto
	4,  // 17: GrpcEnergyMeter.EnergyMeterService.DeleteInstallation:input_type -> GrpcEnergyMeter.InstallationDto
	0,  // 18: GrpcEnergyMeter.EnergyMeterService.Create:output_type -> GrpcEnergyMeter.Dto
	0,  // 19: GrpcEnergyMeter.EnergyMeterService.Get:output_type -> GrpcEnergyMeter.Dto
	0,  // 20: GrpcEnergyMeter.EnergyMeterService.Delete:output_type -> GrpcEnergyMeter.Dto
	2,  // 21: GrpcEnergyMeter.EnergyMeterService.CreateBrand:output_type -> GrpcEnergyMeter.BrandDto
	2,  // 22: GrpcEnergyMeter.EnergyMeterService.GetBrand:output_type -> GrpcEnergyMeter.BrandDto
	2,  // 23: GrpcEnergyMeter.EnergyMeterService.DeleteBrand:output_type -> GrpcEnergyMeter.BrandDto
	4,  // 24: GrpcEnergyMeter.EnergyMeterService.CreateInstallation:output_type -> GrpcEnergyMeter.InstallationDto
	4,  // 25: GrpcEnergyMeter.EnergyMeterService.GetInstallation:output_type -> GrpcEnergyMeter.InstallationDto
	4,  // 26: GrpcEnergyMeter.EnergyMeterService.UpdateInstallation:output_type -> GrpcEnergyMeter.InstallationDto
	4,  // 27: GrpcEnergyMeter.EnergyMeterService.DeleteInstallation:output_type -> GrpcEnergyMeter.InstallationDto
	18, // [18:28] is the sub-list for method output_type
	8,  // [8:18] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_protos_server_proto_init() }
func file_pkg_protos_server_proto_init() {
	if File_pkg_protos_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dto); i {
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
			switch v := v.(*CreateDto); i {
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
			switch v := v.(*BrandDto); i {
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
			switch v := v.(*CreateBrandDto); i {
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
		file_pkg_protos_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallationDto); i {
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
		file_pkg_protos_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstallationDto); i {
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
		file_pkg_protos_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInstallationDto); i {
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
			NumMessages:   7,
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