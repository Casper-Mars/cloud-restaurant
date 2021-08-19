// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/food/v1/food.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ListFoodByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *ListFoodByIdReq) Reset() {
	*x = ListFoodByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_food_v1_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFoodByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFoodByIdReq) ProtoMessage() {}

func (x *ListFoodByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_food_v1_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFoodByIdReq.ProtoReflect.Descriptor instead.
func (*ListFoodByIdReq) Descriptor() ([]byte, []int) {
	return file_api_food_v1_food_proto_rawDescGZIP(), []int{0}
}

func (x *ListFoodByIdReq) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type AddFoodReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AddFoodReq) Reset() {
	*x = AddFoodReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_food_v1_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFoodReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFoodReq) ProtoMessage() {}

func (x *AddFoodReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_food_v1_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFoodReq.ProtoReflect.Descriptor instead.
func (*AddFoodReq) Descriptor() ([]byte, []int) {
	return file_api_food_v1_food_proto_rawDescGZIP(), []int{1}
}

func (x *AddFoodReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddFoodReq) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateFoodReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateFoodReq) Reset() {
	*x = UpdateFoodReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_food_v1_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFoodReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFoodReq) ProtoMessage() {}

func (x *UpdateFoodReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_food_v1_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFoodReq.ProtoReflect.Descriptor instead.
func (*UpdateFoodReq) Descriptor() ([]byte, []int) {
	return file_api_food_v1_food_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFoodReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFoodReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFoodReq) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type FoodModifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FoodModifyResp) Reset() {
	*x = FoodModifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_food_v1_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodModifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodModifyResp) ProtoMessage() {}

func (x *FoodModifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_food_v1_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodModifyResp.ProtoReflect.Descriptor instead.
func (*FoodModifyResp) Descriptor() ([]byte, []int) {
	return file_api_food_v1_food_proto_rawDescGZIP(), []int{3}
}

func (x *FoodModifyResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OnePageFoodListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageSize  *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Name      string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OnePageFoodListReq) Reset() {
	*x = OnePageFoodListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_food_v1_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnePageFoodListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnePageFoodListReq) ProtoMessage() {}

func (x *OnePageFoodListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_food_v1_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnePageFoodListReq.ProtoReflect.Descriptor instead.
func (*OnePageFoodListReq) Descriptor() ([]byte, []int) {
	return file_api_food_v1_food_proto_rawDescGZIP(), []int{4}
}

func (x *OnePageFoodListReq) GetPageIndex() *wrapperspb.UInt32Value {
	if x != nil {
		return x.PageIndex
	}
	return nil
}

func (x *OnePageFoodListReq) GetPageSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.PageSize
	}
	return nil
}

func (x *OnePageFoodListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FoodListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*FoodListResp_FoodItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *FoodListResp) Reset() {
	*x = FoodListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_food_v1_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodListResp) ProtoMessage() {}

func (x *FoodListResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_food_v1_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodListResp.ProtoReflect.Descriptor instead.
func (*FoodListResp) Descriptor() ([]byte, []int) {
	return file_api_food_v1_food_proto_rawDescGZIP(), []int{5}
}

func (x *FoodListResp) GetItems() []*FoodListResp_FoodItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type FoodListResp_FoodItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *FoodListResp_FoodItem) Reset() {
	*x = FoodListResp_FoodItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_food_v1_food_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodListResp_FoodItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodListResp_FoodItem) ProtoMessage() {}

func (x *FoodListResp_FoodItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_food_v1_food_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodListResp_FoodItem.ProtoReflect.Descriptor instead.
func (*FoodListResp_FoodItem) Descriptor() ([]byte, []int) {
	return file_api_food_v1_food_proto_rawDescGZIP(), []int{5, 0}
}

func (x *FoodListResp_FoodItem) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FoodListResp_FoodItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FoodListResp_FoodItem) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_api_food_v1_food_proto protoreflect.FileDescriptor

var file_api_food_v1_food_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f,
	0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x46, 0x6f, 0x6f, 0x64, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x4f, 0x6e,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x39, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8a, 0x01, 0x0a,
	0x0c, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66,
	0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x1a, 0x44, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0xf0, 0x01, 0x0a, 0x04, 0x46, 0x6f,
	0x6f, 0x64, 0x12, 0x33, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x3a, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x66, 0x6f, 0x6f,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x65, 0x50, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x6f, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x18, 0x2e, 0x66, 0x6f,
	0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x73, 0x70, 0x65,
	0x72, 0x2d, 0x4d, 0x61, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6f, 0x6f, 0x64,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_food_v1_food_proto_rawDescOnce sync.Once
	file_api_food_v1_food_proto_rawDescData = file_api_food_v1_food_proto_rawDesc
)

func file_api_food_v1_food_proto_rawDescGZIP() []byte {
	file_api_food_v1_food_proto_rawDescOnce.Do(func() {
		file_api_food_v1_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_food_v1_food_proto_rawDescData)
	})
	return file_api_food_v1_food_proto_rawDescData
}

var file_api_food_v1_food_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_food_v1_food_proto_goTypes = []interface{}{
	(*ListFoodByIdReq)(nil),        // 0: food.v1.ListFoodByIdReq
	(*AddFoodReq)(nil),             // 1: food.v1.AddFoodReq
	(*UpdateFoodReq)(nil),          // 2: food.v1.UpdateFoodReq
	(*FoodModifyResp)(nil),         // 3: food.v1.FoodModifyResp
	(*OnePageFoodListReq)(nil),     // 4: food.v1.OnePageFoodListReq
	(*FoodListResp)(nil),           // 5: food.v1.FoodListResp
	(*FoodListResp_FoodItem)(nil),  // 6: food.v1.FoodListResp.FoodItem
	(*wrapperspb.UInt32Value)(nil), // 7: google.protobuf.UInt32Value
}
var file_api_food_v1_food_proto_depIdxs = []int32{
	7, // 0: food.v1.OnePageFoodListReq.page_index:type_name -> google.protobuf.UInt32Value
	7, // 1: food.v1.OnePageFoodListReq.page_size:type_name -> google.protobuf.UInt32Value
	6, // 2: food.v1.FoodListResp.items:type_name -> food.v1.FoodListResp.FoodItem
	1, // 3: food.v1.Food.Add:input_type -> food.v1.AddFoodReq
	2, // 4: food.v1.Food.Update:input_type -> food.v1.UpdateFoodReq
	4, // 5: food.v1.Food.Page:input_type -> food.v1.OnePageFoodListReq
	0, // 6: food.v1.Food.ListByIds:input_type -> food.v1.ListFoodByIdReq
	3, // 7: food.v1.Food.Add:output_type -> food.v1.FoodModifyResp
	3, // 8: food.v1.Food.Update:output_type -> food.v1.FoodModifyResp
	5, // 9: food.v1.Food.Page:output_type -> food.v1.FoodListResp
	5, // 10: food.v1.Food.ListByIds:output_type -> food.v1.FoodListResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_food_v1_food_proto_init() }
func file_api_food_v1_food_proto_init() {
	if File_api_food_v1_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_food_v1_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFoodByIdReq); i {
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
		file_api_food_v1_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFoodReq); i {
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
		file_api_food_v1_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFoodReq); i {
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
		file_api_food_v1_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodModifyResp); i {
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
		file_api_food_v1_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnePageFoodListReq); i {
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
		file_api_food_v1_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodListResp); i {
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
		file_api_food_v1_food_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodListResp_FoodItem); i {
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
			RawDescriptor: file_api_food_v1_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_food_v1_food_proto_goTypes,
		DependencyIndexes: file_api_food_v1_food_proto_depIdxs,
		MessageInfos:      file_api_food_v1_food_proto_msgTypes,
	}.Build()
	File_api_food_v1_food_proto = out.File
	file_api_food_v1_food_proto_rawDesc = nil
	file_api_food_v1_food_proto_goTypes = nil
	file_api_food_v1_food_proto_depIdxs = nil
}
