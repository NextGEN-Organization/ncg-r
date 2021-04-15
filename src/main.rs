
fn change(value: &mut i64) {
    *value = *value + 1;
}




fn n_lines_in_bytes(_bytes: &mut [u8; 17]) {
    // for i in 0..((usize:<bytes>)%17) {
    //     bytes[i] = 1;
    // }
}





fn main() {
    let mut x = 4;
    //mut here allows us to change the passed-in s value. 
    let y = &mut x;
    

    change(y);
    println!("{}", x);
}
    