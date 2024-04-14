// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: project/v1/project.proto

package projectV1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId string `protobuf:"bytes,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId       string                 `protobuf:"bytes,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
	ProjectId   string                 `protobuf:"bytes,3,opt,name=projectId,proto3" json:"projectId,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Name        string                 `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	ServerName  string                 `protobuf:"bytes,12,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Version     string                 `protobuf:"bytes,13,opt,name=version,proto3" json:"version,omitempty"`
	Environment string                 `protobuf:"bytes,14,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{1}
}

func (x *Host) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Host) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Host) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Host) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Host) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Host) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *Host) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Host) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HostId         string                 `protobuf:"bytes,2,opt,name=hostId,proto3" json:"hostId,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	CompatibleHash string                 `protobuf:"bytes,11,opt,name=compatibleHash,proto3" json:"compatibleHash,omitempty"`
	StartedAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	EndedAt        *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Profile) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *Profile) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Profile) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Profile) GetCompatibleHash() string {
	if x != nil {
		return x.CompatibleHash
	}
	return ""
}

func (x *Profile) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Profile) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

type ProfileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProfileId  string                 `protobuf:"bytes,4,opt,name=profileId,proto3" json:"profileId,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	SampleType string                 `protobuf:"bytes,11,opt,name=sampleType,proto3" json:"sampleType,omitempty"`
	SampleUnit string                 `protobuf:"bytes,12,opt,name=sampleUnit,proto3" json:"sampleUnit,omitempty"`
	TotalValue int64                  `protobuf:"varint,13,opt,name=totalValue,proto3" json:"totalValue,omitempty"`
}

func (x *ProfileMeta) Reset() {
	*x = ProfileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileMeta) ProtoMessage() {}

func (x *ProfileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileMeta.ProtoReflect.Descriptor instead.
func (*ProfileMeta) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProfileMeta) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *ProfileMeta) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProfileMeta) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ProfileMeta) GetSampleType() string {
	if x != nil {
		return x.SampleType
	}
	return ""
}

func (x *ProfileMeta) GetSampleUnit() string {
	if x != nil {
		return x.SampleUnit
	}
	return ""
}

func (x *ProfileMeta) GetTotalValue() int64 {
	if x != nil {
		return x.TotalValue
	}
	return 0
}

type ProfileWithMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile       `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Metas   []*ProfileMeta `protobuf:"bytes,2,rep,name=metas,proto3" json:"metas,omitempty"`
}

func (x *ProfileWithMeta) Reset() {
	*x = ProfileWithMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileWithMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileWithMeta) ProtoMessage() {}

func (x *ProfileWithMeta) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileWithMeta.ProtoReflect.Descriptor instead.
func (*ProfileWithMeta) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{4}
}

func (x *ProfileWithMeta) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *ProfileWithMeta) GetMetas() []*ProfileMeta {
	if x != nil {
		return x.Metas
	}
	return nil
}

var File_project_v1_project_proto protoreflect.FileDescriptor

var file_project_v1_project_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xae, 0x02, 0x0a,
	0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xbd, 0x02,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69,
	0x62, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8f, 0x02,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x6f, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x2d, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x73, 0x6a, 0x64, 0x66, 0x2f, 0x6e, 0x65, 0x6d, 0x65, 0x72, 0x74, 0x65, 0x73, 0x53, 0x44, 0x4b,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_v1_project_proto_rawDescOnce sync.Once
	file_project_v1_project_proto_rawDescData = file_project_v1_project_proto_rawDesc
)

func file_project_v1_project_proto_rawDescGZIP() []byte {
	file_project_v1_project_proto_rawDescOnce.Do(func() {
		file_project_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_v1_project_proto_rawDescData)
	})
	return file_project_v1_project_proto_rawDescData
}

var file_project_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_project_v1_project_proto_goTypes = []interface{}{
	(*Project)(nil),               // 0: project.v1.Project
	(*Host)(nil),                  // 1: project.v1.Host
	(*Profile)(nil),               // 2: project.v1.Profile
	(*ProfileMeta)(nil),           // 3: project.v1.ProfileMeta
	(*ProfileWithMeta)(nil),       // 4: project.v1.ProfileWithMeta
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_project_v1_project_proto_depIdxs = []int32{
	5,  // 0: project.v1.Host.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 1: project.v1.Host.updatedAt:type_name -> google.protobuf.Timestamp
	5,  // 2: project.v1.Profile.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 3: project.v1.Profile.updatedAt:type_name -> google.protobuf.Timestamp
	5,  // 4: project.v1.Profile.startedAt:type_name -> google.protobuf.Timestamp
	5,  // 5: project.v1.Profile.endedAt:type_name -> google.protobuf.Timestamp
	5,  // 6: project.v1.ProfileMeta.createdAt:type_name -> google.protobuf.Timestamp
	5,  // 7: project.v1.ProfileMeta.updatedAt:type_name -> google.protobuf.Timestamp
	2,  // 8: project.v1.ProfileWithMeta.profile:type_name -> project.v1.Profile
	3,  // 9: project.v1.ProfileWithMeta.metas:type_name -> project.v1.ProfileMeta
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_project_v1_project_proto_init() }
func file_project_v1_project_proto_init() {
	if File_project_v1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_project_v1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
		file_project_v1_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_project_v1_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileMeta); i {
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
		file_project_v1_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileWithMeta); i {
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
			RawDescriptor: file_project_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_v1_project_proto_goTypes,
		DependencyIndexes: file_project_v1_project_proto_depIdxs,
		MessageInfos:      file_project_v1_project_proto_msgTypes,
	}.Build()
	File_project_v1_project_proto = out.File
	file_project_v1_project_proto_rawDesc = nil
	file_project_v1_project_proto_goTypes = nil
	file_project_v1_project_proto_depIdxs = nil
}
