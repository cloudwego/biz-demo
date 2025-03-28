use test_rust::cmd::api::hertz_router::generated_register;
use test_rust::cmd::api::hertz_handler::demoapi;
use test_rust::cmd::api::hertz_handler;
use test_rust::cmd::api::rpc;
use test_rust::cmd::api::mw;
use test_rust::pkg::errno;
use test_rust::pkg::consts::API_SERVICE_NAME;
use test_rust::pkg::consts::EXPORT_ENDPOINT;
use volo_http::server::route::Router;
use volo_http::server::route::get_service;
use volo_http::server::response::IntoResponse;
use volo_http::server::Server;
use volo_http::http::StatusCode;
use volo_http::Address;
use motore::service::service_fn;
use std::net::SocketAddr;

// register函数用于在Hertz服务器上注册路由和处理器。
pub fn register() -> Router {
    // 1. 调用GeneratedRegister函数，为Hertz服务器配置路由和中间件
    let router = generated_register();
    
    // 2. 调用customizedRegister函数，在Hertz服务器上注册自定义路由和处理器
    router.merge(customized_register())
}

// customized_register registers custom routers.
pub fn customized_register() -> Router {
    Router::new()
        // 1. 注册 GET 请求的 '/ping' 路由，并指定处理函数
        .route("/ping", get_service(service_fn(hertz_handler::ping)))
        // 2. 为未匹配到任何路由的请求注册处理函数
        .fallback(handle_uri_not_found)
        // 3. 为使用了不被允许的方法的请求注册处理函数
        .fallback(handle_method_not_allowed)
}

// fallback handler for METHOD NOT ALLOWED
async fn handle_method_not_allowed() -> (StatusCode, &'static str) {
    demoapi::send_response(errno::SERVICE_ERR, "Method Not Allowed")
}

// fallback handler for NOT FOUND
async fn handle_uri_not_found() -> (StatusCode, &'static str) {
    demoapi::send_response(errno::SERVICE_ERR, "404 Not Found")
}

pub fn init() {
    // 1. 初始化系统中的用户服务客户端和笔记服务客户端
    rpc::init();

    // 2. 初始化JWT中间件，配置其参数和回调函数
    mw::init_jwt();

    // 3. 创建并设置新的日志记录器
    // 使用 tracing_subscriber 创建一个新的日志记录器
    // let subscriber = FmtSubscriber::builder()
    //     .with_max_level(Level::INFO)
    //     .finish();

    // // 设置全局日志记录器
    // tracing::subscriber::set_global_default(subscriber)
    //     .expect("Failed to set global default subscriber");

    // // // 4. 设置日志记录级别为信息级别
    // info!("System initialized successfully.");
}

#[volo::main]
async fn main() {
    // 1. 创建OpenTelemetry提供者，配置服务名称、导出端点和安全设置
    // let tracer = opentelemetry_otlp::new_pipeline()
    //     .tracing()
    //     .with_endpoint(EXPORT_ENDPOINT)
    //     .with_trace_config(
    //         sdktrace::config().with_resource(Resource::new(vec![
    //             KeyValue::new("service.name", API_SERVICE_NAME),
    //         ])),
    //     )
    //     .install_simple()
    //     .expect("Failed to install OpenTelemetry tracer");

    // let opentelemetry = OpenTelemetryLayer::new(tracer);
    // let subscriber = Registry::default().with(opentelemetry);
    // tracing::subscriber::set_global_default(subscriber).expect("Failed to set tracing subscriber");

    // 2. 注册一个延迟关闭操作以确保OpenTelemetry提供者在程序结束时正确关闭
    // (Note: This is a pseudo-code to indicate the intention)
    // let _guard = opentelemetry::global::shutdown_tracer_provider();

    // 3. 初始化系统组件
    init();

    // 4. 创建一个新的服务器实例，配置端口、处理方法和追踪器
    let app = Router::new()
        .merge(register()); // 6. 在服务器上注册路由和处理器

    let addr = "[::]:8080".parse::<SocketAddr>().unwrap();
    let addr = Address::from(addr);

    println!("Listening on {addr}");

    // 7. 启动服务器
    Server::new(app).run(addr).await.unwrap();
}
