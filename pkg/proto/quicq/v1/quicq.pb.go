// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: proto/quicq.proto

package quicq

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

type RequestType int32

const (
	RequestType_REQUEST_UNKNOWN              RequestType = 0
	RequestType_REQUEST_SUBSCRIBE            RequestType = 1
	RequestType_REQUEST_UNSUBSCRIBE          RequestType = 2
	RequestType_REQUEST_FETCH_TOPIC_METADATA RequestType = 3
	RequestType_REQUEST_POLL                 RequestType = 4
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "REQUEST_UNKNOWN",
		1: "REQUEST_SUBSCRIBE",
		2: "REQUEST_UNSUBSCRIBE",
		3: "REQUEST_FETCH_TOPIC_METADATA",
		4: "REQUEST_POLL",
	}
	RequestType_value = map[string]int32{
		"REQUEST_UNKNOWN":              0,
		"REQUEST_SUBSCRIBE":            1,
		"REQUEST_UNSUBSCRIBE":          2,
		"REQUEST_FETCH_TOPIC_METADATA": 3,
		"REQUEST_POLL":                 4,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_quicq_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_proto_quicq_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{0}
}

type ResponseType int32

const (
	ResponseType_RESPONSE_UNKNOWN              ResponseType = 0
	ResponseType_RESPONSE_SUBSCRIBE            ResponseType = 1
	ResponseType_RESPONSE_UNSUBSCRIBE          ResponseType = 2
	ResponseType_RESPONSE_FETCH_TOPIC_METADATA ResponseType = 3
	ResponseType_RESPONSE_POLL                 ResponseType = 4
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "RESPONSE_UNKNOWN",
		1: "RESPONSE_SUBSCRIBE",
		2: "RESPONSE_UNSUBSCRIBE",
		3: "RESPONSE_FETCH_TOPIC_METADATA",
		4: "RESPONSE_POLL",
	}
	ResponseType_value = map[string]int32{
		"RESPONSE_UNKNOWN":              0,
		"RESPONSE_SUBSCRIBE":            1,
		"RESPONSE_UNSUBSCRIBE":          2,
		"RESPONSE_FETCH_TOPIC_METADATA": 3,
		"RESPONSE_POLL":                 4,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_quicq_proto_enumTypes[1].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_proto_quicq_proto_enumTypes[1]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{1}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestType RequestType `protobuf:"varint,1,opt,name=requestType,proto3,enum=quicq.v1.RequestType" json:"requestType,omitempty"`
	// Types that are assignable to Request:
	//
	//	*Request_SubscribeRequest
	//	*Request_UnsubscribeRequest
	//	*Request_FetchTopicMetadataRequest
	//	*Request_PollRequest
	Request isRequest_Request `protobuf_oneof:"request"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetRequestType() RequestType {
	if x != nil {
		return x.RequestType
	}
	return RequestType_REQUEST_UNKNOWN
}

func (m *Request) GetRequest() isRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Request) GetSubscribeRequest() *SubscribeRequest {
	if x, ok := x.GetRequest().(*Request_SubscribeRequest); ok {
		return x.SubscribeRequest
	}
	return nil
}

func (x *Request) GetUnsubscribeRequest() *UnsubscribeRequest {
	if x, ok := x.GetRequest().(*Request_UnsubscribeRequest); ok {
		return x.UnsubscribeRequest
	}
	return nil
}

func (x *Request) GetFetchTopicMetadataRequest() *FetchTopicMetadataRequest {
	if x, ok := x.GetRequest().(*Request_FetchTopicMetadataRequest); ok {
		return x.FetchTopicMetadataRequest
	}
	return nil
}

func (x *Request) GetPollRequest() *PollRequest {
	if x, ok := x.GetRequest().(*Request_PollRequest); ok {
		return x.PollRequest
	}
	return nil
}

type isRequest_Request interface {
	isRequest_Request()
}

type Request_SubscribeRequest struct {
	SubscribeRequest *SubscribeRequest `protobuf:"bytes,2,opt,name=subscribeRequest,proto3,oneof"`
}

type Request_UnsubscribeRequest struct {
	UnsubscribeRequest *UnsubscribeRequest `protobuf:"bytes,3,opt,name=unsubscribeRequest,proto3,oneof"`
}

type Request_FetchTopicMetadataRequest struct {
	FetchTopicMetadataRequest *FetchTopicMetadataRequest `protobuf:"bytes,4,opt,name=fetchTopicMetadataRequest,proto3,oneof"`
}

type Request_PollRequest struct {
	PollRequest *PollRequest `protobuf:"bytes,5,opt,name=pollRequest,proto3,oneof"`
}

func (*Request_SubscribeRequest) isRequest_Request() {}

func (*Request_UnsubscribeRequest) isRequest_Request() {}

func (*Request_FetchTopicMetadataRequest) isRequest_Request() {}

func (*Request_PollRequest) isRequest_Request() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseType ResponseType `protobuf:"varint,1,opt,name=responseType,proto3,enum=quicq.v1.ResponseType" json:"responseType,omitempty"`
	// Types that are assignable to Response:
	//
	//	*Response_SubscribeResponse
	//	*Response_UnsubscribeResponse
	//	*Response_FetchTopicMetadataResponse
	//	*Response_PollResponse
	Response isResponse_Response `protobuf_oneof:"response"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResponseType() ResponseType {
	if x != nil {
		return x.ResponseType
	}
	return ResponseType_RESPONSE_UNKNOWN
}

func (m *Response) GetResponse() isResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *Response) GetSubscribeResponse() *SubscribeResponse {
	if x, ok := x.GetResponse().(*Response_SubscribeResponse); ok {
		return x.SubscribeResponse
	}
	return nil
}

func (x *Response) GetUnsubscribeResponse() *UnsubscribeResponse {
	if x, ok := x.GetResponse().(*Response_UnsubscribeResponse); ok {
		return x.UnsubscribeResponse
	}
	return nil
}

func (x *Response) GetFetchTopicMetadataResponse() *FetchTopicMetadataResponse {
	if x, ok := x.GetResponse().(*Response_FetchTopicMetadataResponse); ok {
		return x.FetchTopicMetadataResponse
	}
	return nil
}

func (x *Response) GetPollResponse() *PollResponse {
	if x, ok := x.GetResponse().(*Response_PollResponse); ok {
		return x.PollResponse
	}
	return nil
}

type isResponse_Response interface {
	isResponse_Response()
}

type Response_SubscribeResponse struct {
	SubscribeResponse *SubscribeResponse `protobuf:"bytes,2,opt,name=subscribeResponse,proto3,oneof"`
}

type Response_UnsubscribeResponse struct {
	UnsubscribeResponse *UnsubscribeResponse `protobuf:"bytes,3,opt,name=unsubscribeResponse,proto3,oneof"`
}

type Response_FetchTopicMetadataResponse struct {
	FetchTopicMetadataResponse *FetchTopicMetadataResponse `protobuf:"bytes,4,opt,name=fetchTopicMetadataResponse,proto3,oneof"`
}

type Response_PollResponse struct {
	PollResponse *PollResponse `protobuf:"bytes,5,opt,name=pollResponse,proto3,oneof"`
}

func (*Response_SubscribeResponse) isResponse_Response() {}

func (*Response_UnsubscribeResponse) isResponse_Response() {}

func (*Response_FetchTopicMetadataResponse) isResponse_Response() {}

func (*Response_PollResponse) isResponse_Response() {}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID int64  `protobuf:"varint,1,opt,name=consumerID,proto3" json:"consumerID,omitempty"`
	Topic      string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetConsumerID() int64 {
	if x != nil {
		return x.ConsumerID
	}
	return 0
}

func (x *SubscribeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID int64  `protobuf:"varint,1,opt,name=consumerID,proto3" json:"consumerID,omitempty"`
	Topic      string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{3}
}

func (x *UnsubscribeRequest) GetConsumerID() int64 {
	if x != nil {
		return x.ConsumerID
	}
	return 0
}

func (x *UnsubscribeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type FetchTopicMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID int64 `protobuf:"varint,1,opt,name=consumerID,proto3" json:"consumerID,omitempty"`
}

func (x *FetchTopicMetadataRequest) Reset() {
	*x = FetchTopicMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTopicMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTopicMetadataRequest) ProtoMessage() {}

func (x *FetchTopicMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTopicMetadataRequest.ProtoReflect.Descriptor instead.
func (*FetchTopicMetadataRequest) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{4}
}

func (x *FetchTopicMetadataRequest) GetConsumerID() int64 {
	if x != nil {
		return x.ConsumerID
	}
	return 0
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"` // add metadata? What is meant by metadata?
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{5}
}

func (x *SubscribeResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type UnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *UnsubscribeResponse) Reset() {
	*x = UnsubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeResponse) ProtoMessage() {}

func (x *UnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{6}
}

func (x *UnsubscribeResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type FetchTopicMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *FetchTopicMetadataResponse) Reset() {
	*x = FetchTopicMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTopicMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTopicMetadataResponse) ProtoMessage() {}

func (x *FetchTopicMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTopicMetadataResponse.ProtoReflect.Descriptor instead.
func (*FetchTopicMetadataResponse) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{7}
}

func (x *FetchTopicMetadataResponse) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type PollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerID int64 `protobuf:"varint,1,opt,name=consumerID,proto3" json:"consumerID,omitempty"`
}

func (x *PollRequest) Reset() {
	*x = PollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollRequest) ProtoMessage() {}

func (x *PollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollRequest.ProtoReflect.Descriptor instead.
func (*PollRequest) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{8}
}

func (x *PollRequest) GetConsumerID() int64 {
	if x != nil {
		return x.ConsumerID
	}
	return 0
}

type PollResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *PollResponse) Reset() {
	*x = PollResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollResponse) ProtoMessage() {}

func (x *PollResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollResponse.ProtoReflect.Descriptor instead.
func (*PollResponse) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{9}
}

func (x *PollResponse) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quicq_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quicq_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_proto_quicq_proto_rawDescGZIP(), []int{10}
}

func (x *Record) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Record) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_proto_quicq_proto protoreflect.FileDescriptor

var file_proto_quicq_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x22, 0x87, 0x03,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71,
	0x75, 0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x12,
	0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x71,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x12, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x19,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x19, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x39, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x0b, 0x70, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x98, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x71, 0x75, 0x69,
	0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x4b, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x71, 0x75,
	0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x11, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x13, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x71, 0x75, 0x69,
	0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x13, 0x75, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x66, 0x0a, 0x1a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x1a, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x6f, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x48, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x4a, 0x0a, 0x12,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x3b, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x22, 0x2b, 0x0a, 0x13, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x34, 0x0a,
	0x1a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x73, 0x22, 0x2d, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x3a, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x34,
	0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x86, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x55,
	0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43,
	0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x4f, 0x4c, 0x4c, 0x10, 0x04, 0x2a, 0x8c, 0x01,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14,
	0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43,
	0x52, 0x49, 0x42, 0x45, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x5f, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x5f, 0x4d,
	0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x4c, 0x10, 0x04, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x65, 0x64, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x71, 0x2f,
	0x76, 0x31, 0x3b, 0x71, 0x75, 0x69, 0x63, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quicq_proto_rawDescOnce sync.Once
	file_proto_quicq_proto_rawDescData = file_proto_quicq_proto_rawDesc
)

func file_proto_quicq_proto_rawDescGZIP() []byte {
	file_proto_quicq_proto_rawDescOnce.Do(func() {
		file_proto_quicq_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quicq_proto_rawDescData)
	})
	return file_proto_quicq_proto_rawDescData
}

var file_proto_quicq_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_quicq_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_quicq_proto_goTypes = []interface{}{
	(RequestType)(0),                   // 0: quicq.v1.RequestType
	(ResponseType)(0),                  // 1: quicq.v1.ResponseType
	(*Request)(nil),                    // 2: quicq.v1.Request
	(*Response)(nil),                   // 3: quicq.v1.Response
	(*SubscribeRequest)(nil),           // 4: quicq.v1.SubscribeRequest
	(*UnsubscribeRequest)(nil),         // 5: quicq.v1.UnsubscribeRequest
	(*FetchTopicMetadataRequest)(nil),  // 6: quicq.v1.FetchTopicMetadataRequest
	(*SubscribeResponse)(nil),          // 7: quicq.v1.SubscribeResponse
	(*UnsubscribeResponse)(nil),        // 8: quicq.v1.UnsubscribeResponse
	(*FetchTopicMetadataResponse)(nil), // 9: quicq.v1.FetchTopicMetadataResponse
	(*PollRequest)(nil),                // 10: quicq.v1.PollRequest
	(*PollResponse)(nil),               // 11: quicq.v1.PollResponse
	(*Record)(nil),                     // 12: quicq.v1.Record
}
var file_proto_quicq_proto_depIdxs = []int32{
	0,  // 0: quicq.v1.Request.requestType:type_name -> quicq.v1.RequestType
	4,  // 1: quicq.v1.Request.subscribeRequest:type_name -> quicq.v1.SubscribeRequest
	5,  // 2: quicq.v1.Request.unsubscribeRequest:type_name -> quicq.v1.UnsubscribeRequest
	6,  // 3: quicq.v1.Request.fetchTopicMetadataRequest:type_name -> quicq.v1.FetchTopicMetadataRequest
	10, // 4: quicq.v1.Request.pollRequest:type_name -> quicq.v1.PollRequest
	1,  // 5: quicq.v1.Response.responseType:type_name -> quicq.v1.ResponseType
	7,  // 6: quicq.v1.Response.subscribeResponse:type_name -> quicq.v1.SubscribeResponse
	8,  // 7: quicq.v1.Response.unsubscribeResponse:type_name -> quicq.v1.UnsubscribeResponse
	9,  // 8: quicq.v1.Response.fetchTopicMetadataResponse:type_name -> quicq.v1.FetchTopicMetadataResponse
	11, // 9: quicq.v1.Response.pollResponse:type_name -> quicq.v1.PollResponse
	12, // 10: quicq.v1.PollResponse.records:type_name -> quicq.v1.Record
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_proto_quicq_proto_init() }
func file_proto_quicq_proto_init() {
	if File_proto_quicq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quicq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_quicq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_quicq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_proto_quicq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_proto_quicq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTopicMetadataRequest); i {
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
		file_proto_quicq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_proto_quicq_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeResponse); i {
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
		file_proto_quicq_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTopicMetadataResponse); i {
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
		file_proto_quicq_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollRequest); i {
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
		file_proto_quicq_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollResponse); i {
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
		file_proto_quicq_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
	file_proto_quicq_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_SubscribeRequest)(nil),
		(*Request_UnsubscribeRequest)(nil),
		(*Request_FetchTopicMetadataRequest)(nil),
		(*Request_PollRequest)(nil),
	}
	file_proto_quicq_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_SubscribeResponse)(nil),
		(*Response_UnsubscribeResponse)(nil),
		(*Response_FetchTopicMetadataResponse)(nil),
		(*Response_PollResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_quicq_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_quicq_proto_goTypes,
		DependencyIndexes: file_proto_quicq_proto_depIdxs,
		EnumInfos:         file_proto_quicq_proto_enumTypes,
		MessageInfos:      file_proto_quicq_proto_msgTypes,
	}.Build()
	File_proto_quicq_proto = out.File
	file_proto_quicq_proto_rawDesc = nil
	file_proto_quicq_proto_goTypes = nil
	file_proto_quicq_proto_depIdxs = nil
}
