#!/usr/bin/env bash
Version=2.0.0

export GOARCH=amd64
BuildSourcePackage="github.com/davyxu/protoplus/build"
BinaryPackage="github.com/davyxu/protoplus/cmd/protoplus"
BinaryName="protoplus"

BuildBinary()
{
  set -e
  TargetDir=bin/"${1}"
  mkdir -p "${TargetDir}"
  export GOOS=${1}
  BuildTime=$(date -R)
  GitCommit=$(git rev-parse HEAD)
  VersionString="-X \"${BuildSourcePackage}.BuildTime=${BuildTime}\" -X \"${BuildSourcePackage}.Version=${Version}\" -X \"${BuildSourcePackage}.GitCommit=${GitCommit}\""

  go build -v -p 4 -o "${TargetDir}"/${BinaryName} -ldflags "${VersionString}" ${BinaryPackage}
  PackageDir=$(pwd)
  cd "${TargetDir}"
  tar zcvf "${PackageDir}"/${BinaryName}-${Version}-"${1}"-x86_64.tar.gz ${BinaryName}
  cd "${PackageDir}"
}



if [[ ${1} == "" ]]; then
  BuildBinary windows
  BuildBinary linux
  BuildBinary darwin
else
  BuildBinary "${1}"
fi