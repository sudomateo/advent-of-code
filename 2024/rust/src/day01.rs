use std::{collections::HashMap, error::Error, str::FromStr};

struct Solution {
    list1: Vec<isize>,
    list2: Vec<isize>,
}

impl FromStr for Solution {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut list1 = Vec::new();
        let mut list2 = Vec::new();

        // Every line is expected to be in the following format:
        // 123   456
        for line in s.lines() {
            let fields: Vec<&str> = line.split_whitespace().collect();

            if fields.len() != 2 {
                return Err(format!("failed to parse 2 numbers from line: {}", line).into());
            }

            list1.push(fields.first().unwrap().parse::<isize>().unwrap());
            list2.push(fields.last().unwrap().parse::<isize>().unwrap());
        }

        Ok(Solution { list1, list2 })
    }
}

impl Solution {
    #[allow(dead_code)]
    fn part01(self: &mut Self) -> String {
        // The deep copy separates the state for part01 and part02.
        let mut list1: Vec<isize> = self.list1.iter().cloned().collect();
        list1.sort();
        let mut list2: Vec<isize> = self.list2.iter().cloned().collect();
        list2.sort();

        let sum = list1
            .iter()
            .zip(list2.iter())
            .fold(0, |acc, (l1, l2)| acc + (l1 - l2).abs());

        format!("{}", sum)
    }

    #[allow(dead_code)]
    fn part02(self: &mut Self) -> String {
        let mut freq_table = HashMap::new();

        for num in &self.list2 {
            *freq_table.entry(num).or_insert(0) += 1;
        }

        let sum = self.list1.iter().fold(0, |acc, num| {
            let multiplier = freq_table.entry(num).or_default();
            acc + (*num * *multiplier)
        });

        format!("{}", sum)
    }
}

#[cfg(test)]
mod tests {
    use std::{fs::read_to_string, path::Path};

    use super::*;

    #[test]
    fn test_part01() {
        let input = read_to_string(Path::new("../inputs/day01-example.txt")).unwrap();
        let mut s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "11");

        let input = read_to_string(Path::new("../inputs/day01.txt")).unwrap();
        let mut s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "2904518");
    }

    #[test]
    fn test_part02() {
        let input = read_to_string(Path::new("../inputs/day01-example.txt")).unwrap();
        let mut s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "31");

        let input = read_to_string(Path::new("../inputs/day01.txt")).unwrap();
        let mut s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "18650129");
    }
}
