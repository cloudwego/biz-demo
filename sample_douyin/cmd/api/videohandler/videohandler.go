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

package videohandler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/cache"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type VideoHandel struct {
	CommandQueue *cache.CommandQueue
	client       *oss.Client
	bucket       *oss.Bucket
}

type commandState int

const (
	begin commandState = iota
	finishCoverUpLoad
	allFinish
)

type Command struct {
	VideoName string
	UserID    int64
	Title     string
	State     commandState
}

var VH *VideoHandel

func Init() {
	VH = new(VideoHandel)
	VH.CommandQueue = cache.NewCommandQueue(context.Background(), "upload_command")
	// 初始化OSS
	var err error
	VH.client, err = oss.New(consts.Endpoint, consts.AKID, consts.AKS)
	if err != nil {
		panic(fmt.Sprintf("init videohandler error:%v", err))
	}

	// 填写存储空间名称
	VH.bucket, err = VH.client.Bucket(consts.Bucket)
	if err != nil {
		panic(fmt.Sprintf("init videohandler error:%v", err))
	}
	go VH.listen()
}

func (vh *VideoHandel) CommitCommand(VideoName string, UserID int64, Title string) {
	data, _ := json.Marshal(Command{
		VideoName: VideoName,
		UserID:    UserID,
		Title:     Title,
		State:     begin,
	})
	err := vh.CommandQueue.ProductionMessage(data)
	for err != nil {
		time.Sleep(1 * time.Second)
		err = vh.CommandQueue.ProductionMessage(data)
	}
}

func (vh *VideoHandel) listen() {
	for {
		msg, err := vh.CommandQueue.ConsumeMessage()
		if err != nil {
			continue
		}
		var cmd Command
		err = json.Unmarshal(msg, &cmd)
		if err != nil {
			log.Printf("json.Unmarshal error %v", err)
			continue
		}
		log.Printf("[********VideoHandler********] recover command:%v", cmd)
		err = vh.execCommand(&cmd)
		if err != nil || cmd.State != allFinish {
			log.Printf("[********VideoHandler********] command exec fail, error:%v", err)
			data, _ := json.Marshal(cmd)
			err = vh.CommandQueue.ProductionMessage(data)
			if err != nil {
				log.Printf("c.mq.ProductionMessage error %v", err)
				continue
			}
		} else {
			log.Printf("[********VideoHandler********] command exec success!!!")
		}
	}
}

// 执行指令，视频上传成功后service提交指令给videohandler，handler只执行生成封面、入库等操作
func (vh *VideoHandel) execCommand(cmd *Command) error {
	// 执行指令，生成封面
	// 截图格式
	cover_name := "cover/" + time.Now().Format("2006-01-02-15:04:05") + ".jpg"
	switch cmd.State {
	case begin:
		style := "video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast"
		// 根据视频名直接获取截图url
		process := fmt.Sprintf("%s|sys/saveas,o_%v,b_%v", style, base64.URLEncoding.EncodeToString([]byte(cover_name)), base64.URLEncoding.EncodeToString([]byte(consts.Bucket)))
		result, err := VH.bucket.ProcessObject(cmd.VideoName, process)
		if err != nil {
			return err
		}
		log.Println(result.Status)
		cmd.State = finishCoverUpLoad
		fallthrough
	case finishCoverUpLoad:
		// 调rpc写库
		resp, err := rpc.CreateVideo(context.Background(), &douyinvideo.CreateVideoRequest{
			Author:   cmd.UserID,
			PlayUrl:  cmd.VideoName,
			CoverUrl: cover_name,
			Title:    cmd.Title,
		})
		if err != nil {
			return err
		}
		if resp.BaseResp.StatusCode != 0 {
			return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		}
		cmd.State = allFinish
	}
	return nil
}

func (vh *VideoHandel) UpLoadVideo(data *multipart.FileHeader) (videoName string, err error) {
	// 获取文件流
	// 视频文件object名称
	var filepoint multipart.File
	filepoint, err = data.Open()
	if err != nil {
		return
	}
	defer filepoint.Close()
	// 上传视频
	videoName = "videos/" + time.Now().Format("2006-01-02-15:04:05") + ".mp4"
	err = vh.bucket.PutObject(videoName, filepoint)
	return
}

func (vh *VideoHandel) UpLoadVideoV0(data *multipart.FileHeader, userID int64, title string) (err error) {
	var filepoint multipart.File
	filepoint, err = data.Open()
	if err != nil {
		return
	}
	defer filepoint.Close()
	// 上传视频
	videoName := "videos/" + time.Now().Format("2006-01-02-15:04:05") + ".mp4"
	err = vh.bucket.PutObject(videoName, filepoint)

	cover_name := "cover/" + time.Now().Format("2006-01-02-15:04:05") + ".jpg"
	style := "video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast"
	// 根据视频名直接获取截图url
	process := fmt.Sprintf("%s|sys/saveas,o_%v,b_%v", style, base64.URLEncoding.EncodeToString([]byte(cover_name)), base64.URLEncoding.EncodeToString([]byte(consts.Bucket)))
	result, err := VH.bucket.ProcessObject(videoName, process)
	log.Println(result.Status)
	if err != nil {
		return err
	}
	// 调rpc写库
	resp, err := rpc.CreateVideo(context.Background(), &douyinvideo.CreateVideoRequest{
		Author:   userID,
		PlayUrl:  videoName,
		CoverUrl: cover_name,
		Title:    title,
	})
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}
