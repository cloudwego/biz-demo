MODULE = github.com/cloudwego/biz-demo/open-payment-platform
#NACOS_VERSION ?= 2.0.3

.PHONY: prepare
start:
	#docker pull nacos/nacos-server:$(NACOS_VERSION)
	#docker run --rm --name nacos-quick -e MODE=standalone -p 8848:8848 -p 9848:9848 -d nacos/nacos-server:$(NACOS_VERSION)
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: gen
gen:
	kitex -module $(MODULE) idl/payment.thrift

.PHONY: try
try:
	kitex -module $(MODULE) idl/trysvc.thrift

gateway:
	  go run hertz-gateway/*.go
