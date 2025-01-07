use volo_http::http;

// mock for implement stage
pub fn send_response(_err: &'static str, _msg: &'static str) -> (http::StatusCode, &'static str) {
    (http::StatusCode::INTERNAL_SERVER_ERROR, _msg)
}
