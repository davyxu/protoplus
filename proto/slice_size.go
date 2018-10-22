package proto

func SizeBytes(fieldIndex uint64, value []byte) int {

	count := len(value)
	if count == 0 {
		return 0
	}

	size := count * 1

	return SizeVarint(makeWireTag(fieldIndex, WireBytes)) + SizeVarint(uint64(size)) + size
}

func SizeBoolSlice(fieldIndex uint64, value []bool) int {

	count := len(value)
	if count == 0 {
		return 0
	}

	size := count * 1

	return SizeVarint(makeWireTag(fieldIndex, WireBytes)) + SizeVarint(uint64(size)) + size
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

func SizeUInt32Slice(fieldIndex uint64, value []uint32) (ret int) {

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

func SizeInt64Slice(fieldIndex uint64, value []int64) (ret int) {

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

func SizeUInt64Slice(fieldIndex uint64, value []uint64) (ret int) {

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

func SizeFloat64Slice(fieldIndex uint64, value []float64) (ret int) {

	count := len(value)
	if count == 0 {
		return 0
	}

	ret = SizeVarint(makeWireTag(fieldIndex, WireBytes))

	size := count * 8

	ret += SizeVarint(uint64(size))

	ret += size

	return
}
