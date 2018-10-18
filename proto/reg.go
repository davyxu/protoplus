package proto

import (
	"reflect"
	"sync"
)

type Struct interface {
	Marshal(buffer *Buffer) error

	Unmarshal(buffer *Buffer, fieldIndex uint64, wt WireType) error

	Size() int
}

type Meta struct {
	Type reflect.Type
}

var (
	metaByType sync.Map
)

func RegisterType(meta *Meta) {

	if meta.Type.Kind() != reflect.Struct {
		panic("expect struct type")
	}

	metaByType.Store(meta.Type, meta)
}

func MetaByType(t reflect.Type) *Meta {

	if raw, ok := metaByType.Load(t); ok {
		return raw.(*Meta)
	}

	return nil
}
