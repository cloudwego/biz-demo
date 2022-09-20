TOOLS_SHELL="./hack/tools.sh"

.PHONY: test
test:
	@${TOOLS_SHELL} test
	@echo "go test finished"



.PHONY: vet
vet:
	@${TOOLS_SHELL} vet
	@echo "vet check finished"