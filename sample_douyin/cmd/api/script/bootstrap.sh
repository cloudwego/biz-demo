
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=hello.api
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
