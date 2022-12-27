server:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/easy_note idl/user.thrift # execute in the project root directory
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/easy_note -service demouser -use github.com/cloudwego/biz-demo/easy_note/kitex_gen ../../idl/user.thrift # execute in cmd/user