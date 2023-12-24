.PHONY: all
all: help

.PHONY: help
help:
	@echo "Usage: make gen"

# example: make gen svc=product
.PHONY: gen
gen:
	scripts/gen.sh ${svc}

.PHONY: gen-product-client
gen-product-client:
	cd app/frontend && cwgo client -I ../../idl --type RPC --service product --module github.com/baiyutang/gomall/app/frontend --idl ../../idl/product.proto

.PHONY: gen-checkout-client
gen-checkout-client:
	cd app/checkout && cwgo client -I ../../idl --type RPC --service order --module github.com/baiyutang/gomall/app/checkout --idl ../../idl/order.proto

.PHONY: watch-frontend
watch-frontend:
	cd app/frontend && air

.PHONY: tidy
tidy:
	scripts/tidy.sh

.PHONY: lint
lint:
	gofmt -l -w app

# example: `make run svc=cart`
.PHONY: run
run:
	scripts/run.sh ${svc}

.PHONY: env-start
env-start:
	docker-compose up -d

.PHONY: env-stop
env-stop:
	docker-compose down