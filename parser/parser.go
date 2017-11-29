package parser

import (
	"bytes"
	"errors"

	"github.com/davyxu/golexer"
	"github.com/davyxu/protoplus/meta"
)

// 自定义的token id
const (
	Token_EOF = iota
	Token_Unknown
	Token_LineEnd
	Token_Numeral
	Token_String
	Token_WhiteSpace
	Token_Identifier
	Token_UnixComment
	Token_CStyleComment
	Token_Colon       // :
	Token_ParenL      // (
	Token_ParenR      // )
	Token_CurlyBraceL // {
	Token_CurlyBraceR // }
	Token_BracketL    // [
	Token_BracketR    // ]
	Token_Dot         // .
	Token_Enum        // enum
	Token_Struct      // struct
	Token_FileTag     // fileTag
	Token_Assign      // =
)

type protoParser struct {
	*golexer.Parser

	commentsByLine map[int]string
}

func (self *protoParser) Expect(id int) golexer.Token {

	if self.Parser.TokenID() != id {
		panic(errors.New("Expect " + self.Lexer().MatcherString(id)))
	}

	t := self.RawToken()

	self.NextToken()

	return t
}

func (self *protoParser) NextToken() {

	for {
		self.Parser.NextToken()

		switch self.TokenID() {

		case Token_UnixComment,
			Token_CStyleComment:
			self.commentsByLine[self.RawToken().Line()] = self.TokenValue()
		default:
			return
		}
	}

}

func (self *protoParser) CommentGroupByLine(line int) *meta.CommentGroup {

	cg := meta.NewCommentGroup()

	if comment, ok := self.commentsByLine[line]; ok {
		cg.AddLineComment(comment)
		cg.Trailing = comment
	}

	var buff bytes.Buffer

	start := line - 1
	var end int

	for i := line - 1; i >= 1; i-- {

		if _, ok := self.commentsByLine[i]; !ok {
			end = i
			break
		}

	}

	for i := end; i <= start; i++ {

		comment, _ := self.commentsByLine[i]

		if buff.Len() > 0 {
			buff.WriteString("\n")
		}

		cg.AddLineComment(comment)

		buff.WriteString(comment)
	}

	cg.Leading = buff.String()

	if cg.Leading == "" && cg.Trailing == "" {
		return nil
	}

	return cg
}

func newProtoParser(srcName string) *protoParser {

	l := golexer.NewLexer()

	// 匹配顺序从高到低

	l.AddMatcher(golexer.NewNumeralMatcher(Token_Numeral))
	l.AddMatcher(golexer.NewStringMatcher(Token_String))

	l.AddIgnoreMatcher(golexer.NewWhiteSpaceMatcher(Token_WhiteSpace))
	l.AddIgnoreMatcher(golexer.NewLineEndMatcher(Token_LineEnd))
	l.AddMatcher(golexer.NewUnixStyleCommentMatcher(Token_UnixComment))
	l.AddMatcher(golexer.NewCStyleCommentMatcher(Token_CStyleComment))

	l.AddMatcher(golexer.NewSignMatcher(Token_CurlyBraceL, "{"))
	l.AddMatcher(golexer.NewSignMatcher(Token_CurlyBraceR, "}"))
	l.AddMatcher(golexer.NewSignMatcher(Token_ParenL, "("))
	l.AddMatcher(golexer.NewSignMatcher(Token_ParenR, ")"))
	l.AddMatcher(golexer.NewSignMatcher(Token_BracketL, "["))
	l.AddMatcher(golexer.NewSignMatcher(Token_BracketR, "]"))
	l.AddMatcher(golexer.NewSignMatcher(Token_Dot, "."))
	l.AddMatcher(golexer.NewSignMatcher(Token_Assign, "="))
	l.AddMatcher(golexer.NewSignMatcher(Token_Colon, ":"))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_Enum, "enum"))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_Struct, "struct"))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_FileTag, "filetag"))

	l.AddMatcher(golexer.NewIdentifierMatcher(Token_Identifier))

	l.AddMatcher(golexer.NewUnknownMatcher(Token_Unknown))

	return &protoParser{
		Parser:         golexer.NewParser(l, srcName),
		commentsByLine: make(map[int]string),
	}
}
