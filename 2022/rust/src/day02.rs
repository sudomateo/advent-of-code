use std::collections::HashMap;
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

#[derive(Debug, PartialEq, Eq, Hash)]
enum Choice {
    Rock,
    Paper,
    Scissors,
}

pub fn solve(filename: String) -> Result<Solution, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut score1: isize = 0;
    let mut score2: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        let fields: Vec<&str> = line.split(" ").collect();

        let their_choice = get_choice(fields[0]).unwrap();
        let my_choice = get_choice(fields[1]).unwrap();

        score1 = score1 + score_one(&my_choice, &their_choice);
        score2 = score2 + score_two(&their_choice, fields[1]);
    }

    return Ok(Solution {
        part1: score1,
        part2: score2,
    });
}

fn get_choice(s: &str) -> Option<Choice> {
    match s {
        "A" | "X" => Some(Choice::Rock),
        "B" | "Y" => Some(Choice::Paper),
        "C" | "Z" => Some(Choice::Scissors),
        _ => None,
    }
}

fn score_one(my_choice: &Choice, their_choice: &Choice) -> isize {
    let choice_map: HashMap<Choice, (Choice, Choice)> = HashMap::from([
        (Choice::Rock, (Choice::Scissors, Choice::Paper)),
        (Choice::Paper, (Choice::Rock, Choice::Scissors)),
        (Choice::Scissors, (Choice::Paper, Choice::Rock)),
    ]);

    let score_map: HashMap<Choice, isize> =
        HashMap::from([(Choice::Rock, 1), (Choice::Paper, 2), (Choice::Scissors, 3)]);

    if my_choice == their_choice {
        return 3 + score_map.get(&my_choice).unwrap();
    }

    if choice_map.get(&my_choice).unwrap().0 == *their_choice {
        return 6 + score_map.get(&my_choice).unwrap();
    }

    return *score_map.get(&my_choice).unwrap();
}

fn score_two(their_choice: &Choice, s: &str) -> isize {
    let choice_map: HashMap<Choice, (Choice, Choice)> = HashMap::from([
        (Choice::Rock, (Choice::Scissors, Choice::Paper)),
        (Choice::Paper, (Choice::Rock, Choice::Scissors)),
        (Choice::Scissors, (Choice::Paper, Choice::Rock)),
    ]);

    let score_map: HashMap<Choice, isize> =
        HashMap::from([(Choice::Rock, 1), (Choice::Paper, 2), (Choice::Scissors, 3)]);

    match s {
        "X" => *score_map
            .get(&choice_map.get(&their_choice).unwrap().0)
            .unwrap(),
        "Y" => 3 + score_map.get(&their_choice).unwrap(),
        "Z" => {
            6 + *score_map
                .get(&choice_map.get(&their_choice).unwrap().1)
                .unwrap()
        }
        _ => 0,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn sample_input() {
        let solution = match solve(String::from("../inputs/day02.sample.txt")) {
            Ok(s) => s,
            Err(e) => panic!("{}", e),
        };

        let expected_solution = Solution {
            part1: 15,
            part2: 12,
        };

        assert_eq!(solution.part1, expected_solution.part1);
        assert_eq!(solution.part2, expected_solution.part2);
    }

    #[test]
    fn real_input() {
        let solution = match solve(String::from("../inputs/day02.input.txt")) {
            Ok(s) => s,
            Err(e) => panic!("{}", e),
        };

        let expected_solution = Solution {
            part1: 12458,
            part2: 12683,
        };

        assert_eq!(solution.part1, expected_solution.part1);
        assert_eq!(solution.part2, expected_solution.part2);
    }
}
