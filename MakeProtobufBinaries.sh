#!/usr/bin/env bash

Platform=$1


if [[ "$Platform" == "" ]]; then
	Platform=$(go env GOHOSTOS)
fi


function ReportError()
{
	if [ $? -ne 0 ] ; then read -rsp $'Errors occurred...\n' ; fi
}

trap ReportError EXIT
set -e

BuildGoPlugin()
{
  TargetDir=bin/"${1}"
  mkdir -p "${TargetDir}"
  export GOOS=${1}

  # Download & compile protoc-gen-go
  go get google.golang.org/protobuf/cmd/protoc-gen-go
  go build -o "${TargetDir}"/protoc-gen-go google.golang.org/protobuf/cmd/protoc-gen-go
  echo "${TargetDir}"/protoc-gen-go
}




DownloadProtoc()
{
  #export http_proxy="http://127.0.0.1:7890"
  #export https_proxy="http://127.0.0.1:7890"

  wget https://github.com/protocolbuffers/protobuf/releases/download/v21.6/protoc-21.6-win64.zip
}


BuildGoPlugin "${Platform}"