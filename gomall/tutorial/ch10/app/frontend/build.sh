#!/bin/bash
RUN_NAME=frontend
mkdir -p output/bin output/conf
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh
cp -r conf/* output/conf
go build -o output/bin/${RUN_NAME}