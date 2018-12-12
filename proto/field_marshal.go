package proto

import (
	"math"
	"reflect"
)

func MarshalBool(b *Buffer, fieldIndex uint64, value bool) error {

	if !value {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireVarint))

	if value {
		b.buf = append(b.buf, 1)
	} else {
		b.buf = append(b.buf, 0)
	}

	return nil
}

func MarshalInt32(b *Buffer, fieldIndex uint64, value int32) error {

	switch {
	case value > 0:
		b.EncodeVarint(makeWireTag(fieldIndex, WireVarint))
		b.EncodeVarint(uint64(value))
	case value < 0:
		b.EncodeVarint(makeWireTag(fieldIndex, WireFixed32))
		b.EncodeFixed32(uint64(value))
	}

	return nil
}

func MarshalUInt32(b *Buffer, fieldIndex uint64, value uint32) error {

	if value == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireVarint))
	b.EncodeVarint(uint64(value))

	return nil
}

func MarshalInt64(b *Buffer, fieldIndex uint64, value int64) error {

	switch {
	case value > 0:
		b.EncodeVarint(makeWireTag(fieldIndex, WireVarint))
		b.EncodeVarint(uint64(value))
	case value < 0:
		b.EncodeVarint(makeWireTag(fieldIndex, WireFixed64))
		b.EncodeFixed64(uint64(value))
	}

	return nil
}

func MarshalUInt64(b *Buffer, fieldIndex uint64, value uint64) error {

	if value == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireVarint))
	b.EncodeVarint(value)

	return nil
}

func MarshalFloat32(b *Buffer, fieldIndex uint64, value float32) error {

	if value == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireFixed32))
	b.EncodeFixed32(uint64(math.Float32bits(value)))

	return nil
}

func MarshalFloat64(b *Buffer, fieldIndex uint64, value float64) error {

	if value == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireFixed64))
	b.EncodeFixed64(uint64(math.Float64bits(value)))

	return nil
}

func MarshalString(b *Buffer, fieldIndex uint64, value string) error {

	if value == "" {
		return nil
	}
	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))

	b.EncodeStringBytes(value)

	return nil
}

func MarshalStruct(b *Buffer, fieldIndex uint64, msg Struct) error {

	structValue := reflect.ValueOf(msg)

	// *MyType被Message包裹后，判断不为nil
	if structValue.IsNil() {
		return nil
	}

	size := msg.Size()
	if size == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))

	b.EncodeVarint(uint64(size))

	return msg.Marshal(b)
}
