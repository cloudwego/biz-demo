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

.PHONY: run-product
run-product:
	cd app/product && go run .

.PHONY: run-frontend
run-frontend:
	cd app/frontend && go run .

.PHONY: watch-frontend
watch-frontend:
	cd app/frontend && air

.PHONY: gen-cart
gen-cart:
	cd app/cart && cwgo server -I ../../idl --type RPC --service cart --module github.com/baiyutang/gomall/app/cart --idl ../../idl/cart.proto
