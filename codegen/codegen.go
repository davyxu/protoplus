package codegen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type CodeGen struct {
	name string

	buffer bytes.Buffer

	err error

	tpl *template.Template
}

func (self *CodeGen) Data() []byte {
	return self.buffer.Bytes()
}

func (self *CodeGen) Error() error {
	return self.err
}

func (self *CodeGen) ParseTemplate(textTemplate string, modelData interface{}) *CodeGen {

	if self.err != nil {
		return self
	}

	_, self.err = self.tpl.Parse(textTemplate)
	if self.err != nil {
		return self
	}

	self.err = self.tpl.Execute(&self.buffer, modelData)
	if self.err != nil {
		return self
	}

	return self
}

func (self *CodeGen) RegisterTemplateFunc(funcMap template.FuncMap) *CodeGen {
	if self.err != nil {
		return self
	}

	self.tpl.Funcs(funcMap)
	return self
}

func (self *CodeGen) FormatGoCode() *CodeGen {

	if self.err != nil {
		return self
	}

	fset := token.NewFileSet()

	ast, err := parser.ParseFile(fset, "", self.buffer.Bytes(), parser.ParseComments)
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

func (self *CodeGen) Code() string {

	reader := bufio.NewReader(strings.NewReader(string(self.Data())))

	var sb strings.Builder
	line := 0
	for {
		lineStr, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line++
		sb.WriteString(fmt.Sprintf("%d:	%s", line, lineStr))
	}

	return sb.String()

}

func (self *CodeGen) WriteBytes(data *[]byte) *CodeGen {
	if self.err != nil {
		return self
	}

	*data = self.buffer.Bytes()

	return self
}

func (self *CodeGen) WriteOutputFile(outputFileName string) *CodeGen {

	if self.err != nil {
		return self
	}

	os.MkdirAll(filepath.Dir(outputFileName), 0755)

	self.err = ioutil.WriteFile(outputFileName, self.buffer.Bytes(), 0666)

	if self.err != nil {
		return self
	}

	fmt.Printf("%s\n", outputFileName)

	return self

}

func NewCodeGen(name string) *CodeGen {

	self := &CodeGen{
		tpl: template.New(name),
	}

	return self
}
