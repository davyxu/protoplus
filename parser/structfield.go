package parser

import (
	"errors"
	"github.com/davyxu/protoplus/model"
)

func parseStructField(ctx *Context, fd *model.FieldDescriptor) {

	// 注释
	nameToken := ctx.RawToken()

	// 字段名
	fd.Name = ctx.Expect(Token_Identifier).Value()

	if fd.Descriptor.FieldNameExists(fd.Name) {
		panic(errors.New("Duplicate field name: " + fd.Name))
	}

	tp := ctx.TokenPos()

	switch ctx.TokenID() {
	case Token_BracketL: // [  数组类型
		ctx.NextToken()
		ctx.Expect(Token_BracketR)
		fd.Repeatd = true
		// 延后在所有解析完后，检查TypeName是否合法，通过symbol还原位置并报错
		fd.ParseType(ctx.Expect(Token_Identifier).Value())
	case Token_Map:
		ctx.NextToken()
		ctx.Expect(Token_BracketL)
		fd.MapKey = ctx.Expect(Token_Identifier).Value()
		ctx.Expect(Token_BracketR)
		fd.MapValue = ctx.Expect(Token_Identifier).Value()
	default:
		// 延后在所有解析完后，检查TypeName是否合法，通过symbol还原位置并报错
		fd.ParseType(ctx.Expect(Token_Identifier).Value())
	}

	fd.Comment = ctx.CommentGroupByLine(nameToken.Line())

	ctx.AddSymbol(fd, tp)

	fd.Descriptor.AddField(fd)

	return
}
