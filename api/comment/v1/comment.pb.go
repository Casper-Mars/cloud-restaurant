// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/comment/v1/comment.proto

package v1

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

type CommentAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FoodId  uint64 `protobuf:"varint,2,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentAddReq) Reset() {
	*x = CommentAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentAddReq) ProtoMessage() {}

func (x *CommentAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_v1_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentAddReq.ProtoReflect.Descriptor instead.
func (*CommentAddReq) Descriptor() ([]byte, []int) {
	return file_api_comment_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentAddReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentAddReq) GetFoodId() uint64 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

func (x *CommentAddReq) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CommentModifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CommentModifyResp) Reset() {
	*x = CommentModifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentModifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentModifyResp) ProtoMessage() {}

func (x *CommentModifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_v1_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentModifyResp.ProtoReflect.Descriptor instead.
func (*CommentModifyResp) Descriptor() ([]byte, []int) {
	return file_api_comment_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentModifyResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CommentList_CommentListItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CommentList) Reset() {
	*x = CommentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_v1_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentList) ProtoMessage() {}

func (x *CommentList) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_v1_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentList.ProtoReflect.Descriptor instead.
func (*CommentList) Descriptor() ([]byte, []int) {
	return file_api_comment_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentList) GetItems() []*CommentList_CommentListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CommentList_CommentListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FoodId  uint64 `protobuf:"varint,3,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	Comment string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentList_CommentListItem) Reset() {
	*x = CommentList_CommentListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comment_v1_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentList_CommentListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentList_CommentListItem) ProtoMessage() {}

func (x *CommentList_CommentListItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_comment_v1_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentList_CommentListItem.ProtoReflect.Descriptor instead.
func (*CommentList_CommentListItem) Descriptor() ([]byte, []int) {
	return file_api_comment_v1_comment_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CommentList_CommentListItem) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentList_CommentListItem) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentList_CommentListItem) GetFoodId() uint64 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

func (x *CommentList_CommentListItem) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

var File_api_comment_v1_comment_proto protoreflect.FileDescriptor

var file_api_comment_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x6d, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x91, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x3b, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x73, 0x70, 0x65, 0x72,
	0x2d, 0x4d, 0x61, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_comment_v1_comment_proto_rawDescOnce sync.Once
	file_api_comment_v1_comment_proto_rawDescData = file_api_comment_v1_comment_proto_rawDesc
)

func file_api_comment_v1_comment_proto_rawDescGZIP() []byte {
	file_api_comment_v1_comment_proto_rawDescOnce.Do(func() {
		file_api_comment_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_comment_v1_comment_proto_rawDescData)
	})
	return file_api_comment_v1_comment_proto_rawDescData
}

var file_api_comment_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_comment_v1_comment_proto_goTypes = []interface{}{
	(*CommentAddReq)(nil),               // 0: comment.v1.CommentAddReq
	(*CommentModifyResp)(nil),           // 1: comment.v1.CommentModifyResp
	(*CommentList)(nil),                 // 2: comment.v1.CommentList
	(*CommentList_CommentListItem)(nil), // 3: comment.v1.CommentList.CommentListItem
	(*emptypb.Empty)(nil),               // 4: google.protobuf.Empty
}
var file_api_comment_v1_comment_proto_depIdxs = []int32{
	3, // 0: comment.v1.CommentList.items:type_name -> comment.v1.CommentList.CommentListItem
	0, // 1: comment.v1.Comment.AddComment:input_type -> comment.v1.CommentAddReq
	4, // 2: comment.v1.Comment.ListComment:input_type -> google.protobuf.Empty
	1, // 3: comment.v1.Comment.AddComment:output_type -> comment.v1.CommentModifyResp
	2, // 4: comment.v1.Comment.ListComment:output_type -> comment.v1.CommentList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_comment_v1_comment_proto_init() }
func file_api_comment_v1_comment_proto_init() {
	if File_api_comment_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_comment_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentAddReq); i {
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
		file_api_comment_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentModifyResp); i {
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
		file_api_comment_v1_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentList); i {
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
		file_api_comment_v1_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentList_CommentListItem); i {
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
			RawDescriptor: file_api_comment_v1_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_comment_v1_comment_proto_goTypes,
		DependencyIndexes: file_api_comment_v1_comment_proto_depIdxs,
		MessageInfos:      file_api_comment_v1_comment_proto_msgTypes,
	}.Build()
	File_api_comment_v1_comment_proto = out.File
	file_api_comment_v1_comment_proto_rawDesc = nil
	file_api_comment_v1_comment_proto_goTypes = nil
	file_api_comment_v1_comment_proto_depIdxs = nil
}
