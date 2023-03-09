package cache

import (
	"context"
)

//消息队列，调用ProductionCommand将指令结构体放入redis的list中，另起一个线程调用ConsumeCommand接收消息，反序列化指令结构体交给自定义的handler函数执行
//在videoHandel中有使用
type CommandQueue struct {
	ListName string
	ctx      context.Context
}

func NewCommandQueue(ctx context.Context, listName string) *CommandQueue {
	return &CommandQueue{
		ListName: listName,
		ctx:      ctx,
	}
}

func (cq *CommandQueue) ProductionMessage(command []byte) error {
	_, err := redisClient.LPush(cq.ctx, cq.ListName, command).Result()
	return err
}

func (cq *CommandQueue) ConsumeMessage() ([]byte, error) {
	item, err := redisClient.BRPop(cq.ctx, 0, cq.ListName).Result()
	if err != nil {
		return nil, err
	}
	return []byte(item[1]), err
}
