namespace go relation

struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct CreateRelationRequest {
    1: i64 follow_id
    2: i64 follower_id
}

struct DeleteRelationRequest {
    1: i64 follow_id
    2: i64 follower_id 
}

struct GetFollowListRequest {
    1: i64 follower_id
}

struct GetFollowerListRequest {
    1: i64 follow_id
}

struct GetFollowListResponse {
    1: list<i64> follow_ids
    2: BaseResp base_resp
}

struct GetFollowerListResponse {
    1: list<i64> follower_ids
    2: BaseResp base_resp
}

struct GetFriendRequest {
    1: i64 me_id
}

struct GetFriendResponse {
    1: list<i64> friend_ids
    2: BaseResp base_resp
}

struct ValidIfFollowRequest {
    1: i64 follow_id
    2: i64 follower_id
}

struct ValidIfFollowResponse {
    1: bool if_follow
    2: BaseResp base_resp
}

service RelationService {
    BaseResp CreateRelation(1: CreateRelationRequest req)
    BaseResp DeleteRelation(1: DeleteRelationRequest req)
    GetFollowListResponse GetFollow(1: GetFollowListRequest req)
    GetFollowerListResponse GetFollower(1: GetFollowerListRequest req)
    GetFriendResponse GetFriend(1: GetFriendRequest req)
    ValidIfFollowResponse ValidIfFollowRequest(1: ValidIfFollowRequest req)
}