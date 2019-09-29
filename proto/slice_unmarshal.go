package proto

import "io"

func UnmarshalBytes(b *Buffer, wt WireType, ret *[]byte) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		*ret = append(*ret, b.ConsumeBytes(int(size))...)

	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalBoolSlice(b *Buffer, wt WireType, ret *[]bool) error {
	switch wt {
	case WireBytes:
		count, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(int(count)) {
			return io.ErrUnexpectedEOF
		}

		for _, element := range b.ConsumeBytes(int(count)) {

			switch element {
			case 0:
				*ret = append(*ret, false)
			case 1:
				*ret = append(*ret, true)
			default:
				return ErrBadBoolValue
			}

		}

	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalInt32Slice(b *Buffer, wt WireType, ret *[]int32) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		for limitBuffer.BytesRemains() > 0 {
			var element int32
			err = UnmarshalInt32(limitBuffer, WireVarint, &element)
			if err != nil {
				return err
			}

			*ret = append(*ret, element)
		}

	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalUInt32Slice(b *Buffer, wt WireType, ret *[]uint32) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		for limitBuffer.BytesRemains() > 0 {
			var element uint32
			err = UnmarshalUInt32(limitBuffer, WireVarint, &element)
			if err != nil {
				return err
			}

			*ret = append(*ret, element)
		}

	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalInt64Slice(b *Buffer, wt WireType, ret *[]int64) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		for limitBuffer.BytesRemains() > 0 {
			var element int64
			err = UnmarshalInt64(limitBuffer, WireVarint, &element)
			if err != nil {
				return err
			}

			*ret = append(*ret, element)
		}

	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalUInt64Slice(b *Buffer, wt WireType, ret *[]uint64) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		for limitBuffer.BytesRemains() > 0 {
			var element uint64
			err = UnmarshalUInt64(limitBuffer, WireVarint, &element)
			if err != nil {
				return err
			}

			*ret = append(*ret, element)
		}

	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalStringSlice(b *Buffer, wt WireType, ret *[]string) error {
	switch wt {
	case WireBytes:
		v, err := b.DecodeStringBytes()
		if err != nil {
			return err
		}

		*ret = append(*ret, v)
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalFloat32Slice(b *Buffer, wt WireType, ret *[]float32) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		for limitBuffer.BytesRemains() > 0 {
			var element float32
			err = UnmarshalFloat32(limitBuffer, WireFixed32, &element)
			if err != nil {
				return err
			}

			*ret = append(*ret, element)
		}

	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalFloat64Slice(b *Buffer, wt WireType, ret *[]float64) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		for limitBuffer.BytesRemains() > 0 {
			var element float64
			err = UnmarshalFloat64(limitBuffer, WireFixed64, &element)
			if err != nil {
				return err
			}

			*ret = append(*ret, element)
		}

	default:
		return ErrBadWireType
	}

	return nil
}
