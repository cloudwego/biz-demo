// Copyright 2022 CloudWeGo Authors
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

	"github.com/cloudwego/biz-demo/easy_note/cmd/note/dal/db"
	"github.com/cloudwego/biz-demo/easy_note/cmd/note/pack"
	"github.com/cloudwego/biz-demo/easy_note/cmd/note/rpc"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demonote"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demouser"
)

type QueryNoteService struct {
	ctx context.Context
}

// NewQueryNoteService new QueryNoteService
func NewQueryNoteService(ctx context.Context) *QueryNoteService {
	return &QueryNoteService{ctx: ctx}
}

// QueryNoteService query list of note info
func (s *QueryNoteService) QueryNoteService(req *demonote.QueryNoteRequest) ([]*demonote.Note, int64, error) {
	noteModels, total, err := db.QueryNote(s.ctx, req.UserId, req.SearchKey, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, 0, err
	}
	userMap, err := rpc.MGetUser(s.ctx, &demouser.MGetUserRequest{UserIds: []int64{req.UserId}})
	if err != nil {
		return nil, 0, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].Username = u.Username
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, total, nil
}
