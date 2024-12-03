use std::{error::Error, str::FromStr};

use regex::Regex;

struct Solution {
    numbers: Vec<(Action, (isize, isize))>,
}

#[derive(Debug)]
enum Action {
    Do,
    DoNot,
}

impl FromStr for Solution {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut numbers = Vec::new();

        // The number of capture groups must match between the OR conditions or else the extract
        // method below will panic.
        let re = Regex::new(r"(do)(\(\))|(don't)(\(\))|mul\((\d{1,3}),(\d{1,3})\)")?;

        // Action defaults to enabled.
        let mut action = Action::Do;

        for line in s.lines() {
            for capture in re.captures_iter(line) {
                match capture.extract() {
                    ("do()", _) => action = Action::Do,
                    ("don't()", _) => action = Action::DoNot,
                    // TODO: Can we clone this enum instead of using the nested match?
                    (_, [lhs, rhs]) => match action {
                        Action::Do => numbers
                            .push((Action::Do, (lhs.parse::<isize>()?, rhs.parse::<isize>()?))),
                        Action::DoNot => numbers.push((
                            Action::DoNot,
                            (lhs.parse::<isize>()?, rhs.parse::<isize>()?),
                        )),
                    },
                }
            }
        }

        Ok(Solution { numbers })
    }
}

impl Solution {
    #[allow(dead_code)]
    fn part01(self: &Self) -> String {
        let res: isize = self.numbers.iter().map(|(_, (lhs, rhs))| lhs * rhs).sum();

        format!("{}", res)
    }

    #[allow(dead_code)]
    fn part02(self: &Self) -> String {
        let res: isize = self
            .numbers
            .iter()
            .map(|(action, (lhs, rhs))| match action {
                Action::Do => lhs * rhs,
                Action::DoNot => 0,
            })
            .sum();

        format!("{}", res)
    }
}

#[cfg(test)]
mod tests {
    use std::{fs::read_to_string, path::Path};

    use super::*;

    #[test]
    fn test_part01() {
        let input = read_to_string(Path::new("../inputs/day03-part01-example.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "161");

        let input = read_to_string(Path::new("../inputs/day03.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "174336360");
    }

    #[test]
    fn test_part02() {
        let input = read_to_string(Path::new("../inputs/day03-part02-example.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "48");

        let input = read_to_string(Path::new("../inputs/day03.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "88802350");
    }
}
