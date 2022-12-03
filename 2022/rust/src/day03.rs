use std::collections::HashSet;
use std::{
    fs::File,
    io::{BufRead, BufReader},
};

#[allow(dead_code)]
fn part1(filename: String) -> Result<isize, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut left_pouch: HashSet<char> = HashSet::new();

    let mut sum: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        for c in line.chars().take(line.len() / 2) {
            if let None = left_pouch.get(&c) {
                left_pouch.insert(c);
            };
        }

        for c in line.chars().skip(line.len() / 2) {
            if let Some(c) = left_pouch.get(&c) {
                sum = sum + priority(&c);
                _ = left_pouch.drain();
                break;
            };
        }
    }

    return Ok(sum);
}

#[allow(dead_code)]
fn part2(filename: String) -> Result<isize, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut rucksacks: [HashSet<char>; 2] = [HashSet::new(), HashSet::new()];

    let mut sum: isize = 0;
    let mut group: usize = 1;

    for line in r.lines() {
        let line = line.unwrap();

        for c in line.chars() {
            if group % 3 == 0 {
                let in_group_one = match rucksacks[0].get(&c) {
                    Some(_) => true,
                    None => false,
                };

                let in_group_two = match rucksacks[1].get(&c) {
                    Some(_) => true,
                    None => false,
                };

                if in_group_one && in_group_two {
                    sum = sum + priority(&c);
                    _ = rucksacks[0].drain();
                    _ = rucksacks[1].drain();
                }

                continue;
            }

            if let None = rucksacks[(group % 3) - 1].get(&c) {
                rucksacks[(group % 3) - 1].insert(c);
            };
        }

        group = group + 1;
        if group > 3 {
            group = 1
        }
    }

    return Ok(sum);
}

fn priority(c: &char) -> isize {
    if *c >= 'a' && *c <= 'z' {
        return *c as isize - 96;
    }
    if *c >= 'A' && *c <= 'Z' {
        return *c as isize - 38;
    }
    return 0;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(String::from("../inputs/day03.sample.txt")).unwrap();
        assert_eq!(result, 157);

        let result = part1(String::from("../inputs/day03.input.txt")).unwrap();
        assert_eq!(result, 8202);
    }

    #[test]
    fn test_part2() {
        let result = part2(String::from("../inputs/day03.sample.txt")).unwrap();
        assert_eq!(result, 70);

        let result = part2(String::from("../inputs/day03.input.txt")).unwrap();
        assert_eq!(result, 2864);
    }
}
