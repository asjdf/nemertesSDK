// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: team/v1/team_service.proto

package teamV1

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
	TeamService_CreateTeam_FullMethodName        = "/team.v1.TeamService/CreateTeam"
	TeamService_GetTeamList_FullMethodName       = "/team.v1.TeamService/GetTeamList"
	TeamService_GetTeamInfo_FullMethodName       = "/team.v1.TeamService/GetTeamInfo"
	TeamService_UpdateTeamInfo_FullMethodName    = "/team.v1.TeamService/UpdateTeamInfo"
	TeamService_CreateProject_FullMethodName     = "/team.v1.TeamService/CreateProject"
	TeamService_GetProjectList_FullMethodName    = "/team.v1.TeamService/GetProjectList"
	TeamService_GetProjectInfo_FullMethodName    = "/team.v1.TeamService/GetProjectInfo"
	TeamService_UpdateProjectInfo_FullMethodName = "/team.v1.TeamService/UpdateProjectInfo"
)

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error)
	GetTeamList(ctx context.Context, in *GetTeamListRequest, opts ...grpc.CallOption) (*GetTeamListResponse, error)
	GetTeamInfo(ctx context.Context, in *GetTeamInfoRequest, opts ...grpc.CallOption) (*GetTeamInfoResponse, error)
	UpdateTeamInfo(ctx context.Context, in *UpdateTeamInfoRequest, opts ...grpc.CallOption) (*UpdateTeamInfoResponse, error)
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	GetProjectList(ctx context.Context, in *GetProjectListRequest, opts ...grpc.CallOption) (*GetProjectListResponse, error)
	GetProjectInfo(ctx context.Context, in *GetProjectInfoRequest, opts ...grpc.CallOption) (*GetProjectInfoResponse, error)
	UpdateProjectInfo(ctx context.Context, in *UpdateProjectInfoRequest, opts ...grpc.CallOption) (*UpdateProjectInfoResponse, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error) {
	out := new(CreateTeamResponse)
	err := c.cc.Invoke(ctx, TeamService_CreateTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamList(ctx context.Context, in *GetTeamListRequest, opts ...grpc.CallOption) (*GetTeamListResponse, error) {
	out := new(GetTeamListResponse)
	err := c.cc.Invoke(ctx, TeamService_GetTeamList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamInfo(ctx context.Context, in *GetTeamInfoRequest, opts ...grpc.CallOption) (*GetTeamInfoResponse, error) {
	out := new(GetTeamInfoResponse)
	err := c.cc.Invoke(ctx, TeamService_GetTeamInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) UpdateTeamInfo(ctx context.Context, in *UpdateTeamInfoRequest, opts ...grpc.CallOption) (*UpdateTeamInfoResponse, error) {
	out := new(UpdateTeamInfoResponse)
	err := c.cc.Invoke(ctx, TeamService_UpdateTeamInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, TeamService_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetProjectList(ctx context.Context, in *GetProjectListRequest, opts ...grpc.CallOption) (*GetProjectListResponse, error) {
	out := new(GetProjectListResponse)
	err := c.cc.Invoke(ctx, TeamService_GetProjectList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetProjectInfo(ctx context.Context, in *GetProjectInfoRequest, opts ...grpc.CallOption) (*GetProjectInfoResponse, error) {
	out := new(GetProjectInfoResponse)
	err := c.cc.Invoke(ctx, TeamService_GetProjectInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) UpdateProjectInfo(ctx context.Context, in *UpdateProjectInfoRequest, opts ...grpc.CallOption) (*UpdateProjectInfoResponse, error) {
	out := new(UpdateProjectInfoResponse)
	err := c.cc.Invoke(ctx, TeamService_UpdateProjectInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility
type TeamServiceServer interface {
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error)
	GetTeamList(context.Context, *GetTeamListRequest) (*GetTeamListResponse, error)
	GetTeamInfo(context.Context, *GetTeamInfoRequest) (*GetTeamInfoResponse, error)
	UpdateTeamInfo(context.Context, *UpdateTeamInfoRequest) (*UpdateTeamInfoResponse, error)
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	GetProjectList(context.Context, *GetProjectListRequest) (*GetProjectListResponse, error)
	GetProjectInfo(context.Context, *GetProjectInfoRequest) (*GetProjectInfoResponse, error)
	UpdateProjectInfo(context.Context, *UpdateProjectInfoRequest) (*UpdateProjectInfoResponse, error)
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (UnimplementedTeamServiceServer) CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamList(context.Context, *GetTeamListRequest) (*GetTeamListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamList not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamInfo(context.Context, *GetTeamInfoRequest) (*GetTeamInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamInfo not implemented")
}
func (UnimplementedTeamServiceServer) UpdateTeamInfo(context.Context, *UpdateTeamInfoRequest) (*UpdateTeamInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeamInfo not implemented")
}
func (UnimplementedTeamServiceServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedTeamServiceServer) GetProjectList(context.Context, *GetProjectListRequest) (*GetProjectListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectList not implemented")
}
func (UnimplementedTeamServiceServer) GetProjectInfo(context.Context, *GetProjectInfoRequest) (*GetProjectInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectInfo not implemented")
}
func (UnimplementedTeamServiceServer) UpdateProjectInfo(context.Context, *UpdateProjectInfoRequest) (*UpdateProjectInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProjectInfo not implemented")
}
func (UnimplementedTeamServiceServer) mustEmbedUnimplementedTeamServiceServer() {}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	s.RegisterService(&TeamService_ServiceDesc, srv)
}

func _TeamService_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_CreateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamList(ctx, req.(*GetTeamListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamInfo(ctx, req.(*GetTeamInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_UpdateTeamInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).UpdateTeamInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_UpdateTeamInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).UpdateTeamInfo(ctx, req.(*UpdateTeamInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetProjectList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetProjectList(ctx, req.(*GetProjectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetProjectInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetProjectInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetProjectInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetProjectInfo(ctx, req.(*GetProjectInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_UpdateProjectInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).UpdateProjectInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_UpdateProjectInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).UpdateProjectInfo(ctx, req.(*UpdateProjectInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamService_ServiceDesc is the grpc.ServiceDesc for TeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "team.v1.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _TeamService_CreateTeam_Handler,
		},
		{
			MethodName: "GetTeamList",
			Handler:    _TeamService_GetTeamList_Handler,
		},
		{
			MethodName: "GetTeamInfo",
			Handler:    _TeamService_GetTeamInfo_Handler,
		},
		{
			MethodName: "UpdateTeamInfo",
			Handler:    _TeamService_UpdateTeamInfo_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _TeamService_CreateProject_Handler,
		},
		{
			MethodName: "GetProjectList",
			Handler:    _TeamService_GetProjectList_Handler,
		},
		{
			MethodName: "GetProjectInfo",
			Handler:    _TeamService_GetProjectInfo_Handler,
		},
		{
			MethodName: "UpdateProjectInfo",
			Handler:    _TeamService_UpdateProjectInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team/v1/team_service.proto",
}
