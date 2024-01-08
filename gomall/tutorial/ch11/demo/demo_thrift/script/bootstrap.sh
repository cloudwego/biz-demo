#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/demo_thrift"
exec "$CURDIR/bin/demo_thrift"