package codegen

import (
	"github.com/davyxu/protoplus/model"
)

func CSTypeName(fd *model.FieldDescriptor) string {
	switch fd.Type {
	case "int8":
		return "sbyte"
	case "int16":
		return "short"
	case "int32":
		return "int"
	case "int64":
		return "long"
	case "uint8":
		return "byte"
	case "uint16":
		return "ushort"
	case "uint32":
		return "uint"
	case "uint64":
		return "ulong"
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

func CSTypeNameFull(fd *model.FieldDescriptor) (ret string) {

	ret += CSTypeName(fd)

	if fd.Repeatd {
		ret += "[]"
	}

	return
}

func init() {
	UsefulFunc["CSTypeName"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		return CSTypeNameFull(fd)
	}
}
