module github.com/davyxu/protoplus

go 1.12

require (
	github.com/davyxu/cellnet v4.1.0+incompatible
	github.com/davyxu/golexer v0.1.1-0.20200202091144-a15ddde83f6a
	github.com/davyxu/ulexer v0.0.0-20200705151509-86177890ec50
	github.com/stretchr/testify v1.6.1
)

replace github.com/davyxu/ulexer => ../ulexer

replace github.com/davyxu/cellnet => ../cellnet

replace github.com/davyxu/x => ../x
