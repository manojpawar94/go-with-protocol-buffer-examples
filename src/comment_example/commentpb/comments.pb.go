// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: comment_example/comments.proto

package commentpb

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

type ReactionType int32

const (
	ReactionType_UNKNOWN ReactionType = 0
	ReactionType_LIKE    ReactionType = 1
	ReactionType_DISLIKE ReactionType = 2
	ReactionType_LOVE    ReactionType = 3
	ReactionType_HAPPY   ReactionType = 4
	ReactionType_SAD     ReactionType = 5
	ReactionType_ANGRY   ReactionType = 6
)

// Enum value maps for ReactionType.
var (
	ReactionType_name = map[int32]string{
		0: "UNKNOWN",
		1: "LIKE",
		2: "DISLIKE",
		3: "LOVE",
		4: "HAPPY",
		5: "SAD",
		6: "ANGRY",
	}
	ReactionType_value = map[string]int32{
		"UNKNOWN": 0,
		"LIKE":    1,
		"DISLIKE": 2,
		"LOVE":    3,
		"HAPPY":   4,
		"SAD":     5,
		"ANGRY":   6,
	}
)

func (x ReactionType) Enum() *ReactionType {
	p := new(ReactionType)
	*p = x
	return p
}

func (x ReactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_example_comments_proto_enumTypes[0].Descriptor()
}

func (ReactionType) Type() protoreflect.EnumType {
	return &file_comment_example_comments_proto_enumTypes[0]
}

func (x ReactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReactionType.Descriptor instead.
func (ReactionType) EnumDescriptor() ([]byte, []int) {
	return file_comment_example_comments_proto_rawDescGZIP(), []int{0}
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId      int64       `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CommentId   int32       `protobuf:"varint,2,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	CommentText string      `protobuf:"bytes,3,opt,name=comment_text,json=commentText,proto3" json:"comment_text,omitempty"`
	Reactions   []*Reaction `protobuf:"bytes,4,rep,name=reactions,proto3" json:"reactions,omitempty"`
	Comments    []*Comment  `protobuf:"bytes,5,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_example_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_example_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comment_example_comments_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Comment) GetCommentId() int32 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *Comment) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

func (x *Comment) GetReactions() []*Reaction {
	if x != nil {
		return x.Reactions
	}
	return nil
}

func (x *Comment) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ReactionType ReactionType `protobuf:"varint,2,opt,name=reaction_type,json=reactionType,proto3,enum=grpc.comment.ReactionType" json:"reaction_type,omitempty"`
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_example_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_comment_example_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_comment_example_comments_proto_rawDescGZIP(), []int{1}
}

func (x *Reaction) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Reaction) GetReactionType() ReactionType {
	if x != nil {
		return x.ReactionType
	}
	return ReactionType_UNKNOWN
}

var File_comment_example_comments_proto protoreflect.FileDescriptor

var file_comment_example_comments_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xcd,
	0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x64,
	0x0a, 0x08, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x2a, 0x5b, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x49, 0x53, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x56, 0x45,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x41, 0x50, 0x50, 0x59, 0x10, 0x04, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x41, 0x44, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4e, 0x47, 0x52, 0x59, 0x10,
	0x06, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_example_comments_proto_rawDescOnce sync.Once
	file_comment_example_comments_proto_rawDescData = file_comment_example_comments_proto_rawDesc
)

func file_comment_example_comments_proto_rawDescGZIP() []byte {
	file_comment_example_comments_proto_rawDescOnce.Do(func() {
		file_comment_example_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_example_comments_proto_rawDescData)
	})
	return file_comment_example_comments_proto_rawDescData
}

var file_comment_example_comments_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_comment_example_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comment_example_comments_proto_goTypes = []interface{}{
	(ReactionType)(0), // 0: grpc.comment.ReactionType
	(*Comment)(nil),   // 1: grpc.comment.Comment
	(*Reaction)(nil),  // 2: grpc.comment.Reaction
}
var file_comment_example_comments_proto_depIdxs = []int32{
	2, // 0: grpc.comment.Comment.reactions:type_name -> grpc.comment.Reaction
	1, // 1: grpc.comment.Comment.comments:type_name -> grpc.comment.Comment
	0, // 2: grpc.comment.Reaction.reaction_type:type_name -> grpc.comment.ReactionType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_comment_example_comments_proto_init() }
func file_comment_example_comments_proto_init() {
	if File_comment_example_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_example_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comment_example_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
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
			RawDescriptor: file_comment_example_comments_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comment_example_comments_proto_goTypes,
		DependencyIndexes: file_comment_example_comments_proto_depIdxs,
		EnumInfos:         file_comment_example_comments_proto_enumTypes,
		MessageInfos:      file_comment_example_comments_proto_msgTypes,
	}.Build()
	File_comment_example_comments_proto = out.File
	file_comment_example_comments_proto_rawDesc = nil
	file_comment_example_comments_proto_goTypes = nil
	file_comment_example_comments_proto_depIdxs = nil
}