// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: backend/v1/backend.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RegisterConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  string      `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Connection *Connection `protobuf:"bytes,2,opt,name=connection,proto3" json:"connection,omitempty"`
}

func (x *RegisterConnectionRequest) Reset() {
	*x = RegisterConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterConnectionRequest) ProtoMessage() {}

func (x *RegisterConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterConnectionRequest.ProtoReflect.Descriptor instead.
func (*RegisterConnectionRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_backend_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterConnectionRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *RegisterConnectionRequest) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

type GetConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Hostname  string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *GetConnectionRequest) Reset() {
	*x = GetConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionRequest) ProtoMessage() {}

func (x *GetConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetConnectionRequest) Descriptor() ([]byte, []int) {
	return file_backend_v1_backend_proto_rawDescGZIP(), []int{1}
}

func (x *GetConnectionRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetConnectionRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type GetConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *Connection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
}

func (x *GetConnectionResponse) Reset() {
	*x = GetConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_v1_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionResponse) ProtoMessage() {}

func (x *GetConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_v1_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionResponse.ProtoReflect.Descriptor instead.
func (*GetConnectionResponse) Descriptor() ([]byte, []int) {
	return file_backend_v1_backend_proto_rawDescGZIP(), []int{2}
}

func (x *GetConnectionResponse) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

var File_backend_v1_backend_proto protoreflect.FileDescriptor

var file_backend_v1_backend_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x1a, 0x19, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x19, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x4c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xab, 0x01,
	0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x50, 0x0a, 0x12, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x22, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_v1_backend_proto_rawDescOnce sync.Once
	file_backend_v1_backend_proto_rawDescData = file_backend_v1_backend_proto_rawDesc
)

func file_backend_v1_backend_proto_rawDescGZIP() []byte {
	file_backend_v1_backend_proto_rawDescOnce.Do(func() {
		file_backend_v1_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_v1_backend_proto_rawDescData)
	})
	return file_backend_v1_backend_proto_rawDescData
}

var file_backend_v1_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_backend_v1_backend_proto_goTypes = []interface{}{
	(*RegisterConnectionRequest)(nil), // 0: backend.RegisterConnectionRequest
	(*GetConnectionRequest)(nil),      // 1: backend.GetConnectionRequest
	(*GetConnectionResponse)(nil),     // 2: backend.GetConnectionResponse
	(*Connection)(nil),                // 3: backend.Connection
	(*emptypb.Empty)(nil),             // 4: google.protobuf.Empty
}
var file_backend_v1_backend_proto_depIdxs = []int32{
	3, // 0: backend.RegisterConnectionRequest.connection:type_name -> backend.Connection
	3, // 1: backend.GetConnectionResponse.connection:type_name -> backend.Connection
	0, // 2: backend.Backend.RegisterConnection:input_type -> backend.RegisterConnectionRequest
	1, // 3: backend.Backend.GetConnection:input_type -> backend.GetConnectionRequest
	4, // 4: backend.Backend.RegisterConnection:output_type -> google.protobuf.Empty
	2, // 5: backend.Backend.GetConnection:output_type -> backend.GetConnectionResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_backend_v1_backend_proto_init() }
func file_backend_v1_backend_proto_init() {
	if File_backend_v1_backend_proto != nil {
		return
	}
	file_backend_v1_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_backend_v1_backend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterConnectionRequest); i {
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
		file_backend_v1_backend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectionRequest); i {
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
		file_backend_v1_backend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConnectionResponse); i {
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
			RawDescriptor: file_backend_v1_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_v1_backend_proto_goTypes,
		DependencyIndexes: file_backend_v1_backend_proto_depIdxs,
		MessageInfos:      file_backend_v1_backend_proto_msgTypes,
	}.Build()
	File_backend_v1_backend_proto = out.File
	file_backend_v1_backend_proto_rawDesc = nil
	file_backend_v1_backend_proto_goTypes = nil
	file_backend_v1_backend_proto_depIdxs = nil
}
