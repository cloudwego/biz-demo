namespace go douyinvideo

struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct Video {
    1: i64 video_id
    2: i64 author
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
    9: string upload_time
}

struct GetFeedRequest {
    1: string latest_time
    2: i64    user_id
}

struct GetFeedResponse {
    1: BaseResp base_resp
    2: i64 next_time
    3: list<Video> video_list
}

struct CreateVideoRequest {
    1: i64 author
    2: string play_url  (vt.min_size = "1")
    3: string cover_url (vt.min_size = "1")
    8: string title     (vt.min_size = "1")
}

struct CreateVideoResponse {
    1: BaseResp base_resp
    2: list<i64> video_ids
}

struct GetListRequest {
    1: i64 user_id
}

struct GetListResponse {
    1: BaseResp base_resp
    2: list<Video> video_list
}

struct MGetVideoRequest {
    1: list<i64> video_ids (vt.min_size = "1")
}

struct MGetVideoResponse {
    1: list<Video> videos
    2: BaseResp base_resp
}

struct DeleteVideoRequest {
    1: i64 video_id
}

struct DeleteVideoResponse {
    1: BaseResp base_resp
}

struct GetTimeVideosRequest {
    1: BaseResp base_resp
    2: string start
    3: string end
}

struct GetTimeVideosResponse {
    1: BaseResp base_resp
    2: list<Video> video_list
}

service VideoService {
    CreateVideoResponse CreateVideo(1: CreateVideoRequest req)
    GetFeedResponse GetFeed(1: GetFeedRequest req)
    GetListResponse GetList(1: GetListRequest req)
    MGetVideoResponse MGetVideoUser(1: MGetVideoRequest req)
    DeleteVideoResponse DeleteVideo(1: DeleteVideoRequest req)
    GetTimeVideosResponse GetTimeVideos(1: GetTimeVideosRequest req)
}
