#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=frontend
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}