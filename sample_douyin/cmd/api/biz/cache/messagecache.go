package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"log"
	"math"
	"mydouyin/cmd/api/biz/apimodel"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/message"
	"mydouyin/kitex_gen/relation"
	"mydouyin/pkg/errno"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type MessageCache struct {
	keyName         string
	mq              *CommandQueue
	lasted_time_map sync.Map
}

type CreateMessageCommand struct {
	FromUserId int64
	ToUserId   int64
	Content    string
}

var MC *MessageCache

func initMessageCache() {
	MC = new(MessageCache)
	MC.keyName = "message_list"
	MC.mq = NewCommandQueue(context.Background(), "message")

	go MC.listen()
}

// 获取from_user_id最新发给to_uer_id的消息
func (c *MessageCache) GetLastedMsg(from_user_id, to_user_id int64) (lastedTime int64, err error) {
	var msg interface{}
	var ok bool = false
	for !ok {
		msg, ok = c.lasted_time_map.Load(strconv.FormatInt(from_user_id, 10) + strconv.FormatInt(to_user_id, 10))
		time.Sleep(50 * time.Millisecond)
	}
	lastedTime = msg.(*apimodel.Message).CreateTime
	return
}

func (c *MessageCache) PushLastedMsg(msg *apimodel.Message) {
	c.lasted_time_map.Store(strconv.FormatInt(msg.FromUserId, 10)+strconv.FormatInt(msg.ToUserId, 10), msg)
}

func (c *MessageCache) listen() {
	for {
		msg, err := c.mq.ConsumeMessage()
		if err != nil {
			continue
		}
		var cmd CreateMessageCommand
		json.Unmarshal(msg, &cmd)
		log.Printf("[********MessageCache********] recover command:%v", cmd)
		err = c.execCommand(&cmd)
		if err != nil {
			log.Printf("[********MessageCache********] command exec fail, error:%v", err)
			data, _ := json.Marshal(cmd)
			c.mq.ProductionMessage(data)
			time.Sleep(time.Second * 10)
		} else {
			log.Printf("[********MessageCache********] command exec success!!!")
		}
	}
}

func (c *MessageCache) execCommand(cmd *CreateMessageCommand) error {
	resp, err := rpc.CreateMessage(c.mq.ctx, &message.CreateMessageRequest{
		FromUserId: cmd.FromUserId,
		ToUserId:   cmd.ToUserId,
		Content:    cmd.Content,
	})
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	//上传成功，将新消息写入缓存
	msg := apimodel.Message{
		ID:         resp.Id,
		CreateTime: resp.CreateTime,
		ToUserId:   cmd.ToUserId,
		FromUserId: cmd.FromUserId,
		Content:    cmd.Content,
	}
	err = c.SaveMessage([]*apimodel.Message{&msg})
	c.PushLastedMsg(&msg)
	return err
}

func (c *MessageCache) getMsgFromSet(from_id, to_id int64) (msg *apimodel.Message, err error) {
	key := md5.Sum([]byte(strconv.FormatUint(c.hash_key(from_id, to_id), 10) + c.keyName))
	exist, err := redisClient.Exists(c.mq.ctx, string(key[:])).Result()
	if err != nil {
		return nil, err
	}
	if exist != 1 {
		return nil, errors.New("key not exist")
	}
	result, err := redisClient.Get(c.mq.ctx, string(key[:])).Result()
	defer redisClient.Expire(c.mq.ctx, string(key[:]), 12*time.Hour)
	if err != nil {
		return nil, err
	}
	msg = new(apimodel.Message)
	json.Unmarshal([]byte(result), msg)
	return msg, nil
}

func (c *MessageCache) setMsgToSet(msg *apimodel.Message) error {
	key := md5.Sum([]byte(strconv.FormatUint(c.hash_key(msg.FromUserId, msg.ToUserId), 10) + c.keyName))
	val, _ := json.Marshal(msg)
	result, err := redisClient.Set(c.mq.ctx, string(key[:]), val, 12*time.Hour).Result()
	if err != nil || result != "OK" {
		return err
	}
	return nil
}

func (c *MessageCache) hash_key(id1, id2 int64) uint64 {
	if id1 > id2 {
		return uint64(id2 | id1<<32)
	} else {
		return uint64(id1 | (id2 << 32))
	}
}

func (c *MessageCache) CommitCreateMessageCommand(from_user_id, to_user_id int64, content string) error {
	cmd := CreateMessageCommand{
		FromUserId: from_user_id,
		ToUserId:   to_user_id,
		Content:    content,
	}
	data, _ := json.Marshal(cmd)
	return c.mq.ProductionMessage(data)
}

func (c *MessageCache) CommitCreateMessageCommandV0(from_user_id, to_user_id int64, content string) error {
	cmd := CreateMessageCommand{
		FromUserId: from_user_id,
		ToUserId:   to_user_id,
		Content:    content,
	}
	return c.execCommand(&cmd)
	// return c.mq.ProductionMessage(data)
}

func (c *MessageCache) GetFirstMessage(me int64, friendIds []int64) (frist_msg_list []*apimodel.FristMessage) {
	frist_msg_list = make([]*apimodel.FristMessage, 0, len(friendIds))
	for _, friendId := range friendIds {
		msg, err := c.getMsgFromSet(friendId, me)
		if err != nil {
			frist_msg_list = append(frist_msg_list, &apimodel.FristMessage{
				FriendId: friendId,
				MsgType:  -1,
			})
		} else {
			if msg.FromUserId == me {
				frist_msg_list = append(frist_msg_list, &apimodel.FristMessage{
					FriendId: friendId,
					Content:  msg.Content,
					MsgType:  1,
				})
			} else {
				frist_msg_list = append(frist_msg_list, &apimodel.FristMessage{
					FriendId: friendId,
					Content:  msg.Content,
					MsgType:  0,
				})
			}
		}
	}
	return
}

// 设置最新消息
func (c *MessageCache) SetFirstMessage(msg *apimodel.Message) (err error) {
	return c.setMsgToSet(msg)
}

func (c *MessageCache) SaveMessage(messages []*apimodel.Message) error {
	frist_msg := new(apimodel.Message)
	for i := 0; i < len(messages); i++ {
		if messages[i].CreateTime > frist_msg.CreateTime {
			frist_msg = messages[i]
		}
		msgKey := strconv.FormatUint(c.hash_key(messages[i].FromUserId, messages[i].ToUserId), 10) + c.keyName
		message, _ := json.Marshal(messages[i])
		// msgKey := md5.Sum([]byte(strconv.FormatInt(messages[i].ID, 10) + strconv.FormatInt(messages[i].ToUserId, 10) + c.keyname))
		_, err := redisClient.ZAdd(c.mq.ctx, msgKey, &redis.Z{Score: float64(messages[i].CreateTime), Member: message}).Result()
		if err != nil {
			return err
		}
		_, err1 := redisClient.Expire(c.mq.ctx, msgKey, time.Hour*12).Result()
		if err1 != nil {
			return err1
		}
	}
	//更新fristmessage缓存
	return c.SetFirstMessage(frist_msg)
}

func (c *MessageCache) InitMessageFromDB(fromUserID int64) error {
	resp, err := rpc.GetFriend(c.mq.ctx, &relation.GetFriendRequest{
		MeId: fromUserID,
	})
	if err != nil {
		return err
	}
	friendIds := resp.FriendIds
	for i := 0; i < len(friendIds); i++ {
		msgkey := strconv.FormatUint(c.hash_key(fromUserID, friendIds[i]), 10) + c.keyName
		ex, err := redisClient.Exists(c.mq.ctx, msgkey).Result()
		if err != nil {
			return err
		}
		if ex == 0 {
			messageList := make([]*apimodel.Message, 0)
			// 从数据库拉取所有我发的消息
			resp, err := rpc.GetMessageList(c.mq.ctx, &message.GetMessageListRequest{
				FromUserId: fromUserID,
				ToUserId:   friendIds[i],
				PreMsgTime: 0,
			})
			if err != nil {
				return err
			}
			messageList = append(messageList, apimodel.PackMessages(resp.MessageList)...)
			// 从数据库拉取所有发给我的消息
			resp, err = rpc.GetMessageList(c.mq.ctx, &message.GetMessageListRequest{
				FromUserId: friendIds[i],
				ToUserId:   fromUserID,
				PreMsgTime: time.Now().Unix(),
			})
			if err != nil {
				return err
			}
			messageList = append(messageList, apimodel.PackMessages(resp.MessageList)...)
			// 存入缓存
			err = c.SaveMessage(messageList)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (c *MessageCache) GetMessage(fromUserID int64, toUserID int64, preMsgTime int64) ([]*apimodel.Message, bool, error) {
	messageList := make([]*apimodel.Message, 0)
	msgkey := strconv.FormatUint(c.hash_key(fromUserID, toUserID), 10) + c.keyName
	ex, err := redisClient.Exists(c.mq.ctx, msgkey).Result()
	if err != nil {
		return nil, false, err
	}
	if ex == 0 {
		return nil, false, nil
	}
	_, err = redisClient.Expire(c.mq.ctx, msgkey, time.Hour*12).Result()
	if err != nil {
		return nil, false, err
	}
	values, err := redisClient.ZRangeByScore(c.mq.ctx, msgkey, &redis.ZRangeBy{
		Min: strconv.FormatInt(preMsgTime, 10),
		Max: strconv.FormatInt(math.MaxInt64, 10),
	}).Result()
	// log.Println(values)
	if err != nil {
		return nil, false, err
	}

	for i := 0; i < len(values); i++ {
		message := new(apimodel.Message)
		err = json.Unmarshal([]byte(values[i]), &message)
		if err != nil {
			continue
		}
		//由于要返回>preMsgTime的记录，因此需要剔除掉preMsgTime处的记录
		if message.CreateTime == preMsgTime {
			continue
		}
		messageList = append(messageList, message)
	}
	return messageList, true, nil
}
