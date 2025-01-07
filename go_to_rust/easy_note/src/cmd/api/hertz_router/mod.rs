pub mod demoapi;
use volo_http::server::route::Router;



// GeneratedRegister函数的主要功能是为Hertz服务器配置路由和中间件。
pub fn generated_register() -> Router {
    // 1. 调用 Register 函数为 Hertz 服务器配置路由和中间件
    demoapi::register()
}
