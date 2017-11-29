package codegen

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type CodeGen struct {
	name string

	buffer bytes.Buffer

	err error
}

func (self *CodeGen) Data() []byte {
	return self.buffer.Bytes()
}

func (self *CodeGen) Error() error {
	return self.err
}

func (self *CodeGen) ParseTemplate(textTemplate string, modelData interface{}) *CodeGen {

	var tpl *template.Template
	tpl, self.err = template.New(self.name).Parse(textTemplate)
	if self.err != nil {
		return self
	}

	self.err = tpl.Execute(&self.buffer, modelData)
	if self.err != nil {
		return self
	}

	return self
}

func (self *CodeGen) FormatGoCode() *CodeGen {

	fset := token.NewFileSet()

	ast, err := parser.ParseFile(fset, "", self.buffer, parser.ParseComments)
	if err != nil {
		self.err = err
		return self
	}

	self.buffer.Reset()

	err = (&printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}).Fprint(&self.buffer, fset, ast)
	if err != nil {
		self.err = err
		return self
	}

	return self
}

func (self *CodeGen) WriteOutputFile(outputFileName string) *CodeGen {

	os.MkdirAll(filepath.Dir(outputFileName), 666)

	self.err = ioutil.WriteFile(outputFileName, self.buffer.Bytes(), 0666)

	if self.err != nil {
		return self
	}

	return self

}

func NewCodeGen(name string) *CodeGen {

	return &CodeGen{
		name: name,
	}
}
