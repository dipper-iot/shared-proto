// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: services/otp/otp.proto

package otp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type DeviceType int32

const (
	DeviceType_WEB     DeviceType = 0
	DeviceType_ANDROID DeviceType = 1
	DeviceType_IOS     DeviceType = 2
	DeviceType_WINDOWS DeviceType = 3
	DeviceType_LINUX   DeviceType = 4
	DeviceType_UNKNOWN DeviceType = 5
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		0: "WEB",
		1: "ANDROID",
		2: "IOS",
		3: "WINDOWS",
		4: "LINUX",
		5: "UNKNOWN",
	}
	DeviceType_value = map[string]int32{
		"WEB":     0,
		"ANDROID": 1,
		"IOS":     2,
		"WINDOWS": 3,
		"LINUX":   4,
		"UNKNOWN": 5,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_services_otp_otp_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_services_otp_otp_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{0}
}

type OtpType int32

const (
	OtpType_EMAIL OtpType = 0
	OtpType_PHONE OtpType = 1
)

// Enum value maps for OtpType.
var (
	OtpType_name = map[int32]string{
		0: "EMAIL",
		1: "PHONE",
	}
	OtpType_value = map[string]int32{
		"EMAIL": 0,
		"PHONE": 1,
	}
)

func (x OtpType) Enum() *OtpType {
	p := new(OtpType)
	*p = x
	return p
}

func (x OtpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OtpType) Descriptor() protoreflect.EnumDescriptor {
	return file_services_otp_otp_proto_enumTypes[1].Descriptor()
}

func (OtpType) Type() protoreflect.EnumType {
	return &file_services_otp_otp_proto_enumTypes[1]
}

func (x OtpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OtpType.Descriptor instead.
func (OtpType) EnumDescriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{1}
}

type IssuerType int32

const (
	IssuerType_ADMIN IssuerType = 0
	IssuerType_USER  IssuerType = 1
	IssuerType_BASE  IssuerType = 2
)

// Enum value maps for IssuerType.
var (
	IssuerType_name = map[int32]string{
		0: "ADMIN",
		1: "USER",
		2: "BASE",
	}
	IssuerType_value = map[string]int32{
		"ADMIN": 0,
		"USER":  1,
		"BASE":  2,
	}
)

func (x IssuerType) Enum() *IssuerType {
	p := new(IssuerType)
	*p = x
	return p
}

func (x IssuerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssuerType) Descriptor() protoreflect.EnumDescriptor {
	return file_services_otp_otp_proto_enumTypes[2].Descriptor()
}

func (IssuerType) Type() protoreflect.EnumType {
	return &file_services_otp_otp_proto_enumTypes[2]
}

func (x IssuerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssuerType.Descriptor instead.
func (IssuerType) EnumDescriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{2}
}

type OtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Otp        *OtpType    `protobuf:"varint,2,opt,name=otp,proto3,enum=otp.OtpType,oneof" json:"otp,omitempty"`
	Issuer     *IssuerType `protobuf:"varint,3,opt,name=issuer,proto3,enum=otp.IssuerType,oneof" json:"issuer,omitempty"`
	DeviceType *DeviceType `protobuf:"varint,4,opt,name=device_type,json=deviceType,proto3,enum=otp.DeviceType,oneof" json:"device_type,omitempty"`
	Session    *string     `protobuf:"bytes,5,opt,name=session,proto3,oneof" json:"session,omitempty"`
	UserId     *uint64     `protobuf:"varint,6,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
}

func (x *OtpRequest) Reset() {
	*x = OtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_otp_otp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtpRequest) ProtoMessage() {}

func (x *OtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_otp_otp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtpRequest.ProtoReflect.Descriptor instead.
func (*OtpRequest) Descriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{0}
}

func (x *OtpRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *OtpRequest) GetOtp() OtpType {
	if x != nil && x.Otp != nil {
		return *x.Otp
	}
	return OtpType_EMAIL
}

func (x *OtpRequest) GetIssuer() IssuerType {
	if x != nil && x.Issuer != nil {
		return *x.Issuer
	}
	return IssuerType_ADMIN
}

func (x *OtpRequest) GetDeviceType() DeviceType {
	if x != nil && x.DeviceType != nil {
		return *x.DeviceType
	}
	return DeviceType_WEB
}

func (x *OtpRequest) GetSession() string {
	if x != nil && x.Session != nil {
		return *x.Session
	}
	return ""
}

func (x *OtpRequest) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

type Otp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token      string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Otp        OtpType                `protobuf:"varint,3,opt,name=otp,proto3,enum=otp.OtpType" json:"otp,omitempty"`
	Issuer     IssuerType             `protobuf:"varint,4,opt,name=issuer,proto3,enum=otp.IssuerType" json:"issuer,omitempty"`
	DeviceType DeviceType             `protobuf:"varint,5,opt,name=device_type,json=deviceType,proto3,enum=otp.DeviceType" json:"device_type,omitempty"`
	Session    string                 `protobuf:"bytes,6,opt,name=session,proto3" json:"session,omitempty"`
	UserId     *uint64                `protobuf:"varint,7,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	Exp        *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=exp,proto3" json:"exp,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Otp) Reset() {
	*x = Otp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_otp_otp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Otp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Otp) ProtoMessage() {}

func (x *Otp) ProtoReflect() protoreflect.Message {
	mi := &file_services_otp_otp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Otp.ProtoReflect.Descriptor instead.
func (*Otp) Descriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{1}
}

func (x *Otp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Otp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Otp) GetOtp() OtpType {
	if x != nil {
		return x.Otp
	}
	return OtpType_EMAIL
}

func (x *Otp) GetIssuer() IssuerType {
	if x != nil {
		return x.Issuer
	}
	return IssuerType_ADMIN
}

func (x *Otp) GetDeviceType() DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_WEB
}

func (x *Otp) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *Otp) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *Otp) GetExp() *timestamppb.Timestamp {
	if x != nil {
		return x.Exp
	}
	return nil
}

func (x *Otp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Otp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeExp    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_exp,json=timeExp,proto3" json:"time_exp,omitempty"`
	Otp        OtpType                `protobuf:"varint,2,opt,name=otp,proto3,enum=otp.OtpType" json:"otp,omitempty"`
	Issuer     *IssuerType            `protobuf:"varint,3,opt,name=issuer,proto3,enum=otp.IssuerType,oneof" json:"issuer,omitempty"`
	DeviceType *DeviceType            `protobuf:"varint,4,opt,name=device_type,json=deviceType,proto3,enum=otp.DeviceType,oneof" json:"device_type,omitempty"`
	UserId     *uint64                `protobuf:"varint,6,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
}

func (x *CreateOtpRequest) Reset() {
	*x = CreateOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_otp_otp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOtpRequest) ProtoMessage() {}

func (x *CreateOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_otp_otp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOtpRequest.ProtoReflect.Descriptor instead.
func (*CreateOtpRequest) Descriptor() ([]byte, []int) {
	return file_services_otp_otp_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOtpRequest) GetTimeExp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeExp
	}
	return nil
}

func (x *CreateOtpRequest) GetOtp() OtpType {
	if x != nil {
		return x.Otp
	}
	return OtpType_EMAIL
}

func (x *CreateOtpRequest) GetIssuer() IssuerType {
	if x != nil && x.Issuer != nil {
		return *x.Issuer
	}
	return IssuerType_ADMIN
}

func (x *CreateOtpRequest) GetDeviceType() DeviceType {
	if x != nil && x.DeviceType != nil {
		return *x.DeviceType
	}
	return DeviceType_WEB
}

func (x *CreateOtpRequest) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

var File_services_otp_otp_proto protoreflect.FileDescriptor

var file_services_otp_otp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x74, 0x70, 0x2f, 0x6f,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x74, 0x70, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x0a,
	0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x74, 0x70, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x03, 0x6f,
	0x74, 0x70, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x48, 0x01, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x02, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6f, 0x74, 0x70, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x22, 0x8e, 0x03, 0x0a, 0x03, 0x4f, 0x74, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x74, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x6f, 0x74, 0x70,
	0x12, 0x27, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x78,
	0x70, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x93, 0x02, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x74,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x65, 0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x12,
	0x1e, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x6f,
	0x74, 0x70, 0x2e, 0x4f, 0x74, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12,
	0x2c, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x48, 0x00, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2a, 0x50, 0x0a, 0x0a, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x49, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57,
	0x53, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x2a, 0x1f, 0x0a, 0x07, 0x4f,
	0x74, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x2a, 0x2b, 0x0a, 0x0a,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x42, 0x41, 0x53, 0x45, 0x10, 0x02, 0x32, 0x92, 0x01, 0x0a, 0x0a, 0x4f, 0x74,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x74, 0x70, 0x12, 0x15, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x6f,
	0x74, 0x70, 0x2e, 0x4f, 0x74, 0x70, 0x12, 0x23, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x70,
	0x12, 0x0f, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x08, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x74, 0x70, 0x12, 0x31, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x4f, 0x74, 0x70, 0x12, 0x0f, 0x2e, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x74, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0e,
	0x5a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x74, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_otp_otp_proto_rawDescOnce sync.Once
	file_services_otp_otp_proto_rawDescData = file_services_otp_otp_proto_rawDesc
)

func file_services_otp_otp_proto_rawDescGZIP() []byte {
	file_services_otp_otp_proto_rawDescOnce.Do(func() {
		file_services_otp_otp_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_otp_otp_proto_rawDescData)
	})
	return file_services_otp_otp_proto_rawDescData
}

var file_services_otp_otp_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_services_otp_otp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_otp_otp_proto_goTypes = []interface{}{
	(DeviceType)(0),               // 0: otp.DeviceType
	(OtpType)(0),                  // 1: otp.OtpType
	(IssuerType)(0),               // 2: otp.IssuerType
	(*OtpRequest)(nil),            // 3: otp.OtpRequest
	(*Otp)(nil),                   // 4: otp.Otp
	(*CreateOtpRequest)(nil),      // 5: otp.CreateOtpRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_services_otp_otp_proto_depIdxs = []int32{
	1,  // 0: otp.OtpRequest.otp:type_name -> otp.OtpType
	2,  // 1: otp.OtpRequest.issuer:type_name -> otp.IssuerType
	0,  // 2: otp.OtpRequest.device_type:type_name -> otp.DeviceType
	1,  // 3: otp.Otp.otp:type_name -> otp.OtpType
	2,  // 4: otp.Otp.issuer:type_name -> otp.IssuerType
	0,  // 5: otp.Otp.device_type:type_name -> otp.DeviceType
	6,  // 6: otp.Otp.exp:type_name -> google.protobuf.Timestamp
	6,  // 7: otp.Otp.created_at:type_name -> google.protobuf.Timestamp
	6,  // 8: otp.Otp.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 9: otp.CreateOtpRequest.time_exp:type_name -> google.protobuf.Timestamp
	1,  // 10: otp.CreateOtpRequest.otp:type_name -> otp.OtpType
	2,  // 11: otp.CreateOtpRequest.issuer:type_name -> otp.IssuerType
	0,  // 12: otp.CreateOtpRequest.device_type:type_name -> otp.DeviceType
	5,  // 13: otp.OtpService.CreateOtp:input_type -> otp.CreateOtpRequest
	3,  // 14: otp.OtpService.GetOtp:input_type -> otp.OtpRequest
	3,  // 15: otp.OtpService.DelOtp:input_type -> otp.OtpRequest
	4,  // 16: otp.OtpService.CreateOtp:output_type -> otp.Otp
	4,  // 17: otp.OtpService.GetOtp:output_type -> otp.Otp
	7,  // 18: otp.OtpService.DelOtp:output_type -> google.protobuf.Empty
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_services_otp_otp_proto_init() }
func file_services_otp_otp_proto_init() {
	if File_services_otp_otp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_otp_otp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtpRequest); i {
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
		file_services_otp_otp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Otp); i {
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
		file_services_otp_otp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOtpRequest); i {
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
	file_services_otp_otp_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_services_otp_otp_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_services_otp_otp_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_otp_otp_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_otp_otp_proto_goTypes,
		DependencyIndexes: file_services_otp_otp_proto_depIdxs,
		EnumInfos:         file_services_otp_otp_proto_enumTypes,
		MessageInfos:      file_services_otp_otp_proto_msgTypes,
	}.Build()
	File_services_otp_otp_proto = out.File
	file_services_otp_otp_proto_rawDesc = nil
	file_services_otp_otp_proto_goTypes = nil
	file_services_otp_otp_proto_depIdxs = nil
}
