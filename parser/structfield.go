package parser

import (
	"errors"
)

func parseStructField(ctx *Context) {

	// 注释
	nameToken := ctx.RawToken()

	// 字段名
	ctx.FieldDescriptor.Name = ctx.Expect(Token_Identifier).Value()

	if ctx.FieldNameExists(ctx.FieldDescriptor.Name) {
		panic(errors.New("Duplicate field name: " + ctx.FieldDescriptor.Name))
	}

	tp := ctx.TokenPos()

	// [  数组类型
	if ctx.TokenID() == Token_BracketL {
		ctx.NextToken()
		ctx.Expect(Token_BracketR)
		ctx.Repeatd = true
	}

	// 延后在所有解析完后，检查TypeName是否合法，通过symbol还原位置并报错
	ctx.FieldDescriptor.ParseType(ctx.Expect(Token_Identifier).Value())

	ctx.FieldDescriptor.Comment = ctx.CommentGroupByLine(nameToken.Line())

	ctx.AddSymbol(ctx.FieldDescriptor, tp)

	ctx.AddField(ctx.FieldDescriptor)

	return
}
