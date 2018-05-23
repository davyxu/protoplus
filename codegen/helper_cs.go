package codegen

import (
	"github.com/davyxu/protoplus/model"
)

func CSTypeName(fd *model.FieldDescriptor) string {
	switch fd.Type {
	case "int8":
		return "SByte"
	case "int16":
		return "Int16"
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint8":
		return "Byte"
	case "uint16":
		return "UInt16"
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
