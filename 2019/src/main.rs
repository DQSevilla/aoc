use std::io::{self, Write};

mod fuel_counter;

fn main() {
    println!("Advent of Code - 2019");
    println!("=====================");

    loop {
        println!();
        print!("Enter a day to run (q to quit): ");
        io::stdout().flush().expect("Failed to flush stdout");

        let mut input = String::new();
        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read line");

        if input.trim() == "q" {
            return;
        }

        let day: u32 = match input.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Invalid day entered: must be in range [1, 25]");
                continue;
            }
        };

        match day {
            1 => fuel_counter::day_one(),
            2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19
            | 20 | 21 | 22 | 23 | 24 | 25 => {
                println!("Not yet implemented!");
                continue;
            }
            _ => {
                println!("Invalid day entered: must be in range [1, 25]");
                continue;
            }
        }
    }
}
