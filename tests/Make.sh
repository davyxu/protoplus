#!/usr/bin/env bash

set -e

go build -v -o=${GOPATH}/bin/protoplus github.com/davyxu/protoplus/cmd/protoplus

# code.proto为输入文件
${GOPATH}/bin/protoplus -ppgo_out=code_gen.go -package=tests code.proto
${GOPATH}/bin/protoplus -ppgoreg_out=reg_gen.go -package=tests code.proto
${GOPATH}/bin/protoplus -ppcs_out=../example/csharp/Example/ProtoGen.cs -package=Proto code.proto
${GOPATH}/bin/protoplus -ppcsreg_out=../example/csharp/Example/ProtoGenReg.cs -package=Proto code.proto
${GOPATH}/bin/protoplus -pbscheme_out=pb_gen.proto -package=proto code.proto
${GOPATH}/bin/protoplus -route_out=route.json -package=proto code.proto