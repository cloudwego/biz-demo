package main

import (
	"context"

	"mydouyin/cmd/relation/pack"
	"mydouyin/cmd/relation/service"

	relation "mydouyin/kitex_gen/relation"
	"mydouyin/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// CreateRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CreateRelation(ctx context.Context, req *relation.CreateRelationRequest) (resp *relation.BaseResp, err error) {
	// TODO: Your code here...
	if err = req.IsValid(); err != nil {
		resp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewCreateRelationService(ctx).CreateRelation(req)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) DeleteRelation(ctx context.Context, req *relation.DeleteRelationRequest) (resp *relation.BaseResp, err error) {
	// TODO: Your code here...
	if err = req.IsValid(); err != nil {
		resp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewDeleteRelationService(ctx).DeleteRelation(req)
	if err != nil {
		resp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollow(ctx context.Context, req *relation.GetFollowListRequest) (resp *relation.GetFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.GetFollowListResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	resp.FollowIds, err = service.NewGetFollowService(ctx).GetFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollower(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.GetFollowerListResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	resp.FollowerIds, err = service.NewGetFollowerService(ctx).GetFollower(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// ValidIfFollowRequest implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ValidIfFollowRequest(ctx context.Context, req *relation.ValidIfFollowRequest) (resp *relation.ValidIfFollowResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.ValidIfFollowResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	resp.IfFollow, err = service.NewValidIfFollowService(ctx).ValidIfFollowFollower(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFriend implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriend(ctx context.Context, req *relation.GetFriendRequest) (resp *relation.GetFriendResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.GetFriendResponse)
	resp.FriendIds, err = service.NewGetFriendService(ctx).GetFriend(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
