#!/bin/bash

export OLD_DYLD=$DYLD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=$GOPATH/clibs/lib
go build -ldflags -extldflags=-L$DYLD_LIBRARY_PATH
export DYLD_LIBRARY_PATH=$OLD_DYLD


