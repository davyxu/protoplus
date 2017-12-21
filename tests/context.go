package main

import "github.com/davyxu/protoplus/model"

type Context struct {
	*model.DescriptorSet
	OutputFileName string
	PackageName    string
}
