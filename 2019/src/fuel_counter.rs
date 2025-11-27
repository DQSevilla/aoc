use std::fs;

pub fn day_one() {
    println!("Day 1: The Tyranny of the Rocket Equation");

    let masses = read_input("src/input/day_01_masses.txt");

    print!("Part 1: ");
    let fuel = total_naive(&masses);
    println!("Total fuel requirement (naive) is {fuel}");

    print!("Part 2: ");
    let fuel = total_robust(&masses);
    println!("Total fuel requirement (robust) is {fuel}");
}

fn read_input(file_name: &str) -> Vec<u32> {
    let input = fs::read_to_string(file_name).expect("Failed to read file");
    input
        .split_whitespace()
        .map(|s| s.parse::<u32>().expect("Failed to parse u32"))
        .collect()
}

fn total_naive(masses: &[u32]) -> u32 {
    masses.iter().copied().map(|mass| mass / 3 - 2).sum()
}

// fn calc_fuel_robust(mass: u32) -> u32 {
//     std::iter::successors(Some(mass), |&m| {
//         let fuel = (m / 3).saturating_sub(2);
//         if fuel > 0 {
//             Some(fuel)
//         } else {
//             None
//         }
//     })
//     .skip(1)
//     .sum()
// }

fn total_robust(masses: &[u32]) -> u32 {
    let mut sum: u32 = 0;
    for &mass in masses {
        let mut fuel = mass as i32 / 3 - 2;
        while fuel > 0 {
            sum += fuel as u32;
            fuel = fuel / 3 - 2;
        }
    }
    sum
}
