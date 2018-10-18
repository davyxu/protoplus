package proto

func UnmarshalBoolSlice(b *Buffer, wt WireType, ret *[]bool) error {
	switch wt {
	case WireBytes:
		count, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(count)))

		s := make([]bool, count)
		var i uint64
		for i = 0; i < count; i++ {

			if limitBuffer.buf[i] != 0 {
				s[i] = true
			}
		}

		*ret = s

	default:
		return ErrBadWireType
	}

	return nil
}
