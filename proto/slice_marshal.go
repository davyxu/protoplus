package proto

func MarshalBoolSlice(b *Buffer, fieldIndex uint64, value []bool) error {

	count := len(value)
	if count == 0 {
		return nil
	}

	b.EncodeVarint(makeWireTag(fieldIndex, WireBytes))
	b.EncodeVarint(uint64(count))

	for _, v := range value {
		if v {
			b.buf = append(b.buf, 1)
		} else {
			b.buf = append(b.buf, 0)
		}
	}

	return nil
}
