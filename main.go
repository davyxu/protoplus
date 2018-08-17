package main

import (
	"flag"
	"fmt"
	_ "github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/gen/gogopb"
	"github.com/davyxu/protoplus/gen/json"
	_ "github.com/davyxu/protoplus/gen/json"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/util"
	"os"
)

// 显示版本号
var (
	flagVersion = flag.Bool("version", false, "Show version")
	flagPackage = flag.String("package", "", "package name in source files")
	flagPbOut   = flag.String("pb_out", "", "pb schema output to file")
	flagJsonOut = flag.String("json_out", "", "json schema output to file")
	flagJson    = flag.Bool("json", false, "json schema output to std out")
)

const Version = "0.1.0"

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

	err = util.ParseFileList(ctx.DescriptorSet)

	if err != nil {
		goto OnError
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
