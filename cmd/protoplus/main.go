package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/protoplus/build"
	_ "github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/gen/pbscheme"
	"github.com/davyxu/protoplus/gen/ppcs"
	"github.com/davyxu/protoplus/gen/ppgo"
	"github.com/davyxu/protoplus/gen/ppscheme"
	_ "github.com/davyxu/protoplus/gen/ppscheme"
	"github.com/davyxu/protoplus/gen/route"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/util"
	"os"
)

var (
	flagVersion = flag.Bool("version", false, "show version")

	flagPackage   = flag.String("package", "", "package name in source files")
	flagClassBase = flag.String("classbase", "IProtoStruct", "struct inherite class type name in c#")
	flagCodec     = flag.String("codec", "protoplus", "default codec in register entry")
)

type GenEntry struct {
	name        string
	usage       string
	flagOutFile *string
	flagExecute *bool

	outfile func(ctx *gen.Context) error
	execute func(ctx *gen.Context) error
}

var (
	genEntryList = []*GenEntry{
		{name: "ppgo_out", usage: "output protoplus message serialize golang source file", outfile: ppgo.GenGo},
		{name: "ppgoreg_out", usage: "output protoplus message register entry in golang", outfile: ppgo.GenGoReg},
		{name: "ppcs_out", usage: "output protoplus message serialize csharp source file", outfile: ppcs.GenCS},
		{name: "ppcsreg_out", usage: "output protoplus message register entry in csharp", outfile: ppcs.GenCSReg},
		{name: "pbscheme_out", usage: "output google protobuf schema file as single file", outfile: pbscheme.GenProto},

		// 使用例子: protoc $(cat filelist.txt)
		{name: "pbscheme_dir", usage: "output google protobuf schema files into dir", outfile: pbscheme.GenProtoDir},
		{name: "ppscheme_out", usage: "output protoplus scheme json file", outfile: ppscheme.GenJson},
		{name: "route_out", usage: "output route table json file", outfile: route.GenJson},

		{name: "ppscheme", usage: "output protoplus scheme json to std out", execute: ppscheme.OutputJson},
		{name: "route", usage: "output route table json to std out", execute: route.OutputJson},
	}
)

func defineEntryFlag() {
	for _, entry := range genEntryList {
		if entry.outfile != nil {
			entry.flagOutFile = flag.String(entry.name, "", entry.usage)
		}
		if entry.execute != nil {
			entry.flagExecute = flag.Bool(entry.name, false, entry.usage)
		}

	}
}

func runEntry(ctx *gen.Context) error {
	for _, entry := range genEntryList {
		if entry.flagOutFile != nil && *entry.flagOutFile != "" {
			ctx.OutputFileName = *entry.flagOutFile

			fmt.Printf("[%s] %s\n", entry.name, ctx.OutputFileName)
			err := entry.outfile(ctx)
			if err != nil {
				return err
			}
		}

		if entry.flagExecute != nil && *entry.flagExecute {
			err := entry.execute(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {

	defineEntryFlag()

	flag.Parse()

	// 版本
	if *flagVersion {
		build.Print()
		return
	}

	var err error
	var ctx gen.Context
	ctx.DescriptorSet = new(model.DescriptorSet)
	ctx.PackageName = *flagPackage
	ctx.ClassBase = *flagClassBase
	ctx.Codec = *flagCodec

	err = util.ParseFileList(ctx.DescriptorSet)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = runEntry(&ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
