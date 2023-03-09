install_kitex_latest:
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest 

install_thriftgo:
	go install github.com/cloudwego/thriftgo

kitex_gen_user:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/sample_douyin idl/user.thrift # execute in the project root directory

kitex_gen_video:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/sample_douyin idl/video.thrift # execute in the project root directory

kitex_gen_relation:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/sample_douyin idl/relation.thrift # execute in the project root directory

kitex_gen_message:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/sample_douyin idl/message.thrift # execute in the project root directory

kitex_gen_favorite:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/sample_douyin idl/favorite.thrift # execute in the project root directory

kitex_gen_comment:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/sample_douyin idl/comment.thrift # execute in the project root directory
