use std::{
    fs,
    thread::sleep,
    time::{Duration, Instant},
};

use dotenv::dotenv;
use std::env;

fn read_dir() {
    let entries = fs::read_dir("/tmp").unwrap();
    for entry in entries {
        if let Ok(entry) = entry {
            println!("{:?}", entry.path());
        }
    }
}

fn elapsedt() {
    let now = Instant::now();
    sleep(Duration::new(3, 0));
    println!("{:?}", now.elapsed().as_secs());
}

fn envs() {
    dotenv().ok();

    for (key, value) in env::vars() {
        println!("{} : {}", key, value);
    }
}

fn main() {
    //read_dir();
    //elapsedt();
    envs();
}
