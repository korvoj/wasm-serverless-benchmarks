use std::env;

pub fn prime_numbers(max: usize) -> Vec<usize> {
    let mut result: Vec<usize> = Vec::new();

    if max >= 2 {
        result.push(2)
    }
    for i in (3..max + 1).step_by(2) {
        let stop: usize = (i as f64).sqrt() as usize + 1;
        let mut status: bool = true;

        for j in (3..stop).step_by(2) {
            if i % j == 0 {
                status = false;
                break;
            }
        }
        if status {
            result.push(i)
        }
    }

    result
}

pub fn main(){
  let args: Vec<String> = env::args().collect();
  if args.len() == 2 {
    let prime_limit = *(&args[1].parse::<usize>().unwrap());
    let result = prime_numbers(prime_limit);
    println!("{:#?}", result);  
  }
  else {
    panic!("Pass a single number as a CLI parameter to search for primes among the first n numbers...")
  }
}

