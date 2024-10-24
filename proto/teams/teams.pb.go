// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: proto/teams/teams.proto

package teams_proto

import (
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

type GetOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOneRequest) Reset() {
	*x = GetOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_teams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRequest) ProtoMessage() {}

func (x *GetOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_teams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneRequest.ProtoReflect.Descriptor instead.
func (*GetOneRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_teams_proto_rawDescGZIP(), []int{0}
}

func (x *GetOneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Photo       string `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
	IsModerated bool   `protobuf:"varint,4,opt,name=is_moderated,json=isModerated,proto3" json:"is_moderated,omitempty"`
	OwnerId     int32  `protobuf:"varint,5,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *GetOneResponse) Reset() {
	*x = GetOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_teams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneResponse) ProtoMessage() {}

func (x *GetOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_teams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneResponse.ProtoReflect.Descriptor instead.
func (*GetOneResponse) Descriptor() ([]byte, []int) {
	return file_proto_teams_teams_proto_rawDescGZIP(), []int{1}
}

func (x *GetOneResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetOneResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetOneResponse) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *GetOneResponse) GetIsModerated() bool {
	if x != nil {
		return x.IsModerated
	}
	return false
}

func (x *GetOneResponse) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type GetMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId   int32 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	MemberId int32 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *GetMemberRequest) Reset() {
	*x = GetMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_teams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberRequest) ProtoMessage() {}

func (x *GetMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_teams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberRequest.ProtoReflect.Descriptor instead.
func (*GetMemberRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_teams_proto_rawDescGZIP(), []int{2}
}

func (x *GetMemberRequest) GetTeamId() int32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *GetMemberRequest) GetMemberId() int32 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type GetMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *GetMemberResponse) Reset() {
	*x = GetMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_teams_teams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberResponse) ProtoMessage() {}

func (x *GetMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_teams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberResponse.ProtoReflect.Descriptor instead.
func (*GetMemberResponse) Descriptor() ([]byte, []int) {
	return file_proto_teams_teams_proto_rawDescGZIP(), []int{3}
}

func (x *GetMemberResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetMemberResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetMemberResponse) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_proto_teams_teams_proto protoreflect.FileDescriptor

var file_proto_teams_teams_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4d, 0x6f,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x05, 0x54, 0x65, 0x61,
	0x6d, 0x73, 0x12, 0x43, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x1a, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_teams_teams_proto_rawDescOnce sync.Once
	file_proto_teams_teams_proto_rawDescData = file_proto_teams_teams_proto_rawDesc
)

func file_proto_teams_teams_proto_rawDescGZIP() []byte {
	file_proto_teams_teams_proto_rawDescOnce.Do(func() {
		file_proto_teams_teams_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_teams_teams_proto_rawDescData)
	})
	return file_proto_teams_teams_proto_rawDescData
}

var file_proto_teams_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_teams_teams_proto_goTypes = []any{
	(*GetOneRequest)(nil),     // 0: teams_proto.GetOneRequest
	(*GetOneResponse)(nil),    // 1: teams_proto.GetOneResponse
	(*GetMemberRequest)(nil),  // 2: teams_proto.GetMemberRequest
	(*GetMemberResponse)(nil), // 3: teams_proto.GetMemberResponse
}
var file_proto_teams_teams_proto_depIdxs = []int32{
	0, // 0: teams_proto.Teams.GetOne:input_type -> teams_proto.GetOneRequest
	2, // 1: teams_proto.Teams.GetMember:input_type -> teams_proto.GetMemberRequest
	1, // 2: teams_proto.Teams.GetOne:output_type -> teams_proto.GetOneResponse
	3, // 3: teams_proto.Teams.GetMember:output_type -> teams_proto.GetMemberResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_teams_teams_proto_init() }
func file_proto_teams_teams_proto_init() {
	if File_proto_teams_teams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_teams_teams_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetOneRequest); i {
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
		file_proto_teams_teams_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetOneResponse); i {
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
		file_proto_teams_teams_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetMemberRequest); i {
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
		file_proto_teams_teams_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetMemberResponse); i {
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
			RawDescriptor: file_proto_teams_teams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_teams_teams_proto_goTypes,
		DependencyIndexes: file_proto_teams_teams_proto_depIdxs,
		MessageInfos:      file_proto_teams_teams_proto_msgTypes,
	}.Build()
	File_proto_teams_teams_proto = out.File
	file_proto_teams_teams_proto_rawDesc = nil
	file_proto_teams_teams_proto_goTypes = nil
	file_proto_teams_teams_proto_depIdxs = nil
}