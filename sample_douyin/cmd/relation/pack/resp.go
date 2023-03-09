package pack

import (
	"errors"
	"mydouyin/cmd/relation/dal/db"
	"mydouyin/kitex_gen/relation"
	"mydouyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *relation.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *relation.BaseResp {
	return &relation.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg}
}

func Relation2Follow(relations []*db.Relation) []int64 {
	res := make([]int64, 0, len(relations))
	for _, relation := range relations {
		res = append(res, relation.FollowId)
	}
	return res
}

func Relation2Follower(relations []*db.Relation) []int64 {
	res := make([]int64, 0, len(relations))
	for _, relation := range relations {
		res = append(res, relation.FollowerId)
	}
	return res
}
