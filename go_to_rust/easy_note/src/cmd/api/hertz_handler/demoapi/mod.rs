use std::convert::Infallible;
use volo_http::context::ServerContext;
use volo_http::request::ServerRequest;
use volo_http::response::ServerResponse;
use volo_http::http;
use volo_http::server::IntoResponse;

// mock for implement stage
pub async fn check_user(
    _cx: &mut ServerContext, // 上下文 context
    _req: ServerRequest, //  用户请求
) -> Result<ServerResponse, Infallible> {
    // todo!("implement user login logic");
    Ok("todo".into_response())
}

// mock for implement stage
pub async fn create_note(
    _cx: &mut ServerContext, // 上下文 context
    _req: ServerRequest, //  用户请求
) -> Result<ServerResponse, Infallible> {
    // todo!("implement user login logic");
    Ok("todo".into_response())
}

// mock for implement stage
pub async fn create_user(
    _cx: &mut ServerContext, // 上下文 context
    _req: ServerRequest, //  用户请求
) -> Result<ServerResponse, Infallible> {
    // todo!("implement user login logic");
    Ok("todo".into_response())
}
// mock for implement stage
pub async fn delete_note(
    _cx: &mut ServerContext, // 上下文 context
    _req: ServerRequest, //  用户请求
) -> Result<ServerResponse, Infallible> {
    // todo!("implement user login logic");
    Ok("todo".into_response())
}
// mock for implement stage
pub async fn query_note(
    _cx: &mut ServerContext, // 上下文 context
    _req: ServerRequest, //  用户请求
) -> Result<ServerResponse, Infallible> {
    // todo!("implement user login logic");
    Ok("todo".into_response())
}

// mock for implement stage
pub fn send_response(_err: &'static str, _msg: &'static str) -> (http::StatusCode, &'static str) {
    (http::StatusCode::INTERNAL_SERVER_ERROR, _msg)
}

// mock for implement stage
pub async fn update_note(
    _cx: &mut ServerContext, // 上下文 context
    _req: ServerRequest, //  用户请求
) -> Result<ServerResponse, Infallible> {
    // todo!("implement user login logic");
    Ok("todo".into_response())
}