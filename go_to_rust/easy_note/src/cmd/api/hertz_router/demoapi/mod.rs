use volo::service::service_fn;
use volo_http::server::middleware::Next;
use volo_http::server::middleware::from_fn;
use volo_http::server::route::post_service;
use volo_http::server::route::Router;
use volo_http::context::ServerContext;
use volo_http::request::ServerRequest;
use volo_http::response::ServerResponse;
use volo_http::server::IntoResponse;
use crate::cmd::api::hertz_handler::demoapi::create_note;
use crate::cmd::api::hertz_handler::demoapi::update_note;
use crate::cmd::api::hertz_handler::demoapi::delete_note;
use crate::cmd::api::hertz_handler::demoapi::query_note;
use crate::cmd::api::hertz_handler::demoapi::check_user;
use crate::cmd::api::hertz_handler::demoapi::create_user;

// _querynote_mw 函数用于处理查询笔记的中间件逻辑
pub async fn _querynote_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // 目前没有具体的中间件逻辑实现，直接调用下一个中间件或处理函数
    next.run(cx, req).await.into_response()
}

// _user_mw 函数的主要功能是返回一个包含 app.HandlerFunc 类型的切片。
// 目前，该函数的实现是返回 nil，意味着它没有实际的处理逻辑。
pub async fn _user_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // 暂时没有实际的处理逻辑，直接调用 next
    next.run(cx, req).await.into_response()
}

// Register function to configure routes and middleware for the Hertz server.
pub fn register() -> Router {
    Router::new()
    // 1. Configure middleware for the root route
    // .layer(from_fn(root_mw))
    // 2. Configure version route "/v1" and add middleware
    // .route("/v1", post(_v1_mw))
    // 3. Configure note-related routes under "/v1/note"
    .route("/v1/note", post_service(service_fn(create_note)).layer(from_fn(_note_mw)))
    .route("/v1/note/update", post_service(service_fn(update_note)).layer(from_fn(_updatenote_mw)))
    .route("/v1/note/delete", post_service(service_fn(delete_note)).layer(from_fn(_deletenote_mw)))
    .route("/v1/note/query", post_service(service_fn(query_note)).layer(from_fn(_querynote_mw)))
    // 4. Configure user-related routes under "/v1/user"
    .route("/v1/user/login", post_service(service_fn(check_user)).layer(from_fn(_checkuser_mw)))
    .route("/v1/user/register", post_service(service_fn(create_user)).layer(from_fn(_createuser_mw)))
}

// Recovery middleware to handle panics
async fn recovery_middleware(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // Here we should use a real recovery middleware from a library
    // For example, if using `tower-http`, it might look like this:
    // tower_http::middleware::catch_panic::CatchPanicLayer::new()
    // .on_panic(|err| {
    //     error!("Internal server error: {:?}", err);
    //     (http::StatusCode::INTERNAL_SERVER_ERROR, "Internal Server Error").into_response()
    // })
    // .layer(next)
    
    // Mock implementation for now
    let response = next.run(cx, req).await;
    response.into_response()
}

// Request ID middleware
async fn request_id_middleware(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // Here we should use a real request ID middleware
    // For example, if using `tower-http`, it might look like this:
    // tower_http::middleware::request_id::RequestIdLayer::new()
    // .layer(next)

    // Mock implementation for now
    let response = next.run(cx, req).await;
    response.into_response()
}

// Gzip middleware
async fn gzip_middleware(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // Here we should use a real gzip middleware
    // For example, if using `tower-http`, it might look like this:
    // tower_http::middleware::compression::CompressionLayer::new()
    // .layer(next)

    // Mock implementation for now
    let response = next.run(cx, req).await;
    response.into_response()
}

// rootMw function to initialize and return a set of middleware handlers
pub async fn root_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // Apply recovery middleware
    // let response = recovery_middleware(cx, req, next).await.;
    // // Apply request ID middleware
    // let response = request_id_middleware(cx, req, next).await;
    // // Apply gzip middleware
    // let response = gzip_middleware(cx, req, next).await;
    
    next.run(cx, req).await.into_response()
}

// _note_mw 函数用于返回一个包含JWT中间件的处理器函数数组。
pub async fn _note_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // 使用 JWT_MIDDLEWARE 进行 JWT 认证和授权
    next.run(cx, req).await.into_response()
}

// _createuser_mw function creates and returns a slice of app.HandlerFunc type.
// Currently, it returns nil.
pub async fn _createuser_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // Step 1: Return a slice of app.HandlerFunc type, currently returning nil.
    next.run(cx, req).await.into_response()
}

// _deletenote_mw 函数返回一个空的 app.HandlerFunc 类型的切片
pub async fn _deletenote_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // 由于没有具体的业务逻辑，直接调用 next 继续处理请求
    next.run(cx, req).await.into_response()
}

// _updatenote_mw 函数用于返回一个 app.HandlerFunc 类型的切片
pub async fn _updatenote_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // 1. 创建一个 app.HandlerFunc 类型的切片，用于存储中间件函数
    // let mut middlewares: Vec<HandlerFunc> = Vec::new();

    // 2. 将需要的中间件函数添加到切片中
    // 这里可以添加具体的中间件函数，例如身份验证、日志记录等
    // middlewares.push(auth_middleware);
    // middlewares.push(logging_middleware);

    // 3. 返回包含所有中间件函数的切片
    // 由于没有具体的中间件函数可用，暂时直接调用 next
    // 调用下一个中间件或处理函数
    next.run(cx, req).await.into_response()
}

// Define a type alias for the middleware function signature
type HandlerFunc = fn(&mut ServerContext, ServerRequest, Next) -> ServerResponse;

// _v1_mw function returns a vector of middleware functions
pub async fn _v1_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    next.run(cx, req).await.into_response()
}

// _checkuser_mw 函数的主要功能是返回一个包含 app.HandlerFunc 类型的切片。
// 目前函数体内的具体实现代码缺失（即返回 nil），但其设计目的是为了提供一组处理用户检查的中间件函数。
pub async fn _checkuser_mw(cx: &mut ServerContext, req: ServerRequest, next: Next) -> ServerResponse {
    // 1. 返回一个包含处理用户检查逻辑的中间件函数的切片。
    // 这里可以加入用户检查逻辑，例如验证用户身份、检查权限等。
    
    // 调用下一个中间件或处理函数
    next.run(cx, req).await.into_response()
}
