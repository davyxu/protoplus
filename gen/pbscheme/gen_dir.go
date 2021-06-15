package pbscheme

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const protoDirCodeTemplate = `// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
syntax = "proto3";

option go_package= "./;{{.PackageName}}";

package {{.PackageName}};

{{range $a, $enumobj := .DependentSource}}
import "{{.}}"; {{end}}

{{range $a, $enumobj := .Enums}}
enum {{.Name}} {	{{range .Fields}}
	{{.Name}} = {{PbTagNumber $enumobj .}}; {{end}}
}{{end}}

{{range $a, $obj := .Structs}}
{{ObjectLeadingComment .}}
message {{.Name}} {	{{range .Fields}}
	{{PbTypeName .}} {{GoFieldName .}} = {{PbTagNumber $obj .}};{{FieldTrailingComment .}} {{end}}
}
{{end}}
`

func GenProtoDir(ctx *gen.Context) error {

	rootDS := ctx.DescriptorSet

	var sb strings.Builder

	for srcName, ds := range rootDS.DescriptorSetBySource() {

		ctx.DescriptorSet = ds

		gen := codegen.NewCodeGen("dirproto").
			RegisterTemplateFunc(codegen.UsefulFunc).
			RegisterTemplateFunc(UsefulFunc).
			ParseTemplate(protoDirCodeTemplate, ctx)

		if gen.Error() != nil {
			fmt.Println(string(gen.Data()))
			return gen.Error()
		}

		fullPathName := filepath.Join(ctx.OutputFileName, srcName)

		fmt.Fprintf(&sb, "%s ", srcName)

		err := gen.WriteOutputFile(fullPathName).Error()
		if err != nil {
			return err
		}
	}

	err := ioutil.WriteFile(filepath.Join(ctx.OutputFileName, "filelist.txt"), []byte(sb.String()), 0666)

	return err
}
