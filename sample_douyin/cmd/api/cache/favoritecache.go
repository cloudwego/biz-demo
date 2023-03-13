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

package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinfavorite"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type FavoriteCache struct {
	keyName string
	mq      *CommandQueue
}

type FavoriteActionCommand struct {
	UserId     int64
	VideoId    int64
	ActionType string
}

var FC *FavoriteCache

func initFavoriteCache() {
	FC = new(FavoriteCache)
	FC.keyName = "favorite_list"
	FC.mq = NewCommandQueue(context.Background(), "favorite")

	go FC.listen()
}

func (c *FavoriteCache) listen() {
	for {
		msg, err := c.mq.ConsumeMessage()
		if err != nil {
			continue
		}
		var cmd FavoriteActionCommand
		err = json.Unmarshal(msg, &cmd)
		if err != nil {
			log.Printf("json.Unmarshal error %v", err)
			continue
		}
		log.Printf("[********FavoriteCache********] recover command:%v", cmd)
		err = c.execCommand(&cmd)
		if err != nil {
			log.Printf("[********FavoriteCache********] command exec fail, error:%v", err)
			data, _ := json.Marshal(cmd)
			err = c.mq.ProductionMessage(data)
			if err != nil {
				log.Printf("c.mq.ProductionMessage error %v", err)
				continue
			}
			time.Sleep(time.Second * 10)
		} else {
			log.Printf("[********FavoriteCache********] command exec success!!!")
		}
	}
}

func (c *FavoriteCache) execCommand(cmd *FavoriteActionCommand) error {
	resp, err := rpc.FavoriteAction(c.mq.ctx, &douyinfavorite.FavoriteActionRequest{
		UserId:     cmd.UserId,
		VideoId:    cmd.VideoId,
		ActionType: cmd.ActionType,
	})
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func (c *FavoriteCache) CommitFavoriteActionCommand(user_id, video_id int64, action_type string) error {
	cmd := FavoriteActionCommand{
		UserId:     user_id,
		VideoId:    video_id,
		ActionType: action_type,
	}
	data, _ := json.Marshal(cmd)
	return c.mq.ProductionMessage(data)
}
