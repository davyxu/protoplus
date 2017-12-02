package parser

import (
	"errors"
)

func parseEnumField(ctx *Context) {

	// 注释
	nameToken := ctx.RawToken()

	// 字段名
	ctx.FieldDescriptor.Name = ctx.Expect(Token_Identifier).Value()

	if ctx.FieldNameExists(ctx.FieldDescriptor.Name) {
		panic(errors.New("Duplicate field name: " + ctx.FieldDescriptor.Name))
	}

	// 有等号
	if ctx.TokenID() == Token_Assign {
		ctx.NextToken()

		// tag
		ctx.FieldDescriptor.Tag = ctx.Expect(Token_Numeral).ToInt()

	} else { // 没等号自动生成枚举序号

		if len(ctx.Fields) == 0 {
			//fd.AutoTag = 0
		} else {

			// 按前面的序号+1
			//fd.AutoTag = d.MaxTag() + 1
		}

	}

	ctx.FieldDescriptor.Comment = ctx.CommentGroupByLine(nameToken.Line())

	// 枚举值类型，始终为int32
	ctx.FieldDescriptor.ParseType("int32")

	ctx.AddField(ctx.FieldDescriptor)

	return
}
