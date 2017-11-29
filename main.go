package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/golog"
)

var log = golog.New("main")

// 显示版本号
var flagVersion = flag.Bool("version", false, "Show version")

var flagPlugin = flag.String("plugin", "", "")

var flagPluginCmd = flag.String("plugincmd", "", "")

const Version = "0.1.0"

func main() {

	flag.Parse()

	// 版本
	if *flagVersion {
		fmt.Println(Version)
		return
	}

	fmt.Println("protoplus!")

}
