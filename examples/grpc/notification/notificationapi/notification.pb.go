// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package notificationapi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type NotificationStatus int32

const (
	NotificationStatus_UNKNOWN NotificationStatus = 0
	NotificationStatus_SENT    NotificationStatus = 1
	NotificationStatus_PENDING NotificationStatus = 2
	NotificationStatus_FAILED  NotificationStatus = 3
)

var NotificationStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SENT",
	2: "PENDING",
	3: "FAILED",
}

var NotificationStatus_value = map[string]int32{
	"UNKNOWN": 0,
	"SENT":    1,
	"PENDING": 2,
	"FAILED":  3,
}

func (x NotificationStatus) String() string {
	return proto.EnumName(NotificationStatus_name, int32(x))
}

func (NotificationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

type SendEmailRequest struct {
	Email                *EmailMessage `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SendEmailRequest) Reset()         { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()    {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *SendEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailRequest.Unmarshal(m, b)
}
func (m *SendEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailRequest.Marshal(b, m, deterministic)
}
func (m *SendEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailRequest.Merge(m, src)
}
func (m *SendEmailRequest) XXX_Size() int {
	return xxx_messageInfo_SendEmailRequest.Size(m)
}
func (m *SendEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailRequest proto.InternalMessageInfo

func (m *SendEmailRequest) GetEmail() *EmailMessage {
	if m != nil {
		return m.Email
	}
	return nil
}

type EmailMessage struct {
	RecipientEmailAddress string   `protobuf:"bytes,1,opt,name=recipientEmailAddress,proto3" json:"recipientEmailAddress,omitempty"`
	Subject               string   `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Body                  string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *EmailMessage) Reset()         { *m = EmailMessage{} }
func (m *EmailMessage) String() string { return proto.CompactTextString(m) }
func (*EmailMessage) ProtoMessage()    {}
func (*EmailMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *EmailMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailMessage.Unmarshal(m, b)
}
func (m *EmailMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailMessage.Marshal(b, m, deterministic)
}
func (m *EmailMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailMessage.Merge(m, src)
}
func (m *EmailMessage) XXX_Size() int {
	return xxx_messageInfo_EmailMessage.Size(m)
}
func (m *EmailMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EmailMessage proto.InternalMessageInfo

func (m *EmailMessage) GetRecipientEmailAddress() string {
	if m != nil {
		return m.RecipientEmailAddress
	}
	return ""
}

func (m *EmailMessage) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type SendEmailReply struct {
	NotificationUid      string   `protobuf:"bytes,1,opt,name=NotificationUid,proto3" json:"NotificationUid,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailReply) Reset()         { *m = SendEmailReply{} }
func (m *SendEmailReply) String() string { return proto.CompactTextString(m) }
func (*SendEmailReply) ProtoMessage()    {}
func (*SendEmailReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{2}
}

func (m *SendEmailReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailReply.Unmarshal(m, b)
}
func (m *SendEmailReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailReply.Marshal(b, m, deterministic)
}
func (m *SendEmailReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailReply.Merge(m, src)
}
func (m *SendEmailReply) XXX_Size() int {
	return xxx_messageInfo_SendEmailReply.Size(m)
}
func (m *SendEmailReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailReply proto.InternalMessageInfo

func (m *SendEmailReply) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *SendEmailReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type SendSmsRequest struct {
	Sms                  *SmsMessage `protobuf:"bytes,1,opt,name=sms,proto3" json:"sms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SendSmsRequest) Reset()         { *m = SendSmsRequest{} }
func (m *SendSmsRequest) String() string { return proto.CompactTextString(m) }
func (*SendSmsRequest) ProtoMessage()    {}
func (*SendSmsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{3}
}

func (m *SendSmsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSmsRequest.Unmarshal(m, b)
}
func (m *SendSmsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSmsRequest.Marshal(b, m, deterministic)
}
func (m *SendSmsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSmsRequest.Merge(m, src)
}
func (m *SendSmsRequest) XXX_Size() int {
	return xxx_messageInfo_SendSmsRequest.Size(m)
}
func (m *SendSmsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSmsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendSmsRequest proto.InternalMessageInfo

func (m *SendSmsRequest) GetSms() *SmsMessage {
	if m != nil {
		return m.Sms
	}
	return nil
}

type SmsMessage struct {
	RecipientPhoneNumber string   `protobuf:"bytes,1,opt,name=recipientPhoneNumber,proto3" json:"recipientPhoneNumber,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmsMessage) Reset()         { *m = SmsMessage{} }
func (m *SmsMessage) String() string { return proto.CompactTextString(m) }
func (*SmsMessage) ProtoMessage()    {}
func (*SmsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{4}
}

func (m *SmsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmsMessage.Unmarshal(m, b)
}
func (m *SmsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmsMessage.Marshal(b, m, deterministic)
}
func (m *SmsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmsMessage.Merge(m, src)
}
func (m *SmsMessage) XXX_Size() int {
	return xxx_messageInfo_SmsMessage.Size(m)
}
func (m *SmsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SmsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SmsMessage proto.InternalMessageInfo

func (m *SmsMessage) GetRecipientPhoneNumber() string {
	if m != nil {
		return m.RecipientPhoneNumber
	}
	return ""
}

func (m *SmsMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type SendSmsReply struct {
	NotificationUid      string   `protobuf:"bytes,1,opt,name=NotificationUid,proto3" json:"NotificationUid,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendSmsReply) Reset()         { *m = SendSmsReply{} }
func (m *SendSmsReply) String() string { return proto.CompactTextString(m) }
func (*SendSmsReply) ProtoMessage()    {}
func (*SendSmsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{5}
}

func (m *SendSmsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSmsReply.Unmarshal(m, b)
}
func (m *SendSmsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSmsReply.Marshal(b, m, deterministic)
}
func (m *SendSmsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSmsReply.Merge(m, src)
}
func (m *SendSmsReply) XXX_Size() int {
	return xxx_messageInfo_SendSmsReply.Size(m)
}
func (m *SendSmsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSmsReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendSmsReply proto.InternalMessageInfo

func (m *SendSmsReply) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

func (m *SendSmsReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetNotificationStatusRequest struct {
	NotificationUid      string   `protobuf:"bytes,1,opt,name=notificationUid,proto3" json:"notificationUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNotificationStatusRequest) Reset()         { *m = GetNotificationStatusRequest{} }
func (m *GetNotificationStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetNotificationStatusRequest) ProtoMessage()    {}
func (*GetNotificationStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{6}
}

func (m *GetNotificationStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotificationStatusRequest.Unmarshal(m, b)
}
func (m *GetNotificationStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotificationStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetNotificationStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationStatusRequest.Merge(m, src)
}
func (m *GetNotificationStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetNotificationStatusRequest.Size(m)
}
func (m *GetNotificationStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationStatusRequest proto.InternalMessageInfo

func (m *GetNotificationStatusRequest) GetNotificationUid() string {
	if m != nil {
		return m.NotificationUid
	}
	return ""
}

type GetNotificationStatusReply struct {
	Status               NotificationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=notificationapi.NotificationStatus" json:"status,omitempty"`
	Error                *Error             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetNotificationStatusReply) Reset()         { *m = GetNotificationStatusReply{} }
func (m *GetNotificationStatusReply) String() string { return proto.CompactTextString(m) }
func (*GetNotificationStatusReply) ProtoMessage()    {}
func (*GetNotificationStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{7}
}

func (m *GetNotificationStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNotificationStatusReply.Unmarshal(m, b)
}
func (m *GetNotificationStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNotificationStatusReply.Marshal(b, m, deterministic)
}
func (m *GetNotificationStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNotificationStatusReply.Merge(m, src)
}
func (m *GetNotificationStatusReply) XXX_Size() int {
	return xxx_messageInfo_GetNotificationStatusReply.Size(m)
}
func (m *GetNotificationStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNotificationStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetNotificationStatusReply proto.InternalMessageInfo

func (m *GetNotificationStatusReply) GetStatus() NotificationStatus {
	if m != nil {
		return m.Status
	}
	return NotificationStatus_UNKNOWN
}

func (m *GetNotificationStatusReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{8}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("notificationapi.NotificationStatus", NotificationStatus_name, NotificationStatus_value)
	proto.RegisterType((*SendEmailRequest)(nil), "notificationapi.SendEmailRequest")
	proto.RegisterType((*EmailMessage)(nil), "notificationapi.EmailMessage")
	proto.RegisterType((*SendEmailReply)(nil), "notificationapi.SendEmailReply")
	proto.RegisterType((*SendSmsRequest)(nil), "notificationapi.SendSmsRequest")
	proto.RegisterType((*SmsMessage)(nil), "notificationapi.SmsMessage")
	proto.RegisterType((*SendSmsReply)(nil), "notificationapi.SendSmsReply")
	proto.RegisterType((*GetNotificationStatusRequest)(nil), "notificationapi.GetNotificationStatusRequest")
	proto.RegisterType((*GetNotificationStatusReply)(nil), "notificationapi.GetNotificationStatusReply")
	proto.RegisterType((*Error)(nil), "notificationapi.Error")
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07) }

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x93, 0x34, 0xc9, 0xdf, 0x49, 0xd4, 0x5a, 0x23, 0x02, 0x91, 0x49, 0x45, 0x59, 0x24,
	0x14, 0x95, 0x36, 0x91, 0x5c, 0xb8, 0x94, 0x03, 0xaa, 0x94, 0x10, 0x2a, 0xc0, 0x54, 0x76, 0x2b,
	0xce, 0x4e, 0xbc, 0x4d, 0x8d, 0x62, 0xaf, 0xf1, 0x6e, 0x84, 0x22, 0xc4, 0x85, 0x0b, 0x0f, 0xc0,
	0x0b, 0x70, 0xe3, 0x81, 0x78, 0x05, 0x1e, 0x04, 0xed, 0xda, 0x71, 0x8d, 0xed, 0x56, 0x70, 0xe0,
	0xe6, 0x99, 0xf9, 0x66, 0xe6, 0x9b, 0xd9, 0x6f, 0x0c, 0x18, 0x30, 0xe1, 0x5d, 0x78, 0x33, 0x47,
	0x78, 0x2c, 0x18, 0x84, 0x11, 0x13, 0x0c, 0xb7, 0xb3, 0x3e, 0x27, 0xf4, 0xf4, 0xde, 0x9c, 0xb1,
	0xf9, 0x82, 0x0e, 0x9d, 0xd0, 0x1b, 0x3a, 0x41, 0xc0, 0x84, 0x8a, 0xf0, 0x18, 0x4e, 0x26, 0xa0,
	0xd9, 0x34, 0x70, 0xc7, 0xbe, 0xe3, 0x2d, 0x2c, 0xfa, 0x7e, 0x49, 0xb9, 0xc0, 0x43, 0xa8, 0x53,
	0x69, 0x77, 0x2b, 0xbb, 0x95, 0x7e, 0xcb, 0xd8, 0x19, 0xe4, 0x4a, 0x0e, 0x14, 0xfa, 0x35, 0xe5,
	0xdc, 0x99, 0x53, 0x2b, 0xc6, 0x92, 0x08, 0xda, 0x59, 0x37, 0x3e, 0x86, 0x4e, 0x44, 0x67, 0x5e,
	0xe8, 0xd1, 0x40, 0xa8, 0xc0, 0xb1, 0xeb, 0x46, 0x94, 0x73, 0x55, 0x74, 0xd3, 0x2a, 0x0f, 0x62,
	0x17, 0x9a, 0x7c, 0x39, 0x7d, 0x47, 0x67, 0xa2, 0x5b, 0x55, 0xb8, 0xb5, 0x89, 0x08, 0x1b, 0x53,
	0xe6, 0xae, 0xba, 0x35, 0xe5, 0x56, 0xdf, 0xe4, 0x12, 0xb6, 0x32, 0xe4, 0xc3, 0xc5, 0x0a, 0xfb,
	0xb0, 0x6d, 0x66, 0xc8, 0x9e, 0x7b, 0x6e, 0xd2, 0x2f, 0xef, 0xc6, 0x7d, 0xa8, 0xd3, 0x28, 0x62,
	0x91, 0xea, 0xd3, 0x32, 0x6e, 0x17, 0x87, 0x94, 0x51, 0x2b, 0x06, 0x91, 0x67, 0x71, 0x27, 0xdb,
	0xe7, 0xeb, 0x25, 0x1d, 0x40, 0x8d, 0xfb, 0x3c, 0x59, 0xd1, 0xdd, 0x42, 0xb6, 0xed, 0xf3, 0xf5,
	0x82, 0x24, 0x8e, 0x9c, 0x01, 0x5c, 0xb9, 0xd0, 0x80, 0x5b, 0xe9, 0xfc, 0xa7, 0x97, 0x2c, 0xa0,
	0xe6, 0xd2, 0x9f, 0xd2, 0x28, 0xe1, 0x5a, 0x1a, 0x4b, 0x17, 0x50, 0xcd, 0x2c, 0xe0, 0x02, 0xda,
	0x29, 0xad, 0x7f, 0x39, 0xfe, 0x0b, 0xe8, 0x4d, 0xa8, 0xc8, 0xd6, 0xb0, 0x85, 0x23, 0x96, 0xe9,
	0x32, 0xfa, 0xf0, 0x9b, 0xec, 0x32, 0x7d, 0x73, 0x6e, 0xf2, 0xa5, 0x02, 0xfa, 0x35, 0xa5, 0xe4,
	0x00, 0x4f, 0xa1, 0xc1, 0x95, 0xa9, 0xf2, 0xb7, 0x8c, 0x07, 0x05, 0x5e, 0x25, 0x99, 0x49, 0xca,
	0x5f, 0xce, 0xf4, 0x04, 0xea, 0xca, 0x96, 0x8b, 0x9d, 0x31, 0x97, 0xaa, 0x8e, 0x75, 0x4b, 0x7d,
	0x4b, 0x1d, 0xfa, 0xf1, 0x5b, 0xad, 0x75, 0x98, 0x98, 0x7b, 0x23, 0xc0, 0x22, 0x05, 0x6c, 0x41,
	0xf3, 0xdc, 0x7c, 0x69, 0xbe, 0x79, 0x6b, 0x6a, 0xff, 0xe1, 0xff, 0xb0, 0x61, 0x8f, 0xcd, 0x33,
	0xad, 0x22, 0xdd, 0xa7, 0x63, 0x73, 0x74, 0x62, 0x4e, 0xb4, 0x2a, 0x02, 0x34, 0x9e, 0x1f, 0x9f,
	0xbc, 0x1a, 0x8f, 0xb4, 0x9a, 0xf1, 0xad, 0x06, 0xed, 0x6c, 0x19, 0xfc, 0x00, 0x9b, 0xa9, 0x94,
	0xf1, 0x7e, 0x51, 0x4e, 0xb9, 0x1b, 0xd5, 0xef, 0xdd, 0x04, 0x09, 0x17, 0x2b, 0xf2, 0xf0, 0xf3,
	0x8f, 0x9f, 0x5f, 0xab, 0xbb, 0xe4, 0x8e, 0x3a, 0xfc, 0x2c, 0x78, 0xa8, 0x0e, 0xf6, 0x28, 0xbe,
	0x5b, 0xf4, 0xa1, 0x99, 0x48, 0x08, 0xcb, 0x6b, 0x5e, 0x69, 0x5e, 0xdf, 0xb9, 0x1e, 0x20, 0x5b,
	0x12, 0xd5, 0xb2, 0x47, 0x3a, 0xc5, 0x96, 0xdc, 0xe7, 0x47, 0xf2, 0x0e, 0xf0, 0x7b, 0x05, 0x3a,
	0xa5, 0xef, 0x8f, 0x07, 0x85, 0xe2, 0x37, 0x49, 0x4e, 0x7f, 0xf4, 0xa7, 0x70, 0xc9, 0xcc, 0x50,
	0xcc, 0xf6, 0x71, 0xaf, 0x84, 0x99, 0x82, 0x0d, 0x3f, 0xe6, 0x84, 0xfa, 0x69, 0xda, 0x50, 0x3f,
	0xc8, 0xc3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xb5, 0xd6, 0x64, 0x65, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationClient interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailReply, error)
	SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsReply, error)
	GetNotificationStatus(ctx context.Context, in *GetNotificationStatusRequest, opts ...grpc.CallOption) (*GetNotificationStatusReply, error)
}

type notificationClient struct {
	cc *grpc.ClientConn
}

func NewNotificationClient(cc *grpc.ClientConn) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailReply, error) {
	out := new(SendEmailReply)
	err := c.cc.Invoke(ctx, "/notificationapi.Notification/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsReply, error) {
	out := new(SendSmsReply)
	err := c.cc.Invoke(ctx, "/notificationapi.Notification/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) GetNotificationStatus(ctx context.Context, in *GetNotificationStatusRequest, opts ...grpc.CallOption) (*GetNotificationStatusReply, error) {
	out := new(GetNotificationStatusReply)
	err := c.cc.Invoke(ctx, "/notificationapi.Notification/GetNotificationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
type NotificationServer interface {
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailReply, error)
	SendSms(context.Context, *SendSmsRequest) (*SendSmsReply, error)
	GetNotificationStatus(context.Context, *GetNotificationStatusRequest) (*GetNotificationStatusReply, error)
}

// UnimplementedNotificationServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (*UnimplementedNotificationServer) SendEmail(ctx context.Context, req *SendEmailRequest) (*SendEmailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (*UnimplementedNotificationServer) SendSms(ctx context.Context, req *SendSmsRequest) (*SendSmsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (*UnimplementedNotificationServer) GetNotificationStatus(ctx context.Context, req *GetNotificationStatusRequest) (*GetNotificationStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationStatus not implemented")
}

func RegisterNotificationServer(s *grpc.Server, srv NotificationServer) {
	s.RegisterService(&_Notification_serviceDesc, srv)
}

func _Notification_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificationapi.Notification/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificationapi.Notification/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).SendSms(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_GetNotificationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).GetNotificationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notificationapi.Notification/GetNotificationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).GetNotificationStatus(ctx, req.(*GetNotificationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notification_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notificationapi.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _Notification_SendEmail_Handler,
		},
		{
			MethodName: "SendSms",
			Handler:    _Notification_SendSms_Handler,
		},
		{
			MethodName: "GetNotificationStatus",
			Handler:    _Notification_GetNotificationStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}