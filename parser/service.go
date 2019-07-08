package parser

import (
	"errors"
)

func parseSvcCallField(ctx *Context) {

	// 注释
	nameToken := ctx.RawToken()

	if ctx.TokenID() == Token_RPC {
		ctx.NextToken()
		ctx.ServiceCall.IsRPC = true
	}

	// 字段名
	ctx.ServiceCall.Name = ctx.Expect(Token_Identifier).Value()

	if ctx.CallNameExists(ctx.ServiceCall.Name) {
		panic(errors.New("Duplicate svc call name: " + ctx.ServiceCall.Name))
	}

	tp := ctx.TokenPos()

	if ctx.TokenID() == Token_ParenL {
		ctx.NextToken()

		ctx.ServiceCall.RequestName = ctx.Expect(Token_Identifier).Value()
		ctx.Expect(Token_ParenR)
		ctx.ServiceCall.RespondName = ctx.Expect(Token_Identifier).Value()
	} else {
		ctx.ServiceCall.RequestName = ctx.ServiceCall.Name + "REQ"
		ctx.ServiceCall.RespondName = ctx.ServiceCall.Name + "ACK"
	}

	ctx.ServiceCall.Comment = ctx.CommentGroupByLine(nameToken.Line())

	ctx.AddSymbol(ctx.ServiceCall, tp)

	ctx.AddSvcCall(ctx.ServiceCall)

	return
}
