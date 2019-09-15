// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type Account struct {
	AccountId            string               `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash         string               `protobuf:"bytes,4,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	PasswordSalt         string               `protobuf:"bytes,5,opt,name=password_salt,json=passwordSalt,proto3" json:"password_salt,omitempty"`
	Description          string               `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *Account) GetPasswordSalt() string {
	if m != nil {
		return m.PasswordSalt
	}
	return ""
}

func (m *Account) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Account) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Account) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "account.Account")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0xef, 0xd2, 0x30,
	0x18, 0xc6, 0x33, 0xfe, 0x7f, 0x60, 0x14, 0x34, 0xa6, 0xf1, 0x30, 0x38, 0x28, 0xd1, 0x83, 0xc3,
	0x84, 0x15, 0x30, 0x92, 0x88, 0x27, 0x38, 0xe9, 0x75, 0x7a, 0xf2, 0x42, 0xba, 0xae, 0x74, 0x35,
	0x5b, 0xdf, 0xa5, 0xed, 0xd4, 0x8f, 0xe1, 0xd7, 0x81, 0x93, 0xdf, 0xc5, 0x2f, 0x62, 0xd6, 0x6d,
	0x0a, 0x89, 0xd1, 0x4b, 0xd3, 0xb7, 0xfd, 0x3d, 0x4f, 0xdf, 0x3c, 0x6f, 0xd1, 0xb8, 0x80, 0x94,
	0xe7, 0x51, 0xa9, 0xc1, 0x02, 0x1e, 0x52, 0xc6, 0xa0, 0x52, 0x76, 0xb6, 0x13, 0xd2, 0x66, 0x55,
	0x12, 0x31, 0x28, 0x88, 0x54, 0x27, 0x48, 0x72, 0xf8, 0x06, 0x25, 0x57, 0xc4, 0x71, 0x6c, 0x29,
	0xb8, 0x5a, 0x0a, 0xd0, 0x05, 0x81, 0xd2, 0x4a, 0x50, 0x86, 0xd4, 0x45, 0x63, 0x32, 0x7b, 0x2a,
	0x00, 0x44, 0xce, 0x1b, 0x34, 0xa9, 0x4e, 0xc4, 0xca, 0x82, 0x1b, 0x4b, 0x8b, 0xb2, 0x05, 0x96,
	0x57, 0xe6, 0x02, 0x04, 0xfc, 0x21, 0xeb, 0xca, 0x15, 0x6e, 0xd7, 0xe0, 0xcf, 0xbe, 0xdf, 0xa1,
	0xe1, 0xbe, 0xe9, 0x0b, 0xbf, 0x40, 0xa8, 0x6d, 0xf1, 0x28, 0xd3, 0xc0, 0x9b, 0x7b, 0xe1, 0xe8,
	0xe0, 0x5f, 0xce, 0xd3, 0x7b, 0xd4, 0x0b, 0xbd, 0x78, 0xd4, 0xde, 0xbd, 0x4f, 0xf1, 0x0a, 0xf9,
	0x95, 0xe1, 0x5a, 0xd1, 0x82, 0x07, 0x3d, 0x87, 0x3d, 0xbe, 0x9c, 0xa7, 0x8f, 0xd0, 0x43, 0x3c,
	0xf9, 0x42, 0x35, 0xcb, 0xa8, 0x0e, 0xd7, 0x9b, 0xd5, 0x22, 0xfe, 0x4d, 0xe1, 0x97, 0xa8, 0xcf,
	0x0b, 0x2a, 0xf3, 0xe0, 0xee, 0x1f, 0x78, 0x83, 0xe0, 0xe7, 0xe8, 0x41, 0x49, 0x8d, 0xf9, 0x0a,
	0x3a, 0x3d, 0x66, 0xd4, 0x64, 0xc1, 0x7d, 0xad, 0x89, 0x27, 0xdd, 0xe1, 0x3b, 0x6a, 0xb2, 0x1b,
	0xc8, 0xd0, 0xdc, 0x06, 0xfd, 0x5b, 0xe8, 0x03, 0xcd, 0x2d, 0xde, 0xa2, 0x71, 0xca, 0x0d, 0xd3,
	0xd2, 0xe5, 0x18, 0x0c, 0xfe, 0xfe, 0xf6, 0xeb, 0xf5, 0x66, 0x11, 0x5f, 0x83, 0xf8, 0x0d, 0x42,
	0x4c, 0x73, 0x6a, 0x79, 0x7a, 0xa4, 0x36, 0x18, 0xce, 0xbd, 0x70, 0xbc, 0x99, 0x45, 0x4d, 0xf2,
	0x51, 0x97, 0x67, 0xf4, 0xb1, 0x4b, 0x3e, 0x1e, 0xb5, 0xf4, 0xde, 0xd6, 0xd2, 0xaa, 0x4c, 0x3b,
	0xa9, 0xff, 0x7f, 0x69, 0x4b, 0xef, 0xed, 0x6e, 0x70, 0x39, 0x4f, 0x7b, 0xbe, 0x77, 0xd8, 0xfe,
	0xf8, 0xf9, 0xc4, 0xfb, 0xb4, 0xba, 0x9a, 0xa3, 0x92, 0x9a, 0x7e, 0x16, 0x1c, 0xb4, 0x90, 0x54,
	0x91, 0x76, 0x0e, 0xc4, 0x68, 0x46, 0xdc, 0xff, 0x7a, 0xeb, 0xd6, 0x64, 0xe0, 0xec, 0x5f, 0xfd,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xe2, 0x09, 0x8a, 0x75, 0x02, 0x00, 0x00,
}
