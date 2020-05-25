package proto

import "io"

func UnmarshalBytes(b *Buffer, wt WireType) ([]byte, error) {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(size) {
			return nil, io.ErrUnexpectedEOF
		}

		return b.ConsumeBytes(int(size)), nil

	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalBoolSlice(b *Buffer, wt WireType) ([]bool, error) {
	switch wt {
	case WireBytes:
		count, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(int(count)) {
			return nil, io.ErrUnexpectedEOF
		}

		var ret []bool
		for _, element := range b.ConsumeBytes(int(count)) {

			switch element {
			case 0:
				ret = append(ret, false)
			case 1:
				ret = append(ret, true)
			default:
				return nil, ErrBadBoolValue
			}
		}

		return ret, nil

	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalInt32Slice(b *Buffer, wt WireType) ([]int32, error) {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(size) {
			return nil, io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		var ret []int32
		for limitBuffer.BytesRemains() > 0 {
			var element int32
			element, err = UnmarshalInt32(limitBuffer, WireVarint)
			if err != nil {
				return nil, err
			}

			ret = append(ret, element)
		}

		return ret, nil

	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalUInt32Slice(b *Buffer, wt WireType) ([]uint32, error) {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(size) {
			return nil, io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		var ret []uint32
		for limitBuffer.BytesRemains() > 0 {
			var element uint32
			element, err = UnmarshalUInt32(limitBuffer, WireVarint)
			if err != nil {
				return nil, err
			}

			ret = append(ret, element)
		}

		return ret, nil

	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalInt64Slice(b *Buffer, wt WireType) ([]int64, error) {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(size) {
			return nil, io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		var ret []int64
		for limitBuffer.BytesRemains() > 0 {
			var element int64
			element, err = UnmarshalInt64(limitBuffer, WireVarint)
			if err != nil {
				return nil, err
			}

			ret = append(ret, element)
		}
		return ret, nil

	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalUInt64Slice(b *Buffer, wt WireType) ([]uint64, error) {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(size) {
			return nil, io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		var ret []uint64
		for limitBuffer.BytesRemains() > 0 {
			var element uint64
			element, err = UnmarshalUInt64(limitBuffer, WireVarint)
			if err != nil {
				return nil, err
			}

			ret = append(ret, element)
		}

		return ret, nil

	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalStringSlice(b *Buffer, wt WireType) ([]string, error) {
	switch wt {
	case WireBytes:
		v, err := b.DecodeStringBytes()
		if err != nil {
			return nil, err
		}

		return []string{v}, nil
	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalFloat32Slice(b *Buffer, wt WireType) ([]float32, error) {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(size) {
			return nil, io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		var ret []float32
		for limitBuffer.BytesRemains() > 0 {
			var element float32
			element, err = UnmarshalFloat32(limitBuffer, WireFixed32)
			if err != nil {
				return nil, err
			}

			ret = append(ret, element)
		}

		return ret, nil

	default:
		return nil, ErrBadWireType
	}
}

func UnmarshalFloat64Slice(b *Buffer, wt WireType) ([]float64, error) {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return nil, err
		}

		if b.BytesRemains() < int(size) {
			return nil, io.ErrUnexpectedEOF
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		var ret []float64
		for limitBuffer.BytesRemains() > 0 {
			var element float64
			element, err = UnmarshalFloat64(limitBuffer, WireFixed64)
			if err != nil {
				return nil, err
			}

			ret = append(ret, element)
		}

		return ret, nil

	default:
		return nil, ErrBadWireType
	}
}
