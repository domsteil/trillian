// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trillian.proto

package trillian

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import sigpb "github.com/google/trillian/crypto/sigpb"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Defines the way empty / node / leaf hashes are constructed incorporating
// preimage protection, which can be application specific.
type HashStrategy int32

const (
	// Hash strategy cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	HashStrategy_UNKNOWN_HASH_STRATEGY HashStrategy = 0
	// Certificate Transparency strategy: leaf hash prefix = 0x00, node prefix =
	// 0x01, empty hash is digest([]byte{}), as defined in the specification.
	HashStrategy_RFC6962_SHA256 HashStrategy = 1
	// Sparse Merkle Tree strategy:  leaf hash prefix = 0x00, node prefix = 0x01,
	// empty branch is recursively computed from empty leaf nodes.
	// NOT secure in a multi tree environment. For testing only.
	HashStrategy_TEST_MAP_HASHER HashStrategy = 2
	// Append-only log strategy where leaf nodes are defined as the ObjectHash.
	// All other properties are equal to RFC6962_SHA256.
	HashStrategy_OBJECT_RFC6962_SHA256 HashStrategy = 3
	// The CONIKS sparse tree hasher with SHA512_256 as the hash algorithm.
	HashStrategy_CONIKS_SHA512_256 HashStrategy = 4
)

var HashStrategy_name = map[int32]string{
	0: "UNKNOWN_HASH_STRATEGY",
	1: "RFC6962_SHA256",
	2: "TEST_MAP_HASHER",
	3: "OBJECT_RFC6962_SHA256",
	4: "CONIKS_SHA512_256",
}
var HashStrategy_value = map[string]int32{
	"UNKNOWN_HASH_STRATEGY": 0,
	"RFC6962_SHA256":        1,
	"TEST_MAP_HASHER":       2,
	"OBJECT_RFC6962_SHA256": 3,
	"CONIKS_SHA512_256":     4,
}

func (x HashStrategy) String() string {
	return proto.EnumName(HashStrategy_name, int32(x))
}
func (HashStrategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// State of the tree.
type TreeState int32

const (
	// Tree state cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeState_UNKNOWN_TREE_STATE TreeState = 0
	// Active trees are able to respond to both read and write requests.
	TreeState_ACTIVE TreeState = 1
	// Frozen trees are only able to respond to read requests, writing to a frozen
	// tree is forbidden.
	TreeState_FROZEN TreeState = 2
	// Tree was been deleted, therefore is invisible and acts similarly to a
	// non-existing tree for all requests.
	// A soft deleted tree may be undeleted while the soft-deletion period has not
	// passed.
	TreeState_SOFT_DELETED TreeState = 3
	// A hard deleted tree was been definitely deleted and cannot be recovered.
	// Acts an a non-existing tree for all read and write requests, but blocks the
	// tree ID from ever being reused.
	TreeState_HARD_DELETED TreeState = 4
)

var TreeState_name = map[int32]string{
	0: "UNKNOWN_TREE_STATE",
	1: "ACTIVE",
	2: "FROZEN",
	3: "SOFT_DELETED",
	4: "HARD_DELETED",
}
var TreeState_value = map[string]int32{
	"UNKNOWN_TREE_STATE": 0,
	"ACTIVE":             1,
	"FROZEN":             2,
	"SOFT_DELETED":       3,
	"HARD_DELETED":       4,
}

func (x TreeState) String() string {
	return proto.EnumName(TreeState_name, int32(x))
}
func (TreeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

// Type of the tree.
type TreeType int32

const (
	// Tree type cannot be determined. Included to enable detection of
	// mismatched proto versions being used. Represents an invalid value.
	TreeType_UNKNOWN_TREE_TYPE TreeType = 0
	// Tree represents a verifiable log.
	TreeType_LOG TreeType = 1
	// Tree represents a verifiable map.
	TreeType_MAP TreeType = 2
)

var TreeType_name = map[int32]string{
	0: "UNKNOWN_TREE_TYPE",
	1: "LOG",
	2: "MAP",
}
var TreeType_value = map[string]int32{
	"UNKNOWN_TREE_TYPE": 0,
	"LOG":               1,
	"MAP":               2,
}

func (x TreeType) String() string {
	return proto.EnumName(TreeType_name, int32(x))
}
func (TreeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

// Represents a tree, which may be either a verifiable log or map.
// Readonly attributes are assigned at tree creation, after which they may not
// be modified.
type Tree struct {
	// ID of the tree.
	// Readonly.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId" json:"tree_id,omitempty"`
	// State of the tree.
	// Trees are active after creation. At any point the tree may transition
	// between ACTIVE and FROZEN.
	// Deleted trees are set as SOFT_DELETED for a certain time period, after
	// which they'll automatically transition to HARD_DELETED.
	TreeState TreeState `protobuf:"varint,2,opt,name=tree_state,json=treeState,enum=trillian.TreeState" json:"tree_state,omitempty"`
	// Type of the tree.
	// Readonly.
	TreeType TreeType `protobuf:"varint,3,opt,name=tree_type,json=treeType,enum=trillian.TreeType" json:"tree_type,omitempty"`
	// Hash strategy to be used by the tree.
	// Readonly.
	HashStrategy HashStrategy `protobuf:"varint,4,opt,name=hash_strategy,json=hashStrategy,enum=trillian.HashStrategy" json:"hash_strategy,omitempty"`
	// Hash algorithm to be used by the tree.
	// Readonly.
	// TODO(gdbelvin): Deprecate in favor of signature_cipher_suite and hash_strategy.
	HashAlgorithm sigpb.DigitallySigned_HashAlgorithm `protobuf:"varint,5,opt,name=hash_algorithm,json=hashAlgorithm,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	// Signature algorithm to be used by the tree.
	// Readonly.
	// TODO(gdbelvin): Deprecate in favor of signature_cipher_suite.
	SignatureAlgorithm sigpb.DigitallySigned_SignatureAlgorithm `protobuf:"varint,6,opt,name=signature_algorithm,json=signatureAlgorithm,enum=sigpb.DigitallySigned_SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// Signature cipher suite specifies the algorithms used to generate signatures.
	SignatureCipherSuite sigpb.DigitallySigned_SignatureCipherSuite `protobuf:"varint,18,opt,name=signature_cipher_suite,json=signatureCipherSuite,enum=sigpb.DigitallySigned_SignatureCipherSuite" json:"signature_cipher_suite,omitempty"`
	// Display name of the tree.
	// Optional.
	DisplayName string `protobuf:"bytes,8,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Description of the tree,
	// Optional.
	Description string `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	// Identifies the private key used for signing tree heads and entry
	// timestamps.
	// This can be any type of message to accommodate different key management
	// systems, e.g. PEM files, HSMs, etc.
	// Private keys are write-only: they're never returned by RPCs.
	// TODO(RJPercival): Implement sufficient validation to allow this field to be
	// mutable. It should be mutable in the sense that the key can be migrated to
	// a different key management system, but the key itself should never change.
	PrivateKey *google_protobuf.Any `protobuf:"bytes,12,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// Storage-specific settings.
	// Varies according to the storage implementation backing Trillian.
	StorageSettings *google_protobuf.Any `protobuf:"bytes,13,opt,name=storage_settings,json=storageSettings" json:"storage_settings,omitempty"`
	// The public key used for verifying tree heads and entry timestamps.
	// Readonly.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,14,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	// Interval after which a new signed root is produced even if there have been
	// no submission.  If zero, this behavior is disabled.
	MaxRootDuration *google_protobuf1.Duration `protobuf:"bytes,15,opt,name=max_root_duration,json=maxRootDuration" json:"max_root_duration,omitempty"`
	// Time of tree creation.
	// Readonly.
	CreateTime *google_protobuf2.Timestamp `protobuf:"bytes,16,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// Time of last tree update.
	// Readonly (automatically assigned on updates).
	UpdateTime *google_protobuf2.Timestamp `protobuf:"bytes,17,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
}

func (m *Tree) Reset()                    { *m = Tree{} }
func (m *Tree) String() string            { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()               {}
func (*Tree) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Tree) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *Tree) GetTreeState() TreeState {
	if m != nil {
		return m.TreeState
	}
	return TreeState_UNKNOWN_TREE_STATE
}

func (m *Tree) GetTreeType() TreeType {
	if m != nil {
		return m.TreeType
	}
	return TreeType_UNKNOWN_TREE_TYPE
}

func (m *Tree) GetHashStrategy() HashStrategy {
	if m != nil {
		return m.HashStrategy
	}
	return HashStrategy_UNKNOWN_HASH_STRATEGY
}

func (m *Tree) GetHashAlgorithm() sigpb.DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return sigpb.DigitallySigned_NONE
}

func (m *Tree) GetSignatureAlgorithm() sigpb.DigitallySigned_SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return sigpb.DigitallySigned_ANONYMOUS
}

func (m *Tree) GetSignatureCipherSuite() sigpb.DigitallySigned_SignatureCipherSuite {
	if m != nil {
		return m.SignatureCipherSuite
	}
	return sigpb.DigitallySigned_UNKNOWN_CIPHER_SUITE
}

func (m *Tree) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Tree) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tree) GetPrivateKey() *google_protobuf.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *Tree) GetStorageSettings() *google_protobuf.Any {
	if m != nil {
		return m.StorageSettings
	}
	return nil
}

func (m *Tree) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Tree) GetMaxRootDuration() *google_protobuf1.Duration {
	if m != nil {
		return m.MaxRootDuration
	}
	return nil
}

func (m *Tree) GetCreateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Tree) GetUpdateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type SignedEntryTimestamp struct {
	TimestampNanos int64                  `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	LogId          int64                  `protobuf:"varint,2,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	Signature      *sigpb.DigitallySigned `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *SignedEntryTimestamp) Reset()                    { *m = SignedEntryTimestamp{} }
func (m *SignedEntryTimestamp) String() string            { return proto.CompactTextString(m) }
func (*SignedEntryTimestamp) ProtoMessage()               {}
func (*SignedEntryTimestamp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SignedEntryTimestamp) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedEntryTimestamp) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *SignedEntryTimestamp) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

// SignedLogRoot represents a commitment by a Log to a particular tree.
type SignedLogRoot struct {
	// epoch nanoseconds, good until 2500ish
	TimestampNanos int64  `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	RootHash       []byte `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// TreeSize is the number of entries in the tree.
	TreeSize int64 `protobuf:"varint,3,opt,name=tree_size,json=treeSize" json:"tree_size,omitempty"`
	// TODO(al): define serialized format for the signature scheme.
	Signature    *sigpb.DigitallySigned `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	LogId        int64                  `protobuf:"varint,5,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	TreeRevision int64                  `protobuf:"varint,6,opt,name=tree_revision,json=treeRevision" json:"tree_revision,omitempty"`
}

func (m *SignedLogRoot) Reset()                    { *m = SignedLogRoot{} }
func (m *SignedLogRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedLogRoot) ProtoMessage()               {}
func (*SignedLogRoot) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *SignedLogRoot) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedLogRoot) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *SignedLogRoot) GetTreeSize() int64 {
	if m != nil {
		return m.TreeSize
	}
	return 0
}

func (m *SignedLogRoot) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedLogRoot) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *SignedLogRoot) GetTreeRevision() int64 {
	if m != nil {
		return m.TreeRevision
	}
	return 0
}

type MapperMetadata struct {
	SourceLogId                  []byte `protobuf:"bytes,1,opt,name=source_log_id,json=sourceLogId,proto3" json:"source_log_id,omitempty"`
	HighestFullyCompletedSeq     int64  `protobuf:"varint,2,opt,name=highest_fully_completed_seq,json=highestFullyCompletedSeq" json:"highest_fully_completed_seq,omitempty"`
	HighestPartiallyCompletedSeq int64  `protobuf:"varint,3,opt,name=highest_partially_completed_seq,json=highestPartiallyCompletedSeq" json:"highest_partially_completed_seq,omitempty"`
}

func (m *MapperMetadata) Reset()                    { *m = MapperMetadata{} }
func (m *MapperMetadata) String() string            { return proto.CompactTextString(m) }
func (*MapperMetadata) ProtoMessage()               {}
func (*MapperMetadata) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *MapperMetadata) GetSourceLogId() []byte {
	if m != nil {
		return m.SourceLogId
	}
	return nil
}

func (m *MapperMetadata) GetHighestFullyCompletedSeq() int64 {
	if m != nil {
		return m.HighestFullyCompletedSeq
	}
	return 0
}

func (m *MapperMetadata) GetHighestPartiallyCompletedSeq() int64 {
	if m != nil {
		return m.HighestPartiallyCompletedSeq
	}
	return 0
}

// SignedMapRoot represents a commitment by a Map to a particular tree.
type SignedMapRoot struct {
	TimestampNanos int64           `protobuf:"varint,1,opt,name=timestamp_nanos,json=timestampNanos" json:"timestamp_nanos,omitempty"`
	RootHash       []byte          `protobuf:"bytes,2,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	Metadata       *MapperMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	// TODO(al): define serialized format for the signature scheme.
	Signature   *sigpb.DigitallySigned `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	MapId       int64                  `protobuf:"varint,5,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	MapRevision int64                  `protobuf:"varint,6,opt,name=map_revision,json=mapRevision" json:"map_revision,omitempty"`
}

func (m *SignedMapRoot) Reset()                    { *m = SignedMapRoot{} }
func (m *SignedMapRoot) String() string            { return proto.CompactTextString(m) }
func (*SignedMapRoot) ProtoMessage()               {}
func (*SignedMapRoot) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *SignedMapRoot) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *SignedMapRoot) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *SignedMapRoot) GetMetadata() *MapperMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SignedMapRoot) GetSignature() *sigpb.DigitallySigned {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedMapRoot) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *SignedMapRoot) GetMapRevision() int64 {
	if m != nil {
		return m.MapRevision
	}
	return 0
}

func init() {
	proto.RegisterType((*Tree)(nil), "trillian.Tree")
	proto.RegisterType((*SignedEntryTimestamp)(nil), "trillian.SignedEntryTimestamp")
	proto.RegisterType((*SignedLogRoot)(nil), "trillian.SignedLogRoot")
	proto.RegisterType((*MapperMetadata)(nil), "trillian.MapperMetadata")
	proto.RegisterType((*SignedMapRoot)(nil), "trillian.SignedMapRoot")
	proto.RegisterEnum("trillian.HashStrategy", HashStrategy_name, HashStrategy_value)
	proto.RegisterEnum("trillian.TreeState", TreeState_name, TreeState_value)
	proto.RegisterEnum("trillian.TreeType", TreeType_name, TreeType_value)
}

func init() { proto.RegisterFile("trillian.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1056 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x6f, 0xdb, 0xb6,
	0x17, 0xad, 0x62, 0xd7, 0xb1, 0xaf, 0xff, 0x44, 0x61, 0x93, 0xfc, 0x94, 0xf4, 0x87, 0x35, 0xf3,
	0x06, 0x2c, 0xcb, 0x06, 0x7b, 0x73, 0x9a, 0x00, 0x43, 0x31, 0x0c, 0x8e, 0xa3, 0x34, 0x7f, 0x6d,
	0x43, 0xd2, 0x36, 0xb4, 0x2f, 0x04, 0x6d, 0xb3, 0x32, 0x51, 0xc9, 0x52, 0x25, 0xba, 0xa8, 0xfa,
	0xbc, 0xc7, 0x01, 0xfb, 0x3e, 0xfb, 0x3c, 0xfb, 0x16, 0x7b, 0x19, 0x48, 0x51, 0xb2, 0x93, 0x74,
	0x4b, 0x31, 0xec, 0x25, 0x21, 0xcf, 0x3d, 0xe7, 0xe8, 0x8a, 0xf7, 0x5e, 0xca, 0xd0, 0xe0, 0x11,
	0xf3, 0x3c, 0x46, 0x66, 0xad, 0x30, 0x0a, 0x78, 0x80, 0xca, 0xd9, 0x7e, 0xe7, 0xd0, 0x65, 0x7c,
	0x3a, 0x1f, 0xb5, 0xc6, 0x81, 0xdf, 0x76, 0x83, 0xc0, 0xf5, 0x68, 0x3b, 0x8b, 0xb5, 0xc7, 0x51,
	0x12, 0xf2, 0xa0, 0xfd, 0x9a, 0x26, 0x71, 0x38, 0x52, 0xff, 0x52, 0x83, 0x9d, 0x83, 0xfb, 0x65,
	0x31, 0x73, 0xc3, 0x51, 0xfa, 0x57, 0x89, 0xb6, 0x15, 0x53, 0xee, 0x46, 0xf3, 0x57, 0x6d, 0x32,
	0x4b, 0x54, 0xe8, 0x93, 0xdb, 0xa1, 0xc9, 0x3c, 0x22, 0x9c, 0x05, 0x2a, 0xe1, 0x9d, 0x27, 0xb7,
	0xe3, 0x9c, 0xf9, 0x34, 0xe6, 0xc4, 0x0f, 0x53, 0x42, 0xf3, 0xb7, 0x55, 0x28, 0x3a, 0x11, 0xa5,
	0xe8, 0x7f, 0xb0, 0xca, 0x23, 0x4a, 0x31, 0x9b, 0x18, 0xda, 0xae, 0xb6, 0x57, 0xb0, 0x4a, 0x62,
	0x7b, 0x3e, 0x41, 0x1d, 0x00, 0x19, 0x88, 0x39, 0xe1, 0xd4, 0x58, 0xd9, 0xd5, 0xf6, 0x1a, 0x9d,
	0x47, 0xad, 0xfc, 0x60, 0x84, 0xd8, 0x16, 0x21, 0xab, 0xc2, 0xb3, 0x25, 0x6a, 0x83, 0xdc, 0x60,
	0x9e, 0x84, 0xd4, 0x28, 0x48, 0x09, 0xba, 0x29, 0x71, 0x92, 0x90, 0x5a, 0x65, 0xae, 0x56, 0xe8,
	0x19, 0xd4, 0xa7, 0x24, 0x9e, 0xe2, 0x98, 0x47, 0x84, 0x53, 0x37, 0x31, 0x8a, 0x52, 0xb4, 0xb5,
	0x10, 0x9d, 0x91, 0x78, 0x6a, 0xab, 0xa8, 0x55, 0x9b, 0x2e, 0xed, 0xd0, 0x25, 0x34, 0xa4, 0x98,
	0x78, 0x6e, 0x10, 0x31, 0x3e, 0xf5, 0x8d, 0x87, 0x52, 0xfd, 0x79, 0x2b, 0x3d, 0xc5, 0x13, 0xe6,
	0x32, 0x4e, 0x3c, 0x2f, 0xb1, 0x99, 0x3b, 0xa3, 0x13, 0x69, 0xd5, 0xcd, 0xb8, 0x96, 0x7c, 0x70,
	0xbe, 0x45, 0x2f, 0xe1, 0x51, 0xcc, 0xdc, 0x19, 0xe1, 0xf3, 0x88, 0x2e, 0x39, 0x96, 0xa4, 0xe3,
	0x97, 0x7f, 0xe3, 0x68, 0x67, 0x8a, 0x85, 0x2d, 0x8a, 0xef, 0x60, 0x88, 0xc0, 0xd6, 0xc2, 0x7b,
	0xcc, 0xc2, 0x29, 0x8d, 0x70, 0x3c, 0x67, 0x9c, 0x1a, 0x48, 0xda, 0x7f, 0x75, 0x9f, 0x7d, 0x4f,
	0x6a, 0x6c, 0x21, 0xb1, 0x36, 0xe2, 0x0f, 0xa0, 0xe8, 0x53, 0xa8, 0x4d, 0x58, 0x1c, 0x7a, 0x24,
	0xc1, 0x33, 0xe2, 0x53, 0xa3, 0xbc, 0xab, 0xed, 0x55, 0xac, 0xaa, 0xc2, 0xfa, 0xc4, 0xa7, 0x68,
	0x17, 0xaa, 0x13, 0x1a, 0x8f, 0x23, 0x16, 0x8a, 0x46, 0x31, 0x2a, 0x8a, 0xb1, 0x80, 0xd0, 0x21,
	0x54, 0xc3, 0x88, 0xbd, 0x25, 0x9c, 0xe2, 0xd7, 0x34, 0x31, 0x6a, 0xbb, 0xda, 0x5e, 0xb5, 0xb3,
	0xd1, 0x4a, 0x7b, 0xa9, 0x95, 0xf5, 0x52, 0xab, 0x3b, 0x4b, 0x2c, 0x50, 0xc4, 0x4b, 0x9a, 0xa0,
	0x1f, 0x40, 0x8f, 0x79, 0x10, 0x11, 0x97, 0xe2, 0x98, 0x72, 0xce, 0x66, 0x6e, 0x6c, 0xd4, 0xff,
	0x41, 0xbb, 0xa6, 0xd8, 0xb6, 0x22, 0xa3, 0x6f, 0x00, 0xc2, 0xf9, 0xc8, 0x63, 0x63, 0xf9, 0xd8,
	0x86, 0x94, 0xae, 0xb7, 0xd4, 0x00, 0x0d, 0x65, 0xe4, 0x92, 0x26, 0x56, 0x25, 0xcc, 0x96, 0xc8,
	0x84, 0x75, 0x9f, 0xbc, 0xc3, 0x51, 0x10, 0x70, 0x9c, 0xb5, 0xbe, 0xb1, 0x26, 0x85, 0xdb, 0x77,
	0x9e, 0x79, 0xa2, 0x08, 0xd6, 0x9a, 0x4f, 0xde, 0x59, 0x41, 0xc0, 0x33, 0x00, 0x3d, 0x83, 0xea,
	0x38, 0xa2, 0xe2, 0x7d, 0xc5, 0x7c, 0x18, 0xba, 0x34, 0xd8, 0xb9, 0x63, 0xe0, 0x64, 0xc3, 0x63,
	0x41, 0x4a, 0x17, 0x80, 0x10, 0xcf, 0xc3, 0x49, 0x2e, 0x5e, 0xbf, 0x5f, 0x9c, 0xd2, 0x05, 0x70,
	0x51, 0x2c, 0xaf, 0xea, 0xe5, 0x8b, 0x62, 0x19, 0xf4, 0xea, 0x45, 0xb1, 0x5c, 0xd5, 0x6b, 0xcd,
	0x5f, 0x35, 0xd8, 0x48, 0xeb, 0x6e, 0xce, 0x78, 0x94, 0xe4, 0x32, 0xf4, 0x05, 0xac, 0xe5, 0xd3,
	0x8b, 0x67, 0x64, 0x16, 0xc4, 0x6a, 0x52, 0x1b, 0x39, 0xdc, 0x17, 0x28, 0xda, 0x84, 0x92, 0x17,
	0xb8, 0x62, 0x92, 0x57, 0x64, 0xfc, 0xa1, 0x17, 0xb8, 0xe7, 0x13, 0xf4, 0x14, 0x2a, 0x79, 0xcb,
	0xc8, 0xa1, 0xac, 0x76, 0xb6, 0x3e, 0xdc, 0x70, 0xd6, 0x82, 0xd8, 0xfc, 0x43, 0x83, 0x7a, 0x8a,
	0x5e, 0x05, 0xae, 0x38, 0xb4, 0x8f, 0xcf, 0xe3, 0x31, 0x54, 0x64, 0x61, 0xc4, 0x80, 0xc9, 0x54,
	0x6a, 0x56, 0x59, 0x00, 0x62, 0xfe, 0x44, 0x30, 0xbd, 0x56, 0xd8, 0xfb, 0x34, 0x9b, 0x42, 0x7a,
	0x1d, 0xd8, 0xec, 0x3d, 0xbd, 0x99, 0x6a, 0xf1, 0x23, 0x53, 0x5d, 0x7a, 0xef, 0x87, 0xcb, 0xef,
	0xfd, 0x19, 0xd4, 0xe5, 0x93, 0x22, 0xfa, 0x96, 0xc5, 0xa2, 0x3f, 0x4a, 0x32, 0x5a, 0x13, 0xa0,
	0xa5, 0xb0, 0xe6, 0xef, 0x1a, 0x34, 0xae, 0x49, 0x18, 0xd2, 0xe8, 0x9a, 0x72, 0x32, 0x21, 0x9c,
	0xa0, 0x26, 0xd4, 0xe3, 0x60, 0x1e, 0x8d, 0x29, 0x56, 0xae, 0x9a, 0x7c, 0x85, 0x6a, 0x0a, 0x5e,
	0x49, 0xef, 0xef, 0xe1, 0xf1, 0x94, 0xb9, 0x53, 0x1a, 0x73, 0xfc, 0x6a, 0xee, 0x79, 0x09, 0x1e,
	0x07, 0x7e, 0xe8, 0x51, 0x4e, 0x27, 0x38, 0xa6, 0x6f, 0xd4, 0xf9, 0x1b, 0x8a, 0x72, 0x2a, 0x18,
	0xbd, 0x8c, 0x60, 0xd3, 0x37, 0xc8, 0x84, 0x27, 0x99, 0x3c, 0x24, 0x11, 0x67, 0xe4, 0xae, 0x45,
	0x7a, 0x34, 0xff, 0x57, 0xb4, 0x61, 0xc6, 0x5a, 0xb6, 0x69, 0xfe, 0x99, 0xd7, 0xe8, 0x9a, 0x84,
	0xff, 0x61, 0x8d, 0x9e, 0x42, 0xd9, 0x57, 0xa7, 0xa1, 0x1a, 0xc6, 0x58, 0x5c, 0xc8, 0x37, 0x4f,
	0xcb, 0xca, 0x99, 0xff, 0xbe, 0x78, 0x3e, 0x09, 0x97, 0x8a, 0xe7, 0x93, 0xf0, 0x7c, 0x22, 0xee,
	0x33, 0x01, 0xdf, 0xaa, 0x5d, 0xd5, 0x27, 0x61, 0x56, 0xba, 0xfd, 0x5f, 0x34, 0xa8, 0x2d, 0x7f,
	0x1d, 0xd0, 0x36, 0x6c, 0xfe, 0xd8, 0xbf, 0xec, 0x0f, 0x7e, 0xee, 0xe3, 0xb3, 0xae, 0x7d, 0x86,
	0x6d, 0xc7, 0xea, 0x3a, 0xe6, 0xf3, 0x17, 0xfa, 0x03, 0x84, 0xa0, 0x61, 0x9d, 0xf6, 0x8e, 0xbe,
	0x3b, 0xea, 0x60, 0xfb, 0xac, 0xdb, 0x39, 0x3c, 0xd2, 0x35, 0xf4, 0x08, 0xd6, 0x1c, 0xd3, 0x76,
	0xf0, 0x75, 0x77, 0x28, 0xf9, 0xa6, 0xa5, 0xaf, 0x08, 0x8f, 0xc1, 0xf1, 0x85, 0xd9, 0x73, 0xf0,
	0x2d, 0x7e, 0x01, 0x6d, 0xc2, 0x7a, 0x6f, 0xd0, 0x3f, 0xbf, 0xb4, 0x05, 0x74, 0xf8, 0x6d, 0x07,
	0x0b, 0xb8, 0xb8, 0x8f, 0xa1, 0x92, 0x7f, 0x0b, 0xd1, 0x16, 0xa0, 0x2c, 0x05, 0xc7, 0x32, 0x4d,
	0x6c, 0x3b, 0x5d, 0xc7, 0xd4, 0x1f, 0x20, 0x80, 0x52, 0xb7, 0xe7, 0x9c, 0xff, 0x64, 0xea, 0x9a,
	0x58, 0x9f, 0x5a, 0x83, 0x97, 0x66, 0x5f, 0x5f, 0x41, 0x3a, 0xd4, 0xec, 0xc1, 0xa9, 0x83, 0x4f,
	0xcc, 0x2b, 0xd3, 0x31, 0x4f, 0xf4, 0x82, 0x40, 0xce, 0xba, 0xd6, 0x49, 0x8e, 0x14, 0xf7, 0x0f,
	0xa0, 0x9c, 0x7d, 0x39, 0x45, 0x0e, 0x37, 0xfc, 0x9d, 0x17, 0x43, 0x61, 0xbf, 0x0a, 0x85, 0xab,
	0xc1, 0x73, 0x5d, 0x13, 0x8b, 0xeb, 0xee, 0x50, 0x5f, 0x39, 0xfe, 0x1a, 0xb6, 0xc7, 0x81, 0x9f,
	0x5d, 0x46, 0x37, 0x7f, 0xce, 0x1c, 0xd7, 0x1d, 0xb5, 0x1f, 0x8a, 0xed, 0x50, 0x1b, 0x95, 0x24,
	0x7e, 0xf0, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x75, 0x85, 0xca, 0xf8, 0x08, 0x00, 0x00,
}
