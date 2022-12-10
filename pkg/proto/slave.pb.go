// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.11
// source: pkg/proto/slave.proto

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

type MasterToSlaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*MasterToSlaveRequest_Ping
	//	*MasterToSlaveRequest_GetTop
	//	*MasterToSlaveRequest_NewJob
	//	*MasterToSlaveRequest_GetJobLogs
	Request isMasterToSlaveRequest_Request `protobuf_oneof:"request"`
}

func (x *MasterToSlaveRequest) Reset() {
	*x = MasterToSlaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterToSlaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterToSlaveRequest) ProtoMessage() {}

func (x *MasterToSlaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterToSlaveRequest.ProtoReflect.Descriptor instead.
func (*MasterToSlaveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{0}
}

func (m *MasterToSlaveRequest) GetRequest() isMasterToSlaveRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *MasterToSlaveRequest) GetPing() *emptypb.Empty {
	if x, ok := x.GetRequest().(*MasterToSlaveRequest_Ping); ok {
		return x.Ping
	}
	return nil
}

func (x *MasterToSlaveRequest) GetGetTop() *emptypb.Empty {
	if x, ok := x.GetRequest().(*MasterToSlaveRequest_GetTop); ok {
		return x.GetTop
	}
	return nil
}

func (x *MasterToSlaveRequest) GetNewJob() *SlaveNewJobRequest {
	if x, ok := x.GetRequest().(*MasterToSlaveRequest_NewJob); ok {
		return x.NewJob
	}
	return nil
}

func (x *MasterToSlaveRequest) GetGetJobLogs() *GetJobLogsRequest {
	if x, ok := x.GetRequest().(*MasterToSlaveRequest_GetJobLogs); ok {
		return x.GetJobLogs
	}
	return nil
}

type isMasterToSlaveRequest_Request interface {
	isMasterToSlaveRequest_Request()
}

type MasterToSlaveRequest_Ping struct {
	Ping *emptypb.Empty `protobuf:"bytes,1,opt,name=ping,proto3,oneof"`
}

type MasterToSlaveRequest_GetTop struct {
	GetTop *emptypb.Empty `protobuf:"bytes,2,opt,name=getTop,proto3,oneof"`
}

type MasterToSlaveRequest_NewJob struct {
	NewJob *SlaveNewJobRequest `protobuf:"bytes,3,opt,name=newJob,proto3,oneof"`
}

type MasterToSlaveRequest_GetJobLogs struct {
	GetJobLogs *GetJobLogsRequest `protobuf:"bytes,4,opt,name=getJobLogs,proto3,oneof"`
}

func (*MasterToSlaveRequest_Ping) isMasterToSlaveRequest_Request() {}

func (*MasterToSlaveRequest_GetTop) isMasterToSlaveRequest_Request() {}

func (*MasterToSlaveRequest_NewJob) isMasterToSlaveRequest_Request() {}

func (*MasterToSlaveRequest_GetJobLogs) isMasterToSlaveRequest_Request() {}

type SlaveNewJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *UUID          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NewJob *NewJobMessage `protobuf:"bytes,2,opt,name=newJob,proto3" json:"newJob,omitempty"`
}

func (x *SlaveNewJobRequest) Reset() {
	*x = SlaveNewJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveNewJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveNewJobRequest) ProtoMessage() {}

func (x *SlaveNewJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveNewJobRequest.ProtoReflect.Descriptor instead.
func (*SlaveNewJobRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{1}
}

func (x *SlaveNewJobRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SlaveNewJobRequest) GetNewJob() *NewJobMessage {
	if x != nil {
		return x.NewJob
	}
	return nil
}

type SlaveToMasterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*SlaveToMasterRequest_Hello
	//	*SlaveToMasterRequest_JobFinished
	Request isSlaveToMasterRequest_Request `protobuf_oneof:"request"`
}

func (x *SlaveToMasterRequest) Reset() {
	*x = SlaveToMasterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveToMasterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveToMasterRequest) ProtoMessage() {}

func (x *SlaveToMasterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveToMasterRequest.ProtoReflect.Descriptor instead.
func (*SlaveToMasterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{2}
}

func (m *SlaveToMasterRequest) GetRequest() isSlaveToMasterRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *SlaveToMasterRequest) GetHello() *SlaveHello {
	if x, ok := x.GetRequest().(*SlaveToMasterRequest_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *SlaveToMasterRequest) GetJobFinished() *JobFinishedResult {
	if x, ok := x.GetRequest().(*SlaveToMasterRequest_JobFinished); ok {
		return x.JobFinished
	}
	return nil
}

type isSlaveToMasterRequest_Request interface {
	isSlaveToMasterRequest_Request()
}

type SlaveToMasterRequest_Hello struct {
	Hello *SlaveHello `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type SlaveToMasterRequest_JobFinished struct {
	JobFinished *JobFinishedResult `protobuf:"bytes,2,opt,name=jobFinished,proto3,oneof"`
}

func (*SlaveToMasterRequest_Hello) isSlaveToMasterRequest_Request() {}

func (*SlaveToMasterRequest_JobFinished) isSlaveToMasterRequest_Request() {}

type SlaveHello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How should master connect to slave
	ConnectAddress string `protobuf:"bytes,1,opt,name=connect_address,json=connectAddress,proto3" json:"connect_address,omitempty"`
}

func (x *SlaveHello) Reset() {
	*x = SlaveHello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveHello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveHello) ProtoMessage() {}

func (x *SlaveHello) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveHello.ProtoReflect.Descriptor instead.
func (*SlaveHello) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{3}
}

func (x *SlaveHello) GetConnectAddress() string {
	if x != nil {
		return x.ConnectAddress
	}
	return ""
}

// Message sent to slave after SlaveHello
type SlaveHelloMasterAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Slave ID
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SlaveHelloMasterAck) Reset() {
	*x = SlaveHelloMasterAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveHelloMasterAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveHelloMasterAck) ProtoMessage() {}

func (x *SlaveHelloMasterAck) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveHelloMasterAck.ProtoReflect.Descriptor instead.
func (*SlaveHelloMasterAck) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{4}
}

func (x *SlaveHelloMasterAck) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Sent after NewJobMessage is sent to client
type SlaveJobAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*SlaveJobAck_Ack
	//	*SlaveJobAck_InsufficientResource
	//	*SlaveJobAck_TasksFull
	Result isSlaveJobAck_Result `protobuf_oneof:"result"`
}

func (x *SlaveJobAck) Reset() {
	*x = SlaveJobAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveJobAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveJobAck) ProtoMessage() {}

func (x *SlaveJobAck) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveJobAck.ProtoReflect.Descriptor instead.
func (*SlaveJobAck) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{5}
}

func (m *SlaveJobAck) GetResult() isSlaveJobAck_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *SlaveJobAck) GetAck() *emptypb.Empty {
	if x, ok := x.GetResult().(*SlaveJobAck_Ack); ok {
		return x.Ack
	}
	return nil
}

func (x *SlaveJobAck) GetInsufficientResource() *emptypb.Empty {
	if x, ok := x.GetResult().(*SlaveJobAck_InsufficientResource); ok {
		return x.InsufficientResource
	}
	return nil
}

func (x *SlaveJobAck) GetTasksFull() *emptypb.Empty {
	if x, ok := x.GetResult().(*SlaveJobAck_TasksFull); ok {
		return x.TasksFull
	}
	return nil
}

type isSlaveJobAck_Result interface {
	isSlaveJobAck_Result()
}

type SlaveJobAck_Ack struct {
	Ack *emptypb.Empty `protobuf:"bytes,1,opt,name=ack,proto3,oneof"` // sent if job was ok and now it's running!
}

type SlaveJobAck_InsufficientResource struct {
	InsufficientResource *emptypb.Empty `protobuf:"bytes,2,opt,name=insufficient_resource,json=insufficientResource,proto3,oneof"` // sent if slave couldn't run the job. The resource which is insufficient is sent in as data
}

type SlaveJobAck_TasksFull struct {
	TasksFull *emptypb.Empty `protobuf:"bytes,3,opt,name=tasks_full,json=tasksFull,proto3,oneof"` // sent if slave couldn't run the job because all of allowed tasks are full.
}

func (*SlaveJobAck_Ack) isSlaveJobAck_Result() {}

func (*SlaveJobAck_InsufficientResource) isSlaveJobAck_Result() {}

func (*SlaveJobAck_TasksFull) isSlaveJobAck_Result() {}

type JobFinishedResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlaveId uint32 `protobuf:"varint,1,opt,name=slave_id,json=slaveId,proto3" json:"slave_id,omitempty"`
	JobId   *UUID  `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Types that are assignable to Status:
	//	*JobFinishedResult_ExitCode
	//	*JobFinishedResult_RunError
	Status isJobFinishedResult_Status `protobuf_oneof:"status"`
}

func (x *JobFinishedResult) Reset() {
	*x = JobFinishedResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobFinishedResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobFinishedResult) ProtoMessage() {}

func (x *JobFinishedResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobFinishedResult.ProtoReflect.Descriptor instead.
func (*JobFinishedResult) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{6}
}

func (x *JobFinishedResult) GetSlaveId() uint32 {
	if x != nil {
		return x.SlaveId
	}
	return 0
}

func (x *JobFinishedResult) GetJobId() *UUID {
	if x != nil {
		return x.JobId
	}
	return nil
}

func (m *JobFinishedResult) GetStatus() isJobFinishedResult_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (x *JobFinishedResult) GetExitCode() int32 {
	if x, ok := x.GetStatus().(*JobFinishedResult_ExitCode); ok {
		return x.ExitCode
	}
	return 0
}

func (x *JobFinishedResult) GetRunError() string {
	if x, ok := x.GetStatus().(*JobFinishedResult_RunError); ok {
		return x.RunError
	}
	return ""
}

type isJobFinishedResult_Status interface {
	isJobFinishedResult_Status()
}

type JobFinishedResult_ExitCode struct {
	ExitCode int32 `protobuf:"zigzag32,3,opt,name=exit_code,json=exitCode,proto3,oneof"`
}

type JobFinishedResult_RunError struct {
	RunError string `protobuf:"bytes,4,opt,name=run_error,json=runError,proto3,oneof"`
}

func (*JobFinishedResult_ExitCode) isJobFinishedResult_Status() {}

func (*JobFinishedResult_RunError) isJobFinishedResult_Status() {}

// Sent from master to slave after SlaveHello
type SlaveIDAssigned struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlaveId uint32 `protobuf:"varint,1,opt,name=slave_id,json=slaveId,proto3" json:"slave_id,omitempty"`
}

func (x *SlaveIDAssigned) Reset() {
	*x = SlaveIDAssigned{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_slave_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlaveIDAssigned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlaveIDAssigned) ProtoMessage() {}

func (x *SlaveIDAssigned) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_slave_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlaveIDAssigned.ProtoReflect.Descriptor instead.
func (*SlaveIDAssigned) Descriptor() ([]byte, []int) {
	return file_pkg_proto_slave_proto_rawDescGZIP(), []int{7}
}

func (x *SlaveIDAssigned) GetSlaveId() uint32 {
	if x != nil {
		return x.SlaveId
	}
	return 0
}

var File_pkg_proto_slave_proto protoreflect.FileDescriptor

var file_pkg_proto_slave_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6c, 0x61, 0x76,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x54, 0x6f,
	0x53, 0x6c, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x06, 0x67, 0x65,
	0x74, 0x54, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x06, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x12, 0x33, 0x0a, 0x06,
	0x6e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x4a, 0x6f,
	0x62, 0x12, 0x3a, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0a, 0x67, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x12, 0x53, 0x6c, 0x61, 0x76,
	0x65, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x6e,
	0x65, 0x77, 0x4a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x53, 0x6c,
	0x61, 0x76, 0x65, 0x54, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x3c, 0x0a,
	0x0b, 0x6a, 0x6f, 0x62, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0b,
	0x6a, 0x6f, 0x62, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x0a, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x25, 0x0a,
	0x13, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x0b, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x4a, 0x6f,
	0x62, 0x41, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b,
	0x12, 0x4d, 0x0a, 0x15, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x14, 0x69, 0x6e, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x09, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c, 0x61, 0x76,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6c, 0x61, 0x76,
	0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x55, 0x49, 0x44,
	0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x08, 0x65, 0x78,
	0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x72, 0x75, 0x6e, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x72, 0x75, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x2c, 0x0a, 0x0f, 0x53, 0x6c, 0x61, 0x76, 0x65, 0x49, 0x44, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x49, 0x64, 0x42, 0x0f, 0x5a,
	0x0d, 0x57, 0x4c, 0x46, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_slave_proto_rawDescOnce sync.Once
	file_pkg_proto_slave_proto_rawDescData = file_pkg_proto_slave_proto_rawDesc
)

func file_pkg_proto_slave_proto_rawDescGZIP() []byte {
	file_pkg_proto_slave_proto_rawDescOnce.Do(func() {
		file_pkg_proto_slave_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_slave_proto_rawDescData)
	})
	return file_pkg_proto_slave_proto_rawDescData
}

var file_pkg_proto_slave_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_slave_proto_goTypes = []interface{}{
	(*MasterToSlaveRequest)(nil), // 0: proto.MasterToSlaveRequest
	(*SlaveNewJobRequest)(nil),   // 1: proto.SlaveNewJobRequest
	(*SlaveToMasterRequest)(nil), // 2: proto.SlaveToMasterRequest
	(*SlaveHello)(nil),           // 3: proto.SlaveHello
	(*SlaveHelloMasterAck)(nil),  // 4: proto.SlaveHelloMasterAck
	(*SlaveJobAck)(nil),          // 5: proto.SlaveJobAck
	(*JobFinishedResult)(nil),    // 6: proto.JobFinishedResult
	(*SlaveIDAssigned)(nil),      // 7: proto.SlaveIDAssigned
	(*emptypb.Empty)(nil),        // 8: google.protobuf.Empty
	(*GetJobLogsRequest)(nil),    // 9: proto.GetJobLogsRequest
	(*UUID)(nil),                 // 10: proto.UUID
	(*NewJobMessage)(nil),        // 11: proto.NewJobMessage
}
var file_pkg_proto_slave_proto_depIdxs = []int32{
	8,  // 0: proto.MasterToSlaveRequest.ping:type_name -> google.protobuf.Empty
	8,  // 1: proto.MasterToSlaveRequest.getTop:type_name -> google.protobuf.Empty
	1,  // 2: proto.MasterToSlaveRequest.newJob:type_name -> proto.SlaveNewJobRequest
	9,  // 3: proto.MasterToSlaveRequest.getJobLogs:type_name -> proto.GetJobLogsRequest
	10, // 4: proto.SlaveNewJobRequest.id:type_name -> proto.UUID
	11, // 5: proto.SlaveNewJobRequest.newJob:type_name -> proto.NewJobMessage
	3,  // 6: proto.SlaveToMasterRequest.hello:type_name -> proto.SlaveHello
	6,  // 7: proto.SlaveToMasterRequest.jobFinished:type_name -> proto.JobFinishedResult
	8,  // 8: proto.SlaveJobAck.ack:type_name -> google.protobuf.Empty
	8,  // 9: proto.SlaveJobAck.insufficient_resource:type_name -> google.protobuf.Empty
	8,  // 10: proto.SlaveJobAck.tasks_full:type_name -> google.protobuf.Empty
	10, // 11: proto.JobFinishedResult.job_id:type_name -> proto.UUID
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_pkg_proto_slave_proto_init() }
func file_pkg_proto_slave_proto_init() {
	if File_pkg_proto_slave_proto != nil {
		return
	}
	file_pkg_proto_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_slave_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterToSlaveRequest); i {
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
		file_pkg_proto_slave_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveNewJobRequest); i {
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
		file_pkg_proto_slave_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveToMasterRequest); i {
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
		file_pkg_proto_slave_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveHello); i {
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
		file_pkg_proto_slave_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveHelloMasterAck); i {
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
		file_pkg_proto_slave_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveJobAck); i {
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
		file_pkg_proto_slave_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobFinishedResult); i {
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
		file_pkg_proto_slave_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlaveIDAssigned); i {
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
	file_pkg_proto_slave_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MasterToSlaveRequest_Ping)(nil),
		(*MasterToSlaveRequest_GetTop)(nil),
		(*MasterToSlaveRequest_NewJob)(nil),
		(*MasterToSlaveRequest_GetJobLogs)(nil),
	}
	file_pkg_proto_slave_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SlaveToMasterRequest_Hello)(nil),
		(*SlaveToMasterRequest_JobFinished)(nil),
	}
	file_pkg_proto_slave_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*SlaveJobAck_Ack)(nil),
		(*SlaveJobAck_InsufficientResource)(nil),
		(*SlaveJobAck_TasksFull)(nil),
	}
	file_pkg_proto_slave_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*JobFinishedResult_ExitCode)(nil),
		(*JobFinishedResult_RunError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_slave_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_slave_proto_goTypes,
		DependencyIndexes: file_pkg_proto_slave_proto_depIdxs,
		MessageInfos:      file_pkg_proto_slave_proto_msgTypes,
	}.Build()
	File_pkg_proto_slave_proto = out.File
	file_pkg_proto_slave_proto_rawDesc = nil
	file_pkg_proto_slave_proto_goTypes = nil
	file_pkg_proto_slave_proto_depIdxs = nil
}
