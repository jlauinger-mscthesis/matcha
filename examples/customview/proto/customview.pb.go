// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/examples/customview/proto/customview.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	gomatcha.io/matcha/examples/customview/proto/customview.proto

It has these top-level messages:
	View
	Event
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type View struct {
	FrontCamera bool `protobuf:"varint,1,opt,name=frontCamera" json:"frontCamera,omitempty"`
}

func (m *View) Reset()                    { *m = View{} }
func (m *View) String() string            { return proto1.CompactTextString(m) }
func (*View) ProtoMessage()               {}
func (*View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *View) GetFrontCamera() bool {
	if m != nil {
		return m.FrontCamera
	}
	return false
}

type Event struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto1.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto1.RegisterType((*View)(nil), "matcha.examples.customview.View")
	proto1.RegisterType((*Event)(nil), "matcha.examples.customview.Event")
}

func init() {
	proto1.RegisterFile("gomatcha.io/matcha/examples/customview/proto/customview.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x53, 0x2b, 0x12, 0x73, 0x0b,
	0x72, 0x52, 0x8b, 0xf5, 0x93, 0x4b, 0x8b, 0x4b, 0xf2, 0x73, 0xcb, 0x32, 0x53, 0xcb, 0xf5, 0x0b,
	0x8a, 0xf2, 0x4b, 0xf2, 0x91, 0x04, 0xf4, 0xc0, 0x02, 0x42, 0x52, 0x50, 0xcd, 0x30, 0x2d, 0x7a,
	0x08, 0x15, 0x4a, 0x1a, 0x5c, 0x2c, 0x61, 0x99, 0xa9, 0xe5, 0x42, 0x0a, 0x5c, 0xdc, 0x69, 0x45,
	0xf9, 0x79, 0x25, 0xce, 0x89, 0xb9, 0xa9, 0x45, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41,
	0xc8, 0x42, 0x4a, 0xb2, 0x5c, 0xac, 0xae, 0x65, 0xa9, 0x79, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x65,
	0x89, 0x39, 0xa5, 0xa9, 0x50, 0x45, 0x10, 0x8e, 0x93, 0x07, 0x97, 0x4c, 0x66, 0xbe, 0x1e, 0xdc,
	0xa1, 0xe8, 0x8e, 0x70, 0xe2, 0x77, 0x06, 0x8b, 0x80, 0x2c, 0x0b, 0x00, 0x09, 0x44, 0xb1, 0x82,
	0xc5, 0x17, 0x31, 0xa1, 0x4b, 0x24, 0xb1, 0x81, 0x25, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0xb4, 0x16, 0xfb, 0xf6, 0x00, 0x00, 0x00,
}
