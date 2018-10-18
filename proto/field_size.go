package proto

import "reflect"

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
		return SizeVarint(makeWireTag(fieldIndex, WireZigzag32)) + SizeVarint(Zigzag32(uint64(value)))
	}

	return 0
}

func SizeUInt32(fieldIndex uint64, value uint32) int {

	switch {
	case value > 0:
		return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(value))
	case value < 0:
		return SizeVarint(makeWireTag(fieldIndex, WireZigzag32)) + SizeVarint(Zigzag32(uint64(value)))
	}

	return 0
}

func SizeInt64(fieldIndex uint64, value int64) int {

	switch {
	case value > 0:
		return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(value))
	case value < 0:
		return SizeVarint(makeWireTag(fieldIndex, WireZigzag64)) + SizeVarint(Zigzag64(uint64(value)))
	}

	return 0
}

func SizeUInt64(fieldIndex uint64, value uint64) int {

	switch {
	case value > 0:
		return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(value))
	case value < 0:
		return SizeVarint(makeWireTag(fieldIndex, WireZigzag64)) + SizeVarint(Zigzag64(value))
	}

	return 0

}

func SizeFloat32(fieldIndex uint64, value float32) int {

	if value == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + 4
}

func SizeFloat64(fieldIndex uint64, value float64) int {

	if value == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + 8
}

func SizeString(fieldIndex uint64, value string) int {

	l := len(value)

	if l == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(l)) + l
}

func SizeStruct(fieldIndex uint64, msg Struct) int {

	structValue := reflect.ValueOf(msg)

	// *MyType被Message包裹后，判断不为nil
	if structValue.IsNil() {
		return 0
	}

	l := msg.Size()

	if l == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireVarint)) + SizeVarint(uint64(l)) + l
}
