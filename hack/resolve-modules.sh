#!/usr/bin/env bash

# This is used by the linter action.
# Recursively finds all directories with a go.mod file and creates
# a GitHub Actions JSON output option.

set -o errexit

HOME=$(
	cd "$(dirname "${BASH_SOURCE[0]}")" &&
		cd .. &&
		pwd
)

source "${HOME}/hack/util.sh"
all_modules=$(util::find_modules)
PATHS=""
for mod in $all_modules; do
		PATHS+=$(printf '{"workdir":"%s"},' ${mod})
done

echo "::set-output name=matrix::{\"include\":[${PATHS%?}]}"