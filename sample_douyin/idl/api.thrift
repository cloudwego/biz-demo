namespace go douyinapi

/**############ Model ###############**/

struct User {
    1: i64      id
    2: string   name
    3: i64      follow_count
    4: i64      follower_count
    5: bool     is_follow
	6: string   avatar
	7: string   background_image
	8: string   signature
	9: i64      total_favorited
	10: i64     work_count
	11: i64     favorite_count

}

struct FriendUser {
	1: i64      id
    2: string   name
    3: i64      follow_count
    4: i64      follower_count
    5: bool     is_follow
	6: string   avatar
	7: string   background_image
	8: string   signature
	9: i64      total_favorited
	10: i64     work_count
	11: i64     favorite_count
	12: string  message
	13: i64     msgType
}

struct Video {
    1: i64      id
	2: User     author
	3: string   play_url
	4: string   cover_url
	5: i64      favorite_count
	6: i64      comment_count
	7: bool     is_favorite
	8: string   title
	9: string   upload_time
}

struct Comment {
	1: i64      id
	2: User     user
	3: string   content
	4: string   create_date
}

struct Favorite {
	1: i64 id
	2: i64 user_id
	3: i64 video_id
}

struct Message {
	1: i64      id
	2: i64      to_user_id
	3: i64      from_user_id
	4: string   content
	5: i64      create_time
}

/**############ Request ###############**/

struct RegistUserRequest {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct CheckUserRequest {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct GetUserRequest {
    1: string user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct GetFeedRequest {
    1: string latest_time (api.query="latest_time")
	2: string token (api.query="token")         
}

struct PublishVideoRequest {
	1: list<byte> data  (api.form="data")
	2: string token     (api.form="token")      
	3: string title     (api.form="title")
}

struct GetPublishListRequest {
	1: string token (api.query="token")
	2: string user_id (api.query="user_id")
}

struct RelationActionRequest {
	1: string token (api.query="token")
	2: string to_user_id (api.query="to_user_id")
	3: string action_type (api.query="action_type")
}

struct FollowAndFollowerListRequest {
	1: string user_id (api.query="user_id")
	2: string token (api.query="token")
}

struct FriendListRequest {
	1: i64 user_id (api.query="user_id")
	2: string token (api.query="token")
}

struct FavoriteActionRequest {
	1: string token     (api.query="token")
	2: string video_id  (api.query="video_id")
	3: string action_type (api.query="action_type")
}

struct GetFavoriteListRequest {
	1: string token (api.query="token")
	2: string user_id (api.query="user_id")
}

struct CommentActionRequest {
	1: string token (api.query="token")
	2: string video_id (api.query="video_id")
	3: string action_type (api.query="action_type")
	4: string comment_text (api.query="comment_text")
	5: string comment_id (api.query="comment_id")
}

struct CommentListRequest {
	1: string token (api.query="token")
	2: string video_id (api.query="video_id")
}

struct MessageChatRequest {
	1: string token (api.query="token")
	2: i64  to_user_id (api.query="to_user_id")
    3: i64  pre_msg_time (api.query="pre_msg_time")
}

struct MessageActionRequest {
	1: string token (api.query="token")
	2: i64  to_user_id (api.query="to_user_id")
    3: i32  action_type (api.query="action_type")
    4: string content (api.query="content")
}

/**############ Response ###############**/

struct RegistUserResponse {
    1: i64 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct CheckUserResponse {
    1: i64 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct GetUserResponse {
    1: i64 status_code
    2: string status_msg
    3: User user
}

struct GetFeedResponse {
	1: i64 status_code
    2: string status_msg
    3: i64   next_time
	4: list<Video> video_list
}

struct PublishVideoResponse {
	1: i64 status_code
    2: string status_msg
}

struct GetPublishListResponse {
	1: i64 status_code
    2: string status_msg
    3: list<Video> video_list
}

struct CommentActionResponse {
	1: i64 status_code
    2: string status_msg
    3: Comment comment
}

struct CommentListResponse {
	1: i64 status_code
    2: string status_msg	
    3: list<Comment> comment_list
}

struct RelationActionResponse {
	1: i64 status_code
    2: string status_msg	
}

struct FollowAndFollowerListResponse {
    1: i64 status_code
    2: string status_msg
    3: list<User> user_list
}

struct FriendListResponse {
    1: i64 status_code
    2: string status_msg
    3: list<FriendUser> user_list
}

struct FavoriteActionResponse {
    1: i64 status_code
    2: string status_msg
}

struct GetFavoriteListResponse {
	1: i64 status_code
    2: string status_msg
    3: list<Video> video_list
}

struct MessageChatResponse {
	1: i64 status_code
    2: string status_msg
    3: list<Message> message_list
}

struct MessageActionResponse {
    1: i64 status_code
    2: string status_msg
}

service ApiService {

    /**############ 基础接口 ###############**/
    RegistUserResponse RegistUser(1: RegistUserRequest req) (api.post="/douyin/user/register/")
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/douyin/user/login/")
    GetUserResponse GetUser(1: GetUserRequest req) (api.get="/douyin/user/")
    GetFeedResponse GetFeed(1: GetFeedRequest req) (api.get="/douyin/feed/")
    GetPublishListResponse GetPublishList(1: GetPublishListRequest req) (api.get="/douyin/publish/list/")
    PublishVideoResponse PublishVideo(1: PublishVideoRequest req) (api.post="/douyin/publish/action/")

    /**############ 互动接口 ###############**/
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    GetFavoriteListResponse GetFavoriteList(1: GetFavoriteListRequest req) (api.get="/douyin/favorite/list/")
    CommentActionResponse CommentAction(1: CommentActionRequest req) (api.post="/douyin/comment/action/")
    CommentListResponse CommentList(1: CommentListRequest req) (api.get="/douyin/comment/list/")

    /**############ 社交接口 ###############**/
    RelationActionResponse RelationAction(1: RelationActionRequest req) (api.post="/douyin/relation/action/")
    FollowAndFollowerListResponse FollowList(1: FollowAndFollowerListRequest req) (api.get="/douyin/relation/follow/list/")
    FollowAndFollowerListResponse FollowerList(1: FollowAndFollowerListRequest req) (api.get="/douyin/relation/follower/list/")
    FriendListResponse FriendList(1: FriendListRequest req) (api.get="/douyin/relation/friend/list/")
    MessageChatResponse MessageChat(1: MessageChatRequest req) (api.get="/douyin/message/chat/")
    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="/douyin/message/action/")
}

