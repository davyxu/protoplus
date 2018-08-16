package codegen

import (
	"github.com/davyxu/protoplus/model"
	"strings"
)

func ExportSymbolName(name string) string {
	return strings.ToUpper(string(name[0])) + name[1:]
}

func GoTypeName(fd *model.FieldDescriptor) string {
	switch fd.Type {
	case "bytes":
		return "[]byte"
	default:
		return fd.Type
	}
}

func init() {
	UsefulFunc["GoTypeName"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		if fd.Repeatd {
			ret += "[]"
		}

		ret += GoTypeName(fd)
		return
	}

	UsefulFunc["ExportSymbolName"] = ExportSymbolName

	UsefulFunc["GoFieldName"] = func(raw interface{}) string {

		fd := raw.(*model.FieldDescriptor)

		return ExportSymbolName(fd.Name)
	}

	UsefulFunc["GoStructTag"] = func(raw interface{}) string {

		fd := raw.(*model.FieldDescriptor)
		return fd.TagValueString("GoStructTag")
	}
}
