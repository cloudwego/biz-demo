#!/usr/bin/env bash

# find all go mod path
# returns an array contains mod path
function util::find_modules() {
	find . -not \( \
		\( \
		-path './output' \
		-o -path './.git' \
		-o -path '*/third_party/*' \
		-o -path '*/vendor/*' \
		\) -prune \
		\) -name 'go.mod' -print0 | xargs -0 -I {} dirname {}
}