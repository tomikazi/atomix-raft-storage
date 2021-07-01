// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/config/config.proto

package config

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProtocolConfig struct {
	ElectionTimeout   *time.Duration `protobuf:"bytes,1,opt,name=election_timeout,json=electionTimeout,proto3,stdduration" json:"election_timeout,omitempty"`
	HeartbeatInterval *time.Duration `protobuf:"bytes,2,opt,name=heartbeat_interval,json=heartbeatInterval,proto3,stdduration" json:"heartbeat_interval,omitempty"`
	SnapshotInterval  *time.Duration `protobuf:"bytes,3,opt,name=snapshot_interval,json=snapshotInterval,proto3,stdduration" json:"snapshot_interval,omitempty"`
	SnapshotThreshold uint64         `protobuf:"varint,4,opt,name=snapshot_threshold,json=snapshotThreshold,proto3" json:"snapshot_threshold,omitempty"`
}

func (m *ProtocolConfig) Reset()         { *m = ProtocolConfig{} }
func (m *ProtocolConfig) String() string { return proto.CompactTextString(m) }
func (*ProtocolConfig) ProtoMessage()    {}
func (*ProtocolConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac523a84bbf07b3d, []int{0}
}
func (m *ProtocolConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolConfig.Merge(m, src)
}
func (m *ProtocolConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolConfig proto.InternalMessageInfo

func (m *ProtocolConfig) GetElectionTimeout() *time.Duration {
	if m != nil {
		return m.ElectionTimeout
	}
	return nil
}

func (m *ProtocolConfig) GetHeartbeatInterval() *time.Duration {
	if m != nil {
		return m.HeartbeatInterval
	}
	return nil
}

func (m *ProtocolConfig) GetSnapshotInterval() *time.Duration {
	if m != nil {
		return m.SnapshotInterval
	}
	return nil
}

func (m *ProtocolConfig) GetSnapshotThreshold() uint64 {
	if m != nil {
		return m.SnapshotThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtocolConfig)(nil), "atomix.raft.config.ProtocolConfig")
}

func init() { proto.RegisterFile("storage/config/config.proto", fileDescriptor_ac523a84bbf07b3d) }

var fileDescriptor_ac523a84bbf07b3d = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8d, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xeb, 0x12, 0x31, 0x18, 0x09, 0x5a, 0x8b, 0x21, 0x14, 0xc9, 0x54, 0x88, 0xa1, 0x0b,
	0x8e, 0x04, 0x37, 0x28, 0x2c, 0x20, 0x84, 0x50, 0xd5, 0xbd, 0x72, 0x52, 0xc7, 0x89, 0x94, 0xe6,
	0x55, 0xce, 0x0b, 0xe2, 0x18, 0x8c, 0x1c, 0x01, 0x6e, 0xc0, 0x11, 0x18, 0x3b, 0xb2, 0x01, 0xc9,
	0x25, 0x18, 0x51, 0xe2, 0x24, 0x62, 0xec, 0x94, 0x5f, 0x79, 0xfe, 0xbe, 0x8f, 0x1e, 0x67, 0x08,
	0x46, 0x6a, 0xe5, 0x05, 0x90, 0x86, 0xb1, 0x6e, 0x3e, 0x62, 0x6d, 0x00, 0x81, 0x31, 0x89, 0xb0,
	0x8a, 0x9f, 0x84, 0x91, 0x21, 0x0a, 0x7b, 0x19, 0x71, 0x0d, 0xa0, 0x13, 0xe5, 0xd5, 0x2f, 0xfc,
	0x3c, 0xf4, 0x96, 0xb9, 0x91, 0x18, 0x43, 0x6a, 0x99, 0xd1, 0xa1, 0x06, 0x0d, 0xf5, 0xf4, 0xaa,
	0x65, 0xff, 0x9e, 0xbe, 0xf5, 0xe9, 0xfe, 0x43, 0xb5, 0x02, 0x48, 0xae, 0x6a, 0x11, 0xbb, 0xa5,
	0x03, 0x95, 0xa8, 0xa0, 0x42, 0x17, 0x18, 0xaf, 0x14, 0xe4, 0xe8, 0x92, 0x31, 0x99, 0xec, 0x5d,
	0x1c, 0x09, 0xdb, 0x10, 0x6d, 0x43, 0x5c, 0x37, 0x8d, 0xa9, 0xf3, 0xf2, 0x75, 0x42, 0x66, 0x07,
	0x2d, 0x38, 0xb7, 0x1c, 0xbb, 0xa7, 0x2c, 0x52, 0xd2, 0xa0, 0xaf, 0x24, 0x2e, 0xe2, 0x14, 0x95,
	0x79, 0x94, 0x89, 0xdb, 0xdf, 0xce, 0x36, 0xec, 0xd0, 0x9b, 0x86, 0x64, 0x77, 0x74, 0x98, 0xa5,
	0x72, 0x9d, 0x45, 0xf0, 0x4f, 0xb7, 0xb3, 0x9d, 0x6e, 0xd0, 0x92, 0x9d, 0xed, 0x9c, 0xb2, 0xce,
	0x86, 0x91, 0x51, 0x59, 0x04, 0xc9, 0xd2, 0x75, 0xc6, 0x64, 0xe2, 0xcc, 0xba, 0xce, 0xbc, 0x3d,
	0x4c, 0xcf, 0x7e, 0x7f, 0x38, 0x79, 0x2d, 0x38, 0x79, 0x2f, 0x38, 0xf9, 0x28, 0x38, 0xd9, 0x14,
	0x9c, 0x7c, 0x17, 0x9c, 0x3c, 0x97, 0xbc, 0xb7, 0x29, 0x79, 0xef, 0xb3, 0xe4, 0x3d, 0x7f, 0xb7,
	0xee, 0x5f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x3f, 0xb7, 0x72, 0xc1, 0x01, 0x00, 0x00,
}

func (this *ProtocolConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProtocolConfig)
	if !ok {
		that2, ok := that.(ProtocolConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ElectionTimeout != nil && that1.ElectionTimeout != nil {
		if *this.ElectionTimeout != *that1.ElectionTimeout {
			return false
		}
	} else if this.ElectionTimeout != nil {
		return false
	} else if that1.ElectionTimeout != nil {
		return false
	}
	if this.HeartbeatInterval != nil && that1.HeartbeatInterval != nil {
		if *this.HeartbeatInterval != *that1.HeartbeatInterval {
			return false
		}
	} else if this.HeartbeatInterval != nil {
		return false
	} else if that1.HeartbeatInterval != nil {
		return false
	}
	if this.SnapshotInterval != nil && that1.SnapshotInterval != nil {
		if *this.SnapshotInterval != *that1.SnapshotInterval {
			return false
		}
	} else if this.SnapshotInterval != nil {
		return false
	} else if that1.SnapshotInterval != nil {
		return false
	}
	if this.SnapshotThreshold != that1.SnapshotThreshold {
		return false
	}
	return true
}
func (m *ProtocolConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SnapshotThreshold != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.SnapshotThreshold))
		i--
		dAtA[i] = 0x20
	}
	if m.SnapshotInterval != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.SnapshotInterval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.SnapshotInterval):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintConfig(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x1a
	}
	if m.HeartbeatInterval != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.HeartbeatInterval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.HeartbeatInterval):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintConfig(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x12
	}
	if m.ElectionTimeout != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(*m.ElectionTimeout, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ElectionTimeout):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintConfig(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedProtocolConfig(r randyConfig, easy bool) *ProtocolConfig {
	this := &ProtocolConfig{}
	if r.Intn(5) != 0 {
		this.ElectionTimeout = github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	}
	if r.Intn(5) != 0 {
		this.HeartbeatInterval = github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	}
	if r.Intn(5) != 0 {
		this.SnapshotInterval = github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	}
	this.SnapshotThreshold = uint64(uint64(r.Uint32()))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyConfig interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneConfig(r randyConfig) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringConfig(r randyConfig) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneConfig(r)
	}
	return string(tmps)
}
func randUnrecognizedConfig(r randyConfig, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldConfig(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldConfig(dAtA []byte, r randyConfig, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateConfig(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateConfig(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ProtocolConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ElectionTimeout != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.ElectionTimeout)
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.HeartbeatInterval != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.HeartbeatInterval)
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.SnapshotInterval != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdDuration(*m.SnapshotInterval)
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.SnapshotThreshold != 0 {
		n += 1 + sovConfig(uint64(m.SnapshotThreshold))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtocolConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtocolConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElectionTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ElectionTimeout == nil {
				m.ElectionTimeout = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.ElectionTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HeartbeatInterval == nil {
				m.HeartbeatInterval = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.HeartbeatInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SnapshotInterval == nil {
				m.SnapshotInterval = new(time.Duration)
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.SnapshotInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotThreshold", wireType)
			}
			m.SnapshotThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapshotThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)