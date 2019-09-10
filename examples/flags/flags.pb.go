// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flags.proto

package flavortown_flags

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	fmt "fmt"
	_ "github.com/bi-foundation/protobuf-graphql-extension/graphqlproto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_graphql_go_graphql "github.com/graphql-go/graphql"
	io "io"
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

type User struct {
	ScopeA               bool     `protobuf:"varint,1,opt,name=scopeA,proto3" json:"scopeA,omitempty"`
	ScopeB               bool     `protobuf:"varint,2,opt,name=scopeB,proto3" json:"scopeB,omitempty"`
	ScopeC               bool     `protobuf:"varint,3,opt,name=scopeC,proto3" json:"scopeC,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_298e9b561aa019b2, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetScopeA() bool {
	if m != nil {
		return m.ScopeA
	}
	return false
}

func (m *User) GetScopeB() bool {
	if m != nil {
		return m.ScopeB
	}
	return false
}

func (m *User) GetScopeC() bool {
	if m != nil {
		return m.ScopeC
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "flavortown.flags.User")
}

func init() { proto.RegisterFile("flags.proto", fileDescriptor_298e9b561aa019b2) }

var fileDescriptor_298e9b561aa019b2 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcb, 0x49, 0x4c,
	0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x48, 0xcb, 0x49, 0x2c, 0xcb, 0x2f, 0x2a,
	0xc9, 0x2f, 0xcf, 0xd3, 0x03, 0x8b, 0x4b, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x15, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60,
	0x0e, 0x98, 0x05, 0x31, 0x40, 0xca, 0x00, 0x49, 0x79, 0x7e, 0x41, 0x71, 0x6a, 0x2a, 0x42, 0x3d,
	0x98, 0x0b, 0xd1, 0x00, 0x66, 0x42, 0x74, 0x28, 0x85, 0x70, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09,
	0x89, 0x71, 0xb1, 0x15, 0x27, 0xe7, 0x17, 0xa4, 0x3a, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04,
	0x41, 0x79, 0x70, 0x71, 0x27, 0x09, 0x26, 0x24, 0x71, 0x27, 0xb8, 0xb8, 0xb3, 0x04, 0x33, 0x92,
	0xb8, 0xb3, 0x15, 0xcb, 0x89, 0x45, 0xf2, 0x8c, 0x4e, 0x32, 0x3f, 0x1e, 0xca, 0x31, 0xae, 0x78,
	0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x78, 0x60, 0x91, 0x3c, 0x63, 0x12, 0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x19, 0x99, 0xb0, 0xf8, 0xfc, 0x00, 0x00, 0x00,
}

func (this *User) UInt64() uint64 {
	b := uint64(0)
	if this.ScopeA {
		b |= uint64(1) << uint64(0)
	}
	if this.ScopeB {
		b |= uint64(1) << uint64(1)
	}
	if this.ScopeC {
		b |= uint64(1) << uint64(2)
	}

	return b
}
func (this *User) HighFlags() []string {
	var b []string
	if this.ScopeA {
		b = append(b, "scope_a")
	}
	if this.ScopeB {
		b = append(b, "scope_b")
	}
	if this.ScopeC {
		b = append(b, "scope_c")
	}
	return b
}

func (this *User) LowFlags() []string {
	var b []string
	if !this.ScopeA {
		b = append(b, "scope_a")
	}
	if !this.ScopeB {
		b = append(b, "scope_b")
	}
	if !this.ScopeC {
		b = append(b, "scope_c")
	}
	return b
}

func (this *User) SetFlag(flag string) error {
	switch flag {
	case "scope_a":
		this.ScopeA = true
	case "scope_b":
		this.ScopeB = true
	case "scope_c":
		this.ScopeC = true
	default:
		return fmt.Errorf("invalid flag: %v", flag)
	}
	return nil
}
func (this *User) ClearFlag(flag string) error {
	switch flag {
	case "scope_a":
		this.ScopeA = false
	case "scope_b":
		this.ScopeB = false
	case "scope_c":
		this.ScopeC = false
	default:
		return fmt.Errorf("invalid flag: %v", flag)
	}
	return nil
}
func (this *User) SetFlags(flags ...string) []error {
	var errs []error
	for _, f := range flags {
		if err := this.SetFlag(f); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
func (this *User) ClearFlags(flags ...string) []error {
	var errs []error
	for _, f := range flags {
		if err := this.ClearFlag(f); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
func (this *User) TestFlag(flag string) bool {
	switch flag {
	case "scope_a":
		return this.ScopeA
	case "scope_b":
		return this.ScopeB
	case "scope_c":
		return this.ScopeC
	}
	return false
}
func (this *User) TestFlags(flags ...string) bool {
	for _, f := range flags {
		if !this.TestFlag(f) {
			return false
		}
	}
	return true
}
func (this *User) FromUInt64(b uint64) error {
	bb := b
	bb = b
	if bb&(uint64(1)<<uint(0)) > 0 {
		this.ScopeA = true
	} else {
		this.ScopeA = false
	}
	bb = b
	if bb&(uint64(1)<<uint(1)) > 0 {
		this.ScopeB = true
	} else {
		this.ScopeB = false
	}
	bb = b
	if bb&(uint64(1)<<uint(2)) > 0 {
		this.ScopeC = true
	} else {
		this.ScopeC = false
	}

	return nil
}
func (this *User) Scan(i interface{}) error {
	switch v := i.(type) {
	case int:
		return this.FromUInt64(uint64(v))
	case int32:
		return this.FromUInt64(uint64(v))
	case int64:
		return this.FromUInt64(uint64(v))
	case float32:
		return this.FromUInt64(uint64(v))
	case float64:
		return this.FromUInt64(uint64(v))
	}

	return fmt.Errorf("invalid type: %T", i)
}
func (this *User) Value() (database_sql_driver.Value, error) {
	return int64(this.UInt64()), nil
}
func (this *User) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*User)
	if !ok {
		that2, ok := that.(User)
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
	if this.ScopeA != that1.ScopeA {
		return false
	}
	if this.ScopeB != that1.ScopeB {
		return false
	}
	if this.ScopeC != that1.ScopeC {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type UserGetter interface {
	GetUser() *User
}

var GraphQLUserType *github_com_graphql_go_graphql.Object

func init() {
	GraphQLUserType = github_com_graphql_go_graphql.NewObject(github_com_graphql_go_graphql.ObjectConfig{
		Name:        "User",
		Description: "",
		Fields: (github_com_graphql_go_graphql.FieldsThunk)(func() github_com_graphql_go_graphql.Fields {
			return github_com_graphql_go_graphql.Fields{
				"scopeA": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*User)
						if ok {
							return obj.ScopeA, nil
						}
						inter, ok := p.Source.(UserGetter)
						if ok {
							face := inter.GetUser()
							if face == nil {
								return nil, nil
							}
							return face.ScopeA, nil
						}
						return nil, fmt.Errorf("field scopeA not resolved")
					},
				},
				"scopeB": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*User)
						if ok {
							return obj.ScopeB, nil
						}
						inter, ok := p.Source.(UserGetter)
						if ok {
							face := inter.GetUser()
							if face == nil {
								return nil, nil
							}
							return face.ScopeB, nil
						}
						return nil, fmt.Errorf("field scopeB not resolved")
					},
				},
				"scopeC": &github_com_graphql_go_graphql.Field{
					Type:        github_com_graphql_go_graphql.Boolean,
					Description: "",
					Resolve: func(p github_com_graphql_go_graphql.ResolveParams) (interface{}, error) {
						obj, ok := p.Source.(*User)
						if ok {
							return obj.ScopeC, nil
						}
						inter, ok := p.Source.(UserGetter)
						if ok {
							face := inter.GetUser()
							if face == nil {
								return nil, nil
							}
							return face.ScopeC, nil
						}
						return nil, fmt.Errorf("field scopeC not resolved")
					},
				},
			}
		}),
	})
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ScopeA {
		dAtA[i] = 0x8
		i++
		if m.ScopeA {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ScopeB {
		dAtA[i] = 0x10
		i++
		if m.ScopeB {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ScopeC {
		dAtA[i] = 0x18
		i++
		if m.ScopeC {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFlags(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedUser(r randyFlags, easy bool) *User {
	this := &User{}
	this.ScopeA = bool(bool(r.Intn(2) == 0))
	this.ScopeB = bool(bool(r.Intn(2) == 0))
	this.ScopeC = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedFlags(r, 4)
	}
	return this
}

type randyFlags interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneFlags(r randyFlags) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringFlags(r randyFlags) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneFlags(r)
	}
	return string(tmps)
}
func randUnrecognizedFlags(r randyFlags, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldFlags(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldFlags(dAtA []byte, r randyFlags, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateFlags(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateFlags(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateFlags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateFlags(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateFlags(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateFlags(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateFlags(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ScopeA {
		n += 2
	}
	if m.ScopeB {
		n += 2
	}
	if m.ScopeC {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFlags(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFlags(x uint64) (n int) {
	return sovFlags(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlags
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopeA", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ScopeA = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopeB", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ScopeB = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopeC", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ScopeC = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlags
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFlags
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFlags(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlags
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
					return 0, ErrIntOverflowFlags
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFlags
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
				return 0, ErrInvalidLengthFlags
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFlags
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFlags
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFlags(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFlags
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFlags = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlags   = fmt.Errorf("proto: integer overflow")
)