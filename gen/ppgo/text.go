package ppgo

// 报错行号+7
const TemplateText = `// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
package {{.PackageName}}

import (
	"github.com/davyxu/protoplus/api/golang"
	"github.com/davyxu/protoplus/api/golang/wire"
	"unsafe"
)
var (
	_ *wire.Buffer		
	_ = ppgo.Marshal
	_ unsafe.Pointer
)

{{range $a, $enumobj := .Enums}}
type {{.Name}} int32
const (	{{range .Fields}}
	{{$enumobj.Name}}_{{.Name}} {{$enumobj.Name}} = {{PbTagNumber $enumobj .}} {{end}}
)

var (
{{$enumobj.Name}}MapperValueByName = map[string]int32{ {{range .Fields}}
	"{{.Name}}": {{PbTagNumber $enumobj .}}, {{end}}
}

{{$enumobj.Name}}MapperNameByValue = map[int32]string{ {{range .Fields}}
	{{PbTagNumber $enumobj .}}: "{{.Name}}" , {{end}}
}
)

func (self {{$enumobj.Name}}) String() string {
	return {{$enumobj.Name}}MapperNameByValue[int32(self)]
}
{{end}}

{{range $a, $obj := .Structs}}
{{ObjectLeadingComment .}}
type {{.Name}} struct{	{{range .Fields}}
	{{GoFieldName .}} {{ProtoTypeName .}} {{GoStructTag .}}{{FieldTrailingComment .}} {{end}}
}

func (self *{{.Name}}) String() string { return ppgo.CompactTextString(self) }

func (self *{{.Name}}) Size() (ret int) {
{{range .Fields}}
	{{if IsStructSlice .}}
	if len(self.{{GoFieldName .}}) > 0 {
		for _, elm := range self.{{GoFieldName .}} {
			ret += wire.SizeStruct({{PbTagNumber $obj .}}, elm)
		}
	}
	{{else if IsStruct .}}
	ret += wire.Size{{CodecName .}}({{PbTagNumber $obj .}}, self.{{GoFieldName .}})
	{{else if IsEnum .}}
	ret += wire.Size{{CodecName .}}({{PbTagNumber $obj .}}, int32(self.{{GoFieldName .}}))
	{{else if IsEnumSlice .}}
	ret += wire.Size{{CodecName .}}({{PbTagNumber $obj .}}, *(*[]int32)(unsafe.Pointer(&self.{{GoFieldName .}})))
	{{else}}
	ret += wire.Size{{CodecName .}}({{PbTagNumber $obj .}}, self.{{GoFieldName .}})
	{{end}}
{{end}}
	return
}

func (self *{{.Name}}) Marshal(buffer *wire.Buffer) error {
{{range .Fields}}
	{{if IsStructSlice .}}
		for _, elm := range self.{{GoFieldName .}} {
			wire.MarshalStruct(buffer, {{PbTagNumber $obj .}}, elm)
		}
	{{else if IsStruct .}}
		wire.Marshal{{CodecName .}}(buffer, {{PbTagNumber $obj .}}, self.{{GoFieldName .}})
	{{else if IsEnum .}}
		wire.Marshal{{CodecName .}}(buffer, {{PbTagNumber $obj .}}, int32(self.{{GoFieldName .}}))
	{{else if IsEnumSlice .}}
		wire.Marshal{{CodecName .}}(buffer, {{PbTagNumber $obj .}}, *(*[]int32)(unsafe.Pointer(&self.{{GoFieldName .}})))
	{{else}}	
		wire.Marshal{{CodecName .}}(buffer, {{PbTagNumber $obj .}}, self.{{GoFieldName .}})
	{{end}}
{{end}}
	return nil
}

func (self *{{.Name}}) Unmarshal(buffer *wire.Buffer, fieldIndex uint64, wt wire.WireType) error {
	switch fieldIndex {
	{{range .Fields}} case {{PbTagNumber $obj .}}: {{if IsStructSlice .}}
		var elm {{.Type}}
		if err := wire.UnmarshalStruct(buffer, wt, &elm); err != nil {
			return err
		} else {
			self.{{GoFieldName .}} = append(self.{{GoFieldName .}}, &elm)
			return nil
		}{{else if IsEnum .}}
		v, err := wire.Unmarshal{{CodecName .}}(buffer, wt)
		self.{{GoFieldName .}} = {{ProtoTypeName .}}(v)
		return err {{else if IsStruct .}}
		var elm {{.Type}}
		self.{{GoFieldName .}} = &elm
        return wire.Unmarshal{{CodecName .}}(buffer, wt, self.{{GoFieldName .}}) {{else if IsEnumSlice .}}
		v, err := wire.Unmarshal{{CodecName .}}(buffer, wt)
		for _, vv := range v {	
			self.{{GoFieldName .}} = append(self.{{GoFieldName .}}, {{ProtoElementTypeName .}}(vv))
		}
		return err {{else if IsSlice .}}
		v, err := wire.Unmarshal{{CodecName .}}(buffer, wt)	
		self.{{GoFieldName .}} = append(self.{{GoFieldName .}}, v...)
		return err {{else}}
		v, err := wire.Unmarshal{{CodecName .}}(buffer, wt)
		self.{{GoFieldName .}} = v
		return err{{end}}
		{{end}}
	}

	return wire.ErrUnknownField
}
{{end}}
`

const RegTemplateText = `// Generated by github.com/davyxu/protoplus
package {{.PackageName}}

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	"reflect"
)


var (
	_ cellnet.MessageMeta
	_ codec.CodecRecycler
	_ reflect.Kind
)

func init() {
	{{range .Structs}}
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("{{StructCodec .}}"),	
		Type:  reflect.TypeOf((*{{.Name}})(nil)).Elem(),
		ID:    {{StructMsgID .}},
		New:  func() interface{} { return &{{.Name}}{} },
	}) {{end}}
}
`
