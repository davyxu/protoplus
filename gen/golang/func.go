package golang

import (
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func init() {

	// 所有结构生成Entry
	UsefulFunc["GenEntry"] = func(d *model.Descriptor) bool {
		return true
	}

	UsefulFunc["StructCodec"] = func(d *model.Descriptor) string {
		codecName := d.TagValueString("Codec")
		if codecName == "" {
			return d.DescriptorSet.Codec
		}

		return codecName
	}
	UsefulFunc["ProtoTypeName"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		if fd.Repeatd {
			ret += "[]"
		}

		ret += codegen.GoTypeName(fd)
		return
	}

	UsefulFunc["ProtoElementTypeName"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		ret += codegen.GoTypeName(fd)
		return
	}

	UsefulFunc["IsStructSlice"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return fd.Repeatd && fd.Kind == model.Kind_Struct
	}

	UsefulFunc["IsStruct"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return !fd.Repeatd && fd.Kind == model.Kind_Struct
	}

	UsefulFunc["IsSlice"] = func(raw interface{}) bool {

		fd := raw.(*model.FieldDescriptor)

		return fd.Repeatd
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
			ret += "Float32"
		case "float64":
			ret += "Float64"
		case "string":
			ret += "String"
		case "struct":
		case "bytes":
			ret += "Bytes"

		default:
			if fd.Kind == model.Kind_Struct {
				ret += "Struct"
			} else if fd.Kind == model.Kind_Enum {
				ret += "Int32"
			} else {
				panic("unknown Type " + fd.Type)
			}

		}

		if fd.Repeatd {
			ret += "Slice"
		}

		return
	}
}
