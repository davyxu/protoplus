package parser

import (
	"errors"
	"fmt"
	"github.com/davyxu/protoplus/meta"
)

func parseStructField(p *protoParser, d *meta.Descriptor) {

	fd := meta.NewFieldDescriptor(d)

	nameToken := p.RawToken()
	// 字段名
	fd.Name = p.Expect(Token_Identifier).Value()

	if _, ok := d.FieldByName[fd.Name]; ok {
		panic(errors.New("Duplicate field name: " + d.Name))
	}

	// 自动生成字段Tag
	if len(d.Fields) == 0 {
		fd.AutoTag = 0
	} else {
		fd.AutoTag = d.MaxTag() + 1
	}

	tp := p.TokenPos()

	var typeName string

	switch p.TokenID() {
	case Token_BracketL: // [  数组类型
		p.NextToken()
		p.Expect(Token_BracketR)
		fd.Repeatd = true
	}

	typeName = p.Expect(Token_Identifier).Value()

	// 根据类型名查找类型及结构体类型

	pf := meta.NewLazyField(typeName, fd, d, tp)

	fd.CommentGroup = p.CommentGroupByLine(nameToken.Line())

	// 尝试首次解析
	if need2Pass, _ := pf.Resolve(1); need2Pass {
		d.File.FileSet.AddLazyField(pf)
	}

	checkField(d, fd)

	d.AddField(fd)

	return
}

func checkField(d *meta.Descriptor, fd *meta.FieldDescriptor) {

	if _, ok := d.FieldByName[fd.Name]; ok {
		panic(errors.New(fmt.Sprintf("Duplicate field name: %s in %s", fd.Name, d.Name)))
	}

	if _, ok := d.FieldByTag[fd.TagNumber()]; ok {
		panic(errors.New(fmt.Sprintf("Duplicate field tag: %d in %s", fd.TagNumber(), d.Name)))
	}
}
