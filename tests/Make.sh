#!/usr/bin/env bash
CURRDIR=`pwd`
cd ../../../../..
export GOPATH=`pwd`
cd ${CURRDIR}

go build -v -o=${GOPATH}/bin/protoplus github.com/davyxu/protoplus

# code.proto为输入文件
${GOPATH}/bin/protoplus -go_out=code_gen.go -package=tests code.proto
${GOPATH}/bin/protoplus -cs_out=../csharp/Example/ProtoGen.cs -genreg -package=proto code.proto
${GOPATH}/bin/protoplus -pb_out=pb_gen.proto -package=proto code.proto