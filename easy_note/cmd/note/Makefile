server:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/easy_note idl/note.thrift # execute in the project root directory
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/easy_note -service demonote -use github.com/cloudwego/biz-demo/easy_note/kitex_gen ../../idl/note.thrift # execute in cmd/note