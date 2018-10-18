#!/usr/bin/env bash
CURRDIR=`pwd`
cd ../../../../..
export GOPATH=`pwd`
cd ${CURRDIR}

go build -v -o=${GOPATH}/bin/protoplus github.com/davyxu/protoplus


${GOPATH}/bin/protoplus -go_out=code_gen.go -package=tests code.proto