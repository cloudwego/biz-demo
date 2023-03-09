namespace go douyinapi

struct BaseResp {
    1: i64 status_code
    2: string status_msg
}

struct User {
    1: i64 user_id
    2: string username
    3: i64 follow_count
    4: i64 follower_count
}

struct CreateUserRequest {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct CreateUserResponse {
    1: BaseResp base_resp
    2: i64 user_id
    3: string token
}

struct CheckUserRequest {
    1: string username (api.query="username", api.vd="len($) > 0")
    2: string password (api.query="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: BaseResp base_resp
    2: i64 user_id
    3: string token
}

struct GetUserRequest {
    1: string user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct GetUserResponse {
    1: BaseResp base_resp
    2: User user
}

service ApiService {
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/douyin/user/register/")
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/douyin/user/login/")
    GetUserResponse GetUser(1: GetUserRequest req) (api.get="/douyin/user/")
}

