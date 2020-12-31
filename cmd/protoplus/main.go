package main

import (
	"flag"
	"fmt"
	_ "github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/gen/csharp"
	"github.com/davyxu/protoplus/gen/gogopb"
	"github.com/davyxu/protoplus/gen/golang"
	"github.com/davyxu/protoplus/gen/json"
	_ "github.com/davyxu/protoplus/gen/json"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/util"
	"os"
)

// 显示版本号
var (
	flagVersion    = flag.Bool("version", false, "Show version")
	flagPackage    = flag.String("package", "", "package name in source files")
	flagPbOut      = flag.String("pb_out", "", "output google protobuf schema file")
	flagGoOut      = flag.String("go_out", "", "output golang source file")
	flagCSOut      = flag.String("cs_out", "", "output csharp source file")
	flagJsonOut    = flag.String("json_out", "", "output json file")
	flagGoRegOut   = flag.String("goreg_out", "", "output golang message register source file")
	flagJson       = flag.Bool("json", false, "output json to std out")
	flagGenReg     = flag.Bool("genreg", false, "gen message register entry")
	flagStructBase = flag.String("structbase", "IProtoStruct", "struct inherite class type name in c#")
	flagCodec      = flag.String("codec", "protoplus", "default codec in register entry")
)

const Version = "2.0.0"

func main() {

	flag.Parse()

	// 版本
	if *flagVersion {
		fmt.Println(Version)
		return
	}

	var err error
	var ctx gen.Context
	ctx.DescriptorSet = new(model.DescriptorSet)
	ctx.PackageName = *flagPackage
	ctx.StructBase = *flagStructBase
	ctx.RegEntry = *flagGenReg
	ctx.Codec = *flagCodec

	err = util.ParseFileList(ctx.DescriptorSet)

	if err != nil {
		goto OnError
	}

	if *flagGoOut != "" {
		ctx.OutputFileName = *flagGoOut

		err = golang.GenGo(&ctx)

		if err != nil {
			goto OnError
		}
	}

	if *flagGoRegOut != "" {
		ctx.OutputFileName = *flagGoRegOut

		err = golang.GenGoReg(&ctx)

		if err != nil {
			goto OnError
		}
	}

	if *flagCSOut != "" {
		ctx.OutputFileName = *flagCSOut

		err = csharp.GenCSharp(&ctx)

		if err != nil {
			goto OnError
		}
	}

	if *flagPbOut != "" {
		ctx.OutputFileName = *flagPbOut

		err = gogopb.GenProto(&ctx)

		if err != nil {
			goto OnError
		}
	}

	if *flagJsonOut != "" {
		ctx.OutputFileName = *flagJsonOut

		err = json.GenJson(&ctx)

		if err != nil {
			goto OnError
		}
	}

	if *flagJson {

		err = json.OutputJson(&ctx)

		if err != nil {
			goto OnError
		}
	}

	return

OnError:
	fmt.Println(err)
	os.Exit(1)
}
