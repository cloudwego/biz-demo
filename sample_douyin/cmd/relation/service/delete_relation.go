// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package service

import (
	"context"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/relation/dal/db"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/relation"
)

type DeleteRelationService struct {
	ctx context.Context
}

func NewDeleteRelationService(ctx context.Context) *DeleteRelationService {
	return &DeleteRelationService{
		ctx: ctx,
	}
}

func (s *DeleteRelationService) DeleteRelation(req *relation.DeleteRelationRequest) error {
	return db.DeleteRelation(s.ctx, req.FollowId, req.FollowerId)
}
