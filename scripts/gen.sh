#!/bin/bash

svcName=${1}

cd app/${svcName}
cwgo client -I ../../idl --type RPC --service ${svcName} --module github.com/baiyutang/gomall/app/${svcName} --idl ../../idl/${svcName}.proto
