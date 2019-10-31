// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/model.proto

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
	VerificationCode     string               `protobuf:"bytes,7,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`
	Verified             bool                 `protobuf:"varint,8,opt,name=verified,proto3" json:"verified,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8c387d18afeb6b2, []int{0}
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

func (m *Account) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

func (m *Account) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
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

func init() { proto.RegisterFile("account/model.proto", fileDescriptor_d8c387d18afeb6b2) }

var fileDescriptor_d8c387d18afeb6b2 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd2, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x07, 0x70, 0x39, 0xa4, 0x89, 0xbd, 0x09, 0xa5, 0x2c, 0x1c, 0xec, 0x1c, 0x20, 0x82, 0x03,
	0x0e, 0x28, 0x76, 0x1a, 0x44, 0x45, 0xcb, 0x01, 0x25, 0x5c, 0xe0, 0x6a, 0x38, 0x71, 0x89, 0xd6,
	0xbb, 0x1b, 0x7b, 0x91, 0xed, 0xb1, 0x76, 0xd7, 0x85, 0x07, 0xe1, 0x65, 0xec, 0x13, 0xef, 0xc2,
	0x8b, 0x20, 0x7f, 0xb5, 0x3e, 0xa0, 0xf6, 0x62, 0x79, 0xd6, 0xbf, 0xff, 0x68, 0xe5, 0x19, 0xf4,
	0x84, 0x50, 0x0a, 0x45, 0xa6, 0xfd, 0x14, 0x18, 0x4f, 0xbc, 0x5c, 0x82, 0x06, 0x3c, 0xed, 0x0e,
	0x17, 0x57, 0x91, 0xd0, 0x71, 0x11, 0x7a, 0x14, 0x52, 0x5f, 0x64, 0x47, 0x08, 0x13, 0xf8, 0x05,
	0x39, 0xcf, 0xfc, 0xc6, 0xd1, 0x75, 0xc4, 0xb3, 0x75, 0x04, 0x32, 0xf5, 0x21, 0xd7, 0x02, 0x32,
	0xe5, 0xd7, 0x45, 0xdb, 0x64, 0xf1, 0x3c, 0x02, 0x88, 0x12, 0xde, 0xd2, 0xb0, 0x38, 0xfa, 0x5a,
	0xa4, 0x5c, 0x69, 0x92, 0xe6, 0x1d, 0x58, 0x0f, 0x9a, 0x47, 0x10, 0xc1, 0xad, 0xac, 0xab, 0xa6,
	0x68, 0xde, 0x5a, 0xfe, 0xe2, 0xf7, 0x18, 0x4d, 0x77, 0xed, 0xbd, 0xf0, 0x2b, 0x84, 0xba, 0x2b,
	0x1e, 0x04, 0xb3, 0x8d, 0xa5, 0xe1, 0x5a, 0x7b, 0xb3, 0x2a, 0x9d, 0x31, 0x1a, 0xb9, 0x46, 0x60,
	0x75, 0xdf, 0xbe, 0x30, 0xbc, 0x41, 0x66, 0xa1, 0xb8, 0xcc, 0x48, 0xca, 0xed, 0x51, 0xc3, 0x9e,
	0x56, 0xa5, 0x73, 0x86, 0x4e, 0xf1, 0xfc, 0x9a, 0x48, 0x1a, 0x13, 0xe9, 0x9e, 0x6f, 0x37, 0xab,
	0xe0, 0x46, 0xe1, 0xd7, 0xe8, 0x84, 0xa7, 0x44, 0x24, 0xf6, 0x83, 0x3b, 0x78, 0x4b, 0xf0, 0x25,
	0x7a, 0x98, 0x13, 0xa5, 0x7e, 0x82, 0x64, 0x87, 0x98, 0xa8, 0xd8, 0x1e, 0xdf, 0x91, 0x99, 0xf7,
	0xf4, 0x33, 0x51, 0x31, 0x7e, 0x39, 0x88, 0x2a, 0x92, 0x68, 0xfb, 0xa4, 0x8e, 0xde, 0xa2, 0xaf,
	0x24, 0xd1, 0xf8, 0x02, 0xcd, 0x18, 0x57, 0x54, 0x8a, 0xe6, 0xef, 0xda, 0x93, 0xff, 0x77, 0x7f,
	0x77, 0xbe, 0x5d, 0x05, 0x43, 0x88, 0x3f, 0xa2, 0xc7, 0xd7, 0x5c, 0x8a, 0xa3, 0xa0, 0xa4, 0xae,
	0x0f, 0x14, 0x18, 0xb7, 0xa7, 0x4d, 0x1a, 0x57, 0xa5, 0x73, 0x8a, 0xe6, 0x18, 0xf5, 0xe9, 0xf7,
	0xab, 0xe0, 0x6c, 0x88, 0x3f, 0x01, 0xe3, 0xf8, 0x0d, 0x32, 0xdb, 0x33, 0xce, 0x6c, 0x73, 0x69,
	0xb8, 0xe6, 0xfe, 0x51, 0x55, 0x3a, 0x33, 0x64, 0xe1, 0x69, 0x08, 0x90, 0x70, 0x92, 0x05, 0x37,
	0x00, 0x5f, 0x22, 0x44, 0x25, 0x27, 0x9a, 0xb3, 0x03, 0xd1, 0xb6, 0xb5, 0x34, 0xdc, 0xd9, 0x76,
	0xe1, 0xb5, 0xd3, 0xf7, 0xfa, 0x99, 0x7a, 0xdf, 0xfa, 0xe9, 0x07, 0x56, 0xa7, 0x77, 0xba, 0x8e,
	0x16, 0x39, 0xeb, 0xa3, 0xe8, 0xfe, 0x68, 0xa7, 0x77, 0xfa, 0x6a, 0x52, 0x95, 0xce, 0xc8, 0x34,
	0xf6, 0x17, 0x7f, 0xfe, 0x3e, 0x33, 0xbe, 0x6f, 0x06, 0xbb, 0x94, 0x09, 0x49, 0x7e, 0x44, 0x1c,
	0x64, 0x24, 0x48, 0xe6, 0xf7, 0xfb, 0xad, 0x24, 0x6d, 0x77, 0xfc, 0x43, 0xf3, 0x0c, 0x27, 0x4d,
	0xfb, 0xb7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x35, 0xa3, 0x10, 0x01, 0x01, 0x03, 0x00, 0x00,
}
