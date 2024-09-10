use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    let file = File::open("../../input.txt")?;
    let reader = BufReader::new(file);

    let mut lines = 0;

    for _line in reader.lines() {
        lines += 1;
    }

    println!("{}", lines);

    Ok(())
}
