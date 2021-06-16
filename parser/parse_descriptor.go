package parser

import (
	"errors"
	"github.com/davyxu/protoplus/model"
)

func parseObject(ctx *Context, d *model.Descriptor) {

	keywordToken := ctx.RawToken()

	ctx.NextToken()

	// 名字
	d.Name = ctx.Expect(Token_Identifier).Value()

	// 名字上面的注释

	// {
	ctx.Expect(Token_CurlyBraceL)

	for ctx.TokenID() != Token_CurlyBraceR {

		var fd model.FieldDescriptor
		fd.Descriptor = d

		switch d.Kind {
		case model.Kind_Struct:
			parseStructField(ctx, &fd)
		case model.Kind_Enum:
			parseEnumField(ctx, &fd)
		}

		// 读取字段后面的[Tag项]
		if ctx.TokenID() == Token_BracketL {
			fd.TagSet = parseTagSet(ctx)
		}

	}

	ctx.Expect(Token_CurlyBraceR)

	// }

	d.Comment = ctx.CommentGroupByLine(keywordToken.Line())

	// 名字重复检查

	if ctx.DescriptorSet.ObjectNameExists(d.Name) {
		panic(errors.New("Duplicate name: " + d.Name))
	}

	ctx.AddObject(d)

}
