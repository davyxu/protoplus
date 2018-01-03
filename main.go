package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/golog"
	_ "github.com/davyxu/protoplus/codegen"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/util"
	"os"
)

var log = golog.New("main")

// 显示版本号
var flagVersion = flag.Bool("version", false, "Show version")

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

	if err := util.RunGenerator(&dset); err != nil {
		fmt.Println("Generate error: ", err)
		os.Exit(1)
	}

}
