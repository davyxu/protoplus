package codegen

import (
	"github.com/davyxu/protoplus/model"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func init() {
	UsefulFunc["ObjectLeadingComment"] = func(raw interface{}) (ret string) {

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
	}

	UsefulFunc["FieldTrailingComment"] = func(raw interface{}) string {

		fd := raw.(*model.FieldDescriptor)

		if fd.Trailing != "" {
			return "// " + fd.Trailing
		}

		return ""
	}

	UsefulFunc["TagNumber"] = func(rawD, rawFD interface{}) (tag int) {
		d := rawD.(*model.Descriptor)
		fd := rawFD.(*model.FieldDescriptor)

		tag = -1
		for _, libfd := range d.Fields {

			if libfd.Tag != 0 {
				tag = libfd.Tag
			} else {
				tag++
			}

			if libfd == fd {
				return tag
			}
		}

		return 0
	}

	UsefulFunc["IsMessage"] = func(raw interface{}) bool {
		d := raw.(*model.Descriptor)
		return strings.HasSuffix(d.Name, "REQ") || strings.HasSuffix(d.Name, "ACK")
	}
}
