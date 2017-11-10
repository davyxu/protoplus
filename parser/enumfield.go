package parser

import (
	"errors"
	"github.com/davyxu/protoplus/meta"
)

func parseEnumField(p *protoParser, d *meta.Descriptor) {

	fd := meta.NewFieldDescriptor(d)

	nameToken := p.RawToken()

	// 字段名
	fd.Name = p.Expect(Token_Identifier).Value()

	if _, ok := d.FieldByName[fd.Name]; ok {
		panic(errors.New("Duplicate field name: " + d.Name))
	}

	// 有等号
	if p.TokenID() == Token_Assign {
		p.NextToken()

		// tag
		fd.Tag = p.Expect(Token_Numeral).ToInt()

	} else { // 没等号自动生成枚举序号

		if len(d.Fields) == 0 {
			fd.AutoTag = 0
		} else {

			// 按前面的序号+1
			fd.AutoTag = d.MaxTag() + 1
		}

	}

	fd.Type = meta.FieldType_Int32

	fd.CommentGroup = p.CommentGroupByLine(nameToken.Line())

	checkField(d, fd)

	d.AddField(fd)

	return
}
