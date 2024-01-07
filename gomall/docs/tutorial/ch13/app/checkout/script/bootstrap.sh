#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/checkout"
exec "$CURDIR/bin/checkout"
