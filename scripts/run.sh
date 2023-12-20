#!/bin/bash

svcName=${1}

if [ -d "app/${svcName}" ];then
    go run app/${svcName}/*.go
fi