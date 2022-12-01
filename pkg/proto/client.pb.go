// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: pkg/proto/client.proto

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

// Clients authorization for accessing master
type ClientAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ClientAuthorization) Reset() {
	*x = ClientAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAuthorization) ProtoMessage() {}

func (x *ClientAuthorization) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAuthorization.ProtoReflect.Descriptor instead.
func (*ClientAuthorization) Descriptor() ([]byte, []int) {
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{0}
}

func (x *ClientAuthorization) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ClientAuthorization) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// A job to send to a slave
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The command to run on node
	Cmd string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	// The executable to send to node
	Program []byte `protobuf:"bytes,2,opt,name=program,proto3,oneof" json:"program,omitempty"`
	// Needed free memory to run this application
	NeededMemory *uint64 `protobuf:"varint,3,opt,name=needed_memory,json=neededMemory,proto3,oneof" json:"needed_memory,omitempty"`
	// Needed cores to run this application
	NeededCores *uint64 `protobuf:"varint,4,opt,name=needed_cores,json=neededCores,proto3,oneof" json:"needed_cores,omitempty"`
	// Needed HDD space to run this application
	NeededSpace *uint64 `protobuf:"varint,5,opt,name=needed_space,json=neededSpace,proto3,oneof" json:"needed_space,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{1}
}

func (x *Job) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *Job) GetProgram() []byte {
	if x != nil {
		return x.Program
	}
	return nil
}

func (x *Job) GetNeededMemory() uint64 {
	if x != nil && x.NeededMemory != nil {
		return *x.NeededMemory
	}
	return 0
}

func (x *Job) GetNeededCores() uint64 {
	if x != nil && x.NeededCores != nil {
		return *x.NeededCores
	}
	return 0
}

func (x *Job) GetNeededSpace() uint64 {
	if x != nil && x.NeededSpace != nil {
		return *x.NeededSpace
	}
	return 0
}

var File_pkg_proto_client_proto protoreflect.FileDescriptor

var file_pkg_proto_client_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4d, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xf0,
	0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x6e, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01,
	0x52, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6e, 0x65, 0x65,
	0x64, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x48,
	0x03, 0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x57, 0x4c, 0x46, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_client_proto_rawDescOnce sync.Once
	file_pkg_proto_client_proto_rawDescData = file_pkg_proto_client_proto_rawDesc
)

func file_pkg_proto_client_proto_rawDescGZIP() []byte {
	file_pkg_proto_client_proto_rawDescOnce.Do(func() {
		file_pkg_proto_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_client_proto_rawDescData)
	})
	return file_pkg_proto_client_proto_rawDescData
}

var file_pkg_proto_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_client_proto_goTypes = []interface{}{
	(*ClientAuthorization)(nil), // 0: proto.ClientAuthorization
	(*Job)(nil),                 // 1: proto.Job
}
var file_pkg_proto_client_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_client_proto_init() }
func file_pkg_proto_client_proto_init() {
	if File_pkg_proto_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAuthorization); i {
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
		file_pkg_proto_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
	file_pkg_proto_client_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_client_proto_goTypes,
		DependencyIndexes: file_pkg_proto_client_proto_depIdxs,
		MessageInfos:      file_pkg_proto_client_proto_msgTypes,
	}.Build()
	File_pkg_proto_client_proto = out.File
	file_pkg_proto_client_proto_rawDesc = nil
	file_pkg_proto_client_proto_goTypes = nil
	file_pkg_proto_client_proto_depIdxs = nil
}
