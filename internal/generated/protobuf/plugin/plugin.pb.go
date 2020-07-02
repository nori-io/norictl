// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

package plugin

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

	common "github.com/nori-io/norictl/internal/generated/protobuf/common"
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
	Id                   *common.ID `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagVerbose          bool       `protobuf:"varint,3,opt,name=FlagVerbose,json=flagVerbose,proto3" json:"FlagVerbose,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PluginGetRequest) Reset()         { *m = PluginGetRequest{} }
func (m *PluginGetRequest) String() string { return proto.CompactTextString(m) }
func (*PluginGetRequest) ProtoMessage()    {}
func (*PluginGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0}
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

func (m *PluginGetRequest) GetId() *common.ID {
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
	Id                   *common.ID `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagDeps             bool       `protobuf:"varint,2,opt,name=FlagDeps,json=flagDeps,proto3" json:"FlagDeps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PluginPullRequest) Reset()         { *m = PluginPullRequest{} }
func (m *PluginPullRequest) String() string { return proto.CompactTextString(m) }
func (*PluginPullRequest) ProtoMessage()    {}
func (*PluginPullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{1}
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

func (m *PluginPullRequest) GetId() *common.ID {
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

// cmd: plugin install
type PluginInstallRequest struct {
	Id                   *common.ID `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagVerbose          bool       `protobuf:"varint,2,opt,name=FlagVerbose,json=flagVerbose,proto3" json:"FlagVerbose,omitempty"`
	FlagDeps             bool       `protobuf:"varint,3,opt,name=FlagDeps,json=flagDeps,proto3" json:"FlagDeps,omitempty"`
	FlagAll              bool       `protobuf:"varint,4,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PluginInstallRequest) Reset()         { *m = PluginInstallRequest{} }
func (m *PluginInstallRequest) String() string { return proto.CompactTextString(m) }
func (*PluginInstallRequest) ProtoMessage()    {}
func (*PluginInstallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{2}
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

func (m *PluginInstallRequest) GetId() *common.ID {
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

type PluginInterfaceRequest struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=InterfaceName,json=interfaceName,proto3" json:"InterfaceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInterfaceRequest) Reset()         { *m = PluginInterfaceRequest{} }
func (m *PluginInterfaceRequest) String() string { return proto.CompactTextString(m) }
func (*PluginInterfaceRequest) ProtoMessage()    {}
func (*PluginInterfaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{3}
}

func (m *PluginInterfaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInterfaceRequest.Unmarshal(m, b)
}
func (m *PluginInterfaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInterfaceRequest.Marshal(b, m, deterministic)
}
func (m *PluginInterfaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInterfaceRequest.Merge(m, src)
}
func (m *PluginInterfaceRequest) XXX_Size() int {
	return xxx_messageInfo_PluginInterfaceRequest.Size(m)
}
func (m *PluginInterfaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInterfaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInterfaceRequest proto.InternalMessageInfo

func (m *PluginInterfaceRequest) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

// cmd: plugin list
type PluginListRequest struct {
	FlagAll              bool     `protobuf:"varint,1,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	FlagError            bool     `protobuf:"varint,2,opt,name=FlagError,json=flagError,proto3" json:"FlagError,omitempty"`
	FlagInstalled        bool     `protobuf:"varint,3,opt,name=FlagInstalled,json=flagInstalled,proto3" json:"FlagInstalled,omitempty"`
	FlagRunning          bool     `protobuf:"varint,4,opt,name=FlagRunning,json=flagRunning,proto3" json:"FlagRunning,omitempty"`
	FlagInstallable      bool     `protobuf:"varint,5,opt,name=FlagInstallable,json=flagInstallable,proto3" json:"FlagInstallable,omitempty"`
	FlagInactive         bool     `protobuf:"varint,6,opt,name=FlagInactive,json=flagInactive,proto3" json:"FlagInactive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginListRequest) Reset()         { *m = PluginListRequest{} }
func (m *PluginListRequest) String() string { return proto.CompactTextString(m) }
func (*PluginListRequest) ProtoMessage()    {}
func (*PluginListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{4}
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

func (m *PluginListRequest) GetFlagInstalled() bool {
	if m != nil {
		return m.FlagInstalled
	}
	return false
}

func (m *PluginListRequest) GetFlagRunning() bool {
	if m != nil {
		return m.FlagRunning
	}
	return false
}

func (m *PluginListRequest) GetFlagInstallable() bool {
	if m != nil {
		return m.FlagInstallable
	}
	return false
}

func (m *PluginListRequest) GetFlagInactive() bool {
	if m != nil {
		return m.FlagInactive
	}
	return false
}

type PluginListWithStatus struct {
	MetaID               *common.ID    `protobuf:"bytes,1,opt,name=metaID,proto3" json:"metaID,omitempty"`
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
	return fileDescriptor_22a625af4bc1cc87, []int{5}
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

func (m *PluginListWithStatus) GetMetaID() *common.ID {
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
	MetaID               *common.ID    `protobuf:"bytes,1,opt,name=metaID,proto3" json:"metaID,omitempty"`
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
	return fileDescriptor_22a625af4bc1cc87, []int{6}
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

func (m *PluginListWithoutStatus) GetMetaID() *common.ID {
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
	ErrorPlugin          *common.ErrorReply      `protobuf:"bytes,2,opt,name=ErrorPlugin,json=errorPlugin,proto3" json:"ErrorPlugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PluginListReply) Reset()         { *m = PluginListReply{} }
func (m *PluginListReply) String() string { return proto.CompactTextString(m) }
func (*PluginListReply) ProtoMessage()    {}
func (*PluginListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{7}
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

func (m *PluginListReply) GetErrorPlugin() *common.ErrorReply {
	if m != nil {
		return m.ErrorPlugin
	}
	return nil
}

// cmd: plugin remove
type PluginRemoveRequest struct {
	Id                   *common.ID `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PluginRemoveRequest) Reset()         { *m = PluginRemoveRequest{} }
func (m *PluginRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*PluginRemoveRequest) ProtoMessage()    {}
func (*PluginRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{8}
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

func (m *PluginRemoveRequest) GetId() *common.ID {
	if m != nil {
		return m.Id
	}
	return nil
}

type PluginStartRequest struct {
	Id                   *common.ID `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagAll              bool       `protobuf:"varint,2,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PluginStartRequest) Reset()         { *m = PluginStartRequest{} }
func (m *PluginStartRequest) String() string { return proto.CompactTextString(m) }
func (*PluginStartRequest) ProtoMessage()    {}
func (*PluginStartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{9}
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

func (m *PluginStartRequest) GetId() *common.ID {
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
	Id                   *common.ID `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagAll              bool       `protobuf:"varint,2,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PluginStopRequest) Reset()         { *m = PluginStopRequest{} }
func (m *PluginStopRequest) String() string { return proto.CompactTextString(m) }
func (*PluginStopRequest) ProtoMessage()    {}
func (*PluginStopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{10}
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

func (m *PluginStopRequest) GetId() *common.ID {
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
	Id                   *common.ID `protobuf:"bytes,1,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
	FlagAll              bool       `protobuf:"varint,2,opt,name=FlagAll,json=flagAll,proto3" json:"FlagAll,omitempty"`
	FlagDependent        bool       `protobuf:"varint,3,opt,name=FlagDependent,json=flagDependent,proto3" json:"FlagDependent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PluginUninstallRequest) Reset()         { *m = PluginUninstallRequest{} }
func (m *PluginUninstallRequest) String() string { return proto.CompactTextString(m) }
func (*PluginUninstallRequest) ProtoMessage()    {}
func (*PluginUninstallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{11}
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

func (m *PluginUninstallRequest) GetId() *common.ID {
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
	return fileDescriptor_22a625af4bc1cc87, []int{12}
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
	return fileDescriptor_22a625af4bc1cc87, []int{13}
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
	return fileDescriptor_22a625af4bc1cc87, []int{14}
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
	proto.RegisterType((*PluginGetRequest)(nil), "plugin.PluginGetRequest")
	proto.RegisterType((*PluginPullRequest)(nil), "plugin.PluginPullRequest")
	proto.RegisterType((*PluginInstallRequest)(nil), "plugin.PluginInstallRequest")
	proto.RegisterType((*PluginInterfaceRequest)(nil), "plugin.PluginInterfaceRequest")
	proto.RegisterType((*PluginListRequest)(nil), "plugin.PluginListRequest")
	proto.RegisterType((*PluginListWithStatus)(nil), "plugin.PluginListWithStatus")
	proto.RegisterType((*PluginListWithoutStatus)(nil), "plugin.PluginListWithoutStatus")
	proto.RegisterType((*PluginListReply)(nil), "plugin.PluginListReply")
	proto.RegisterType((*PluginRemoveRequest)(nil), "plugin.PluginRemoveRequest")
	proto.RegisterType((*PluginStartRequest)(nil), "plugin.PluginStartRequest")
	proto.RegisterType((*PluginStopRequest)(nil), "plugin.PluginStopRequest")
	proto.RegisterType((*PluginUninstallRequest)(nil), "plugin.PluginUninstallRequest")
	proto.RegisterType((*PluginUploadRequest)(nil), "plugin.PluginUploadRequest")
	proto.RegisterType((*PluginMetaRequest)(nil), "plugin.PluginMetaRequest")
	proto.RegisterType((*PluginMetaReply)(nil), "plugin.PluginMetaReply")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor_22a625af4bc1cc87) }

var fileDescriptor_22a625af4bc1cc87 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xdb, 0x6e, 0xeb, 0x44,
	0x14, 0x55, 0x12, 0xe7, 0xb6, 0x9d, 0x4b, 0x33, 0xa9, 0xc0, 0x8a, 0x8a, 0x88, 0x0c, 0xaa, 0xc2,
	0x4b, 0x5b, 0x0c, 0xbc, 0x22, 0x15, 0x42, 0x51, 0xa0, 0xa0, 0x6a, 0x4a, 0xe1, 0xb1, 0x9a, 0xda,
	0xe3, 0x74, 0x54, 0xc7, 0x36, 0xf6, 0xa4, 0x52, 0x7e, 0x81, 0x0f, 0xe0, 0xa7, 0x78, 0xe2, 0x13,
	0xce, 0x9f, 0x1c, 0x79, 0x2e, 0xbe, 0xb4, 0x69, 0x8f, 0xa5, 0xf3, 0x74, 0xa4, 0xf3, 0x64, 0xcf,
	0x5e, 0x6b, 0xd6, 0xcc, 0xde, 0xb3, 0xc6, 0xdb, 0x30, 0x88, 0x83, 0xed, 0x9a, 0x85, 0x27, 0x71,
	0x12, 0xf1, 0x08, 0x75, 0xe4, 0x68, 0x36, 0x75, 0xa3, 0xcd, 0x26, 0x0a, 0x4f, 0xe5, 0x43, 0x82,
	0xb3, 0xcf, 0x24, 0x78, 0xbb, 0xa1, 0x9c, 0xdc, 0xb2, 0xd0, 0x8f, 0x92, 0x0d, 0xe1, 0x4c, 0xc3,
	0xf6, 0x15, 0x1c, 0x5c, 0x09, 0xc2, 0xcf, 0x94, 0x63, 0xfa, 0xf7, 0x96, 0xa6, 0x1c, 0xcd, 0xa0,
	0xb9, 0xf2, 0xac, 0xc6, 0xbc, 0xb1, 0x30, 0x1d, 0x38, 0x51, 0x6a, 0xab, 0x25, 0x6e, 0x32, 0x0f,
	0xcd, 0xc1, 0xbc, 0x08, 0xc8, 0xfa, 0x4f, 0x9a, 0xdc, 0x45, 0x29, 0xb5, 0x5a, 0xf3, 0xc6, 0xa2,
	0x87, 0x4d, 0xbf, 0x08, 0xd9, 0xbf, 0xc2, 0x44, 0x2a, 0x5e, 0x6d, 0x83, 0xa0, 0x8e, 0xe4, 0x0c,
	0x7a, 0x99, 0xe4, 0x92, 0xc6, 0xa9, 0xd5, 0x14, 0x7a, 0x3d, 0x5f, 0x8d, 0xed, 0x7f, 0x1a, 0x70,
	0x28, 0xd5, 0x56, 0x61, 0xca, 0x49, 0x3d, 0xc1, 0x27, 0x7b, 0x6c, 0x3e, 0xdb, 0x63, 0x65, 0xc9,
	0x56, 0x75, 0x49, 0x64, 0x41, 0x37, 0xc3, 0xce, 0x83, 0xc0, 0x32, 0x04, 0xd4, 0xf5, 0xe5, 0xd0,
	0xfe, 0x1e, 0x3e, 0xd1, 0x7b, 0xe1, 0x34, 0xf1, 0x89, 0x4b, 0xf5, 0x6e, 0xbe, 0x84, 0x61, 0x1e,
	0xfb, 0x9d, 0x6c, 0xa8, 0xd8, 0x58, 0x1f, 0x0f, 0x59, 0x39, 0x68, 0xbf, 0x69, 0xe8, 0xd2, 0x5c,
	0xb2, 0x34, 0xaf, 0x76, 0x69, 0xbd, 0x46, 0x65, 0x3d, 0x74, 0x04, 0xfd, 0x0c, 0xf9, 0x29, 0x49,
	0xa2, 0x44, 0x65, 0xd1, 0xf7, 0x75, 0x20, 0x5b, 0x33, 0x43, 0x55, 0x5d, 0xa8, 0xa7, 0x12, 0x19,
	0xfa, 0xe5, 0xa0, 0xae, 0x05, 0xde, 0x86, 0x21, 0x0b, 0xd7, 0x2a, 0x23, 0x51, 0x0b, 0x15, 0x42,
	0x0b, 0x18, 0x97, 0x74, 0xc8, 0x5d, 0x40, 0xad, 0xb6, 0x60, 0x8d, 0xfd, 0x6a, 0x18, 0xd9, 0x30,
	0x90, 0x4c, 0xe2, 0x72, 0xf6, 0x48, 0xad, 0x8e, 0xa0, 0x0d, 0xfc, 0x52, 0xcc, 0xfe, 0xb7, 0xad,
	0x0f, 0x2c, 0xcb, 0xf1, 0x2f, 0xc6, 0xef, 0xaf, 0x39, 0xe1, 0xdb, 0x14, 0xd9, 0xd0, 0xc9, 0x2c,
	0xb8, 0x5a, 0xee, 0x39, 0x34, 0x85, 0xa0, 0x63, 0xe8, 0x9c, 0x6f, 0xf9, 0xbd, 0xca, 0xd6, 0x74,
	0x46, 0x27, 0xca, 0xe7, 0x32, 0x8a, 0x3b, 0x44, 0x3c, 0xd1, 0x0f, 0x30, 0x59, 0xd2, 0x98, 0x86,
	0x1e, 0x0d, 0x5d, 0x46, 0xd3, 0xf3, 0x24, 0x21, 0x3b, 0x91, 0xbe, 0xe9, 0x1c, 0xea, 0x29, 0x65,
	0x02, 0x9e, 0x78, 0x4f, 0xe9, 0xe8, 0x3b, 0x30, 0x97, 0x34, 0x75, 0x13, 0x16, 0x67, 0xb7, 0x41,
	0x14, 0xc6, 0x74, 0xa6, 0xc5, 0xec, 0x1c, 0xc2, 0xa6, 0x57, 0x0c, 0xd0, 0x1c, 0x8c, 0x1f, 0xa3,
	0x44, 0x96, 0xc8, 0x74, 0x06, 0x9a, 0x9f, 0xc5, 0xb0, 0xe1, 0x46, 0x09, 0x45, 0xa7, 0xd0, 0xcf,
	0xbd, 0x20, 0x4a, 0x64, 0x3a, 0x13, 0x4d, 0x2b, 0x8c, 0xd3, 0xcf, 0xad, 0x81, 0xbe, 0x82, 0xee,
	0x25, 0x73, 0x69, 0x98, 0x52, 0xab, 0x2b, 0xe8, 0x63, 0x4d, 0x57, 0x61, 0xdc, 0x0d, 0xe4, 0x0b,
	0xfa, 0x02, 0xda, 0x97, 0x2c, 0x7c, 0x48, 0xad, 0x9e, 0x20, 0x0e, 0x0b, 0x62, 0xf8, 0x90, 0xe2,
	0x76, 0x90, 0x3d, 0x90, 0x03, 0x80, 0x69, 0x1c, 0xa5, 0x8c, 0x47, 0xc9, 0xce, 0xea, 0x0b, 0x26,
	0xd2, 0xcc, 0x02, 0xc1, 0x90, 0xe4, 0xef, 0x08, 0x81, 0xf1, 0x07, 0x59, 0xa7, 0x16, 0xcc, 0x5b,
	0x8b, 0x3e, 0x36, 0x38, 0x59, 0x57, 0x2e, 0x82, 0xf9, 0x8a, 0x31, 0x07, 0xef, 0x34, 0xe6, 0xb0,
	0x86, 0x31, 0x47, 0xb5, 0x8c, 0x39, 0xae, 0x67, 0xcc, 0x83, 0x3d, 0xc6, 0xfc, 0xaf, 0x05, 0x9f,
	0x56, 0x8d, 0x19, 0x6d, 0xf9, 0x47, 0x6f, 0x7e, 0x98, 0xde, 0xb4, 0x77, 0x30, 0x2e, 0x7f, 0x49,
	0xe3, 0x60, 0x87, 0xce, 0xc0, 0x58, 0x12, 0x4e, 0xac, 0xc6, 0xbc, 0xb5, 0x30, 0x9d, 0x23, 0x2d,
	0xba, 0xef, 0x63, 0x84, 0x0d, 0x8f, 0x70, 0x82, 0xbe, 0x05, 0x53, 0x38, 0x56, 0x52, 0xd4, 0xb9,
	0x22, 0x7d, 0xf6, 0x02, 0x12, 0xd2, 0xd8, 0xa4, 0x05, 0xcd, 0xfe, 0x1a, 0xa6, 0xf2, 0x0d, 0xd3,
	0x4d, 0xf4, 0x48, 0x6b, 0x34, 0x24, 0xfb, 0x17, 0x40, 0x72, 0xca, 0x35, 0x27, 0x49, 0xad, 0x36,
	0x5b, 0xba, 0x7b, 0xcd, 0x6a, 0x13, 0x5a, 0xe9, 0x1e, 0x72, 0xcd, 0xa3, 0xf8, 0xfd, 0xa4, 0xb8,
	0xee, 0x67, 0x37, 0x21, 0xab, 0xdf, 0x5d, 0x5f, 0xd4, 0xd3, 0x17, 0x5f, 0xbb, 0x9b, 0x97, 0x3b,
	0x52, 0x1e, 0x2c, 0xea, 0x77, 0x13, 0x07, 0x11, 0xf1, 0x8a, 0x25, 0x7b, 0x17, 0x2c, 0xa0, 0x31,
	0xe1, 0xf7, 0xaa, 0x7b, 0xf6, 0x7c, 0x35, 0xb6, 0xff, 0xcf, 0x1b, 0xe7, 0x6f, 0x94, 0x13, 0x3d,
	0x63, 0x0e, 0xcd, 0xfc, 0xc6, 0x1e, 0x54, 0x8f, 0x5b, 0x6c, 0x75, 0xf9, 0xda, 0x9f, 0x05, 0x3a,
	0x86, 0x91, 0xc6, 0xa4, 0x29, 0xd4, 0x6e, 0x47, 0x7e, 0x25, 0xfa, 0x3c, 0x29, 0x63, 0x4f, 0x52,
	0xe8, 0x0c, 0xa6, 0x15, 0x96, 0x92, 0x94, 0x8d, 0x74, 0xea, 0x3f, 0x87, 0xec, 0x47, 0xed, 0x60,
	0x99, 0x52, 0xe6, 0x60, 0x17, 0x8e, 0x48, 0x76, 0xff, 0x5f, 0xf8, 0x4c, 0x29, 0x67, 0x7f, 0xbe,
	0xdf, 0xd9, 0x39, 0x0d, 0xbf, 0x2a, 0x72, 0xd7, 0x11, 0xff, 0x7d, 0xdf, 0xbc, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x36, 0xd2, 0x14, 0xee, 0x43, 0x0a, 0x00, 0x00,
}