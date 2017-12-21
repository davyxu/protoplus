package main

import (
	"flag"
	"fmt"
	_ "github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/util"
	"os"
)

var (
	flagPackage = flag.String("package", "", "package name in source files")
)

func genAdaptor(ctx *Context, f func(*Context) error) util.GenFunc {

	return func(dset *model.DescriptorSet, fileName string) error {
		ctx.OutputFileName = fileName
		ctx.DescriptorSet = dset
		ctx.PackageName = *flagPackage

		return f(ctx)
	}
}

func main() {

	var ctx Context

	util.RegisterGenerator(
		&util.Generator{
			FlagName:    "go_out",
			FlagComment: "go source output",
			GenFunc:     genAdaptor(&ctx, gen_go),
		},
	)

	err := util.RunGenerator()

	if err != nil {
		fmt.Println("Generate error: ", err)
		os.Exit(1)
	}

}
