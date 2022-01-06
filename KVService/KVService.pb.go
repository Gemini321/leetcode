// Code generated by protoc-gen-go. DO NOT EDIT.
// source: KVService.proto

package KVService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type KV struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{0}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	ReturnKV             *KV      `protobuf:"bytes,2,opt,name=ReturnKV,proto3" json:"ReturnKV,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetReturnKV() *KV {
	if m != nil {
		return m.ReturnKV
	}
	return nil
}

type RequestVoteArgs struct {
	Term                 int64    `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	CandidateId          int64    `protobuf:"varint,2,opt,name=CandidateId,proto3" json:"CandidateId,omitempty"`
	LastLogIndex         int64    `protobuf:"varint,3,opt,name=LastLogIndex,proto3" json:"LastLogIndex,omitempty"`
	LastLogTerm          int64    `protobuf:"varint,4,opt,name=LastLogTerm,proto3" json:"LastLogTerm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteArgs) Reset()         { *m = RequestVoteArgs{} }
func (m *RequestVoteArgs) String() string { return proto.CompactTextString(m) }
func (*RequestVoteArgs) ProtoMessage()    {}
func (*RequestVoteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{2}
}

func (m *RequestVoteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteArgs.Unmarshal(m, b)
}
func (m *RequestVoteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteArgs.Marshal(b, m, deterministic)
}
func (m *RequestVoteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteArgs.Merge(m, src)
}
func (m *RequestVoteArgs) XXX_Size() int {
	return xxx_messageInfo_RequestVoteArgs.Size(m)
}
func (m *RequestVoteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteArgs proto.InternalMessageInfo

func (m *RequestVoteArgs) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteArgs) GetCandidateId() int64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVoteArgs) GetLastLogIndex() int64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *RequestVoteArgs) GetLastLogTerm() int64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

type RequestVoteReply struct {
	Term                 int64    `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	VoteGranted          bool     `protobuf:"varint,2,opt,name=VoteGranted,proto3" json:"VoteGranted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteReply) Reset()         { *m = RequestVoteReply{} }
func (m *RequestVoteReply) String() string { return proto.CompactTextString(m) }
func (*RequestVoteReply) ProtoMessage()    {}
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{3}
}

func (m *RequestVoteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteReply.Unmarshal(m, b)
}
func (m *RequestVoteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteReply.Marshal(b, m, deterministic)
}
func (m *RequestVoteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteReply.Merge(m, src)
}
func (m *RequestVoteReply) XXX_Size() int {
	return xxx_messageInfo_RequestVoteReply.Size(m)
}
func (m *RequestVoteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteReply.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteReply proto.InternalMessageInfo

func (m *RequestVoteReply) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteReply) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

type LogEntry struct {
	Command              []byte   `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	Term                 int64    `protobuf:"varint,2,opt,name=Term,proto3" json:"Term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{4}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetCommand() []byte {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *LogEntry) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type AppendEntriesArg struct {
	Term                 int64       `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	LeaderId             int64       `protobuf:"varint,2,opt,name=LeaderId,proto3" json:"LeaderId,omitempty"`
	PrevLogIndex         int64       `protobuf:"varint,3,opt,name=PrevLogIndex,proto3" json:"PrevLogIndex,omitempty"`
	PrevLogTerm          int64       `protobuf:"varint,4,opt,name=PrevLogTerm,proto3" json:"PrevLogTerm,omitempty"`
	Entries              []*LogEntry `protobuf:"bytes,5,rep,name=Entries,proto3" json:"Entries,omitempty"`
	LeaderCommit         int64       `protobuf:"varint,6,opt,name=LeaderCommit,proto3" json:"LeaderCommit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppendEntriesArg) Reset()         { *m = AppendEntriesArg{} }
func (m *AppendEntriesArg) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesArg) ProtoMessage()    {}
func (*AppendEntriesArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{5}
}

func (m *AppendEntriesArg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesArg.Unmarshal(m, b)
}
func (m *AppendEntriesArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesArg.Marshal(b, m, deterministic)
}
func (m *AppendEntriesArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesArg.Merge(m, src)
}
func (m *AppendEntriesArg) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesArg.Size(m)
}
func (m *AppendEntriesArg) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesArg.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesArg proto.InternalMessageInfo

func (m *AppendEntriesArg) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesArg) GetLeaderId() int64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *AppendEntriesArg) GetPrevLogIndex() int64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesArg) GetPrevLogTerm() int64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesArg) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AppendEntriesArg) GetLeaderCommit() int64 {
	if m != nil {
		return m.LeaderCommit
	}
	return 0
}

type AppendEntriesReply struct {
	Term                 int64    `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntriesReply) Reset()         { *m = AppendEntriesReply{} }
func (m *AppendEntriesReply) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesReply) ProtoMessage()    {}
func (*AppendEntriesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{6}
}

func (m *AppendEntriesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesReply.Unmarshal(m, b)
}
func (m *AppendEntriesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesReply.Marshal(b, m, deterministic)
}
func (m *AppendEntriesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesReply.Merge(m, src)
}
func (m *AppendEntriesReply) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesReply.Size(m)
}
func (m *AppendEntriesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesReply.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesReply proto.InternalMessageInfo

func (m *AppendEntriesReply) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type LeaderID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaderID) Reset()         { *m = LeaderID{} }
func (m *LeaderID) String() string { return proto.CompactTextString(m) }
func (*LeaderID) ProtoMessage()    {}
func (*LeaderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{7}
}

func (m *LeaderID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaderID.Unmarshal(m, b)
}
func (m *LeaderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaderID.Marshal(b, m, deterministic)
}
func (m *LeaderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaderID.Merge(m, src)
}
func (m *LeaderID) XXX_Size() int {
	return xxx_messageInfo_LeaderID.Size(m)
}
func (m *LeaderID) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaderID.DiscardUnknown(m)
}

var xxx_messageInfo_LeaderID proto.InternalMessageInfo

func (m *LeaderID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type State struct {
	State                int64    `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_2113408249da62ec, []int{8}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

func init() {
	proto.RegisterType((*KV)(nil), "KVService.KV")
	proto.RegisterType((*Response)(nil), "KVService.Response")
	proto.RegisterType((*RequestVoteArgs)(nil), "KVService.RequestVoteArgs")
	proto.RegisterType((*RequestVoteReply)(nil), "KVService.RequestVoteReply")
	proto.RegisterType((*LogEntry)(nil), "KVService.LogEntry")
	proto.RegisterType((*AppendEntriesArg)(nil), "KVService.AppendEntriesArg")
	proto.RegisterType((*AppendEntriesReply)(nil), "KVService.AppendEntriesReply")
	proto.RegisterType((*LeaderID)(nil), "KVService.LeaderID")
	proto.RegisterType((*State)(nil), "KVService.State")
}

func init() { proto.RegisterFile("KVService.proto", fileDescriptor_2113408249da62ec) }

var fileDescriptor_2113408249da62ec = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0x06, 0x1c, 0x7e, 0x32, 0x0e, 0x0d, 0xda, 0xf6, 0x60, 0xb9, 0x8a, 0x84, 0xf6, 0x94, 0xaa,
	0x2d, 0xad, 0xe8, 0xa5, 0x3d, 0x92, 0xa4, 0x02, 0x04, 0x91, 0xa2, 0xa5, 0xf2, 0xa1, 0x37, 0x97,
	0x1d, 0x21, 0x24, 0xb0, 0xe9, 0xee, 0x3a, 0x2a, 0x0f, 0xd1, 0x07, 0xe8, 0xb3, 0xf5, 0x65, 0xaa,
	0x5d, 0xaf, 0xcd, 0x3a, 0x05, 0x29, 0xb7, 0x9d, 0x6f, 0xbe, 0x19, 0xcf, 0xf7, 0xcd, 0x00, 0x5c,
	0xce, 0xa2, 0x05, 0x8a, 0xc7, 0xf5, 0x12, 0x07, 0x3b, 0x91, 0xaa, 0x94, 0x9c, 0x97, 0x00, 0x7d,
	0x07, 0x8d, 0x59, 0x44, 0x7a, 0xe0, 0xcd, 0x70, 0x1f, 0xd4, 0xfb, 0xf5, 0xeb, 0x0b, 0xa6, 0x9f,
	0xe4, 0x15, 0x34, 0xa3, 0x78, 0x93, 0x61, 0xd0, 0x30, 0x58, 0x1e, 0xd0, 0x31, 0x74, 0x18, 0xca,
	0x5d, 0x9a, 0x48, 0xd4, 0x35, 0xf7, 0x72, 0x65, 0x6a, 0xce, 0x99, 0x7e, 0x92, 0x37, 0x3a, 0xab,
	0x32, 0x91, 0xcc, 0x22, 0x53, 0xe6, 0x0f, 0xbb, 0x83, 0xc3, 0xa7, 0x67, 0x11, 0x2b, 0xd3, 0xf4,
	0x77, 0x1d, 0x2e, 0x19, 0xfe, 0xcc, 0x50, 0xaa, 0x28, 0x55, 0x38, 0x12, 0x2b, 0x49, 0x08, 0x9c,
	0x7d, 0x43, 0xb1, 0x35, 0x1d, 0x3d, 0x66, 0xde, 0xa4, 0x0f, 0xfe, 0x6d, 0x9c, 0xf0, 0x35, 0x8f,
	0x15, 0x4e, 0xb9, 0xe9, 0xea, 0x31, 0x17, 0x22, 0x14, 0x2e, 0xe6, 0xb1, 0x54, 0xf3, 0x74, 0x35,
	0x4d, 0x38, 0xfe, 0x0a, 0x3c, 0x43, 0xa9, 0x60, 0xba, 0x8b, 0x8d, 0xcd, 0x07, 0xce, 0xf2, 0x2e,
	0x0e, 0x44, 0x27, 0xd0, 0x73, 0xc6, 0x61, 0xb8, 0xdb, 0xec, 0x4f, 0xcd, 0xa3, 0x09, 0x63, 0x11,
	0x27, 0x0a, 0xf3, 0x79, 0x3a, 0xcc, 0x85, 0xe8, 0x67, 0xe8, 0xcc, 0xd3, 0xd5, 0xd7, 0x44, 0x89,
	0x3d, 0x09, 0xa0, 0x7d, 0x9b, 0x6e, 0xb7, 0x71, 0xc2, 0xad, 0xb5, 0x45, 0x58, 0xf6, 0x6e, 0x1c,
	0x7a, 0xd3, 0xbf, 0x75, 0xe8, 0x8d, 0x76, 0x3b, 0x4c, 0xb8, 0xae, 0x5e, 0xa3, 0x1c, 0x89, 0xd5,
	0xd1, 0x21, 0x42, 0xe8, 0xcc, 0x31, 0xe6, 0x28, 0x4a, 0x47, 0xca, 0x58, 0xdb, 0xf1, 0x20, 0xf0,
	0xf1, 0xa9, 0x1d, 0x2e, 0xa6, 0x45, 0xd8, 0xd8, 0xb5, 0xc3, 0x81, 0xc8, 0x7b, 0x68, 0xdb, 0x19,
	0x82, 0x66, 0xdf, 0xbb, 0xf6, 0x87, 0x2f, 0x9d, 0x45, 0x16, 0xf2, 0x58, 0xc1, 0x31, 0x3b, 0x30,
	0x03, 0x68, 0x79, 0x6b, 0x15, 0xb4, 0xec, 0x0e, 0x1c, 0x8c, 0xde, 0x00, 0xa9, 0x88, 0x3b, 0xed,
	0x71, 0x00, 0xed, 0x45, 0xb6, 0x5c, 0xa2, 0x94, 0xd6, 0xdf, 0x22, 0xa4, 0x07, 0xe1, 0x77, 0xe4,
	0x05, 0x34, 0xa6, 0xdc, 0xd6, 0x35, 0xa6, 0x9c, 0x5e, 0x41, 0x73, 0xa1, 0x62, 0x85, 0xfa, 0x72,
	0xa5, 0x7e, 0xd8, 0x5c, 0x1e, 0x0c, 0xff, 0x78, 0xd0, 0x9e, 0xc4, 0x09, 0xdf, 0xa0, 0x20, 0x6f,
	0xc1, 0x7b, 0xc8, 0x14, 0xa9, 0x1e, 0x67, 0xe8, 0x4a, 0x2c, 0x8e, 0x9c, 0xd6, 0x34, 0x79, 0x8c,
	0xcf, 0x25, 0x0f, 0xa0, 0x75, 0x87, 0x1b, 0x54, 0xf8, 0x4c, 0xfe, 0x17, 0xf0, 0xc7, 0xa8, 0x4a,
	0x4d, 0x15, 0x97, 0x2d, 0x18, 0x1e, 0x03, 0x69, 0x8d, 0x7c, 0x84, 0xce, 0x18, 0x55, 0x2e, 0xb9,
	0xe7, 0x50, 0x0c, 0x12, 0xfe, 0x87, 0xd0, 0x1a, 0x99, 0x80, 0xef, 0xdc, 0x38, 0x09, 0x2b, 0x23,
	0x55, 0x7e, 0x8a, 0xe1, 0xeb, 0xe3, 0x39, 0xb3, 0x33, 0x5a, 0x23, 0xf7, 0xd0, 0xad, 0xec, 0x92,
	0xb8, 0xfc, 0xa7, 0x27, 0x1c, 0x5e, 0x9d, 0x4a, 0xda, 0x76, 0x37, 0xdd, 0xef, 0xfe, 0xe0, 0x43,
	0xc9, 0xf9, 0xd1, 0x32, 0x7f, 0x52, 0x9f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xd5, 0xc6,
	0x1e, 0xb7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HandlerClient is the client API for Handler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HandlerClient interface {
	Put(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Response, error)
	GetLeaderID(ctx context.Context, in *LeaderID, opts ...grpc.CallOption) (*LeaderID, error)
	GetState(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error)
	RequestVote(ctx context.Context, in *RequestVoteArgs, opts ...grpc.CallOption) (*RequestVoteReply, error)
	AppendEntries(ctx context.Context, in *AppendEntriesArg, opts ...grpc.CallOption) (*AppendEntriesReply, error)
}

type handlerClient struct {
	cc *grpc.ClientConn
}

func NewHandlerClient(cc *grpc.ClientConn) HandlerClient {
	return &handlerClient{cc}
}

func (c *handlerClient) Put(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/KVService.Handler/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) Get(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/KVService.Handler/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) Delete(ctx context.Context, in *KV, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/KVService.Handler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) GetLeaderID(ctx context.Context, in *LeaderID, opts ...grpc.CallOption) (*LeaderID, error) {
	out := new(LeaderID)
	err := c.cc.Invoke(ctx, "/KVService.Handler/GetLeaderID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) GetState(ctx context.Context, in *State, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/KVService.Handler/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) RequestVote(ctx context.Context, in *RequestVoteArgs, opts ...grpc.CallOption) (*RequestVoteReply, error) {
	out := new(RequestVoteReply)
	err := c.cc.Invoke(ctx, "/KVService.Handler/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerClient) AppendEntries(ctx context.Context, in *AppendEntriesArg, opts ...grpc.CallOption) (*AppendEntriesReply, error) {
	out := new(AppendEntriesReply)
	err := c.cc.Invoke(ctx, "/KVService.Handler/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandlerServer is the server API for Handler service.
type HandlerServer interface {
	Put(context.Context, *KV) (*Response, error)
	Get(context.Context, *KV) (*Response, error)
	Delete(context.Context, *KV) (*Response, error)
	GetLeaderID(context.Context, *LeaderID) (*LeaderID, error)
	GetState(context.Context, *State) (*State, error)
	RequestVote(context.Context, *RequestVoteArgs) (*RequestVoteReply, error)
	AppendEntries(context.Context, *AppendEntriesArg) (*AppendEntriesReply, error)
}

// UnimplementedHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedHandlerServer struct {
}

func (*UnimplementedHandlerServer) Put(ctx context.Context, req *KV) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedHandlerServer) Get(ctx context.Context, req *KV) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedHandlerServer) Delete(ctx context.Context, req *KV) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedHandlerServer) GetLeaderID(ctx context.Context, req *LeaderID) (*LeaderID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderID not implemented")
}
func (*UnimplementedHandlerServer) GetState(ctx context.Context, req *State) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (*UnimplementedHandlerServer) RequestVote(ctx context.Context, req *RequestVoteArgs) (*RequestVoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (*UnimplementedHandlerServer) AppendEntries(ctx context.Context, req *AppendEntriesArg) (*AppendEntriesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}

func RegisterHandlerServer(s *grpc.Server, srv HandlerServer) {
	s.RegisterService(&_Handler_serviceDesc, srv)
}

func _Handler_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVService.Handler/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).Put(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

func _Handler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVService.Handler/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).Get(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

func _Handler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVService.Handler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).Delete(ctx, req.(*KV))
	}
	return interceptor(ctx, in, info, handler)
}

func _Handler_GetLeaderID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).GetLeaderID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVService.Handler/GetLeaderID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).GetLeaderID(ctx, req.(*LeaderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Handler_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(State)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVService.Handler/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).GetState(ctx, req.(*State))
	}
	return interceptor(ctx, in, info, handler)
}

func _Handler_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVService.Handler/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).RequestVote(ctx, req.(*RequestVoteArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Handler_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KVService.Handler/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerServer).AppendEntries(ctx, req.(*AppendEntriesArg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Handler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KVService.Handler",
	HandlerType: (*HandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Handler_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Handler_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Handler_Delete_Handler,
		},
		{
			MethodName: "GetLeaderID",
			Handler:    _Handler_GetLeaderID_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Handler_GetState_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _Handler_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _Handler_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "KVService.proto",
}
