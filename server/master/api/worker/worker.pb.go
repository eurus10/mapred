// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: worker.proto

package worker

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GenericResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GenericResp) Reset() {
	*x = GenericResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResp) ProtoMessage() {}

func (x *GenericResp) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResp.ProtoReflect.Descriptor instead.
func (*GenericResp) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *GenericResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GenericResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HeartBeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *HeartBeatReq) Reset() {
	*x = HeartBeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatReq) ProtoMessage() {}

func (x *HeartBeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatReq.ProtoReflect.Descriptor instead.
func (*HeartBeatReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *HeartBeatReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ApplyForJobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ApplyForJobReq) Reset() {
	*x = ApplyForJobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyForJobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyForJobReq) ProtoMessage() {}

func (x *ApplyForJobReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyForJobReq.ProtoReflect.Descriptor instead.
func (*ApplyForJobReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *ApplyForJobReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[3]
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
	return file_worker_proto_rawDescGZIP(), []int{3}
}

func (x *Job) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Job) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ApplyForJobResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Job     *Job   `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *ApplyForJobResp) Reset() {
	*x = ApplyForJobResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyForJobResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyForJobResp) ProtoMessage() {}

func (x *ApplyForJobResp) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyForJobResp.ProtoReflect.Descriptor instead.
func (*ApplyForJobResp) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyForJobResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ApplyForJobResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApplyForJobResp) GetJob() *Job {
	if x != nil {
		return x.Job
	}
	return nil
}

type DoneJobReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId int32 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *DoneJobReq) Reset() {
	*x = DoneJobReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoneJobReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneJobReq) ProtoMessage() {}

func (x *DoneJobReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneJobReq.ProtoReflect.Descriptor instead.
func (*DoneJobReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{5}
}

func (x *DoneJobReq) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

type ReportFailureReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	JobId    int32  `protobuf:"varint,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReportFailureReq) Reset() {
	*x = ReportFailureReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFailureReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFailureReq) ProtoMessage() {}

func (x *ReportFailureReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFailureReq.ProtoReflect.Descriptor instead.
func (*ReportFailureReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{6}
}

func (x *ReportFailureReq) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *ReportFailureReq) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *ReportFailureReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PullFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName  string `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *PullFileReq) Reset() {
	*x = PullFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFileReq) ProtoMessage() {}

func (x *PullFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFileReq.ProtoReflect.Descriptor instead.
func (*PullFileReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{7}
}

func (x *PullFileReq) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *PullFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type PullFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PullFileResp) Reset() {
	*x = PullFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullFileResp) ProtoMessage() {}

func (x *PullFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullFileResp.ProtoReflect.Descriptor instead.
func (*PullFileResp) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{8}
}

func (x *PullFileResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PullFileResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PullFileResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type WriteFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName  string `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *WriteFileReq) Reset() {
	*x = WriteFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileReq) ProtoMessage() {}

func (x *WriteFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileReq.ProtoReflect.Descriptor instead.
func (*WriteFileReq) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{9}
}

func (x *WriteFileReq) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *WriteFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *WriteFileReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x46, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x03, 0x4a,
	0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x64, 0x0a, 0x0f, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x46, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1d, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62,
	0x22, 0x23, 0x0a, 0x0a, 0x44, 0x6f, 0x6e, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x12, 0x15,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x56,
	0x0a, 0x0c, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0xea, 0x02, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61,
	0x74, 0x12, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0b,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x46, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x12, 0x16, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x46, 0x6f, 0x72, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x46, 0x6f, 0x72, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x07,
	0x44, 0x6f, 0x6e, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x35, 0x0a, 0x08, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_worker_proto_goTypes = []interface{}{
	(*GenericResp)(nil),      // 0: worker.GenericResp
	(*HeartBeatReq)(nil),     // 1: worker.HeartBeatReq
	(*ApplyForJobReq)(nil),   // 2: worker.ApplyForJobReq
	(*Job)(nil),              // 3: worker.Job
	(*ApplyForJobResp)(nil),  // 4: worker.ApplyForJobResp
	(*DoneJobReq)(nil),       // 5: worker.DoneJobReq
	(*ReportFailureReq)(nil), // 6: worker.ReportFailureReq
	(*PullFileReq)(nil),      // 7: worker.PullFileReq
	(*PullFileResp)(nil),     // 8: worker.PullFileResp
	(*WriteFileReq)(nil),     // 9: worker.WriteFileReq
}
var file_worker_proto_depIdxs = []int32{
	3, // 0: worker.ApplyForJobResp.job:type_name -> worker.Job
	1, // 1: worker.WorkerService.HeartBeat:input_type -> worker.HeartBeatReq
	2, // 2: worker.WorkerService.ApplyForJob:input_type -> worker.ApplyForJobReq
	5, // 3: worker.WorkerService.DoneJob:input_type -> worker.DoneJobReq
	6, // 4: worker.WorkerService.ReportFailure:input_type -> worker.ReportFailureReq
	7, // 5: worker.WorkerService.PullFile:input_type -> worker.PullFileReq
	9, // 6: worker.WorkerService.WriteFile:input_type -> worker.WriteFileReq
	0, // 7: worker.WorkerService.HeartBeat:output_type -> worker.GenericResp
	4, // 8: worker.WorkerService.ApplyForJob:output_type -> worker.ApplyForJobResp
	0, // 9: worker.WorkerService.DoneJob:output_type -> worker.GenericResp
	0, // 10: worker.WorkerService.ReportFailure:output_type -> worker.GenericResp
	8, // 11: worker.WorkerService.PullFile:output_type -> worker.PullFileResp
	0, // 12: worker.WorkerService.WriteFile:output_type -> worker.GenericResp
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericResp); i {
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
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatReq); i {
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
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyForJobReq); i {
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
		file_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyForJobResp); i {
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
		file_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoneJobReq); i {
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
		file_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFailureReq); i {
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
		file_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFileReq); i {
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
		file_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullFileResp); i {
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
		file_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFileReq); i {
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
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerServiceClient interface {
	HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*GenericResp, error)
	ApplyForJob(ctx context.Context, in *ApplyForJobReq, opts ...grpc.CallOption) (*ApplyForJobResp, error)
	DoneJob(ctx context.Context, in *DoneJobReq, opts ...grpc.CallOption) (*GenericResp, error)
	ReportFailure(ctx context.Context, in *ReportFailureReq, opts ...grpc.CallOption) (*GenericResp, error)
	PullFile(ctx context.Context, in *PullFileReq, opts ...grpc.CallOption) (*PullFileResp, error)
	WriteFile(ctx context.Context, in *WriteFileReq, opts ...grpc.CallOption) (*GenericResp, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*GenericResp, error) {
	out := new(GenericResp)
	err := c.cc.Invoke(ctx, "/worker.WorkerService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) ApplyForJob(ctx context.Context, in *ApplyForJobReq, opts ...grpc.CallOption) (*ApplyForJobResp, error) {
	out := new(ApplyForJobResp)
	err := c.cc.Invoke(ctx, "/worker.WorkerService/ApplyForJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) DoneJob(ctx context.Context, in *DoneJobReq, opts ...grpc.CallOption) (*GenericResp, error) {
	out := new(GenericResp)
	err := c.cc.Invoke(ctx, "/worker.WorkerService/DoneJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) ReportFailure(ctx context.Context, in *ReportFailureReq, opts ...grpc.CallOption) (*GenericResp, error) {
	out := new(GenericResp)
	err := c.cc.Invoke(ctx, "/worker.WorkerService/ReportFailure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) PullFile(ctx context.Context, in *PullFileReq, opts ...grpc.CallOption) (*PullFileResp, error) {
	out := new(PullFileResp)
	err := c.cc.Invoke(ctx, "/worker.WorkerService/PullFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) WriteFile(ctx context.Context, in *WriteFileReq, opts ...grpc.CallOption) (*GenericResp, error) {
	out := new(GenericResp)
	err := c.cc.Invoke(ctx, "/worker.WorkerService/WriteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
type WorkerServiceServer interface {
	HeartBeat(context.Context, *HeartBeatReq) (*GenericResp, error)
	ApplyForJob(context.Context, *ApplyForJobReq) (*ApplyForJobResp, error)
	DoneJob(context.Context, *DoneJobReq) (*GenericResp, error)
	ReportFailure(context.Context, *ReportFailureReq) (*GenericResp, error)
	PullFile(context.Context, *PullFileReq) (*PullFileResp, error)
	WriteFile(context.Context, *WriteFileReq) (*GenericResp, error)
}

// UnimplementedWorkerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (*UnimplementedWorkerServiceServer) HeartBeat(context.Context, *HeartBeatReq) (*GenericResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (*UnimplementedWorkerServiceServer) ApplyForJob(context.Context, *ApplyForJobReq) (*ApplyForJobResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyForJob not implemented")
}
func (*UnimplementedWorkerServiceServer) DoneJob(context.Context, *DoneJobReq) (*GenericResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoneJob not implemented")
}
func (*UnimplementedWorkerServiceServer) ReportFailure(context.Context, *ReportFailureReq) (*GenericResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportFailure not implemented")
}
func (*UnimplementedWorkerServiceServer) PullFile(context.Context, *PullFileReq) (*PullFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullFile not implemented")
}
func (*UnimplementedWorkerServiceServer) WriteFile(context.Context, *WriteFileReq) (*GenericResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteFile not implemented")
}

func RegisterWorkerServiceServer(s *grpc.Server, srv WorkerServiceServer) {
	s.RegisterService(&_WorkerService_serviceDesc, srv)
}

func _WorkerService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.WorkerService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).HeartBeat(ctx, req.(*HeartBeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_ApplyForJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyForJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).ApplyForJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.WorkerService/ApplyForJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).ApplyForJob(ctx, req.(*ApplyForJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_DoneJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoneJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).DoneJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.WorkerService/DoneJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).DoneJob(ctx, req.(*DoneJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_ReportFailure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportFailureReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).ReportFailure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.WorkerService/ReportFailure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).ReportFailure(ctx, req.(*ReportFailureReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_PullFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).PullFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.WorkerService/PullFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).PullFile(ctx, req.(*PullFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_WriteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).WriteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.WorkerService/WriteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).WriteFile(ctx, req.(*WriteFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "worker.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _WorkerService_HeartBeat_Handler,
		},
		{
			MethodName: "ApplyForJob",
			Handler:    _WorkerService_ApplyForJob_Handler,
		},
		{
			MethodName: "DoneJob",
			Handler:    _WorkerService_DoneJob_Handler,
		},
		{
			MethodName: "ReportFailure",
			Handler:    _WorkerService_ReportFailure_Handler,
		},
		{
			MethodName: "PullFile",
			Handler:    _WorkerService_PullFile_Handler,
		},
		{
			MethodName: "WriteFile",
			Handler:    _WorkerService_WriteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "worker.proto",
}
