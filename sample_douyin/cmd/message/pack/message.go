package pack

import (
	"mydouyin/cmd/message/dal/db"
	"mydouyin/kitex_gen/message"
)

// User pack user info
func Message(u *db.Message) *message.Message {
	if u == nil {
		return nil
	}

	return &message.Message{
		Id:         int64(u.ID),
		ToUserId:   int64(u.ToUserID),
		FromUserId: int64(u.FromUserID),
		Content:    u.Content,
		CreateTime: u.CreatedAt.Unix(),
	}
}

func Messages(us []*db.Message) []*message.Message {
	messages := make([]*message.Message, 0)
	for _, m := range us {
		if temp := Message(m); temp != nil {
			messages = append(messages, temp)
		}
	}
	return messages
}
