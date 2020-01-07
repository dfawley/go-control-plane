// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/common/fault/v3alpha/fault.proto

package envoy_extensions_filters_common_fault_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type FaultDelay_FaultDelayType int32

const (
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}

var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}

func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3d725bf1f9b5b26, []int{0, 0}
}

type FaultDelay struct {
	HiddenEnvoyDeprecatedType FaultDelay_FaultDelayType `protobuf:"varint,1,opt,name=hidden_envoy_deprecated_type,json=hiddenEnvoyDeprecatedType,proto3,enum=envoy.extensions.filters.common.fault.v3alpha.FaultDelay_FaultDelayType" json:"hidden_envoy_deprecated_type,omitempty"` // Deprecated: Do not use.
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	//	*FaultDelay_HeaderDelay_
	FaultDelaySecifier   isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	Percentage           *v3alpha.FractionalPercent      `protobuf:"bytes,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FaultDelay) Reset()         { *m = FaultDelay{} }
func (m *FaultDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()    {}
func (*FaultDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d725bf1f9b5b26, []int{0}
}

func (m *FaultDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultDelay.Unmarshal(m, b)
}
func (m *FaultDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultDelay.Marshal(b, m, deterministic)
}
func (m *FaultDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay.Merge(m, src)
}
func (m *FaultDelay) XXX_Size() int {
	return xxx_messageInfo_FaultDelay.Size(m)
}
func (m *FaultDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *FaultDelay) GetHiddenEnvoyDeprecatedType() FaultDelay_FaultDelayType {
	if m != nil {
		return m.HiddenEnvoyDeprecatedType
	}
	return FaultDelay_FIXED
}

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
}

type FaultDelay_FixedDelay struct {
	FixedDelay *duration.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,proto3,oneof"`
}

type FaultDelay_HeaderDelay_ struct {
	HeaderDelay *FaultDelay_HeaderDelay `protobuf:"bytes,5,opt,name=header_delay,json=headerDelay,proto3,oneof"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (*FaultDelay_HeaderDelay_) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetFixedDelay() *duration.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *FaultDelay) GetHeaderDelay() *FaultDelay_HeaderDelay {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_HeaderDelay_); ok {
		return x.HeaderDelay
	}
	return nil
}

func (m *FaultDelay) GetPercentage() *v3alpha.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultDelay_FixedDelay)(nil),
		(*FaultDelay_HeaderDelay_)(nil),
	}
}

type FaultDelay_HeaderDelay struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultDelay_HeaderDelay) Reset()         { *m = FaultDelay_HeaderDelay{} }
func (m *FaultDelay_HeaderDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay_HeaderDelay) ProtoMessage()    {}
func (*FaultDelay_HeaderDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d725bf1f9b5b26, []int{0, 0}
}

func (m *FaultDelay_HeaderDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Unmarshal(m, b)
}
func (m *FaultDelay_HeaderDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Marshal(b, m, deterministic)
}
func (m *FaultDelay_HeaderDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay_HeaderDelay.Merge(m, src)
}
func (m *FaultDelay_HeaderDelay) XXX_Size() int {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Size(m)
}
func (m *FaultDelay_HeaderDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay_HeaderDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay_HeaderDelay proto.InternalMessageInfo

type FaultRateLimit struct {
	// Types that are valid to be assigned to LimitType:
	//	*FaultRateLimit_FixedLimit_
	//	*FaultRateLimit_HeaderLimit_
	LimitType            isFaultRateLimit_LimitType `protobuf_oneof:"limit_type"`
	Percentage           *v3alpha.FractionalPercent `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FaultRateLimit) Reset()         { *m = FaultRateLimit{} }
func (m *FaultRateLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit) ProtoMessage()    {}
func (*FaultRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d725bf1f9b5b26, []int{1}
}

func (m *FaultRateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit.Merge(m, src)
}
func (m *FaultRateLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit.Size(m)
}
func (m *FaultRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit proto.InternalMessageInfo

type isFaultRateLimit_LimitType interface {
	isFaultRateLimit_LimitType()
}

type FaultRateLimit_FixedLimit_ struct {
	FixedLimit *FaultRateLimit_FixedLimit `protobuf:"bytes,1,opt,name=fixed_limit,json=fixedLimit,proto3,oneof"`
}

type FaultRateLimit_HeaderLimit_ struct {
	HeaderLimit *FaultRateLimit_HeaderLimit `protobuf:"bytes,3,opt,name=header_limit,json=headerLimit,proto3,oneof"`
}

func (*FaultRateLimit_FixedLimit_) isFaultRateLimit_LimitType() {}

func (*FaultRateLimit_HeaderLimit_) isFaultRateLimit_LimitType() {}

func (m *FaultRateLimit) GetLimitType() isFaultRateLimit_LimitType {
	if m != nil {
		return m.LimitType
	}
	return nil
}

func (m *FaultRateLimit) GetFixedLimit() *FaultRateLimit_FixedLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_FixedLimit_); ok {
		return x.FixedLimit
	}
	return nil
}

func (m *FaultRateLimit) GetHeaderLimit() *FaultRateLimit_HeaderLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_HeaderLimit_); ok {
		return x.HeaderLimit
	}
	return nil
}

func (m *FaultRateLimit) GetPercentage() *v3alpha.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultRateLimit) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultRateLimit_FixedLimit_)(nil),
		(*FaultRateLimit_HeaderLimit_)(nil),
	}
}

type FaultRateLimit_FixedLimit struct {
	LimitKbps            uint64   `protobuf:"varint,1,opt,name=limit_kbps,json=limitKbps,proto3" json:"limit_kbps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_FixedLimit) Reset()         { *m = FaultRateLimit_FixedLimit{} }
func (m *FaultRateLimit_FixedLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_FixedLimit) ProtoMessage()    {}
func (*FaultRateLimit_FixedLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d725bf1f9b5b26, []int{1, 0}
}

func (m *FaultRateLimit_FixedLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit_FixedLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit_FixedLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_FixedLimit.Merge(m, src)
}
func (m *FaultRateLimit_FixedLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Size(m)
}
func (m *FaultRateLimit_FixedLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_FixedLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_FixedLimit proto.InternalMessageInfo

func (m *FaultRateLimit_FixedLimit) GetLimitKbps() uint64 {
	if m != nil {
		return m.LimitKbps
	}
	return 0
}

type FaultRateLimit_HeaderLimit struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_HeaderLimit) Reset()         { *m = FaultRateLimit_HeaderLimit{} }
func (m *FaultRateLimit_HeaderLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_HeaderLimit) ProtoMessage()    {}
func (*FaultRateLimit_HeaderLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3d725bf1f9b5b26, []int{1, 1}
}

func (m *FaultRateLimit_HeaderLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.Merge(m, src)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Size(m)
}
func (m *FaultRateLimit_HeaderLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_HeaderLimit proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.extensions.filters.common.fault.v3alpha.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
	proto.RegisterType((*FaultDelay)(nil), "envoy.extensions.filters.common.fault.v3alpha.FaultDelay")
	proto.RegisterType((*FaultDelay_HeaderDelay)(nil), "envoy.extensions.filters.common.fault.v3alpha.FaultDelay.HeaderDelay")
	proto.RegisterType((*FaultRateLimit)(nil), "envoy.extensions.filters.common.fault.v3alpha.FaultRateLimit")
	proto.RegisterType((*FaultRateLimit_FixedLimit)(nil), "envoy.extensions.filters.common.fault.v3alpha.FaultRateLimit.FixedLimit")
	proto.RegisterType((*FaultRateLimit_HeaderLimit)(nil), "envoy.extensions.filters.common.fault.v3alpha.FaultRateLimit.HeaderLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/common/fault/v3alpha/fault.proto", fileDescriptor_d3d725bf1f9b5b26)
}

var fileDescriptor_d3d725bf1f9b5b26 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdf, 0x8e, 0xd2, 0x4e,
	0x14, 0xa6, 0x6c, 0xf9, 0xfd, 0x76, 0x07, 0xb3, 0x59, 0x1b, 0x13, 0x59, 0xd6, 0x6c, 0x90, 0x44,
	0x45, 0xcd, 0x4e, 0x13, 0x88, 0x7f, 0x16, 0xa3, 0x17, 0x0d, 0x10, 0x70, 0x35, 0x21, 0x8d, 0x17,
	0xde, 0x35, 0x43, 0x7b, 0x0a, 0xe3, 0x96, 0x99, 0xa6, 0x2d, 0x04, 0xde, 0xc0, 0x1b, 0x5f, 0xc0,
	0x47, 0xf0, 0x25, 0xbc, 0xf2, 0x05, 0x7c, 0x1a, 0xb3, 0x57, 0x66, 0xfe, 0x14, 0xd8, 0x44, 0xa3,
	0xe0, 0xdd, 0x9c, 0x4e, 0xbf, 0xef, 0x3b, 0xe7, 0xfb, 0x4e, 0x06, 0x9d, 0x03, 0x9b, 0xf3, 0xa5,
	0x0d, 0x8b, 0x0c, 0x58, 0x4a, 0x39, 0x4b, 0xed, 0x90, 0x46, 0x19, 0x24, 0xa9, 0xed, 0xf3, 0xe9,
	0x94, 0x33, 0x3b, 0x24, 0xb3, 0x28, 0xb3, 0xe7, 0x2d, 0x12, 0xc5, 0x13, 0xa2, 0x2a, 0x1c, 0x27,
	0x3c, 0xe3, 0xd6, 0x99, 0x84, 0xe2, 0x35, 0x14, 0x6b, 0x28, 0x56, 0x50, 0xac, 0x7e, 0xd6, 0xd0,
	0x6a, 0x4d, 0x29, 0x65, 0xcb, 0x18, 0x56, 0x74, 0x31, 0x24, 0x3e, 0x30, 0x4d, 0x58, 0x3d, 0x1d,
	0x73, 0x3e, 0x8e, 0xc0, 0x96, 0xd5, 0x68, 0x16, 0xda, 0xc1, 0x2c, 0x21, 0x19, 0xe5, 0x4c, 0xdf,
	0xdf, 0x9d, 0x05, 0x31, 0xb1, 0x09, 0x63, 0x3c, 0x93, 0x9f, 0x53, 0x7b, 0x0e, 0x89, 0x50, 0xa6,
	0x6c, 0xac, 0x7f, 0xb9, 0x3d, 0x27, 0x11, 0x0d, 0x48, 0x06, 0x76, 0x7e, 0x50, 0x17, 0xf5, 0xef,
	0x26, 0x42, 0x3d, 0xd1, 0x4f, 0x07, 0x22, 0xb2, 0xb4, 0x3e, 0x19, 0xe8, 0xce, 0x84, 0x06, 0x01,
	0x30, 0x4f, 0xb6, 0xe5, 0x05, 0x10, 0x27, 0xe0, 0x93, 0x0c, 0x02, 0x4f, 0x74, 0x58, 0x31, 0x6a,
	0x46, 0xe3, 0xb0, 0xd9, 0xc7, 0x5b, 0xcd, 0x88, 0xd7, 0x0a, 0x1b, 0xc7, 0x77, 0xcb, 0x18, 0x9c,
	0x62, 0xc5, 0x70, 0x8f, 0x95, 0x62, 0x57, 0x50, 0x76, 0x56, 0x7a, 0xe2, 0xda, 0xea, 0xa1, 0x72,
	0x48, 0x17, 0x10, 0x78, 0x81, 0x40, 0x54, 0xf6, 0x6a, 0x46, 0xa3, 0xdc, 0x3c, 0xc6, 0xca, 0x10,
	0x9c, 0x1b, 0x82, 0x3b, 0xda, 0x10, 0x67, 0xff, 0xca, 0x29, 0x7d, 0x31, 0x8a, 0x8f, 0x0a, 0xfd,
	0x82, 0x8b, 0x24, 0x52, 0xcd, 0xf5, 0x01, 0xdd, 0x98, 0x00, 0x09, 0x20, 0xd1, 0x44, 0x25, 0x49,
	0xd4, 0xdd, 0x7d, 0x8c, 0xbe, 0x64, 0x93, 0xe7, 0x7e, 0xc1, 0x2d, 0x4f, 0xd6, 0xa5, 0xd5, 0x45,
	0x48, 0xe7, 0x47, 0xc6, 0x50, 0x31, 0xa5, 0xd2, 0x3d, 0xad, 0x24, 0x3c, 0x5c, 0xd3, 0x25, 0xc4,
	0x17, 0x5d, 0x93, 0x68, 0xa8, 0xfe, 0x77, 0x37, 0x80, 0xd5, 0x01, 0x2a, 0x6f, 0x88, 0xb4, 0xdb,
	0x9f, 0xbf, 0x7d, 0x3c, 0x7d, 0x82, 0x5a, 0x8a, 0xc7, 0xe7, 0x2c, 0xa4, 0x63, 0xdd, 0x6d, 0xde,
	0x65, 0xf3, 0x37, 0x0d, 0xd6, 0x4f, 0xd0, 0xe1, 0x75, 0xdb, 0xad, 0x03, 0x54, 0xea, 0x0d, 0xde,
	0x77, 0x3b, 0x47, 0x85, 0x36, 0x16, 0xc4, 0x0f, 0xd1, 0x83, 0xbf, 0x24, 0x76, 0x4e, 0xd0, 0x2d,
	0xf9, 0x59, 0x39, 0xe9, 0xa5, 0xe0, 0xd3, 0x90, 0x42, 0x62, 0xed, 0xfd, 0x70, 0x8c, 0xd7, 0xe6,
	0x7e, 0xf1, 0x68, 0xaf, 0xfe, 0xd5, 0xd4, 0x82, 0x2e, 0xc9, 0xe0, 0x0d, 0x9d, 0xd2, 0xcc, 0xba,
	0xcc, 0x83, 0x8c, 0x44, 0x29, 0xd7, 0xa8, 0xbc, 0xdb, 0x1a, 0xad, 0x38, 0x71, 0x4f, 0x10, 0xca,
	0xe3, 0x2a, 0x6d, 0x25, 0xc6, 0x56, 0x69, 0x2b, 0x35, 0xb5, 0x36, 0x83, 0x7f, 0x53, 0x53, 0x86,
	0xe6, 0x72, 0x3a, 0x71, 0xa5, 0x77, 0x3d, 0xf1, 0xe2, 0xae, 0x89, 0xa7, 0x08, 0xad, 0x47, 0xb2,
	0xee, 0x23, 0x24, 0xbb, 0xf7, 0x2e, 0x47, 0x71, 0x2a, 0x0d, 0x33, 0x9d, 0xff, 0xaf, 0x1c, 0xb3,
	0x59, 0x6c, 0x18, 0xee, 0x81, 0xbc, 0xba, 0x18, 0xc5, 0x69, 0xfb, 0xa5, 0xc8, 0xef, 0x39, 0x7a,
	0xfa, 0xe7, 0xfc, 0x7e, 0xe5, 0x5c, 0xf5, 0x6d, 0xbe, 0x66, 0xb2, 0x6c, 0xbf, 0x12, 0x6c, 0xe7,
	0xe8, 0xd9, 0x36, 0x6c, 0x9b, 0xf8, 0xa6, 0xc0, 0x9f, 0xa1, 0xc7, 0x5b, 0xe0, 0x9d, 0x9b, 0xf9,
	0xa4, 0xc2, 0x2b, 0xb9, 0x47, 0xce, 0x05, 0x7a, 0x41, 0xb9, 0x72, 0x30, 0x4e, 0xf8, 0x62, 0xb9,
	0x5d, 0x74, 0x8e, 0x7a, 0xd2, 0x86, 0xe2, 0x79, 0x18, 0x1a, 0xa3, 0xff, 0xe4, 0x3b, 0xd1, 0xfa,
	0x19, 0x00, 0x00, 0xff, 0xff, 0x36, 0x22, 0xf5, 0x54, 0xd4, 0x05, 0x00, 0x00,
}