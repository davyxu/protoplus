package parser

import (
	"errors"
	"github.com/davyxu/protoplus/model"
)

func parseEnumField(ctx *Context, fd *model.FieldDescriptor) {

	// 注释
	nameToken := ctx.RawToken()

	// 字段名
	fd.Name = ctx.Expect(Token_Identifier).Value()

	if fd.Descriptor.FieldNameExists(fd.Name) {
		panic(errors.New("Duplicate field name: " + fd.Name))
	}

	// 有等号
	if ctx.TokenID() == Token_Assign {
		ctx.NextToken()

		// tag
		fd.Tag = ctx.Expect(Token_Numeral).ToInt()

	} else { // 没等号自动生成枚举序号

		//if len(ctx.Fields) == 0 {
		//	//fd.AutoTag = 0
		//} else {
		//
		//	// 按前面的序号+1
		//	//fd.AutoTag = d.MaxTag() + 1
		//}

	}

	fd.Comment = ctx.CommentGroupByLine(nameToken.Line())

	// 枚举值类型，始终为int32
	fd.ParseType("int32")

	fd.Descriptor.AddField(fd)

	return
}
