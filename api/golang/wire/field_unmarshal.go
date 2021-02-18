package wire

import (
	"io"
	"math"
)

func UnmarshalBool(b *Buffer, wt WireType) (bool, error) {
	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return false, err
		}
		return v != 0, nil
	default:
		return false, ErrBadWireType
	}
}

func UnmarshalInt32(b *Buffer, wt WireType) (int32, error) {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return 0, err
		}
		return int32(v), nil

	case WireFixed32:
		v, err := b.DecodeFixed32()
		if err != nil {
			return 0, err
		}

		return int32(v), nil

	default:
		return 0, ErrBadWireType
	}
}

func UnmarshalUInt32(b *Buffer, wt WireType) (uint32, error) {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return 0, err
		}

		return uint32(v), nil
	default:
		return 0, ErrBadWireType
	}
}

func UnmarshalInt64(b *Buffer, wt WireType) (int64, error) {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return 0, err
		}
		return int64(v), nil

	case WireFixed64:
		v, err := b.DecodeFixed64()
		if err != nil {
			return 0, err
		}

		return int64(v), nil
	default:
		return 0, ErrBadWireType
	}
}

func UnmarshalUInt64(b *Buffer, wt WireType) (uint64, error) {

	switch wt {
	case WireVarint:
		v, err := b.DecodeVarint()
		if err != nil {
			return 0, err
		}
		return v, nil
	default:
		return 0, ErrBadWireType
	}
}

func UnmarshalFloat32(b *Buffer, wt WireType) (float32, error) {

	switch wt {
	case WireFixed32:
		v, err := b.DecodeFixed32()
		if err != nil {
			return 0, err
		}

		return math.Float32frombits(uint32(v)), nil
	default:
		return 0, ErrBadWireType
	}

}

func UnmarshalFloat64(b *Buffer, wt WireType) (float64, error) {

	switch wt {
	case WireFixed64:
		v, err := b.DecodeFixed64()
		if err != nil {
			return 0, err
		}

		return math.Float64frombits(uint64(v)), nil
	default:
		return 0, ErrBadWireType
	}
}

func UnmarshalString(b *Buffer, wt WireType) (string, error) {
	switch wt {
	case WireBytes:
		v, err := b.DecodeStringBytes()
		if err != nil {
			return "", err
		}

		return v, nil
	default:
		return "", ErrBadWireType
	}
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
}
