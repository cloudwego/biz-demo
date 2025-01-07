pub mod demoapi;

use std::convert::Infallible;
use volo_http::context::ServerContext;
use volo_http::http::StatusCode;
use volo_http::request::ServerRequest;
use volo_http::response::ServerResponse;
use volo_http::server::response::IntoResponse;
use volo_http::json::Json;
use serde_json;

// Ping函数用于处理传入的HTTP请求，并返回一个JSON响应。
pub async fn ping(
    _ctx: &mut ServerContext, // 上下文信息
    _req: ServerRequest,      // 请求对象
) -> Result<ServerResponse, Infallible> {
    // 1. 调用 RequestContext 的 JSON 方法，传入状态码 200 和包含键值对 {"message": "pong"} 的对象，来生成并发送 JSON 响应。
    let response = Json(serde_json::json!({ "message": "pong" })).into_response();
    Ok(response)
}
