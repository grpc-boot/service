// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: gateway.proto

package pb

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

type GwReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *GwInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GwReply) Reset() {
	*x = GwReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GwReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GwReply) ProtoMessage() {}

func (x *GwReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GwReply.ProtoReflect.Descriptor instead.
func (*GwReply) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *GwReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GwReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GwReply) GetData() *GwInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GwInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qps        int32         `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
	Total      uint64        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	MethodList []*MethodInfo `protobuf:"bytes,3,rep,name=method_list,json=methodList,proto3" json:"method_list,omitempty"`
}

func (x *GwInfo) Reset() {
	*x = GwInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GwInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GwInfo) ProtoMessage() {}

func (x *GwInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GwInfo.ProtoReflect.Descriptor instead.
func (*GwInfo) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GwInfo) GetQps() int32 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *GwInfo) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GwInfo) GetMethodList() []*MethodInfo {
	if x != nil {
		return x.MethodList
	}
	return nil
}

type MethodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path        string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	SecondLimit int32            `protobuf:"varint,3,opt,name=second_limit,json=secondLimit,proto3" json:"second_limit,omitempty"`
	Qps         int32            `protobuf:"varint,4,opt,name=qps,proto3" json:"qps,omitempty"`
	Total       uint64           `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	Avg         string           `protobuf:"bytes,6,opt,name=avg,proto3" json:"avg,omitempty"`
	Min         string           `protobuf:"bytes,7,opt,name=min,proto3" json:"min,omitempty"`
	Max         string           `protobuf:"bytes,8,opt,name=max,proto3" json:"max,omitempty"`
	Line_90     string           `protobuf:"bytes,9,opt,name=line_90,json=line90,proto3" json:"line_90,omitempty"`
	Line_95     string           `protobuf:"bytes,10,opt,name=line_95,json=line95,proto3" json:"line_95,omitempty"`
	CodeMap     map[int32]uint64 `protobuf:"bytes,11,rep,name=code_map,json=codeMap,proto3" json:"code_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *MethodInfo) Reset() {
	*x = MethodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodInfo) ProtoMessage() {}

func (x *MethodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodInfo.ProtoReflect.Descriptor instead.
func (*MethodInfo) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *MethodInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MethodInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MethodInfo) GetSecondLimit() int32 {
	if x != nil {
		return x.SecondLimit
	}
	return 0
}

func (x *MethodInfo) GetQps() int32 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *MethodInfo) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MethodInfo) GetAvg() string {
	if x != nil {
		return x.Avg
	}
	return ""
}

func (x *MethodInfo) GetMin() string {
	if x != nil {
		return x.Min
	}
	return ""
}

func (x *MethodInfo) GetMax() string {
	if x != nil {
		return x.Max
	}
	return ""
}

func (x *MethodInfo) GetLine_90() string {
	if x != nil {
		return x.Line_90
	}
	return ""
}

func (x *MethodInfo) GetLine_95() string {
	if x != nil {
		return x.Line_95
	}
	return ""
}

func (x *MethodInfo) GetCodeMap() map[int32]uint64 {
	if x != nil {
		return x.CodeMap
	}
	return nil
}

var File_gateway_proto protoreflect.FileDescriptor

var file_gateway_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4f, 0x0a, 0x07, 0x47, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x61, 0x0a, 0x06, 0x47, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x71,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x70, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0xdb, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x71, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x70, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x76, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61,
	0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x39, 0x30, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x69, 0x6e, 0x65, 0x39, 0x30, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x39, 0x35,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x39, 0x35, 0x12, 0x36,
	0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63,
	0x6f, 0x64, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0x36, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x2b, 0x0a,
	0x02, 0x47, 0x77, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_proto_rawDescOnce sync.Once
	file_gateway_proto_rawDescData = file_gateway_proto_rawDesc
)

func file_gateway_proto_rawDescGZIP() []byte {
	file_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_proto_rawDescData)
	})
	return file_gateway_proto_rawDescData
}

var file_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gateway_proto_goTypes = []interface{}{
	(*GwReply)(nil),       // 0: pb.GwReply
	(*GwInfo)(nil),        // 1: pb.GwInfo
	(*MethodInfo)(nil),    // 2: pb.MethodInfo
	nil,                   // 3: pb.MethodInfo.CodeMapEntry
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_gateway_proto_depIdxs = []int32{
	1, // 0: pb.GwReply.data:type_name -> pb.GwInfo
	2, // 1: pb.GwInfo.method_list:type_name -> pb.MethodInfo
	3, // 2: pb.MethodInfo.code_map:type_name -> pb.MethodInfo.CodeMapEntry
	4, // 3: pb.Gateway.Gw:input_type -> google.protobuf.Empty
	0, // 4: pb.Gateway.Gw:output_type -> pb.GwReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gateway_proto_init() }
func file_gateway_proto_init() {
	if File_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GwReply); i {
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
		file_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GwInfo); i {
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
		file_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodInfo); i {
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
			RawDescriptor: file_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_proto_msgTypes,
	}.Build()
	File_gateway_proto = out.File
	file_gateway_proto_rawDesc = nil
	file_gateway_proto_goTypes = nil
	file_gateway_proto_depIdxs = nil
}