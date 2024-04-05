// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: project/v1/project_service.proto

package projectV1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	TeamId  string   `protobuf:"bytes,2,opt,name=teamId,proto3" json:"teamId,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectRequest) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *CreateProjectRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

type CreateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *CreateProjectResponse) Reset() {
	*x = CreateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectResponse) ProtoMessage() {}

func (x *CreateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectResponse.ProtoReflect.Descriptor instead.
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type GetProjectListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId string `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
}

func (x *GetProjectListRequest) Reset() {
	*x = GetProjectListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectListRequest) ProtoMessage() {}

func (x *GetProjectListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectListRequest.ProtoReflect.Descriptor instead.
func (*GetProjectListRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetProjectListRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

type GetProjectListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *GetProjectListResponse) Reset() {
	*x = GetProjectListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectListResponse) ProtoMessage() {}

func (x *GetProjectListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectListResponse.ProtoReflect.Descriptor instead.
func (*GetProjectListResponse) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetProjectListResponse) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type GetProjectInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // project id
}

func (x *GetProjectInfoRequest) Reset() {
	*x = GetProjectInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectInfoRequest) ProtoMessage() {}

func (x *GetProjectInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectInfoRequest.ProtoReflect.Descriptor instead.
func (*GetProjectInfoRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetProjectInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProjectInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GetProjectInfoResponse) Reset() {
	*x = GetProjectInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectInfoResponse) ProtoMessage() {}

func (x *GetProjectInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectInfoResponse.ProtoReflect.Descriptor instead.
func (*GetProjectInfoResponse) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetProjectInfoResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type UpdateProjectInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *UpdateProjectInfoRequest) Reset() {
	*x = UpdateProjectInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectInfoRequest) ProtoMessage() {}

func (x *UpdateProjectInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectInfoRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateProjectInfoRequest) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

type UpdateProjectInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *UpdateProjectInfoResponse) Reset() {
	*x = UpdateProjectInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectInfoResponse) ProtoMessage() {}

func (x *UpdateProjectInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateProjectInfoResponse) Descriptor() ([]byte, []int) {
	return file_project_v1_project_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateProjectInfoResponse) GetProject() *Project {
	if x != nil {
		return x.Project
	}
	return nil
}

var File_project_v1_project_service_proto protoreflect.FileDescriptor

var file_project_v1_project_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2f, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x49,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x47, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x49, 0x0a, 0x18, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x4a, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x32, 0x8d, 0x04, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x22, 0x1c, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x7b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x7b,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12,
	0x1a, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x84, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x1a, 0x1a, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e,
	0x66, 0x6f, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x73, 0x6a, 0x64, 0x66, 0x2f, 0x6e, 0x65, 0x6d, 0x65, 0x72, 0x74, 0x65, 0x73, 0x53,
	0x44, 0x4b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_v1_project_service_proto_rawDescOnce sync.Once
	file_project_v1_project_service_proto_rawDescData = file_project_v1_project_service_proto_rawDesc
)

func file_project_v1_project_service_proto_rawDescGZIP() []byte {
	file_project_v1_project_service_proto_rawDescOnce.Do(func() {
		file_project_v1_project_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_v1_project_service_proto_rawDescData)
	})
	return file_project_v1_project_service_proto_rawDescData
}

var file_project_v1_project_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_project_v1_project_service_proto_goTypes = []interface{}{
	(*CreateProjectRequest)(nil),      // 0: project.v1.CreateProjectRequest
	(*CreateProjectResponse)(nil),     // 1: project.v1.CreateProjectResponse
	(*GetProjectListRequest)(nil),     // 2: project.v1.GetProjectListRequest
	(*GetProjectListResponse)(nil),    // 3: project.v1.GetProjectListResponse
	(*GetProjectInfoRequest)(nil),     // 4: project.v1.GetProjectInfoRequest
	(*GetProjectInfoResponse)(nil),    // 5: project.v1.GetProjectInfoResponse
	(*UpdateProjectInfoRequest)(nil),  // 6: project.v1.UpdateProjectInfoRequest
	(*UpdateProjectInfoResponse)(nil), // 7: project.v1.UpdateProjectInfoResponse
	(*Project)(nil),                   // 8: project.v1.Project
}
var file_project_v1_project_service_proto_depIdxs = []int32{
	8,  // 0: project.v1.CreateProjectRequest.project:type_name -> project.v1.Project
	8,  // 1: project.v1.CreateProjectResponse.project:type_name -> project.v1.Project
	8,  // 2: project.v1.GetProjectListResponse.projects:type_name -> project.v1.Project
	8,  // 3: project.v1.GetProjectInfoResponse.project:type_name -> project.v1.Project
	8,  // 4: project.v1.UpdateProjectInfoRequest.project:type_name -> project.v1.Project
	8,  // 5: project.v1.UpdateProjectInfoResponse.project:type_name -> project.v1.Project
	0,  // 6: project.v1.ProjectService.CreateProject:input_type -> project.v1.CreateProjectRequest
	2,  // 7: project.v1.ProjectService.GetProjectList:input_type -> project.v1.GetProjectListRequest
	4,  // 8: project.v1.ProjectService.GetProjectInfo:input_type -> project.v1.GetProjectInfoRequest
	6,  // 9: project.v1.ProjectService.UpdateProjectInfo:input_type -> project.v1.UpdateProjectInfoRequest
	1,  // 10: project.v1.ProjectService.CreateProject:output_type -> project.v1.CreateProjectResponse
	3,  // 11: project.v1.ProjectService.GetProjectList:output_type -> project.v1.GetProjectListResponse
	5,  // 12: project.v1.ProjectService.GetProjectInfo:output_type -> project.v1.GetProjectInfoResponse
	7,  // 13: project.v1.ProjectService.UpdateProjectInfo:output_type -> project.v1.UpdateProjectInfoResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_project_v1_project_service_proto_init() }
func file_project_v1_project_service_proto_init() {
	if File_project_v1_project_service_proto != nil {
		return
	}
	file_project_v1_project_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_project_v1_project_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_v1_project_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_v1_project_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_v1_project_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_v1_project_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_v1_project_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_v1_project_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProjectInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_project_v1_project_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProjectInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_v1_project_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_v1_project_service_proto_goTypes,
		DependencyIndexes: file_project_v1_project_service_proto_depIdxs,
		MessageInfos:      file_project_v1_project_service_proto_msgTypes,
	}.Build()
	File_project_v1_project_service_proto = out.File
	file_project_v1_project_service_proto_rawDesc = nil
	file_project_v1_project_service_proto_goTypes = nil
	file_project_v1_project_service_proto_depIdxs = nil
}
