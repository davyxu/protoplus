package proto

func SizeBoolSlice(fieldIndex uint64, value []bool) int {

	count := len(value)
	if count == 0 {
		return 0
	}

	return SizeVarint(makeWireTag(fieldIndex, WireBytes)) + SizeVarint(uint64(count)) + count
}
