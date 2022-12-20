namespace go demonote

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
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

struct CreateNoteRequest {
    1: string title (vt.min_size = "1")
    2: string content (vt.min_size = "1")
    3: i64 user_id (vt.gt = "0")
}

struct CreateNoteResponse {
    1: BaseResp base_resp
}

struct DeleteNoteRequest {
    1: i64 note_id (vt.gt = "0")
    2: i64 user_id
}

struct DeleteNoteResponse {
    1: BaseResp base_resp
}

struct UpdateNoteRequest {
    1: i64 note_id (vt.gt = "0")
    2: i64 user_id
    3: optional string title
    4: optional string content
}

struct UpdateNoteResponse {
    1: BaseResp base_resp
}

struct QueryNoteRequest {
    1: i64 user_id (vt.gt = "0")
    2: optional string search_key
    3: i64 offset (vt.ge = "0")
    4: i64 limit (vt.ge = "0")
}

struct QueryNoteResponse {
    1: list<Note> notes
    2: i64 total
    3: BaseResp base_resp
}

struct MGetNoteRequest {
    1: list<i64> note_ids (vt.min_size = "1")
}

struct MGetNoteResponse {
    1: list<Note> notes
    2: BaseResp base_resp
}

service NoteService {
    CreateNoteResponse CreateNote(1: CreateNoteRequest req)
    DeleteNoteResponse DeleteNote(1: DeleteNoteRequest req)
    UpdateNoteResponse UpdateNote(1: UpdateNoteRequest req)
    QueryNoteResponse QueryNote(1: QueryNoteRequest req)
    MGetNoteResponse MGetNote(1: MGetNoteRequest req)
}
