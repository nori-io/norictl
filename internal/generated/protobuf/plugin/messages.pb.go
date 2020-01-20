// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package protoNori

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// cmd: plugin get
type PluginGetRequest struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagVerbose          bool     `protobuf:"varint,3,opt,name=FlagVerbose,json=flagVerbose,proto3" json:"FlagVerbose,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginGetRequest) Reset()         { *m = PluginGetRequest{} }
func (m *PluginGetRequest) String() string { return proto.CompactTextString(m) }
func (*PluginGetRequest) ProtoMessage()    {}
func (*PluginGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *PluginGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginGetRequest.Unmarshal(m, b)
}
func (m *PluginGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginGetRequest.Marshal(b, m, deterministic)
}
func (m *PluginGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginGetRequest.Merge(m, src)
}
func (m *PluginGetRequest) XXX_Size() int {
	return xxx_messageInfo_PluginGetRequest.Size(m)
}
func (m *PluginGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginGetRequest proto.InternalMessageInfo

func (m *PluginGetRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PluginGetRequest) GetFlagVerbose() bool {
	if m != nil {
		return m.FlagVerbose
	}
	return false
}

type PluginPullRequest struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagDeps             bool     `protobuf:"varint,2,opt,name=FlagDeps,json=flagDeps,proto3" json:"FlagDeps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginPullRequest) Reset()         { *m = PluginPullRequest{} }
func (m *PluginPullRequest) String() string { return proto.CompactTextString(m) }
func (*PluginPullRequest) ProtoMessage()    {}
func (*PluginPullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *PluginPullRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginPullRequest.Unmarshal(m, b)
}
func (m *PluginPullRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginPullRequest.Marshal(b, m, deterministic)
}
func (m *PluginPullRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginPullRequest.Merge(m, src)
}
func (m *PluginPullRequest) XXX_Size() int {
	return xxx_messageInfo_PluginPullRequest.Size(m)
}
func (m *PluginPullRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginPullRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginPullRequest proto.InternalMessageInfo

func (m *PluginPullRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PluginPullRequest) GetFlagDeps() bool {
	if m != nil {
		return m.FlagDeps
	}
	return false
}

type ErrorReply struct {
	Status               bool     `protobuf:"varint,1,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,json=error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorReply) Reset()         { *m = ErrorReply{} }
func (m *ErrorReply) String() string { return proto.CompactTextString(m) }
func (*ErrorReply) ProtoMessage()    {}
func (*ErrorReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *ErrorReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorReply.Unmarshal(m, b)
}
func (m *ErrorReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorReply.Marshal(b, m, deterministic)
}
func (m *ErrorReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorReply.Merge(m, src)
}
func (m *ErrorReply) XXX_Size() int {
	return xxx_messageInfo_ErrorReply.Size(m)
}
func (m *ErrorReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorReply.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorReply proto.InternalMessageInfo

func (m *ErrorReply) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ErrorReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// cmd: plugin install
type PluginInstallRequest struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagVerbose          bool     `protobuf:"varint,2,opt,name=FlagVerbose,json=flagVerbose,proto3" json:"FlagVerbose,omitempty"`
	FlagDeps             bool     `protobuf:"varint,3,opt,name=FlagDeps,json=flagDeps,proto3" json:"FlagDeps,omitempty"`
	FlagAll              bool     `protobuf:"varint,4,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInstallRequest) Reset()         { *m = PluginInstallRequest{} }
func (m *PluginInstallRequest) String() string { return proto.CompactTextString(m) }
func (*PluginInstallRequest) ProtoMessage()    {}
func (*PluginInstallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *PluginInstallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInstallRequest.Unmarshal(m, b)
}
func (m *PluginInstallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInstallRequest.Marshal(b, m, deterministic)
}
func (m *PluginInstallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInstallRequest.Merge(m, src)
}
func (m *PluginInstallRequest) XXX_Size() int {
	return xxx_messageInfo_PluginInstallRequest.Size(m)
}
func (m *PluginInstallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInstallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInstallRequest proto.InternalMessageInfo

func (m *PluginInstallRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PluginInstallRequest) GetFlagVerbose() bool {
	if m != nil {
		return m.FlagVerbose
	}
	return false
}

func (m *PluginInstallRequest) GetFlagDeps() bool {
	if m != nil {
		return m.FlagDeps
	}
	return false
}

func (m *PluginInstallRequest) GetFlagAll() bool {
	if m != nil {
		return m.FlagAll
	}
	return false
}

// cmd: plugin list
type PluginListRequest struct {
	FlagAll              bool     `protobuf:"varint,1,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	FlagError            bool     `protobuf:"varint,2,opt,name=FlagError,json=flagError,proto3" json:"FlagError,omitempty"`
	Installed            bool     `protobuf:"varint,3,opt,name=Installed,json=installed,proto3" json:"Installed,omitempty"`
	Running              bool     `protobuf:"varint,4,opt,name=Running,json=running,proto3" json:"Running,omitempty"`
	Installable          bool     `protobuf:"varint,5,opt,name=Installable,json=installable,proto3" json:"Installable,omitempty"`
	Inactive             bool     `protobuf:"varint,6,opt,name=Inactive,json=inactive,proto3" json:"Inactive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginListRequest) Reset()         { *m = PluginListRequest{} }
func (m *PluginListRequest) String() string { return proto.CompactTextString(m) }
func (*PluginListRequest) ProtoMessage()    {}
func (*PluginListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *PluginListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginListRequest.Unmarshal(m, b)
}
func (m *PluginListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginListRequest.Marshal(b, m, deterministic)
}
func (m *PluginListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginListRequest.Merge(m, src)
}
func (m *PluginListRequest) XXX_Size() int {
	return xxx_messageInfo_PluginListRequest.Size(m)
}
func (m *PluginListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginListRequest proto.InternalMessageInfo

func (m *PluginListRequest) GetFlagAll() bool {
	if m != nil {
		return m.FlagAll
	}
	return false
}

func (m *PluginListRequest) GetFlagError() bool {
	if m != nil {
		return m.FlagError
	}
	return false
}

func (m *PluginListRequest) GetInstalled() bool {
	if m != nil {
		return m.Installed
	}
	return false
}

func (m *PluginListRequest) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *PluginListRequest) GetInstallable() bool {
	if m != nil {
		return m.Installable
	}
	return false
}

func (m *PluginListRequest) GetInactive() bool {
	if m != nil {
		return m.Inactive
	}
	return false
}

type PluginListWithStatus struct {
	MetaID               *ID           `protobuf:"bytes,1,opt,name=metaID,proto3" json:"metaID,omitempty"`
	Author               *Author       `protobuf:"bytes,2,opt,name=Author,json=author,proto3" json:"Author,omitempty"`
	DependenciesArray    *Dependencies `protobuf:"bytes,3,opt,name=DependenciesArray,json=dependenciesArray,proto3" json:"DependenciesArray,omitempty"`
	Description          *Description  `protobuf:"bytes,4,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	Core                 *Core         `protobuf:"bytes,5,opt,name=Core,json=core,proto3" json:"Core,omitempty"`
	Interface            *Interface    `protobuf:"bytes,6,opt,name=Interface,json=interface,proto3" json:"Interface,omitempty"`
	License              *License      `protobuf:"bytes,7,opt,name=License,json=license,proto3" json:"License,omitempty"`
	Links                *Links        `protobuf:"bytes,8,opt,name=Links,json=links,proto3" json:"Links,omitempty"`
	Repository           *Repository   `protobuf:"bytes,9,opt,name=Repository,json=repository,proto3" json:"Repository,omitempty"`
	Tags                 []string      `protobuf:"bytes,10,rep,name=Tags,json=tags,proto3" json:"Tags,omitempty"`
	FlagAll              bool          `protobuf:"varint,11,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	FlagError            bool          `protobuf:"varint,12,opt,name=FlagError,json=flagError,proto3" json:"FlagError,omitempty"`
	FlagInstalled        bool          `protobuf:"varint,13,opt,name=FlagInstalled,json=flagInstalled,proto3" json:"FlagInstalled,omitempty"`
	FlagRunning          bool          `protobuf:"varint,14,opt,name=FlagRunning,json=flagRunning,proto3" json:"FlagRunning,omitempty"`
	FlagInstallable      bool          `protobuf:"varint,15,opt,name=FlagInstallable,json=flagInstallable,proto3" json:"FlagInstallable,omitempty"`
	FlagInactive         bool          `protobuf:"varint,16,opt,name=FlagInactive,json=flagInactive,proto3" json:"FlagInactive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PluginListWithStatus) Reset()         { *m = PluginListWithStatus{} }
func (m *PluginListWithStatus) String() string { return proto.CompactTextString(m) }
func (*PluginListWithStatus) ProtoMessage()    {}
func (*PluginListWithStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{5}
}

func (m *PluginListWithStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginListWithStatus.Unmarshal(m, b)
}
func (m *PluginListWithStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginListWithStatus.Marshal(b, m, deterministic)
}
func (m *PluginListWithStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginListWithStatus.Merge(m, src)
}
func (m *PluginListWithStatus) XXX_Size() int {
	return xxx_messageInfo_PluginListWithStatus.Size(m)
}
func (m *PluginListWithStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginListWithStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PluginListWithStatus proto.InternalMessageInfo

func (m *PluginListWithStatus) GetMetaID() *ID {
	if m != nil {
		return m.MetaID
	}
	return nil
}

func (m *PluginListWithStatus) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *PluginListWithStatus) GetDependenciesArray() *Dependencies {
	if m != nil {
		return m.DependenciesArray
	}
	return nil
}

func (m *PluginListWithStatus) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *PluginListWithStatus) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func (m *PluginListWithStatus) GetInterface() *Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *PluginListWithStatus) GetLicense() *License {
	if m != nil {
		return m.License
	}
	return nil
}

func (m *PluginListWithStatus) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PluginListWithStatus) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PluginListWithStatus) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PluginListWithStatus) GetFlagAll() bool {
	if m != nil {
		return m.FlagAll
	}
	return false
}

func (m *PluginListWithStatus) GetFlagError() bool {
	if m != nil {
		return m.FlagError
	}
	return false
}

func (m *PluginListWithStatus) GetFlagInstalled() bool {
	if m != nil {
		return m.FlagInstalled
	}
	return false
}

func (m *PluginListWithStatus) GetFlagRunning() bool {
	if m != nil {
		return m.FlagRunning
	}
	return false
}

func (m *PluginListWithStatus) GetFlagInstallable() bool {
	if m != nil {
		return m.FlagInstallable
	}
	return false
}

func (m *PluginListWithStatus) GetFlagInactive() bool {
	if m != nil {
		return m.FlagInactive
	}
	return false
}

type PluginListWithoutStatus struct {
	MetaID               *ID           `protobuf:"bytes,1,opt,name=metaID,proto3" json:"metaID,omitempty"`
	Author               *Author       `protobuf:"bytes,2,opt,name=Author,json=author,proto3" json:"Author,omitempty"`
	DependenciesArray    *Dependencies `protobuf:"bytes,3,opt,name=DependenciesArray,json=dependenciesArray,proto3" json:"DependenciesArray,omitempty"`
	Description          *Description  `protobuf:"bytes,4,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	Core                 *Core         `protobuf:"bytes,5,opt,name=Core,json=core,proto3" json:"Core,omitempty"`
	Interface            *Interface    `protobuf:"bytes,6,opt,name=Interface,json=interface,proto3" json:"Interface,omitempty"`
	License              *License      `protobuf:"bytes,7,opt,name=License,json=license,proto3" json:"License,omitempty"`
	Links                *Links        `protobuf:"bytes,8,opt,name=Links,json=links,proto3" json:"Links,omitempty"`
	Repository           *Repository   `protobuf:"bytes,9,opt,name=Repository,json=repository,proto3" json:"Repository,omitempty"`
	Tags                 []string      `protobuf:"bytes,10,rep,name=Tags,json=tags,proto3" json:"Tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PluginListWithoutStatus) Reset()         { *m = PluginListWithoutStatus{} }
func (m *PluginListWithoutStatus) String() string { return proto.CompactTextString(m) }
func (*PluginListWithoutStatus) ProtoMessage()    {}
func (*PluginListWithoutStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{6}
}

func (m *PluginListWithoutStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginListWithoutStatus.Unmarshal(m, b)
}
func (m *PluginListWithoutStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginListWithoutStatus.Marshal(b, m, deterministic)
}
func (m *PluginListWithoutStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginListWithoutStatus.Merge(m, src)
}
func (m *PluginListWithoutStatus) XXX_Size() int {
	return xxx_messageInfo_PluginListWithoutStatus.Size(m)
}
func (m *PluginListWithoutStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginListWithoutStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PluginListWithoutStatus proto.InternalMessageInfo

func (m *PluginListWithoutStatus) GetMetaID() *ID {
	if m != nil {
		return m.MetaID
	}
	return nil
}

func (m *PluginListWithoutStatus) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *PluginListWithoutStatus) GetDependenciesArray() *Dependencies {
	if m != nil {
		return m.DependenciesArray
	}
	return nil
}

func (m *PluginListWithoutStatus) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *PluginListWithoutStatus) GetCore() *Core {
	if m != nil {
		return m.Core
	}
	return nil
}

func (m *PluginListWithoutStatus) GetInterface() *Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *PluginListWithoutStatus) GetLicense() *License {
	if m != nil {
		return m.License
	}
	return nil
}

func (m *PluginListWithoutStatus) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PluginListWithoutStatus) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PluginListWithoutStatus) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type PluginListReply struct {
	Data                 []*PluginListWithStatus `protobuf:"bytes,1,rep,name=Data,json=data,proto3" json:"Data,omitempty"`
	ErrorPlugin          *ErrorReply             `protobuf:"bytes,2,opt,name=ErrorPlugin,json=errorPlugin,proto3" json:"ErrorPlugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PluginListReply) Reset()         { *m = PluginListReply{} }
func (m *PluginListReply) String() string { return proto.CompactTextString(m) }
func (*PluginListReply) ProtoMessage()    {}
func (*PluginListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{7}
}

func (m *PluginListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginListReply.Unmarshal(m, b)
}
func (m *PluginListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginListReply.Marshal(b, m, deterministic)
}
func (m *PluginListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginListReply.Merge(m, src)
}
func (m *PluginListReply) XXX_Size() int {
	return xxx_messageInfo_PluginListReply.Size(m)
}
func (m *PluginListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginListReply.DiscardUnknown(m)
}

var xxx_messageInfo_PluginListReply proto.InternalMessageInfo

func (m *PluginListReply) GetData() []*PluginListWithStatus {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PluginListReply) GetErrorPlugin() *ErrorReply {
	if m != nil {
		return m.ErrorPlugin
	}
	return nil
}

// cmd: plugin remove
type PluginRemoveRequest struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginRemoveRequest) Reset()         { *m = PluginRemoveRequest{} }
func (m *PluginRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*PluginRemoveRequest) ProtoMessage()    {}
func (*PluginRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{8}
}

func (m *PluginRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginRemoveRequest.Unmarshal(m, b)
}
func (m *PluginRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginRemoveRequest.Marshal(b, m, deterministic)
}
func (m *PluginRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginRemoveRequest.Merge(m, src)
}
func (m *PluginRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_PluginRemoveRequest.Size(m)
}
func (m *PluginRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginRemoveRequest proto.InternalMessageInfo

func (m *PluginRemoveRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

type PluginStartRequest struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagAll              bool     `protobuf:"varint,2,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginStartRequest) Reset()         { *m = PluginStartRequest{} }
func (m *PluginStartRequest) String() string { return proto.CompactTextString(m) }
func (*PluginStartRequest) ProtoMessage()    {}
func (*PluginStartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{9}
}

func (m *PluginStartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginStartRequest.Unmarshal(m, b)
}
func (m *PluginStartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginStartRequest.Marshal(b, m, deterministic)
}
func (m *PluginStartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginStartRequest.Merge(m, src)
}
func (m *PluginStartRequest) XXX_Size() int {
	return xxx_messageInfo_PluginStartRequest.Size(m)
}
func (m *PluginStartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginStartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginStartRequest proto.InternalMessageInfo

func (m *PluginStartRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PluginStartRequest) GetFlagAll() bool {
	if m != nil {
		return m.FlagAll
	}
	return false
}

type PluginStopRequest struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagAll              bool     `protobuf:"varint,2,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginStopRequest) Reset()         { *m = PluginStopRequest{} }
func (m *PluginStopRequest) String() string { return proto.CompactTextString(m) }
func (*PluginStopRequest) ProtoMessage()    {}
func (*PluginStopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{10}
}

func (m *PluginStopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginStopRequest.Unmarshal(m, b)
}
func (m *PluginStopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginStopRequest.Marshal(b, m, deterministic)
}
func (m *PluginStopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginStopRequest.Merge(m, src)
}
func (m *PluginStopRequest) XXX_Size() int {
	return xxx_messageInfo_PluginStopRequest.Size(m)
}
func (m *PluginStopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginStopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginStopRequest proto.InternalMessageInfo

func (m *PluginStopRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PluginStopRequest) GetFlagAll() bool {
	if m != nil {
		return m.FlagAll
	}
	return false
}

// cmd: plugin uninstall
type PluginUninstallRequest struct {
	Id                   *ID      `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagAll              bool     `protobuf:"varint,2,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	FlagDependent        bool     `protobuf:"varint,3,opt,name=FlagDependent,json=flagDependent,proto3" json:"FlagDependent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginUninstallRequest) Reset()         { *m = PluginUninstallRequest{} }
func (m *PluginUninstallRequest) String() string { return proto.CompactTextString(m) }
func (*PluginUninstallRequest) ProtoMessage()    {}
func (*PluginUninstallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{11}
}

func (m *PluginUninstallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginUninstallRequest.Unmarshal(m, b)
}
func (m *PluginUninstallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginUninstallRequest.Marshal(b, m, deterministic)
}
func (m *PluginUninstallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginUninstallRequest.Merge(m, src)
}
func (m *PluginUninstallRequest) XXX_Size() int {
	return xxx_messageInfo_PluginUninstallRequest.Size(m)
}
func (m *PluginUninstallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginUninstallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginUninstallRequest proto.InternalMessageInfo

func (m *PluginUninstallRequest) GetId() *ID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PluginUninstallRequest) GetFlagAll() bool {
	if m != nil {
		return m.FlagAll
	}
	return false
}

func (m *PluginUninstallRequest) GetFlagDependent() bool {
	if m != nil {
		return m.FlagDependent
	}
	return false
}

// cmd: plugin upload
type PluginUploadRequest struct {
	Filepath             string   `protobuf:"bytes,1,opt,name=Filepath,json=filepath,proto3" json:"Filepath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginUploadRequest) Reset()         { *m = PluginUploadRequest{} }
func (m *PluginUploadRequest) String() string { return proto.CompactTextString(m) }
func (*PluginUploadRequest) ProtoMessage()    {}
func (*PluginUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{12}
}

func (m *PluginUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginUploadRequest.Unmarshal(m, b)
}
func (m *PluginUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginUploadRequest.Marshal(b, m, deterministic)
}
func (m *PluginUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginUploadRequest.Merge(m, src)
}
func (m *PluginUploadRequest) XXX_Size() int {
	return xxx_messageInfo_PluginUploadRequest.Size(m)
}
func (m *PluginUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginUploadRequest proto.InternalMessageInfo

func (m *PluginUploadRequest) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

// cmd: plugin meta
type PluginMetaRequest struct {
	ID                   *PluginID `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	FlagDeps             bool      `protobuf:"varint,2,opt,name=FlagDeps,json=flagDeps,proto3" json:"FlagDeps,omitempty"`
	FlagDepsStatus       bool      `protobuf:"varint,3,opt,name=FlagDepsStatus,json=flagDepsStatus,proto3" json:"FlagDepsStatus,omitempty"`
	FlagDependent        bool      `protobuf:"varint,4,opt,name=FlagDependent,json=flagDependent,proto3" json:"FlagDependent,omitempty"`
	FlagDependentStatus  bool      `protobuf:"varint,5,opt,name=FlagDependentStatus,json=flagDependentStatus,proto3" json:"FlagDependentStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PluginMetaRequest) Reset()         { *m = PluginMetaRequest{} }
func (m *PluginMetaRequest) String() string { return proto.CompactTextString(m) }
func (*PluginMetaRequest) ProtoMessage()    {}
func (*PluginMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{13}
}

func (m *PluginMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginMetaRequest.Unmarshal(m, b)
}
func (m *PluginMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginMetaRequest.Marshal(b, m, deterministic)
}
func (m *PluginMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginMetaRequest.Merge(m, src)
}
func (m *PluginMetaRequest) XXX_Size() int {
	return xxx_messageInfo_PluginMetaRequest.Size(m)
}
func (m *PluginMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginMetaRequest proto.InternalMessageInfo

func (m *PluginMetaRequest) GetID() *PluginID {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *PluginMetaRequest) GetFlagDeps() bool {
	if m != nil {
		return m.FlagDeps
	}
	return false
}

func (m *PluginMetaRequest) GetFlagDepsStatus() bool {
	if m != nil {
		return m.FlagDepsStatus
	}
	return false
}

func (m *PluginMetaRequest) GetFlagDependent() bool {
	if m != nil {
		return m.FlagDependent
	}
	return false
}

func (m *PluginMetaRequest) GetFlagDependentStatus() bool {
	if m != nil {
		return m.FlagDependentStatus
	}
	return false
}

type PluginMetaReply struct {
	ArrayPluginListWithoutStatus []*PluginListWithoutStatus `protobuf:"bytes,1,rep,name=arrayPluginListWithoutStatus,proto3" json:"arrayPluginListWithoutStatus,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                   `json:"-"`
	XXX_unrecognized             []byte                     `json:"-"`
	XXX_sizecache                int32                      `json:"-"`
}

func (m *PluginMetaReply) Reset()         { *m = PluginMetaReply{} }
func (m *PluginMetaReply) String() string { return proto.CompactTextString(m) }
func (*PluginMetaReply) ProtoMessage()    {}
func (*PluginMetaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{14}
}

func (m *PluginMetaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginMetaReply.Unmarshal(m, b)
}
func (m *PluginMetaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginMetaReply.Marshal(b, m, deterministic)
}
func (m *PluginMetaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginMetaReply.Merge(m, src)
}
func (m *PluginMetaReply) XXX_Size() int {
	return xxx_messageInfo_PluginMetaReply.Size(m)
}
func (m *PluginMetaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginMetaReply.DiscardUnknown(m)
}

var xxx_messageInfo_PluginMetaReply proto.InternalMessageInfo

func (m *PluginMetaReply) GetArrayPluginListWithoutStatus() []*PluginListWithoutStatus {
	if m != nil {
		return m.ArrayPluginListWithoutStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*PluginGetRequest)(nil), "protoNori.PluginGetRequest")
	proto.RegisterType((*PluginPullRequest)(nil), "protoNori.PluginPullRequest")
	proto.RegisterType((*ErrorReply)(nil), "protoNori.ErrorReply")
	proto.RegisterType((*PluginInstallRequest)(nil), "protoNori.PluginInstallRequest")
	proto.RegisterType((*PluginListRequest)(nil), "protoNori.PluginListRequest")
	proto.RegisterType((*PluginListWithStatus)(nil), "protoNori.PluginListWithStatus")
	proto.RegisterType((*PluginListWithoutStatus)(nil), "protoNori.PluginListWithoutStatus")
	proto.RegisterType((*PluginListReply)(nil), "protoNori.PluginListReply")
	proto.RegisterType((*PluginRemoveRequest)(nil), "protoNori.PluginRemoveRequest")
	proto.RegisterType((*PluginStartRequest)(nil), "protoNori.PluginStartRequest")
	proto.RegisterType((*PluginStopRequest)(nil), "protoNori.PluginStopRequest")
	proto.RegisterType((*PluginUninstallRequest)(nil), "protoNori.PluginUninstallRequest")
	proto.RegisterType((*PluginUploadRequest)(nil), "protoNori.PluginUploadRequest")
	proto.RegisterType((*PluginMetaRequest)(nil), "protoNori.PluginMetaRequest")
	proto.RegisterType((*PluginMetaReply)(nil), "protoNori.PluginMetaReply")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 849 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0x12, 0xe7, 0xc7, 0xc7, 0x4d, 0xd3, 0x4e, 0x4b, 0xd7, 0x5a, 0xed, 0x8a, 0xc8, 0x0b,
	0xab, 0x20, 0xa1, 0x0a, 0xbc, 0x20, 0x10, 0x77, 0x15, 0x59, 0x50, 0xa4, 0xee, 0x6a, 0x35, 0x65,
	0xe1, 0x7a, 0x1a, 0x8f, 0xd3, 0xd1, 0xba, 0xb6, 0x99, 0x99, 0x2c, 0xea, 0x15, 0x4f, 0xc0, 0xc3,
	0xf0, 0x08, 0xbc, 0x03, 0xcf, 0xc0, 0x73, 0xa0, 0xf9, 0x8b, 0xc7, 0x4d, 0x5b, 0x05, 0x71, 0x87,
	0xb8, 0xb2, 0xe7, 0x9c, 0xef, 0x7c, 0x3e, 0xe7, 0xcc, 0x37, 0x73, 0x0c, 0xfb, 0xd7, 0x54, 0x08,
	0xb2, 0xa2, 0xe2, 0xb4, 0xe6, 0x95, 0xac, 0x50, 0xa8, 0x1f, 0xaf, 0x2b, 0xce, 0x1e, 0x3f, 0x75,
	0xae, 0x57, 0x54, 0x92, 0x45, 0x99, 0x57, 0xfc, 0x9a, 0x48, 0x56, 0x95, 0x06, 0x99, 0x5c, 0xc0,
	0xc1, 0x9b, 0x62, 0xbd, 0x62, 0xe5, 0xf7, 0x54, 0x62, 0xfa, 0xf3, 0x9a, 0x0a, 0x89, 0x9e, 0x42,
	0x77, 0x91, 0xc5, 0x9d, 0x69, 0x67, 0x16, 0xa5, 0xe3, 0xd3, 0x0d, 0xd5, 0xe9, 0x62, 0x8e, 0xbb,
	0x2c, 0x43, 0x53, 0x88, 0xbe, 0x2b, 0xc8, 0xea, 0x47, 0xca, 0x2f, 0x2b, 0x41, 0xe3, 0xde, 0xb4,
	0x33, 0x1b, 0xe1, 0x28, 0x6f, 0x4c, 0xc9, 0x6b, 0x38, 0x34, 0xa4, 0x6f, 0xd6, 0x45, 0xb1, 0x23,
	0xeb, 0x63, 0x18, 0x29, 0xd6, 0x39, 0xad, 0x45, 0xdc, 0xd5, 0x94, 0xa3, 0xdc, 0xae, 0x93, 0x6f,
	0x00, 0x5e, 0x72, 0x5e, 0x71, 0x4c, 0xeb, 0xe2, 0x06, 0x9d, 0xc0, 0xe0, 0x42, 0x12, 0xb9, 0x16,
	0x9a, 0x6c, 0x84, 0x07, 0x42, 0xaf, 0xd0, 0x31, 0xf4, 0x35, 0x4a, 0x87, 0x87, 0xb8, 0x4f, 0xd5,
	0x22, 0xf9, 0xad, 0x03, 0xc7, 0x26, 0x99, 0x45, 0x29, 0x24, 0xd9, 0x39, 0x9f, 0x5b, 0x55, 0x76,
	0xb7, 0xaa, 0x6c, 0x65, 0xdc, 0x6b, 0x67, 0x8c, 0x62, 0x18, 0x2a, 0xdf, 0x59, 0x51, 0xc4, 0x81,
	0x76, 0x0d, 0x73, 0xb3, 0x4c, 0xfe, 0xe8, 0xb8, 0xe6, 0x9c, 0x33, 0xb1, 0x69, 0xb9, 0x87, 0xef,
	0xb4, 0xf0, 0xe8, 0x09, 0x84, 0xca, 0xd3, 0x54, 0x36, 0xc2, 0x61, 0xee, 0x0c, 0xca, 0x6b, 0xcb,
	0xa2, 0x99, 0x4d, 0x22, 0x64, 0xce, 0xa0, 0x58, 0xf1, 0xba, 0x2c, 0x59, 0xb9, 0x72, 0x59, 0x70,
	0xb3, 0x54, 0xd5, 0xd9, 0x38, 0x72, 0x59, 0xd0, 0xb8, 0x6f, 0xaa, 0x63, 0x8d, 0x49, 0x55, 0xb7,
	0x28, 0xc9, 0x52, 0xb2, 0xf7, 0x34, 0x1e, 0x98, 0xea, 0x98, 0x5d, 0x27, 0xbf, 0xf7, 0x5d, 0x4f,
	0x55, 0x0d, 0x3f, 0x31, 0x79, 0x65, 0x36, 0x04, 0x7d, 0x0c, 0x83, 0x6b, 0x25, 0xb3, 0xf9, 0xdd,
	0x7d, 0xb5, 0x4e, 0xf4, 0x09, 0x0c, 0xce, 0xd6, 0xf2, 0xca, 0x16, 0x14, 0xa5, 0x87, 0x1e, 0xcc,
	0x38, 0xf0, 0x80, 0xe8, 0x27, 0x7a, 0x09, 0x87, 0x73, 0x5a, 0xd3, 0x32, 0xa3, 0xe5, 0x92, 0x51,
	0x71, 0xc6, 0x39, 0xb9, 0xd1, 0x85, 0x46, 0xe9, 0x23, 0x2f, 0xca, 0xc7, 0xe0, 0xc3, 0xec, 0x76,
	0x04, 0xfa, 0x1a, 0xa2, 0x39, 0x15, 0x4b, 0xce, 0x6a, 0xa5, 0x7d, 0xdd, 0x8d, 0x28, 0x3d, 0x69,
	0x11, 0x6c, 0xbc, 0x38, 0xca, 0x9a, 0x05, 0x7a, 0x06, 0xc1, 0xb7, 0x15, 0x37, 0x2d, 0x8a, 0xd2,
	0x89, 0x17, 0xa2, 0xcc, 0x38, 0x58, 0x56, 0x9c, 0xa2, 0x54, 0x6d, 0x83, 0xa4, 0x3c, 0x27, 0x4b,
	0xd3, 0xad, 0x28, 0x3d, 0xf6, 0x4b, 0x77, 0x3e, 0xb5, 0x39, 0xf6, 0x15, 0x7d, 0x0a, 0xc3, 0x73,
	0xb6, 0xa4, 0xa5, 0xa0, 0xf1, 0x50, 0x47, 0x20, 0x2f, 0xc2, 0x7a, 0xf0, 0xb0, 0x30, 0x2f, 0xe8,
	0x39, 0xf4, 0xcf, 0x59, 0xf9, 0x4e, 0xc4, 0x23, 0x8d, 0x3d, 0x68, 0x61, 0xcb, 0x77, 0x02, 0xf7,
	0x0b, 0xf5, 0x40, 0x5f, 0x02, 0x60, 0x5a, 0x57, 0x82, 0xc9, 0x8a, 0xdf, 0xc4, 0xa1, 0x06, 0x7f,
	0xe0, 0x81, 0x1b, 0x27, 0x06, 0xbe, 0x79, 0x47, 0x08, 0x82, 0x1f, 0xc8, 0x4a, 0xc4, 0x30, 0xed,
	0xcd, 0x42, 0x1c, 0x48, 0xb2, 0x6a, 0x69, 0x38, 0x7a, 0x40, 0x93, 0x7b, 0xb7, 0x35, 0xf9, 0x11,
	0x8c, 0x95, 0xb7, 0xd1, 0xe5, 0x58, 0x23, 0xc6, 0xb9, 0x6f, 0x74, 0xe7, 0xcb, 0xe9, 0x73, 0xbf,
	0x39, 0x5f, 0xd6, 0x84, 0x66, 0x30, 0xf1, 0x78, 0xb4, 0x4e, 0x27, 0x1a, 0x35, 0xc9, 0xdb, 0x66,
	0x94, 0xc0, 0x9e, 0x41, 0x5a, 0xbd, 0x1e, 0x68, 0xd8, 0x5e, 0xee, 0xd9, 0x92, 0xbf, 0x7a, 0xf0,
	0xa8, 0xad, 0xd9, 0x6a, 0x2d, 0xff, 0x97, 0xed, 0x7f, 0x4e, 0xb6, 0xc9, 0xaf, 0x30, 0xf1, 0xef,
	0x57, 0x35, 0x31, 0x5e, 0x40, 0x30, 0x27, 0x92, 0xc4, 0x9d, 0x69, 0x6f, 0x16, 0xa5, 0x1f, 0x7a,
	0xbc, 0x77, 0xdd, 0x62, 0x38, 0xc8, 0x88, 0x24, 0xe8, 0x2b, 0x88, 0xb4, 0x9e, 0x0d, 0xc4, 0x6e,
	0xb9, 0x9f, 0x53, 0x33, 0x92, 0x70, 0x44, 0x1b, 0x64, 0xf2, 0x05, 0x1c, 0x99, 0x37, 0x4c, 0xaf,
	0xab, 0xf7, 0x74, 0xb7, 0x79, 0x93, 0xbc, 0x02, 0x64, 0xa2, 0x2e, 0x24, 0xe1, 0xbb, 0x8e, 0x62,
	0xef, 0x88, 0x76, 0xdb, 0x63, 0xe6, 0xdc, 0x4d, 0x99, 0x0b, 0x59, 0xd5, 0xff, 0x9a, 0xed, 0x17,
	0x38, 0x31, 0x6c, 0x6f, 0x4b, 0xf6, 0x8f, 0xa6, 0xe8, 0xbd, 0x94, 0xee, 0x96, 0x70, 0x27, 0x40,
	0xda, 0xe9, 0x35, 0xce, 0x7d, 0x63, 0xf2, 0xb9, 0xeb, 0xe5, 0xdb, 0xba, 0xa8, 0x48, 0xe6, 0xbe,
	0xaa, 0x46, 0x2f, 0x2b, 0x68, 0x4d, 0xe4, 0x95, 0xfe, 0x76, 0x88, 0x47, 0xb9, 0x5d, 0x27, 0x7f,
	0x6e, 0x06, 0xac, 0xfa, 0xe3, 0x71, 0x11, 0xcf, 0xa0, 0xbb, 0x39, 0xde, 0x47, 0x5b, 0x02, 0xd0,
	0xd9, 0xce, 0x1f, 0xfa, 0x07, 0x41, 0xcf, 0x61, 0xdf, 0xf9, 0xec, 0xdf, 0x87, 0x49, 0x78, 0x3f,
	0x6f, 0x59, 0xb7, 0xeb, 0x0a, 0xee, 0xa8, 0x0b, 0x7d, 0x06, 0x47, 0x2d, 0x94, 0xa5, 0x34, 0x73,
	0xf8, 0x28, 0xdf, 0x76, 0x25, 0x37, 0x4e, 0xd6, 0xa6, 0x2a, 0x25, 0xeb, 0x1c, 0x9e, 0x10, 0x75,
	0x4d, 0xdc, 0x73, 0xad, 0x59, 0xb9, 0x27, 0xf7, 0xca, 0x7d, 0x83, 0xc4, 0x0f, 0xf2, 0x5c, 0x0e,
	0x34, 0xc1, 0x8b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xee, 0x83, 0xa7, 0x66, 0x0a, 0x00,
	0x00,
}
