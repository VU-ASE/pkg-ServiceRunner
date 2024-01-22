// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: servicediscovery.proto

package protobuf_msgs

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

type ServiceDiscovery_Information_Status int32

const (
	ServiceDiscovery_Information_UNKNOWN ServiceDiscovery_Information_Status = 0
	ServiceDiscovery_Information_RUNNING ServiceDiscovery_Information_Status = 1
	ServiceDiscovery_Information_STOPPED ServiceDiscovery_Information_Status = 2
)

// Enum value maps for ServiceDiscovery_Information_Status.
var (
	ServiceDiscovery_Information_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "RUNNING",
		2: "STOPPED",
	}
	ServiceDiscovery_Information_Status_value = map[string]int32{
		"UNKNOWN": 0,
		"RUNNING": 1,
		"STOPPED": 2,
	}
)

func (x ServiceDiscovery_Information_Status) Enum() *ServiceDiscovery_Information_Status {
	p := new(ServiceDiscovery_Information_Status)
	*p = x
	return p
}

func (x ServiceDiscovery_Information_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceDiscovery_Information_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_servicediscovery_proto_enumTypes[0].Descriptor()
}

func (ServiceDiscovery_Information_Status) Type() protoreflect.EnumType {
	return &file_servicediscovery_proto_enumTypes[0]
}

func (x ServiceDiscovery_Information_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceDiscovery_Information_Status.Descriptor instead.
func (ServiceDiscovery_Information_Status) EnumDescriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3, 2, 0}
}

type ServiceDiscovery_Order_OrderType int32

const (
	ServiceDiscovery_Order_STOP          ServiceDiscovery_Order_OrderType = 0 // will attempt a graceful shutdown
	ServiceDiscovery_Order_RESTART       ServiceDiscovery_Order_OrderType = 1 // will attempt a graceful shutdown and restart
	ServiceDiscovery_Order_KILL          ServiceDiscovery_Order_OrderType = 2 // will kill the service immediately
	ServiceDiscovery_Order_FORCE_RESTART ServiceDiscovery_Order_OrderType = 3 // will kill the service immediately and restart
)

// Enum value maps for ServiceDiscovery_Order_OrderType.
var (
	ServiceDiscovery_Order_OrderType_name = map[int32]string{
		0: "STOP",
		1: "RESTART",
		2: "KILL",
		3: "FORCE_RESTART",
	}
	ServiceDiscovery_Order_OrderType_value = map[string]int32{
		"STOP":          0,
		"RESTART":       1,
		"KILL":          2,
		"FORCE_RESTART": 3,
	}
)

func (x ServiceDiscovery_Order_OrderType) Enum() *ServiceDiscovery_Order_OrderType {
	p := new(ServiceDiscovery_Order_OrderType)
	*p = x
	return p
}

func (x ServiceDiscovery_Order_OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceDiscovery_Order_OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_servicediscovery_proto_enumTypes[1].Descriptor()
}

func (ServiceDiscovery_Order_OrderType) Type() protoreflect.EnumType {
	return &file_servicediscovery_proto_enumTypes[1]
}

func (x ServiceDiscovery_Order_OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceDiscovery_Order_OrderType.Descriptor instead.
func (ServiceDiscovery_Order_OrderType) EnumDescriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3, 3, 0}
}

// Used to identify a service within the system
type ServiceIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pid  int32  `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *ServiceIdentifier) Reset() {
	*x = ServiceIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceIdentifier) ProtoMessage() {}

func (x *ServiceIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceIdentifier.ProtoReflect.Descriptor instead.
func (*ServiceIdentifier) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceIdentifier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceIdentifier) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

// An endpoint that is made available by a service
type ServiceEndpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // the identifier to select this endpoint
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // the zmq address to connect to
}

func (x *ServiceEndpoint) Reset() {
	*x = ServiceEndpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceEndpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceEndpoint) ProtoMessage() {}

func (x *ServiceEndpoint) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceEndpoint.ProtoReflect.Descriptor instead.
func (*ServiceEndpoint) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceEndpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceEndpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// A description of a service
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier *ServiceIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Endpoints  []*ServiceEndpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{2}
}

func (x *Service) GetIdentifier() *ServiceIdentifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *Service) GetEndpoints() []*ServiceEndpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

// A message sent back and forth for the service discovery
type ServiceDiscovery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*ServiceDiscovery_Registration_
	//	*ServiceDiscovery_InformationRequest_
	//	*ServiceDiscovery_Information_
	//	*ServiceDiscovery_Order_
	//	*ServiceDiscovery_Error_
	Content isServiceDiscovery_Content `protobuf_oneof:"content"`
}

func (x *ServiceDiscovery) Reset() {
	*x = ServiceDiscovery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDiscovery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDiscovery) ProtoMessage() {}

func (x *ServiceDiscovery) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDiscovery.ProtoReflect.Descriptor instead.
func (*ServiceDiscovery) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3}
}

func (m *ServiceDiscovery) GetContent() isServiceDiscovery_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ServiceDiscovery) GetRegistration() *ServiceDiscovery_Registration {
	if x, ok := x.GetContent().(*ServiceDiscovery_Registration_); ok {
		return x.Registration
	}
	return nil
}

func (x *ServiceDiscovery) GetInformationRequest() *ServiceDiscovery_InformationRequest {
	if x, ok := x.GetContent().(*ServiceDiscovery_InformationRequest_); ok {
		return x.InformationRequest
	}
	return nil
}

func (x *ServiceDiscovery) GetInformation() *ServiceDiscovery_Information {
	if x, ok := x.GetContent().(*ServiceDiscovery_Information_); ok {
		return x.Information
	}
	return nil
}

func (x *ServiceDiscovery) GetOrder() *ServiceDiscovery_Order {
	if x, ok := x.GetContent().(*ServiceDiscovery_Order_); ok {
		return x.Order
	}
	return nil
}

func (x *ServiceDiscovery) GetError() *ServiceDiscovery_Error {
	if x, ok := x.GetContent().(*ServiceDiscovery_Error_); ok {
		return x.Error
	}
	return nil
}

type isServiceDiscovery_Content interface {
	isServiceDiscovery_Content()
}

type ServiceDiscovery_Registration_ struct {
	Registration *ServiceDiscovery_Registration `protobuf:"bytes,1,opt,name=registration,proto3,oneof"`
}

type ServiceDiscovery_InformationRequest_ struct {
	InformationRequest *ServiceDiscovery_InformationRequest `protobuf:"bytes,2,opt,name=information_request,json=informationRequest,proto3,oneof"`
}

type ServiceDiscovery_Information_ struct {
	Information *ServiceDiscovery_Information `protobuf:"bytes,3,opt,name=information,proto3,oneof"`
}

type ServiceDiscovery_Order_ struct {
	Order *ServiceDiscovery_Order `protobuf:"bytes,4,opt,name=order,proto3,oneof"`
}

type ServiceDiscovery_Error_ struct {
	Error *ServiceDiscovery_Error `protobuf:"bytes,5,opt,name=error,proto3,oneof"`
}

func (*ServiceDiscovery_Registration_) isServiceDiscovery_Content() {}

func (*ServiceDiscovery_InformationRequest_) isServiceDiscovery_Content() {}

func (*ServiceDiscovery_Information_) isServiceDiscovery_Content() {}

func (*ServiceDiscovery_Order_) isServiceDiscovery_Content() {}

func (*ServiceDiscovery_Error_) isServiceDiscovery_Content() {}

// When a service registers itself, it sends a ServiceRegistration
type ServiceDiscovery_Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *ServiceDiscovery_Registration) Reset() {
	*x = ServiceDiscovery_Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDiscovery_Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDiscovery_Registration) ProtoMessage() {}

func (x *ServiceDiscovery_Registration) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDiscovery_Registration.ProtoReflect.Descriptor instead.
func (*ServiceDiscovery_Registration) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ServiceDiscovery_Registration) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

// When a service requests information about other services, it sends an InformationRequest message
type ServiceDiscovery_InformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requested *ServiceIdentifier `protobuf:"bytes,1,opt,name=requested,proto3" json:"requested,omitempty"`
}

func (x *ServiceDiscovery_InformationRequest) Reset() {
	*x = ServiceDiscovery_InformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDiscovery_InformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDiscovery_InformationRequest) ProtoMessage() {}

func (x *ServiceDiscovery_InformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDiscovery_InformationRequest.ProtoReflect.Descriptor instead.
func (*ServiceDiscovery_InformationRequest) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ServiceDiscovery_InformationRequest) GetRequested() *ServiceIdentifier {
	if x != nil {
		return x.Requested
	}
	return nil
}

// When the system manager sends information about a service, it sends an Information message
// Also used to broadcast registrations to all services
type ServiceDiscovery_Information struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *Service                            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Status  ServiceDiscovery_Information_Status `protobuf:"varint,2,opt,name=status,proto3,enum=protobuf_msgs.ServiceDiscovery_Information_Status" json:"status,omitempty"`
}

func (x *ServiceDiscovery_Information) Reset() {
	*x = ServiceDiscovery_Information{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDiscovery_Information) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDiscovery_Information) ProtoMessage() {}

func (x *ServiceDiscovery_Information) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDiscovery_Information.ProtoReflect.Descriptor instead.
func (*ServiceDiscovery_Information) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3, 2}
}

func (x *ServiceDiscovery_Information) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ServiceDiscovery_Information) GetStatus() ServiceDiscovery_Information_Status {
	if x != nil {
		return x.Status
	}
	return ServiceDiscovery_Information_UNKNOWN
}

// The system manager can order services to stop/restart by sending a service order
type ServiceDiscovery_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The service this order is for
	Service *ServiceIdentifier               `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Order   ServiceDiscovery_Order_OrderType `protobuf:"varint,2,opt,name=order,proto3,enum=protobuf_msgs.ServiceDiscovery_Order_OrderType" json:"order,omitempty"`
}

func (x *ServiceDiscovery_Order) Reset() {
	*x = ServiceDiscovery_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDiscovery_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDiscovery_Order) ProtoMessage() {}

func (x *ServiceDiscovery_Order) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDiscovery_Order.ProtoReflect.Descriptor instead.
func (*ServiceDiscovery_Order) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3, 3}
}

func (x *ServiceDiscovery_Order) GetService() *ServiceIdentifier {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ServiceDiscovery_Order) GetOrder() ServiceDiscovery_Order_OrderType {
	if x != nil {
		return x.Order
	}
	return ServiceDiscovery_Order_STOP
}

// A reply to indicate that the order or request did not succeed
type ServiceDiscovery_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceDiscovery_Error) Reset() {
	*x = ServiceDiscovery_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servicediscovery_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDiscovery_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDiscovery_Error) ProtoMessage() {}

func (x *ServiceDiscovery_Error) ProtoReflect() protoreflect.Message {
	mi := &file_servicediscovery_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDiscovery_Error.ProtoReflect.Descriptor instead.
func (*ServiceDiscovery_Error) Descriptor() ([]byte, []int) {
	return file_servicediscovery_proto_rawDescGZIP(), []int{3, 4}
}

var File_servicediscovery_proto protoreflect.FileDescriptor

var file_servicediscovery_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x22, 0x39, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d,
	0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x3c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22,
	0xd5, 0x07, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x12, 0x52, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x65, 0x0a, 0x13, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x12, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4f, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3d, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x3d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x40,
	0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x54, 0x0a, 0x12, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x1a, 0xbc, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50,
	0x50, 0x45, 0x44, 0x10, 0x02, 0x1a, 0xcb, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x3a, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x3f, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4b, 0x49, 0x4c, 0x4c, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x10, 0x03, 0x1a, 0x07, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x73, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servicediscovery_proto_rawDescOnce sync.Once
	file_servicediscovery_proto_rawDescData = file_servicediscovery_proto_rawDesc
)

func file_servicediscovery_proto_rawDescGZIP() []byte {
	file_servicediscovery_proto_rawDescOnce.Do(func() {
		file_servicediscovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_servicediscovery_proto_rawDescData)
	})
	return file_servicediscovery_proto_rawDescData
}

var file_servicediscovery_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_servicediscovery_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_servicediscovery_proto_goTypes = []interface{}{
	(ServiceDiscovery_Information_Status)(0),    // 0: protobuf_msgs.ServiceDiscovery.Information.Status
	(ServiceDiscovery_Order_OrderType)(0),       // 1: protobuf_msgs.ServiceDiscovery.Order.OrderType
	(*ServiceIdentifier)(nil),                   // 2: protobuf_msgs.ServiceIdentifier
	(*ServiceEndpoint)(nil),                     // 3: protobuf_msgs.ServiceEndpoint
	(*Service)(nil),                             // 4: protobuf_msgs.Service
	(*ServiceDiscovery)(nil),                    // 5: protobuf_msgs.ServiceDiscovery
	(*ServiceDiscovery_Registration)(nil),       // 6: protobuf_msgs.ServiceDiscovery.Registration
	(*ServiceDiscovery_InformationRequest)(nil), // 7: protobuf_msgs.ServiceDiscovery.InformationRequest
	(*ServiceDiscovery_Information)(nil),        // 8: protobuf_msgs.ServiceDiscovery.Information
	(*ServiceDiscovery_Order)(nil),              // 9: protobuf_msgs.ServiceDiscovery.Order
	(*ServiceDiscovery_Error)(nil),              // 10: protobuf_msgs.ServiceDiscovery.Error
}
var file_servicediscovery_proto_depIdxs = []int32{
	2,  // 0: protobuf_msgs.Service.identifier:type_name -> protobuf_msgs.ServiceIdentifier
	3,  // 1: protobuf_msgs.Service.endpoints:type_name -> protobuf_msgs.ServiceEndpoint
	6,  // 2: protobuf_msgs.ServiceDiscovery.registration:type_name -> protobuf_msgs.ServiceDiscovery.Registration
	7,  // 3: protobuf_msgs.ServiceDiscovery.information_request:type_name -> protobuf_msgs.ServiceDiscovery.InformationRequest
	8,  // 4: protobuf_msgs.ServiceDiscovery.information:type_name -> protobuf_msgs.ServiceDiscovery.Information
	9,  // 5: protobuf_msgs.ServiceDiscovery.order:type_name -> protobuf_msgs.ServiceDiscovery.Order
	10, // 6: protobuf_msgs.ServiceDiscovery.error:type_name -> protobuf_msgs.ServiceDiscovery.Error
	4,  // 7: protobuf_msgs.ServiceDiscovery.Registration.service:type_name -> protobuf_msgs.Service
	2,  // 8: protobuf_msgs.ServiceDiscovery.InformationRequest.requested:type_name -> protobuf_msgs.ServiceIdentifier
	4,  // 9: protobuf_msgs.ServiceDiscovery.Information.service:type_name -> protobuf_msgs.Service
	0,  // 10: protobuf_msgs.ServiceDiscovery.Information.status:type_name -> protobuf_msgs.ServiceDiscovery.Information.Status
	2,  // 11: protobuf_msgs.ServiceDiscovery.Order.service:type_name -> protobuf_msgs.ServiceIdentifier
	1,  // 12: protobuf_msgs.ServiceDiscovery.Order.order:type_name -> protobuf_msgs.ServiceDiscovery.Order.OrderType
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_servicediscovery_proto_init() }
func file_servicediscovery_proto_init() {
	if File_servicediscovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_servicediscovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceIdentifier); i {
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
		file_servicediscovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceEndpoint); i {
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
		file_servicediscovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_servicediscovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDiscovery); i {
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
		file_servicediscovery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDiscovery_Registration); i {
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
		file_servicediscovery_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDiscovery_InformationRequest); i {
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
		file_servicediscovery_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDiscovery_Information); i {
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
		file_servicediscovery_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDiscovery_Order); i {
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
		file_servicediscovery_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDiscovery_Error); i {
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
	file_servicediscovery_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ServiceDiscovery_Registration_)(nil),
		(*ServiceDiscovery_InformationRequest_)(nil),
		(*ServiceDiscovery_Information_)(nil),
		(*ServiceDiscovery_Order_)(nil),
		(*ServiceDiscovery_Error_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_servicediscovery_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_servicediscovery_proto_goTypes,
		DependencyIndexes: file_servicediscovery_proto_depIdxs,
		EnumInfos:         file_servicediscovery_proto_enumTypes,
		MessageInfos:      file_servicediscovery_proto_msgTypes,
	}.Build()
	File_servicediscovery_proto = out.File
	file_servicediscovery_proto_rawDesc = nil
	file_servicediscovery_proto_goTypes = nil
	file_servicediscovery_proto_depIdxs = nil
}
