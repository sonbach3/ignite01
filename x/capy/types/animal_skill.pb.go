// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: capy/capy/animal_skill.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AnimalSkill struct {
	Id      uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Skill   []uint64 `protobuf:"varint,2,rep,packed,name=skill,proto3" json:"skill,omitempty"`
	Creator string   `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *AnimalSkill) Reset()         { *m = AnimalSkill{} }
func (m *AnimalSkill) String() string { return proto.CompactTextString(m) }
func (*AnimalSkill) ProtoMessage()    {}
func (*AnimalSkill) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad9ed40d2a9b09f, []int{0}
}
func (m *AnimalSkill) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AnimalSkill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AnimalSkill.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AnimalSkill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnimalSkill.Merge(m, src)
}
func (m *AnimalSkill) XXX_Size() int {
	return m.Size()
}
func (m *AnimalSkill) XXX_DiscardUnknown() {
	xxx_messageInfo_AnimalSkill.DiscardUnknown(m)
}

var xxx_messageInfo_AnimalSkill proto.InternalMessageInfo

func (m *AnimalSkill) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AnimalSkill) GetSkill() []uint64 {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *AnimalSkill) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*AnimalSkill)(nil), "capy.capy.AnimalSkill")
}

func init() { proto.RegisterFile("capy/capy/animal_skill.proto", fileDescriptor_0ad9ed40d2a9b09f) }

var fileDescriptor_0ad9ed40d2a9b09f = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4e, 0x2c, 0xa8,
	0xd4, 0x07, 0x13, 0x89, 0x79, 0x99, 0xb9, 0x89, 0x39, 0xf1, 0xc5, 0xd9, 0x99, 0x39, 0x39, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x20, 0x09, 0x3d, 0x10, 0xa1, 0xe4, 0xcb, 0xc5, 0xed,
	0x08, 0x56, 0x10, 0x0c, 0x92, 0x17, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0x05, 0x6b, 0x94, 0x60, 0x52, 0x60, 0xd6,
	0x60, 0x09, 0x82, 0x70, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24,
	0x98, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x27, 0xed, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x88, 0x12, 0x04, 0x3b, 0xa6, 0x02, 0xe2, 0xa6, 0x92, 0xca, 0x82, 0xd4, 0xe2,
	0x24, 0x36, 0xb0, 0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x57, 0x8c, 0xb3, 0xad,
	0x00, 0x00, 0x00,
}

func (m *AnimalSkill) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AnimalSkill) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AnimalSkill) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAnimalSkill(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Skill) > 0 {
		dAtA2 := make([]byte, len(m.Skill)*10)
		var j1 int
		for _, num := range m.Skill {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintAnimalSkill(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAnimalSkill(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAnimalSkill(dAtA []byte, offset int, v uint64) int {
	offset -= sovAnimalSkill(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AnimalSkill) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAnimalSkill(uint64(m.Id))
	}
	if len(m.Skill) > 0 {
		l = 0
		for _, e := range m.Skill {
			l += sovAnimalSkill(uint64(e))
		}
		n += 1 + sovAnimalSkill(uint64(l)) + l
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAnimalSkill(uint64(l))
	}
	return n
}

func sovAnimalSkill(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAnimalSkill(x uint64) (n int) {
	return sovAnimalSkill(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AnimalSkill) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAnimalSkill
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
			return fmt.Errorf("proto: AnimalSkill: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AnimalSkill: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimalSkill
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAnimalSkill
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Skill = append(m.Skill, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAnimalSkill
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthAnimalSkill
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthAnimalSkill
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Skill) == 0 {
					m.Skill = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAnimalSkill
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Skill = append(m.Skill, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Skill", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAnimalSkill
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAnimalSkill
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAnimalSkill
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAnimalSkill(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAnimalSkill
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
func skipAnimalSkill(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAnimalSkill
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
					return 0, ErrIntOverflowAnimalSkill
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
					return 0, ErrIntOverflowAnimalSkill
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
				return 0, ErrInvalidLengthAnimalSkill
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAnimalSkill
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAnimalSkill
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAnimalSkill        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAnimalSkill          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAnimalSkill = fmt.Errorf("proto: unexpected end of group")
)