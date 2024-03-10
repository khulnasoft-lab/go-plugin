// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v4.25.2
// source: tests/well-known/proto/known.proto

package proto

import (
	context "context"
	durationpb "github.com/khulnasoft-lab/go-plugin/types/known/durationpb"
	emptypb "github.com/khulnasoft-lab/go-plugin/types/known/emptypb"
	structpb "github.com/khulnasoft-lab/go-plugin/types/known/structpb"
	timestamppb "github.com/khulnasoft-lab/go-plugin/types/known/timestamppb"
	wrapperspb "github.com/khulnasoft-lab/go-plugin/types/known/wrapperspb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// duration
	A *durationpb.Duration `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	// timestamp
	B *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	// struct
	C *structpb.Value `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	// wrappers
	D *wrapperspb.BoolValue   `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E *wrapperspb.BytesValue  `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
	F *wrapperspb.DoubleValue `protobuf:"bytes,6,opt,name=f,proto3" json:"f,omitempty"`
	G *wrapperspb.FloatValue  `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
	H *wrapperspb.Int32Value  `protobuf:"bytes,8,opt,name=h,proto3" json:"h,omitempty"`
	I *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=i,proto3" json:"i,omitempty"`
	J *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=j,proto3" json:"j,omitempty"`
	K *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=k,proto3" json:"k,omitempty"`
	L *wrapperspb.UInt64Value `protobuf:"bytes,12,opt,name=l,proto3" json:"l,omitempty"`
}

func (x *Request) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Request) GetA() *durationpb.Duration {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Request) GetB() *timestamppb.Timestamp {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *Request) GetC() *structpb.Value {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *Request) GetD() *wrapperspb.BoolValue {
	if x != nil {
		return x.D
	}
	return nil
}

func (x *Request) GetE() *wrapperspb.BytesValue {
	if x != nil {
		return x.E
	}
	return nil
}

func (x *Request) GetF() *wrapperspb.DoubleValue {
	if x != nil {
		return x.F
	}
	return nil
}

func (x *Request) GetG() *wrapperspb.FloatValue {
	if x != nil {
		return x.G
	}
	return nil
}

func (x *Request) GetH() *wrapperspb.Int32Value {
	if x != nil {
		return x.H
	}
	return nil
}

func (x *Request) GetI() *wrapperspb.Int64Value {
	if x != nil {
		return x.I
	}
	return nil
}

func (x *Request) GetJ() *wrapperspb.StringValue {
	if x != nil {
		return x.J
	}
	return nil
}

func (x *Request) GetK() *wrapperspb.UInt32Value {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Request) GetL() *wrapperspb.UInt64Value {
	if x != nil {
		return x.L
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *durationpb.Duration    `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B *timestamppb.Timestamp  `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C *structpb.Value         `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	D *wrapperspb.BoolValue   `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E *wrapperspb.BytesValue  `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
	F *wrapperspb.DoubleValue `protobuf:"bytes,6,opt,name=f,proto3" json:"f,omitempty"`
	G *wrapperspb.FloatValue  `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
	H *wrapperspb.Int32Value  `protobuf:"bytes,8,opt,name=h,proto3" json:"h,omitempty"`
	I *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=i,proto3" json:"i,omitempty"`
	J *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=j,proto3" json:"j,omitempty"`
	K *wrapperspb.UInt32Value `protobuf:"bytes,11,opt,name=k,proto3" json:"k,omitempty"`
	L *wrapperspb.UInt64Value `protobuf:"bytes,12,opt,name=l,proto3" json:"l,omitempty"`
}

func (x *Response) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Response) GetA() *durationpb.Duration {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Response) GetB() *timestamppb.Timestamp {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *Response) GetC() *structpb.Value {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *Response) GetD() *wrapperspb.BoolValue {
	if x != nil {
		return x.D
	}
	return nil
}

func (x *Response) GetE() *wrapperspb.BytesValue {
	if x != nil {
		return x.E
	}
	return nil
}

func (x *Response) GetF() *wrapperspb.DoubleValue {
	if x != nil {
		return x.F
	}
	return nil
}

func (x *Response) GetG() *wrapperspb.FloatValue {
	if x != nil {
		return x.G
	}
	return nil
}

func (x *Response) GetH() *wrapperspb.Int32Value {
	if x != nil {
		return x.H
	}
	return nil
}

func (x *Response) GetI() *wrapperspb.Int64Value {
	if x != nil {
		return x.I
	}
	return nil
}

func (x *Response) GetJ() *wrapperspb.StringValue {
	if x != nil {
		return x.J
	}
	return nil
}

func (x *Response) GetK() *wrapperspb.UInt32Value {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Response) GetL() *wrapperspb.UInt64Value {
	if x != nil {
		return x.L
	}
	return nil
}

// go:plugin type=plugin version=1
type KnownTypesTest interface {
	Test(context.Context, *Request) (*Response, error)
}

// go:plugin type=plugin version=1
type EmptyTest interface {
	DoNothing(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}
