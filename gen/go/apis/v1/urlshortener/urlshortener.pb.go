// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: urlshortener.proto

package urlshortener

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{1}
}

type MakeShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// short url 생성할 long url 요청
	LongUrl string `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
}

func (x *MakeShortUrlRequest) Reset() {
	*x = MakeShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeShortUrlRequest) ProtoMessage() {}

func (x *MakeShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeShortUrlRequest.ProtoReflect.Descriptor instead.
func (*MakeShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{2}
}

func (x *MakeShortUrlRequest) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

type MakeShortUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 생성된 short url 반환
	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *MakeShortUrlResponse) Reset() {
	*x = MakeShortUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeShortUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeShortUrlResponse) ProtoMessage() {}

func (x *MakeShortUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeShortUrlResponse.ProtoReflect.Descriptor instead.
func (*MakeShortUrlResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{3}
}

func (x *MakeShortUrlResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetLongUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// short url로 매핑된(저장된) long url 얻기 요청
	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetLongUrlRequest) Reset() {
	*x = GetLongUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLongUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLongUrlRequest) ProtoMessage() {}

func (x *GetLongUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLongUrlRequest.ProtoReflect.Descriptor instead.
func (*GetLongUrlRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{4}
}

func (x *GetLongUrlRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetLongUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// short url에 매핑된(저장된) long url 반환
	LongUrl string `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
}

func (x *GetLongUrlResponse) Reset() {
	*x = GetLongUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLongUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLongUrlResponse) ProtoMessage() {}

func (x *GetLongUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLongUrlResponse.ProtoReflect.Descriptor instead.
func (*GetLongUrlResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{5}
}

func (x *GetLongUrlResponse) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

type RedirectShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// short url을 입력 받아서 long url을 찾으면 redirect
	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *RedirectShortUrlRequest) Reset() {
	*x = RedirectShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectShortUrlRequest) ProtoMessage() {}

func (x *RedirectShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectShortUrlRequest.ProtoReflect.Descriptor instead.
func (*RedirectShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{6}
}

func (x *RedirectShortUrlRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

// If OK, HTTP status code is redirect(302)
// short url로 들어온 요청을 long url로 redirect
type RedirectShortUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *RedirectShortUrlResponse) Reset() {
	*x = RedirectShortUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_urlshortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectShortUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectShortUrlResponse) ProtoMessage() {}

func (x *RedirectShortUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectShortUrlResponse.ProtoReflect.Descriptor instead.
func (*RedirectShortUrlResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{7}
}

func (x *RedirectShortUrlResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

var File_urlshortener_proto protoreflect.FileDescriptor

var file_urlshortener_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x30, 0x0a, 0x13, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55,
	0x72, 0x6c, 0x22, 0x33, 0x0a, 0x14, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2f, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x36, 0x0a, 0x17, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x3d, 0x0a, 0x18, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x32, 0x80, 0x05, 0x0a, 0x0c, 0x55, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x12, 0x69, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x23, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0xac, 0x01,
	0x0a, 0x0c, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x24,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x49, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75,
	0x72, 0x6c, 0x5a, 0x28, 0x12, 0x26, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f,
	0x6d, 0x61, 0x6b, 0x65, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x12, 0xa2, 0x01, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x12, 0x1b, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x67,
	0x65, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x75, 0x72, 0x6c, 0x5a, 0x26, 0x12, 0x24, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x6c, 0x6f, 0x6e, 0x67, 0x75, 0x72,
	0x6c, 0x12, 0xb0, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x41, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5a, 0x24,
	0x12, 0x22, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x42, 0x81, 0x01, 0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x61, 0x6c, 0x61, 0x64, 0x2e, 0x69, 0x64,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x42, 0x11, 0x55, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x61, 0x6e, 0x67, 0x54, 0x61, 0x65, 0x6b, 0x69, 0x2f,
	0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x6c, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_urlshortener_proto_rawDescOnce sync.Once
	file_urlshortener_proto_rawDescData = file_urlshortener_proto_rawDesc
)

func file_urlshortener_proto_rawDescGZIP() []byte {
	file_urlshortener_proto_rawDescOnce.Do(func() {
		file_urlshortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_urlshortener_proto_rawDescData)
	})
	return file_urlshortener_proto_rawDescData
}

var file_urlshortener_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_urlshortener_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),       // 0: v1.urlshortener.HealthCheckRequest
	(*HealthCheckResponse)(nil),      // 1: v1.urlshortener.HealthCheckResponse
	(*MakeShortUrlRequest)(nil),      // 2: v1.urlshortener.MakeShortUrlRequest
	(*MakeShortUrlResponse)(nil),     // 3: v1.urlshortener.MakeShortUrlResponse
	(*GetLongUrlRequest)(nil),        // 4: v1.urlshortener.GetLongUrlRequest
	(*GetLongUrlResponse)(nil),       // 5: v1.urlshortener.GetLongUrlResponse
	(*RedirectShortUrlRequest)(nil),  // 6: v1.urlshortener.RedirectShortUrlRequest
	(*RedirectShortUrlResponse)(nil), // 7: v1.urlshortener.RedirectShortUrlResponse
}
var file_urlshortener_proto_depIdxs = []int32{
	0, // 0: v1.urlshortener.Urlshortener.HealthCheck:input_type -> v1.urlshortener.HealthCheckRequest
	2, // 1: v1.urlshortener.Urlshortener.MakeShortUrl:input_type -> v1.urlshortener.MakeShortUrlRequest
	4, // 2: v1.urlshortener.Urlshortener.GetLongUrl:input_type -> v1.urlshortener.GetLongUrlRequest
	6, // 3: v1.urlshortener.Urlshortener.RedirectShortUrl:input_type -> v1.urlshortener.RedirectShortUrlRequest
	1, // 4: v1.urlshortener.Urlshortener.HealthCheck:output_type -> v1.urlshortener.HealthCheckResponse
	3, // 5: v1.urlshortener.Urlshortener.MakeShortUrl:output_type -> v1.urlshortener.MakeShortUrlResponse
	5, // 6: v1.urlshortener.Urlshortener.GetLongUrl:output_type -> v1.urlshortener.GetLongUrlResponse
	7, // 7: v1.urlshortener.Urlshortener.RedirectShortUrl:output_type -> v1.urlshortener.RedirectShortUrlResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_urlshortener_proto_init() }
func file_urlshortener_proto_init() {
	if File_urlshortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_urlshortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_urlshortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_urlshortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeShortUrlRequest); i {
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
		file_urlshortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeShortUrlResponse); i {
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
		file_urlshortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLongUrlRequest); i {
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
		file_urlshortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLongUrlResponse); i {
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
		file_urlshortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectShortUrlRequest); i {
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
		file_urlshortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectShortUrlResponse); i {
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
			RawDescriptor: file_urlshortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_urlshortener_proto_goTypes,
		DependencyIndexes: file_urlshortener_proto_depIdxs,
		MessageInfos:      file_urlshortener_proto_msgTypes,
	}.Build()
	File_urlshortener_proto = out.File
	file_urlshortener_proto_rawDesc = nil
	file_urlshortener_proto_goTypes = nil
	file_urlshortener_proto_depIdxs = nil
}
