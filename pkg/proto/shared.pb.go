// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: pkg/proto/shared.proto

package proto

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

type Resource int32

const (
	Resource_CPU    Resource = 0
	Resource_Memory Resource = 1
	Resource_Disk   Resource = 2
)

// Enum value maps for Resource.
var (
	Resource_name = map[int32]string{
		0: "CPU",
		1: "Memory",
		2: "Disk",
	}
	Resource_value = map[string]int32{
		"CPU":    0,
		"Memory": 1,
		"Disk":   2,
	}
)

func (x Resource) Enum() *Resource {
	p := new(Resource)
	*p = x
	return p
}

func (x Resource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Resource) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_shared_proto_enumTypes[0].Descriptor()
}

func (Resource) Type() protoreflect.EnumType {
	return &file_pkg_proto_shared_proto_enumTypes[0]
}

func (x Resource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Resource.Descriptor instead.
func (Resource) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_shared_proto_rawDescGZIP(), []int{0}
}

// A job to send to a slave
type NewJobMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The command to run on node
	Cmd string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	// The executable to send to node
	Program *PayloadProgram `protobuf:"bytes,2,opt,name=program,proto3,oneof" json:"program,omitempty"`
	// Needed free memory to run this application
	NeededMemory *uint64 `protobuf:"varint,3,opt,name=needed_memory,json=neededMemory,proto3,oneof" json:"needed_memory,omitempty"`
	// Needed cores to run this application
	NeededCores *uint32 `protobuf:"varint,4,opt,name=needed_cores,json=neededCores,proto3,oneof" json:"needed_cores,omitempty"`
	// Needed HDD space to run this application
	NeededSpace *uint64 `protobuf:"varint,5,opt,name=needed_space,json=neededSpace,proto3,oneof" json:"needed_space,omitempty"`
}

func (x *NewJobMessage) Reset() {
	*x = NewJobMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewJobMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJobMessage) ProtoMessage() {}

func (x *NewJobMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewJobMessage.ProtoReflect.Descriptor instead.
func (*NewJobMessage) Descriptor() ([]byte, []int) {
	return file_pkg_proto_shared_proto_rawDescGZIP(), []int{0}
}

func (x *NewJobMessage) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *NewJobMessage) GetProgram() *PayloadProgram {
	if x != nil {
		return x.Program
	}
	return nil
}

func (x *NewJobMessage) GetNeededMemory() uint64 {
	if x != nil && x.NeededMemory != nil {
		return *x.NeededMemory
	}
	return 0
}

func (x *NewJobMessage) GetNeededCores() uint32 {
	if x != nil && x.NeededCores != nil {
		return *x.NeededCores
	}
	return 0
}

func (x *NewJobMessage) GetNeededSpace() uint64 {
	if x != nil && x.NeededSpace != nil {
		return *x.NeededSpace
	}
	return 0
}

// Jobs might have a program in them
type PayloadProgram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Binary of application
	ProgramBin []byte `protobuf:"bytes,1,opt,name=program_bin,json=programBin,proto3" json:"program_bin,omitempty"`
	// Name of the program
	ProgramName string `protobuf:"bytes,2,opt,name=program_name,json=programName,proto3" json:"program_name,omitempty"`
}

func (x *PayloadProgram) Reset() {
	*x = PayloadProgram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadProgram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadProgram) ProtoMessage() {}

func (x *PayloadProgram) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadProgram.ProtoReflect.Descriptor instead.
func (*PayloadProgram) Descriptor() ([]byte, []int) {
	return file_pkg_proto_shared_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadProgram) GetProgramBin() []byte {
	if x != nil {
		return x.ProgramBin
	}
	return nil
}

func (x *PayloadProgram) GetProgramName() string {
	if x != nil {
		return x.ProgramName
	}
	return ""
}

// Brief resource status of a slave
type SlaveTop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobLimit    uint32 `protobuf:"varint,1,opt,name=job_limit,json=jobLimit,proto3" json:"job_limit,omitempty"`
	RunningJobs uint32 `protobuf:"varint,2,opt,name=running_jobs,json=runningJobs,proto3" json:"running_jobs,omitempty"`
	Cores       uint32 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	FreeMemory  uint64 `protobuf:"varint,4,opt,name=free_memory,json=freeMemory,proto3" json:"free_memory,omitempty"`
	FreeDisk    uint64 `protobuf:"varint,5,opt,name=free_disk,json=freeDisk,proto3" json:"free_disk,omitempty"`
}

func (x *SlaveTop) Reset() {
	*x = SlaveTop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveTop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveTop) ProtoMessage() {}

func (x *SlaveTop) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveTop.ProtoReflect.Descriptor instead.
func (*SlaveTop) Descriptor() ([]byte, []int) {
	return file_pkg_proto_shared_proto_rawDescGZIP(), []int{2}
}

func (x *SlaveTop) GetJobLimit() uint32 {
	if x != nil {
		return x.JobLimit
	}
	return 0
}

func (x *SlaveTop) GetRunningJobs() uint32 {
	if x != nil {
		return x.RunningJobs
	}
	return 0
}

func (x *SlaveTop) GetCores() uint32 {
	if x != nil {
		return x.Cores
	}
	return 0
}

func (x *SlaveTop) GetFreeMemory() uint64 {
	if x != nil {
		return x.FreeMemory
	}
	return 0
}

func (x *SlaveTop) GetFreeDisk() uint64 {
	if x != nil {
		return x.FreeDisk
	}
	return 0
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_pkg_proto_shared_proto_rawDescGZIP(), []int{3}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pkg_proto_shared_proto protoreflect.FileDescriptor

var file_pkg_proto_shared_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x91, 0x02, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x6e, 0x65, 0x65,
	0x64, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x01, 0x52, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0b, 0x6e, 0x65, 0x65,
	0x64, 0x65, 0x64, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6e,
	0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x03, 0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x5f, 0x62, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x42, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x08, 0x53, 0x6c,
	0x61, 0x76, 0x65, 0x54, 0x6f, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6a,
	0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x72, 0x65, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x66, 0x72, 0x65, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x29, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x69, 0x73,
	0x6b, 0x10, 0x02, 0x42, 0x0f, 0x5a, 0x0d, 0x57, 0x4c, 0x46, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_shared_proto_rawDescOnce sync.Once
	file_pkg_proto_shared_proto_rawDescData = file_pkg_proto_shared_proto_rawDesc
)

func file_pkg_proto_shared_proto_rawDescGZIP() []byte {
	file_pkg_proto_shared_proto_rawDescOnce.Do(func() {
		file_pkg_proto_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_shared_proto_rawDescData)
	})
	return file_pkg_proto_shared_proto_rawDescData
}

var file_pkg_proto_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_shared_proto_goTypes = []interface{}{
	(Resource)(0),          // 0: proto.Resource
	(*NewJobMessage)(nil),  // 1: proto.NewJobMessage
	(*PayloadProgram)(nil), // 2: proto.PayloadProgram
	(*SlaveTop)(nil),       // 3: proto.SlaveTop
	(*UUID)(nil),           // 4: proto.UUID
}
var file_pkg_proto_shared_proto_depIdxs = []int32{
	2, // 0: proto.NewJobMessage.program:type_name -> proto.PayloadProgram
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_shared_proto_init() }
func file_pkg_proto_shared_proto_init() {
	if File_pkg_proto_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewJobMessage); i {
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
		file_pkg_proto_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadProgram); i {
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
		file_pkg_proto_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveTop); i {
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
		file_pkg_proto_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
	file_pkg_proto_shared_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_shared_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_shared_proto_goTypes,
		DependencyIndexes: file_pkg_proto_shared_proto_depIdxs,
		EnumInfos:         file_pkg_proto_shared_proto_enumTypes,
		MessageInfos:      file_pkg_proto_shared_proto_msgTypes,
	}.Build()
	File_pkg_proto_shared_proto = out.File
	file_pkg_proto_shared_proto_rawDesc = nil
	file_pkg_proto_shared_proto_goTypes = nil
	file_pkg_proto_shared_proto_depIdxs = nil
}
