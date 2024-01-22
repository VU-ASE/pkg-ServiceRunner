// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: tuningstate.proto

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

type TuningState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp         uint64                   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DynamicParameters []*TuningState_Parameter `protobuf:"bytes,2,rep,name=dynamicParameters,proto3" json:"dynamicParameters,omitempty"`
}

func (x *TuningState) Reset() {
	*x = TuningState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuningstate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState) ProtoMessage() {}

func (x *TuningState) ProtoReflect() protoreflect.Message {
	mi := &file_tuningstate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState.ProtoReflect.Descriptor instead.
func (*TuningState) Descriptor() ([]byte, []int) {
	return file_tuningstate_proto_rawDescGZIP(), []int{0}
}

func (x *TuningState) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TuningState) GetDynamicParameters() []*TuningState_Parameter {
	if x != nil {
		return x.DynamicParameters
	}
	return nil
}

type TuningState_Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Parameter:
	//
	//	*TuningState_Parameter_Float
	//	*TuningState_Parameter_Int
	//	*TuningState_Parameter_String_
	Parameter isTuningState_Parameter_Parameter `protobuf_oneof:"parameter"`
}

func (x *TuningState_Parameter) Reset() {
	*x = TuningState_Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuningstate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningState_Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState_Parameter) ProtoMessage() {}

func (x *TuningState_Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_tuningstate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState_Parameter.ProtoReflect.Descriptor instead.
func (*TuningState_Parameter) Descriptor() ([]byte, []int) {
	return file_tuningstate_proto_rawDescGZIP(), []int{0, 0}
}

func (m *TuningState_Parameter) GetParameter() isTuningState_Parameter_Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (x *TuningState_Parameter) GetFloat() *TuningState_Parameter_FloatParameter {
	if x, ok := x.GetParameter().(*TuningState_Parameter_Float); ok {
		return x.Float
	}
	return nil
}

func (x *TuningState_Parameter) GetInt() *TuningState_Parameter_IntParameter {
	if x, ok := x.GetParameter().(*TuningState_Parameter_Int); ok {
		return x.Int
	}
	return nil
}

func (x *TuningState_Parameter) GetString_() *TuningState_Parameter_StringParameter {
	if x, ok := x.GetParameter().(*TuningState_Parameter_String_); ok {
		return x.String_
	}
	return nil
}

type isTuningState_Parameter_Parameter interface {
	isTuningState_Parameter_Parameter()
}

type TuningState_Parameter_Float struct {
	Float *TuningState_Parameter_FloatParameter `protobuf:"bytes,1,opt,name=float,proto3,oneof"`
}

type TuningState_Parameter_Int struct {
	Int *TuningState_Parameter_IntParameter `protobuf:"bytes,2,opt,name=int,proto3,oneof"`
}

type TuningState_Parameter_String_ struct {
	String_ *TuningState_Parameter_StringParameter `protobuf:"bytes,3,opt,name=string,proto3,oneof"`
}

func (*TuningState_Parameter_Float) isTuningState_Parameter_Parameter() {}

func (*TuningState_Parameter_Int) isTuningState_Parameter_Parameter() {}

func (*TuningState_Parameter_String_) isTuningState_Parameter_Parameter() {}

type TuningState_Parameter_FloatParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TuningState_Parameter_FloatParameter) Reset() {
	*x = TuningState_Parameter_FloatParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuningstate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningState_Parameter_FloatParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState_Parameter_FloatParameter) ProtoMessage() {}

func (x *TuningState_Parameter_FloatParameter) ProtoReflect() protoreflect.Message {
	mi := &file_tuningstate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState_Parameter_FloatParameter.ProtoReflect.Descriptor instead.
func (*TuningState_Parameter_FloatParameter) Descriptor() ([]byte, []int) {
	return file_tuningstate_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *TuningState_Parameter_FloatParameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TuningState_Parameter_FloatParameter) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TuningState_Parameter_IntParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TuningState_Parameter_IntParameter) Reset() {
	*x = TuningState_Parameter_IntParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuningstate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningState_Parameter_IntParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState_Parameter_IntParameter) ProtoMessage() {}

func (x *TuningState_Parameter_IntParameter) ProtoReflect() protoreflect.Message {
	mi := &file_tuningstate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState_Parameter_IntParameter.ProtoReflect.Descriptor instead.
func (*TuningState_Parameter_IntParameter) Descriptor() ([]byte, []int) {
	return file_tuningstate_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *TuningState_Parameter_IntParameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TuningState_Parameter_IntParameter) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TuningState_Parameter_StringParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TuningState_Parameter_StringParameter) Reset() {
	*x = TuningState_Parameter_StringParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tuningstate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningState_Parameter_StringParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningState_Parameter_StringParameter) ProtoMessage() {}

func (x *TuningState_Parameter_StringParameter) ProtoReflect() protoreflect.Message {
	mi := &file_tuningstate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningState_Parameter_StringParameter.ProtoReflect.Descriptor instead.
func (*TuningState_Parameter_StringParameter) Descriptor() ([]byte, []int) {
	return file_tuningstate_proto_rawDescGZIP(), []int{0, 0, 2}
}

func (x *TuningState_Parameter_StringParameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TuningState_Parameter_StringParameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_tuningstate_proto protoreflect.FileDescriptor

var file_tuningstate_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73,
	0x67, 0x73, 0x22, 0xab, 0x04, 0x0a, 0x0b, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x52, 0x0a, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x11, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x1a, 0xa9, 0x03, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x4b, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67,
	0x73, 0x2e, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12,
	0x45, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x54, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x49, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x38, 0x0a, 0x0e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x36, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x39, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x42, 0x13, 0x5a, 0x11, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x5f, 0x6d, 0x73, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tuningstate_proto_rawDescOnce sync.Once
	file_tuningstate_proto_rawDescData = file_tuningstate_proto_rawDesc
)

func file_tuningstate_proto_rawDescGZIP() []byte {
	file_tuningstate_proto_rawDescOnce.Do(func() {
		file_tuningstate_proto_rawDescData = protoimpl.X.CompressGZIP(file_tuningstate_proto_rawDescData)
	})
	return file_tuningstate_proto_rawDescData
}

var file_tuningstate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tuningstate_proto_goTypes = []interface{}{
	(*TuningState)(nil),                           // 0: protobuf_msgs.TuningState
	(*TuningState_Parameter)(nil),                 // 1: protobuf_msgs.TuningState.Parameter
	(*TuningState_Parameter_FloatParameter)(nil),  // 2: protobuf_msgs.TuningState.Parameter.FloatParameter
	(*TuningState_Parameter_IntParameter)(nil),    // 3: protobuf_msgs.TuningState.Parameter.IntParameter
	(*TuningState_Parameter_StringParameter)(nil), // 4: protobuf_msgs.TuningState.Parameter.StringParameter
}
var file_tuningstate_proto_depIdxs = []int32{
	1, // 0: protobuf_msgs.TuningState.dynamicParameters:type_name -> protobuf_msgs.TuningState.Parameter
	2, // 1: protobuf_msgs.TuningState.Parameter.float:type_name -> protobuf_msgs.TuningState.Parameter.FloatParameter
	3, // 2: protobuf_msgs.TuningState.Parameter.int:type_name -> protobuf_msgs.TuningState.Parameter.IntParameter
	4, // 3: protobuf_msgs.TuningState.Parameter.string:type_name -> protobuf_msgs.TuningState.Parameter.StringParameter
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tuningstate_proto_init() }
func file_tuningstate_proto_init() {
	if File_tuningstate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tuningstate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TuningState); i {
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
		file_tuningstate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TuningState_Parameter); i {
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
		file_tuningstate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TuningState_Parameter_FloatParameter); i {
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
		file_tuningstate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TuningState_Parameter_IntParameter); i {
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
		file_tuningstate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TuningState_Parameter_StringParameter); i {
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
	file_tuningstate_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TuningState_Parameter_Float)(nil),
		(*TuningState_Parameter_Int)(nil),
		(*TuningState_Parameter_String_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tuningstate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tuningstate_proto_goTypes,
		DependencyIndexes: file_tuningstate_proto_depIdxs,
		MessageInfos:      file_tuningstate_proto_msgTypes,
	}.Build()
	File_tuningstate_proto = out.File
	file_tuningstate_proto_rawDesc = nil
	file_tuningstate_proto_goTypes = nil
	file_tuningstate_proto_depIdxs = nil
}
