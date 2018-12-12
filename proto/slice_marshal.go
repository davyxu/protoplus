package proto

import (
	"math"
)

func MarshalBytes(b *Buffer, fieldIndex uint64, value []byte) error {

	size := len(value)
	if size == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))

	// 因为bool每一个value都是1个字节，size=1*count
	b.EncodeVarint(uint64(size))
	b.buf = append(b.buf, value...)

	return nil
}

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
		b.EncodeVarint(uint64(v)) // TODO 负数会导致编码很大
	}

	return nil
}

func MarshalUInt32Slice(b *Buffer, fieldIndex uint64, value []uint32) error {

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

func MarshalInt64Slice(b *Buffer, fieldIndex uint64, value []int64) error {

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

func MarshalUInt64Slice(b *Buffer, fieldIndex uint64, value []uint64) error {

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
		b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))
		b.EncodeStringBytes(v)
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

func MarshalFloat64Slice(b *Buffer, fieldIndex uint64, value []float64) error {

	count := len(value)
	if count == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))

	// 写入长度
	b.EncodeVarint(uint64(count * 8))

	for _, v := range value {
		b.EncodeFixed64(uint64(math.Float64bits(v)))
	}

	return nil
}
