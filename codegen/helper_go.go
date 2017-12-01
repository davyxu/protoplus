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
