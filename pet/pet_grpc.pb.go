// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0--rc1
// source: pet.proto

package pet_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PetService_CreatePets_FullMethodName = "/pet.PetService/CreatePets"
	PetService_GetPet_FullMethodName     = "/pet.PetService/GetPet"
)

// PetServiceClient is the client API for PetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 宠物服务定义
type PetServiceClient interface {
	// 批量创建宠物
	CreatePets(ctx context.Context, in *CreatePetsRequest, opts ...grpc.CallOption) (*CreatePetsResponse, error)
	// 根据ID获取宠物
	GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error)
}

type petServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetServiceClient(cc grpc.ClientConnInterface) PetServiceClient {
	return &petServiceClient{cc}
}

func (c *petServiceClient) CreatePets(ctx context.Context, in *CreatePetsRequest, opts ...grpc.CallOption) (*CreatePetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePetsResponse)
	err := c.cc.Invoke(ctx, PetService_CreatePets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPetResponse)
	err := c.cc.Invoke(ctx, PetService_GetPet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetServiceServer is the server API for PetService service.
// All implementations must embed UnimplementedPetServiceServer
// for forward compatibility.
//
// 宠物服务定义
type PetServiceServer interface {
	// 批量创建宠物
	CreatePets(context.Context, *CreatePetsRequest) (*CreatePetsResponse, error)
	// 根据ID获取宠物
	GetPet(context.Context, *GetPetRequest) (*GetPetResponse, error)
	mustEmbedUnimplementedPetServiceServer()
}

// UnimplementedPetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPetServiceServer struct{}

func (UnimplementedPetServiceServer) CreatePets(context.Context, *CreatePetsRequest) (*CreatePetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePets not implemented")
}
func (UnimplementedPetServiceServer) GetPet(context.Context, *GetPetRequest) (*GetPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPet not implemented")
}
func (UnimplementedPetServiceServer) mustEmbedUnimplementedPetServiceServer() {}
func (UnimplementedPetServiceServer) testEmbeddedByValue()                    {}

// UnsafePetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetServiceServer will
// result in compilation errors.
type UnsafePetServiceServer interface {
	mustEmbedUnimplementedPetServiceServer()
}

func RegisterPetServiceServer(s grpc.ServiceRegistrar, srv PetServiceServer) {
	// If the following call pancis, it indicates UnimplementedPetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PetService_ServiceDesc, srv)
}

func _PetService_CreatePets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).CreatePets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_CreatePets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).CreatePets(ctx, req.(*CreatePetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_GetPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).GetPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_GetPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).GetPet(ctx, req.(*GetPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetService_ServiceDesc is the grpc.ServiceDesc for PetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pet.PetService",
	HandlerType: (*PetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePets",
			Handler:    _PetService_CreatePets_Handler,
		},
		{
			MethodName: "GetPet",
			Handler:    _PetService_GetPet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pet.proto",
}
