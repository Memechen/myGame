// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: proto/messageId.proto

package messageId

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

type MessageId int32

const (
	MessageId_None    MessageId = 0
	MessageId_CSLogin MessageId = 50001
	MessageId_SCLogin MessageId = 50002
	// player
	MessageId_CSAddFriend   MessageId = 100001
	MessageId_SCAddFriend   MessageId = 100002
	MessageId_CSDelFriend   MessageId = 100003
	MessageId_SCDelFriend   MessageId = 100004
	MessageId_CSSendChatMsg MessageId = 100005
	MessageId_SCSendChatMsg MessageId = 100006
)

// Enum value maps for MessageId.
var (
	MessageId_name = map[int32]string{
		0:      "None",
		50001:  "CSLogin",
		50002:  "SCLogin",
		100001: "CSAddFriend",
		100002: "SCAddFriend",
		100003: "CSDelFriend",
		100004: "SCDelFriend",
		100005: "CSSendChatMsg",
		100006: "SCSendChatMsg",
	}
	MessageId_value = map[string]int32{
		"None":          0,
		"CSLogin":       50001,
		"SCLogin":       50002,
		"CSAddFriend":   100001,
		"SCAddFriend":   100002,
		"CSDelFriend":   100003,
		"SCDelFriend":   100004,
		"CSSendChatMsg": 100005,
		"SCSendChatMsg": 100006,
	}
)

func (x MessageId) Enum() *MessageId {
	p := new(MessageId)
	*p = x
	return p
}

func (x MessageId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageId) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_messageId_proto_enumTypes[0].Descriptor()
}

func (MessageId) Type() protoreflect.EnumType {
	return &file_proto_messageId_proto_enumTypes[0]
}

func (x MessageId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageId.Descriptor instead.
func (MessageId) EnumDescriptor() ([]byte, []int) {
	return file_proto_messageId_proto_rawDescGZIP(), []int{0}
}

var File_proto_messageId_proto protoreflect.FileDescriptor

var file_proto_messageId_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x2a, 0xa9, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x07, 0x43, 0x53,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0xd1, 0x86, 0x03, 0x12, 0x0d, 0x0a, 0x07, 0x53, 0x43, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x10, 0xd2, 0x86, 0x03, 0x12, 0x11, 0x0a, 0x0b, 0x43, 0x53, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa1, 0x8d, 0x06, 0x12, 0x11, 0x0a, 0x0b, 0x53,
	0x43, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa2, 0x8d, 0x06, 0x12, 0x11,
	0x0a, 0x0b, 0x43, 0x53, 0x44, 0x65, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x10, 0xa3, 0x8d,
	0x06, 0x12, 0x11, 0x0a, 0x0b, 0x53, 0x43, 0x44, 0x65, 0x6c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x10, 0xa4, 0x8d, 0x06, 0x12, 0x13, 0x0a, 0x0d, 0x43, 0x53, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x73, 0x67, 0x10, 0xa5, 0x8d, 0x06, 0x12, 0x13, 0x0a, 0x0d, 0x53, 0x43, 0x53,
	0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x10, 0xa6, 0x8d, 0x06, 0x42, 0x0c,
	0x5a, 0x0a, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_messageId_proto_rawDescOnce sync.Once
	file_proto_messageId_proto_rawDescData = file_proto_messageId_proto_rawDesc
)

func file_proto_messageId_proto_rawDescGZIP() []byte {
	file_proto_messageId_proto_rawDescOnce.Do(func() {
		file_proto_messageId_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messageId_proto_rawDescData)
	})
	return file_proto_messageId_proto_rawDescData
}

var file_proto_messageId_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_messageId_proto_goTypes = []interface{}{
	(MessageId)(0), // 0: messageId.MessageId
}
var file_proto_messageId_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_messageId_proto_init() }
func file_proto_messageId_proto_init() {
	if File_proto_messageId_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_messageId_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_messageId_proto_goTypes,
		DependencyIndexes: file_proto_messageId_proto_depIdxs,
		EnumInfos:         file_proto_messageId_proto_enumTypes,
	}.Build()
	File_proto_messageId_proto = out.File
	file_proto_messageId_proto_rawDesc = nil
	file_proto_messageId_proto_goTypes = nil
	file_proto_messageId_proto_depIdxs = nil
}
