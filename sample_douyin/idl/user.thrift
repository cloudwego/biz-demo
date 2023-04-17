namespace go douyinuser

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
    QueryErrCode               = 10005
}

struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct User {
    1: i64 user_id
    2: string username
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: i64 favorite_count
    7: i64 work_count
    8: i64 total_favorited
    9: string background_image
    10: string avatar
    11: string signature
}

struct CreateUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
    3: string background_image
    4: string avatar
    5: string signature
}

struct CreateUserResponse {
    1: BaseResp base_resp
    2: i64 user_id
    3: string token
}

struct CheckUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct CheckUserResponse {
    1: i64 user_id
    2: BaseResp base_resp
}

struct MGetUserRequest {
    1: list<i64> user_ids (vt.min_size = "1")
}

struct MGetUserResponse {
    1: list<User> users
    2: BaseResp base_resp
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    CheckUserResponse CheckUser(1: CheckUserRequest req)
    MGetUserResponse MGetUser(1: MGetUserRequest req)
}