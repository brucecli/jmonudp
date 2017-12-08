// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry_top.proto

/*
Package telemetry_top is a generated protocol buffer package.

It is generated from these files:
	telemetry_top.proto

It has these top-level messages:
	TelemetryFieldOptions
	TelemetryStream
	IETFSensors
	EnterpriseSensors
	JuniperNetworksSensors
*/
package telemetry_top

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

//import _ "."
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TelemetryFieldOptions struct {
	IsKey            *bool  `protobuf:"varint,1,opt,name=is_key,json=isKey" json:"is_key,omitempty"`
	IsTimestamp      *bool  `protobuf:"varint,2,opt,name=is_timestamp,json=isTimestamp" json:"is_timestamp,omitempty"`
	IsCounter        *bool  `protobuf:"varint,3,opt,name=is_counter,json=isCounter" json:"is_counter,omitempty"`
	IsGauge          *bool  `protobuf:"varint,4,opt,name=is_gauge,json=isGauge" json:"is_gauge,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TelemetryFieldOptions) Reset()                    { *m = TelemetryFieldOptions{} }
func (m *TelemetryFieldOptions) String() string            { return proto.CompactTextString(m) }
func (*TelemetryFieldOptions) ProtoMessage()               {}
func (*TelemetryFieldOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TelemetryFieldOptions) GetIsKey() bool {
	if m != nil && m.IsKey != nil {
		return *m.IsKey
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsTimestamp() bool {
	if m != nil && m.IsTimestamp != nil {
		return *m.IsTimestamp
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsCounter() bool {
	if m != nil && m.IsCounter != nil {
		return *m.IsCounter
	}
	return false
}

func (m *TelemetryFieldOptions) GetIsGauge() bool {
	if m != nil && m.IsGauge != nil {
		return *m.IsGauge
	}
	return false
}

type TelemetryStream struct {
	// router name or export IP address
	SystemId *string `protobuf:"bytes,1,req,name=system_id,json=systemId" json:"system_id,omitempty"`
	// line card / RE (slot number)
	ComponentId *uint32 `protobuf:"varint,2,opt,name=component_id,json=componentId" json:"component_id,omitempty"`
	// PFE (if applicable)
	SubComponentId *uint32 `protobuf:"varint,3,opt,name=sub_component_id,json=subComponentId" json:"sub_component_id,omitempty"`
	// configured sensor name
	SensorName *string `protobuf:"bytes,4,opt,name=sensor_name,json=sensorName" json:"sensor_name,omitempty"`
	// sequence number, monotonically increasesing for each
	// system_id, component_id, sub_component_id + sensor_name.
	SequenceNumber *uint32 `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// timestamp (milliseconds since 00:00:00 UTC 1/1/1970)
	Timestamp *uint64 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	// major version
	VersionMajor *uint32 `protobuf:"varint,7,opt,name=version_major,json=versionMajor" json:"version_major,omitempty"`
	// minor version
	VersionMinor     *uint32            `protobuf:"varint,8,opt,name=version_minor,json=versionMinor" json:"version_minor,omitempty"`
	Ietf             *IETFSensors       `protobuf:"bytes,100,opt,name=ietf" json:"ietf,omitempty"`
	Enterprise       *EnterpriseSensors `protobuf:"bytes,101,opt,name=enterprise" json:"enterprise,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *TelemetryStream) Reset()                    { *m = TelemetryStream{} }
func (m *TelemetryStream) String() string            { return proto.CompactTextString(m) }
func (*TelemetryStream) ProtoMessage()               {}
func (*TelemetryStream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TelemetryStream) GetSystemId() string {
	if m != nil && m.SystemId != nil {
		return *m.SystemId
	}
	return ""
}

func (m *TelemetryStream) GetComponentId() uint32 {
	if m != nil && m.ComponentId != nil {
		return *m.ComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSubComponentId() uint32 {
	if m != nil && m.SubComponentId != nil {
		return *m.SubComponentId
	}
	return 0
}

func (m *TelemetryStream) GetSensorName() string {
	if m != nil && m.SensorName != nil {
		return *m.SensorName
	}
	return ""
}

func (m *TelemetryStream) GetSequenceNumber() uint32 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

func (m *TelemetryStream) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *TelemetryStream) GetVersionMajor() uint32 {
	if m != nil && m.VersionMajor != nil {
		return *m.VersionMajor
	}
	return 0
}

func (m *TelemetryStream) GetVersionMinor() uint32 {
	if m != nil && m.VersionMinor != nil {
		return *m.VersionMinor
	}
	return 0
}

func (m *TelemetryStream) GetIetf() *IETFSensors {
	if m != nil {
		return m.Ietf
	}
	return nil
}

func (m *TelemetryStream) GetEnterprise() *EnterpriseSensors {
	if m != nil {
		return m.Enterprise
	}
	return nil
}

type IETFSensors struct {
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *IETFSensors) Reset()                    { *m = IETFSensors{} }
func (m *IETFSensors) String() string            { return proto.CompactTextString(m) }
func (*IETFSensors) ProtoMessage()               {}
func (*IETFSensors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

var extRange_IETFSensors = []proto.ExtensionRange{
	{1, 536870911},
}

func (*IETFSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_IETFSensors
}

type EnterpriseSensors struct {
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *EnterpriseSensors) Reset()                    { *m = EnterpriseSensors{} }
func (m *EnterpriseSensors) String() string            { return proto.CompactTextString(m) }
func (*EnterpriseSensors) ProtoMessage()               {}
func (*EnterpriseSensors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

var extRange_EnterpriseSensors = []proto.ExtensionRange{
	{1, 536870911},
}

func (*EnterpriseSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EnterpriseSensors
}

type JuniperNetworksSensors struct {
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *JuniperNetworksSensors) Reset()                    { *m = JuniperNetworksSensors{} }
func (m *JuniperNetworksSensors) String() string            { return proto.CompactTextString(m) }
func (*JuniperNetworksSensors) ProtoMessage()               {}
func (*JuniperNetworksSensors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

var extRange_JuniperNetworksSensors = []proto.ExtensionRange{
	{1, 536870911},
}

func (*JuniperNetworksSensors) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_JuniperNetworksSensors
}

var E_TelemetryOptions = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*TelemetryFieldOptions)(nil),
	Field:         1024,
	Name:          "telemetry_options",
	Tag:           "bytes,1024,opt,name=telemetry_options,json=telemetryOptions",
	Filename:      "telemetry_top.proto",
}

var E_JuniperNetworks = &proto.ExtensionDesc{
	ExtendedType:  (*EnterpriseSensors)(nil),
	ExtensionType: (*JuniperNetworksSensors)(nil),
	Field:         2636,
	Name:          "juniperNetworks",
	Tag:           "bytes,2636,opt,name=juniperNetworks",
	Filename:      "telemetry_top.proto",
}

func init() {
	proto.RegisterType((*TelemetryFieldOptions)(nil), "TelemetryFieldOptions")
	proto.RegisterType((*TelemetryStream)(nil), "TelemetryStream")
	proto.RegisterType((*IETFSensors)(nil), "IETFSensors")
	proto.RegisterType((*EnterpriseSensors)(nil), "EnterpriseSensors")
	proto.RegisterType((*JuniperNetworksSensors)(nil), "JuniperNetworksSensors")
	proto.RegisterExtension(E_TelemetryOptions)
	proto.RegisterExtension(E_JuniperNetworks)
}

func init() { proto.RegisterFile("telemetry_top.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xda, 0x4c,
	0x14, 0xc5, 0x65, 0x42, 0x12, 0x73, 0x21, 0x81, 0xcc, 0xf7, 0x25, 0x99, 0x56, 0x42, 0x72, 0xc9,
	0x22, 0xa8, 0x95, 0x8c, 0x94, 0x65, 0x36, 0xad, 0x12, 0x25, 0x15, 0xad, 0x4a, 0x25, 0x07, 0x75,
	0x6b, 0x19, 0x7c, 0x83, 0x86, 0xe0, 0x19, 0x77, 0xee, 0xb8, 0x15, 0x3b, 0xd4, 0x6d, 0x5f, 0xa6,
	0x0f, 0xd2, 0xa7, 0xe8, 0x93, 0x54, 0x1e, 0xf3, 0xc7, 0x34, 0xec, 0x98, 0x73, 0x7e, 0x73, 0x38,
	0xb6, 0x0f, 0xfc, 0x67, 0x70, 0x86, 0x09, 0x1a, 0x3d, 0x0f, 0x8d, 0x4a, 0xfd, 0x54, 0x2b, 0xa3,
	0x5e, 0xd6, 0xd2, 0xd1, 0x74, 0xf9, 0xd3, 0x9b, 0x28, 0x35, 0x99, 0x61, 0xcf, 0x9e, 0x46, 0xd9,
	0x63, 0x2f, 0x46, 0x1a, 0x6b, 0x91, 0x1a, 0xa5, 0x0b, 0xa2, 0xf3, 0xd3, 0x81, 0xd3, 0xe1, 0x2a,
	0xe4, 0x5e, 0xe0, 0x2c, 0xfe, 0x9c, 0x1a, 0xa1, 0x24, 0xb1, 0x53, 0x38, 0x10, 0x14, 0x3e, 0xe1,
	0x9c, 0x3b, 0x9e, 0xd3, 0x75, 0x83, 0x7d, 0x41, 0x1f, 0x71, 0xce, 0x5e, 0x41, 0x43, 0x50, 0x68,
	0x44, 0x82, 0x64, 0xa2, 0x24, 0xe5, 0x15, 0x6b, 0xd6, 0x05, 0x0d, 0x57, 0x12, 0x6b, 0x03, 0x08,
	0x0a, 0xc7, 0x2a, 0x93, 0x06, 0x35, 0xdf, 0xb3, 0x40, 0x4d, 0xd0, 0x6d, 0x21, 0xb0, 0x17, 0xe0,
	0x0a, 0x0a, 0x27, 0x51, 0x36, 0x41, 0x5e, 0xb5, 0xe6, 0xa1, 0xa0, 0xf7, 0xf9, 0xb1, 0xf3, 0x6b,
	0x0f, 0x9a, 0xeb, 0x36, 0x0f, 0x46, 0x63, 0x94, 0xb0, 0x4b, 0xa8, 0xd1, 0x9c, 0x0c, 0x26, 0xa1,
	0x88, 0xb9, 0xe3, 0x55, 0xba, 0xb5, 0x1b, 0xf8, 0xf3, 0xb6, 0xc2, 0xab, 0x3f, 0xde, 0x55, 0x5c,
	0x27, 0x70, 0x0b, 0xb3, 0x1f, 0xb3, 0x2e, 0x34, 0xc6, 0x2a, 0x49, 0x95, 0x44, 0x69, 0x72, 0x36,
	0x6f, 0x76, 0x74, 0xb3, 0x5f, 0x60, 0xf5, 0xb5, 0xd5, 0x8f, 0x59, 0x0f, 0x5a, 0x94, 0x8d, 0xc2,
	0x2d, 0x7a, 0xaf, 0x4c, 0x1f, 0x53, 0x36, 0xba, 0x2d, 0x5d, 0x78, 0x03, 0x75, 0x42, 0x49, 0x4a,
	0x87, 0x32, 0x4a, 0x8a, 0xd6, 0xdb, 0x2d, 0xa0, 0xb0, 0x07, 0x51, 0x82, 0xec, 0x12, 0x9a, 0x84,
	0x5f, 0x33, 0x94, 0x63, 0x0c, 0x65, 0x96, 0x8c, 0x50, 0xf3, 0xfd, 0x3c, 0x3c, 0x38, 0x5e, 0xc9,
	0x03, 0xab, 0xb2, 0x0b, 0xa8, 0x6d, 0xde, 0xe3, 0x81, 0xe7, 0x74, 0xab, 0xf6, 0xff, 0x5b, 0x4e,
	0xb0, 0xd1, 0xd9, 0x05, 0x1c, 0x7d, 0x43, 0x4d, 0x42, 0xc9, 0x30, 0x89, 0xa6, 0x4a, 0xf3, 0x43,
	0x9b, 0xd5, 0x58, 0x8a, 0x9f, 0x72, 0x6d, 0x0b, 0x12, 0x52, 0x69, 0xee, 0x6e, 0x43, 0xb9, 0xc6,
	0x3c, 0xa8, 0x0a, 0x34, 0x8f, 0x3c, 0xf6, 0x9c, 0x6e, 0xfd, 0xaa, 0xe1, 0xf7, 0xef, 0x86, 0xf7,
	0x0f, 0xb6, 0x36, 0x05, 0xd6, 0x61, 0x57, 0x00, 0x98, 0x7f, 0xa2, 0x54, 0x0b, 0x42, 0x8e, 0x96,
	0x63, 0xfe, 0xdd, 0x5a, 0x5a, 0xd1, 0x25, 0xaa, 0x73, 0x0e, 0xf5, 0x52, 0xd0, 0x6b, 0xd7, 0x75,
	0x5a, 0x8b, 0xc5, 0x62, 0x51, 0xe9, 0xb4, 0xe1, 0xe4, 0xd9, 0xcd, 0x92, 0xdd, 0x81, 0xb3, 0x0f,
	0x99, 0x14, 0x29, 0xea, 0x01, 0x9a, 0xef, 0x4a, 0x3f, 0xd1, 0x33, 0xe6, 0x7a, 0x0c, 0x27, 0x9b,
	0x81, 0xab, 0xe5, 0x2e, 0xdb, 0x7e, 0x31, 0x6a, 0x7f, 0x35, 0x6a, 0xbf, 0x3c, 0x5b, 0xbe, 0x70,
	0x6d, 0xed, 0x33, 0x7f, 0xe7, 0xaa, 0x83, 0xd6, 0x3a, 0x70, 0xa9, 0x5c, 0x7f, 0x81, 0xe6, 0x74,
	0xbb, 0x08, 0xdb, 0xf1, 0xcc, 0xfc, 0xf7, 0xff, 0x36, 0xf7, 0xdc, 0xdf, 0xdd, 0x3a, 0xf8, 0x37,
	0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x47, 0x6a, 0xc1, 0x9c, 0x03, 0x00, 0x00,
}
