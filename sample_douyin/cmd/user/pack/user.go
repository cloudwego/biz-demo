package pack

import (
	"mydouyin/cmd/user/dal/db"
	"mydouyin/kitex_gen/douyinuser"
)

// User pack user info
func User(u *db.User) *douyinuser.User {
	if u == nil {
		return nil
	}

	return &douyinuser.User{
		UserId:        int64(u.ID),
		Username:      u.Username,
		FollowCount:   int64(u.FollowCount),
		FollowerCount: int64(u.FollowerCount),
		TotalFavorited: int64(u.TotalFavorited),
		FavoriteCount: int64(u.FavoriteCount),
		WorkCount: int64(u.WorkCount),
		Avatar: u.Avatar,
		Signature: u.Signature,
		BackgroundImage: u.BackgroundImage,
	}
}

func Users(us []*db.User) []*douyinuser.User {
	users := make([]*douyinuser.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
