mod interpretmath;

use interpretmath::tokenizer::Tokenizer;

fn main() {
    let expression = "3 + 5 * (10 - 4)";
    let mut tokenizer = Tokenizer::new(expression);
    println!("Hello, world!");
}
