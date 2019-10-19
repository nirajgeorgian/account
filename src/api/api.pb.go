// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	model "github.com/nirajgeorgian/account/src/model"
	grpc "google.golang.org/grpc"
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

// create an account
type CreateAccountReq struct {
	Account              *model.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateAccountReq) Reset()         { *m = CreateAccountReq{} }
func (m *CreateAccountReq) String() string { return proto.CompactTextString(m) }
func (*CreateAccountReq) ProtoMessage()    {}
func (*CreateAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *CreateAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountReq.Unmarshal(m, b)
}
func (m *CreateAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountReq.Marshal(b, m, deterministic)
}
func (m *CreateAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountReq.Merge(m, src)
}
func (m *CreateAccountReq) XXX_Size() int {
	return xxx_messageInfo_CreateAccountReq.Size(m)
}
func (m *CreateAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountReq proto.InternalMessageInfo

func (m *CreateAccountReq) GetAccount() *model.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type CreateAccountRes struct {
	Account              *model.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateAccountRes) Reset()         { *m = CreateAccountRes{} }
func (m *CreateAccountRes) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRes) ProtoMessage()    {}
func (*CreateAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *CreateAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRes.Unmarshal(m, b)
}
func (m *CreateAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRes.Marshal(b, m, deterministic)
}
func (m *CreateAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRes.Merge(m, src)
}
func (m *CreateAccountRes) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRes.Size(m)
}
func (m *CreateAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRes proto.InternalMessageInfo

func (m *CreateAccountRes) GetAccount() *model.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

// read an account
type ReadAccountReq struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAccountReq) Reset()         { *m = ReadAccountReq{} }
func (m *ReadAccountReq) String() string { return proto.CompactTextString(m) }
func (*ReadAccountReq) ProtoMessage()    {}
func (*ReadAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *ReadAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAccountReq.Unmarshal(m, b)
}
func (m *ReadAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAccountReq.Marshal(b, m, deterministic)
}
func (m *ReadAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAccountReq.Merge(m, src)
}
func (m *ReadAccountReq) XXX_Size() int {
	return xxx_messageInfo_ReadAccountReq.Size(m)
}
func (m *ReadAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAccountReq proto.InternalMessageInfo

func (m *ReadAccountReq) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type ReadAccountRes struct {
	Account              *model.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReadAccountRes) Reset()         { *m = ReadAccountRes{} }
func (m *ReadAccountRes) String() string { return proto.CompactTextString(m) }
func (*ReadAccountRes) ProtoMessage()    {}
func (*ReadAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *ReadAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAccountRes.Unmarshal(m, b)
}
func (m *ReadAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAccountRes.Marshal(b, m, deterministic)
}
func (m *ReadAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAccountRes.Merge(m, src)
}
func (m *ReadAccountRes) XXX_Size() int {
	return xxx_messageInfo_ReadAccountRes.Size(m)
}
func (m *ReadAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAccountRes proto.InternalMessageInfo

func (m *ReadAccountRes) GetAccount() *model.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

// update an account
type UpdateAccountReq struct {
	Account              *model.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateAccountReq) Reset()         { *m = UpdateAccountReq{} }
func (m *UpdateAccountReq) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountReq) ProtoMessage()    {}
func (*UpdateAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *UpdateAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountReq.Unmarshal(m, b)
}
func (m *UpdateAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountReq.Marshal(b, m, deterministic)
}
func (m *UpdateAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountReq.Merge(m, src)
}
func (m *UpdateAccountReq) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountReq.Size(m)
}
func (m *UpdateAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountReq proto.InternalMessageInfo

func (m *UpdateAccountReq) GetAccount() *model.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type UpdateAccountRes struct {
	Success              bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Account              *model.Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateAccountRes) Reset()         { *m = UpdateAccountRes{} }
func (m *UpdateAccountRes) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountRes) ProtoMessage()    {}
func (*UpdateAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *UpdateAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountRes.Unmarshal(m, b)
}
func (m *UpdateAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountRes.Marshal(b, m, deterministic)
}
func (m *UpdateAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountRes.Merge(m, src)
}
func (m *UpdateAccountRes) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountRes.Size(m)
}
func (m *UpdateAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountRes proto.InternalMessageInfo

func (m *UpdateAccountRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateAccountRes) GetAccount() *model.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

// delete an account
type DeleteAccountReq struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountReq) Reset()         { *m = DeleteAccountReq{} }
func (m *DeleteAccountReq) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountReq) ProtoMessage()    {}
func (*DeleteAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *DeleteAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountReq.Unmarshal(m, b)
}
func (m *DeleteAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountReq.Marshal(b, m, deterministic)
}
func (m *DeleteAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountReq.Merge(m, src)
}
func (m *DeleteAccountReq) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountReq.Size(m)
}
func (m *DeleteAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountReq proto.InternalMessageInfo

func (m *DeleteAccountReq) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type DeleteAccountRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAccountRes) Reset()         { *m = DeleteAccountRes{} }
func (m *DeleteAccountRes) String() string { return proto.CompactTextString(m) }
func (*DeleteAccountRes) ProtoMessage()    {}
func (*DeleteAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *DeleteAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAccountRes.Unmarshal(m, b)
}
func (m *DeleteAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAccountRes.Marshal(b, m, deterministic)
}
func (m *DeleteAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAccountRes.Merge(m, src)
}
func (m *DeleteAccountRes) XXX_Size() int {
	return xxx_messageInfo_DeleteAccountRes.Size(m)
}
func (m *DeleteAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAccountRes proto.InternalMessageInfo

func (m *DeleteAccountRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListAccountsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAccountsReq) Reset()         { *m = ListAccountsReq{} }
func (m *ListAccountsReq) String() string { return proto.CompactTextString(m) }
func (*ListAccountsReq) ProtoMessage()    {}
func (*ListAccountsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *ListAccountsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccountsReq.Unmarshal(m, b)
}
func (m *ListAccountsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccountsReq.Marshal(b, m, deterministic)
}
func (m *ListAccountsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsReq.Merge(m, src)
}
func (m *ListAccountsReq) XXX_Size() int {
	return xxx_messageInfo_ListAccountsReq.Size(m)
}
func (m *ListAccountsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsReq proto.InternalMessageInfo

type ListAccountsRes struct {
	Account              []*model.Account `protobuf:"bytes,1,rep,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListAccountsRes) Reset()         { *m = ListAccountsRes{} }
func (m *ListAccountsRes) String() string { return proto.CompactTextString(m) }
func (*ListAccountsRes) ProtoMessage()    {}
func (*ListAccountsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *ListAccountsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAccountsRes.Unmarshal(m, b)
}
func (m *ListAccountsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAccountsRes.Marshal(b, m, deterministic)
}
func (m *ListAccountsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAccountsRes.Merge(m, src)
}
func (m *ListAccountsRes) XXX_Size() int {
	return xxx_messageInfo_ListAccountsRes.Size(m)
}
func (m *ListAccountsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAccountsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListAccountsRes proto.InternalMessageInfo

func (m *ListAccountsRes) GetAccount() []*model.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AuthReq struct {
	Account              *model.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AuthReq) Reset()         { *m = AuthReq{} }
func (m *AuthReq) String() string { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()    {}
func (*AuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *AuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReq.Unmarshal(m, b)
}
func (m *AuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReq.Marshal(b, m, deterministic)
}
func (m *AuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReq.Merge(m, src)
}
func (m *AuthReq) XXX_Size() int {
	return xxx_messageInfo_AuthReq.Size(m)
}
func (m *AuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReq proto.InternalMessageInfo

func (m *AuthReq) GetAccount() *model.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AuthRes struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRes) Reset()         { *m = AuthRes{} }
func (m *AuthRes) String() string { return proto.CompactTextString(m) }
func (*AuthRes) ProtoMessage()    {}
func (*AuthRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *AuthRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRes.Unmarshal(m, b)
}
func (m *AuthRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRes.Marshal(b, m, deterministic)
}
func (m *AuthRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRes.Merge(m, src)
}
func (m *AuthRes) XXX_Size() int {
	return xxx_messageInfo_AuthRes.Size(m)
}
func (m *AuthRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRes.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRes proto.InternalMessageInfo

func (m *AuthRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthRes) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type ValidateTokenReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenReq) Reset()         { *m = ValidateTokenReq{} }
func (m *ValidateTokenReq) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenReq) ProtoMessage()    {}
func (*ValidateTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *ValidateTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenReq.Unmarshal(m, b)
}
func (m *ValidateTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenReq.Marshal(b, m, deterministic)
}
func (m *ValidateTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenReq.Merge(m, src)
}
func (m *ValidateTokenReq) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenReq.Size(m)
}
func (m *ValidateTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenReq proto.InternalMessageInfo

func (m *ValidateTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateTokenRes struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRes) Reset()         { *m = ValidateTokenRes{} }
func (m *ValidateTokenRes) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRes) ProtoMessage()    {}
func (*ValidateTokenRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{13}
}

func (m *ValidateTokenRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRes.Unmarshal(m, b)
}
func (m *ValidateTokenRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRes.Marshal(b, m, deterministic)
}
func (m *ValidateTokenRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRes.Merge(m, src)
}
func (m *ValidateTokenRes) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenRes.Size(m)
}
func (m *ValidateTokenRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenRes.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenRes proto.InternalMessageInfo

func (m *ValidateTokenRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ValidateTokenRes) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*CreateAccountReq)(nil), "account.CreateAccountReq")
	proto.RegisterType((*CreateAccountRes)(nil), "account.CreateAccountRes")
	proto.RegisterType((*ReadAccountReq)(nil), "account.ReadAccountReq")
	proto.RegisterType((*ReadAccountRes)(nil), "account.ReadAccountRes")
	proto.RegisterType((*UpdateAccountReq)(nil), "account.UpdateAccountReq")
	proto.RegisterType((*UpdateAccountRes)(nil), "account.UpdateAccountRes")
	proto.RegisterType((*DeleteAccountReq)(nil), "account.DeleteAccountReq")
	proto.RegisterType((*DeleteAccountRes)(nil), "account.DeleteAccountRes")
	proto.RegisterType((*ListAccountsReq)(nil), "account.ListAccountsReq")
	proto.RegisterType((*ListAccountsRes)(nil), "account.ListAccountsRes")
	proto.RegisterType((*AuthReq)(nil), "account.AuthReq")
	proto.RegisterType((*AuthRes)(nil), "account.AuthRes")
	proto.RegisterType((*ValidateTokenReq)(nil), "account.ValidateTokenReq")
	proto.RegisterType((*ValidateTokenRes)(nil), "account.ValidateTokenRes")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x4e, 0xe3, 0x30,
	0x10, 0x6d, 0xbb, 0x97, 0x6e, 0xa7, 0x6a, 0x37, 0x6b, 0xad, 0xb4, 0xdd, 0x48, 0xbb, 0x42, 0x79,
	0xaa, 0x50, 0x48, 0x44, 0x11, 0x4f, 0x5c, 0xa4, 0x02, 0x52, 0x85, 0xc4, 0x53, 0xb8, 0x3c, 0xf0,
	0x82, 0xdc, 0xc4, 0xa4, 0x86, 0x36, 0x4e, 0x63, 0xa7, 0x3f, 0xc2, 0x0f, 0xf1, 0x2f, 0xfc, 0x08,
	0xb2, 0x93, 0x54, 0x49, 0x7a, 0x41, 0xe1, 0x2d, 0x67, 0xe6, 0x9c, 0x19, 0x1f, 0xcf, 0xc4, 0xd0,
	0xc2, 0x21, 0xb5, 0xc2, 0x88, 0x09, 0x86, 0x9a, 0xd8, 0x75, 0x59, 0x1c, 0x08, 0x7d, 0xcf, 0xa7,
	0x62, 0x12, 0x8f, 0x2d, 0x97, 0xcd, 0x6c, 0x9f, 0xf9, 0xcc, 0x56, 0xf9, 0x71, 0xfc, 0xa8, 0x90,
	0x02, 0xea, 0x2b, 0xd1, 0xe9, 0xed, 0x19, 0xf3, 0xc8, 0x34, 0x01, 0xc6, 0x29, 0x68, 0xe7, 0x11,
	0xc1, 0x82, 0x0c, 0x93, 0x62, 0x0e, 0x99, 0xa3, 0x5d, 0xc8, 0x4a, 0xf7, 0xea, 0x3b, 0xf5, 0x7e,
	0x7b, 0xa0, 0x59, 0x29, 0xb6, 0x32, 0x56, 0x46, 0x58, 0xa3, 0xe7, 0x95, 0xf4, 0x36, 0x74, 0x1d,
	0x82, 0xbd, 0x5c, 0xf7, 0x7f, 0x00, 0x69, 0xf2, 0x81, 0x7a, 0xaa, 0x40, 0xcb, 0x69, 0xa5, 0x91,
	0x4b, 0xcf, 0x38, 0x2e, 0x09, 0x78, 0xd5, 0xe3, 0xde, 0x86, 0xde, 0xe7, 0xed, 0x9a, 0x2b, 0x7a,
	0x8e, 0x7a, 0xd0, 0xe4, 0xb1, 0xeb, 0x12, 0xce, 0x95, 0xfe, 0x87, 0x93, 0x41, 0x63, 0x1f, 0xb4,
	0x0b, 0x32, 0x25, 0x85, 0x6e, 0x1f, 0xd8, 0x33, 0x57, 0x24, 0xdb, 0x1a, 0xfc, 0x82, 0x9f, 0x57,
	0x94, 0x8b, 0x94, 0xcb, 0x1d, 0x32, 0x37, 0x4e, 0xca, 0xa1, 0xd2, 0x05, 0x7d, 0xd9, 0x6e, 0xf0,
	0x10, 0x9a, 0xc3, 0x58, 0x4c, 0xaa, 0xde, 0xcb, 0x52, 0xc6, 0xd1, 0x6f, 0xf8, 0x26, 0xd8, 0x33,
	0x09, 0x52, 0x6f, 0x09, 0x90, 0xd1, 0x05, 0x9e, 0x52, 0xaf, 0xd7, 0x50, 0x0e, 0x12, 0x60, 0xf4,
	0x41, 0xbb, 0x93, 0x1f, 0x58, 0x90, 0x1b, 0x49, 0x93, 0x6d, 0xd7, 0xea, 0xe5, 0xe0, 0x4a, 0xcc,
	0x4a, 0x9d, 0x06, 0x2f, 0x0d, 0xe8, 0xa6, 0xa7, 0xbe, 0x26, 0xd1, 0x82, 0xba, 0x04, 0x8d, 0xa0,
	0x53, 0x58, 0x5d, 0xf4, 0x77, 0xe9, 0xaf, 0xfc, 0x4b, 0xe8, 0x1b, 0x53, 0xdc, 0xa8, 0xa1, 0x21,
	0xb4, 0x73, 0x2b, 0x89, 0xfe, 0x2c, 0xb9, 0xc5, 0xcd, 0xd6, 0x37, 0x24, 0x64, 0x89, 0x11, 0x74,
	0x0a, 0x7b, 0x95, 0x3b, 0x4b, 0x79, 0x5f, 0xf5, 0x8d, 0x29, 0x59, 0xc8, 0x84, 0xaf, 0x72, 0x10,
	0x28, 0x37, 0xab, 0x64, 0x9c, 0x7a, 0x39, 0xc2, 0x8d, 0xda, 0xd9, 0xe0, 0xf5, 0xed, 0x7f, 0xfd,
	0xde, 0xcc, 0xbd, 0x1f, 0x01, 0x8d, 0xf0, 0x93, 0x4f, 0x58, 0xe4, 0x53, 0x1c, 0xd8, 0xa9, 0xc2,
	0xe6, 0x91, 0x6b, 0xe3, 0x90, 0x1e, 0xe1, 0x90, 0x8e, 0xbf, 0xab, 0x87, 0xe3, 0xe0, 0x3d, 0x00,
	0x00, 0xff, 0xff, 0x6b, 0xdb, 0xd5, 0x04, 0x8a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error)
	ReadAccount(ctx context.Context, in *ReadAccountReq, opts ...grpc.CallOption) (*ReadAccountRes, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error)
	// rpc DeleteAccount (DeleteAccountReq) returns (DeleteAccountRes) {};
	// rpc ListAccounts (ListAccountsReq) returns (ListAccountsRes) {};
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRes, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountReq, opts ...grpc.CallOption) (*CreateAccountRes, error) {
	out := new(CreateAccountRes)
	err := c.cc.Invoke(ctx, "/account.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ReadAccount(ctx context.Context, in *ReadAccountReq, opts ...grpc.CallOption) (*ReadAccountRes, error) {
	out := new(ReadAccountRes)
	err := c.cc.Invoke(ctx, "/account.AccountService/ReadAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error) {
	out := new(UpdateAccountRes)
	err := c.cc.Invoke(ctx, "/account.AccountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRes, error) {
	out := new(AuthRes)
	err := c.cc.Invoke(ctx, "/account.AccountService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRes, error)
	ReadAccount(context.Context, *ReadAccountReq) (*ReadAccountRes, error)
	UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRes, error)
	// rpc DeleteAccount (DeleteAccountReq) returns (DeleteAccountRes) {};
	// rpc ListAccounts (ListAccountsReq) returns (ListAccountsRes) {};
	Auth(context.Context, *AuthReq) (*AuthRes, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ReadAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ReadAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/ReadAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ReadAccount(ctx, req.(*ReadAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "ReadAccount",
			Handler:    _AccountService_ReadAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _AccountService_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
