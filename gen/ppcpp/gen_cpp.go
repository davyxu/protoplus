package ppcpp

import (
	"fmt"
	"github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
)

func GenCppReg(ctx *gen.Context) error {

	codeGen := codegen.NewCodeGen("cppreg").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc).
		ParseTemplate(RegTemplateText, ctx)

	if codeGen.Error() != nil {
		fmt.Println(codeGen.Code())
		return codeGen.Error()
	}

	return codeGen.WriteOutputFile(ctx.OutputFileName).Error()
}
