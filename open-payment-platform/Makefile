MODULE = github.com/cloudwego/biz-demo/open-payment-platform

.PHONY: prepare
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: gen
gen:
	kitex -module $(MODULE) idl/payment.thrift

gateway:
	  go run hertz-gateway/*.go
