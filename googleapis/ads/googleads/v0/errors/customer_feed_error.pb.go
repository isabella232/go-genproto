// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/customer_feed_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible customer feed errors.
type CustomerFeedErrorEnum_CustomerFeedError int32

const (
	// Enum unspecified.
	CustomerFeedErrorEnum_UNSPECIFIED CustomerFeedErrorEnum_CustomerFeedError = 0
	// The received error code is not known in this version.
	CustomerFeedErrorEnum_UNKNOWN CustomerFeedErrorEnum_CustomerFeedError = 1
	// An active feed already exists for this customer and place holder type.
	CustomerFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 2
	// The specified feed is removed.
	CustomerFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED CustomerFeedErrorEnum_CustomerFeedError = 3
	// The CustomerFeed already exists. Update should be used to modify the
	// existing CustomerFeed.
	CustomerFeedErrorEnum_CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 4
	// Cannot update removed customer feed.
	CustomerFeedErrorEnum_CANNOT_MODIFY_REMOVED_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 5
	// Invalid placeholder type.
	CustomerFeedErrorEnum_INVALID_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 6
	// Feed mapping for this placeholder type does not exist.
	CustomerFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 7
	// Placeholder not allowed at the account level.
	CustomerFeedErrorEnum_PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 8
)

var CustomerFeedErrorEnum_CustomerFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	3: "CANNOT_CREATE_FOR_REMOVED_FEED",
	4: "CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED",
	5: "CANNOT_MODIFY_REMOVED_CUSTOMER_FEED",
	6: "INVALID_PLACEHOLDER_TYPE",
	7: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
	8: "PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED",
}
var CustomerFeedErrorEnum_CustomerFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":      2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":                3,
	"CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED":  4,
	"CANNOT_MODIFY_REMOVED_CUSTOMER_FEED":           5,
	"INVALID_PLACEHOLDER_TYPE":                      6,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE":      7,
	"PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED": 8,
}

func (x CustomerFeedErrorEnum_CustomerFeedError) String() string {
	return proto.EnumName(CustomerFeedErrorEnum_CustomerFeedError_name, int32(x))
}
func (CustomerFeedErrorEnum_CustomerFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_customer_feed_error_7d219f169025f9b2, []int{0, 0}
}

// Container for enum describing possible customer feed errors.
type CustomerFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerFeedErrorEnum) Reset()         { *m = CustomerFeedErrorEnum{} }
func (m *CustomerFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerFeedErrorEnum) ProtoMessage()    {}
func (*CustomerFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_feed_error_7d219f169025f9b2, []int{0}
}
func (m *CustomerFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerFeedErrorEnum.Unmarshal(m, b)
}
func (m *CustomerFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerFeedErrorEnum.Marshal(b, m, deterministic)
}
func (dst *CustomerFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerFeedErrorEnum.Merge(dst, src)
}
func (m *CustomerFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerFeedErrorEnum.Size(m)
}
func (m *CustomerFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomerFeedErrorEnum)(nil), "google.ads.googleads.v0.errors.CustomerFeedErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.CustomerFeedErrorEnum_CustomerFeedError", CustomerFeedErrorEnum_CustomerFeedError_name, CustomerFeedErrorEnum_CustomerFeedError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/customer_feed_error.proto", fileDescriptor_customer_feed_error_7d219f169025f9b2)
}

var fileDescriptor_customer_feed_error_7d219f169025f9b2 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x69, 0x06, 0x1b, 0xf2, 0x0e, 0x14, 0x4b, 0x20, 0x0e, 0xa8, 0x87, 0x72, 0x80, 0xc3,
	0x70, 0x82, 0xb8, 0x20, 0x71, 0xf2, 0xe2, 0x7f, 0x8a, 0x45, 0x62, 0x47, 0x49, 0xda, 0x51, 0x14,
	0xc9, 0x2a, 0x8b, 0x89, 0x90, 0xd6, 0x7a, 0x8a, 0xb7, 0x3d, 0x10, 0x47, 0x1e, 0x05, 0xf1, 0x0a,
	0x3c, 0x07, 0x57, 0xe4, 0xb8, 0xab, 0xd4, 0x16, 0x76, 0xca, 0xa7, 0x7f, 0x7e, 0xdf, 0xff, 0xb3,
	0xf5, 0x19, 0xbd, 0x6b, 0x8d, 0x69, 0x2f, 0x74, 0xb8, 0x68, 0x6c, 0xe8, 0xa5, 0x53, 0x37, 0x51,
	0xa8, 0xbb, 0xce, 0x74, 0x36, 0x3c, 0xbf, 0xb6, 0x57, 0x66, 0xa9, 0x3b, 0xf5, 0x55, 0xeb, 0x46,
	0xf5, 0x43, 0x72, 0xd9, 0x99, 0x2b, 0x83, 0x47, 0x1e, 0x27, 0x8b, 0xc6, 0x92, 0x8d, 0x93, 0xdc,
	0x44, 0xc4, 0x3b, 0xc7, 0x7f, 0x02, 0xf4, 0x24, 0x5e, 0xbb, 0x13, 0xad, 0x1b, 0x70, 0x63, 0x58,
	0x5d, 0x2f, 0xc7, 0xbf, 0x03, 0xf4, 0x78, 0xef, 0x0f, 0x7e, 0x84, 0x8e, 0xa7, 0xa2, 0xcc, 0x21,
	0xe6, 0x09, 0x07, 0x36, 0xbc, 0x87, 0x8f, 0xd1, 0xd1, 0x54, 0x7c, 0x14, 0xf2, 0x4c, 0x0c, 0x07,
	0xf8, 0x04, 0xbd, 0x4a, 0x00, 0x98, 0xa2, 0x69, 0x01, 0x94, 0xcd, 0x15, 0x7c, 0xe2, 0x65, 0x55,
	0xaa, 0x44, 0x16, 0x2a, 0x4f, 0x69, 0x0c, 0x1f, 0x64, 0xca, 0xa0, 0x50, 0xd5, 0x3c, 0x87, 0x61,
	0x80, 0xc7, 0x68, 0x14, 0x53, 0x21, 0x64, 0xa5, 0xe2, 0x02, 0x68, 0x05, 0x3d, 0x57, 0x40, 0x26,
	0x67, 0xc0, 0x94, 0xdb, 0x33, 0x3c, 0xc0, 0x11, 0x3a, 0xd9, 0x66, 0xb6, 0x56, 0x73, 0x31, 0x51,
	0xf1, 0xb4, 0xac, 0x64, 0x06, 0x85, 0x77, 0xdc, 0xc7, 0x2f, 0xd1, 0x8b, 0xb5, 0x23, 0x93, 0x8c,
	0x27, 0xf3, 0xcd, 0xc6, 0x6d, 0xf0, 0x01, 0x7e, 0x8e, 0x9e, 0x71, 0x31, 0xa3, 0x29, 0x67, 0xfb,
	0x87, 0x3b, 0x74, 0x57, 0xc9, 0x78, 0x59, 0xba, 0x04, 0xc7, 0x67, 0x34, 0xcf, 0x7b, 0xfd, 0xaf,
	0xab, 0x1c, 0xe1, 0x37, 0xe8, 0xf5, 0xee, 0x54, 0xb9, 0x23, 0xd0, 0x34, 0x95, 0x67, 0xc0, 0x94,
	0x14, 0x3b, 0xf1, 0x0f, 0x4f, 0x7f, 0x0d, 0xd0, 0xf8, 0xdc, 0x2c, 0xc9, 0xdd, 0x05, 0x9d, 0x3e,
	0xdd, 0xeb, 0x20, 0x77, 0xc5, 0xe6, 0x83, 0xcf, 0x6c, 0xed, 0x6c, 0xcd, 0xc5, 0x62, 0xd5, 0x12,
	0xd3, 0xb5, 0x61, 0xab, 0x57, 0x7d, 0xed, 0xb7, 0x8f, 0xe4, 0xf2, 0x9b, 0xfd, 0xdf, 0x9b, 0x79,
	0xef, 0x3f, 0xdf, 0x83, 0x83, 0x09, 0xa5, 0x3f, 0x82, 0xd1, 0xc4, 0x2f, 0xa3, 0x8d, 0x25, 0x5e,
	0x3a, 0x35, 0x8b, 0x48, 0x1f, 0x69, 0x7f, 0xde, 0x02, 0x35, 0x6d, 0x6c, 0xbd, 0x01, 0xea, 0x59,
	0x54, 0x7b, 0xe0, 0xcb, 0x61, 0x1f, 0xfc, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x13,
	0xb5, 0x65, 0xab, 0x02, 0x00, 0x00,
}