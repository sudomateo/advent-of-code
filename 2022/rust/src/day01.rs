use std::{
    fs::File,
    io::{BufRead, BufReader},
};

#[allow(dead_code)]
fn part1(filename: String) -> Result<isize, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut max_calories: isize = 0;
    let mut sum: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        if line == "" {
            if sum > max_calories {
                max_calories = sum;
            }
            sum = 0;
            continue;
        }

        sum = sum + line.parse::<isize>().unwrap();
    }

    return Ok(max_calories);
}

#[allow(dead_code)]
fn part2(filename: String) -> Result<isize, &'static str> {
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

    if calories.len() < 3 {
        return Err("not enough elves to count");
    }

    return Ok(calories[0] + calories[1] + calories[2]);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(String::from("../inputs/day01.sample.txt")).unwrap();
        assert_eq!(result, 24000);

        let result = part1(String::from("../inputs/day01.input.txt")).unwrap();
        assert_eq!(result, 72240);
    }

    #[test]
    fn test_part2() {
        let part2_result = part2(String::from("../inputs/day01.sample.txt")).unwrap();
        assert_eq!(part2_result, 45000);

        let part2_result = part2(String::from("../inputs/day01.input.txt")).unwrap();
        assert_eq!(part2_result, 210957);
    }
}
