package proto

import (
	"io"
	"math"
)

func UnmarshalBool(b *Buffer, wt WireType, ret *bool) error {
	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return err
		}

		if v != 0 {
			*ret = true
		} else {
			*ret = false
		}
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalInt32(b *Buffer, wt WireType, ret *int32) error {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return err
		}
		*ret = int32(v)

	case WireFixed32:
		v, err := b.DecodeFixed32()
		if err != nil {
			break
		}

		*ret = int32(v)
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalUInt32(b *Buffer, wt WireType, ret *uint32) error {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return err
		}
		*ret = uint32(v)
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalInt64(b *Buffer, wt WireType, ret *int64) error {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return err
		}
		*ret = int64(v)

	case WireFixed64:
		v, err := b.DecodeFixed64()
		if err != nil {
			break
		}

		*ret = int64(v)
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalUInt64(b *Buffer, wt WireType, ret *uint64) error {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return err
		}
		*ret = v
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalFloat32(b *Buffer, wt WireType, ret *float32) error {

	switch wt {
	case WireFixed32:
		v, err := b.DecodeFixed32()
		if err != nil {
			return err
		}

		*ret = math.Float32frombits(uint32(v))
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalFloat64(b *Buffer, wt WireType, ret *float64) error {

	switch wt {
	case WireFixed64:
		v, err := b.DecodeFixed64()
		if err != nil {
			return err
		}

		*ret = math.Float64frombits(uint64(v))
	default:
		return ErrBadWireType
	}

	return nil
}

func UnmarshalString(b *Buffer, wt WireType, ret *string) error {
	switch wt {
	case WireBytes:
		v, err := b.DecodeStringBytes()
		if err != nil {
			return err
		}

		*ret = v
	default:
		return ErrBadWireType
	}

	return nil
}

func skipField(b *Buffer, wt WireType) error {

	switch wt {
	case WireVarint:
		_, err := b.DecodeVarint()
		return err
	case WireBytes:
		size, err := b.DecodeVarint()

		if b.BytesRemains() < int(size) {
			return io.ErrUnexpectedEOF
		}

		b.ConsumeBytes(int(size))
		return err
	case WireZigzag32:
		_, err := b.DecodeZigzag32()
		return err
	case WireZigzag64:
		_, err := b.DecodeZigzag64()
		return err
	case WireFixed32:
		_, err := b.DecodeFixed32()
		return err
	case WireFixed64:
		_, err := b.DecodeFixed64()
		return err
	default:
		return ErrBadWireType
	}

	return nil
}

func rawUnmarshalStruct(b *Buffer, msg Struct) error {

	for b.BytesRemains() > 0 {
		wireTag, err := b.DecodeVarint()

		if err != nil {
			return err
		}

		fieldIndex, wt := parseWireTag(wireTag)

		err = msg.Unmarshal(b, fieldIndex, wt)

		if err == ErrUnknownField {
			err = skipField(b, wt)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func UnmarshalStruct(b *Buffer, wt WireType, msgPtr Struct) error {
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

		return rawUnmarshalStruct(limitBuffer, msgPtr)

	default:
		return ErrBadWireType
	}
}
