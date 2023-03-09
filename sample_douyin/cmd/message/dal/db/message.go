package db

import (
	"context"
	"log"
	"mydouyin/pkg/consts"
	"time"
)

type Message struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	FromUserID int    `json:"from_user_id"`
	ToUserID   int    `json:"to_user_id"`
	Content    string `json:"content"`
}

type FirstMessage struct {
	Message  string `json:"message"`
	MsgType  int64  `json:"msgtype"`
	FirendID int64  `json:"friend_id"`
}

func (u *Message) TableName() string {
	return consts.MessageTableName
}

// CreateMessage create message info
func CreateMessage(ctx context.Context, messages []*Message) (id, create_time int64, err error) {
	log.Println(messages[0].Content)
	result := DB.WithContext(ctx).Create(messages)
	if result.Error != nil {
		err = result.Error
		return
	}
	return int64(messages[0].ID), messages[0].CreatedAt.Unix(), err
}

// QueryUser query list of user info
func QueryMessage(ctx context.Context, fromUserID int64, toUserID int64, preMsgTime int64) ([]*Message, error) {
	res := make([]*Message, 0)
	if err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ? AND created_at > ?", fromUserID, toUserID, time.Unix(preMsgTime, 0).Format("2006-01-02 15:04:05")).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//
func MGetFirstMessage(ctx context.Context, userID int64, friendIDs []int64) ([]*FirstMessage, error) {
	res := make([]*FirstMessage, len(friendIDs))
	tx := DB.WithContext(ctx)
	msg1 := make([]*Message, 0) // userID 发送的最近消息
	msg2 := make([]*Message, 0) // userID 接收的最近消息

	if err := tx.Select("to_user_id,max(created_at) as created_at").Where("from_user_id = ? AND to_user_id in ?", userID, friendIDs).Group("to_user_id").Find(&msg1).Error; err != nil {
		return nil, err
	}
	// log.Println(len(msg1), msg1[0].CreatedAt, msg1[0].ToUserID)
	if err := tx.Select("from_user_id,max(created_at) as created_at").Where("to_user_id = ? AND from_user_id in ?", userID, friendIDs).Group("from_user_id").Find(&msg2).Error; err != nil {
		return nil, err
	}
	// log.Println(msg2[0].CreatedAt, msg2[0].FromUserID)
	// if len(msg1) == 0 && len(msg2) == 0 {
	// 	return
	// }
	msg1_map, msg2_map := make(map[int64]time.Time, 0), make(map[int64]time.Time, 0)
	for _, m := range msg1 {
		msg1_map[int64(m.ToUserID)] = m.CreatedAt
	}

	for _, m := range msg2 {
		msg2_map[int64(m.FromUserID)] = m.CreatedAt
	}

	for _, friend_id := range friendIDs {
		first_message := make([]*FirstMessage, 0, 1)
		time1, ok1 := msg1_map[friend_id]
		time2, ok2 := msg2_map[friend_id]
		if (!ok1) && (!ok2) {
			res = append(res, &FirstMessage{FirendID: friend_id, MsgType: -1})
		} else if ok1 && (!ok2) {
			if err := tx.Model(&Message{}).Select("to_user_id as friend_id,content as message").Where("created_at = ?", time1).Find(&first_message).Error; err != nil {
				return nil, err
			}
			first_message[0].MsgType = 1
			first_message[0].FirendID = friend_id
			res = append(res, append([]*FirstMessage{}, first_message[0])...)
		} else if (!ok1) && ok2 {
			if err := tx.Model(&Message{}).Select("from_user_id as friend_id,content as message").Where("created_at = ?", time2).Find(&first_message).Error; err != nil {
				return nil, err
			}
			first_message[0].MsgType = 0
			first_message[0].FirendID = friend_id
			res = append(res, append([]*FirstMessage{}, first_message[0])...)
		} else {
			if time1.Unix() > time2.Unix() {
				if err := tx.Model(&Message{}).Select("to_user_id as friend_id,content as message").Where("created_at = ?", time1).Find(&first_message).Error; err != nil {
					return nil, err
				}
				first_message[0].MsgType = 1
			} else {
				if err := tx.Model(&Message{}).Select("from_user_id as friend_id,content as message").Where("created_at = ?", time2).Find(&first_message).Error; err != nil {
					return nil, err
				}
				first_message[0].MsgType = 0
			}
			first_message[0].FirendID = friend_id
			res = append(res, append([]*FirstMessage{}, first_message[0])...)
		}
	}
	return res, nil
}
