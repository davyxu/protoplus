package pbscheme

import (
	"fmt"
	"github.com/davyxu/protoplus/model"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func PrimitiveToPbType(primitiveType string) string {
	switch primitiveType {
	case "uint16":
		return "uint32"
	case "int16":
		return "int32"
	case "float32":
		return "float"
	case "float64":
		return "double"
	default:
		return primitiveType
	}
}

func init() {
	UsefulFunc["PbTypeName"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		switch {
		case fd.IsMap():
			ret = fmt.Sprintf("map<%s,%s>", PrimitiveToPbType(fd.MapKey), PrimitiveToPbType(fd.MapValue))
		case fd.Repeatd:
			ret += "repeated "
			ret += PrimitiveToPbType(fd.Type)
		default:
			ret = PrimitiveToPbType(fd.Type)
		}

		return
	}

}
