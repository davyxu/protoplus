#!/usr/bin/env bash

set -e

go build -v -o=${GOPATH}/bin/protoplus github.com/davyxu/protoplus/cmd/protoplus

# code.proto为输入文件
ProtoPlusBin=${GOPATH}/bin/protoplus

# 原生输出请使用pb
#${ProtoPlusBin} -ppgo_out=code_gen.go -package=tests filelist.proto
#${ProtoPlusBin} -ppcs_out=../example/csharp/Example/ProtoGen.cs -package=Proto filelist.proto

${ProtoPlusBin} -ppgoreg_out=reg_gen.go -package=tests filelist.proto
${ProtoPlusBin} -ppcsreg_out=../example/csharp/Example/ProtoGenReg.cs -package=Proto filelist.proto
${ProtoPlusBin} -pbscheme_out=pb_gen.proto -package=proto filelist.proto
${ProtoPlusBin} -route_out=route.json -package=proto filelist.proto