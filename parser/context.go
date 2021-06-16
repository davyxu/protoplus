package parser

import (
	"github.com/davyxu/golexer"
	"github.com/davyxu/protoplus/model"
	"strings"
)

type Context struct {

	// 每个文件对应的属性
	SourceName string

	*protoParser

	*model.DescriptorSet

	// 全局属性
	symbolPos map[interface{}]golexer.TokenPos

	sourceByName map[string]struct{}
}

func (self *Context) AddSource(sourceName string) bool {
	lowerName := strings.ToLower(sourceName)
	if _, ok := self.sourceByName[lowerName]; ok {
		return false
	}

	self.sourceByName[lowerName] = struct{}{}

	return true
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

func (self *Context) Clone(srcName string) *Context {

	return &Context{
		SourceName:    srcName,
		symbolPos:     self.symbolPos,
		sourceByName:  self.sourceByName,
		DescriptorSet: self.DescriptorSet,
	}
}

func newContext() *Context {

	return &Context{
		symbolPos:    map[interface{}]golexer.TokenPos{},
		sourceByName: map[string]struct{}{},
	}
}
