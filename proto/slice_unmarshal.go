package proto

import "reflect"

func UnmarshalBoolSlice(b *Buffer, wt WireType, ret *[]bool) error {
	switch wt {
	case WireBytes:
		count, err := b.DecodeVarint()
		if err != nil {
			return err
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

func UnmarshalStructSlice(b *Buffer, wt WireType, ret interface{}) error {

	switch wt {
	case WireBytes:
		size, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		limitBuffer := NewBuffer(b.ConsumeBytes(int(size)))

		sliceType := reflect.TypeOf(ret).Elem()

		elementType := sliceType.Elem().Elem()

		sliceValue := reflect.ValueOf(ret).Elem()

		// msgIns:  *MyType
		msgIns := reflect.New(elementType)

		err = rawUnmarshalStruct(limitBuffer, msgIns.Interface().(Struct))
		if err != nil {
			return err
		}

		sliceValue.Set(reflect.Append(sliceValue, msgIns))

	default:
		return ErrBadWireType
	}

	return nil
}
