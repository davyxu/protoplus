package model

type Kind string

const (
	Kind_None      Kind = ""
	Kind_Primitive      = "primitive"
	Kind_Struct         = "struct"
	Kind_Enum           = "enum"
)

var SchemeType2Type = map[string]string{
	"int32":   "int32",
	"int64":   "int64",
	"uint32":  "uint32",
	"uint64":  "uint64",
	"float32": "float32",
	"float64": "float64",
	"bool":    "bool",
	"string":  "string",
	"bytes":   "bytes",
}
