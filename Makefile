.PHONY: all
all: help

.PHONY: help
help:
	@echo "Usage: make gen"

.PHONY: gen-product
gen-product:
	cd app/product && cwgo server -I ../../idl --type RPC --service product --module github.com/baiyutang/gomall/app/product --idl ../../idl/product.proto

.PHONY: gen-product-client
gen-product-client:
	cd app/frontend && cwgo client -I ../../idl --type RPC --service product --module github.com/baiyutang/gomall/app/frontend --idl ../../idl/product.proto

.PHONY: watch-frontend
watch-frontend:
	cd app/frontend && air

.PHONY: gen-cart
gen-cart:
	cd app/cart && cwgo server -I ../../idl --type RPC --service cart --module github.com/baiyutang/gomall/app/cart --idl ../../idl/cart.proto

.PHONY: gen-order
gen-order:
	cd app/order && cwgo server -I ../../idl --type RPC --service order --module github.com/baiyutang/gomall/app/order --idl ../../idl/order.proto

.PHONY: gen-payment
gen-payment:
	cd app/payment && cwgo server -I ../../idl --type RPC --service payment --module github.com/baiyutang/gomall/app/payment --idl ../../idl/payment.proto

.PHONY: tidy
tidy:
	scripts/tidy.sh

.PHONY: lint
lint:
	scripts/lint.sh

# example: `make run svc=cart`
.PHONY: run
run:
	scripts/run.sh ${svc}