// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/proto/client.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// True or false if authorization is ok
type ClientAuthorizationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ClientAuthorizationResult) Reset() {
	*x = ClientAuthorizationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAuthorizationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAuthorizationResult) ProtoMessage() {}

func (x *ClientAuthorizationResult) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClientAuthorizationResult.ProtoReflect.Descriptor instead.
func (*ClientAuthorizationResult) Descriptor() ([]byte, []int) {
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientAuthorizationResult) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// The request which user has
type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*ClientRequest_NewJob
	//	*ClientRequest_JobList
	//	*ClientRequest_JobLog
	//	*ClientRequest_NodeList
	//	*ClientRequest_NodeTop
	Request isClientRequest_Request `protobuf_oneof:"request"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{2}
}

func (m *ClientRequest) GetRequest() isClientRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *ClientRequest) GetNewJob() *NewJobMessage {
	if x, ok := x.GetRequest().(*ClientRequest_NewJob); ok {
		return x.NewJob
	}
	return nil
}

func (x *ClientRequest) GetJobList() *emptypb.Empty {
	if x, ok := x.GetRequest().(*ClientRequest_JobList); ok {
		return x.JobList
	}
	return nil
}

func (x *ClientRequest) GetJobLog() *JobLogRequestMessage {
	if x, ok := x.GetRequest().(*ClientRequest_JobLog); ok {
		return x.JobLog
	}
	return nil
}

func (x *ClientRequest) GetNodeList() *emptypb.Empty {
	if x, ok := x.GetRequest().(*ClientRequest_NodeList); ok {
		return x.NodeList
	}
	return nil
}

func (x *ClientRequest) GetNodeTop() *emptypb.Empty {
	if x, ok := x.GetRequest().(*ClientRequest_NodeTop); ok {
		return x.NodeTop
	}
	return nil
}

type isClientRequest_Request interface {
	isClientRequest_Request()
}

type ClientRequest_NewJob struct {
	NewJob *NewJobMessage `protobuf:"bytes,1,opt,name=newJob,proto3,oneof"`
}

type ClientRequest_JobList struct {
	JobList *emptypb.Empty `protobuf:"bytes,2,opt,name=jobList,proto3,oneof"`
}

type ClientRequest_JobLog struct {
	JobLog *JobLogRequestMessage `protobuf:"bytes,3,opt,name=jobLog,proto3,oneof"`
}

type ClientRequest_NodeList struct {
	NodeList *emptypb.Empty `protobuf:"bytes,4,opt,name=nodeList,proto3,oneof"`
}

type ClientRequest_NodeTop struct {
	NodeTop *emptypb.Empty `protobuf:"bytes,5,opt,name=nodeTop,proto3,oneof"`
}

func (*ClientRequest_NewJob) isClientRequest_Request() {}

func (*ClientRequest_JobList) isClientRequest_Request() {}

func (*ClientRequest_JobLog) isClientRequest_Request() {}

func (*ClientRequest_NodeList) isClientRequest_Request() {}

func (*ClientRequest_NodeTop) isClientRequest_Request() {}

// A job to send to a slave
type NewJobMessage struct {
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

func (x *NewJobMessage) Reset() {
	*x = NewJobMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewJobMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewJobMessage) ProtoMessage() {}

func (x *NewJobMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_client_proto_msgTypes[3]
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
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{3}
}

func (x *NewJobMessage) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *NewJobMessage) GetProgram() []byte {
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

func (x *NewJobMessage) GetNeededCores() uint64 {
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

type JobLogRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// tail or head?
	Tail bool `protobuf:"varint,2,opt,name=tail,proto3" json:"tail,omitempty"`
	// Show live logs?
	Live      bool   `protobuf:"varint,3,opt,name=live,proto3" json:"live,omitempty"`
	LineCount uint32 `protobuf:"varint,4,opt,name=line_count,json=lineCount,proto3" json:"line_count,omitempty"`
}

func (x *JobLogRequestMessage) Reset() {
	*x = JobLogRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobLogRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobLogRequestMessage) ProtoMessage() {}

func (x *JobLogRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobLogRequestMessage.ProtoReflect.Descriptor instead.
func (*JobLogRequestMessage) Descriptor() ([]byte, []int) {
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{4}
}

func (x *JobLogRequestMessage) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *JobLogRequestMessage) GetTail() bool {
	if x != nil {
		return x.Tail
	}
	return false
}

func (x *JobLogRequestMessage) GetLive() bool {
	if x != nil {
		return x.Live
	}
	return false
}

func (x *JobLogRequestMessage) GetLineCount() uint32 {
	if x != nil {
		return x.LineCount
	}
	return 0
}

// Status of a slave to send to client with node list
type SlaveStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Dead    bool   `protobuf:"varint,3,opt,name=dead,proto3" json:"dead,omitempty"`
}

func (x *SlaveStatus) Reset() {
	*x = SlaveStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveStatus) ProtoMessage() {}

func (x *SlaveStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveStatus.ProtoReflect.Descriptor instead.
func (*SlaveStatus) Descriptor() ([]byte, []int) {
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{5}
}

func (x *SlaveStatus) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SlaveStatus) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SlaveStatus) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

// An array of SlaveStatus
type SlavesStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status []*SlaveStatus `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
}

func (x *SlavesStatus) Reset() {
	*x = SlavesStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlavesStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlavesStatus) ProtoMessage() {}

func (x *SlavesStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlavesStatus.ProtoReflect.Descriptor instead.
func (*SlavesStatus) Descriptor() ([]byte, []int) {
	return file_pkg_proto_client_proto_rawDescGZIP(), []int{6}
}

func (x *SlavesStatus) GetStatus() []*SlaveStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_pkg_proto_client_proto protoreflect.FileDescriptor

var file_pkg_proto_client_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x13,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x19, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x9f, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x6e, 0x65,
	0x77, 0x4a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x12, 0x32, 0x0a, 0x07, 0x6a, 0x6f,
	0x62, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x06, 0x6a,
	0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x12, 0x34, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48,
	0x00, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x54, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x70, 0x42,
	0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfa, 0x01, 0x0a, 0x0d, 0x4e,
	0x65, 0x77, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x1d,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a,
	0x0d, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52,
	0x0b, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x26, 0x0a, 0x0c, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x48, 0x03, 0x52, 0x0b, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x74, 0x0a, 0x14, 0x4a, 0x6f, 0x62, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4b, 0x0a,
	0x0b, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x61, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x61, 0x64, 0x22, 0x3a, 0x0a, 0x0c, 0x53, 0x6c,
	0x61, 0x76, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x57, 0x4c, 0x46, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pkg_proto_client_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_client_proto_goTypes = []interface{}{
	(*ClientAuthorization)(nil),       // 0: proto.ClientAuthorization
	(*ClientAuthorizationResult)(nil), // 1: proto.ClientAuthorizationResult
	(*ClientRequest)(nil),             // 2: proto.ClientRequest
	(*NewJobMessage)(nil),             // 3: proto.NewJobMessage
	(*JobLogRequestMessage)(nil),      // 4: proto.JobLogRequestMessage
	(*SlaveStatus)(nil),               // 5: proto.SlaveStatus
	(*SlavesStatus)(nil),              // 6: proto.SlavesStatus
	(*emptypb.Empty)(nil),             // 7: google.protobuf.Empty
}
var file_pkg_proto_client_proto_depIdxs = []int32{
	3, // 0: proto.ClientRequest.newJob:type_name -> proto.NewJobMessage
	7, // 1: proto.ClientRequest.jobList:type_name -> google.protobuf.Empty
	4, // 2: proto.ClientRequest.jobLog:type_name -> proto.JobLogRequestMessage
	7, // 3: proto.ClientRequest.nodeList:type_name -> google.protobuf.Empty
	7, // 4: proto.ClientRequest.nodeTop:type_name -> google.protobuf.Empty
	5, // 5: proto.SlavesStatus.status:type_name -> proto.SlaveStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
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
			switch v := v.(*ClientAuthorizationResult); i {
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
		file_pkg_proto_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_pkg_proto_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_proto_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobLogRequestMessage); i {
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
		file_pkg_proto_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveStatus); i {
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
		file_pkg_proto_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlavesStatus); i {
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
	file_pkg_proto_client_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ClientRequest_NewJob)(nil),
		(*ClientRequest_JobList)(nil),
		(*ClientRequest_JobLog)(nil),
		(*ClientRequest_NodeList)(nil),
		(*ClientRequest_NodeTop)(nil),
	}
	file_pkg_proto_client_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
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
