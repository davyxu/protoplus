package util

import (
	"github.com/davyxu/protoplus/meta"
	"github.com/davyxu/protoplus/parser"
	"os"
)

func ParseInputFiles() (fileset *meta.FileDescriptorSet, err error) {
	fileset = meta.NewFileDescriptorSet()

	err = parser.ParseFileList(fileset, os.Args[1:]...)

	return
}

//
//var flagHost = flag.String("protoplus", "", "out file")
//
//func ReadHostData() (string, error) {
//
//	if !flag.Parsed() {
//		flag.Parse()
//	}
//
//	data, err := ReadProcessOutput(*flagHost)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	return string(data), err
//}
//func ReadProcessOutput(filename string) ([]byte, error) {
//
//	cmd := exec.Command(filename)
//
//	data, err := cmd.CombinedOutput()
//
//	return data, err
//}
