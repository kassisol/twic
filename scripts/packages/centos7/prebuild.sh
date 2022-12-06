#!/usr/bin/env bash

ROOTDIR=$(dirname $0)/../../..
cd $(dirname $0)

if [ -d "build" ]; then
	rm -rf build
fi
mkdir -p build

cp ${ROOTDIR}/bin/twic build/

export GO111MODULE=auto
go run ${ROOTDIR}/gen/man/genman.go
cp -r /tmp/twic/man build/

go run ${ROOTDIR}/gen/shellcompletion/genshellcompletion.go
cp -r /tmp/twic/shellcompletion build/
