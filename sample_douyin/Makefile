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

install_hz_latest:
	go install github.com/cloudwego/hertz/cmd/hz@latest

hertz_gen_model:
	hz model --idl=idl/api.thrift --mod=github.com/cloudwego/biz-demo/sample_douyin --model_dir=hertz_gen

hertz_gen_client:
	hz client --idl=idl/api.thrift --base_domain=http://127.0.0.1:8080 --client_dir=api_request --mod=github.com/cloudwego/biz-demo/sample_douyin --model_dir=hertz_gen