// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: repoupdater.proto

package v1

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

type RepoUpdateSchedulerInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the repo to lookup the schedule for.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RepoUpdateSchedulerInfoRequest) Reset() {
	*x = RepoUpdateSchedulerInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repoupdater_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoUpdateSchedulerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoUpdateSchedulerInfoRequest) ProtoMessage() {}

func (x *RepoUpdateSchedulerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repoupdater_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoUpdateSchedulerInfoRequest.ProtoReflect.Descriptor instead.
func (*RepoUpdateSchedulerInfoRequest) Descriptor() ([]byte, []int) {
	return file_repoupdater_proto_rawDescGZIP(), []int{0}
}

func (x *RepoUpdateSchedulerInfoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RepoUpdateSchedulerInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedule *RepoScheduleState `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Queue    *RepoQueueState    `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *RepoUpdateSchedulerInfoResponse) Reset() {
	*x = RepoUpdateSchedulerInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repoupdater_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoUpdateSchedulerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoUpdateSchedulerInfoResponse) ProtoMessage() {}

func (x *RepoUpdateSchedulerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repoupdater_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoUpdateSchedulerInfoResponse.ProtoReflect.Descriptor instead.
func (*RepoUpdateSchedulerInfoResponse) Descriptor() ([]byte, []int) {
	return file_repoupdater_proto_rawDescGZIP(), []int{1}
}

func (x *RepoUpdateSchedulerInfoResponse) GetSchedule() *RepoScheduleState {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *RepoUpdateSchedulerInfoResponse) GetQueue() *RepoQueueState {
	if x != nil {
		return x.Queue
	}
	return nil
}

type RepoScheduleState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index           int64                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Total           int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	IntervalSeconds int64                  `protobuf:"varint,3,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	Due             *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=due,proto3" json:"due,omitempty"`
}

func (x *RepoScheduleState) Reset() {
	*x = RepoScheduleState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repoupdater_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoScheduleState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoScheduleState) ProtoMessage() {}

func (x *RepoScheduleState) ProtoReflect() protoreflect.Message {
	mi := &file_repoupdater_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoScheduleState.ProtoReflect.Descriptor instead.
func (*RepoScheduleState) Descriptor() ([]byte, []int) {
	return file_repoupdater_proto_rawDescGZIP(), []int{2}
}

func (x *RepoScheduleState) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *RepoScheduleState) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RepoScheduleState) GetIntervalSeconds() int64 {
	if x != nil {
		return x.IntervalSeconds
	}
	return 0
}

func (x *RepoScheduleState) GetDue() *timestamppb.Timestamp {
	if x != nil {
		return x.Due
	}
	return nil
}

type RepoQueueState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Total    int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Updating bool  `protobuf:"varint,3,opt,name=updating,proto3" json:"updating,omitempty"`
	Priority int64 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *RepoQueueState) Reset() {
	*x = RepoQueueState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repoupdater_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoQueueState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoQueueState) ProtoMessage() {}

func (x *RepoQueueState) ProtoReflect() protoreflect.Message {
	mi := &file_repoupdater_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoQueueState.ProtoReflect.Descriptor instead.
func (*RepoQueueState) Descriptor() ([]byte, []int) {
	return file_repoupdater_proto_rawDescGZIP(), []int{3}
}

func (x *RepoQueueState) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *RepoQueueState) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RepoQueueState) GetUpdating() bool {
	if x != nil {
		return x.Updating
	}
	return false
}

func (x *RepoQueueState) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

var File_repoupdater_proto protoreflect.FileDescriptor

var file_repoupdater_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x1e, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x1f, 0x52, 0x65, 0x70, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72,
	0x65, 0x70, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x03,
	0x64, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x64, 0x75, 0x65, 0x22, 0x74, 0x0a, 0x0e, 0x52, 0x65,
	0x70, 0x6f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x32, 0x90, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repoupdater_proto_rawDescOnce sync.Once
	file_repoupdater_proto_rawDescData = file_repoupdater_proto_rawDesc
)

func file_repoupdater_proto_rawDescGZIP() []byte {
	file_repoupdater_proto_rawDescOnce.Do(func() {
		file_repoupdater_proto_rawDescData = protoimpl.X.CompressGZIP(file_repoupdater_proto_rawDescData)
	})
	return file_repoupdater_proto_rawDescData
}

var file_repoupdater_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_repoupdater_proto_goTypes = []interface{}{
	(*RepoUpdateSchedulerInfoRequest)(nil),  // 0: repoupdater.v1.RepoUpdateSchedulerInfoRequest
	(*RepoUpdateSchedulerInfoResponse)(nil), // 1: repoupdater.v1.RepoUpdateSchedulerInfoResponse
	(*RepoScheduleState)(nil),               // 2: repoupdater.v1.RepoScheduleState
	(*RepoQueueState)(nil),                  // 3: repoupdater.v1.RepoQueueState
	(*timestamppb.Timestamp)(nil),           // 4: google.protobuf.Timestamp
}
var file_repoupdater_proto_depIdxs = []int32{
	2, // 0: repoupdater.v1.RepoUpdateSchedulerInfoResponse.schedule:type_name -> repoupdater.v1.RepoScheduleState
	3, // 1: repoupdater.v1.RepoUpdateSchedulerInfoResponse.queue:type_name -> repoupdater.v1.RepoQueueState
	4, // 2: repoupdater.v1.RepoScheduleState.due:type_name -> google.protobuf.Timestamp
	0, // 3: repoupdater.v1.RepoUpdaterService.RepoUpdateSchedulerInfo:input_type -> repoupdater.v1.RepoUpdateSchedulerInfoRequest
	1, // 4: repoupdater.v1.RepoUpdaterService.RepoUpdateSchedulerInfo:output_type -> repoupdater.v1.RepoUpdateSchedulerInfoResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_repoupdater_proto_init() }
func file_repoupdater_proto_init() {
	if File_repoupdater_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repoupdater_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoUpdateSchedulerInfoRequest); i {
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
		file_repoupdater_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoUpdateSchedulerInfoResponse); i {
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
		file_repoupdater_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoScheduleState); i {
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
		file_repoupdater_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoQueueState); i {
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
			RawDescriptor: file_repoupdater_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repoupdater_proto_goTypes,
		DependencyIndexes: file_repoupdater_proto_depIdxs,
		MessageInfos:      file_repoupdater_proto_msgTypes,
	}.Build()
	File_repoupdater_proto = out.File
	file_repoupdater_proto_rawDesc = nil
	file_repoupdater_proto_goTypes = nil
	file_repoupdater_proto_depIdxs = nil
}
