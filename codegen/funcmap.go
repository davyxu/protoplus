package codegen

import (
	"github.com/davyxu/protoplus/model"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{
	"ObjectLeadingComment": func(raw interface{}) (ret string) {

		d := raw.(*model.Descriptor)

		if d.Leading != "" {

			for index, line := range strings.Split(d.Leading, "\n") {
				if index > 0 {
					ret += "\n"
				}
				ret += "// " + line
			}

			return
		}

		return
	},

	"FieldTrailingComment": func(raw interface{}) string {

		fd := raw.(*model.FieldDescriptor)

		if fd.Trailing != "" {
			return "// " + fd.Trailing
		}

		return ""
	},

	"GoFieldName": func(raw interface{}) string {

		fd := raw.(*model.FieldDescriptor)

		return ExportSymbolName(fd.Name)
	},

	"GoTypeName": func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		if fd.Repeatd {
			ret += "[]"
		}

		ret += GoTypeName(fd)
		return
	},

	"CSTypeName": func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		ret += CSTypeName(fd)

		if fd.Repeatd {
			ret += "[]"
		}

		return
	},
}
