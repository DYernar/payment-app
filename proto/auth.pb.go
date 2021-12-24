// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AuthResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Id                   int64                `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Iin                  string               `protobuf:"bytes,2,opt,name=Iin,proto3" json:"Iin,omitempty"`
	Login                string               `protobuf:"bytes,3,opt,name=Login,proto3" json:"Login,omitempty"`
	Role                 *Role                `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetIin() string {
	if m != nil {
		return m.Iin
	}
	return ""
}

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type Role struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*AuthResponse)(nil), "auth.AuthResponse")
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Role)(nil), "auth.Role")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x9b, 0x5a, 0x32, 0x15, 0x91, 0xc5, 0xc3, 0x52, 0x50, 0x4b, 0x4e, 0x45, 0x70,
	0x03, 0xf5, 0x22, 0xe8, 0xa5, 0x7a, 0x31, 0xa0, 0x1e, 0xd6, 0xea, 0xc1, 0x5b, 0x6a, 0xa6, 0x69,
	0xb0, 0xc9, 0x86, 0xdd, 0x8d, 0xe0, 0x3f, 0xf1, 0xe7, 0xca, 0xce, 0x26, 0x58, 0xf0, 0x36, 0xef,
	0xf1, 0x76, 0xe7, 0x7b, 0x03, 0x90, 0xb5, 0x76, 0x2b, 0x1a, 0xad, 0xac, 0x62, 0xa1, 0x9b, 0xa7,
	0xe7, 0x85, 0x52, 0xc5, 0x0e, 0x13, 0xf2, 0xd6, 0xed, 0x26, 0xb1, 0x65, 0x85, 0xc6, 0x66, 0x55,
	0xe3, 0x63, 0xf1, 0x29, 0x8c, 0x56, 0xea, 0x13, 0x6b, 0x76, 0xd2, 0x0d, 0x3c, 0x98, 0x05, 0xf3,
	0x48, 0x7a, 0x11, 0x3f, 0xc0, 0xe1, 0xb2, 0xb5, 0x5b, 0x89, 0xa6, 0x51, 0xb5, 0x41, 0xc6, 0x61,
	0xfc, 0x84, 0xc6, 0x64, 0x05, 0x76, 0xb9, 0x5e, 0xb2, 0x33, 0x08, 0x5b, 0x83, 0x9a, 0x0f, 0x66,
	0xc1, 0x7c, 0xb2, 0x00, 0x41, 0x28, 0xaf, 0x06, 0xb5, 0x24, 0x3f, 0xfe, 0x09, 0x20, 0x74, 0x92,
	0x1d, 0xc1, 0x20, 0xcd, 0xe9, 0xf5, 0x50, 0x0e, 0xd2, 0x9c, 0x1d, 0xc3, 0x30, 0x2d, 0x6b, 0x7a,
	0x17, 0x49, 0x37, 0x3a, 0x94, 0x47, 0x55, 0x94, 0x35, 0x1f, 0x7a, 0x14, 0x12, 0x6e, 0x81, 0x54,
	0x3b, 0xe4, 0xe1, 0xfe, 0x02, 0xe7, 0x48, 0xf2, 0xd9, 0x35, 0x44, 0xf7, 0x1a, 0x33, 0x8b, 0xf9,
	0xd2, 0xf2, 0x11, 0x85, 0xa6, 0xc2, 0xd7, 0x17, 0x7d, 0x7d, 0xb1, 0xea, 0xeb, 0xcb, 0xbf, 0x70,
	0x7c, 0xe1, 0x7f, 0xfe, 0x47, 0xc6, 0x20, 0x7c, 0xce, 0x2a, 0xec, 0xd0, 0x68, 0x5e, 0xdc, 0xc2,
	0xc4, 0x1d, 0xe4, 0x05, 0xf5, 0x57, 0xf9, 0x81, 0xec, 0x12, 0xe0, 0x0d, 0x75, 0xb9, 0xf9, 0xa6,
	0x6a, 0x13, 0x0f, 0x45, 0xa7, 0x9b, 0x32, 0x2f, 0xf6, 0xcf, 0x77, 0x17, 0xbd, 0x8f, 0x45, 0x72,
	0xe3, 0xfc, 0xf5, 0x01, 0x31, 0x5d, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x02, 0xb1, 0x5f,
	0xb4, 0x01, 0x00, 0x00,
}
