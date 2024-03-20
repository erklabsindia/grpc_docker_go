// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: attachments.proto

package generated

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

// Define the Attachment message
type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url            string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Ref            string `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type           string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	BlurHash       string `protobuf:"bytes,5,opt,name=blur_hash,json=blurHash,proto3" json:"blur_hash,omitempty"`
	Thumbnail      string `protobuf:"bytes,6,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	LocalUploadRef string `protobuf:"bytes,7,opt,name=local_upload_ref,json=localUploadRef,proto3" json:"local_upload_ref,omitempty"`
	Id             int64  `protobuf:"varint,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_attachments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_attachments_proto_rawDescGZIP(), []int{0}
}

func (x *Attachment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Attachment) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *Attachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attachment) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Attachment) GetBlurHash() string {
	if x != nil {
		return x.BlurHash
	}
	return ""
}

func (x *Attachment) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Attachment) GetLocalUploadRef() string {
	if x != nil {
		return x.LocalUploadRef
	}
	return ""
}

func (x *Attachment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// AttachMents
type AttachmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttachmentId string `protobuf:"bytes,1,opt,name=attachment_id,json=attachmentId,proto3" json:"attachment_id,omitempty"`
	PostId       string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Uuid         string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *AttachmentRequest) Reset() {
	*x = AttachmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentRequest) ProtoMessage() {}

func (x *AttachmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attachments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentRequest.ProtoReflect.Descriptor instead.
func (*AttachmentRequest) Descriptor() ([]byte, []int) {
	return file_attachments_proto_rawDescGZIP(), []int{1}
}

func (x *AttachmentRequest) GetAttachmentId() string {
	if x != nil {
		return x.AttachmentId
	}
	return ""
}

func (x *AttachmentRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *AttachmentRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type AttachmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message     string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type        string        `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Attachments []*Attachment `protobuf:"bytes,4,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *AttachmentResponse) Reset() {
	*x = AttachmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attachments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachmentResponse) ProtoMessage() {}

func (x *AttachmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attachments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachmentResponse.ProtoReflect.Descriptor instead.
func (*AttachmentResponse) Descriptor() ([]byte, []int) {
	return file_attachments_proto_rawDescGZIP(), []int{2}
}

func (x *AttachmentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AttachmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AttachmentResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttachmentResponse) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

var File_attachments_proto protoreflect.FileDescriptor

var file_attachments_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x6c, 0x75, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x6c, 0x75, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x65, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x12, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xcf, 0x03, 0x0a, 0x19, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x57, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x20, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attachments_proto_rawDescOnce sync.Once
	file_attachments_proto_rawDescData = file_attachments_proto_rawDesc
)

func file_attachments_proto_rawDescGZIP() []byte {
	file_attachments_proto_rawDescOnce.Do(func() {
		file_attachments_proto_rawDescData = protoimpl.X.CompressGZIP(file_attachments_proto_rawDescData)
	})
	return file_attachments_proto_rawDescData
}

var file_attachments_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_attachments_proto_goTypes = []interface{}{
	(*Attachment)(nil),         // 0: worklen.proto.Attachment
	(*AttachmentRequest)(nil),  // 1: worklen.proto.AttachmentRequest
	(*AttachmentResponse)(nil), // 2: worklen.proto.AttachmentResponse
}
var file_attachments_proto_depIdxs = []int32{
	0, // 0: worklen.proto.AttachmentResponse.attachments:type_name -> worklen.proto.Attachment
	1, // 1: worklen.proto.AttachmentProtobufService.ListAttachment:input_type -> worklen.proto.AttachmentRequest
	0, // 2: worklen.proto.AttachmentProtobufService.CreateAttachment:input_type -> worklen.proto.Attachment
	1, // 3: worklen.proto.AttachmentProtobufService.GetAttachment:input_type -> worklen.proto.AttachmentRequest
	0, // 4: worklen.proto.AttachmentProtobufService.UpdateAttachment:input_type -> worklen.proto.Attachment
	1, // 5: worklen.proto.AttachmentProtobufService.DeleteAttachment:input_type -> worklen.proto.AttachmentRequest
	2, // 6: worklen.proto.AttachmentProtobufService.ListAttachment:output_type -> worklen.proto.AttachmentResponse
	2, // 7: worklen.proto.AttachmentProtobufService.CreateAttachment:output_type -> worklen.proto.AttachmentResponse
	2, // 8: worklen.proto.AttachmentProtobufService.GetAttachment:output_type -> worklen.proto.AttachmentResponse
	2, // 9: worklen.proto.AttachmentProtobufService.UpdateAttachment:output_type -> worklen.proto.AttachmentResponse
	2, // 10: worklen.proto.AttachmentProtobufService.DeleteAttachment:output_type -> worklen.proto.AttachmentResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_attachments_proto_init() }
func file_attachments_proto_init() {
	if File_attachments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attachments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_attachments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentRequest); i {
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
		file_attachments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachmentResponse); i {
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
			RawDescriptor: file_attachments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attachments_proto_goTypes,
		DependencyIndexes: file_attachments_proto_depIdxs,
		MessageInfos:      file_attachments_proto_msgTypes,
	}.Build()
	File_attachments_proto = out.File
	file_attachments_proto_rawDesc = nil
	file_attachments_proto_goTypes = nil
	file_attachments_proto_depIdxs = nil
}
