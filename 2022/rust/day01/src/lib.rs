use std::{
    fs::File,
    io::{BufRead, BufReader},
};

#[derive(Debug)]
#[allow(dead_code)]
pub struct Solution {
    part1: isize,
    part2: isize,
}

pub fn solve(filename: String, k: isize) -> Result<Solution, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut calories: Vec<isize> = Vec::new();

    let mut sum: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        if line == "" {
            calories.push(sum);
            sum = 0;
            continue;
        }

        sum = sum + line.parse::<isize>().unwrap();
    }
    calories.push(sum);

    calories.sort_by(|a, b| b.cmp(a));

    if calories.len() < k as usize {
        return Err("not enough elves to count");
    }

    return Ok(Solution {
        part1: calories[0],
        part2: calories[0] + calories[1] + calories[2],
    });
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_input() {
        let solution = match solve(String::from("test/example.txt"), 3) {
            Ok(s) => s,
            Err(e) => panic!("{}", e),
        };

        let expected_solution = Solution {
            part1: 24000,
            part2: 45000,
        };

        assert_eq!(solution.part1, expected_solution.part1);
        assert_eq!(solution.part2, expected_solution.part2);
    }

    #[test]
    fn real_input() {
        let solution = match solve(String::from("test/input.txt"), 3) {
            Ok(s) => s,
            Err(e) => panic!("{}", e),
        };

        let expected_solution = Solution {
            part1: 72240,
            part2: 210957,
        };

        assert_eq!(solution.part1, expected_solution.part1);
        assert_eq!(solution.part2, expected_solution.part2);
    }
}
