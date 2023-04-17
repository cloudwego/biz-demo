namespace go douyinfavorite

struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct Favorite {
    1: i64 favorite_id
    2: i64 user_id
    3: i64 video_id
}

struct FavoriteActionRequest {
    1: i64 user_id
    2: i64 video_id
    3: string action_type
}

struct FavoriteActionResponse {
    1: BaseResp base_resp
}

struct GetListRequest {
    1: i64 user_id
}

struct GetListResponse {
    1: BaseResp base_resp
    2: list<i64> video_ids
}

struct GetIsFavoriteRequest{
    1: list<Favorite> favorite_list
}

struct GetIsFavoriteResponse {
    1: list<bool> is_favorites
    2: BaseResp base_resp
}

service FavoriteService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    GetListResponse GetList(1: GetListRequest req)
    GetIsFavoriteResponse GetIsFavorite(1: GetIsFavoriteRequest req)
}


