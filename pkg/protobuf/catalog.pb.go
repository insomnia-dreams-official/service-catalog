// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pkg/protobuf/catalog.proto

package protobuf

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Structure for site navigation's item.
type NavigationItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Link  string            `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	Items []*NavigationItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *NavigationItem) Reset() {
	*x = NavigationItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NavigationItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NavigationItem) ProtoMessage() {}

func (x *NavigationItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NavigationItem.ProtoReflect.Descriptor instead.
func (*NavigationItem) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *NavigationItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NavigationItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NavigationItem) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *NavigationItem) GetItems() []*NavigationItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetNavigationItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNavigationItemsRequest) Reset() {
	*x = GetNavigationItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNavigationItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNavigationItemsRequest) ProtoMessage() {}

func (x *GetNavigationItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNavigationItemsRequest.ProtoReflect.Descriptor instead.
func (*GetNavigationItemsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{1}
}

type GetNavigationItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NavigationItems []*NavigationItem `protobuf:"bytes,1,rep,name=navigation_items,json=navigationItems,proto3" json:"navigation_items,omitempty"`
}

func (x *GetNavigationItemsResponse) Reset() {
	*x = GetNavigationItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNavigationItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNavigationItemsResponse) ProtoMessage() {}

func (x *GetNavigationItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNavigationItemsResponse.ProtoReflect.Descriptor instead.
func (*GetNavigationItemsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *GetNavigationItemsResponse) GetNavigationItems() []*NavigationItem {
	if x != nil {
		return x.NavigationItems
	}
	return nil
}

// Structure for site category; category is an attribute of every product
type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Articul  string `protobuf:"bytes,1,opt,name=articul,proto3" json:"articul,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Link     string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	FullLink string `protobuf:"bytes,5,opt,name=full_link,json=fullLink,proto3" json:"full_link,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *Category) GetArticul() string {
	if x != nil {
		return x.Articul
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Category) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Category) GetFullLink() string {
	if x != nil {
		return x.FullLink
	}
	return ""
}

type GetRootcategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRootcategoriesRequest) Reset() {
	*x = GetRootcategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRootcategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRootcategoriesRequest) ProtoMessage() {}

func (x *GetRootcategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRootcategoriesRequest.ProtoReflect.Descriptor instead.
func (*GetRootcategoriesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{4}
}

type GetRootcategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rootcategories []*Category `protobuf:"bytes,1,rep,name=rootcategories,proto3" json:"rootcategories,omitempty"`
}

func (x *GetRootcategoriesResponse) Reset() {
	*x = GetRootcategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRootcategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRootcategoriesResponse) ProtoMessage() {}

func (x *GetRootcategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRootcategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetRootcategoriesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *GetRootcategoriesResponse) GetRootcategories() []*Category {
	if x != nil {
		return x.Rootcategories
	}
	return nil
}

type GetCategoryChildsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *GetCategoryChildsRequest) Reset() {
	*x = GetCategoryChildsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryChildsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryChildsRequest) ProtoMessage() {}

func (x *GetCategoryChildsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryChildsRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryChildsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{6}
}

func (x *GetCategoryChildsRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetCategoryChildsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryChilds []*Category `protobuf:"bytes,1,rep,name=categoryChilds,proto3" json:"categoryChilds,omitempty"`
}

func (x *GetCategoryChildsResponse) Reset() {
	*x = GetCategoryChildsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_catalog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryChildsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryChildsResponse) ProtoMessage() {}

func (x *GetCategoryChildsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_catalog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryChildsResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryChildsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_catalog_proto_rawDescGZIP(), []int{7}
}

func (x *GetCategoryChildsResponse) GetCategoryChilds() []*Category {
	if x != nil {
		return x.CategoryChilds
	}
	return nil
}

var File_pkg_protobuf_catalog_proto protoreflect.FileDescriptor

var file_pkg_protobuf_catalog_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x78, 0x0a, 0x0e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x6e,
	0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x0f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x7d, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x75, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x22,
	0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x57, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x69,
	0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x32, 0xac, 0x02,
	0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x61, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61,
	0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x73, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_protobuf_catalog_proto_rawDescOnce sync.Once
	file_pkg_protobuf_catalog_proto_rawDescData = file_pkg_protobuf_catalog_proto_rawDesc
)

func file_pkg_protobuf_catalog_proto_rawDescGZIP() []byte {
	file_pkg_protobuf_catalog_proto_rawDescOnce.Do(func() {
		file_pkg_protobuf_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protobuf_catalog_proto_rawDescData)
	})
	return file_pkg_protobuf_catalog_proto_rawDescData
}

var file_pkg_protobuf_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_protobuf_catalog_proto_goTypes = []interface{}{
	(*NavigationItem)(nil),             // 0: protobuf.NavigationItem
	(*GetNavigationItemsRequest)(nil),  // 1: protobuf.GetNavigationItemsRequest
	(*GetNavigationItemsResponse)(nil), // 2: protobuf.GetNavigationItemsResponse
	(*Category)(nil),                   // 3: protobuf.Category
	(*GetRootcategoriesRequest)(nil),   // 4: protobuf.GetRootcategoriesRequest
	(*GetRootcategoriesResponse)(nil),  // 5: protobuf.GetRootcategoriesResponse
	(*GetCategoryChildsRequest)(nil),   // 6: protobuf.GetCategoryChildsRequest
	(*GetCategoryChildsResponse)(nil),  // 7: protobuf.GetCategoryChildsResponse
}
var file_pkg_protobuf_catalog_proto_depIdxs = []int32{
	0, // 0: protobuf.NavigationItem.items:type_name -> protobuf.NavigationItem
	0, // 1: protobuf.GetNavigationItemsResponse.navigation_items:type_name -> protobuf.NavigationItem
	3, // 2: protobuf.GetRootcategoriesResponse.rootcategories:type_name -> protobuf.Category
	3, // 3: protobuf.GetCategoryChildsResponse.categoryChilds:type_name -> protobuf.Category
	1, // 4: protobuf.Catalog.GetNavigationItems:input_type -> protobuf.GetNavigationItemsRequest
	4, // 5: protobuf.Catalog.GetRootcategories:input_type -> protobuf.GetRootcategoriesRequest
	6, // 6: protobuf.Catalog.GetCategoryChilds:input_type -> protobuf.GetCategoryChildsRequest
	2, // 7: protobuf.Catalog.GetNavigationItems:output_type -> protobuf.GetNavigationItemsResponse
	5, // 8: protobuf.Catalog.GetRootcategories:output_type -> protobuf.GetRootcategoriesResponse
	7, // 9: protobuf.Catalog.GetCategoryChilds:output_type -> protobuf.GetCategoryChildsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_protobuf_catalog_proto_init() }
func file_pkg_protobuf_catalog_proto_init() {
	if File_pkg_protobuf_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protobuf_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NavigationItem); i {
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
		file_pkg_protobuf_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNavigationItemsRequest); i {
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
		file_pkg_protobuf_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNavigationItemsResponse); i {
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
		file_pkg_protobuf_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_pkg_protobuf_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRootcategoriesRequest); i {
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
		file_pkg_protobuf_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRootcategoriesResponse); i {
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
		file_pkg_protobuf_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryChildsRequest); i {
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
		file_pkg_protobuf_catalog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryChildsResponse); i {
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
			RawDescriptor: file_pkg_protobuf_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protobuf_catalog_proto_goTypes,
		DependencyIndexes: file_pkg_protobuf_catalog_proto_depIdxs,
		MessageInfos:      file_pkg_protobuf_catalog_proto_msgTypes,
	}.Build()
	File_pkg_protobuf_catalog_proto = out.File
	file_pkg_protobuf_catalog_proto_rawDesc = nil
	file_pkg_protobuf_catalog_proto_goTypes = nil
	file_pkg_protobuf_catalog_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogClient interface {
	GetNavigationItems(ctx context.Context, in *GetNavigationItemsRequest, opts ...grpc.CallOption) (*GetNavigationItemsResponse, error)
	GetRootcategories(ctx context.Context, in *GetRootcategoriesRequest, opts ...grpc.CallOption) (*GetRootcategoriesResponse, error)
	GetCategoryChilds(ctx context.Context, in *GetCategoryChildsRequest, opts ...grpc.CallOption) (*GetCategoryChildsResponse, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) GetNavigationItems(ctx context.Context, in *GetNavigationItemsRequest, opts ...grpc.CallOption) (*GetNavigationItemsResponse, error) {
	out := new(GetNavigationItemsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Catalog/GetNavigationItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetRootcategories(ctx context.Context, in *GetRootcategoriesRequest, opts ...grpc.CallOption) (*GetRootcategoriesResponse, error) {
	out := new(GetRootcategoriesResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Catalog/GetRootcategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) GetCategoryChilds(ctx context.Context, in *GetCategoryChildsRequest, opts ...grpc.CallOption) (*GetCategoryChildsResponse, error) {
	out := new(GetCategoryChildsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Catalog/GetCategoryChilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
type CatalogServer interface {
	GetNavigationItems(context.Context, *GetNavigationItemsRequest) (*GetNavigationItemsResponse, error)
	GetRootcategories(context.Context, *GetRootcategoriesRequest) (*GetRootcategoriesResponse, error)
	GetCategoryChilds(context.Context, *GetCategoryChildsRequest) (*GetCategoryChildsResponse, error)
}

// UnimplementedCatalogServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (*UnimplementedCatalogServer) GetNavigationItems(context.Context, *GetNavigationItemsRequest) (*GetNavigationItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNavigationItems not implemented")
}
func (*UnimplementedCatalogServer) GetRootcategories(context.Context, *GetRootcategoriesRequest) (*GetRootcategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootcategories not implemented")
}
func (*UnimplementedCatalogServer) GetCategoryChilds(context.Context, *GetCategoryChildsRequest) (*GetCategoryChildsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryChilds not implemented")
}

func RegisterCatalogServer(s *grpc.Server, srv CatalogServer) {
	s.RegisterService(&_Catalog_serviceDesc, srv)
}

func _Catalog_GetNavigationItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNavigationItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetNavigationItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Catalog/GetNavigationItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetNavigationItems(ctx, req.(*GetNavigationItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetRootcategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRootcategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetRootcategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Catalog/GetRootcategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetRootcategories(ctx, req.(*GetRootcategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_GetCategoryChilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryChildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).GetCategoryChilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Catalog/GetCategoryChilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).GetCategoryChilds(ctx, req.(*GetCategoryChildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Catalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNavigationItems",
			Handler:    _Catalog_GetNavigationItems_Handler,
		},
		{
			MethodName: "GetRootcategories",
			Handler:    _Catalog_GetRootcategories_Handler,
		},
		{
			MethodName: "GetCategoryChilds",
			Handler:    _Catalog_GetCategoryChilds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/catalog.proto",
}
