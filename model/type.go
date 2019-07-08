package model

type Kind string

const (
	Kind_None      Kind = ""
	Kind_Primitive      = "primitive"
	Kind_Struct         = "struct"
	Kind_Enum           = "enum"
	Kind_Service        = "service"
)

var SchemeType2Type = map[string]string{
	"int8":    "int8",
	"int16":   "int16",
	"int32":   "int32",
	"int64":   "int64",
	"uint8":   "uint8",
	"uint16":  "uint16",
	"uint32":  "uint32",
	"uint64":  "uint64",
	"float32": "float32",
	"float64": "float64",
	"bool":    "bool",
	"string":  "string",
	"bytes":   "bytes",
}

func TypeSize(t string) int32 {
	switch t {
	case "int8":
		return 1
	case "int16":
		return 2
	case "int32":
		return 4
	case "int64":
		return 8
	case "uint8":
		return 1
	case "uint16":
		return 2
	case "uint32":
		return 4
	case "uint64":
		return 8
	case "float32":
		return 4
	case "float64":
		return 8
	case "bool":
		return 1
	default:
		return 0

	}
}
