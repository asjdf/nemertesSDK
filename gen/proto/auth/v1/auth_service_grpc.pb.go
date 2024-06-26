// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: auth/v1/auth_service.proto

package authV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthService_SwitchOrg_FullMethodName     = "/auth.v1.AuthService/SwitchOrg"
	AuthService_RefreshToken_FullMethodName  = "/auth.v1.AuthService/RefreshToken"
	AuthService_OAuthRedirect_FullMethodName = "/auth.v1.AuthService/OAuthRedirect"
	AuthService_OAuthCallback_FullMethodName = "/auth.v1.AuthService/OAuthCallback"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	SwitchOrg(ctx context.Context, in *SwitchOrgRequest, opts ...grpc.CallOption) (*SwitchOrgResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	OAuthRedirect(ctx context.Context, in *OAuthRedirectRequest, opts ...grpc.CallOption) (*OAuthRedirectResponse, error)
	OAuthCallback(ctx context.Context, in *OAuthCallbackRequest, opts ...grpc.CallOption) (*OAuthCallbackResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SwitchOrg(ctx context.Context, in *SwitchOrgRequest, opts ...grpc.CallOption) (*SwitchOrgResponse, error) {
	out := new(SwitchOrgResponse)
	err := c.cc.Invoke(ctx, AuthService_SwitchOrg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OAuthRedirect(ctx context.Context, in *OAuthRedirectRequest, opts ...grpc.CallOption) (*OAuthRedirectResponse, error) {
	out := new(OAuthRedirectResponse)
	err := c.cc.Invoke(ctx, AuthService_OAuthRedirect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) OAuthCallback(ctx context.Context, in *OAuthCallbackRequest, opts ...grpc.CallOption) (*OAuthCallbackResponse, error) {
	out := new(OAuthCallbackResponse)
	err := c.cc.Invoke(ctx, AuthService_OAuthCallback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	SwitchOrg(context.Context, *SwitchOrgRequest) (*SwitchOrgResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	OAuthRedirect(context.Context, *OAuthRedirectRequest) (*OAuthRedirectResponse, error)
	OAuthCallback(context.Context, *OAuthCallbackRequest) (*OAuthCallbackResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) SwitchOrg(context.Context, *SwitchOrgRequest) (*SwitchOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchOrg not implemented")
}
func (UnimplementedAuthServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServiceServer) OAuthRedirect(context.Context, *OAuthRedirectRequest) (*OAuthRedirectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OAuthRedirect not implemented")
}
func (UnimplementedAuthServiceServer) OAuthCallback(context.Context, *OAuthCallbackRequest) (*OAuthCallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OAuthCallback not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_SwitchOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SwitchOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SwitchOrg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SwitchOrg(ctx, req.(*SwitchOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OAuthRedirect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthRedirectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OAuthRedirect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OAuthRedirect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OAuthRedirect(ctx, req.(*OAuthRedirectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_OAuthCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).OAuthCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_OAuthCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).OAuthCallback(ctx, req.(*OAuthCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SwitchOrg",
			Handler:    _AuthService_SwitchOrg_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
		{
			MethodName: "OAuthRedirect",
			Handler:    _AuthService_OAuthRedirect_Handler,
		},
		{
			MethodName: "OAuthCallback",
			Handler:    _AuthService_OAuthCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/auth_service.proto",
}
