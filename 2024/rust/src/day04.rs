use std::{collections::HashSet, error::Error, str::FromStr};

use itertools::iproduct;

struct Solution {
    grid: Vec<Vec<char>>,
}

impl FromStr for Solution {
    type Err = Box<dyn Error>;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut grid = Vec::new();

        for line in s.lines() {
            let chars = line.chars().collect();
            grid.push(chars);
        }

        Ok(Solution { grid })
    }
}

impl Solution {
    #[allow(dead_code)]
    fn part01(self: &Self) -> String {
        let h = self.grid.len();
        let w = self.grid[0].len();

        let res: usize = iproduct!(0..h, 0..w)
            // Always start at an `X`.
            .filter(|&(row, col)| self.grid[row][col] == 'X')
            .map(|(row, col)| {
                // Visit all 8 neighbors.
                [
                    (1, 0),
                    (1, 1),
                    (0, 1),
                    (-1, 1),
                    (-1, 0),
                    (-1, -1),
                    (0, -1),
                    (1, -1),
                ]
                .into_iter()
                .filter(|(row_delta, col_delta)| {
                    // Continue walking in the direction looking for `XMAS` in order.
                    // Finding the char_indices method was a blessing.
                    "XMAS".char_indices().all(|(idx, ch)| {
                        // The checked_add_signed is really used to catch the overflow when going
                        // negative. For example, when row or col and 0 and we add -1 to them.
                        match (
                            row.checked_add_signed(row_delta * idx as isize),
                            col.checked_add_signed(col_delta * idx as isize),
                        ) {
                            // If we got back valid values from checked_add_signed then we just
                            // have to confirm that they are within grid bounds and that the
                            // character we're at matches the character from `XMAS`.
                            (Some(new_row), Some(new_col)) => {
                                new_row < h && new_col < w && self.grid[new_row][new_col] == ch
                            }

                            // Any other case is an invalid match so we can move to the next
                            // coordinate on the grid.
                            (_, _) => false,
                        }
                    })
                })
                .count()
            })
            .sum();

        format!("{}", res)
    }

    #[allow(dead_code)]
    fn part02(self: &Self) -> String {
        let h = self.grid.len();
        let w = self.grid[0].len();

        // There can't possibly be an X anywhere on the border since it wouldn't leave enough room.
        let res: usize = iproduct!(1..h - 1, 1..w - 1)
            // Always start at an `A` since that's the middle of the X.
            // M . M
            // . A .
            // S . S
            .filter(|&(row, col)| self.grid[row][col] == 'A')
            .filter(|&(row, col)| {
                // Diag 1: top left, bottom right
                // Diag 2: top right, bottom left
                [[(1, -1), (-1, 1)], [(1, 1), (-1, -1)]]
                    .into_iter()
                    .all(|diag_deltas| {
                        let mut set = HashSet::from(['M', 'S']);

                        diag_deltas.iter().for_each(|(row_delta, col_delta)| {
                            let new_row = row.checked_add_signed(*row_delta).unwrap();
                            let new_col = col.checked_add_signed(*col_delta).unwrap();
                            set.remove(&self.grid[new_row][new_col]);
                        });

                        set.is_empty()
                    })
            })
            .count();

        format!("{}", res)
    }
}

#[cfg(test)]
mod tests {
    use std::{fs::read_to_string, path::Path};

    use super::*;

    #[test]
    fn test_part01() {
        let input = read_to_string(Path::new("../inputs/day04-example.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "18");

        let input = read_to_string(Path::new("../inputs/day04.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part01(), "2662");
    }

    #[test]
    fn test_part02() {
        let input = read_to_string(Path::new("../inputs/day04-example.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "9");

        let input = read_to_string(Path::new("../inputs/day04.txt")).unwrap();
        let s = Solution::from_str(&input).unwrap();
        assert_eq!(s.part02(), "2034");
    }
}
