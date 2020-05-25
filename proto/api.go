package proto

import (
	"errors"
	"reflect"
)

type Struct interface {
	Marshal(buffer *Buffer) error

	Unmarshal(buffer *Buffer, fieldIndex uint64, wt WireType) error

	Size() int
}

const (
	TopFieldIndex = 7
)

func Marshal(raw interface{}) ([]byte, error) {

	switch msg := raw.(type) {
	case Struct:
		l := msg.Size()

		data := make([]byte, 0, l)

		buffer := NewBuffer(data)

		err := msg.Marshal(buffer)
		if err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	default:
		size, err := sizeTypes(TopFieldIndex, raw)
		if err != nil {
			return nil, err
		}

		b := NewBufferBySize(size)

		err = marshalTypes(b, TopFieldIndex, raw)
		if err != nil {
			return nil, err
		}

		return b.Bytes(), nil
	}

}

func Size(raw interface{}) int {
	msg := raw.(Struct)

	return msg.Size()
}

func Unmarshal(data []byte, raw interface{}) (err error) {

	buffer := NewBuffer(data)

	switch msg := raw.(type) {
	case Struct:
		return rawUnmarshalStruct(buffer, msg)
	default:

		for buffer.BytesRemains() > 0 {
			wireTag, err := buffer.DecodeVarint()

			if err != nil {
				return err
			}

			fieldIndex, wt := parseWireTag(wireTag)
			if fieldIndex != TopFieldIndex {
				return errors.New("invalid top field tag")
			}

			err = unmarshalTypes(buffer, wt, raw)
			if err != nil {
				return err
			}
		}

		return nil

	}
}

func sizeTypes(fieldIndex uint64, raw interface{}) (size int, err error) {
	switch v := raw.(type) {
	case Struct:
		return SizeStruct(fieldIndex, v), nil
	case int32:
		return SizeInt32(fieldIndex, v), nil
	case int64:
		return SizeInt64(fieldIndex, v), nil
	case uint32:
		return SizeUInt32(fieldIndex, v), nil
	case uint64:
		return SizeUInt64(fieldIndex, v), nil
	case float32:
		return SizeFloat32(fieldIndex, v), nil
	case float64:
		return SizeFloat64(fieldIndex, v), nil
	case bool:
		return SizeBool(fieldIndex, v), nil
	case string:
		return SizeString(fieldIndex, v), nil
	case []byte:
		return SizeBytes(fieldIndex, v), nil
	case []int32:
		return SizeInt32Slice(fieldIndex, v), nil
	case []int64:
		return SizeInt64Slice(fieldIndex, v), nil
	case []uint32:
		return SizeUInt32Slice(fieldIndex, v), nil
	case []uint64:
		return SizeUInt64Slice(fieldIndex, v), nil
	case []float32:
		return SizeFloat32Slice(fieldIndex, v), nil
	case []float64:
		return SizeFloat64Slice(fieldIndex, v), nil
	case []bool:
		return SizeBoolSlice(fieldIndex, v), nil
	case []string:
		return SizeStringSlice(fieldIndex, v), nil
	}

	tRaw := reflect.TypeOf(raw)
	switch tRaw.Kind() {
	case reflect.Array, reflect.Slice:

		vRaw := reflect.ValueOf(raw)

		for i := 0; i < vRaw.Len(); i++ {
			vElement := vRaw.Index(i)
			var thisSize int
			thisSize, err = sizeTypes(fieldIndex, vElement.Interface())
			if err != nil {
				return 0, err
			}

			size += thisSize
		}
		return

	default:
		return 0, errors.New("unsupport type: " + tRaw.Kind().String())
	}
}

func unmarshalTypes(buffer *Buffer, wt WireType, raw interface{}) error {
	switch v := raw.(type) {
	case Struct:
		return UnmarshalStruct(buffer, wt, v)
	case *int32:
		vv, err := UnmarshalInt32(buffer, wt)
		*v = vv
		return err
	case *int64:
		vv, err := UnmarshalInt64(buffer, wt)
		*v = vv
		return err
	case *uint32:
		vv, err := UnmarshalUInt32(buffer, wt)
		*v = vv
		return err
	case *uint64:
		vv, err := UnmarshalUInt64(buffer, wt)
		*v = vv
		return err
	case *float32:
		vv, err := UnmarshalFloat32(buffer, wt)
		*v = vv
		return err
	case *float64:
		vv, err := UnmarshalFloat64(buffer, wt)
		*v = vv
		return err
	case *bool:
		vv, err := UnmarshalBool(buffer, wt)
		*v = vv
		return err
	case *string:
		vv, err := UnmarshalString(buffer, wt)
		*v = vv
		return err
	case *[]byte:
		vv, err := UnmarshalBytes(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]int32:
		vv, err := UnmarshalInt32Slice(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]int64:
		vv, err := UnmarshalInt64Slice(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]uint32:
		vv, err := UnmarshalUInt32Slice(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]uint64:
		vv, err := UnmarshalUInt64Slice(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]float32:
		vv, err := UnmarshalFloat32Slice(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]float64:
		vv, err := UnmarshalFloat64Slice(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]bool:
		vv, err := UnmarshalBoolSlice(buffer, wt)
		*v = append(*v, vv...)
		return err
	case *[]string:
		vv, err := UnmarshalStringSlice(buffer, wt)
		*v = append(*v, vv...)
		return err
	}

	tRaw := reflect.TypeOf(raw)
	switch tRaw.Kind() {
	case reflect.Array, reflect.Slice:

		vRaw := reflect.ValueOf(raw)

		vElement := reflect.New(reflect.TypeOf(raw).Elem())
		err := unmarshalTypes(buffer, wt, vElement.Interface())
		if err != nil {
			return err
		}

		vRaw.Set(reflect.AppendSlice(vRaw, vElement))

		return nil
	default:
		return errors.New("unsupport type: " + tRaw.Kind().String())
	}
}

func marshalTypes(buffer *Buffer, fieldIndex uint64, raw interface{}) error {

	switch v := raw.(type) {
	case Struct:
		return MarshalStruct(buffer, fieldIndex, v)
	case int32:
		return MarshalInt32(buffer, fieldIndex, v)
	case int64:
		return MarshalInt64(buffer, fieldIndex, v)
	case uint32:
		return MarshalUInt32(buffer, fieldIndex, v)
	case uint64:
		return MarshalUInt64(buffer, fieldIndex, v)
	case float32:
		return MarshalFloat32(buffer, fieldIndex, v)
	case float64:
		return MarshalFloat64(buffer, fieldIndex, v)
	case bool:
		return MarshalBool(buffer, fieldIndex, v)
	case string:
		return MarshalString(buffer, fieldIndex, v)
	case []byte:
		return MarshalBytes(buffer, fieldIndex, v)
	case []int32:
		return MarshalInt32Slice(buffer, fieldIndex, v)
	case []int64:
		return MarshalInt64Slice(buffer, fieldIndex, v)
	case []uint32:
		return MarshalUInt32Slice(buffer, fieldIndex, v)
	case []uint64:
		return MarshalUInt64Slice(buffer, fieldIndex, v)
	case []float32:
		return MarshalFloat32Slice(buffer, fieldIndex, v)
	case []float64:
		return MarshalFloat64Slice(buffer, fieldIndex, v)
	case []bool:
		return MarshalBoolSlice(buffer, fieldIndex, v)
	case []string:
		return MarshalStringSlice(buffer, fieldIndex, v)
	}

	tRaw := reflect.TypeOf(raw)
	switch tRaw.Kind() {
	case reflect.Array, reflect.Slice:

		vRaw := reflect.ValueOf(raw)

		for i := 0; i < vRaw.Len(); i++ {
			vElement := vRaw.Index(i)
			err := marshalTypes(buffer, fieldIndex, vElement.Interface())
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return errors.New("unsupport type: " + tRaw.Kind().String())
	}
}
