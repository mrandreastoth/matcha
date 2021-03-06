// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/proto/layout/layout.proto

/*
Package layout is a generated protocol buffer package.

It is generated from these files:
	gomatcha.io/matcha/proto/layout/layout.proto

It has these top-level messages:
	Point
	Rect
	Insets
	Guide
*/
package layout

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

type Point struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y" json:"y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Point) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Rect struct {
	Min *Point `protobuf:"bytes,1,opt,name=min" json:"min,omitempty"`
	Max *Point `protobuf:"bytes,2,opt,name=max" json:"max,omitempty"`
}

func (m *Rect) Reset()                    { *m = Rect{} }
func (m *Rect) String() string            { return proto.CompactTextString(m) }
func (*Rect) ProtoMessage()               {}
func (*Rect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rect) GetMin() *Point {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Rect) GetMax() *Point {
	if m != nil {
		return m.Max
	}
	return nil
}

type Insets struct {
	Top    float64 `protobuf:"fixed64,1,opt,name=top" json:"top,omitempty"`
	Left   float64 `protobuf:"fixed64,2,opt,name=left" json:"left,omitempty"`
	Bottom float64 `protobuf:"fixed64,3,opt,name=bottom" json:"bottom,omitempty"`
	Right  float64 `protobuf:"fixed64,4,opt,name=right" json:"right,omitempty"`
}

func (m *Insets) Reset()                    { *m = Insets{} }
func (m *Insets) String() string            { return proto.CompactTextString(m) }
func (*Insets) ProtoMessage()               {}
func (*Insets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Insets) GetTop() float64 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Insets) GetLeft() float64 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Insets) GetBottom() float64 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Insets) GetRight() float64 {
	if m != nil {
		return m.Right
	}
	return 0
}

type Guide struct {
	Frame  *Rect `protobuf:"bytes,1,opt,name=frame" json:"frame,omitempty"`
	ZIndex int64 `protobuf:"varint,3,opt,name=zIndex" json:"zIndex,omitempty"`
}

func (m *Guide) Reset()                    { *m = Guide{} }
func (m *Guide) String() string            { return proto.CompactTextString(m) }
func (*Guide) ProtoMessage()               {}
func (*Guide) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Guide) GetFrame() *Rect {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *Guide) GetZIndex() int64 {
	if m != nil {
		return m.ZIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "matcha.layout.Point")
	proto.RegisterType((*Rect)(nil), "matcha.layout.Rect")
	proto.RegisterType((*Insets)(nil), "matcha.layout.Insets")
	proto.RegisterType((*Guide)(nil), "matcha.layout.Guide")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/proto/layout/layout.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xdb, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x73, 0x40, 0xc6, 0x03, 0xb2, 0x16, 0xc9, 0x9d, 0x12, 0x41, 0x14, 0x24, 0x01,
	0x7d, 0x83, 0x20, 0x48, 0x45, 0x21, 0xec, 0x85, 0x17, 0xe2, 0xcd, 0xa6, 0xdd, 0xb6, 0x0b, 0x4d,
	0xa6, 0xa4, 0x53, 0xd8, 0xf8, 0x38, 0x3e, 0xa9, 0xec, 0xec, 0x22, 0x28, 0x78, 0xb5, 0xff, 0xff,
	0xf3, 0xcd, 0x61, 0x07, 0xee, 0x56, 0xd8, 0x29, 0x9a, 0xaf, 0x55, 0x69, 0xb0, 0xf2, 0xaa, 0xda,
	0x0e, 0x48, 0x58, 0x6d, 0xd4, 0x88, 0x7b, 0x0a, 0x4f, 0xc9, 0x99, 0x38, 0x0e, 0xac, 0x0f, 0x8b,
	0x2b, 0x48, 0x1b, 0x34, 0x3d, 0x89, 0x23, 0x88, 0x6c, 0x1e, 0x5d, 0x46, 0x37, 0x91, 0x8c, 0xac,
	0x73, 0x63, 0x3e, 0xf1, 0x6e, 0x2c, 0xde, 0x20, 0x91, 0x7a, 0x4e, 0xe2, 0x1a, 0xe2, 0xce, 0xf4,
	0x4c, 0x1d, 0xde, 0x4f, 0xcb, 0x5f, 0x9d, 0x4a, 0x6e, 0x23, 0x1d, 0xc0, 0x9c, 0xb2, 0x5c, 0xff,
	0x3f, 0xa7, 0x6c, 0xf1, 0x01, 0xd9, 0xac, 0xdf, 0x69, 0xda, 0x89, 0x53, 0x88, 0x09, 0xb7, 0x61,
	0xbe, 0x93, 0x42, 0x40, 0xb2, 0xd1, 0x4b, 0x0a, 0x4b, 0xb0, 0x16, 0xe7, 0x90, 0xb5, 0x48, 0x84,
	0x5d, 0x1e, 0x73, 0x1a, 0x9c, 0x98, 0x42, 0x3a, 0x98, 0xd5, 0x9a, 0xf2, 0x84, 0x63, 0x6f, 0x8a,
	0x67, 0x48, 0x9f, 0xf6, 0x66, 0xa1, 0xc5, 0x2d, 0xa4, 0xcb, 0x41, 0x75, 0x3a, 0x2c, 0x7e, 0xf6,
	0x67, 0x21, 0xf7, 0x35, 0xe9, 0x09, 0x37, 0xe1, 0x73, 0xd6, 0x2f, 0xb4, 0xe5, 0x09, 0xb1, 0x0c,
	0xae, 0x7e, 0x84, 0x0b, 0x83, 0xe5, 0xcf, 0xa1, 0xc3, 0xc3, 0x17, 0x0d, 0x6d, 0xea, 0x83, 0xa6,
	0x7d, 0x61, 0xf5, 0x9e, 0xf9, 0xe4, 0x6b, 0x72, 0xf2, 0xca, 0x9c, 0x8f, 0x9b, 0xba, 0xcd, 0xb8,
	0xe0, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x03, 0x23, 0x11, 0xb2, 0x01, 0x00, 0x00,
}
