use std::fs;

#[no_mangle]
pub extern "C" fn generate_fibonacci(n: u32) {
    let fibonacci = calculate_fibonacci(n);
    write_fibonacci_to_file(&fibonacci);
}

fn calculate_fibonacci(n: u32) -> Vec<u64> {
    let mut sequence = vec![0, 1];

    for i in 2..=n as usize {
        let next = sequence[i - 1] + sequence[i - 2];
        sequence.push(next);
    }

    sequence
}

fn write_fibonacci_to_file(fibonacci: &[u64]) {
    let file_path = "output.txt";
    let content = fibonacci
        .iter()
        .map(|num| num.to_string())
        .collect::<Vec<String>>()
        .join(", ");

    fs::write(file_path, content).expect("Failed to write file");
}


//wasm-rust-setup/wasm/wasm_rust_setup_bg.wasm