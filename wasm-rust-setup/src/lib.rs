extern "C" {
    fn load_input(pointer: *mut u8);
    fn dump_output(pointer: *const u8, length: usize);
}

#[no_mangle]
pub extern "C" fn handler(input_length: usize) {
    // load input data
    let mut input = Vec::with_capacity(input_length);
    unsafe {
        load_input(input.as_mut_ptr());
        input.set_len(input_length);
    }

    // process app data
    let output = input.to_ascii_uppercase();

    // dump output data
    unsafe {
        dump_output(output.as_ptr(), output.len());
    }
}