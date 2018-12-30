// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3.proto

package proto3_proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	test_proto "github.com/gogo/protobuf/proto/test_proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}

var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}

func (Message_Humour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{0, 0}
}

type Message struct {
	Name                 string                             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity             Message_Humour                     `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm           uint32                             `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm,proto3" json:"height_in_cm,omitempty"`
	Data                 []byte                             `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount          int64                              `protobuf:"varint,7,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	TrueScotsman         bool                               `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman,proto3" json:"true_scotsman,omitempty"`
	Score                float32                            `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key                  []uint64                           `protobuf:"varint,5,rep,packed,name=key,proto3" json:"key,omitempty"`
	ShortKey             []int32                            `protobuf:"varint,19,rep,packed,name=short_key,json=shortKey,proto3" json:"short_key,omitempty"`
	Nested               *Nested                            `protobuf:"bytes,6,opt,name=nested,proto3" json:"nested,omitempty"`
	RFunny               []Message_Humour                   `protobuf:"varint,16,rep,packed,name=r_funny,json=rFunny,proto3,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain              map[string]*Nested                 `protobuf:"bytes,10,rep,name=terrain,proto3" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Proto2Field          *test_proto.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field,proto3" json:"proto2_field,omitempty"`
	Proto2Value          map[string]*test_proto.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value,proto3" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Anything             *types.Any                         `protobuf:"bytes,14,opt,name=anything,proto3" json:"anything,omitempty"`
	ManyThings           []*types.Any                       `protobuf:"bytes,15,rep,name=many_things,json=manyThings,proto3" json:"many_things,omitempty"`
	Submessage           *Message                           `protobuf:"bytes,17,opt,name=submessage,proto3" json:"submessage,omitempty"`
	Children             []*Message                         `protobuf:"bytes,18,rep,name=children,proto3" json:"children,omitempty"`
	StringMap            map[string]string                  `protobuf:"bytes,20,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetHilarity() Message_Humour {
	if m != nil {
		return m.Hilarity
	}
	return Message_UNKNOWN
}

func (m *Message) GetHeightInCm() uint32 {
	if m != nil {
		return m.HeightInCm
	}
	return 0
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetResultCount() int64 {
	if m != nil {
		return m.ResultCount
	}
	return 0
}

func (m *Message) GetTrueScotsman() bool {
	if m != nil {
		return m.TrueScotsman
	}
	return false
}

func (m *Message) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Message) GetKey() []uint64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetShortKey() []int32 {
	if m != nil {
		return m.ShortKey
	}
	return nil
}

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetRFunny() []Message_Humour {
	if m != nil {
		return m.RFunny
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

func (m *Message) GetAnything() *types.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Message) GetManyThings() []*types.Any {
	if m != nil {
		return m.ManyThings
	}
	return nil
}

func (m *Message) GetSubmessage() *Message {
	if m != nil {
		return m.Submessage
	}
	return nil
}

func (m *Message) GetChildren() []*Message {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Message) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type Nested struct {
	Bunny                string   `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
	Cute                 bool     `protobuf:"varint,2,opt,name=cute,proto3" json:"cute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{1}
}
func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetBunny() string {
	if m != nil {
		return m.Bunny
	}
	return ""
}

func (m *Nested) GetCute() bool {
	if m != nil {
		return m.Cute
	}
	return false
}

type MessageWithMap struct {
	ByteMapping          map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping,proto3" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MessageWithMap) Reset()         { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()    {}
func (*MessageWithMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{2}
}
func (m *MessageWithMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithMap.Unmarshal(m, b)
}
func (m *MessageWithMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithMap.Marshal(b, m, deterministic)
}
func (m *MessageWithMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithMap.Merge(m, src)
}
func (m *MessageWithMap) XXX_Size() int {
	return xxx_messageInfo_MessageWithMap.Size(m)
}
func (m *MessageWithMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithMap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithMap proto.InternalMessageInfo

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

type IntMap struct {
	Rtt                  map[int32]int32 `protobuf:"bytes,1,rep,name=rtt,proto3" json:"rtt,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{3}
}
func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (m *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(m, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetRtt() map[int32]int32 {
	if m != nil {
		return m.Rtt
	}
	return nil
}

type IntMaps struct {
	Maps                 []*IntMap `protobuf:"bytes,1,rep,name=maps,proto3" json:"maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IntMaps) Reset()         { *m = IntMaps{} }
func (m *IntMaps) String() string { return proto.CompactTextString(m) }
func (*IntMaps) ProtoMessage()    {}
func (*IntMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{4}
}
func (m *IntMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMaps.Unmarshal(m, b)
}
func (m *IntMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMaps.Marshal(b, m, deterministic)
}
func (m *IntMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMaps.Merge(m, src)
}
func (m *IntMaps) XXX_Size() int {
	return xxx_messageInfo_IntMaps.Size(m)
}
func (m *IntMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMaps.DiscardUnknown(m)
}

var xxx_messageInfo_IntMaps proto.InternalMessageInfo

func (m *IntMaps) GetMaps() []*IntMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

type TestUTF8 struct {
	Scalar string   `protobuf:"bytes,1,opt,name=scalar,proto3" json:"scalar,omitempty"`
	Vector []string `protobuf:"bytes,2,rep,name=vector,proto3" json:"vector,omitempty"`
	// Types that are valid to be assigned to Oneof:
	//	*TestUTF8_Field
	Oneof                isTestUTF8_Oneof `protobuf_oneof:"oneof"`
	MapKey               map[string]int64 `protobuf:"bytes,4,rep,name=map_key,json=mapKey,proto3" json:"map_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapValue             map[int64]string `protobuf:"bytes,5,rep,name=map_value,json=mapValue,proto3" json:"map_value,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestUTF8) Reset()         { *m = TestUTF8{} }
func (m *TestUTF8) String() string { return proto.CompactTextString(m) }
func (*TestUTF8) ProtoMessage()    {}
func (*TestUTF8) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{5}
}
func (m *TestUTF8) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestUTF8.Unmarshal(m, b)
}
func (m *TestUTF8) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestUTF8.Marshal(b, m, deterministic)
}
func (m *TestUTF8) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUTF8.Merge(m, src)
}
func (m *TestUTF8) XXX_Size() int {
	return xxx_messageInfo_TestUTF8.Size(m)
}
func (m *TestUTF8) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUTF8.DiscardUnknown(m)
}

var xxx_messageInfo_TestUTF8 proto.InternalMessageInfo

type isTestUTF8_Oneof interface {
	isTestUTF8_Oneof()
}

type TestUTF8_Field struct {
	Field string `protobuf:"bytes,3,opt,name=field,proto3,oneof"`
}

func (*TestUTF8_Field) isTestUTF8_Oneof() {}

func (m *TestUTF8) GetOneof() isTestUTF8_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (m *TestUTF8) GetScalar() string {
	if m != nil {
		return m.Scalar
	}
	return ""
}

func (m *TestUTF8) GetVector() []string {
	if m != nil {
		return m.Vector
	}
	return nil
}

func (m *TestUTF8) GetField() string {
	if x, ok := m.GetOneof().(*TestUTF8_Field); ok {
		return x.Field
	}
	return ""
}

func (m *TestUTF8) GetMapKey() map[string]int64 {
	if m != nil {
		return m.MapKey
	}
	return nil
}

func (m *TestUTF8) GetMapValue() map[int64]string {
	if m != nil {
		return m.MapValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TestUTF8) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TestUTF8_OneofMarshaler, _TestUTF8_OneofUnmarshaler, _TestUTF8_OneofSizer, []interface{}{
		(*TestUTF8_Field)(nil),
	}
}

func _TestUTF8_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TestUTF8)
	// oneof
	switch x := m.Oneof.(type) {
	case *TestUTF8_Field:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Field)
	case nil:
	default:
		return fmt.Errorf("TestUTF8.Oneof has unexpected type %T", x)
	}
	return nil
}

func _TestUTF8_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TestUTF8)
	switch tag {
	case 3: // oneof.field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Oneof = &TestUTF8_Field{x}
		return true, err
	default:
		return false, nil
	}
}

func _TestUTF8_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TestUTF8)
	// oneof
	switch x := m.Oneof.(type) {
	case *TestUTF8_Field:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Field)))
		n += len(x.Field)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterMapType((map[string]*test_proto.SubDefaults)(nil), "proto3_proto.Message.Proto2ValueEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto3_proto.Message.StringMapEntry")
	proto.RegisterMapType((map[string]*Nested)(nil), "proto3_proto.Message.TerrainEntry")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterMapType((map[bool][]byte)(nil), "proto3_proto.MessageWithMap.ByteMappingEntry")
	proto.RegisterType((*IntMap)(nil), "proto3_proto.IntMap")
	proto.RegisterMapType((map[int32]int32)(nil), "proto3_proto.IntMap.RttEntry")
	proto.RegisterType((*IntMaps)(nil), "proto3_proto.IntMaps")
	proto.RegisterType((*TestUTF8)(nil), "proto3_proto.TestUTF8")
	proto.RegisterMapType((map[string]int64)(nil), "proto3_proto.TestUTF8.MapKeyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "proto3_proto.TestUTF8.MapValueEntry")
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
}

func init() { proto.RegisterFile("proto3.proto", fileDescriptor_4fee6d65e34a64b6) }

var fileDescriptor_4fee6d65e34a64b6 = []byte{
	// 891 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xff, 0x8e, 0xdb, 0x44,
	0x10, 0xae, 0xe3, 0xfc, 0x70, 0x26, 0xc9, 0x35, 0x2c, 0x69, 0x59, 0x02, 0x48, 0x26, 0x20, 0x64,
	0x21, 0xea, 0x83, 0x54, 0x87, 0x8e, 0xb6, 0x02, 0xdd, 0x1d, 0x3d, 0x35, 0xba, 0x4b, 0x88, 0x36,
	0x39, 0x4e, 0xfc, 0x65, 0x6d, 0x72, 0x9b, 0xc4, 0x22, 0x5e, 0x07, 0xef, 0xba, 0x92, 0x5f, 0x80,
	0x07, 0xe1, 0x95, 0x78, 0x21, 0xb4, 0xbb, 0xce, 0xd5, 0xa9, 0x5c, 0xee, 0x2f, 0xef, 0x7c, 0xfe,
	0x66, 0xbe, 0xd9, 0x99, 0xd9, 0x81, 0xf6, 0x2e, 0x89, 0x65, 0xfc, 0xdc, 0xd7, 0x1f, 0x94, 0x5b,
	0x81, 0xfe, 0xf4, 0x3f, 0x5d, 0xc7, 0xf1, 0x7a, 0xcb, 0x8e, 0xb5, 0xb5, 0x48, 0x57, 0xc7, 0x94,
	0x67, 0x86, 0xd8, 0x7f, 0x22, 0x99, 0x90, 0x86, 0x76, 0xac, 0x8e, 0x06, 0x1e, 0xfc, 0xdd, 0x84,
	0xc6, 0x98, 0x09, 0x41, 0xd7, 0x0c, 0x21, 0xa8, 0x72, 0x1a, 0x31, 0x6c, 0xb9, 0x96, 0xd7, 0x24,
	0xfa, 0x8c, 0x4e, 0xc1, 0xd9, 0x84, 0x5b, 0x9a, 0x84, 0x32, 0xc3, 0x15, 0xd7, 0xf2, 0x8e, 0x86,
	0x9f, 0xfb, 0x45, 0x49, 0x3f, 0x77, 0xf6, 0xdf, 0xa4, 0x51, 0x9c, 0x26, 0xe4, 0x9e, 0x8d, 0x5c,
	0x68, 0x6f, 0x58, 0xb8, 0xde, 0xc8, 0x20, 0xe4, 0xc1, 0x32, 0xc2, 0xb6, 0x6b, 0x79, 0x1d, 0x02,
	0x06, 0x1b, 0xf1, 0x8b, 0x48, 0xe9, 0xdd, 0x51, 0x49, 0x71, 0xd5, 0xb5, 0xbc, 0x36, 0xd1, 0x67,
	0xf4, 0x25, 0xb4, 0x13, 0x26, 0xd2, 0xad, 0x0c, 0x96, 0x71, 0xca, 0x25, 0x6e, 0xb8, 0x96, 0x67,
	0x93, 0x96, 0xc1, 0x2e, 0x14, 0x84, 0xbe, 0x82, 0x8e, 0x4c, 0x52, 0x16, 0x88, 0x65, 0x2c, 0x45,
	0x44, 0x39, 0x76, 0x5c, 0xcb, 0x73, 0x48, 0x5b, 0x81, 0xb3, 0x1c, 0x43, 0x3d, 0xa8, 0x89, 0x65,
	0x9c, 0x30, 0xdc, 0x74, 0x2d, 0xaf, 0x42, 0x8c, 0x81, 0xba, 0x60, 0xff, 0xc9, 0x32, 0x5c, 0x73,
	0x6d, 0xaf, 0x4a, 0xd4, 0x11, 0x7d, 0x06, 0x4d, 0xb1, 0x89, 0x13, 0x19, 0x28, 0xfc, 0x63, 0xd7,
	0xf6, 0x6a, 0xc4, 0xd1, 0xc0, 0x15, 0xcb, 0xd0, 0x77, 0x50, 0xe7, 0x4c, 0x48, 0x76, 0x87, 0xeb,
	0xae, 0xe5, 0xb5, 0x86, 0xbd, 0xc3, 0xab, 0x4f, 0xf4, 0x3f, 0x92, 0x73, 0xd0, 0x09, 0x34, 0x92,
	0x60, 0x95, 0x72, 0x9e, 0xe1, 0xae, 0x6b, 0x3f, 0x58, 0xa9, 0x7a, 0x72, 0xa9, 0xb8, 0xe8, 0x15,
	0x34, 0x24, 0x4b, 0x12, 0x1a, 0x72, 0x0c, 0xae, 0xed, 0xb5, 0x86, 0x83, 0x72, 0xb7, 0xb9, 0x21,
	0xbd, 0xe6, 0x32, 0xc9, 0xc8, 0xde, 0x05, 0xbd, 0xc8, 0xe7, 0x61, 0x18, 0xac, 0x42, 0xb6, 0xbd,
	0xc3, 0x2d, 0x9d, 0xe8, 0x27, 0xfe, 0xbb, 0x6e, 0xfb, 0xb3, 0x74, 0xf1, 0x2b, 0x5b, 0xd1, 0x74,
	0x2b, 0x05, 0x69, 0x19, 0xf2, 0xa5, 0xe2, 0xa2, 0xd1, 0xbd, 0xef, 0x5b, 0xba, 0x4d, 0x19, 0xee,
	0x68, 0xf9, 0x6f, 0xca, 0xe5, 0xa7, 0x9a, 0xf9, 0xbb, 0x22, 0x9a, 0x14, 0xf2, 0x50, 0x1a, 0x41,
	0xdf, 0x83, 0x43, 0x79, 0x26, 0x37, 0x21, 0x5f, 0xe3, 0xa3, 0xbc, 0x56, 0x66, 0x16, 0xfd, 0xfd,
	0x2c, 0xfa, 0x67, 0x3c, 0x23, 0xf7, 0x2c, 0x74, 0x02, 0xad, 0x88, 0xf2, 0x2c, 0xd0, 0x96, 0xc0,
	0x8f, 0xb5, 0x76, 0xb9, 0x13, 0x28, 0xe2, 0x5c, 0xf3, 0xd0, 0x09, 0x80, 0x48, 0x17, 0x91, 0x49,
	0x0a, 0x7f, 0xa4, 0xa5, 0x9e, 0x94, 0x66, 0x4c, 0x0a, 0x44, 0xf4, 0x03, 0x38, 0xcb, 0x4d, 0xb8,
	0xbd, 0x4b, 0x18, 0xc7, 0x48, 0x4b, 0x7d, 0xc0, 0xe9, 0x9e, 0x86, 0x2e, 0x00, 0x84, 0x4c, 0x42,
	0xbe, 0x0e, 0x22, 0xba, 0xc3, 0x3d, 0xed, 0xf4, 0x75, 0x79, 0x6d, 0x66, 0x9a, 0x37, 0xa6, 0x3b,
	0x53, 0x99, 0xa6, 0xd8, 0xdb, 0xfd, 0x29, 0xb4, 0x8b, 0x7d, 0xdb, 0x0f, 0xa0, 0x79, 0x61, 0x7a,
	0x00, 0xbf, 0x85, 0x9a, 0xa9, 0x7e, 0xe5, 0x7f, 0x46, 0xcc, 0x50, 0x5e, 0x54, 0x4e, 0xad, 0xfe,
	0x2d, 0x74, 0xdf, 0x6f, 0x45, 0x49, 0xd4, 0x67, 0x87, 0x51, 0x3f, 0x38, 0x0f, 0x85, 0xc0, 0xaf,
	0xe0, 0xe8, 0xf0, 0x1e, 0x25, 0x61, 0x7b, 0xc5, 0xb0, 0xcd, 0x82, 0xf7, 0xe0, 0x17, 0xa8, 0x9b,
	0xb9, 0x46, 0x2d, 0x68, 0xdc, 0x4c, 0xae, 0x26, 0xbf, 0xdd, 0x4e, 0xba, 0x8f, 0x90, 0x03, 0xd5,
	0xe9, 0xcd, 0x64, 0xd6, 0xb5, 0x50, 0x07, 0x9a, 0xb3, 0xeb, 0xb3, 0xe9, 0x6c, 0x3e, 0xba, 0xb8,
	0xea, 0x56, 0xd0, 0x63, 0x68, 0x9d, 0x8f, 0xae, 0xaf, 0x83, 0xf3, 0xb3, 0xd1, 0xf5, 0xeb, 0x3f,
	0xba, 0xf6, 0x60, 0x08, 0x75, 0x73, 0x59, 0x25, 0xb2, 0xd0, 0xaf, 0xc8, 0x08, 0x1b, 0x43, 0x2d,
	0x8b, 0x65, 0x2a, 0x8d, 0xb2, 0x43, 0xf4, 0x79, 0xf0, 0x8f, 0x05, 0x47, 0x79, 0x0f, 0x6e, 0x43,
	0xb9, 0x19, 0xd3, 0x1d, 0x9a, 0x42, 0x7b, 0x91, 0x49, 0xa6, 0x7a, 0xb6, 0x53, 0xc3, 0x68, 0xe9,
	0xbe, 0x3d, 0x2b, 0xed, 0x5b, 0xee, 0xe3, 0x9f, 0x67, 0x92, 0x8d, 0x0d, 0x3f, 0x1f, 0xed, 0xc5,
	0x3b, 0xa4, 0xff, 0x33, 0x74, 0xdf, 0x27, 0x14, 0x2b, 0xe3, 0x94, 0x54, 0xa6, 0x5d, 0xac, 0xcc,
	0x5f, 0x50, 0x1f, 0x71, 0xa9, 0x72, 0x3b, 0x06, 0x3b, 0x91, 0x32, 0x4f, 0xe9, 0x8b, 0xc3, 0x94,
	0x0c, 0xc5, 0x27, 0x52, 0x9a, 0x14, 0x14, 0xb3, 0xff, 0x23, 0x38, 0x7b, 0xa0, 0x28, 0x59, 0x2b,
	0x91, 0xac, 0x15, 0x25, 0x9f, 0x43, 0xc3, 0xc4, 0x13, 0xc8, 0x83, 0x6a, 0x44, 0x77, 0x22, 0x17,
	0xed, 0x95, 0x89, 0x12, 0xcd, 0x18, 0xfc, 0x5b, 0x01, 0x67, 0xce, 0x84, 0xbc, 0x99, 0x5f, 0x9e,
	0xa2, 0xa7, 0x50, 0x17, 0x4b, 0xba, 0xa5, 0x49, 0xde, 0x84, 0xdc, 0x52, 0xf8, 0x5b, 0xb6, 0x94,
	0x71, 0x82, 0x2b, 0xae, 0xad, 0x70, 0x63, 0xa1, 0xa7, 0x50, 0x33, 0xfb, 0x47, 0x6d, 0xf9, 0xe6,
	0x9b, 0x47, 0xc4, 0x98, 0xe8, 0x25, 0x34, 0x22, 0xba, 0xd3, 0xcb, 0xb5, 0x5a, 0xb6, 0xdc, 0xf6,
	0x82, 0xfe, 0x98, 0xee, 0xae, 0x58, 0x66, 0xee, 0x5e, 0x8f, 0xb4, 0x81, 0xce, 0xa0, 0xa9, 0x9c,
	0xcd, 0x25, 0x6b, 0x65, 0x0f, 0xb0, 0xe8, 0x5e, 0x58, 0x4d, 0x4e, 0x94, 0x9b, 0xfd, 0x9f, 0xa0,
	0x55, 0x88, 0xfc, 0xd0, 0x44, 0xdb, 0xc5, 0xf7, 0xf0, 0x12, 0x3a, 0x07, 0x51, 0x8b, 0xce, 0xf6,
	0x03, 0xcf, 0xe1, 0xbc, 0x01, 0xb5, 0x98, 0xb3, 0x78, 0xb5, 0xa8, 0x9b, 0x7c, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0x22, 0xea, 0x15, 0xb6, 0x07, 0x00, 0x00,
}
