use serde_json::json;
use volo::context::Context;
use volo_http::context::ServerContext;
use volo_http::http;
use volo_http::http::StatusCode;
use volo_http::request::ServerRequest;
use volo_http::response::ServerResponse;
use volo_http::json::Json;
use volo_http::server::extract::FromRequest;
use volo_http::server::response::IntoResponse;
use volo_http::http::Response as http_response;
use std::convert::Infallible;
use crate::cmd::api::mw::JWT_MIDDLEWARE;
use crate::cmd::api::rpc::create_note as rpc_create_note;
use crate::cmd::api::rpc;
use crate::pkg::errno;
use crate::pkg::consts;
use crate::pkg::errno::ErrNo;
use serde::Serialize;
use serde::Deserialize;

// CheckUser函数用于处理用户登录请求，通过JwtMiddleware的LoginHandler方法进行JWT认证和授权。
pub async fn check_user(
    cx: &mut ServerContext,
    req: ServerRequest,
) -> Result<ServerResponse, Infallible> {
    // 调用JwtMiddleware的LoginHandler方法，处理用户登录请求并进行JWT认证和授权
    JWT_MIDDLEWARE.login_handler(cx, req).await
}

#[derive(Serialize, Deserialize)]
pub struct Response {
    code: i64,
    message: String,
    data: String,
}

// CreateNote函数的主要功能是创建一条新笔记记录，并将结果返回给客户端。
pub async fn create_note(
    cx: &mut ServerContext,
    req: ServerRequest,
) -> Result<ServerResponse, Infallible> {
    // 1. 从请求上下文中绑定并验证请求数据
    let (parts, body) = req.into_parts();
    let data = match Json::<NoteData>::from_request(cx, parts, body).await {
        Ok(data) => data,
        Err(err) => {
            // 2. 如果验证失败，将错误信息转换为特定错误类型，并返回给客户端
            let err = ErrNo::default();
            return Ok(send_response(cx, err, "Invalid data".to_string()).await.into_response());
        }
    };

    // 3. 从请求上下文中获取用户身份信息
    let user_identity = cx.extensions().get::<String>().expect("Identity not found");

    // 4. 调用远程过程创建笔记记录
    match rpc_create_note(user_identity, data).await {
        Ok(_) => {
            // 6. 如果操作成功，返回成功状态给客户端
            let err = ErrNo::default();
            Ok(send_response(cx, err, "Note created successfully".to_string()).await.into_response())
        }
        Err(err) => {
            // 5. 如果创建笔记过程中发生错误，将错误信息转换为特定错误类型，并返回给客户端
            let err = ErrNo::default();
            Ok(send_response(cx, err, "Failed to create note".to_string()).await.into_response())
        }
    }
}

// mock struct for NoteData, replace with actual structure
#[derive(Debug, Deserialize)]
struct NoteData {
    // fields for note data
}

// mock for implement stage
#[derive(Debug, Deserialize)]
struct CreateUserRequest {
    // Define the fields based on your requirements
}

// CreateUser handles the user creation request
pub async fn create_user(
    cx: &mut ServerContext,
    req: ServerRequest,
) -> Result<ServerResponse, Infallible> {
    let (parts, body) = req.into_parts();

    // 1. 使用请求上下文中的方法绑定并验证创建用户请求数据
    let create_user_req = match Json::<CreateUserRequest>::from_request(cx, parts, body).await {
        Ok(data) => data,
        Err(err) => {
            // 2. 如果验证失败，转换错误并发送错误响应给客户端
            let err = ErrNo::default();
            return Ok(send_response(cx, err, "Invalid data".to_string()).await.into_response());
        }
    };

    // 3. 调用RPC服务创建用户
    match rpc::create_user(cx, create_user_req).await {
        Ok(_) => {
            // 5. 如果创建成功，发送成功状态响应给客户端
            let err = ErrNo::default();
            Ok(send_response(cx, err, "successfully".to_string()).await.into_response())
        }
        Err(err) => {
            // 4. 如果创建失败，转换错误并发送错误响应给客户端
            let err = ErrNo::default();
            Ok(send_response(cx, err, "failed to create user".to_string()).await.into_response())
        }
    }
}

// DeleteNote函数的主要功能是处理删除笔记的请求。
pub async fn delete_note(
    cx: &mut ServerContext,
    req: ServerRequest,
) -> Result<ServerResponse, Infallible> {
    // 1. 从请求上下文中绑定并验证请求数据
    let (parts, body) = req.into_parts();
    let request_data = match Json::<NoteRequest>::from_request(cx, parts, body).await {
        Ok(data) => data,
        Err(err) => {
            // 2. 如果绑定或验证失败，转换错误信息并通过请求上下文返回给客户端
            let err = ErrNo::default();
            return Ok(send_response(cx, err, "Invalid data".to_string()).await.into_response());
        }
    };

    // 3. 从请求上下文中获取用户身份信息
    let identity = cx.extensions().get::<String>().cloned().unwrap_or_default();

    // 4. 调用RPC接口删除指定笔记
    match rpc::delete_note(&identity, &request_data.0.note_id).await {
        Ok(_) => {
            // 5. 如果删除成功，返回成功状态给客户端
            let err = ErrNo::default();
            Ok(send_response(cx, err, "successfully".to_string()).await.into_response())
        }
        Err(err) => {
            // 6. 如果删除过程中出现错误，转换错误信息并通过请求上下文返回给客户端
            let err = ErrNo::default();
            Ok(send_response(cx, err, "failed tp delete note".to_string()).await.into_response())
        }
    }
}

// Mock struct for NoteRequest
#[derive(Debug, Deserialize)]
struct NoteRequest {
    note_id: String,
}

// QueryNote函数的主要功能是处理查询笔记的请求。
pub async fn query_note(
    cx: &mut ServerContext,
    req: ServerRequest,
) -> Result<ServerResponse, Infallible> {
    // 1. 定义错误变量err和请求变量req
    let mut err = None;
    let mut request_data = Json::default();

    // 2. 调用请求上下文的BindAndValidate方法绑定并验证请求数据
    let (parts, body) = req.into_parts();
    match Json::from_request(cx, parts, body).await {
        Ok(data) => request_data = data,
        Err(e) => {
            // 3. 如果验证失败，调用SendResponse函数返回错误响应
            let err = ErrNo::default();
            return Ok(send_response(cx, err, "Invalid data".to_string()).await.into_response());
        }
    }

    // 4. 从请求上下文中获取用户身份信息
    let user_id: String = cx
        .extensions()
        .get::<String>()
        .cloned()
        .unwrap_or_default();

    // 5. 调用QueryNotes函数查询笔记，传入用户ID、搜索关键字、偏移量和限制量
    match rpc::query_notes(user_id, request_data).await {
        Ok(response) => {
            let err = ErrNo::default();
            Ok(send_response(cx, err, "successfully".to_string()).await.into_response())
        }
        
        Err(e) => {
            // 6. 如果查询出错，调用SendResponse函数返回错误响应
            let err = ErrNo::default();
            Ok(send_response(cx, err, "failed to query note".to_string()).await.into_response())
        }
    }
}


pub async fn update_note(
    cx: &mut ServerContext,
    req: ServerRequest,
) -> Result<ServerResponse, Infallible> {
    // 2. 使用请求上下文的BindAndValidate方法绑定并验证请求数据到请求数据结构    

    let (parts, body) = req.into_parts();
    let data = match Json::<UpdateNoteRequest>::from_request(cx, parts, body).await {
        Ok(data) =>  data,
        Err(_) => {
            let err = ErrNo::default();
            return Ok(send_response(cx, err, "Invalid data".to_string()).await.into_response());
        }
    };

    // 4. 从请求上下文中获取用户身份信息
    let user_identity = cx
        .extensions()
        .get::<String>()
        .unwrap_or(&consts::IDENTITY_KEY.to_string())
        .clone();

    // 5. 调用rpc.UpdateNote函数更新笔记内容
    match rpc::update_note(&user_identity, &data).await {
        Ok(_) => {
            // 7. 如果更新成功，调用SendResponse函数返回成功响应
            let err = ErrNo::default();
            Ok(send_response(cx, err, "successfully".to_string()).await.into_response())
        }
        Err(e) => {
            // 6. 如果更新失败，调用SendResponse函数返回错误响应
            let err = ErrNo::default();
            Ok(send_response(cx, err, "failed to update note".to_string()).await.into_response())
        }
    }
}

pub async fn send_response(
    _cx: &mut ServerContext,
    err_no: errno::ErrNo,
    data: String,
) -> Result<http_response<String>, Infallible> {
    let res = Response{
        code: err_no.err_code,
        message: err_no.err_msg,
        data: data,
    };
    let response_body = json!(res).to_string();
    let resp = http_response::builder()
        .status(StatusCode::OK)
        .header(http::header::CONTENT_TYPE,mime::APPLICATION_JSON.essence_str())
        .body(response_body)
        .unwrap();
    Ok(resp)
}

// mock for implement stage
#[derive(Debug, Deserialize)]
struct UpdateNoteRequest {
    // define fields here
}
