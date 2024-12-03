use std::{cmp::Ordering, error::Error, str::FromStr};

struct Solution {
    reports: Vec<Vec<isize>>,
}

impl FromStr for Solution {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut reports = Vec::new();

        // Every line is expected to be in the following format:
        // 1 2 3 4 5
        for line in s.lines() {
            let levels: Vec<&str> = line.split_whitespace().collect();
            let levels: Vec<isize> = levels
                .iter()
                .map(|elem| elem.parse::<isize>().unwrap())
                .collect();

            reports.push(levels);
        }

        Ok(Solution { reports })
    }
}

impl Solution {
    #[allow(dead_code)]
    fn part01(self: &Self) -> String {
        let safe_reports = self
            .reports
            .iter()
            .filter(|report| is_safe_report(report))
            .count();

        format!("{}", safe_reports)
    }

    #[allow(dead_code)]
    fn part02(self: &Self) -> String {
        let safe_reports = self
            .reports
            .iter()
            .filter(|report| {
                is_safe_report(report)
                    || (0..report.len()).any(|i| {
                        // Generate a report without the specified index to see if it's now safe.
                        let filtered_report = report
                            .iter()
                            .enumerate()
                            .filter_map(|(idx, val)| (i != idx).then_some(*val))
                            .collect::<Vec<isize>>();
                        is_safe_report(&filtered_report)
                    })
            })
            .count();

        format!("{}", safe_reports)
    }
}

// is_safe_report determines whether a given report is safe according to the following rules.
// - The levels in the report must be strictly increasing or strictly decreasing.
// - The difference between adjacent levels must have an absolute value of 1..=3.
fn is_safe_report(report: &[isize]) -> bool {
    // To keep track of the first Ordering we've seen during iteration so subsequent iterations can
    // be compared against it.
    let mut ord: Option<Ordering> = None;

    report.windows(2).all(|levels| {
        // Give me them sweet, sweet tuples.
        let (a, b) = (&levels[0], &levels[1]);

        // The .or ensures we update the Ordering exactly once from None to Some.
        ord = ord.or(Some(a.cmp(b)));

        match a - b {
            -3..=-1 => ord == Some(Ordering::Less),
            1..=3 => ord == Some(Ordering::Greater),
            _ => false,
        }
    })
}

#[cfg(test)]
mod tests {
    use std::{fs::read_to_string, path::Path};

    use super::*;

    #[test]
    fn test_part01() {
        let input = read_to_string(Path::new("../inputs/day02-example.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "2");

        let input = read_to_string(Path::new("../inputs/day02.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "334");
    }

    #[test]
    fn test_part02() {
        let input = read_to_string(Path::new("../inputs/day02-example.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "4");

        let input = read_to_string(Path::new("../inputs/day02.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "400");
    }
}
