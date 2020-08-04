package parser

import (
	"errors"
	"github.com/davyxu/protoplus/model"
)

func parseObject(ctx *Context) {

	dotToken := ctx.RawToken()

	ctx.NextToken()

	// 名字
	ctx.Descriptor.Name = ctx.Expect(Token_Identifier).Value()
	ctx.SrcName = ctx.SourceName

	// 名字上面的注释

	// {
	ctx.Expect(Token_CurlyBraceL)

	for ctx.TokenID() != Token_CurlyBraceR {

		switch ctx.Descriptor.Kind {
		case model.Kind_Struct:
			var fd model.FieldDescriptor
			ctx.FieldDescriptor = &fd
			parseStructField(ctx)
		case model.Kind_Enum:
			var fd model.FieldDescriptor
			ctx.FieldDescriptor = &fd
			parseEnumField(ctx)
		case model.Kind_Service:
			var sc model.ServiceCall
			ctx.ServiceCall = &sc
			parseSvcCallField(ctx)
		}

		// 读取字段后面的[Tag项]
		if ctx.TokenID() == Token_BracketL {
			ctx.FieldDescriptor.TagSet = parseTagSet(ctx)
		}

	}

	ctx.Expect(Token_CurlyBraceR)

	// }

	ctx.Descriptor.Comment = ctx.CommentGroupByLine(dotToken.Line())

	// 名字重复检查

	if ctx.DescriptorSet.ObjectNameExists(ctx.Descriptor.Name) {
		panic(errors.New("Duplicate name: " + ctx.Descriptor.Name))
	}

	ctx.AddObject(ctx.Descriptor)

}
