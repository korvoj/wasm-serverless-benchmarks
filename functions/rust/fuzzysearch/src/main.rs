use simsearch::SimSearch;
use std::fs::File;
use std::env;
use std::path::Path;
use std::io::{self, prelude::*, BufReader};

fn read_file(filename: impl AsRef<Path>) -> Vec<String> {
    let file = File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect()
}

fn main() -> io::Result<()> {

    let args: Vec<String> = env::args().collect();
    if args.len() == 3 {
      let mut engine: SimSearch<u32> = SimSearch::new();
      let lines = read_file(&args[1]);
      let mut counter = 0;
      for line in lines {
//          println!("{}", line);
          engine.insert(counter, &line);
          counter+=1;
      }

      let mut results: Vec<u32> = engine.search(&args[2]);
      results.sort();
      println!("{:?}", results);
      Ok(())
    }
    else {
      panic!("Pass the path to the input file to be used for searching...")
    }
}
