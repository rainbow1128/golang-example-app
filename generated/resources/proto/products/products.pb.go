// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources/proto/products/products.proto

package products

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

type ListProductOut_Status int32

const (
	ListProductOut_OK        ListProductOut_Status = 0
	ListProductOut_NOT_FOUND ListProductOut_Status = 2
)

var ListProductOut_Status_name = map[int32]string{
	0: "OK",
	2: "NOT_FOUND",
}
var ListProductOut_Status_value = map[string]int32{
	"OK":        0,
	"NOT_FOUND": 2,
}

func (x ListProductOut_Status) String() string {
	return proto.EnumName(ListProductOut_Status_name, int32(x))
}
func (ListProductOut_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_products_61900bba6a66f708, []int{2, 0}
}

type Product struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_61900bba6a66f708, []int{0}
}
func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (dst *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(dst, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListProductIn struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListProductIn) Reset()         { *m = ListProductIn{} }
func (m *ListProductIn) String() string { return proto.CompactTextString(m) }
func (*ListProductIn) ProtoMessage()    {}
func (*ListProductIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_61900bba6a66f708, []int{1}
}
func (m *ListProductIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductIn.Unmarshal(m, b)
}
func (m *ListProductIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductIn.Marshal(b, m, deterministic)
}
func (dst *ListProductIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductIn.Merge(dst, src)
}
func (m *ListProductIn) XXX_Size() int {
	return xxx_messageInfo_ListProductIn.Size(m)
}
func (m *ListProductIn) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductIn.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductIn proto.InternalMessageInfo

func (m *ListProductIn) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListProductOut struct {
	Status               ListProductOut_Status `protobuf:"varint,1,opt,name=status,proto3,enum=products.ListProductOut_Status" json:"status,omitempty"`
	Products             []*Product            `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListProductOut) Reset()         { *m = ListProductOut{} }
func (m *ListProductOut) String() string { return proto.CompactTextString(m) }
func (*ListProductOut) ProtoMessage()    {}
func (*ListProductOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_61900bba6a66f708, []int{2}
}
func (m *ListProductOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProductOut.Unmarshal(m, b)
}
func (m *ListProductOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProductOut.Marshal(b, m, deterministic)
}
func (dst *ListProductOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProductOut.Merge(dst, src)
}
func (m *ListProductOut) XXX_Size() int {
	return xxx_messageInfo_ListProductOut.Size(m)
}
func (m *ListProductOut) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProductOut.DiscardUnknown(m)
}

var xxx_messageInfo_ListProductOut proto.InternalMessageInfo

func (m *ListProductOut) GetStatus() ListProductOut_Status {
	if m != nil {
		return m.Status
	}
	return ListProductOut_OK
}

func (m *ListProductOut) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "products.Product")
	proto.RegisterType((*ListProductIn)(nil), "products.ListProductIn")
	proto.RegisterType((*ListProductOut)(nil), "products.ListProductOut")
	proto.RegisterEnum("products.ListProductOut_Status", ListProductOut_Status_name, ListProductOut_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductsClient is the client API for Products service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductsClient interface {
	ListProduct(ctx context.Context, in *ListProductIn, opts ...grpc.CallOption) (*ListProductOut, error)
}

type productsClient struct {
	cc *grpc.ClientConn
}

func NewProductsClient(cc *grpc.ClientConn) ProductsClient {
	return &productsClient{cc}
}

func (c *productsClient) ListProduct(ctx context.Context, in *ListProductIn, opts ...grpc.CallOption) (*ListProductOut, error) {
	out := new(ListProductOut)
	err := c.cc.Invoke(ctx, "/products.Products/ListProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServer is the server API for Products service.
type ProductsServer interface {
	ListProduct(context.Context, *ListProductIn) (*ListProductOut, error)
}

func RegisterProductsServer(s *grpc.Server, srv ProductsServer) {
	s.RegisterService(&_Products_serviceDesc, srv)
}

func _Products_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products.Products/ListProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServer).ListProduct(ctx, req.(*ListProductIn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Products_serviceDesc = grpc.ServiceDesc{
	ServiceName: "products.Products",
	HandlerType: (*ProductsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProduct",
			Handler:    _Products_ListProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources/proto/products/products.proto",
}

func init() {
	proto.RegisterFile("resources/proto/products/products.proto", fileDescriptor_products_61900bba6a66f708)
}

var fileDescriptor_products_61900bba6a66f708 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07, 0x91, 0x29, 0xa5, 0xc9,
	0x25, 0xc5, 0x70, 0x86, 0x1e, 0x58, 0x5c, 0x88, 0x03, 0xc6, 0x57, 0xd2, 0xe5, 0x62, 0x0f, 0x80,
	0xb0, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x98, 0x32,
	0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x25, 0x79, 0x2e, 0x5e, 0x9f, 0xcc, 0xe2, 0x12, 0xa8, 0x16, 0xcf, 0x3c, 0x74, 0x4d,
	0x4a, 0x33, 0x19, 0xb9, 0xf8, 0x90, 0x54, 0xf8, 0x97, 0x96, 0x08, 0x99, 0x73, 0xb1, 0x15, 0x97,
	0x24, 0x96, 0x94, 0x16, 0x83, 0x95, 0xf1, 0x19, 0xc9, 0xeb, 0xc1, 0x5d, 0x83, 0xaa, 0x52, 0x2f,
	0x18, 0xac, 0x2c, 0x08, 0xaa, 0x5c, 0x48, 0x97, 0x0b, 0xee, 0x4e, 0x09, 0x26, 0x05, 0x66, 0x0d,
	0x6e, 0x23, 0x41, 0x84, 0x56, 0xa8, 0xb6, 0x20, 0x84, 0x57, 0xe4, 0xb9, 0xd8, 0x20, 0x06, 0x08,
	0xb1, 0x71, 0x31, 0xf9, 0x7b, 0x0b, 0x30, 0x08, 0xf1, 0x72, 0x71, 0xfa, 0xf9, 0x87, 0xc4, 0xbb,
	0xf9, 0x87, 0xfa, 0xb9, 0x08, 0x30, 0x19, 0xf9, 0x71, 0x71, 0x40, 0x75, 0x15, 0x0b, 0x39, 0x71,
	0x71, 0x23, 0x59, 0x2e, 0x24, 0x8e, 0xd5, 0x4d, 0x9e, 0x79, 0x52, 0x12, 0xb8, 0x1c, 0xab, 0xc4,
	0x90, 0xc4, 0x06, 0x0e, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xc6, 0xbe, 0x6e,
	0x77, 0x01, 0x00, 0x00,
}
