package proto

import (
	"math"
	"reflect"
)

func SizeBool(fieldIndex uint64, value bool) (ret int) {

	if value == false {
		return 0
	}

	if fieldIndex != math.MaxUint64 {
		ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
	}

	ret += 1

	return
}

func SizeInt32(fieldIndex uint64, value int32) (ret int) {

	switch {
	case value > 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
		}

		ret += SizeVarint(uint64(value))

	case value < 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireZigzag32))
		}

		ret += SizeVarint(Zigzag32(uint64(value)))
	default:
		if fieldIndex == math.MaxUint64 {
			ret += 1
		}
	}

	return
}

func SizeUInt32(fieldIndex uint64, value uint32) (ret int) {

	switch {
	case value > 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
		}

		ret += SizeVarint(uint64(value))

	case value < 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireZigzag32))
		}

		ret += SizeVarint(Zigzag32(uint64(value)))
	default:
		if fieldIndex == math.MaxUint64 {
			ret += 1
		}
	}

	return
}

func SizeInt64(fieldIndex uint64, value int64) (ret int) {

	switch {
	case value > 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
		}

		ret += SizeVarint(uint64(value))

	case value < 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireZigzag64))
		}

		ret += SizeVarint(Zigzag64(uint64(value)))
	default:
		if fieldIndex == math.MaxUint64 {
			ret += 1
		}
	}

	return
}

func SizeUInt64(fieldIndex uint64, value uint64) (ret int) {

	switch {
	case value > 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
		}

		ret += SizeVarint(uint64(value))

	case value < 0:

		if fieldIndex != math.MaxUint64 {
			ret = SizeVarint(makeWireTag(fieldIndex, WireZigzag64))
		}

		ret += SizeVarint(Zigzag64(uint64(value)))
	default:
		if fieldIndex == math.MaxUint64 {
			ret += 1
		}
	}

	return

}

func SizeFloat32(fieldIndex uint64, value float32) (ret int) {

	if value == 0 {
		return 0
	}

	if fieldIndex != math.MaxUint64 {
		ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
	}

	ret += 4

	return
}

func SizeFloat64(fieldIndex uint64, value float64) (ret int) {

	if value == 0 {
		return 0
	}

	if fieldIndex != math.MaxUint64 {
		ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
	}

	ret += 8

	return
}

func SizeString(fieldIndex uint64, value string) (ret int) {

	size := len(value)

	if size == 0 {
		return 0
	}

	if fieldIndex != math.MaxUint64 {
		ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
	}

	ret += SizeVarint(uint64(size)) + size

	return
}

func SizeStruct(fieldIndex uint64, msg Struct) (ret int) {

	structValue := reflect.ValueOf(msg)

	// *MyType被Message包裹后，判断不为nil
	if structValue.IsNil() {
		return 0
	}

	size := msg.Size()

	if size == 0 {
		return 0
	}

	if fieldIndex != math.MaxUint64 {
		ret = SizeVarint(makeWireTag(fieldIndex, WireVarint))
	}

	ret += SizeVarint(uint64(size)) + size

	return
}
