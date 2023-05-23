use std::env;
use whatlang::{detect};

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() == 2 {
      let input_text = &args[1];
      let info = detect(input_text).unwrap();
      println!("{}", info.lang());
    }
    else {
      panic!("Pass the text to be used for language detection as first argument...");
    }
}
