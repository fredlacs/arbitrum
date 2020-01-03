// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rollup.proto

package rollup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	value "github.com/offchainlabs/arbitrum/packages/arb-util/value"
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

type ChainObserverBuf struct {
	ContractAddress      []byte           `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	VmParams             *ChainParamsBuf  `protobuf:"bytes,2,opt,name=vmParams,proto3" json:"vmParams,omitempty"`
	PendingInbox         *PendingInboxBuf `protobuf:"bytes,3,opt,name=pendingInbox,proto3" json:"pendingInbox,omitempty"`
	Nodes                []*NodeBuf       `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	LatestConfirmedHash  *value.HashBuf   `protobuf:"bytes,5,opt,name=latestConfirmedHash,proto3" json:"latestConfirmedHash,omitempty"`
	LeafHashes           []*value.HashBuf `protobuf:"bytes,6,rep,name=leafHashes,proto3" json:"leafHashes,omitempty"`
	Stakers              []*StakerBuf     `protobuf:"bytes,7,rep,name=stakers,proto3" json:"stakers,omitempty"`
	Challenges           []*ChallengeBuf  `protobuf:"bytes,8,rep,name=challenges,proto3" json:"challenges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChainObserverBuf) Reset()         { *m = ChainObserverBuf{} }
func (m *ChainObserverBuf) String() string { return proto.CompactTextString(m) }
func (*ChainObserverBuf) ProtoMessage()    {}
func (*ChainObserverBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{0}
}

func (m *ChainObserverBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainObserverBuf.Unmarshal(m, b)
}
func (m *ChainObserverBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainObserverBuf.Marshal(b, m, deterministic)
}
func (m *ChainObserverBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainObserverBuf.Merge(m, src)
}
func (m *ChainObserverBuf) XXX_Size() int {
	return xxx_messageInfo_ChainObserverBuf.Size(m)
}
func (m *ChainObserverBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainObserverBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChainObserverBuf proto.InternalMessageInfo

func (m *ChainObserverBuf) GetContractAddress() []byte {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

func (m *ChainObserverBuf) GetVmParams() *ChainParamsBuf {
	if m != nil {
		return m.VmParams
	}
	return nil
}

func (m *ChainObserverBuf) GetPendingInbox() *PendingInboxBuf {
	if m != nil {
		return m.PendingInbox
	}
	return nil
}

func (m *ChainObserverBuf) GetNodes() []*NodeBuf {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *ChainObserverBuf) GetLatestConfirmedHash() *value.HashBuf {
	if m != nil {
		return m.LatestConfirmedHash
	}
	return nil
}

func (m *ChainObserverBuf) GetLeafHashes() []*value.HashBuf {
	if m != nil {
		return m.LeafHashes
	}
	return nil
}

func (m *ChainObserverBuf) GetStakers() []*StakerBuf {
	if m != nil {
		return m.Stakers
	}
	return nil
}

func (m *ChainObserverBuf) GetChallenges() []*ChallengeBuf {
	if m != nil {
		return m.Challenges
	}
	return nil
}

type ChainParamsBuf struct {
	StakeRequirement     *value.BigIntegerBuf `protobuf:"bytes,1,opt,name=stakeRequirement,proto3" json:"stakeRequirement,omitempty"`
	GracePeriod          *RollupTimeBuf       `protobuf:"bytes,2,opt,name=gracePeriod,proto3" json:"gracePeriod,omitempty"`
	MaxExecutionSteps    uint32               `protobuf:"varint,3,opt,name=maxExecutionSteps,proto3" json:"maxExecutionSteps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChainParamsBuf) Reset()         { *m = ChainParamsBuf{} }
func (m *ChainParamsBuf) String() string { return proto.CompactTextString(m) }
func (*ChainParamsBuf) ProtoMessage()    {}
func (*ChainParamsBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{1}
}

func (m *ChainParamsBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainParamsBuf.Unmarshal(m, b)
}
func (m *ChainParamsBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainParamsBuf.Marshal(b, m, deterministic)
}
func (m *ChainParamsBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainParamsBuf.Merge(m, src)
}
func (m *ChainParamsBuf) XXX_Size() int {
	return xxx_messageInfo_ChainParamsBuf.Size(m)
}
func (m *ChainParamsBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainParamsBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChainParamsBuf proto.InternalMessageInfo

func (m *ChainParamsBuf) GetStakeRequirement() *value.BigIntegerBuf {
	if m != nil {
		return m.StakeRequirement
	}
	return nil
}

func (m *ChainParamsBuf) GetGracePeriod() *RollupTimeBuf {
	if m != nil {
		return m.GracePeriod
	}
	return nil
}

func (m *ChainParamsBuf) GetMaxExecutionSteps() uint32 {
	if m != nil {
		return m.MaxExecutionSteps
	}
	return 0
}

type NodeBuf struct {
	Hash                 *value.HashBuf     `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	DisputableNode       *DisputableNodeBuf `protobuf:"bytes,2,opt,name=disputableNode,proto3" json:"disputableNode,omitempty"`
	MachineHash          *value.HashBuf     `protobuf:"bytes,3,opt,name=machineHash,proto3" json:"machineHash,omitempty"`
	PendingTopHash       *value.HashBuf     `protobuf:"bytes,4,opt,name=pendingTopHash,proto3" json:"pendingTopHash,omitempty"`
	LinkType             uint32             `protobuf:"varint,5,opt,name=linkType,proto3" json:"linkType,omitempty"`
	PrevHash             *value.HashBuf     `protobuf:"bytes,6,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NodeBuf) Reset()         { *m = NodeBuf{} }
func (m *NodeBuf) String() string { return proto.CompactTextString(m) }
func (*NodeBuf) ProtoMessage()    {}
func (*NodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{2}
}

func (m *NodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeBuf.Unmarshal(m, b)
}
func (m *NodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeBuf.Marshal(b, m, deterministic)
}
func (m *NodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeBuf.Merge(m, src)
}
func (m *NodeBuf) XXX_Size() int {
	return xxx_messageInfo_NodeBuf.Size(m)
}
func (m *NodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_NodeBuf proto.InternalMessageInfo

func (m *NodeBuf) GetHash() *value.HashBuf {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *NodeBuf) GetDisputableNode() *DisputableNodeBuf {
	if m != nil {
		return m.DisputableNode
	}
	return nil
}

func (m *NodeBuf) GetMachineHash() *value.HashBuf {
	if m != nil {
		return m.MachineHash
	}
	return nil
}

func (m *NodeBuf) GetPendingTopHash() *value.HashBuf {
	if m != nil {
		return m.PendingTopHash
	}
	return nil
}

func (m *NodeBuf) GetLinkType() uint32 {
	if m != nil {
		return m.LinkType
	}
	return 0
}

func (m *NodeBuf) GetPrevHash() *value.HashBuf {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

type StakerBuf struct {
	Address              []byte         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location             *value.HashBuf `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	CreationTime         *RollupTimeBuf `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	InChallenge          bool           `protobuf:"varint,4,opt,name=inChallenge,proto3" json:"inChallenge,omitempty"`
	ChallengeAddr        []byte         `protobuf:"bytes,5,opt,name=challengeAddr,proto3" json:"challengeAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StakerBuf) Reset()         { *m = StakerBuf{} }
func (m *StakerBuf) String() string { return proto.CompactTextString(m) }
func (*StakerBuf) ProtoMessage()    {}
func (*StakerBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{3}
}

func (m *StakerBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakerBuf.Unmarshal(m, b)
}
func (m *StakerBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakerBuf.Marshal(b, m, deterministic)
}
func (m *StakerBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakerBuf.Merge(m, src)
}
func (m *StakerBuf) XXX_Size() int {
	return xxx_messageInfo_StakerBuf.Size(m)
}
func (m *StakerBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_StakerBuf.DiscardUnknown(m)
}

var xxx_messageInfo_StakerBuf proto.InternalMessageInfo

func (m *StakerBuf) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *StakerBuf) GetLocation() *value.HashBuf {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *StakerBuf) GetCreationTime() *RollupTimeBuf {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *StakerBuf) GetInChallenge() bool {
	if m != nil {
		return m.InChallenge
	}
	return false
}

func (m *StakerBuf) GetChallengeAddr() []byte {
	if m != nil {
		return m.ChallengeAddr
	}
	return nil
}

type ChallengeBuf struct {
	Contract             []byte   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Asserter             []byte   `protobuf:"bytes,2,opt,name=asserter,proto3" json:"asserter,omitempty"`
	Challenger           []byte   `protobuf:"bytes,3,opt,name=challenger,proto3" json:"challenger,omitempty"`
	Kind                 uint32   `protobuf:"varint,4,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChallengeBuf) Reset()         { *m = ChallengeBuf{} }
func (m *ChallengeBuf) String() string { return proto.CompactTextString(m) }
func (*ChallengeBuf) ProtoMessage()    {}
func (*ChallengeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{4}
}

func (m *ChallengeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeBuf.Unmarshal(m, b)
}
func (m *ChallengeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeBuf.Marshal(b, m, deterministic)
}
func (m *ChallengeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeBuf.Merge(m, src)
}
func (m *ChallengeBuf) XXX_Size() int {
	return xxx_messageInfo_ChallengeBuf.Size(m)
}
func (m *ChallengeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeBuf proto.InternalMessageInfo

func (m *ChallengeBuf) GetContract() []byte {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *ChallengeBuf) GetAsserter() []byte {
	if m != nil {
		return m.Asserter
	}
	return nil
}

func (m *ChallengeBuf) GetChallenger() []byte {
	if m != nil {
		return m.Challenger
	}
	return nil
}

func (m *ChallengeBuf) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type DisputableNodeBuf struct {
	PrevNodeHash          *value.HashBuf          `protobuf:"bytes,1,opt,name=prevNodeHash,proto3" json:"prevNodeHash,omitempty"`
	TimeLowerBound        *RollupTimeBuf          `protobuf:"bytes,2,opt,name=timeLowerBound,proto3" json:"timeLowerBound,omitempty"`
	TimeUpperBound        *RollupTimeBuf          `protobuf:"bytes,3,opt,name=timeUpperBound,proto3" json:"timeUpperBound,omitempty"`
	AfterPendingTop       *value.HashBuf          `protobuf:"bytes,4,opt,name=afterPendingTop,proto3" json:"afterPendingTop,omitempty"`
	ImportedMessagesSlice *value.HashBuf          `protobuf:"bytes,5,opt,name=importedMessagesSlice,proto3" json:"importedMessagesSlice,omitempty"`
	ImportedMessageCount  *value.BigIntegerBuf    `protobuf:"bytes,6,opt,name=importedMessageCount,proto3" json:"importedMessageCount,omitempty"`
	AssertionStub         *protocol.AssertionStub `protobuf:"bytes,7,opt,name=assertionStub,proto3" json:"assertionStub,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      []byte                  `json:"-"`
	XXX_sizecache         int32                   `json:"-"`
}

func (m *DisputableNodeBuf) Reset()         { *m = DisputableNodeBuf{} }
func (m *DisputableNodeBuf) String() string { return proto.CompactTextString(m) }
func (*DisputableNodeBuf) ProtoMessage()    {}
func (*DisputableNodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{5}
}

func (m *DisputableNodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisputableNodeBuf.Unmarshal(m, b)
}
func (m *DisputableNodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisputableNodeBuf.Marshal(b, m, deterministic)
}
func (m *DisputableNodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputableNodeBuf.Merge(m, src)
}
func (m *DisputableNodeBuf) XXX_Size() int {
	return xxx_messageInfo_DisputableNodeBuf.Size(m)
}
func (m *DisputableNodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputableNodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_DisputableNodeBuf proto.InternalMessageInfo

func (m *DisputableNodeBuf) GetPrevNodeHash() *value.HashBuf {
	if m != nil {
		return m.PrevNodeHash
	}
	return nil
}

func (m *DisputableNodeBuf) GetTimeLowerBound() *RollupTimeBuf {
	if m != nil {
		return m.TimeLowerBound
	}
	return nil
}

func (m *DisputableNodeBuf) GetTimeUpperBound() *RollupTimeBuf {
	if m != nil {
		return m.TimeUpperBound
	}
	return nil
}

func (m *DisputableNodeBuf) GetAfterPendingTop() *value.HashBuf {
	if m != nil {
		return m.AfterPendingTop
	}
	return nil
}

func (m *DisputableNodeBuf) GetImportedMessagesSlice() *value.HashBuf {
	if m != nil {
		return m.ImportedMessagesSlice
	}
	return nil
}

func (m *DisputableNodeBuf) GetImportedMessageCount() *value.BigIntegerBuf {
	if m != nil {
		return m.ImportedMessageCount
	}
	return nil
}

func (m *DisputableNodeBuf) GetAssertionStub() *protocol.AssertionStub {
	if m != nil {
		return m.AssertionStub
	}
	return nil
}

type PendingInboxBuf struct {
	Items                [][]byte       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	HashOfRest           *value.HashBuf `protobuf:"bytes,2,opt,name=hashOfRest,proto3" json:"hashOfRest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PendingInboxBuf) Reset()         { *m = PendingInboxBuf{} }
func (m *PendingInboxBuf) String() string { return proto.CompactTextString(m) }
func (*PendingInboxBuf) ProtoMessage()    {}
func (*PendingInboxBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{6}
}

func (m *PendingInboxBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingInboxBuf.Unmarshal(m, b)
}
func (m *PendingInboxBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingInboxBuf.Marshal(b, m, deterministic)
}
func (m *PendingInboxBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingInboxBuf.Merge(m, src)
}
func (m *PendingInboxBuf) XXX_Size() int {
	return xxx_messageInfo_PendingInboxBuf.Size(m)
}
func (m *PendingInboxBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingInboxBuf.DiscardUnknown(m)
}

var xxx_messageInfo_PendingInboxBuf proto.InternalMessageInfo

func (m *PendingInboxBuf) GetItems() [][]byte {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *PendingInboxBuf) GetHashOfRest() *value.HashBuf {
	if m != nil {
		return m.HashOfRest
	}
	return nil
}

type RollupTimeBuf struct {
	Val                  *value.BigIntegerBuf `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RollupTimeBuf) Reset()         { *m = RollupTimeBuf{} }
func (m *RollupTimeBuf) String() string { return proto.CompactTextString(m) }
func (*RollupTimeBuf) ProtoMessage()    {}
func (*RollupTimeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{7}
}

func (m *RollupTimeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollupTimeBuf.Unmarshal(m, b)
}
func (m *RollupTimeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollupTimeBuf.Marshal(b, m, deterministic)
}
func (m *RollupTimeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollupTimeBuf.Merge(m, src)
}
func (m *RollupTimeBuf) XXX_Size() int {
	return xxx_messageInfo_RollupTimeBuf.Size(m)
}
func (m *RollupTimeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_RollupTimeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_RollupTimeBuf proto.InternalMessageInfo

func (m *RollupTimeBuf) GetVal() *value.BigIntegerBuf {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*ChainObserverBuf)(nil), "rollup.ChainObserverBuf")
	proto.RegisterType((*ChainParamsBuf)(nil), "rollup.ChainParamsBuf")
	proto.RegisterType((*NodeBuf)(nil), "rollup.NodeBuf")
	proto.RegisterType((*StakerBuf)(nil), "rollup.StakerBuf")
	proto.RegisterType((*ChallengeBuf)(nil), "rollup.ChallengeBuf")
	proto.RegisterType((*DisputableNodeBuf)(nil), "rollup.DisputableNodeBuf")
	proto.RegisterType((*PendingInboxBuf)(nil), "rollup.PendingInboxBuf")
	proto.RegisterType((*RollupTimeBuf)(nil), "rollup.RollupTimeBuf")
}

func init() { proto.RegisterFile("rollup.proto", fileDescriptor_037f188b50610c79) }

var fileDescriptor_037f188b50610c79 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x57, 0x2e, 0x6d, 0x13, 0x26, 0x4e, 0x7a, 0x1d, 0x7a, 0x9c, 0xe9, 0x03, 0xaa, 0x2c, 0x40,
	0x11, 0x7f, 0x12, 0x54, 0x10, 0x07, 0x42, 0x87, 0xda, 0xf4, 0x90, 0x7a, 0x12, 0x70, 0x95, 0x5b,
	0x84, 0xc4, 0xdb, 0xc6, 0x9e, 0x24, 0xab, 0xda, 0xbb, 0x66, 0x77, 0x1d, 0x0a, 0x1f, 0x87, 0x37,
	0x3e, 0x02, 0x5f, 0x82, 0x27, 0x3e, 0x10, 0xda, 0xf5, 0x9f, 0xb3, 0xd3, 0x84, 0x7b, 0x49, 0x76,
	0xf6, 0xf7, 0xfb, 0xcd, 0xce, 0xce, 0xcc, 0x8e, 0xc1, 0x53, 0x32, 0x49, 0xf2, 0x6c, 0x92, 0x29,
	0x69, 0x24, 0x1e, 0x14, 0xd6, 0xc9, 0xd1, 0x9a, 0x25, 0x39, 0x4d, 0xdd, 0x6f, 0x01, 0x9d, 0x3c,
	0x75, 0x7f, 0x91, 0x4c, 0xa6, 0xd5, 0xa2, 0x00, 0x82, 0xbf, 0xba, 0xf0, 0xf8, 0x72, 0xc5, 0xb8,
	0x78, 0x35, 0xd7, 0xa4, 0xd6, 0xa4, 0x66, 0xf9, 0x02, 0xc7, 0x70, 0x18, 0x49, 0x61, 0x14, 0x8b,
	0xcc, 0x45, 0x1c, 0x2b, 0xd2, 0xda, 0xef, 0x9c, 0x76, 0xc6, 0x5e, 0xb8, 0xb9, 0x8d, 0x67, 0xd0,
	0x5f, 0xa7, 0xd7, 0x4c, 0xb1, 0x54, 0xfb, 0x8f, 0x4e, 0x3b, 0xe3, 0xc1, 0xd9, 0x3b, 0x93, 0x32,
	0x26, 0xe7, 0xb5, 0x80, 0x66, 0xf9, 0x22, 0xac, 0x79, 0xf8, 0x0d, 0x78, 0x19, 0x89, 0x98, 0x8b,
	0xe5, 0x4b, 0x31, 0x97, 0xf7, 0x7e, 0xd7, 0xe9, 0x9e, 0x56, 0xba, 0xeb, 0x06, 0x66, 0x85, 0x2d,
	0x32, 0x7e, 0x00, 0xfb, 0x42, 0xc6, 0xa4, 0xfd, 0xbd, 0xd3, 0xee, 0x78, 0x70, 0x76, 0x58, 0xa9,
	0x7e, 0x94, 0x31, 0x59, 0x76, 0x81, 0xe2, 0x39, 0xbc, 0x9d, 0x30, 0x43, 0xda, 0x5c, 0x4a, 0xb1,
	0xe0, 0x2a, 0xa5, 0xf8, 0x8a, 0xe9, 0x95, 0xbf, 0xef, 0x8e, 0x1a, 0x4d, 0x8a, 0xd4, 0xd8, 0x2d,
	0xab, 0xd9, 0x46, 0xc5, 0x09, 0x40, 0x42, 0x6c, 0x61, 0xd7, 0xa4, 0xfd, 0x03, 0x77, 0xda, 0xa6,
	0xb0, 0xc1, 0xc0, 0x8f, 0xa1, 0xa7, 0x0d, 0xbb, 0x23, 0xa5, 0xfd, 0x9e, 0x23, 0x1f, 0x55, 0xa1,
	0xdd, 0xb8, 0x6d, 0xcb, 0xaf, 0x18, 0xf8, 0x05, 0x40, 0xb4, 0x62, 0x49, 0x42, 0x62, 0x49, 0xda,
	0xef, 0x3b, 0xfe, 0x71, 0x23, 0x71, 0x05, 0xe2, 0x8e, 0x78, 0xcd, 0x0b, 0xfe, 0xee, 0xc0, 0xa8,
	0x9d, 0x55, 0x3c, 0x87, 0xc7, 0xce, 0x67, 0x48, 0xbf, 0xe6, 0x5c, 0x51, 0x4a, 0xc2, 0xb8, 0x52,
	0x59, 0x77, 0x45, 0xac, 0x33, 0xbe, 0x7c, 0x29, 0x0c, 0x2d, 0x8b, 0x08, 0x1e, 0xb0, 0xf1, 0x19,
	0x0c, 0x96, 0x8a, 0x45, 0x74, 0x4d, 0x8a, 0xcb, 0xb8, 0x2c, 0xe2, 0x93, 0x2a, 0x96, 0xd0, 0xfd,
	0xdd, 0xf2, 0xd4, 0x05, 0xd3, 0x64, 0xe2, 0x27, 0x70, 0x94, 0xb2, 0xfb, 0xef, 0xee, 0x29, 0xca,
	0x0d, 0x97, 0xe2, 0xc6, 0x50, 0xa6, 0x5d, 0x2d, 0x87, 0xe1, 0x43, 0x20, 0xf8, 0xf3, 0x11, 0xf4,
	0xca, 0x1a, 0x61, 0x00, 0x7b, 0x2b, 0x5b, 0x8d, 0xce, 0xd6, 0x6a, 0x38, 0x0c, 0x2f, 0x60, 0x14,
	0x73, 0x9d, 0xe5, 0x86, 0xcd, 0x13, 0xb2, 0xc2, 0x32, 0xb2, 0x77, 0xab, 0xc8, 0x5e, 0xb4, 0x50,
	0x2b, 0xdc, 0x10, 0xe0, 0x67, 0x30, 0x48, 0x59, 0xb4, 0xe2, 0x82, 0x5c, 0xed, 0xbb, 0x5b, 0x4f,
	0x6b, 0x52, 0xf0, 0x4b, 0x18, 0x95, 0xcd, 0x76, 0x2b, 0x33, 0x27, 0xda, 0xdb, 0x2a, 0xda, 0x60,
	0xe1, 0x09, 0xf4, 0x13, 0x2e, 0xee, 0x6e, 0x7f, 0xcf, 0xc8, 0xb5, 0xd8, 0x30, 0xac, 0x6d, 0xfc,
	0x08, 0xfa, 0x99, 0xa2, 0xb5, 0xf3, 0x76, 0xb0, 0xd5, 0x5b, 0x8d, 0x07, 0xff, 0x76, 0xe0, 0xad,
	0xba, 0x5b, 0xd0, 0x87, 0x1e, 0x6b, 0xbd, 0xbe, 0xca, 0xb4, 0x3e, 0x13, 0x19, 0x31, 0x9b, 0xdd,
	0x32, 0x2d, 0x0f, 0x7c, 0x56, 0x38, 0x7e, 0x0d, 0x5e, 0xa4, 0xc8, 0xad, 0x6d, 0x19, 0xcb, 0x34,
	0xec, 0x28, 0x70, 0x8b, 0x8a, 0xa7, 0x30, 0xe0, 0xa2, 0xee, 0x46, 0x97, 0x8b, 0x7e, 0xd8, 0xdc,
	0xc2, 0xf7, 0x61, 0x58, 0xf7, 0xa7, 0x1d, 0x09, 0xee, 0xf6, 0x5e, 0xd8, 0xde, 0x0c, 0xfe, 0x00,
	0xaf, 0xd9, 0xd3, 0x36, 0x5d, 0xd5, 0x1c, 0x29, 0x6f, 0x56, 0xdb, 0x16, 0x63, 0x5a, 0x93, 0x32,
	0xa4, 0xdc, 0xd5, 0xbc, 0xb0, 0xb6, 0xf1, 0xbd, 0xc6, 0xab, 0x51, 0xee, 0x22, 0x5e, 0xe3, 0x7d,
	0x28, 0x44, 0xd8, 0xbb, 0xe3, 0x22, 0x76, 0x81, 0x0e, 0x43, 0xb7, 0x0e, 0xfe, 0xe9, 0xc2, 0xd1,
	0x83, 0x56, 0xc1, 0x33, 0xf0, 0x6c, 0xd2, 0xad, 0x79, 0xb5, 0xbb, 0x13, 0x5b, 0x1c, 0x7c, 0x0e,
	0x23, 0xc3, 0x53, 0xfa, 0x5e, 0xfe, 0x46, 0x6a, 0x26, 0x73, 0xf1, 0x86, 0xb7, 0xb2, 0x41, 0xae,
	0xe4, 0x3f, 0x65, 0x59, 0x25, 0xef, 0xbe, 0x51, 0xfe, 0x9a, 0x8c, 0x5f, 0xc1, 0x21, 0x5b, 0x18,
	0x52, 0xd7, 0x75, 0xe7, 0xed, 0xe8, 0xcd, 0x4d, 0x1a, 0xbe, 0x80, 0x27, 0x3c, 0xcd, 0xa4, 0x32,
	0x14, 0xff, 0x40, 0x5a, 0xb3, 0x25, 0xe9, 0x9b, 0x84, 0x47, 0xb4, 0x63, 0x18, 0x6e, 0x27, 0xe3,
	0x15, 0x1c, 0x6f, 0x00, 0x97, 0x32, 0x17, 0xa6, 0x6c, 0xe9, 0xed, 0xc3, 0x66, 0xab, 0x02, 0x9f,
	0xc3, 0xb0, 0xa8, 0xa8, 0x9b, 0x0d, 0xf9, 0xdc, 0xef, 0x95, 0xf3, 0xbf, 0xfe, 0x32, 0x5d, 0x34,
	0xe1, 0xb0, 0xcd, 0x0e, 0x7e, 0x86, 0xc3, 0x8d, 0x2f, 0x04, 0x1e, 0xc3, 0x3e, 0x37, 0x94, 0xda,
	0x67, 0xd2, 0x1d, 0x7b, 0x61, 0x61, 0xd8, 0x01, 0x6e, 0x27, 0xc9, 0xab, 0x45, 0x48, 0xda, 0xec,
	0x78, 0x26, 0x0d, 0x46, 0xf0, 0x0c, 0x86, 0xad, 0x12, 0xe0, 0x87, 0xd0, 0x5d, 0xb3, 0xe4, 0x7f,
	0xc7, 0xa9, 0x25, 0xcc, 0xce, 0x7f, 0xf9, 0x76, 0xc9, 0xcd, 0x2a, 0x9f, 0x4f, 0x22, 0x99, 0x4e,
	0xe5, 0x62, 0x11, 0xd9, 0x19, 0x9d, 0xb0, 0xb9, 0x9e, 0x32, 0x35, 0xe7, 0x46, 0xe5, 0xe9, 0x34,
	0x63, 0xd1, 0x9d, 0x4d, 0xa6, 0xdd, 0xf9, 0x74, 0xcd, 0x12, 0x1e, 0x33, 0x23, 0xd5, 0xb4, 0x28,
	0xfd, 0xfc, 0xc0, 0x5d, 0xfd, 0xf3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x28, 0x9c, 0x6c, 0x1e,
	0xcf, 0x07, 0x00, 0x00,
}