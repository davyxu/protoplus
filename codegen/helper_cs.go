package codegen

import (
	"github.com/davyxu/protoplus/model"
)

func CSTypeName(fd *model.FieldDescriptor) string {
	switch fd.Type {
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint32":
		return "UInt32"
	case "uint64":
		return "UInt64"
	case "float32":
		return "float"
	case "float64":
		return "double"
	case "string":
		return "string"
	case "bool":
		return "bool"
	case "bytes":
		return "byte[]"
	default:
		return fd.Type
	}
}
