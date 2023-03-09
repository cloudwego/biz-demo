package pack

import (
	"mydouyin/cmd/favorite/dal/db"
	"mydouyin/kitex_gen/douyinfavorite"
)

// Favorite pack video info
func Favorite(f *douyinfavorite.Favorite) *db.Favorite{
	if f == nil {
		return nil
	}
	return &db.Favorite{
		UserId:  	int64(f.UserId),	
		VideoId: 	int64(f.VideoId),
	}
}

func Favorites(fs []*douyinfavorite.Favorite)  []*db.Favorite{
	favorites := make([]*db.Favorite, 0)
	for _, f := range fs {
		if temp := Favorite(f); temp != nil {
			favorites = append(favorites, temp)
		}
	}
	return favorites
}


func FavoriteToVideoids(favorites []*db.Favorite) []int64 {
	vids := make([]int64, 0)
	for i:= 0 ;i < len(favorites); i++{
		if temp := int64(favorites[i].VideoId); temp != 0 {
			vids = append(vids, temp)
		}
	}
	return vids
}

