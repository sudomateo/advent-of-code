use std::collections::HashSet;
use std::fs;

#[allow(dead_code)]
fn part1(filename: &str) -> Result<usize, &'static str> {
    let datastream = fs::read_to_string(filename).unwrap();
    Ok(decode(datastream, 4))
}

#[allow(dead_code)]
fn part2(filename: &str) -> Result<usize, &'static str> {
    let datastream = fs::read_to_string(filename).unwrap();
    Ok(decode(datastream, 14))
}

fn decode(datastream: String, window: usize) -> usize {
    let mut seen = HashSet::new();

    for (i, nums) in datastream.into_bytes().windows(window).enumerate() {
        for num in nums.iter() {
            seen.insert(num);
        }

        if seen.len() == window {
            return i + window;
        }

        seen.clear();
    }

    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(&String::from("../inputs/day06.sample.txt")).unwrap();
        assert_eq!(result, 7);

        let result = part1(&String::from("../inputs/day06.input.txt")).unwrap();
        assert_eq!(result, 1757);
    }

    #[test]
    fn test_part2() {
        let result = part2(&String::from("../inputs/day06.sample.txt")).unwrap();
        assert_eq!(result, 19);

        let result = part2(&String::from("../inputs/day06.input.txt")).unwrap();
        assert_eq!(result, 2950);
    }
}
