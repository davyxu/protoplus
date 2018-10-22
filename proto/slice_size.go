package proto

import "reflect"

func SizeBoolSlice(fieldIndex uint64, value []bool) (ret int) {

	count := len(value)
	if count == 0 {
		return 0
	}

	ret = SizeVarint(makeWireTag(fieldIndex, WireBytes))

	size := count * 1

	ret += SizeVarint(uint64(size))

	ret += size

	return
}

func SizeInt32Slice(fieldIndex uint64, value []int32) (ret int) {

	count := len(value)
	if count == 0 {
		return 0
	}

	ret = SizeVarint(makeWireTag(fieldIndex, WireBytes))

	// 后部分的长度
	size := 0
	for _, v := range value {
		size += SizeVarint(uint64(v))
	}

	ret += SizeVarint(uint64(size))

	ret += size

	return
}

func SizeStringSlice(fieldIndex uint64, value []string) (ret int) {

	if len(value) == 0 {
		return 0
	}

	for _, v := range value {
		ret += SizeString(fieldIndex, v)
	}

	return
}

func SizeFloat32Slice(fieldIndex uint64, value []float32) (ret int) {

	count := len(value)
	if count == 0 {
		return 0
	}

	ret = SizeVarint(makeWireTag(fieldIndex, WireBytes))

	size := count * 4

	ret += SizeVarint(uint64(size))

	ret += size

	return
}

func SizeStructSlice(fieldIndex uint64, value interface{}) (ret int) {

	slice := reflect.ValueOf(value)
	count := slice.Len()

	if count == 0 {
		return 0
	}

	for i := 0; i < count; i++ {
		v := slice.Index(i)
		ret += SizeStruct(fieldIndex, v.Interface().(Struct))
	}

	return
}
