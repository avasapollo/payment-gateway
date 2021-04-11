// Code generated by protoc-gen-go. DO NOT EDIT.
// source: web/proto/v1/payment-gateway.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HealthResp struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResp) Reset()         { *m = HealthResp{} }
func (m *HealthResp) String() string { return proto.CompactTextString(m) }
func (*HealthResp) ProtoMessage()    {}
func (*HealthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{0}
}

func (m *HealthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResp.Unmarshal(m, b)
}
func (m *HealthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResp.Marshal(b, m, deterministic)
}
func (m *HealthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResp.Merge(m, src)
}
func (m *HealthResp) XXX_Size() int {
	return xxx_messageInfo_HealthResp.Size(m)
}
func (m *HealthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResp.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResp proto.InternalMessageInfo

func (m *HealthResp) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Card struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CardNumber           string   `protobuf:"bytes,2,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpireMonth          string   `protobuf:"bytes,3,opt,name=expire_month,json=expireMonth,proto3" json:"expire_month,omitempty"`
	ExpireYear           string   `protobuf:"bytes,4,opt,name=expire_year,json=expireYear,proto3" json:"expire_year,omitempty"`
	Cvv                  string   `protobuf:"bytes,5,opt,name=cvv,proto3" json:"cvv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{1}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Card) GetCardNumber() string {
	if m != nil {
		return m.CardNumber
	}
	return ""
}

func (m *Card) GetExpireMonth() string {
	if m != nil {
		return m.ExpireMonth
	}
	return ""
}

func (m *Card) GetExpireYear() string {
	if m != nil {
		return m.ExpireYear
	}
	return ""
}

func (m *Card) GetCvv() string {
	if m != nil {
		return m.Cvv
	}
	return ""
}

type Amount struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Amount) Reset()         { *m = Amount{} }
func (m *Amount) String() string { return proto.CompactTextString(m) }
func (*Amount) ProtoMessage()    {}
func (*Amount) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{2}
}

func (m *Amount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Amount.Unmarshal(m, b)
}
func (m *Amount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Amount.Marshal(b, m, deterministic)
}
func (m *Amount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Amount.Merge(m, src)
}
func (m *Amount) XXX_Size() int {
	return xxx_messageInfo_Amount.Size(m)
}
func (m *Amount) XXX_DiscardUnknown() {
	xxx_messageInfo_Amount.DiscardUnknown(m)
}

var xxx_messageInfo_Amount proto.InternalMessageInfo

func (m *Amount) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Amount) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type AuthorizeReq struct {
	Card                 *Card    `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Amount               *Amount  `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeReq) Reset()         { *m = AuthorizeReq{} }
func (m *AuthorizeReq) String() string { return proto.CompactTextString(m) }
func (*AuthorizeReq) ProtoMessage()    {}
func (*AuthorizeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{3}
}

func (m *AuthorizeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeReq.Unmarshal(m, b)
}
func (m *AuthorizeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeReq.Marshal(b, m, deterministic)
}
func (m *AuthorizeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeReq.Merge(m, src)
}
func (m *AuthorizeReq) XXX_Size() int {
	return xxx_messageInfo_AuthorizeReq.Size(m)
}
func (m *AuthorizeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeReq.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeReq proto.InternalMessageInfo

func (m *AuthorizeReq) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *AuthorizeReq) GetAmount() *Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

type AuthorizeResp struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	AuthorizationId      string   `protobuf:"bytes,2,opt,name=authorization_id,json=authorizationId,proto3" json:"authorization_id,omitempty"`
	Amount               *Amount  `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeResp) Reset()         { *m = AuthorizeResp{} }
func (m *AuthorizeResp) String() string { return proto.CompactTextString(m) }
func (*AuthorizeResp) ProtoMessage()    {}
func (*AuthorizeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{4}
}

func (m *AuthorizeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeResp.Unmarshal(m, b)
}
func (m *AuthorizeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeResp.Marshal(b, m, deterministic)
}
func (m *AuthorizeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeResp.Merge(m, src)
}
func (m *AuthorizeResp) XXX_Size() int {
	return xxx_messageInfo_AuthorizeResp.Size(m)
}
func (m *AuthorizeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeResp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeResp proto.InternalMessageInfo

func (m *AuthorizeResp) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *AuthorizeResp) GetAuthorizationId() string {
	if m != nil {
		return m.AuthorizationId
	}
	return ""
}

func (m *AuthorizeResp) GetAmount() *Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

type AmountResp struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Amount               *Amount  `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AmountResp) Reset()         { *m = AmountResp{} }
func (m *AmountResp) String() string { return proto.CompactTextString(m) }
func (*AmountResp) ProtoMessage()    {}
func (*AmountResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{5}
}

func (m *AmountResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmountResp.Unmarshal(m, b)
}
func (m *AmountResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmountResp.Marshal(b, m, deterministic)
}
func (m *AmountResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmountResp.Merge(m, src)
}
func (m *AmountResp) XXX_Size() int {
	return xxx_messageInfo_AmountResp.Size(m)
}
func (m *AmountResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AmountResp.DiscardUnknown(m)
}

var xxx_messageInfo_AmountResp proto.InternalMessageInfo

func (m *AmountResp) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *AmountResp) GetAmount() *Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

type VoidReq struct {
	AuthorizationId      string   `protobuf:"bytes,1,opt,name=authorization_id,json=authorizationId,proto3" json:"authorization_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoidReq) Reset()         { *m = VoidReq{} }
func (m *VoidReq) String() string { return proto.CompactTextString(m) }
func (*VoidReq) ProtoMessage()    {}
func (*VoidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{6}
}

func (m *VoidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoidReq.Unmarshal(m, b)
}
func (m *VoidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoidReq.Marshal(b, m, deterministic)
}
func (m *VoidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoidReq.Merge(m, src)
}
func (m *VoidReq) XXX_Size() int {
	return xxx_messageInfo_VoidReq.Size(m)
}
func (m *VoidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VoidReq.DiscardUnknown(m)
}

var xxx_messageInfo_VoidReq proto.InternalMessageInfo

func (m *VoidReq) GetAuthorizationId() string {
	if m != nil {
		return m.AuthorizationId
	}
	return ""
}

type CaptureReq struct {
	AuthorizationId      string   `protobuf:"bytes,1,opt,name=authorization_id,json=authorizationId,proto3" json:"authorization_id,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CaptureReq) Reset()         { *m = CaptureReq{} }
func (m *CaptureReq) String() string { return proto.CompactTextString(m) }
func (*CaptureReq) ProtoMessage()    {}
func (*CaptureReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{7}
}

func (m *CaptureReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CaptureReq.Unmarshal(m, b)
}
func (m *CaptureReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CaptureReq.Marshal(b, m, deterministic)
}
func (m *CaptureReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CaptureReq.Merge(m, src)
}
func (m *CaptureReq) XXX_Size() int {
	return xxx_messageInfo_CaptureReq.Size(m)
}
func (m *CaptureReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CaptureReq.DiscardUnknown(m)
}

var xxx_messageInfo_CaptureReq proto.InternalMessageInfo

func (m *CaptureReq) GetAuthorizationId() string {
	if m != nil {
		return m.AuthorizationId
	}
	return ""
}

func (m *CaptureReq) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type RefundReq struct {
	AuthorizationId      string   `protobuf:"bytes,1,opt,name=authorization_id,json=authorizationId,proto3" json:"authorization_id,omitempty"`
	Amount               float64  `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefundReq) Reset()         { *m = RefundReq{} }
func (m *RefundReq) String() string { return proto.CompactTextString(m) }
func (*RefundReq) ProtoMessage()    {}
func (*RefundReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_55b1559c3d23429d, []int{8}
}

func (m *RefundReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefundReq.Unmarshal(m, b)
}
func (m *RefundReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefundReq.Marshal(b, m, deterministic)
}
func (m *RefundReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefundReq.Merge(m, src)
}
func (m *RefundReq) XXX_Size() int {
	return xxx_messageInfo_RefundReq.Size(m)
}
func (m *RefundReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RefundReq.DiscardUnknown(m)
}

var xxx_messageInfo_RefundReq proto.InternalMessageInfo

func (m *RefundReq) GetAuthorizationId() string {
	if m != nil {
		return m.AuthorizationId
	}
	return ""
}

func (m *RefundReq) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*HealthResp)(nil), "payment.gateway.v1.HealthResp")
	proto.RegisterType((*Card)(nil), "payment.gateway.v1.Card")
	proto.RegisterType((*Amount)(nil), "payment.gateway.v1.Amount")
	proto.RegisterType((*AuthorizeReq)(nil), "payment.gateway.v1.AuthorizeReq")
	proto.RegisterType((*AuthorizeResp)(nil), "payment.gateway.v1.AuthorizeResp")
	proto.RegisterType((*AmountResp)(nil), "payment.gateway.v1.AmountResp")
	proto.RegisterType((*VoidReq)(nil), "payment.gateway.v1.VoidReq")
	proto.RegisterType((*CaptureReq)(nil), "payment.gateway.v1.CaptureReq")
	proto.RegisterType((*RefundReq)(nil), "payment.gateway.v1.RefundReq")
}

func init() { proto.RegisterFile("web/proto/v1/payment-gateway.proto", fileDescriptor_55b1559c3d23429d) }

var fileDescriptor_55b1559c3d23429d = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x96, 0x9b, 0xd4, 0x6d, 0x26, 0x49, 0x13, 0xcd, 0xfb, 0x12, 0x59, 0x2e, 0x94, 0x76, 0xc5,
	0x01, 0x2a, 0xb0, 0xd5, 0xc0, 0x29, 0xb7, 0x52, 0x21, 0xe0, 0x40, 0x01, 0x1f, 0x50, 0xe1, 0x40,
	0xb4, 0x89, 0xb7, 0x89, 0x51, 0xfc, 0xc1, 0x7a, 0xed, 0x12, 0x8e, 0x1c, 0xb8, 0x23, 0x7e, 0x1a,
	0x7f, 0x81, 0x23, 0x3f, 0x02, 0xed, 0x47, 0xd2, 0x20, 0x2c, 0x02, 0x88, 0xdb, 0xce, 0x33, 0xb3,
	0xcf, 0x3c, 0x33, 0xb3, 0xb3, 0x40, 0x2e, 0xd8, 0xc8, 0xcf, 0x78, 0x2a, 0x52, 0xbf, 0x3c, 0xf2,
	0x33, 0x3a, 0x8f, 0x59, 0x22, 0xee, 0x4c, 0xa8, 0x60, 0x17, 0x74, 0xee, 0x29, 0x07, 0xa2, 0x81,
	0xbd, 0x05, 0x5c, 0x1e, 0xb9, 0x57, 0x27, 0x69, 0x3a, 0x99, 0x31, 0x9f, 0x66, 0x91, 0x4f, 0x93,
	0x24, 0x15, 0x54, 0x44, 0x69, 0x92, 0xeb, 0x1b, 0xee, 0xae, 0xf1, 0x2a, 0x6b, 0x54, 0x9c, 0xfb,
	0x2c, 0xce, 0x84, 0xa1, 0x23, 0x37, 0x00, 0x1e, 0x31, 0x3a, 0x13, 0xd3, 0x80, 0xe5, 0x19, 0xf6,
	0xc0, 0xce, 0x05, 0x15, 0x45, 0xee, 0x58, 0xfb, 0xd6, 0xcd, 0x46, 0x60, 0x2c, 0xf2, 0xc9, 0x82,
	0xfa, 0x09, 0xe5, 0x21, 0x22, 0xd4, 0x13, 0x1a, 0x33, 0xe3, 0x56, 0x67, 0xbc, 0x0e, 0xcd, 0x31,
	0xe5, 0xe1, 0x30, 0x29, 0xe2, 0x11, 0xe3, 0xce, 0x86, 0x72, 0x81, 0x84, 0x4e, 0x15, 0x82, 0x07,
	0xd0, 0x62, 0xef, 0xb2, 0x88, 0xb3, 0x61, 0x9c, 0x26, 0x62, 0xea, 0xd4, 0x54, 0x44, 0x53, 0x63,
	0x4f, 0x24, 0x24, 0x39, 0x4c, 0xc8, 0x9c, 0x51, 0xee, 0xd4, 0x35, 0x87, 0x86, 0x5e, 0x32, 0xca,
	0xb1, 0x0b, 0xb5, 0x71, 0x59, 0x3a, 0x9b, 0xca, 0x21, 0x8f, 0x64, 0x00, 0xf6, 0x71, 0x9c, 0x16,
	0x89, 0xc0, 0xff, 0x61, 0xb3, 0xa4, 0xb3, 0x42, 0xab, 0xb2, 0x02, 0x6d, 0xa0, 0x0b, 0xdb, 0xe3,
	0x82, 0x73, 0x96, 0x8c, 0xe7, 0x46, 0xd3, 0xd2, 0x26, 0x19, 0xb4, 0x8e, 0x0b, 0x31, 0x4d, 0x79,
	0xf4, 0x9e, 0x05, 0xec, 0x2d, 0xde, 0x86, 0xba, 0xd4, 0xab, 0x08, 0x9a, 0x7d, 0xc7, 0xfb, 0xb9,
	0xc7, 0x9e, 0x2c, 0x3f, 0x50, 0x51, 0xd8, 0x07, 0x9b, 0xaa, 0xcc, 0x8a, 0xb7, 0xd9, 0x77, 0xab,
	0xe2, 0xb5, 0xb6, 0xc0, 0x44, 0x92, 0x8f, 0x16, 0xb4, 0x57, 0x52, 0xea, 0x5e, 0x73, 0x96, 0x17,
	0x33, 0xb1, 0xe8, 0xb5, 0xb6, 0xf0, 0x16, 0x74, 0xa9, 0x09, 0x54, 0x63, 0x1c, 0x46, 0xa1, 0xd1,
	0xdf, 0xf9, 0x01, 0x7f, 0xbc, 0x2a, 0xa4, 0xf6, 0xdb, 0x42, 0xce, 0x00, 0x0c, 0xf2, 0x2b, 0x11,
	0x7f, 0x53, 0xe2, 0x3d, 0xd8, 0x7a, 0x91, 0x46, 0xa1, 0xec, 0x67, 0x55, 0x0d, 0x56, 0x65, 0x0d,
	0xe4, 0x29, 0xc0, 0x09, 0xcd, 0x44, 0xc1, 0xd9, 0x9f, 0x5d, 0x94, 0xd2, 0x57, 0x24, 0x5a, 0x4b,
	0x19, 0xa7, 0xd0, 0x08, 0xd8, 0x79, 0x91, 0x84, 0xff, 0x86, 0xaf, 0xff, 0xad, 0x06, 0x3b, 0xcf,
	0x74, 0xf1, 0x0f, 0x75, 0xed, 0xf8, 0x1c, 0x6c, 0xbd, 0x34, 0xd8, 0xf3, 0xf4, 0x72, 0x79, 0x8b,
	0xe5, 0xf2, 0x1e, 0xc8, 0xe5, 0x72, 0xf7, 0xaa, 0xfa, 0x75, 0xb9, 0x68, 0xa4, 0xf3, 0xe1, 0xcb,
	0xd7, 0xcf, 0x1b, 0x0d, 0xdc, 0xf2, 0xa7, 0x9a, 0xe8, 0x0d, 0x34, 0x96, 0xcf, 0x03, 0xf7, 0x2b,
	0xbb, 0xbd, 0xf2, 0x60, 0xdd, 0x83, 0x35, 0x11, 0x79, 0x46, 0x1c, 0x95, 0x02, 0x49, 0x5b, 0xfe,
	0x25, 0x8b, 0x62, 0xd9, 0xc0, 0x3a, 0xc4, 0x33, 0xa8, 0xcb, 0x41, 0xe1, 0x6e, 0x15, 0x89, 0x19,
	0x61, 0x75, 0x05, 0x97, 0x2f, 0x87, 0xfc, 0xa7, 0xe8, 0xdb, 0x64, 0x5b, 0xd2, 0x97, 0x69, 0x14,
	0x4a, 0x66, 0x0a, 0x5b, 0x66, 0x98, 0xb8, 0x57, 0xbd, 0x44, 0x8b, 0x49, 0xaf, 0xe5, 0xef, 0x29,
	0xfe, 0x2e, 0x69, 0x4a, 0xfe, 0xb1, 0xbe, 0x27, 0x53, 0xbc, 0x06, 0x5b, 0x8f, 0x17, 0xaf, 0x55,
	0x31, 0x2c, 0x47, 0xbf, 0x36, 0xc1, 0x15, 0x95, 0xa0, 0x43, 0x40, 0x26, 0xe0, 0xea, 0xda, 0xc0,
	0x3a, 0xbc, 0xbf, 0xf3, 0xaa, 0xb5, 0xfa, 0x0b, 0x8f, 0x6c, 0x75, 0xba, 0xfb, 0x3d, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0x62, 0x24, 0x12, 0x9c, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentGatewayClient is the client API for PaymentGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentGatewayClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResp, error)
	Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeResp, error)
	Void(ctx context.Context, in *VoidReq, opts ...grpc.CallOption) (*AmountResp, error)
	Capture(ctx context.Context, in *CaptureReq, opts ...grpc.CallOption) (*AmountResp, error)
	Refund(ctx context.Context, in *RefundReq, opts ...grpc.CallOption) (*AmountResp, error)
}

type paymentGatewayClient struct {
	cc *grpc.ClientConn
}

func NewPaymentGatewayClient(cc *grpc.ClientConn) PaymentGatewayClient {
	return &paymentGatewayClient{cc}
}

func (c *paymentGatewayClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResp, error) {
	out := new(HealthResp)
	err := c.cc.Invoke(ctx, "/payment.gateway.v1.PaymentGateway/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) Authorize(ctx context.Context, in *AuthorizeReq, opts ...grpc.CallOption) (*AuthorizeResp, error) {
	out := new(AuthorizeResp)
	err := c.cc.Invoke(ctx, "/payment.gateway.v1.PaymentGateway/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) Void(ctx context.Context, in *VoidReq, opts ...grpc.CallOption) (*AmountResp, error) {
	out := new(AmountResp)
	err := c.cc.Invoke(ctx, "/payment.gateway.v1.PaymentGateway/Void", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) Capture(ctx context.Context, in *CaptureReq, opts ...grpc.CallOption) (*AmountResp, error) {
	out := new(AmountResp)
	err := c.cc.Invoke(ctx, "/payment.gateway.v1.PaymentGateway/Capture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) Refund(ctx context.Context, in *RefundReq, opts ...grpc.CallOption) (*AmountResp, error) {
	out := new(AmountResp)
	err := c.cc.Invoke(ctx, "/payment.gateway.v1.PaymentGateway/Refund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentGatewayServer is the server API for PaymentGateway service.
type PaymentGatewayServer interface {
	Health(context.Context, *empty.Empty) (*HealthResp, error)
	Authorize(context.Context, *AuthorizeReq) (*AuthorizeResp, error)
	Void(context.Context, *VoidReq) (*AmountResp, error)
	Capture(context.Context, *CaptureReq) (*AmountResp, error)
	Refund(context.Context, *RefundReq) (*AmountResp, error)
}

// UnimplementedPaymentGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentGatewayServer struct {
}

func (*UnimplementedPaymentGatewayServer) Health(ctx context.Context, req *empty.Empty) (*HealthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedPaymentGatewayServer) Authorize(ctx context.Context, req *AuthorizeReq) (*AuthorizeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (*UnimplementedPaymentGatewayServer) Void(ctx context.Context, req *VoidReq) (*AmountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Void not implemented")
}
func (*UnimplementedPaymentGatewayServer) Capture(ctx context.Context, req *CaptureReq) (*AmountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capture not implemented")
}
func (*UnimplementedPaymentGatewayServer) Refund(ctx context.Context, req *RefundReq) (*AmountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refund not implemented")
}

func RegisterPaymentGatewayServer(s *grpc.Server, srv PaymentGatewayServer) {
	s.RegisterService(&_PaymentGateway_serviceDesc, srv)
}

func _PaymentGateway_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.gateway.v1.PaymentGateway/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.gateway.v1.PaymentGateway/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).Authorize(ctx, req.(*AuthorizeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_Void_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).Void(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.gateway.v1.PaymentGateway/Void",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).Void(ctx, req.(*VoidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_Capture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptureReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).Capture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.gateway.v1.PaymentGateway/Capture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).Capture(ctx, req.(*CaptureReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_Refund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).Refund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.gateway.v1.PaymentGateway/Refund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).Refund(ctx, req.(*RefundReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payment.gateway.v1.PaymentGateway",
	HandlerType: (*PaymentGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _PaymentGateway_Health_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _PaymentGateway_Authorize_Handler,
		},
		{
			MethodName: "Void",
			Handler:    _PaymentGateway_Void_Handler,
		},
		{
			MethodName: "Capture",
			Handler:    _PaymentGateway_Capture_Handler,
		},
		{
			MethodName: "Refund",
			Handler:    _PaymentGateway_Refund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web/proto/v1/payment-gateway.proto",
}
