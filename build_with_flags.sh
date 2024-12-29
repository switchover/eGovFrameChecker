#!/usr/bin/env bash

TARGET=github.com/switchover/eGovFrameChecker/cmd/ver

BUILD_TIME=$(date +"%Y-%m-%d %H:%M:%S.%2N %Z")

GO_VERSION=$(go version | sed -e "s/go version //g")

COMMIT_HASH=$(git log -1 --pretty=format:%h)

echo Build time : $BUILD_TIME
echo Go ver : $GO_VERSION
echo Commit hash : $COMMIT_HASH

FLAG="-X '$TARGET.BuildTime=$BUILD_TIME'"
FLAG="$FLAG -X '$TARGET.GoVersion=$GO_VERSION'"
FLAG="$FLAG -X '$TARGET.CommitHash=$COMMIT_HASH'"

go build -v -ldflags "$FLAG" "$1" "$2" "$3" "$4" "$5" "$6" "$7" "$8" "$9"
