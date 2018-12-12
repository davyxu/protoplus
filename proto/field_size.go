package proto

import (
	"reflect"
)

func SizeBool(fieldIndex uint64, value bool) int {

	if value == false {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + 1
}

func SizeInt32(fieldIndex uint64, value int32) int {

	switch {
	case value > 0:
		return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(value))
	case value < 0:
		return SizeVarint(makeWireTag(fieldIndex, WireFixed32)) + 4
	default:
		return 0
	}
}

func SizeUInt32(fieldIndex uint64, value uint32) int {

	if value == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(value))
}

func SizeInt64(fieldIndex uint64, value int64) int {

	switch {
	case value > 0:
		return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(value))
	case value < 0:
		return SizeVarint(makeWireTag(fieldIndex, WireFixed64)) + 8
	default:
		return 0
	}
}

func SizeUInt64(fieldIndex uint64, value uint64) int {

	if value == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(value))
}

func SizeFloat32(fieldIndex uint64, value float32) int {

	if value == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireFixed32)) + 4
}

func SizeFloat64(fieldIndex uint64, value float64) int {

	if value == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireFixed64)) + 8
}

func SizeString(fieldIndex uint64, value string) int {

	size := len(value)

	if size == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(size)) + size
}

func SizeStruct(fieldIndex uint64, msg Struct) int {

	structValue := reflect.ValueOf(msg)

	// *MyType被Message包裹后，判断不为nil
	if structValue.IsNil() {
		return 0
	}

	size := msg.Size()

	if size == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(size)) + size
}
