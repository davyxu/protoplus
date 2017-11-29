package codegen

import "testing"

func TestCodeGen(t *testing.T) {

	codegen := NewCodeGen("test")
	codegen.ParseTemplate("xx.pp").FormatGoCode().WriteOutputFile("out.go")

}
