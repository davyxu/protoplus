package proto

import (
	"math"
	"reflect"
)

func MarshalBoolSlice(b *Buffer, fieldIndex uint64, value []bool) error {

	size := len(value)
	if size == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))

	// 因为bool每一个value都是1个字节，size=1*count
	b.EncodeVarint(uint64(size))

	for _, v := range value {

		if v {
			b.buf = append(b.buf, 1)
		} else {
			b.buf = append(b.buf, 0)
		}
	}

	return nil
}

func MarshalInt32Slice(b *Buffer, fieldIndex uint64, value []int32) error {

	if len(value) == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))

	// 计算变长单元大小
	var size int
	for _, v := range value {
		size += SizeVarint(uint64(v))
	}

	// 写入长度
	b.EncodeVarint(uint64(size))
	for _, v := range value {
		b.EncodeVarint(uint64(v))
	}

	return nil
}

func MarshalStringSlice(b *Buffer, fieldIndex uint64, value []string) error {

	if len(value) == 0 {
		return nil
	}

	for _, v := range value {
		MarshalString(b, fieldIndex, v)
	}

	return nil
}

func MarshalFloat32Slice(b *Buffer, fieldIndex uint64, value []float32) error {

	count := len(value)
	if count == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))

	// 写入长度
	b.EncodeVarint(uint64(count * 4))

	for _, v := range value {
		b.EncodeFixed32(uint64(math.Float32bits(v)))
	}

	return nil
}

func MarshalStructSlice(b *Buffer, fieldIndex uint64, raw interface{}) error {

	slice := reflect.ValueOf(raw)
	count := slice.Len()

	if count == 0 {
		return nil
	}

	for i := 0; i < count; i++ {
		v := slice.Index(i)

		MarshalStruct(b, fieldIndex, v.Interface().(Struct))
	}
	return nil
}
