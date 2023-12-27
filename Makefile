.PHONY: all
all: help

default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 - .]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: gen
gen: # gen client code of {svc}. example: make gen svc=product
	scripts/gen.sh ${svc}

.PHONY: gen-product-client
gen-product-client:
	cd app/frontend && cwgo client -I ../../idl --type RPC --service product --module github.com/baiyutang/gomall/app/frontend --idl ../../idl/product.proto

.PHONY: gen-checkout-client
gen-checkout-client:
	cd app/checkout && cwgo client -I ../../idl --type RPC --service payment --module github.com/baiyutang/gomall/app/checkout --idl ../../idl/payment.proto

.PHONY: gen-order-client
gen-order-client:
	cd app/frontend && cwgo client -I ../../idl --type RPC --service order --module github.com/baiyutang/gomall/app/frontend --idl ../../idl/order.proto

.PHONY: watch-frontend
watch-frontend:
	cd app/frontend && air

.PHONY: tidy
tidy: # run `go mod tidy` for all go mudule
	scripts/tidy.sh

.PHONY: lint
lint: # run `gofmt` for all go mudule
	gofmt -l -w app

# example: `make run svc=cart`
.PHONY: run
run: # run {svc} server. example: make run svc=product
	scripts/run.sh ${svc}

.PHONY: env-start
env-start:  # launch all middleware software as the docker 
	docker-compose up -d

.PHONY: env-stop
env-stop: # stop all docker
	docker-compose down

.PHONY: open.gomall
open.gomall: # open `gomall` website in the default browser
	open "http://localhost:8080/"

.PHONY: open.consul
open.consul: # open `consul ui` in the default browser
	open "http://localhost:8500/ui/"

.PHONY: open.jaeger
open.jaeger: # open `jaeger ui` in the default browser
	open "http://localhost:16686/search"

.PHONY: open.prometheus
open.prometheus: # open `prometheus ui` in the default browser
	open "http://localhost:9090"

