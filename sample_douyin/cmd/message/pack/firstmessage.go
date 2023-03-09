package pack

import (
	"mydouyin/cmd/message/dal/db"
	"mydouyin/kitex_gen/message"
)

// User pack user info
func FirstMessage(u *db.FirstMessage) *message.FirstMessage {
	if u == nil {
		return nil
	}

	return &message.FirstMessage{
		Message:  u.Message,
		MsgType:  int64(u.MsgType),
		FriendId: int64(u.FirendID),
	}
}

func FirstMessages(us []*db.FirstMessage) []*message.FirstMessage {
	firstmessages := make([]*message.FirstMessage, 0)
	for _, m := range us {
		if temp := FirstMessage(m); temp != nil {
			firstmessages = append(firstmessages, temp)
		}
	}
	return firstmessages
}
