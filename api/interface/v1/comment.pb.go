// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/interface/v1/comment.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

	UserId  *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FoodId  *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	Comment *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentAddReq) Reset() {
	*x = CommentAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interface_v1_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentAddReq) ProtoMessage() {}

func (x *CommentAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_interface_v1_comment_proto_msgTypes[0]
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
	return file_api_interface_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentAddReq) GetUserId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *CommentAddReq) GetFoodId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.FoodId
	}
	return nil
}

func (x *CommentAddReq) GetComment() *wrapperspb.StringValue {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentModifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId uint64 `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
}

func (x *CommentModifyResp) Reset() {
	*x = CommentModifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interface_v1_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentModifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentModifyResp) ProtoMessage() {}

func (x *CommentModifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_interface_v1_comment_proto_msgTypes[1]
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
	return file_api_interface_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentModifyResp) GetCommentId() uint64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type ListCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ListCommentResp_ListCommentItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCommentResp) Reset() {
	*x = ListCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interface_v1_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentResp) ProtoMessage() {}

func (x *ListCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_interface_v1_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentResp.ProtoReflect.Descriptor instead.
func (*ListCommentResp) Descriptor() ([]byte, []int) {
	return file_api_interface_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *ListCommentResp) GetItems() []*ListCommentResp_ListCommentItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListCommentResp_ListCommentItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FoodId   uint64 `protobuf:"varint,3,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	FoodName string `protobuf:"bytes,4,opt,name=food_name,json=foodName,proto3" json:"food_name,omitempty"`
	Comment  string `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *ListCommentResp_ListCommentItem) Reset() {
	*x = ListCommentResp_ListCommentItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_interface_v1_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentResp_ListCommentItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentResp_ListCommentItem) ProtoMessage() {}

func (x *ListCommentResp_ListCommentItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_interface_v1_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentResp_ListCommentItem.ProtoReflect.Descriptor instead.
func (*ListCommentResp_ListCommentItem) Descriptor() ([]byte, []int) {
	return file_api_interface_v1_comment_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListCommentResp_ListCommentItem) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListCommentResp_ListCommentItem) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ListCommentResp_ListCommentItem) GetFoodId() uint64 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

func (x *ListCommentResp_ListCommentItem) GetFoodName() string {
	if x != nil {
		return x.FoodName
	}
	return ""
}

func (x *ListCommentResp_ListCommentItem) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

var File_api_interface_v1_comment_proto protoreflect.FileDescriptor

var file_api_interface_v1_comment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x32, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xf0, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x97,
	0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xc7, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x1f, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x61, 0x73, 0x70, 0x65, 0x72, 0x2d, 0x4d, 0x61, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_interface_v1_comment_proto_rawDescOnce sync.Once
	file_api_interface_v1_comment_proto_rawDescData = file_api_interface_v1_comment_proto_rawDesc
)

func file_api_interface_v1_comment_proto_rawDescGZIP() []byte {
	file_api_interface_v1_comment_proto_rawDescOnce.Do(func() {
		file_api_interface_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_interface_v1_comment_proto_rawDescData)
	})
	return file_api_interface_v1_comment_proto_rawDescData
}

var file_api_interface_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_interface_v1_comment_proto_goTypes = []interface{}{
	(*CommentAddReq)(nil),                   // 0: interface.v1.CommentAddReq
	(*CommentModifyResp)(nil),               // 1: interface.v1.CommentModifyResp
	(*ListCommentResp)(nil),                 // 2: interface.v1.ListCommentResp
	(*ListCommentResp_ListCommentItem)(nil), // 3: interface.v1.ListCommentResp.ListCommentItem
	(*wrapperspb.UInt64Value)(nil),          // 4: google.protobuf.UInt64Value
	(*wrapperspb.StringValue)(nil),          // 5: google.protobuf.StringValue
	(*emptypb.Empty)(nil),                   // 6: google.protobuf.Empty
}
var file_api_interface_v1_comment_proto_depIdxs = []int32{
	4, // 0: interface.v1.CommentAddReq.user_id:type_name -> google.protobuf.UInt64Value
	4, // 1: interface.v1.CommentAddReq.food_id:type_name -> google.protobuf.UInt64Value
	5, // 2: interface.v1.CommentAddReq.comment:type_name -> google.protobuf.StringValue
	3, // 3: interface.v1.ListCommentResp.items:type_name -> interface.v1.ListCommentResp.ListCommentItem
	0, // 4: interface.v1.Comment.AddComment:input_type -> interface.v1.CommentAddReq
	6, // 5: interface.v1.Comment.ListComment:input_type -> google.protobuf.Empty
	1, // 6: interface.v1.Comment.AddComment:output_type -> interface.v1.CommentModifyResp
	2, // 7: interface.v1.Comment.ListComment:output_type -> interface.v1.ListCommentResp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_interface_v1_comment_proto_init() }
func file_api_interface_v1_comment_proto_init() {
	if File_api_interface_v1_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_interface_v1_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_interface_v1_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_interface_v1_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentResp); i {
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
		file_api_interface_v1_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentResp_ListCommentItem); i {
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
			RawDescriptor: file_api_interface_v1_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_interface_v1_comment_proto_goTypes,
		DependencyIndexes: file_api_interface_v1_comment_proto_depIdxs,
		MessageInfos:      file_api_interface_v1_comment_proto_msgTypes,
	}.Build()
	File_api_interface_v1_comment_proto = out.File
	file_api_interface_v1_comment_proto_rawDesc = nil
	file_api_interface_v1_comment_proto_goTypes = nil
	file_api_interface_v1_comment_proto_depIdxs = nil
}
