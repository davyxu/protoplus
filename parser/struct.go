package parser

import (
	"errors"
	"github.com/davyxu/protoplus/meta"
)

func parseStruct(p *protoParser, fileD *meta.FileDescriptor, srcName string) {

	dotToken := p.RawToken()

	p.NextToken()

	d := meta.NewDescriptor(fileD)
	d.Type = meta.DescriptorType_Struct

	// 名字
	d.Name = p.Expect(Token_Identifier).Value()

	// 名字上面的注释
	d.CommentGroup = p.CommentGroupByLine(dotToken.Line())

	// {
	p.Expect(Token_CurlyBraceL)

	for p.TokenID() != Token_CurlyBraceR {

		// 枚举字段
		parseStructField(p, d)

	}

	p.Expect(Token_CurlyBraceR)

	// }

	// 名字重复检查

	if fileD.NameExists(d.Name) {
		panic(errors.New("Duplicate name: " + d.Name))
	}

	fileD.AddObject(d, srcName)

}
