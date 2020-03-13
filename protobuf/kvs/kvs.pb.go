// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/kvs/kvs.proto

package kvs

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KVSCommand_Type int32

const (
	KVSCommand_UNKNOWN KVSCommand_Type = 0
	KVSCommand_JOIN    KVSCommand_Type = 1
	KVSCommand_LEAVE   KVSCommand_Type = 2
	KVSCommand_PUT     KVSCommand_Type = 3
	KVSCommand_DELETE  KVSCommand_Type = 4
)

var KVSCommand_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "JOIN",
	2: "LEAVE",
	3: "PUT",
	4: "DELETE",
}

var KVSCommand_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"JOIN":    1,
	"LEAVE":   2,
	"PUT":     3,
	"DELETE":  4,
}

func (x KVSCommand_Type) String() string {
	return proto.EnumName(KVSCommand_Type_name, int32(x))
}

func (KVSCommand_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{9, 0}
}

type WatchResponse_Event int32

const (
	WatchResponse_UNKNOWN WatchResponse_Event = 0
	WatchResponse_JOIN    WatchResponse_Event = 1
	WatchResponse_LEAVE   WatchResponse_Event = 2
	WatchResponse_PUT     WatchResponse_Event = 3
	WatchResponse_DELETE  WatchResponse_Event = 4
)

var WatchResponse_Event_name = map[int32]string{
	0: "UNKNOWN",
	1: "JOIN",
	2: "LEAVE",
	3: "PUT",
	4: "DELETE",
}

var WatchResponse_Event_value = map[string]int32{
	"UNKNOWN": 0,
	"JOIN":    1,
	"LEAVE":   2,
	"PUT":     3,
	"DELETE":  4,
}

func (x WatchResponse_Event) String() string {
	return proto.EnumName(WatchResponse_Event_name, int32(x))
}

func (WatchResponse_Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{10, 0}
}

type Node struct {
	BindAddr             string   `protobuf:"bytes,1,opt,name=bind_addr,json=bindAddr,proto3" json:"bind_addr,omitempty"`
	GrpcAddr             string   `protobuf:"bytes,2,opt,name=grpc_addr,json=grpcAddr,proto3" json:"grpc_addr,omitempty"`
	HttpAddr             string   `protobuf:"bytes,3,opt,name=http_addr,json=httpAddr,proto3" json:"http_addr,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

func (m *Node) GetGrpcAddr() string {
	if m != nil {
		return m.GrpcAddr
	}
	return ""
}

func (m *Node) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

func (m *Node) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type JoinRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BindAddr             string   `protobuf:"bytes,2,opt,name=bind_addr,json=bindAddr,proto3" json:"bind_addr,omitempty"`
	GrpcAddr             string   `protobuf:"bytes,3,opt,name=grpc_addr,json=grpcAddr,proto3" json:"grpc_addr,omitempty"`
	HttpAddr             string   `protobuf:"bytes,4,opt,name=http_addr,json=httpAddr,proto3" json:"http_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{1}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JoinRequest) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

func (m *JoinRequest) GetGrpcAddr() string {
	if m != nil {
		return m.GrpcAddr
	}
	return ""
}

func (m *JoinRequest) GetHttpAddr() string {
	if m != nil {
		return m.HttpAddr
	}
	return ""
}

type LeaveRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRequest) Reset()         { *m = LeaveRequest{} }
func (m *LeaveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveRequest) ProtoMessage()    {}
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{2}
}

func (m *LeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRequest.Unmarshal(m, b)
}
func (m *LeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRequest.Marshal(b, m, deterministic)
}
func (m *LeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRequest.Merge(m, src)
}
func (m *LeaveRequest) XXX_Size() int {
	return xxx_messageInfo_LeaveRequest.Size(m)
}
func (m *LeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRequest proto.InternalMessageInfo

func (m *LeaveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NodeResponse struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeResponse) Reset()         { *m = NodeResponse{} }
func (m *NodeResponse) String() string { return proto.CompactTextString(m) }
func (*NodeResponse) ProtoMessage()    {}
func (*NodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{3}
}

func (m *NodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeResponse.Unmarshal(m, b)
}
func (m *NodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeResponse.Marshal(b, m, deterministic)
}
func (m *NodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeResponse.Merge(m, src)
}
func (m *NodeResponse) XXX_Size() int {
	return xxx_messageInfo_NodeResponse.Size(m)
}
func (m *NodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeResponse proto.InternalMessageInfo

func (m *NodeResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type ClusterResponse struct {
	Nodes                map[string]*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ClusterResponse) Reset()         { *m = ClusterResponse{} }
func (m *ClusterResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterResponse) ProtoMessage()    {}
func (*ClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{4}
}

func (m *ClusterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResponse.Unmarshal(m, b)
}
func (m *ClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResponse.Marshal(b, m, deterministic)
}
func (m *ClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResponse.Merge(m, src)
}
func (m *ClusterResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterResponse.Size(m)
}
func (m *ClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResponse proto.InternalMessageInfo

func (m *ClusterResponse) GetNodes() map[string]*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GetRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{7}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

func (m *PutRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type DeleteRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{8}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type KVSCommand struct {
	Type                 KVSCommand_Type `protobuf:"varint,1,opt,name=type,proto3,enum=kvs.KVSCommand_Type" json:"type,omitempty"`
	Data                 *any.Any        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *KVSCommand) Reset()         { *m = KVSCommand{} }
func (m *KVSCommand) String() string { return proto.CompactTextString(m) }
func (*KVSCommand) ProtoMessage()    {}
func (*KVSCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{9}
}

func (m *KVSCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVSCommand.Unmarshal(m, b)
}
func (m *KVSCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVSCommand.Marshal(b, m, deterministic)
}
func (m *KVSCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVSCommand.Merge(m, src)
}
func (m *KVSCommand) XXX_Size() int {
	return xxx_messageInfo_KVSCommand.Size(m)
}
func (m *KVSCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_KVSCommand.DiscardUnknown(m)
}

var xxx_messageInfo_KVSCommand proto.InternalMessageInfo

func (m *KVSCommand) GetType() KVSCommand_Type {
	if m != nil {
		return m.Type
	}
	return KVSCommand_UNKNOWN
}

func (m *KVSCommand) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type WatchResponse struct {
	Event                WatchResponse_Event `protobuf:"varint,1,opt,name=event,proto3,enum=kvs.WatchResponse_Event" json:"event,omitempty"`
	Data                 *any.Any            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{10}
}

func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchResponse.Unmarshal(m, b)
}
func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
}
func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}
func (m *WatchResponse) XXX_Size() int {
	return xxx_messageInfo_WatchResponse.Size(m)
}
func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetEvent() WatchResponse_Event {
	if m != nil {
		return m.Event
	}
	return WatchResponse_UNKNOWN
}

func (m *WatchResponse) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type KeyValuePair struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e9909cfc2f34163, []int{11}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValuePair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("kvs.KVSCommand_Type", KVSCommand_Type_name, KVSCommand_Type_value)
	proto.RegisterEnum("kvs.WatchResponse_Event", WatchResponse_Event_name, WatchResponse_Event_value)
	proto.RegisterType((*Node)(nil), "kvs.Node")
	proto.RegisterType((*JoinRequest)(nil), "kvs.JoinRequest")
	proto.RegisterType((*LeaveRequest)(nil), "kvs.LeaveRequest")
	proto.RegisterType((*NodeResponse)(nil), "kvs.NodeResponse")
	proto.RegisterType((*ClusterResponse)(nil), "kvs.ClusterResponse")
	proto.RegisterMapType((map[string]*Node)(nil), "kvs.ClusterResponse.NodesEntry")
	proto.RegisterType((*GetRequest)(nil), "kvs.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "kvs.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "kvs.PutRequest")
	proto.RegisterType((*DeleteRequest)(nil), "kvs.DeleteRequest")
	proto.RegisterType((*KVSCommand)(nil), "kvs.KVSCommand")
	proto.RegisterType((*WatchResponse)(nil), "kvs.WatchResponse")
	proto.RegisterType((*KeyValuePair)(nil), "kvs.KeyValuePair")
}

func init() { proto.RegisterFile("protobuf/kvs/kvs.proto", fileDescriptor_6e9909cfc2f34163) }

var fileDescriptor_6e9909cfc2f34163 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x6f, 0x4f, 0xda, 0x5e,
	0x14, 0xa6, 0xff, 0x54, 0x0e, 0xa8, 0xf5, 0x86, 0x18, 0x7e, 0x98, 0x9f, 0xba, 0x6b, 0x96, 0x98,
	0x25, 0x2b, 0x06, 0x9d, 0xfb, 0x93, 0x6d, 0x89, 0xd3, 0xc6, 0x4c, 0x0c, 0x92, 0x8a, 0x98, 0xec,
	0xcd, 0x52, 0xe8, 0x1d, 0x10, 0xa0, 0xed, 0xda, 0x5b, 0x92, 0x7e, 0x88, 0x7d, 0x86, 0xbd, 0xde,
	0x77, 0xdb, 0x87, 0x58, 0xee, 0xbd, 0x85, 0x82, 0xae, 0x64, 0xdb, 0x0b, 0x12, 0x7a, 0x9e, 0xe7,
	0x39, 0xf7, 0xb9, 0x87, 0xf3, 0x14, 0xd8, 0xf6, 0x03, 0x8f, 0x7a, 0x9d, 0xe8, 0x4b, 0x75, 0x38,
	0x09, 0xd9, 0xc7, 0xe0, 0x05, 0xa4, 0x0c, 0x27, 0x61, 0xe5, 0xbf, 0x9e, 0xe7, 0xf5, 0x46, 0xa4,
	0x3a, 0xe3, 0xd8, 0x6e, 0x2c, 0xf0, 0xca, 0xce, 0x43, 0x88, 0x8c, 0x7d, 0x9a, 0x80, 0x38, 0x04,
	0xb5, 0xe1, 0x39, 0x04, 0xed, 0x40, 0xbe, 0x33, 0x70, 0x9d, 0xcf, 0xb6, 0xe3, 0x04, 0x65, 0x69,
	0x5f, 0x3a, 0xcc, 0x5b, 0x6b, 0xac, 0x70, 0xe6, 0x38, 0x01, 0x03, 0x7b, 0x81, 0xdf, 0x15, 0xa0,
	0x2c, 0x40, 0x56, 0x98, 0x82, 0x7d, 0x4a, 0x7d, 0x01, 0x2a, 0x02, 0x64, 0x05, 0x0e, 0x96, 0x40,
	0x0b, 0xa9, 0x4d, 0x49, 0x59, 0xe5, 0x80, 0x78, 0xc0, 0x14, 0x0a, 0x57, 0xde, 0xc0, 0xb5, 0xc8,
	0xd7, 0x88, 0x84, 0x14, 0x6d, 0x80, 0x3c, 0x70, 0x92, 0x43, 0xe5, 0x81, 0xb3, 0xe8, 0x45, 0x5e,
	0xe6, 0x45, 0x59, 0xe6, 0x45, 0x5d, 0xf4, 0x82, 0x77, 0xa1, 0x78, 0x4d, 0xec, 0x09, 0xc9, 0x38,
	0x16, 0x3f, 0x87, 0x22, 0x1b, 0x85, 0x45, 0x42, 0xdf, 0x73, 0x43, 0x82, 0xfe, 0x07, 0xd5, 0xf5,
	0x1c, 0xc2, 0x19, 0x85, 0x5a, 0xde, 0x60, 0x13, 0xe7, 0x04, 0x5e, 0xc6, 0xdf, 0x24, 0xd8, 0x3c,
	0x1f, 0x45, 0x21, 0x25, 0xc1, 0x4c, 0xf2, 0x02, 0x34, 0x86, 0x85, 0x65, 0x69, 0x5f, 0x39, 0x2c,
	0xd4, 0xf6, 0xb8, 0xe6, 0x01, 0x89, 0xf7, 0x08, 0x4d, 0x97, 0x06, 0xb1, 0x25, 0xd8, 0x95, 0x73,
	0x80, 0xb4, 0x88, 0x74, 0x50, 0x86, 0x24, 0x4e, 0x8c, 0xb1, 0xaf, 0x68, 0x0f, 0xb4, 0x89, 0x3d,
	0x8a, 0x08, 0x1f, 0xc6, 0x82, 0x15, 0x51, 0x7f, 0x23, 0xbf, 0x92, 0xf0, 0x2e, 0xc0, 0x25, 0xa1,
	0xd3, 0xcb, 0xcd, 0x35, 0x29, 0xf2, 0x26, 0xf8, 0x00, 0x0a, 0x1c, 0x4f, 0xac, 0x96, 0xa6, 0x3d,
	0x05, 0x45, 0x3c, 0xe0, 0x13, 0x80, 0x66, 0x94, 0xdd, 0x24, 0x55, 0xc9, 0xf3, 0xaa, 0x27, 0xb0,
	0x7e, 0x41, 0x46, 0x84, 0x92, 0xec, 0xd3, 0xbf, 0x4b, 0x00, 0xf5, 0xf6, 0xed, 0xb9, 0x37, 0x1e,
	0xdb, 0xae, 0x83, 0x0e, 0x41, 0xa5, 0xb1, 0x2f, 0x0e, 0xdf, 0xa8, 0x95, 0xf8, 0x85, 0x52, 0xd8,
	0x68, 0xc5, 0x3e, 0xb1, 0x38, 0x83, 0x31, 0x1d, 0x9b, 0xda, 0xc9, 0xd5, 0x4b, 0x86, 0x58, 0x66,
	0x63, 0xba, 0xcc, 0xc6, 0x99, 0x1b, 0x5b, 0x9c, 0x81, 0xdf, 0x81, 0xca, 0x74, 0xa8, 0x00, 0xab,
	0x77, 0x8d, 0x7a, 0xe3, 0xe6, 0xbe, 0xa1, 0xe7, 0xd0, 0x1a, 0xa8, 0x57, 0x37, 0x1f, 0x1b, 0xba,
	0x84, 0xf2, 0xa0, 0x5d, 0x9b, 0x67, 0x6d, 0x53, 0x97, 0xd1, 0x2a, 0x28, 0xcd, 0xbb, 0x96, 0xae,
	0x20, 0x80, 0x95, 0x0b, 0xf3, 0xda, 0x6c, 0x99, 0xba, 0x8a, 0x7f, 0x48, 0xb0, 0x7e, 0x6f, 0xd3,
	0x6e, 0x7f, 0x36, 0x22, 0x03, 0x34, 0x32, 0x21, 0x2e, 0x4d, 0x5c, 0x96, 0xb9, 0xcb, 0x05, 0x8a,
	0x61, 0x32, 0xdc, 0x12, 0xb4, 0xbf, 0xb0, 0xfa, 0x1e, 0x34, 0xae, 0xfc, 0x57, 0xaf, 0xa7, 0x50,
	0xac, 0x93, 0xb8, 0xcd, 0x86, 0xdf, 0xb4, 0x07, 0xc1, 0x9f, 0xfe, 0x50, 0xb5, 0x9f, 0x0a, 0x28,
	0xf5, 0xf6, 0x2d, 0xaa, 0x81, 0xca, 0x02, 0x88, 0x74, 0x7e, 0xa5, 0xb9, 0x2c, 0x56, 0xb6, 0x1f,
	0xb9, 0x36, 0xd9, 0xdb, 0x02, 0xe7, 0xd0, 0x09, 0x68, 0x3c, 0x3e, 0x68, 0x8b, 0x8b, 0xe6, 0xa3,
	0xb4, 0x44, 0x75, 0x9c, 0xbc, 0x5f, 0x32, 0x18, 0x95, 0xad, 0x74, 0x97, 0x93, 0x99, 0xe2, 0x1c,
	0x7a, 0x0d, 0xab, 0x49, 0x68, 0x32, 0x75, 0xa5, 0xdf, 0x45, 0x0b, 0xe7, 0xd0, 0x5b, 0x58, 0xbb,
	0x75, 0x6d, 0x3f, 0xec, 0x7b, 0x34, 0x53, 0x9b, 0xed, 0xf6, 0x19, 0x28, 0x97, 0x84, 0xa2, 0x4d,
	0xde, 0x3c, 0x4d, 0x53, 0x45, 0x4f, 0x0b, 0xb3, 0x93, 0x8e, 0x40, 0x69, 0x46, 0x53, 0x6e, 0x1a,
	0x9a, 0x25, 0xdd, 0x4f, 0x61, 0x45, 0xc4, 0x04, 0x21, 0x2e, 0x5a, 0xc8, 0xcc, 0x12, 0xdd, 0x4b,
	0xd0, 0xf8, 0xd6, 0x65, 0x5e, 0x08, 0x3d, 0xde, 0x4c, 0x9c, 0x3b, 0x92, 0x3e, 0x3c, 0xfd, 0x74,
	0xd0, 0x1b, 0xd0, 0x7e, 0xd4, 0x31, 0xba, 0xde, 0xb8, 0x3a, 0xf6, 0xc2, 0x68, 0x68, 0x57, 0xbb,
	0x84, 0xce, 0xfd, 0x17, 0x0c, 0x27, 0x61, 0x67, 0x85, 0x3f, 0x1d, 0xff, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x1a, 0x6e, 0x23, 0x0d, 0x61, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KVSClient is the client API for KVS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVSClient interface {
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Node(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NodeResponse, error)
	Cluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ClusterResponse, error)
	Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Watch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (KVS_WatchClient, error)
}

type kVSClient struct {
	cc *grpc.ClientConn
}

func NewKVSClient(cc *grpc.ClientConn) KVSClient {
	return &kVSClient{cc}
}

func (c *kVSClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Node(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NodeResponse, error) {
	out := new(NodeResponse)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Node", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Cluster(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ClusterResponse, error) {
	out := new(ClusterResponse)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Cluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Snapshot(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Snapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kvs.KVS/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVSClient) Watch(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (KVS_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KVS_serviceDesc.Streams[0], "/kvs.KVS/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVSWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KVS_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type kVSWatchClient struct {
	grpc.ClientStream
}

func (x *kVSWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVSServer is the server API for KVS service.
type KVSServer interface {
	Join(context.Context, *JoinRequest) (*empty.Empty, error)
	Leave(context.Context, *LeaveRequest) (*empty.Empty, error)
	Node(context.Context, *empty.Empty) (*NodeResponse, error)
	Cluster(context.Context, *empty.Empty) (*ClusterResponse, error)
	Snapshot(context.Context, *empty.Empty) (*empty.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	Watch(*empty.Empty, KVS_WatchServer) error
}

// UnimplementedKVSServer can be embedded to have forward compatible implementations.
type UnimplementedKVSServer struct {
}

func (*UnimplementedKVSServer) Join(ctx context.Context, req *JoinRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedKVSServer) Leave(ctx context.Context, req *LeaveRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (*UnimplementedKVSServer) Node(ctx context.Context, req *empty.Empty) (*NodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Node not implemented")
}
func (*UnimplementedKVSServer) Cluster(ctx context.Context, req *empty.Empty) (*ClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cluster not implemented")
}
func (*UnimplementedKVSServer) Snapshot(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (*UnimplementedKVSServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVSServer) Put(ctx context.Context, req *PutRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedKVSServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedKVSServer) Watch(req *empty.Empty, srv KVS_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterKVSServer(s *grpc.Server, srv KVSServer) {
	s.RegisterService(&_KVS_serviceDesc, srv)
}

func _KVS_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Node_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Node(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Node",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Node(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Cluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Cluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Cluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Cluster(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Snapshot(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVSServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvs.KVS/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVSServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVS_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVSServer).Watch(m, &kVSWatchServer{stream})
}

type KVS_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type kVSWatchServer struct {
	grpc.ServerStream
}

func (x *kVSWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _KVS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kvs.KVS",
	HandlerType: (*KVSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _KVS_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _KVS_Leave_Handler,
		},
		{
			MethodName: "Node",
			Handler:    _KVS_Node_Handler,
		},
		{
			MethodName: "Cluster",
			Handler:    _KVS_Cluster_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _KVS_Snapshot_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KVS_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KVS_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KVS_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _KVS_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/kvs/kvs.proto",
}