package csharp

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func CSTypeNameFull(fd *model.FieldDescriptor) (ret string) {

	if fd.Repeatd {
		return fmt.Sprintf("List<%s>", codegen.CSTypeName(fd))
	}

	return codegen.CSTypeName(fd)
}

func init() {

	UsefulFunc["CSTypeNameFull"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		return CSTypeNameFull(fd)
	}

	//UsefulFunc["ObjectInitList"] = func(raw interface{}) bool {
	//
	//	d := raw.(*model.Descriptor)
	//	for _, fd := range d.Fields {
	//		if fd.Repeatd && fd.Kind != model.Kind_Struct || fd.Kind == model.Kind_Struct{
	//
	//		}
	//	}
	//
	//	return fd.Repeatd && fd.Kind == model.Kind_Struct
	//}

	UsefulFunc["IsPrimitiveSlice"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return fd.Repeatd && fd.Kind != model.Kind_Struct
	}

	UsefulFunc["IsStructSlice"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return fd.Repeatd && fd.Kind == model.Kind_Struct
	}

	UsefulFunc["IsStruct"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return !fd.Repeatd && fd.Kind == model.Kind_Struct
	}

	UsefulFunc["IsEnum"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return !fd.Repeatd && fd.Kind == model.Kind_Enum
	}

	UsefulFunc["IsEnumSlice"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return fd.Repeatd && fd.Kind == model.Kind_Enum
	}

	UsefulFunc["CodecName"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		switch fd.Type {
		case "bool":
			ret += "Bool"
		case "int32":
			ret += "Int32"
		case "uint32":
			ret += "UInt32"
		case "int64":
			ret += "Int64"
		case "uint64":
			ret += "UInt64"
		case "float32":
			ret += "Float"
		case "float64":
			ret += "Double"
		case "string":
			ret += "String"
		case "struct":
		case "bytes":
			ret += "Bytes"

		default:
			if fd.Kind == model.Kind_Struct {
				ret += "Struct"
			} else if fd.Kind == model.Kind_Enum {
				ret += "Enum"
			} else {
				panic("unknown Type " + fd.Type)
			}

		}

		return
	}
}
