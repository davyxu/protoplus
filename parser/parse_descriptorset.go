package parser

import (
	"errors"
	"fmt"
	"github.com/davyxu/golexer"
	"github.com/davyxu/protoplus/model"
	"io"
	"io/ioutil"
	"strings"
)

// 解析字符串
func rawParse(ctx *Context, reader io.Reader) (retErr error) {

	data, err := ioutil.ReadAll(reader)

	if err != nil {
		return retErr
	}

	ctx.protoParser = newProtoParser(ctx.SourceName)

	defer golexer.ErrorCatcher(func(err error) {

		retErr = fmt.Errorf("%s %s", ctx.PreTokenPos().String(), err.Error())

	})

	ctx.Lexer().Start(string(data))

	ctx.NextToken()

	for ctx.TokenID() != Token_EOF {

		var d model.Descriptor
		ctx.Descriptor = &d

		if ctx.TokenID() == Token_BracketL {
			d.TagSet = parseTagSet(ctx)
		}

		switch ctx.TokenID() {
		case Token_Struct:
			d.Kind = model.Kind_Struct
			parseObject(ctx)
		case Token_Enum:
			d.Kind = model.Kind_Enum
			parseObject(ctx)
		case Token_Service:
			d.Kind = model.Kind_Service
			parseObject(ctx)
		default:
			panic(errors.New("Unknown token: " + ctx.TokenValue()))
		}

	}

	return nil
}

func checkAndFix(ctx *Context) error {

	for _, d := range ctx.Objects {

		d.DescriptorSet = ctx.DescriptorSet

		for _, fd := range d.Fields {
			fd.Descriptor = d

			if fd.Kind == "" {

				// 将字段中使用的结构体的Kind确认为struct
				findD := ctx.ObjectByName(fd.Type)
				if findD == nil {
					return errors.New(fmt.Sprintf("type not found: %s at %s", fd.Type, ctx.QuerySymbolPosString(fd)))
				}

				fd.Kind = findD.Kind
			}
		}

		for _, sc := range d.SvcCall {
			findREQ := ctx.ObjectByName(sc.RequestName)
			if findREQ == nil {
				return errors.New(fmt.Sprintf("type not found: %s at %s", sc.RequestName, ctx.QuerySymbolPosString(sc)))
			}

			findACK := ctx.ObjectByName(sc.RespondName)
			if findACK == nil {
				return errors.New(fmt.Sprintf("type not found: %s at %s", sc.RespondName, ctx.QuerySymbolPosString(sc)))
			}
		}

		if !d.TagExists("AutoMsgID") && !d.TagExists("MsgID") && (strings.HasSuffix(d.Name, "REQ") || strings.HasSuffix(d.Name, "ACK")) {
			return fmt.Errorf("struct like message but not gen msgid: %s", d.Name)
		}

		if d.Kind == model.Kind_Struct && (d.TagExists("AutoMsgID") || d.TagExists("MsgID")) && (!strings.HasSuffix(d.Name, "REQ") && !strings.HasSuffix(d.Name, "ACK")) {
			return fmt.Errorf("struct gen msgid, but not like a message(REQ/ACK) : %s", d.Name)
		}
	}

	return nil

}
