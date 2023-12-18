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

.PHONY: gen-order
gen-order:
	cd app/order && cwgo server -I ../../idl --type RPC --service order --module github.com/baiyutang/gomall/app/order --idl ../../idl/order.proto

.PHONY: gen-payment
gen-payment:
	cd app/payment && cwgo server -I ../../idl --type RPC --service payment --module github.com/baiyutang/gomall/app/payment --idl ../../idl/payment.proto


.PHONY: run-frontend
run-frontend:
	cd app/frontend && go run .

.PHONY: run-order
run-order:
	cd app/order && go run .

.PHONY: run-payment
run-payment:
	cd app/payment && go run .