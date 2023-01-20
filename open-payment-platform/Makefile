MODULE = github.com/cloudwego/biz-demo/open-payment-platform

# start the environment of demo
.PHONY: start
start:
	docker-compose up -d

# stop the environment of demo
.PHONY: stop
stop:
	docker-compose down

# generate client code by IDL.
.PHONY: gen
gen:
	kitex -module $(MODULE) idl/payment.thrift
	kitex -module $(MODULE) idl/common.thrift

# run the gateway
.PHONY: gateway
gateway:
	  go run hertz-gateway/*.go
