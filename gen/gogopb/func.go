package gogopb

import (
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
		case "float64":
			ret += "double"
		default:
			ret += fd.Type
		}

		return
	}

}
