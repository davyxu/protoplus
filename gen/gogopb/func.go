package gogopb

import (
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func init() {
	UsefulFunc["PbTypeName"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		if fd.Repeatd {
			ret += "repeated "
		}

		switch fd.Type {
		case "uint16":
			ret += "uint32"
		case "int16":
			ret += "int32"
		case "float32":
			ret += "float"
		default:
			ret += fd.Type
		}

		return
	}

	UsefulFunc["PbTagNumber"] = func(rawD, rawFD interface{}) (tag int) {
		d := rawD.(*model.Descriptor)
		fd := rawFD.(*model.FieldDescriptor)

		if d.Kind == model.Kind_Enum {
			return codegen.TagNumber(d, fd)
		} else {
			return codegen.TagNumber(d, fd) + 1
		}

	}
}
