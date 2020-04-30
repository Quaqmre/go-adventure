// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: message.proto

package main

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*Message_Hello
	//	*Message_Pong
	//	*Message_Space
	Content isMessage_Content `protobuf_oneof:"content"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Message) GetHello() *Hello {
	if x, ok := x.GetContent().(*Message_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *Message) GetPong() *Pong {
	if x, ok := x.GetContent().(*Message_Pong); ok {
		return x.Pong
	}
	return nil
}

func (x *Message) GetSpace() *Space {
	if x, ok := x.GetContent().(*Message_Space); ok {
		return x.Space
	}
	return nil
}

type isMessage_Content interface {
	isMessage_Content()
}

type Message_Hello struct {
	Hello *Hello `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type Message_Pong struct {
	Pong *Pong `protobuf:"bytes,2,opt,name=pong,proto3,oneof"`
}

type Message_Space struct {
	Space *Space `protobuf:"bytes,3,opt,name=space,proto3,oneof"`
}

func (*Message_Hello) isMessage_Content() {}

func (*Message_Pong) isMessage_Content() {}

func (*Message_Space) isMessage_Content() {}

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyId         uint32            `protobuf:"varint,1,opt,name=myId,proto3" json:"myId,omitempty"`
	IdToUsername map[uint32]string `protobuf:"bytes,2,rep,name=idToUsername,proto3" json:"idToUsername,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Hello) GetMyId() uint32 {
	if x != nil {
		return x.MyId
	}
	return 0
}

func (x *Hello) GetIdToUsername() map[uint32]string {
	if x != nil {
		return x.IdToUsername
	}
	return nil
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *Pong) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Space struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhysicsFrameID uint32   `protobuf:"varint,1,opt,name=physicsFrameID,proto3" json:"physicsFrameID,omitempty"`
	Spaceships     []int32  `protobuf:"varint,2,rep,packed,name=spaceships,proto3" json:"spaceships,omitempty"`
	Asteroids      []string `protobuf:"bytes,3,rep,name=asteroids,proto3" json:"asteroids,omitempty"`
}

func (x *Space) Reset() {
	*x = Space{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Space) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Space) ProtoMessage() {}

func (x *Space) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Space.ProtoReflect.Descriptor instead.
func (*Space) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *Space) GetPhysicsFrameID() uint32 {
	if x != nil {
		return x.PhysicsFrameID
	}
	return 0
}

func (x *Space) GetSpaceships() []int32 {
	if x != nil {
		return x.Spaceships
	}
	return nil
}

func (x *Space) GetAsteroids() []string {
	if x != nil {
		return x.Asteroids
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x00, 0x52,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x6e, 0x67,
	0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x05, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x6d, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0c, 0x69, 0x64, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x49, 0x64, 0x54, 0x6f, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x69, 0x64, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x3f, 0x0a, 0x11, 0x49, 0x64, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x16, 0x0a, 0x04, 0x50, 0x6f,
	0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x64, 0x22, 0x6d, 0x0a, 0x05, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x73, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x73, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6f, 0x69, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x74, 0x65, 0x72, 0x6f, 0x69, 0x64,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_message_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: main.Message
	(*Hello)(nil),   // 1: main.Hello
	(*Pong)(nil),    // 2: main.Pong
	(*Space)(nil),   // 3: main.Space
	nil,             // 4: main.Hello.IdToUsernameEntry
}
var file_message_proto_depIdxs = []int32{
	1, // 0: main.Message.hello:type_name -> main.Hello
	2, // 1: main.Message.pong:type_name -> main.Pong
	3, // 2: main.Message.space:type_name -> main.Space
	4, // 3: main.Hello.idToUsername:type_name -> main.Hello.IdToUsernameEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Space); i {
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
	file_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Hello)(nil),
		(*Message_Pong)(nil),
		(*Message_Space)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}