package rpc

import (
	"context"
	"mydouyin/kitex_gen/relation"
	"mydouyin/kitex_gen/relation/relationservice"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		consts.RelationServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func CreateRelation(ctx context.Context, req *relation.CreateRelationRequest) (r *relation.BaseResp, err error) {
	return relationClient.CreateRelation(ctx, req)
}

func DeleteRelation(ctx context.Context, req *relation.DeleteRelationRequest) (r *relation.BaseResp, err error) {
	return relationClient.DeleteRelation(ctx, req)
}

func GetFollower(ctx context.Context, req *relation.GetFollowerListRequest) (r *relation.GetFollowerListResponse, err error) {
	return relationClient.GetFollower(ctx, req)
}

func GetFollow(ctx context.Context, req *relation.GetFollowListRequest) (r *relation.GetFollowListResponse, err error) {
	return relationClient.GetFollow(ctx, req)
}

func ValidIfFollowRequest(ctx context.Context, req *relation.ValidIfFollowRequest) (r *relation.ValidIfFollowResponse, err error) {
	return relationClient.ValidIfFollowRequest(ctx, req)
}

// func GetFriendList(ctx context.Context, req *relation.GetFollowerListRequest) ([]*apimodel.FriendUser, error) {
// 	// resp, err := relationClient.GetFollower(ctx, req)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// if resp.BaseResp.StatusCode != 0 {
// 	// 	return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
// 	// }
// 	// if len(resp.FollowerIds) < 1 {
// 	// 	return []*apimodel.FriendUser{}, nil
// 	// }
// 	// ur, err := userClient.MGetUser(ctx, &douyinuser.MGetUserRequest{
// 	// 	UserIds: resp.FollowerIds,
// 	// })
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// if ur.BaseResp.StatusCode != 0 {
// 	// 	return nil, errno.NewErrNo(ur.BaseResp.StatusCode, ur.BaseResp.StatusMessage)
// 	// }
// 	// res := make([]*apimodel.FriendUser, 0, 30)
// 	// for _, rpc_user := range ur.Users {
// 	// 	user := apimodel.PackFriendUser(rpc_user)
// 	// 	r, err := relationClient.ValidIfFollowRequest(ctx, &relation.ValidIfFollowRequest{
// 	// 		FollowId:   user.UserID,
// 	// 		FollowerId: req.FollowId,
// 	// 	})
// 	// 	if err != nil || r.BaseResp.StatusCode != 0 {
// 	// 		continue
// 	// 	}
// 	// 	user.IsFollow = r.IfFollow
// 	// 	res = append(res, user)
// 	// }
// 	return res, nil
// }
func GetFriend(ctx context.Context, req *relation.GetFriendRequest) (r *relation.GetFriendResponse, err error) {
	return relationClient.GetFriend(ctx, req)
}
