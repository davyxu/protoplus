package util

import (
	"flag"
	"github.com/davyxu/protoplus/model"
	"github.com/davyxu/protoplus/parser"
)

func ParseFileList(dset *model.DescriptorSet) (retErr error) {

	err := parser.ParseFileList(dset, flag.Args()...)
	if err != nil {
		return err
	}

	return

}
