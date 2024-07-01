export ROOT_MOD=github.com/cloudwego/biz-demo/gomall
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --module github.com/cloudwego/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	cd demo/demo_proto && golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/order_page.proto

.PHONY: gen-user
gen-user: 
	@cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto

.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto


.PHONY: gen-cart
gen-cart: 
	@cd rpc_gen && cwgo client --type RPC --service cart --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module ${ROOT_MOD}/app/cart --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto


.PHONY: gen-checkout
gen-checkout: 
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module ${ROOT_MOD}/app/checkout --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto


.PHONY: gen-payment
gen-payment: 
	@cd rpc_gen && cwgo client --type RPC --service payment --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module ${ROOT_MOD}/app/payment --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto

.PHONY: gen-order
gen-order: 
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module ${ROOT_MOD}/app/order --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-email
gen-email: 
	@cd rpc_gen && cwgo client --type RPC --service email --module ${ROOT_MOD}/rpc_gen  -I ../idl  --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module ${ROOT_MOD}/app/email --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/email.proto

.PHONY: tidy-all
tidy-all: 
	cd app/cart && go mod tidy
	cd app/checkout && go mod tidy
	cd app/email && go mod tidy
	cd app/frontend && go mod tidy
	cd app/order && go mod tidy
	cd app/payment && go mod tidy
	cd app/product && go mod tidy
	cd app/user && go mod tidy
	cd common && go mod tidy

.PHONY: build-frontend
build-frontend:
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .

.PHONY: build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .

.PHONY: build-all
build-all:
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .
	docker build -f ./deploy/Dockerfile.svc -t cart:${v} --build-arg SVC=cart .
	docker build -f ./deploy/Dockerfile.svc -t checkout:${v} --build-arg SVC=checkout .
	docker build -f ./deploy/Dockerfile.svc -t email:${v} --build-arg SVC=email .
	docker build -f ./deploy/Dockerfile.svc -t order:${v} --build-arg SVC=order .
	docker build -f ./deploy/Dockerfile.svc -t payment:${v} --build-arg SVC=payment .
	docker build -f ./deploy/Dockerfile.svc -t product:${v} --build-arg SVC=product .
	docker build -f ./deploy/Dockerfile.svc -t user:${v} --build-arg SVC=user .

.PHONY: kind-load-image
kind-load-image:
	kind load docker-image --name gomall-dev \
	frontend:${v} cart:${v} checkout:${v} email:${v} order:${v} payment:${v} product:${v} user:${v}