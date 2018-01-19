package main

import (
	"flag"
	"fmt"
	_ "github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/msgidutil"
	"github.com/davyxu/protoplus/util"
	"os"
)

var (
	flagPackage         = flag.String("package", "", "package name in source files")
	flagGenSuggestMsgID = flag.Bool("GenSuggestMsgID", false, "Generate suggest msgid, default is false")
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
	ctx.DescriptorSet = new(model.DescriptorSet)

	util.RegisterGenerator(
		&util.Generator{
			FlagName:    "go_out",
			FlagComment: "go source output",
			GenFunc:     genAdaptor(&ctx, gen_go),
		},
	)

	flag.Parse()

	if err := util.ParseFileList(ctx.DescriptorSet); err != nil {
		fmt.Println("ParseFileList error: ", err)
		os.Exit(1)
	}

	if *flagGenSuggestMsgID {
		msgidutil.GenSuggestMsgID(ctx.DescriptorSet)
		return
	}

	if err := util.RunGenerator(ctx.DescriptorSet); err != nil {
		fmt.Println("Generate error: ", err)
		os.Exit(1)
	}

}
