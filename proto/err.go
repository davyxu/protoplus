package proto

import "errors"

var (
	ErrBadWireType  = errors.New("bad wire type")
	ErrUnknownField = errors.New("unknown field")
)
