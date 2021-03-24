module github.com/davyxu/protoplus

go 1.12

require (
	github.com/davyxu/cellnet v4.1.0+incompatible
	github.com/davyxu/golexer v0.0.0-20180314091252-f048a86ae200
	github.com/davyxu/ulexer v0.0.0-20200705151509-86177890ec50
	github.com/stretchr/testify v1.6.1
)

replace github.com/davyxu/ulexer => ../ulexer

replace github.com/davyxu/cellnet => ../cellnet

replace github.com/davyxu/x => ../x
