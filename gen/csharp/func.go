package csharp

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func CSTypeNameFull(fd *model.FieldDescriptor) (ret string) {

	if fd.Repeatd {
		return fmt.Sprintf("List<%s>", codegen.CSTypeName(fd))
	}

	return codegen.CSTypeName(fd)
}

func getEndPointPair(d *model.Descriptor) (from, to string) {

	msgdir := d.TagValueString("MsgDir")
	endPoints := strings.Split(msgdir, "->")
	if len(endPoints) >= 2 {

		from = strings.TrimSpace(endPoints[0])

		to = strings.TrimSpace(endPoints[1])
	}

	return
}

func init() {

	UsefulFunc["CSTypeNameFull"] = func(raw interface{}) (ret string) {

		fd := raw.(*model.FieldDescriptor)

		return CSTypeNameFull(fd)
	}

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

	UsefulFunc["GetSourcePeer"] = func(raw interface{}) string {

		d := raw.(*model.Descriptor)

		from, _ := getEndPointPair(d)

		return from
	}

	UsefulFunc["GetTargetPeer"] = func(raw interface{}) string {

		d := raw.(*model.Descriptor)

		_, to := getEndPointPair(d)

		return to
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
