// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: sso/auth/service.proto

package auth

import (
	user "cinematic.back/sso/user"
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

type UserContactTypeEnum int32

const (
	UserContactTypeEnum_USER_CONTACT_TYPE_ENUM_UNSPECIFIED UserContactTypeEnum = 0
	UserContactTypeEnum_USER_CONTACT_TYPE_ENUM_LOGIN       UserContactTypeEnum = 1
)

// Enum value maps for UserContactTypeEnum.
var (
	UserContactTypeEnum_name = map[int32]string{
		0: "USER_CONTACT_TYPE_ENUM_UNSPECIFIED",
		1: "USER_CONTACT_TYPE_ENUM_LOGIN",
	}
	UserContactTypeEnum_value = map[string]int32{
		"USER_CONTACT_TYPE_ENUM_UNSPECIFIED": 0,
		"USER_CONTACT_TYPE_ENUM_LOGIN":       1,
	}
)

func (x UserContactTypeEnum) Enum() *UserContactTypeEnum {
	p := new(UserContactTypeEnum)
	*p = x
	return p
}

func (x UserContactTypeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserContactTypeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_sso_auth_service_proto_enumTypes[0].Descriptor()
}

func (UserContactTypeEnum) Type() protoreflect.EnumType {
	return &file_sso_auth_service_proto_enumTypes[0]
}

func (x UserContactTypeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserContactTypeEnum.Descriptor instead.
func (UserContactTypeEnum) EnumDescriptor() ([]byte, []int) {
	return file_sso_auth_service_proto_rawDescGZIP(), []int{0}
}

type RegisterUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credentials *UserCredentials `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_sso_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterUserRequest) GetCredentials() *UserCredentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type RegisterUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *user.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AccessToken  *Token     `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken *Token     `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_sso_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterUserResponse) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RegisterUserResponse) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *RegisterUserResponse) GetRefreshToken() *Token {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

type LoginByCredsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credentials *UserCredentials `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (x *LoginByCredsReq) Reset() {
	*x = LoginByCredsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByCredsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByCredsReq) ProtoMessage() {}

func (x *LoginByCredsReq) ProtoReflect() protoreflect.Message {
	mi := &file_sso_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByCredsReq.ProtoReflect.Descriptor instead.
func (*LoginByCredsReq) Descriptor() ([]byte, []int) {
	return file_sso_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *LoginByCredsReq) GetCredentials() *UserCredentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type LoginByCredsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  *Token `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken *Token `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *LoginByCredsRes) Reset() {
	*x = LoginByCredsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginByCredsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginByCredsRes) ProtoMessage() {}

func (x *LoginByCredsRes) ProtoReflect() protoreflect.Message {
	mi := &file_sso_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginByCredsRes.ProtoReflect.Descriptor instead.
func (*LoginByCredsRes) Descriptor() ([]byte, []int) {
	return file_sso_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *LoginByCredsRes) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *LoginByCredsRes) GetRefreshToken() *Token {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

type RefreshTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenReq) Reset() {
	*x = RefreshTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenReq) ProtoMessage() {}

func (x *RefreshTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_sso_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenReq.ProtoReflect.Descriptor instead.
func (*RefreshTokenReq) Descriptor() ([]byte, []int) {
	return file_sso_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken *Token `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *RefreshTokenRes) Reset() {
	*x = RefreshTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRes) ProtoMessage() {}

func (x *RefreshTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_sso_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRes.ProtoReflect.Descriptor instead.
func (*RefreshTokenRes) Descriptor() ([]byte, []int) {
	return file_sso_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshTokenRes) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

var File_sso_auth_service_proto protoreflect.FileDescriptor

var file_sso_auth_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x73, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x14,
	0x73, 0x73, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x73, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x13, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4a, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79,
	0x43, 0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x22, 0x73, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x43, 0x72, 0x65, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x41,
	0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x2e, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x2a, 0x5f, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x22, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e,
	0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x43, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x10, 0x01, 0x32, 0xd6, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x12, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x43, 0x72,
	0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x79, 0x43, 0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a,
	0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x63,
	0x69, 0x6e, 0x65, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sso_auth_service_proto_rawDescOnce sync.Once
	file_sso_auth_service_proto_rawDescData = file_sso_auth_service_proto_rawDesc
)

func file_sso_auth_service_proto_rawDescGZIP() []byte {
	file_sso_auth_service_proto_rawDescOnce.Do(func() {
		file_sso_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_auth_service_proto_rawDescData)
	})
	return file_sso_auth_service_proto_rawDescData
}

var file_sso_auth_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sso_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sso_auth_service_proto_goTypes = []any{
	(UserContactTypeEnum)(0),     // 0: auth.UserContactTypeEnum
	(*RegisterUserRequest)(nil),  // 1: auth.RegisterUserRequest
	(*RegisterUserResponse)(nil), // 2: auth.RegisterUserResponse
	(*LoginByCredsReq)(nil),      // 3: auth.LoginByCredsReq
	(*LoginByCredsRes)(nil),      // 4: auth.LoginByCredsRes
	(*RefreshTokenReq)(nil),      // 5: auth.RefreshTokenReq
	(*RefreshTokenRes)(nil),      // 6: auth.RefreshTokenRes
	(*UserCredentials)(nil),      // 7: auth.UserCredentials
	(*user.User)(nil),            // 8: user.User
	(*Token)(nil),                // 9: auth.Token
}
var file_sso_auth_service_proto_depIdxs = []int32{
	7,  // 0: auth.RegisterUserRequest.credentials:type_name -> auth.UserCredentials
	8,  // 1: auth.RegisterUserResponse.user:type_name -> user.User
	9,  // 2: auth.RegisterUserResponse.access_token:type_name -> auth.Token
	9,  // 3: auth.RegisterUserResponse.refresh_token:type_name -> auth.Token
	7,  // 4: auth.LoginByCredsReq.credentials:type_name -> auth.UserCredentials
	9,  // 5: auth.LoginByCredsRes.access_token:type_name -> auth.Token
	9,  // 6: auth.LoginByCredsRes.refresh_token:type_name -> auth.Token
	9,  // 7: auth.RefreshTokenRes.access_token:type_name -> auth.Token
	1,  // 8: auth.AuthService.RegisterUser:input_type -> auth.RegisterUserRequest
	3,  // 9: auth.AuthService.LoginByCredentials:input_type -> auth.LoginByCredsReq
	5,  // 10: auth.AuthService.RefreshToken:input_type -> auth.RefreshTokenReq
	2,  // 11: auth.AuthService.RegisterUser:output_type -> auth.RegisterUserResponse
	4,  // 12: auth.AuthService.LoginByCredentials:output_type -> auth.LoginByCredsRes
	6,  // 13: auth.AuthService.RefreshToken:output_type -> auth.RefreshTokenRes
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_sso_auth_service_proto_init() }
func file_sso_auth_service_proto_init() {
	if File_sso_auth_service_proto != nil {
		return
	}
	file_sso_auth_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sso_auth_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserRequest); i {
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
		file_sso_auth_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserResponse); i {
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
		file_sso_auth_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LoginByCredsReq); i {
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
		file_sso_auth_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LoginByCredsRes); i {
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
		file_sso_auth_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshTokenReq); i {
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
		file_sso_auth_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshTokenRes); i {
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
			RawDescriptor: file_sso_auth_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_auth_service_proto_goTypes,
		DependencyIndexes: file_sso_auth_service_proto_depIdxs,
		EnumInfos:         file_sso_auth_service_proto_enumTypes,
		MessageInfos:      file_sso_auth_service_proto_msgTypes,
	}.Build()
	File_sso_auth_service_proto = out.File
	file_sso_auth_service_proto_rawDesc = nil
	file_sso_auth_service_proto_goTypes = nil
	file_sso_auth_service_proto_depIdxs = nil
}
