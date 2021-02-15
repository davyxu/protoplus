package ppgo

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
)

func GenGo(ctx *gen.Context) error {

	codeGen := codegen.NewCodeGen("go").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc).
		ParseTemplate(TemplateText, ctx).
		FormatGoCode()

	if codeGen.Error() != nil {
		fmt.Println(string(codeGen.Code()))
		return codeGen.Error()
	}

	return codeGen.WriteOutputFile(ctx.OutputFileName).Error()
}

func GenGoReg(ctx *gen.Context) error {

	codeGen := codegen.NewCodeGen("goreg").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc).
		ParseTemplate(RegTemplateText, ctx).
		FormatGoCode()

	if codeGen.Error() != nil {
		fmt.Println(codeGen.Code())
		return codeGen.Error()
	}

	return codeGen.WriteOutputFile(ctx.OutputFileName).Error()
}
