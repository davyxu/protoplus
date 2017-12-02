package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/davyxu/protoplus/util"
	"os"
)

var log = golog.New("main")

// 显示版本号
var flagVersion = flag.Bool("version", false, "Show version")

const Version = "0.1.0"

func main() {

	flag.Parse()

	// 版本
	if *flagVersion {
		fmt.Println(Version)
		return
	}

	if err := util.RunGenerator(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
