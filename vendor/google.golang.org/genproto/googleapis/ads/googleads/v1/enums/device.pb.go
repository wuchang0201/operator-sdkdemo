// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/device.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enumerates Google Ads devices available for targeting.
type DeviceEnum_Device int32

const (
	// Not specified.
	DeviceEnum_UNSPECIFIED DeviceEnum_Device = 0
	// The value is unknown in this version.
	DeviceEnum_UNKNOWN DeviceEnum_Device = 1
	// Mobile devices with full browsers.
	DeviceEnum_MOBILE DeviceEnum_Device = 2
	// Tablets with full browsers.
	DeviceEnum_TABLET DeviceEnum_Device = 3
	// Computers.
	DeviceEnum_DESKTOP DeviceEnum_Device = 4
	// Other device types.
	DeviceEnum_OTHER DeviceEnum_Device = 5
)

var DeviceEnum_Device_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MOBILE",
	3: "TABLET",
	4: "DESKTOP",
	5: "OTHER",
}
var DeviceEnum_Device_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MOBILE":      2,
	"TABLET":      3,
	"DESKTOP":     4,
	"OTHER":       5,
}

func (x DeviceEnum_Device) String() string {
	return proto.EnumName(DeviceEnum_Device_name, int32(x))
}
func (DeviceEnum_Device) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_device_0774693a7118df63, []int{0, 0}
}

// Container for enumeration of Google Ads devices available for targeting.
type DeviceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceEnum) Reset()         { *m = DeviceEnum{} }
func (m *DeviceEnum) String() string { return proto.CompactTextString(m) }
func (*DeviceEnum) ProtoMessage()    {}
func (*DeviceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_0774693a7118df63, []int{0}
}
func (m *DeviceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceEnum.Unmarshal(m, b)
}
func (m *DeviceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceEnum.Marshal(b, m, deterministic)
}
func (dst *DeviceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceEnum.Merge(dst, src)
}
func (m *DeviceEnum) XXX_Size() int {
	return xxx_messageInfo_DeviceEnum.Size(m)
}
func (m *DeviceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeviceEnum)(nil), "google.ads.googleads.v1.enums.DeviceEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.DeviceEnum_Device", DeviceEnum_Device_name, DeviceEnum_Device_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/device.proto", fileDescriptor_device_0774693a7118df63)
}

var fileDescriptor_device_0774693a7118df63 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x1c, 0xc5, 0xbf, 0xa6, 0x5f, 0x8b, 0x70, 0x07, 0xac, 0x8c, 0x88, 0x0e, 0xed, 0xc8, 0x60, 0x2b,
	0x62, 0x33, 0x93, 0x43, 0x4d, 0xa9, 0x5a, 0x92, 0x88, 0xa6, 0x41, 0x42, 0x59, 0x42, 0x1d, 0x59,
	0x91, 0x5a, 0xbb, 0xaa, 0xd3, 0x1c, 0x88, 0x91, 0xa3, 0x70, 0x0a, 0x66, 0x4e, 0x81, 0x6c, 0x93,
	0x6c, 0xb0, 0x58, 0xcf, 0x7e, 0x3f, 0x3f, 0x3f, 0xff, 0xc1, 0xb5, 0x50, 0x4a, 0xec, 0x4a, 0x5c,
	0x70, 0x8d, 0x9d, 0x34, 0xaa, 0x09, 0x70, 0x29, 0x4f, 0x7b, 0x8d, 0x79, 0xd9, 0x54, 0xdb, 0x12,
	0x1d, 0x8e, 0xaa, 0x56, 0xfe, 0xd8, 0x01, 0xa8, 0xe0, 0x1a, 0x75, 0x2c, 0x6a, 0x02, 0x64, 0xd9,
	0xcb, 0xab, 0x36, 0xea, 0x50, 0xe1, 0x42, 0x4a, 0x55, 0x17, 0x75, 0xa5, 0xa4, 0x76, 0x97, 0xa7,
	0x1c, 0x80, 0x99, 0x0d, 0x63, 0xf2, 0xb4, 0x9f, 0x66, 0x60, 0xe8, 0x76, 0xfe, 0x05, 0x18, 0x6d,
	0xa2, 0x75, 0xc2, 0xee, 0x16, 0xf7, 0x0b, 0x36, 0x83, 0xff, 0xfc, 0x11, 0x38, 0xdb, 0x44, 0xcb,
	0x28, 0x7e, 0x8e, 0x60, 0xcf, 0x07, 0x60, 0xf8, 0x18, 0x87, 0x8b, 0x15, 0x83, 0x9e, 0xd1, 0x29,
	0x0d, 0x57, 0x2c, 0x85, 0x7d, 0x03, 0xcd, 0xd8, 0x7a, 0x99, 0xc6, 0x09, 0xfc, 0xef, 0x9f, 0x83,
	0x41, 0x9c, 0x3e, 0xb0, 0x27, 0x38, 0x08, 0x3f, 0x7b, 0x60, 0xb2, 0x55, 0x7b, 0xf4, 0x67, 0xd3,
	0x70, 0xe4, 0xde, 0x4e, 0x4c, 0xb1, 0xa4, 0xf7, 0x12, 0xfe, 0xd0, 0x42, 0xed, 0x0a, 0x29, 0x90,
	0x3a, 0x0a, 0x2c, 0x4a, 0x69, 0x6b, 0xb7, 0x33, 0x39, 0x54, 0xfa, 0x97, 0x11, 0xdd, 0xda, 0xf5,
	0xcd, 0xeb, 0xcf, 0x29, 0x7d, 0xf7, 0xc6, 0x73, 0x17, 0x45, 0xb9, 0x46, 0x4e, 0x1a, 0x95, 0x05,
	0xc8, 0xfc, 0x5a, 0x7f, 0xb4, 0x7e, 0x4e, 0xb9, 0xce, 0x3b, 0x3f, 0xcf, 0x82, 0xdc, 0xfa, 0x5f,
	0xde, 0xc4, 0x1d, 0x12, 0x42, 0xb9, 0x26, 0xa4, 0x23, 0x08, 0xc9, 0x02, 0x42, 0x2c, 0xf3, 0x3a,
	0xb4, 0xc5, 0x6e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0x5e, 0x04, 0x03, 0xba, 0x01, 0x00,
	0x00,
}
