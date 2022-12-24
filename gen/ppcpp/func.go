package ppcpp

import (
	"github.com/davyxu/protoplus/gen"
	"github.com/davyxu/protoplus/model"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var UsefulFunc = template.FuncMap{}

func ChangeExtension(filename, newExt string) string {

	file := filepath.Base(filename)

	return strings.TrimSuffix(file, path.Ext(file)) + newExt
}

func SourceList(ctx *gen.Context) (ret []string) {

	rootDS := &model.PBDescriptorSet{DescriptorSet: *ctx.DescriptorSet}
	for _, protoFile := range rootDS.SourceList() {
		ret = append(ret, ChangeExtension(protoFile, ".pb.h"))
	}

	return
}

func init() {
	UsefulFunc["SourceList"] = SourceList
}
