package parser

import (
	"github.com/davyxu/protoplus/meta"
	"strings"
)

func parseFileTag(p *protoParser, fileD *meta.FileDescriptor, srcName string) {

	p.Expect(Token_FileTag)

	rawTagStr := p.Expect(Token_String).Value()
	for _, tagStr := range strings.Split(rawTagStr, " ") {
		fileD.FileTag = append(fileD.FileTag, strings.TrimSpace(tagStr))
	}

}
