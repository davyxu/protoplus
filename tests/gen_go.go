package main

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
)

// 报错行号+7
const goCodeTemplate = `
package main

func init() {
	{{range .Structs}} {{ if IsMessage . }}
	//cellnet.RegisterMessageMeta("binary","{{$.PackageName}}.{{.Name}}", reflect.TypeOf((*{{.Name}})(nil)).Elem(), {{StructMsgID .}}) {{end}} {{end}}
}

`

func gen_go(ctx *Context) error {

	gen := codegen.NewCodeGen("go").
		RegisterTemplateFunc(codegen.UsefulFunc).
		ParseTemplate(goCodeTemplate, ctx)

	if gen.Error() != nil {
		fmt.Println(string(gen.Data()))
		return gen.Error()
	}

	return gen.WriteOutputFile(ctx.OutputFileName).Error()
}
