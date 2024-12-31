// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc/egress.proto

package rpc

import (
	livekit "github.com/livekit/protocol/livekit"
	_ "github.com/livekit/psrpc/protoc-gen-psrpc/options"
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

type StartEgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request metadata
	EgressId string `protobuf:"bytes,1,opt,name=egress_id,json=egressId,proto3" json:"egress_id,omitempty"`
	// request
	//
	// Types that are assignable to Request:
	//
	//	*StartEgressRequest_RoomComposite
	//	*StartEgressRequest_Web
	//	*StartEgressRequest_Participant
	//	*StartEgressRequest_TrackComposite
	//	*StartEgressRequest_Track
	Request isStartEgressRequest_Request `protobuf_oneof:"request"`
	// connection info
	RoomId string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Token  string `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	WsUrl  string `protobuf:"bytes,9,opt,name=ws_url,json=wsUrl,proto3" json:"ws_url,omitempty"`
	// cloud only
	CloudBackupEnabled bool    `protobuf:"varint,10,opt,name=cloud_backup_enabled,json=cloudBackupEnabled,proto3" json:"cloud_backup_enabled,omitempty"`
	EstimatedCpu       float64 `protobuf:"fixed64,14,opt,name=estimated_cpu,json=estimatedCpu,proto3" json:"estimated_cpu,omitempty"`
}

func (x *StartEgressRequest) Reset() {
	*x = StartEgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_egress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartEgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartEgressRequest) ProtoMessage() {}

func (x *StartEgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_egress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartEgressRequest.ProtoReflect.Descriptor instead.
func (*StartEgressRequest) Descriptor() ([]byte, []int) {
	return file_rpc_egress_proto_rawDescGZIP(), []int{0}
}

func (x *StartEgressRequest) GetEgressId() string {
	if x != nil {
		return x.EgressId
	}
	return ""
}

func (m *StartEgressRequest) GetRequest() isStartEgressRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *StartEgressRequest) GetRoomComposite() *livekit.RoomCompositeEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_RoomComposite); ok {
		return x.RoomComposite
	}
	return nil
}

func (x *StartEgressRequest) GetWeb() *livekit.WebEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_Web); ok {
		return x.Web
	}
	return nil
}

func (x *StartEgressRequest) GetParticipant() *livekit.ParticipantEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_Participant); ok {
		return x.Participant
	}
	return nil
}

func (x *StartEgressRequest) GetTrackComposite() *livekit.TrackCompositeEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_TrackComposite); ok {
		return x.TrackComposite
	}
	return nil
}

func (x *StartEgressRequest) GetTrack() *livekit.TrackEgressRequest {
	if x, ok := x.GetRequest().(*StartEgressRequest_Track); ok {
		return x.Track
	}
	return nil
}

func (x *StartEgressRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *StartEgressRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StartEgressRequest) GetWsUrl() string {
	if x != nil {
		return x.WsUrl
	}
	return ""
}

func (x *StartEgressRequest) GetCloudBackupEnabled() bool {
	if x != nil {
		return x.CloudBackupEnabled
	}
	return false
}

func (x *StartEgressRequest) GetEstimatedCpu() float64 {
	if x != nil {
		return x.EstimatedCpu
	}
	return 0
}

type isStartEgressRequest_Request interface {
	isStartEgressRequest_Request()
}

type StartEgressRequest_RoomComposite struct {
	RoomComposite *livekit.RoomCompositeEgressRequest `protobuf:"bytes,5,opt,name=room_composite,json=roomComposite,proto3,oneof"`
}

type StartEgressRequest_Web struct {
	Web *livekit.WebEgressRequest `protobuf:"bytes,11,opt,name=web,proto3,oneof"`
}

type StartEgressRequest_Participant struct {
	Participant *livekit.ParticipantEgressRequest `protobuf:"bytes,13,opt,name=participant,proto3,oneof"`
}

type StartEgressRequest_TrackComposite struct {
	TrackComposite *livekit.TrackCompositeEgressRequest `protobuf:"bytes,6,opt,name=track_composite,json=trackComposite,proto3,oneof"`
}

type StartEgressRequest_Track struct {
	Track *livekit.TrackEgressRequest `protobuf:"bytes,7,opt,name=track,proto3,oneof"`
}

func (*StartEgressRequest_RoomComposite) isStartEgressRequest_Request() {}

func (*StartEgressRequest_Web) isStartEgressRequest_Request() {}

func (*StartEgressRequest_Participant) isStartEgressRequest_Request() {}

func (*StartEgressRequest_TrackComposite) isStartEgressRequest_Request() {}

func (*StartEgressRequest_Track) isStartEgressRequest_Request() {}

type ListActiveEgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListActiveEgressRequest) Reset() {
	*x = ListActiveEgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_egress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveEgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveEgressRequest) ProtoMessage() {}

func (x *ListActiveEgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_egress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveEgressRequest.ProtoReflect.Descriptor instead.
func (*ListActiveEgressRequest) Descriptor() ([]byte, []int) {
	return file_rpc_egress_proto_rawDescGZIP(), []int{1}
}

type ListActiveEgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EgressIds []string `protobuf:"bytes,1,rep,name=egress_ids,json=egressIds,proto3" json:"egress_ids,omitempty"`
}

func (x *ListActiveEgressResponse) Reset() {
	*x = ListActiveEgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_egress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActiveEgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActiveEgressResponse) ProtoMessage() {}

func (x *ListActiveEgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_egress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActiveEgressResponse.ProtoReflect.Descriptor instead.
func (*ListActiveEgressResponse) Descriptor() ([]byte, []int) {
	return file_rpc_egress_proto_rawDescGZIP(), []int{2}
}

func (x *ListActiveEgressResponse) GetEgressIds() []string {
	if x != nil {
		return x.EgressIds
	}
	return nil
}

var File_rpc_egress_proto protoreflect.FileDescriptor

var file_rpc_egress_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f,
	0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x04, 0x0a,
	0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x4c, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x0d, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x12, 0x2d,
	0x0a, 0x03, 0x77, 0x65, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x57, 0x65, 0x62, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x03, 0x77, 0x65, 0x62, 0x12, 0x45, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x77, 0x73, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x73, 0x55, 0x72, 0x6c,
	0x12, 0x30, 0x0a, 0x14, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x63, 0x70, 0x75, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x73, 0x74, 0x69, 0x6d,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x70, 0x75, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x64, 0x73, 0x32, 0xb2, 0x01, 0x0a, 0x0e, 0x45, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x45, 0x0a, 0x0b, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x08, 0xb2, 0x89, 0x01, 0x04, 0x10, 0x01,
	0x30, 0x01, 0x12, 0x59, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x08, 0xb2, 0x89, 0x01, 0x04, 0x10, 0x01, 0x28, 0x01, 0x32, 0xa1, 0x01,
	0x0a, 0x0d, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x1c, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x10, 0x01, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x74,
	0x6f, 0x70, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x06, 0xb2, 0x89, 0x01, 0x02, 0x10,
	0x01, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_egress_proto_rawDescOnce sync.Once
	file_rpc_egress_proto_rawDescData = file_rpc_egress_proto_rawDesc
)

func file_rpc_egress_proto_rawDescGZIP() []byte {
	file_rpc_egress_proto_rawDescOnce.Do(func() {
		file_rpc_egress_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_egress_proto_rawDescData)
	})
	return file_rpc_egress_proto_rawDescData
}

var file_rpc_egress_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_egress_proto_goTypes = []interface{}{
	(*StartEgressRequest)(nil),                  // 0: rpc.StartEgressRequest
	(*ListActiveEgressRequest)(nil),             // 1: rpc.ListActiveEgressRequest
	(*ListActiveEgressResponse)(nil),            // 2: rpc.ListActiveEgressResponse
	(*livekit.RoomCompositeEgressRequest)(nil),  // 3: livekit.RoomCompositeEgressRequest
	(*livekit.WebEgressRequest)(nil),            // 4: livekit.WebEgressRequest
	(*livekit.ParticipantEgressRequest)(nil),    // 5: livekit.ParticipantEgressRequest
	(*livekit.TrackCompositeEgressRequest)(nil), // 6: livekit.TrackCompositeEgressRequest
	(*livekit.TrackEgressRequest)(nil),          // 7: livekit.TrackEgressRequest
	(*livekit.UpdateStreamRequest)(nil),         // 8: livekit.UpdateStreamRequest
	(*livekit.StopEgressRequest)(nil),           // 9: livekit.StopEgressRequest
	(*livekit.EgressInfo)(nil),                  // 10: livekit.EgressInfo
}
var file_rpc_egress_proto_depIdxs = []int32{
	3,  // 0: rpc.StartEgressRequest.room_composite:type_name -> livekit.RoomCompositeEgressRequest
	4,  // 1: rpc.StartEgressRequest.web:type_name -> livekit.WebEgressRequest
	5,  // 2: rpc.StartEgressRequest.participant:type_name -> livekit.ParticipantEgressRequest
	6,  // 3: rpc.StartEgressRequest.track_composite:type_name -> livekit.TrackCompositeEgressRequest
	7,  // 4: rpc.StartEgressRequest.track:type_name -> livekit.TrackEgressRequest
	0,  // 5: rpc.EgressInternal.StartEgress:input_type -> rpc.StartEgressRequest
	1,  // 6: rpc.EgressInternal.ListActiveEgress:input_type -> rpc.ListActiveEgressRequest
	8,  // 7: rpc.EgressHandler.UpdateStream:input_type -> livekit.UpdateStreamRequest
	9,  // 8: rpc.EgressHandler.StopEgress:input_type -> livekit.StopEgressRequest
	10, // 9: rpc.EgressInternal.StartEgress:output_type -> livekit.EgressInfo
	2,  // 10: rpc.EgressInternal.ListActiveEgress:output_type -> rpc.ListActiveEgressResponse
	10, // 11: rpc.EgressHandler.UpdateStream:output_type -> livekit.EgressInfo
	10, // 12: rpc.EgressHandler.StopEgress:output_type -> livekit.EgressInfo
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_egress_proto_init() }
func file_rpc_egress_proto_init() {
	if File_rpc_egress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_egress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartEgressRequest); i {
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
		file_rpc_egress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveEgressRequest); i {
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
		file_rpc_egress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActiveEgressResponse); i {
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
	file_rpc_egress_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StartEgressRequest_RoomComposite)(nil),
		(*StartEgressRequest_Web)(nil),
		(*StartEgressRequest_Participant)(nil),
		(*StartEgressRequest_TrackComposite)(nil),
		(*StartEgressRequest_Track)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_egress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rpc_egress_proto_goTypes,
		DependencyIndexes: file_rpc_egress_proto_depIdxs,
		MessageInfos:      file_rpc_egress_proto_msgTypes,
	}.Build()
	File_rpc_egress_proto = out.File
	file_rpc_egress_proto_rawDesc = nil
	file_rpc_egress_proto_goTypes = nil
	file_rpc_egress_proto_depIdxs = nil
}
