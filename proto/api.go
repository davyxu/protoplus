package proto

import (
	"github.com/davyxu/protoplus/wire"
	"github.com/davyxu/ulexer"
)

type Struct = wire.Struct

func Marshal(msg Struct) ([]byte, error) {

	l := msg.Size()

	data := make([]byte, 0, l)

	buffer := wire.NewBuffer(data)

	err := msg.Marshal(buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil

}

func Size(msg Struct) int {

	return msg.Size()
}

func Unmarshal(data []byte, msg Struct) (err error) {

	buffer := wire.NewBuffer(data)

	return wire.UnmarshalStructObject(buffer, msg)
}

var (
	defaultTextMarshaler = TextMarshaler{}
	compactTextMarshaler = TextMarshaler{Compact: true, IgnoreDefault: true, CompactBytesSize: 50}
)

func MarshalTextString(obj interface{}) string {
	return defaultTextMarshaler.Text(obj)
}

func CompactTextString(obj interface{}) string { return compactTextMarshaler.Text(obj) }

func UnmarshalText(s string, obj interface{}) error {

	vObj := safeValueOf(obj)
	tObj := safeTypeOf(obj)

	lex := ulexer.NewLexer(s)

	// 读之前, 清掉空白
	lex.SetPreHook(func(lex *ulexer.Lexer) *ulexer.Token {

		lex.Read(ulexer.WhiteSpace())
		return nil
	})

	err := lex.Run(func(lex *ulexer.Lexer) {

		parseStruct(lex, tObj, vObj, "")

	})

	return err
}
