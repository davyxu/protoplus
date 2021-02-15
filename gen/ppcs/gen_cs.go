package ppcs

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
)

func GenCSharp(ctx *gen.Context) error {

	gen := codegen.NewCodeGen("cs").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc).
		ParseTemplate(TemplateText, ctx)

	if gen.Error() != nil {
		fmt.Println(string(gen.Code()))
		return gen.Error()
	}

	return gen.WriteOutputFile(ctx.OutputFileName).Error()
}
