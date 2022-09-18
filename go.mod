module github.com/davyxu/protoplus

go 1.12

require (
	// 测试用
	github.com/davyxu/golexer v0.1.1-0.20200202091144-a15ddde83f6a
	github.com/davyxu/ulexer v0.0.0-20200713054812-c9bb8db3521f
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/stretchr/testify v1.7.0
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/davyxu/ulexer => ../ulexer
