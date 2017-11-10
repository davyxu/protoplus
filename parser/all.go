package parser

import (
	"errors"
	"fmt"
	"github.com/davyxu/golexer"
	"github.com/davyxu/protoplus/meta"
	"io/ioutil"
)

func ParseFile(fileName string) (*meta.FileDescriptorSet, error) {

	fileset := meta.NewFileDescriptorSet()

	fileD := meta.NewFileDescriptor()
	fileD.FileName = fileName
	fileset.AddFile(fileD)

	err := rawPaseFile(fileD, fileName)
	if err != nil {
		return nil, err
	}

	return fileset, fileset.ResolveAll()
}

func ParseFileList(fileset *meta.FileDescriptorSet, filelist []string) (string, error) {

	for _, filename := range filelist {

		fileD := meta.NewFileDescriptor()
		fileD.FileName = filename
		fileset.AddFile(fileD)

		if err := rawPaseFile(fileD, filename); err != nil {
			return filename, err
		}

	}

	return "", fileset.ResolveAll()

}

// 从文件解析
func rawPaseFile(fileD *meta.FileDescriptor, fileName string) error {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return rawParse(fileD, string(data), fileName)
}

// 解析字符串
func rawParse(fileD *meta.FileDescriptor, data string, srcName string) (retErr error) {

	p := newProtoParser(srcName)

	defer golexer.ErrorCatcher(func(err error) {

		retErr = fmt.Errorf("%s %s", p.PreTokenPos().String(), err.Error())

	})

	p.Lexer().Start(data)

	p.NextToken()

	for p.TokenID() != Token_EOF {

		switch p.TokenID() {
		case Token_Struct:
			parseStruct(p, fileD, srcName)
		case Token_Enum:
			parseEnum(p, fileD, srcName)
		case Token_FileTag:
			parseFileTag(p, fileD, srcName)
		default:
			panic(errors.New("Unknown token: " + p.TokenValue()))
		}

	}

	return nil
}
