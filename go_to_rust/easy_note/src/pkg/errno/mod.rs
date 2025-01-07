use std::sync::Arc;

use serde::{Deserialize, Serialize};

// mock for implement stage
#[derive(Default, Serialize, Deserialize)]
pub struct ErrNo {
    pub err_code: i64,
    pub err_msg: String,
}


// mock for implement stage
pub fn convert_err(_err: impl Into<ErrNo>) -> ErrNo {
    // todo!("implement it");
    ErrNo::default()
}

// mock for implement stage
pub static SUCCESS: Arc<str> = Arc::from("Success");

// mock for implement stage
pub static SERVICE_ERR: &'static str = "Service error";
