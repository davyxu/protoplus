// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
package tests

import (
	"github.com/davyxu/protoplus/proto"
	"fmt"
)

var (
	_ *proto.Buffer
	_ fmt.Stringer
)

type MyTypeMini struct {
	Bool bool
}

func (self *MyTypeMini) String() string { return fmt.Sprintf("%+v", *self) }

func (self *MyTypeMini) Size() (ret int) {

	ret += proto.SizeBool(0, self.Bool)

	return
}

func (self *MyTypeMini) Marshal(buffer *proto.Buffer) error {

	proto.MarshalBool(buffer, 0, self.Bool)

	return nil
}

func (self *MyTypeMini) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalBool(buffer, wt, &self.Bool)

	}

	return proto.ErrUnknownField
}

type MyType struct {
	Bool         bool
	Int32        int32
	UInt32       uint32
	Int64        int64
	UInt64       uint64
	Float32      float32
	Float64      float64
	Str          string
	Struct       *MyType
	BytesSlice   []byte
	BoolSlice    []bool
	Int32Slice   []int32
	UInt32Slice  []uint32
	Int64Slice   []int64
	UInt64Slice  []uint64
	Float32Slice []float32
	Float64Slice []float64
	StrSlice     []string
	StructSlice  []*MyType
}

func (self *MyType) String() string { return fmt.Sprintf("%+v", *self) }

func (self *MyType) Size() (ret int) {

	ret += proto.SizeBool(0, self.Bool)

	ret += proto.SizeInt32(1, self.Int32)

	ret += proto.SizeUInt32(2, self.UInt32)

	ret += proto.SizeInt64(3, self.Int64)

	ret += proto.SizeUInt64(4, self.UInt64)

	ret += proto.SizeFloat32(5, self.Float32)

	ret += proto.SizeFloat64(6, self.Float64)

	ret += proto.SizeString(7, self.Str)

	ret += proto.SizeStruct(8, self.Struct)

	ret += proto.SizeBytes(9, self.BytesSlice)

	ret += proto.SizeBoolSlice(10, self.BoolSlice)

	ret += proto.SizeInt32Slice(11, self.Int32Slice)

	ret += proto.SizeUInt32Slice(12, self.UInt32Slice)

	ret += proto.SizeInt64Slice(13, self.Int64Slice)

	ret += proto.SizeUInt64Slice(14, self.UInt64Slice)

	ret += proto.SizeFloat32Slice(15, self.Float32Slice)

	ret += proto.SizeFloat64Slice(16, self.Float64Slice)

	ret += proto.SizeStringSlice(17, self.StrSlice)

	if len(self.StructSlice) > 0 {
		for _, elm := range self.StructSlice {
			ret += proto.SizeStruct(18, elm)
		}
	}

	return
}

func (self *MyType) Marshal(buffer *proto.Buffer) error {

	proto.MarshalBool(buffer, 0, self.Bool)

	proto.MarshalInt32(buffer, 1, self.Int32)

	proto.MarshalUInt32(buffer, 2, self.UInt32)

	proto.MarshalInt64(buffer, 3, self.Int64)

	proto.MarshalUInt64(buffer, 4, self.UInt64)

	proto.MarshalFloat32(buffer, 5, self.Float32)

	proto.MarshalFloat64(buffer, 6, self.Float64)

	proto.MarshalString(buffer, 7, self.Str)

	proto.MarshalStruct(buffer, 8, self.Struct)

	proto.MarshalBytes(buffer, 9, self.BytesSlice)

	proto.MarshalBoolSlice(buffer, 10, self.BoolSlice)

	proto.MarshalInt32Slice(buffer, 11, self.Int32Slice)

	proto.MarshalUInt32Slice(buffer, 12, self.UInt32Slice)

	proto.MarshalInt64Slice(buffer, 13, self.Int64Slice)

	proto.MarshalUInt64Slice(buffer, 14, self.UInt64Slice)

	proto.MarshalFloat32Slice(buffer, 15, self.Float32Slice)

	proto.MarshalFloat64Slice(buffer, 16, self.Float64Slice)

	proto.MarshalStringSlice(buffer, 17, self.StrSlice)

	for _, elm := range self.StructSlice {
		proto.MarshalStruct(buffer, 18, elm)
	}

	return nil
}

func (self *MyType) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalBool(buffer, wt, &self.Bool)
	case 1:
		return proto.UnmarshalInt32(buffer, wt, &self.Int32)
	case 2:
		return proto.UnmarshalUInt32(buffer, wt, &self.UInt32)
	case 3:
		return proto.UnmarshalInt64(buffer, wt, &self.Int64)
	case 4:
		return proto.UnmarshalUInt64(buffer, wt, &self.UInt64)
	case 5:
		return proto.UnmarshalFloat32(buffer, wt, &self.Float32)
	case 6:
		return proto.UnmarshalFloat64(buffer, wt, &self.Float64)
	case 7:
		return proto.UnmarshalString(buffer, wt, &self.Str)
	case 8:
		self.Struct = new(MyType)
		return proto.UnmarshalStruct(buffer, wt, self.Struct)
	case 9:
		return proto.UnmarshalBytes(buffer, wt, &self.BytesSlice)
	case 10:
		return proto.UnmarshalBoolSlice(buffer, wt, &self.BoolSlice)
	case 11:
		return proto.UnmarshalInt32Slice(buffer, wt, &self.Int32Slice)
	case 12:
		return proto.UnmarshalUInt32Slice(buffer, wt, &self.UInt32Slice)
	case 13:
		return proto.UnmarshalInt64Slice(buffer, wt, &self.Int64Slice)
	case 14:
		return proto.UnmarshalUInt64Slice(buffer, wt, &self.UInt64Slice)
	case 15:
		return proto.UnmarshalFloat32Slice(buffer, wt, &self.Float32Slice)
	case 16:
		return proto.UnmarshalFloat64Slice(buffer, wt, &self.Float64Slice)
	case 17:
		return proto.UnmarshalStringSlice(buffer, wt, &self.StrSlice)
	case 18:
		elm := new(MyType)
		if err := proto.UnmarshalStruct(buffer, wt, elm); err != nil {
			return err
		} else {
			self.StructSlice = append(self.StructSlice, elm)
			return nil
		}

	}

	return proto.ErrUnknownField
}