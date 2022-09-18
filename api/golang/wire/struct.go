package wire

import (
	"io"
	"reflect"
)

type Struct interface {
	Marshal(buffer *Buffer) error

	Unmarshal(buffer *Buffer, fieldIndex uint64, wt WireType) error

	Size() int
}

func MarshalStruct(b *Buffer, fieldIndex uint64, msg Struct) error {

	structValue := reflect.ValueOf(msg)

	// *MyType被Message包裹后，判断不为nil
	if structValue.IsNil() {
		return nil
	}

	size := msg.Size()
	if size == 0 {
		return nil
	}

	b.EncodeVarint(MakeTag(fieldIndex, WireBytes))

	b.EncodeVarint(uint64(size))

	return msg.Marshal(b)
}

func SizeStruct(fieldIndex uint64, msg Struct) int {

	structValue := reflect.ValueOf(msg)

	// *MyType被Message包裹后，判断不为nil
	if structValue.IsNil() {
		return 0
	}

	size := msg.Size()

	if size == 0 {
		return 0
	}

	return SizeVarint(MakeTag(fieldIndex, WireVarint)) + SizeVarint(uint64(size)) + size
}

func UnmarshalStructObject(b *Buffer, msg Struct) error {

	for b.BytesRemains() > 0 {
		wireTag, err := b.DecodeVarint()

		if err != nil {
			return err
		}

		fieldIndex, wt := ParseTag(wireTag)

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

		return UnmarshalStructObject(limitBuffer, msgPtr)

	default:
		return ErrBadWireType
	}
}
