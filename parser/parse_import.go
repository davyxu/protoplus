package parser

import (
	"os"
)

func parseImport(ctx *Context) error {

	ctx.NextToken()

	fileName := ctx.Expect(Token_String).Value()

	return parseFile(ctx, fileName)
}

func parseFile(ctx *Context, fileName string) error {
	if !ctx.AddSource(fileName) {
		return nil
	}

	file, err := os.Open(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	newCtx := ctx.Clone(fileName)

	return rawParse(newCtx, file)
}
