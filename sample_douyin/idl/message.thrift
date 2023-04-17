namespace go message


struct BaseResp {
    1: i64 status_code
    2: string status_message
}

struct Message {
    1: i64 id
    2: i64 to_user_id
    3: i64 from_user_id
    4: string content
    5: i64 create_time
}

struct FirstMessage {
    1: string message (vt.min_size = "1")
    2: i64 msgType (vt.min_size = "1")
    3: i64 friend_id
}



struct CreateMessageRequest {
    1: i64 from_user_id
    2: i64 to_user_id
    3: string content 
}

struct CreateMessageResponse {
    1: i64 id
    2: i64 create_time
    3: BaseResp base_resp
}

struct GetMessageListRequest {
    1: i64 from_user_id
    2: i64 to_user_id
    3: i64 pre_msg_time
}

struct GetMessageListResponse {
    1: BaseResp base_resp
    2: list<Message> message_list
}

struct GetFirstMessageRequest {
    1: i64 id
    2: list<i64> friend_ids
}

struct GetFirstMessageResponse {
    1: list<FirstMessage> first_message_list
    2: BaseResp base_resp
}



service MessageService {
    CreateMessageResponse CreateMessage(1: CreateMessageRequest req)
    GetMessageListResponse GetMessageList(1: GetMessageListRequest req)
    GetFirstMessageResponse GetFirstMessage(1: GetFirstMessageRequest req)
}