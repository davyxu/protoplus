#!/usr/bin/env bash
Version=2.0.0


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

if [[ "$Platform" == "" ]]; then
	Platform=$(go env GOHOSTOS)
fi

#export GOARCH=amd64

BuildSourcePackage="github.com/davyxu/protoplus/build"
BinaryPackage="github.com/davyxu/protoplus/cmd/protoplus"
BinaryName="protoplus"

BuildBinary()
{
  TargetDir=bin/"${1}"
  mkdir -p "${TargetDir}"
  export GOOS=${1}

  BuildTime=$(date -R)
  GitCommit=$(git rev-parse HEAD)
  VersionString="-X \"${BuildSourcePackage}.BuildTime=${BuildTime}\" -X \"${BuildSourcePackage}.Version=${Version}\" -X \"${BuildSourcePackage}.GitCommit=${GitCommit}\""

  go build -p 4 -o "${TargetDir}"/${BinaryName} -ldflags "${VersionString}" ${BinaryPackage}
  echo "${TargetDir}"/${BinaryName}
}


BuildBinary "${Platform}"