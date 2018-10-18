package proto

import "errors"

var (
	ErrFieldIndexOutOfRange = errors.New("fieldindex out of range ")
	ErrBadWireType          = errors.New("bad wire type")
	ErrMetaNotFound         = errors.New("meta not found")
)
