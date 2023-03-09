package cache

import (
	"context"
	"encoding/json"
	"log"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinfavorite"
	"mydouyin/pkg/errno"
	"time"
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
		json.Unmarshal(msg, &cmd)
		log.Printf("[********FavoriteCache********] recover command:%v", cmd)
		err = c.execCommand(&cmd)
		if err != nil {
			log.Printf("[********FavoriteCache********] command exec fail, error:%v", err)
			data, _ := json.Marshal(cmd)
			c.mq.ProductionMessage(data)
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
