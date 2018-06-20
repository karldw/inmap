// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eioserve.proto

/*
Package eiopb is a generated protocol buffer package.

It is generated from these files:
	eioserve.proto

It has these top-level messages:
	Selectors
	Selection
	Year
	Point
	Rectangle
	ColorInfo
*/
package eiopb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Selectors struct {
	Codes  []string  `protobuf:"bytes,1,rep,name=Codes" json:"Codes,omitempty"`
	Names  []string  `protobuf:"bytes,2,rep,name=Names" json:"Names,omitempty"`
	Values []float32 `protobuf:"fixed32,3,rep,packed,name=Values" json:"Values,omitempty"`
}

func (m *Selectors) Reset()                    { *m = Selectors{} }
func (m *Selectors) String() string            { return proto.CompactTextString(m) }
func (*Selectors) ProtoMessage()               {}
func (*Selectors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Selectors) GetCodes() []string {
	if m != nil {
		return m.Codes
	}
	return nil
}

func (m *Selectors) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Selectors) GetValues() []float32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Selection struct {
	DemandGroup      string `protobuf:"bytes,1,opt,name=DemandGroup" json:"DemandGroup,omitempty"`
	DemandSector     string `protobuf:"bytes,2,opt,name=DemandSector" json:"DemandSector,omitempty"`
	ProductionGroup  string `protobuf:"bytes,3,opt,name=ProductionGroup" json:"ProductionGroup,omitempty"`
	ProductionSector string `protobuf:"bytes,4,opt,name=ProductionSector" json:"ProductionSector,omitempty"`
	ImpactType       string `protobuf:"bytes,5,opt,name=ImpactType" json:"ImpactType,omitempty"`
	DemandType       string `protobuf:"bytes,6,opt,name=DemandType" json:"DemandType,omitempty"`
	Year             int32  `protobuf:"varint,7,opt,name=Year" json:"Year,omitempty"`
	Population       string `protobuf:"bytes,8,opt,name=Population" json:"Population,omitempty"`
	Pollutant        int32  `protobuf:"varint,9,opt,name=Pollutant" json:"Pollutant,omitempty"`
}

func (m *Selection) Reset()                    { *m = Selection{} }
func (m *Selection) String() string            { return proto.CompactTextString(m) }
func (*Selection) ProtoMessage()               {}
func (*Selection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Selection) GetDemandGroup() string {
	if m != nil {
		return m.DemandGroup
	}
	return ""
}

func (m *Selection) GetDemandSector() string {
	if m != nil {
		return m.DemandSector
	}
	return ""
}

func (m *Selection) GetProductionGroup() string {
	if m != nil {
		return m.ProductionGroup
	}
	return ""
}

func (m *Selection) GetProductionSector() string {
	if m != nil {
		return m.ProductionSector
	}
	return ""
}

func (m *Selection) GetImpactType() string {
	if m != nil {
		return m.ImpactType
	}
	return ""
}

func (m *Selection) GetDemandType() string {
	if m != nil {
		return m.DemandType
	}
	return ""
}

func (m *Selection) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Selection) GetPopulation() string {
	if m != nil {
		return m.Population
	}
	return ""
}

func (m *Selection) GetPollutant() int32 {
	if m != nil {
		return m.Pollutant
	}
	return 0
}

type Year struct {
	Years []int32 `protobuf:"varint,1,rep,packed,name=Years" json:"Years,omitempty"`
}

func (m *Year) Reset()                    { *m = Year{} }
func (m *Year) String() string            { return proto.CompactTextString(m) }
func (*Year) ProtoMessage()               {}
func (*Year) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Year) GetYears() []int32 {
	if m != nil {
		return m.Years
	}
	return nil
}

type Point struct {
	X float32 `protobuf:"fixed32,1,opt,name=X" json:"X,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=Y" json:"Y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Point) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Rectangle struct {
	LL *Point `protobuf:"bytes,1,opt,name=LL" json:"LL,omitempty"`
	LR *Point `protobuf:"bytes,2,opt,name=LR" json:"LR,omitempty"`
	UR *Point `protobuf:"bytes,3,opt,name=UR" json:"UR,omitempty"`
	UL *Point `protobuf:"bytes,4,opt,name=UL" json:"UL,omitempty"`
}

func (m *Rectangle) Reset()                    { *m = Rectangle{} }
func (m *Rectangle) String() string            { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()               {}
func (*Rectangle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Rectangle) GetLL() *Point {
	if m != nil {
		return m.LL
	}
	return nil
}

func (m *Rectangle) GetLR() *Point {
	if m != nil {
		return m.LR
	}
	return nil
}

func (m *Rectangle) GetUR() *Point {
	if m != nil {
		return m.UR
	}
	return nil
}

func (m *Rectangle) GetUL() *Point {
	if m != nil {
		return m.UL
	}
	return nil
}

type ColorInfo struct {
	RGB    [][]byte `protobuf:"bytes,1,rep,name=RGB,proto3" json:"RGB,omitempty"`
	Legend string   `protobuf:"bytes,2,opt,name=Legend" json:"Legend,omitempty"`
}

func (m *ColorInfo) Reset()                    { *m = ColorInfo{} }
func (m *ColorInfo) String() string            { return proto.CompactTextString(m) }
func (*ColorInfo) ProtoMessage()               {}
func (*ColorInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ColorInfo) GetRGB() [][]byte {
	if m != nil {
		return m.RGB
	}
	return nil
}

func (m *ColorInfo) GetLegend() string {
	if m != nil {
		return m.Legend
	}
	return ""
}

func init() {
	proto.RegisterType((*Selectors)(nil), "eiopb.Selectors")
	proto.RegisterType((*Selection)(nil), "eiopb.Selection")
	proto.RegisterType((*Year)(nil), "eiopb.Year")
	proto.RegisterType((*Point)(nil), "eiopb.Point")
	proto.RegisterType((*Rectangle)(nil), "eiopb.Rectangle")
	proto.RegisterType((*ColorInfo)(nil), "eiopb.ColorInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EIOServe service

type EIOServeClient interface {
	DemandGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	DemandSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	ProdGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	ProdSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	Years(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Year, error)
	DefaultSelection(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selection, error)
	Populations(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error)
	MapInfo(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*ColorInfo, error)
	Geometry(ctx context.Context, in *Selection, opts ...grpc.CallOption) (EIOServe_GeometryClient, error)
}

type eIOServeClient struct {
	cc *grpc.ClientConn
}

func NewEIOServeClient(cc *grpc.ClientConn) EIOServeClient {
	return &eIOServeClient{cc}
}

func (c *eIOServeClient) DemandGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/DemandGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) DemandSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/DemandSectors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) ProdGroups(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/ProdGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) ProdSectors(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/ProdSectors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) Years(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Year, error) {
	out := new(Year)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/Years", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) DefaultSelection(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selection, error) {
	out := new(Selection)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/DefaultSelection", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) Populations(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*Selectors, error) {
	out := new(Selectors)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/Populations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) MapInfo(ctx context.Context, in *Selection, opts ...grpc.CallOption) (*ColorInfo, error) {
	out := new(ColorInfo)
	err := grpc.Invoke(ctx, "/eiopb.EIOServe/MapInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eIOServeClient) Geometry(ctx context.Context, in *Selection, opts ...grpc.CallOption) (EIOServe_GeometryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EIOServe_serviceDesc.Streams[0], c.cc, "/eiopb.EIOServe/Geometry", opts...)
	if err != nil {
		return nil, err
	}
	x := &eIOServeGeometryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EIOServe_GeometryClient interface {
	Recv() (*Rectangle, error)
	grpc.ClientStream
}

type eIOServeGeometryClient struct {
	grpc.ClientStream
}

func (x *eIOServeGeometryClient) Recv() (*Rectangle, error) {
	m := new(Rectangle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EIOServe service

type EIOServeServer interface {
	DemandGroups(context.Context, *Selection) (*Selectors, error)
	DemandSectors(context.Context, *Selection) (*Selectors, error)
	ProdGroups(context.Context, *Selection) (*Selectors, error)
	ProdSectors(context.Context, *Selection) (*Selectors, error)
	Years(context.Context, *Selection) (*Year, error)
	DefaultSelection(context.Context, *Selection) (*Selection, error)
	Populations(context.Context, *Selection) (*Selectors, error)
	MapInfo(context.Context, *Selection) (*ColorInfo, error)
	Geometry(*Selection, EIOServe_GeometryServer) error
}

func RegisterEIOServeServer(s *grpc.Server, srv EIOServeServer) {
	s.RegisterService(&_EIOServe_serviceDesc, srv)
}

func _EIOServe_DemandGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).DemandGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/DemandGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).DemandGroups(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_DemandSectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).DemandSectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/DemandSectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).DemandSectors(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_ProdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).ProdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/ProdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).ProdGroups(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_ProdSectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).ProdSectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/ProdSectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).ProdSectors(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_Years_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).Years(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/Years",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).Years(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_DefaultSelection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).DefaultSelection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/DefaultSelection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).DefaultSelection(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_Populations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).Populations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/Populations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).Populations(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_MapInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Selection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EIOServeServer).MapInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eiopb.EIOServe/MapInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EIOServeServer).MapInfo(ctx, req.(*Selection))
	}
	return interceptor(ctx, in, info, handler)
}

func _EIOServe_Geometry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Selection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EIOServeServer).Geometry(m, &eIOServeGeometryServer{stream})
}

type EIOServe_GeometryServer interface {
	Send(*Rectangle) error
	grpc.ServerStream
}

type eIOServeGeometryServer struct {
	grpc.ServerStream
}

func (x *eIOServeGeometryServer) Send(m *Rectangle) error {
	return x.ServerStream.SendMsg(m)
}

var _EIOServe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eiopb.EIOServe",
	HandlerType: (*EIOServeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DemandGroups",
			Handler:    _EIOServe_DemandGroups_Handler,
		},
		{
			MethodName: "DemandSectors",
			Handler:    _EIOServe_DemandSectors_Handler,
		},
		{
			MethodName: "ProdGroups",
			Handler:    _EIOServe_ProdGroups_Handler,
		},
		{
			MethodName: "ProdSectors",
			Handler:    _EIOServe_ProdSectors_Handler,
		},
		{
			MethodName: "Years",
			Handler:    _EIOServe_Years_Handler,
		},
		{
			MethodName: "DefaultSelection",
			Handler:    _EIOServe_DefaultSelection_Handler,
		},
		{
			MethodName: "Populations",
			Handler:    _EIOServe_Populations_Handler,
		},
		{
			MethodName: "MapInfo",
			Handler:    _EIOServe_MapInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Geometry",
			Handler:       _EIOServe_Geometry_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eioserve.proto",
}

func init() { proto.RegisterFile("eioserve.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xed, 0x3a, 0x8d, 0x27, 0x01, 0xa2, 0x15, 0x42, 0x16, 0x8a, 0x90, 0xb5, 0x5c, 0x2c,
	0x0e, 0x01, 0xa5, 0x54, 0xe2, 0x4c, 0x8b, 0xa2, 0x48, 0x86, 0x46, 0x1b, 0x8a, 0x9a, 0xe3, 0x36,
	0x99, 0x56, 0x96, 0x1c, 0xaf, 0x65, 0xaf, 0x91, 0x7a, 0xe6, 0xab, 0xf8, 0x19, 0xbe, 0x05, 0xed,
	0xac, 0x1b, 0xa7, 0x04, 0x68, 0x7b, 0xca, 0xce, 0x7b, 0xf3, 0x66, 0x26, 0x33, 0x2f, 0x81, 0xa7,
	0x98, 0xaa, 0x0a, 0xcb, 0xef, 0x38, 0x2e, 0x4a, 0xa5, 0x15, 0xf3, 0x31, 0x55, 0xc5, 0x25, 0x3f,
	0x83, 0x60, 0x81, 0x19, 0xae, 0xb4, 0x2a, 0x2b, 0xf6, 0x1c, 0xfc, 0x13, 0xb5, 0xc6, 0x2a, 0x74,
	0x22, 0x2f, 0x0e, 0x84, 0x0d, 0x0c, 0xfa, 0x45, 0x6e, 0xb0, 0x0a, 0x5d, 0x8b, 0x52, 0xc0, 0x5e,
	0x40, 0xf7, 0x9b, 0xcc, 0x6a, 0xac, 0x42, 0x2f, 0xf2, 0x62, 0x57, 0x34, 0x11, 0xff, 0xe9, 0xde,
	0x56, 0x4c, 0x55, 0xce, 0x22, 0xe8, 0x9f, 0xe2, 0x46, 0xe6, 0xeb, 0x69, 0xa9, 0xea, 0x22, 0x74,
	0x22, 0x27, 0x0e, 0xc4, 0x2e, 0xc4, 0x38, 0x0c, 0x6c, 0xb8, 0xa0, 0x21, 0x42, 0x97, 0x52, 0xee,
	0x60, 0x2c, 0x86, 0x67, 0xf3, 0x52, 0xad, 0x6b, 0xaa, 0x69, 0x2b, 0x79, 0x94, 0xf6, 0x27, 0xcc,
	0xde, 0xc0, 0xb0, 0x85, 0x9a, 0x8a, 0x07, 0x94, 0xba, 0x87, 0xb3, 0x57, 0x00, 0xb3, 0x4d, 0x21,
	0x57, 0xfa, 0xeb, 0x4d, 0x81, 0xa1, 0x4f, 0x59, 0x3b, 0x88, 0xe1, 0xed, 0x14, 0xc4, 0x77, 0x2d,
	0xdf, 0x22, 0x8c, 0xc1, 0xc1, 0x12, 0x65, 0x19, 0x1e, 0x46, 0x4e, 0xec, 0x0b, 0x7a, 0x1b, 0xcd,
	0x5c, 0x15, 0x75, 0x26, 0x4d, 0x9f, 0xb0, 0x67, 0x35, 0x2d, 0xc2, 0x46, 0x10, 0xcc, 0x55, 0x96,
	0xd5, 0x5a, 0xe6, 0x3a, 0x0c, 0x48, 0xd8, 0x02, 0x7c, 0x64, 0x2b, 0x9a, 0x8d, 0x9b, 0x4f, 0x7b,
	0x07, 0x5f, 0xd8, 0x80, 0xbf, 0x06, 0x7f, 0xae, 0xd2, 0x5c, 0xb3, 0x01, 0x38, 0x17, 0xb4, 0x4a,
	0x57, 0x38, 0x17, 0x26, 0x5a, 0xd2, 0xd6, 0x5c, 0xe1, 0x2c, 0xf9, 0x0f, 0x07, 0x02, 0x81, 0x2b,
	0x2d, 0xf3, 0xeb, 0x0c, 0xd9, 0x08, 0xdc, 0x24, 0xa1, 0xd4, 0xfe, 0x64, 0x30, 0xa6, 0x8b, 0x8f,
	0xa9, 0x86, 0x70, 0x93, 0x84, 0x58, 0x41, 0xd2, 0x7d, 0x56, 0x18, 0xf6, 0x5c, 0xd0, 0x9e, 0xf7,
	0xd8, 0x73, 0xcb, 0x26, 0xb4, 0xda, 0x7d, 0x36, 0xe1, 0xc7, 0x10, 0x9c, 0xa8, 0x4c, 0x95, 0xb3,
	0xfc, 0x4a, 0xb1, 0x21, 0x78, 0x62, 0xfa, 0x91, 0xbe, 0xcb, 0x40, 0x98, 0xa7, 0xf1, 0x4e, 0x82,
	0xd7, 0x98, 0xaf, 0x9b, 0x6b, 0x37, 0xd1, 0xe4, 0x97, 0x07, 0xbd, 0x4f, 0xb3, 0xb3, 0x85, 0xb1,
	0x29, 0x7b, 0x7f, 0x6b, 0x0c, 0xba, 0x6c, 0xc5, 0x86, 0x4d, 0x97, 0xad, 0xb9, 0x5e, 0xde, 0x45,
	0x54, 0x59, 0xf1, 0x0e, 0x3b, 0x86, 0x27, 0xbb, 0xd6, 0x79, 0xa8, 0x6c, 0x02, 0x60, 0xfc, 0xf1,
	0xa8, 0x56, 0x47, 0xd0, 0x37, 0x9a, 0xc7, 0x35, 0x8a, 0x9b, 0xd3, 0xfe, 0x25, 0xbd, 0xdf, 0x20,
	0x86, 0xe7, 0x1d, 0xf6, 0x01, 0x86, 0xa7, 0x78, 0x25, 0xeb, 0x4c, 0xb7, 0x3f, 0xa7, 0xfb, 0x7a,
	0xa4, 0x2a, 0x6f, 0x06, 0xdb, 0x5a, 0xee, 0xa1, 0x83, 0xbd, 0x85, 0xc3, 0xcf, 0xb2, 0xb0, 0x07,
	0xfb, 0xa7, 0x60, 0x7b, 0x54, 0x5a, 0x59, 0x6f, 0x8a, 0x6a, 0x83, 0xba, 0xbc, 0xf9, 0x8f, 0x62,
	0xeb, 0x45, 0xde, 0x79, 0xe7, 0x5c, 0x76, 0xe9, 0xbf, 0xe7, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x78, 0x60, 0x91, 0x8d, 0x04, 0x00, 0x00,
}