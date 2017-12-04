package parser

import (
	"github.com/davyxu/protoplus/model"
	"os"
	"strings"
)

func ParseFile(fileName string) (*model.DescriptorSet, error) {

	var dset model.DescriptorSet

	return &dset, ParseFileList(&dset, fileName)
}

func ParseString(script string) (*model.DescriptorSet, error) {

	ctx := newContext()
	ctx.SourceName = "string"
	ctx.DescriptorSet = new(model.DescriptorSet)

	if err := rawParse(ctx, strings.NewReader(script)); err != nil {
		return nil, err
	}

	return ctx.DescriptorSet, checkAndFix(ctx)
}

func ParseFileList(dset *model.DescriptorSet, filelist ...string) error {

	ctx := newContext()
	ctx.DescriptorSet = dset

	for _, filename := range filelist {

		ctx.SourceName = filename

		if file, err := os.Open(filename); err != nil {
			return err
		} else {

			if err := rawParse(ctx, file); err != nil {
				file.Close()

				return err
			}

			file.Close()

		}

	}

	return checkAndFix(ctx)
}
