package parser

import "github.com/davyxu/protoplus/model"

func parseTagSet(ctx *Context) (ts model.TagSet) {

	// [
	if ctx.TokenID() != Token_BracketL {
		return
	}

	ctx.NextToken()

	for ctx.TokenID() != Token_BracketR {

		var tag model.Tag
		tag.Key = ctx.Expect(Token_Identifier).Value()

		ctx.Expect(Token_Colon).Value()

		tag.Value = ctx.TokenValue()

		ts.AddTag(tag)

		ctx.NextToken()
	}

	ctx.Expect(Token_BracketR)

	return
}
