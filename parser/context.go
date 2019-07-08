package parser

import (
	"github.com/davyxu/golexer"
	"github.com/davyxu/protoplus/model"
)

type Context struct {
	SourceName string

	*protoParser

	*model.DescriptorSet

	*model.Descriptor

	*model.FieldDescriptor

	*model.ServiceCall

	symbolPos map[interface{}]golexer.TokenPos
}

func (self *Context) QuerySymbolPosString(v interface{}) string {

	if s, ok := self.symbolPos[v]; ok {
		return s.String()
	}

	return ""
}

func (self *Context) AddSymbol(v interface{}, pos golexer.TokenPos) {

	self.symbolPos[v] = pos
}

func newContext() *Context {

	return &Context{
		symbolPos: make(map[interface{}]golexer.TokenPos),
	}
}
