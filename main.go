package main

import (
	"flag"
	"fmt"
	_ "github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/gen/gogopb"
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
)

const Version = "0.1.0"

func main() {

	flag.Parse()

	var dset model.DescriptorSet
	if err := util.ParseFileList(&dset); err != nil {
		fmt.Println("ParseFileList error: ", err)
		os.Exit(1)
	}

	// 版本
	if *flagVersion {
		fmt.Println(Version)
		return
	}

	if *flagPbOut != "" {
		if err := gogopb.Run(&gogopb.Context{
			DescriptorSet:  &dset,
			OutputFileName: *flagPbOut,
			PackageName:    *flagPackage,
		}); err != nil {
			fmt.Println(err)
		}
	}

}
