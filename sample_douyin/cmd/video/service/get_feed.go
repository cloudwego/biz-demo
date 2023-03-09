package service

import (
	"context"
	"fmt"
	"log"
	"mydouyin/cmd/video/dal/db"
	"mydouyin/cmd/video/pack"
	"mydouyin/kitex_gen/douyinvideo"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type GetFeedService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{ctx: ctx}
}

// GetFeedService.
func (s *GetFeedService) GetFeed(req *douyinvideo.GetFeedRequest) (int64, []*douyinvideo.Video, error) {
	fmt.Printf("video")
	log.Println(req.LatestTime)
	latestTime, err := strconv.Atoi(req.LatestTime)
	if err != nil {
		return time.Now().Unix(), nil, err
	}
	latestTimeStr := time.Unix(int64(latestTime), 0).Format("2006-01-02 15:04:05")
	videos, err := db.GetFeed(s.ctx, latestTimeStr)
	klog.Infof("请求时间：%v", req.LatestTime)
	if err != nil {
		return time.Now().Unix(), nil, err
	}
	if len(videos) < 1 {
		return time.Now().Unix(), nil, nil
	}
	var index int = len(videos) - 1
	return videos[index].CreatedAt.Unix(), pack.Videos(videos), nil
}
