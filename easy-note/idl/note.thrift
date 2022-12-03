// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

namespace go notedemo

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
    1: string title (api.form="title", api.vd="len($) > 0")
    2: string content (api.form="content", api.vd="len($) > 0")
    3: i64 user_id
}

struct CreateNoteResponse {
    1: BaseResp base_resp
}

struct DeleteNoteRequest {
    1: i64 note_id (api.path="note_id")
    2: i64 user_id
}

struct DeleteNoteResponse {
    1: BaseResp base_resp
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

struct MGetNoteRequest {
    1: list<i64> note_ids
}

struct MGetNoteResponse {
    1: list<Note> notes
    2: BaseResp base_resp
}

service NoteService {
    CreateNoteResponse CreateNote(1: CreateNoteRequest req) (api.post="/v2/note")
    DeleteNoteResponse DeleteNote(1: DeleteNoteRequest req) (api.delete="/v2/note/:note_id")
    UpdateNoteResponse UpdateNote(1: UpdateNoteRequest req) (api.put="/v2/note/:note_id")
    QueryNoteResponse QueryNote(1: QueryNoteRequest req) (api.get="/v2/note/query")
    MGetNoteResponse MGetNote(1: MGetNoteRequest req)
}
