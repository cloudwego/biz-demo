namespace go demoapi

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 user_id
    2: string username
    3: string avatar
}

struct Note {
    1: i64 note_id
    2: i64 user_id
    3: string username
    4: string user_avatar
    5: string title
    6: string content
    7: i64 create_time
}

struct CreateUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CreateUserResponse {
    1: BaseResp base_resp
}

struct CheckUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: BaseResp base_resp
}

struct CreateNoteRequest {
    1: string title (api.vd="len($) > 0")
    2: string content (api.vd="len($) > 0")
    3: i64 user_id
}

struct CreateNoteResponse {
    1: BaseResp base_resp
}

struct QueryNoteRequest {
    1: i64 user_id
    2: optional string search_key (api.query="search_key", api.vd="len($) > 0")
    3: i64 offset (api.query="offset", api.vd="len($) >= 0")
    4: i64 limit (api.query="limit", api.vd="len($) >= 0")
}

struct QueryNoteResponse {
    1: list<Note> notes
    2: i64 total
    3: BaseResp base_resp
}

struct UpdateNoteRequest {
    1: i64 note_id (api.path="note_id")
    2: i64 user_id
    3: optional string title (api.form="title", api.vd="len($) > 0")
    4: optional string content (api.form="content", api.vd="len($) > 0")
}

struct UpdateNoteResponse {
    1: BaseResp base_resp
}

struct DeleteNoteRequest {
    1: i64 note_id (api.path="note_id")
    2: i64 user_id
}

struct DeleteNoteResponse {
    1: BaseResp base_resp
}

service ApiService {
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/v1/user/register")
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/v1/user/login")
    CreateNoteResponse CreateNote(1: CreateNoteRequest req) (api.post="/v1/note")
    QueryNoteResponse QueryNote(1: QueryNoteRequest req) (api.get="/v1/note/query")
    UpdateNoteResponse UpdateNote(1: UpdateNoteRequest req) (api.put="/v1/note/:note_id")
    DeleteNoteResponse DeleteNote(1: DeleteNoteRequest req) (api.delete="/v1/note/:note_id")
}