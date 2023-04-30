// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: services/auth/proto/auth.proto

package auth

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

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier    string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	OldCredential string `protobuf:"bytes,2,opt,name=oldCredential,proto3" json:"oldCredential,omitempty"`
	NewCredential string `protobuf:"bytes,3,opt,name=newCredential,proto3" json:"newCredential,omitempty"`
	Token         string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ChangePasswordRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *ChangePasswordRequest) GetOldCredential() string {
	if x != nil {
		return x.OldCredential
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewCredential() string {
	if x != nil {
		return x.NewCredential
	}
	return ""
}

func (x *ChangePasswordRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignupRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *SignupRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *LoginRequest) GetCredential() string {
	if x != nil {
		return x.Credential
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *RefreshTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type IdentifierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *IdentifierResponse) Reset() {
	*x = IdentifierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentifierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentifierResponse) ProtoMessage() {}

func (x *IdentifierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentifierResponse.ProtoReflect.Descriptor instead.
func (*IdentifierResponse) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *IdentifierResponse) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_auth_proto_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_auth_proto_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_services_auth_proto_auth_proto_rawDescGZIP(), []int{7}
}

func (x *TokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_services_auth_proto_auth_proto protoreflect.FileDescriptor

var file_services_auth_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a,
	0x69, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22, 0x99, 0x01, 0x0a,
	0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x6c, 0x64, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x6c, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d,
	0x6e, 0x65, 0x77, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4f, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x4e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x45, 0x0a, 0x0d, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x4b, 0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a,
	0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x34, 0x0a, 0x12, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x22, 0x25, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x84, 0x05, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x75, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x31, 0x2e, 0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f,
	0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e,
	0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x28, 0x2e, 0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61,
	0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x65,
	0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x12, 0x29, 0x2e, 0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b,
	0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x0c, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x2e, 0x61, 0x65, 0x72,
	0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x65,
	0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x12, 0x29, 0x2e, 0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b,
	0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x0d, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x2e, 0x61, 0x65,
	0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x61, 0x65, 0x72, 0x69, 0x71, 0x75, 0x2e, 0x77, 0x6b, 0x5f, 0x6b, 0x61, 0x6e, 0x6a, 0x69, 0x5f,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x65,
	0x72, 0x69, 0x71, 0x75, 0x2f, 0x6b, 0x61, 0x6e, 0x69, 0x6b, 0x61, 0x6b, 0x69, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_services_auth_proto_auth_proto_rawDescOnce sync.Once
	file_services_auth_proto_auth_proto_rawDescData = file_services_auth_proto_auth_proto_rawDesc
)

func file_services_auth_proto_auth_proto_rawDescGZIP() []byte {
	file_services_auth_proto_auth_proto_rawDescOnce.Do(func() {
		file_services_auth_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_auth_proto_auth_proto_rawDescData)
	})
	return file_services_auth_proto_auth_proto_rawDescData
}

var file_services_auth_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_auth_proto_auth_proto_goTypes = []interface{}{
	(*ChangePasswordRequest)(nil), // 0: aeriqu.wk_kanji_write.auth.ChangePasswordRequest
	(*SignupRequest)(nil),         // 1: aeriqu.wk_kanji_write.auth.SignupRequest
	(*LoginRequest)(nil),          // 2: aeriqu.wk_kanji_write.auth.LoginRequest
	(*LogoutRequest)(nil),         // 3: aeriqu.wk_kanji_write.auth.LogoutRequest
	(*RefreshTokenRequest)(nil),   // 4: aeriqu.wk_kanji_write.auth.RefreshTokenRequest
	(*ValidateTokenRequest)(nil),  // 5: aeriqu.wk_kanji_write.auth.ValidateTokenRequest
	(*IdentifierResponse)(nil),    // 6: aeriqu.wk_kanji_write.auth.IdentifierResponse
	(*TokenResponse)(nil),         // 7: aeriqu.wk_kanji_write.auth.TokenResponse
}
var file_services_auth_proto_auth_proto_depIdxs = []int32{
	0, // 0: aeriqu.wk_kanji_write.auth.Auth.ChangePassword:input_type -> aeriqu.wk_kanji_write.auth.ChangePasswordRequest
	2, // 1: aeriqu.wk_kanji_write.auth.Auth.Login:input_type -> aeriqu.wk_kanji_write.auth.LoginRequest
	3, // 2: aeriqu.wk_kanji_write.auth.Auth.Logout:input_type -> aeriqu.wk_kanji_write.auth.LogoutRequest
	4, // 3: aeriqu.wk_kanji_write.auth.Auth.RefreshToken:input_type -> aeriqu.wk_kanji_write.auth.RefreshTokenRequest
	1, // 4: aeriqu.wk_kanji_write.auth.Auth.Signup:input_type -> aeriqu.wk_kanji_write.auth.SignupRequest
	5, // 5: aeriqu.wk_kanji_write.auth.Auth.ValidateToken:input_type -> aeriqu.wk_kanji_write.auth.ValidateTokenRequest
	6, // 6: aeriqu.wk_kanji_write.auth.Auth.ChangePassword:output_type -> aeriqu.wk_kanji_write.auth.IdentifierResponse
	7, // 7: aeriqu.wk_kanji_write.auth.Auth.Login:output_type -> aeriqu.wk_kanji_write.auth.TokenResponse
	7, // 8: aeriqu.wk_kanji_write.auth.Auth.Logout:output_type -> aeriqu.wk_kanji_write.auth.TokenResponse
	7, // 9: aeriqu.wk_kanji_write.auth.Auth.RefreshToken:output_type -> aeriqu.wk_kanji_write.auth.TokenResponse
	7, // 10: aeriqu.wk_kanji_write.auth.Auth.Signup:output_type -> aeriqu.wk_kanji_write.auth.TokenResponse
	6, // 11: aeriqu.wk_kanji_write.auth.Auth.ValidateToken:output_type -> aeriqu.wk_kanji_write.auth.IdentifierResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_auth_proto_auth_proto_init() }
func file_services_auth_proto_auth_proto_init() {
	if File_services_auth_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_auth_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_services_auth_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
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
		file_services_auth_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_services_auth_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_services_auth_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenRequest); i {
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
		file_services_auth_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
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
		file_services_auth_proto_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentifierResponse); i {
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
		file_services_auth_proto_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResponse); i {
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
			RawDescriptor: file_services_auth_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_auth_proto_auth_proto_goTypes,
		DependencyIndexes: file_services_auth_proto_auth_proto_depIdxs,
		MessageInfos:      file_services_auth_proto_auth_proto_msgTypes,
	}.Build()
	File_services_auth_proto_auth_proto = out.File
	file_services_auth_proto_auth_proto_rawDesc = nil
	file_services_auth_proto_auth_proto_goTypes = nil
	file_services_auth_proto_auth_proto_depIdxs = nil
}