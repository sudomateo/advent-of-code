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
#[allow(dead_code)]
enum Choice {
    Rock,
    Paper,
    Scissors,
}

#[allow(dead_code)]
fn part1(filename: String) -> Result<isize, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut score: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        let fields: Vec<&str> = line.split(" ").collect();

        let their_choice = get_choice(fields[0]).unwrap();
        let my_choice = get_choice(fields[1]).unwrap();

        score = score + score_one(&my_choice, &their_choice);
    }

    return Ok(score);
}

#[allow(dead_code)]
pub fn part2(filename: String) -> Result<isize, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut score: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        let fields: Vec<&str> = line.split(" ").collect();

        let their_choice = get_choice(fields[0]).unwrap();

        score = score + score_two(&their_choice, fields[1]);
    }

    return Ok(score);
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
    fn test_part1() {
        let result = part1(String::from("../inputs/day02.sample.txt")).unwrap();
        assert_eq!(result, 15);

        let result = part1(String::from("../inputs/day02.input.txt")).unwrap();
        assert_eq!(result, 12458);
    }

    #[test]
    fn test_part2() {
        let result = part2(String::from("../inputs/day02.sample.txt")).unwrap();
        assert_eq!(result, 12);

        let result = part2(String::from("../inputs/day02.input.txt")).unwrap();
        assert_eq!(result, 12683);
    }
}
