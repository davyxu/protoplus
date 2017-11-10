package parser

import (
	"errors"
	"github.com/davyxu/protoplus/meta"
)

func parseEnum(p *protoParser, fileD *meta.FileDescriptor, srcName string) {

	// enum
	enumToken := p.Expect(Token_Enum)

	d := meta.NewDescriptor(fileD)
	d.Type = meta.DescriptorType_Enum

	// 名字
	d.Name = p.Expect(Token_Identifier).Value()

	d.CommentGroup = p.CommentGroupByLine(enumToken.Line())

	// {
	p.Expect(Token_CurlyBraceL)

	for p.TokenID() != Token_CurlyBraceR {

		// 字段
		parseEnumField(p, d)

	}

	p.Expect(Token_CurlyBraceR)

	// }

	// 名字重复检查

	if fileD.NameExists(d.Name) {
		panic(errors.New("Duplicate name: " + d.Name))
	}

	fileD.AddObject(d, srcName)

}
