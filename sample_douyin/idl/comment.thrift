namespace go douyincomment

struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct Comment {
    1: i64 comment_id
    2: i64 video
    3: i64 user
    4: string content   
    5: string create_date
}

struct CreateCommentRequest {
    1: i64 video
    2: i64 user
    3: string content (vt.min_size = "1")
    4: string create_date   (vt.min_size = "1")
}

struct DeleteCommentRequest {
    1: i64 comment_id
}

struct GetVideoCommentsRequest {
    1: i64 video
}

struct CreateCommentResponse {
    1: BaseResp base_resp
    2: i64 comment_id
}

struct DeleteCommentResponse {
    1: BaseResp base_resp
}

struct GetVideoCommentsResponse {
    1: BaseResp base_resp
    2: list<Comment> comments
}


service CommentService {
    CreateCommentResponse CreateComment(1: CreateCommentRequest req)
    DeleteCommentResponse DeleteComment(1: DeleteCommentRequest req)
    GetVideoCommentsResponse GetVideoComments(1: GetVideoCommentsRequest req)
}