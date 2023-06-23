use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn fibonacci(n: i32) -> i32 {
    if n < 2 {
        return n;
    }
    fibonacci(n - 1) + fibonacci(n - 2)
}