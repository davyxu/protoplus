package parser

import (
	"errors"
	"fmt"
	"github.com/davyxu/golexer"
	"github.com/davyxu/protoplus/model"
	"io"
	"io/ioutil"
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

	return ctx.DescriptorSet, check(ctx)
}

func ParseFileList(dset *model.DescriptorSet, filelist ...string) error {

	ctx := newContext()

	for _, filename := range filelist {

		ctx.SourceName = filename
		ctx.DescriptorSet = dset

		if file, err := os.Open(filename); err != nil {
			return err
		} else {

			rawParse(ctx, file)

			file.Close()
		}

	}

	return check(ctx)
}

// 解析字符串
func rawParse(ctx *Context, reader io.Reader) (retErr error) {

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return retErr
	}

	ctx.protoParser = newProtoParser(ctx.SourceName)

	defer golexer.ErrorCatcher(func(err error) {

		retErr = fmt.Errorf("%s %s", ctx.PreTokenPos().String(), err.Error())

	})

	ctx.Lexer().Start(string(data))

	ctx.NextToken()

	for ctx.TokenID() != Token_EOF {

		var d model.Descriptor
		ctx.Descriptor = &d

		if ctx.TokenID() == Token_BracketL {
			d.TagSet = parseTagSet(ctx)
		}

		switch ctx.TokenID() {
		case Token_Struct:
			d.Kind = model.Kind_Struct
			parseObject(ctx)
		case Token_Enum:
			d.Kind = model.Kind_Enum
			parseObject(ctx)
		default:
			panic(errors.New("Unknown token: " + ctx.TokenValue()))
		}

	}

	return nil
}

func check(ctx *Context) error {

	for _, d := range ctx.Objects {

		for _, fd := range d.Fields {
			if fd.Kind == "" {

				// 将字段中使用的结构体的Kind确认为struct
				findD := ctx.ObjectByName(fd.Type)
				if findD == nil {
					return errors.New(fmt.Sprintf("type not found: %s at %s", d.Name, ctx.QuerySymbolPosString(fd)))
				}

				fd.Kind = findD.Kind
			}
		}

	}

	return nil

}
