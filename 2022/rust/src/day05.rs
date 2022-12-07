use std::{
    fs::File,
    io::{BufRead, BufReader},
};

#[allow(dead_code)]
fn part1(filename: &str) -> Result<String, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut count = 0;

    let mut header: Vec<String> = Vec::new();
    for line in r.lines() {
        count = count + 1;
        let line = line.unwrap();

        if line == "" {
            break;
        }

        header.push(line);
    }

    let mut stacks = build_stacks(header).unwrap();

    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);
    for line in r.lines() {
        if count > 0 {
            count = count - 1;
            continue;
        }
        let line = line.unwrap();

        let m = parse_move(line).unwrap();

        for _ in 0..m.num_crates {
            if let Some(c) = stacks[m.from_stack].pop() {
                stacks[m.to_stack].push(c);
            };
        }
    }

    let mut res = String::new();

    for s in stacks.iter_mut() {
        if let Some(c) = s.pop() {
            res.push_str(c.as_str());
        };
    }

    return Ok(res);
}

#[allow(dead_code)]
fn part2(filename: &str) -> Result<String, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut count = 0;

    let mut header: Vec<String> = Vec::new();
    for line in r.lines() {
        count = count + 1;
        let line = line.unwrap();

        if line == "" {
            break;
        }

        header.push(line);
    }

    let mut stacks = build_stacks(header).unwrap();

    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);
    for line in r.lines() {
        if count > 0 {
            count = count - 1;
            continue;
        }
        let line = line.unwrap();

        let m = parse_move(line).unwrap();

        let mut tmp: Vec<String> = Vec::new();

        for _ in 0..m.num_crates {
            if let Some(c) = stacks[m.from_stack].pop() {
                tmp.push(c);
            };
        }

        for c in tmp.iter().rev() {
            stacks[m.to_stack].push(c.to_string())
        }
    }

    let mut res = String::new();

    for s in stacks.iter_mut() {
        if let Some(c) = s.pop() {
            res.push_str(c.as_str());
        };
    }

    return Ok(res);
}

struct Move {
    num_crates: isize,
    from_stack: usize,
    to_stack: usize,
}

fn parse_move(line: String) -> Result<Move, &'static str> {
    let moves: Vec<&str> = line.split(" ").collect();
    let num_crates = moves[1].parse::<isize>().unwrap();
    let from_stack = moves[3].parse::<usize>().unwrap();
    let to_stack = moves[5].parse::<usize>().unwrap();

    return Ok(Move {
        num_crates,
        from_stack,
        to_stack,
    });
}

fn build_stacks(header: Vec<String>) -> Result<Vec<Vec<String>>, &'static str> {
    let mut stacks: Vec<Vec<String>> = Vec::new();
    stacks.push(Vec::new());

    let mut crate_index = 0;
    let mut in_num = false;
    let mut num_arr: Vec<char> = Vec::new();

    for (idx, ch) in header[header.len() - 1].chars().enumerate() {
        if ch.is_digit(10) {
            if !in_num {
                in_num = true;
                crate_index = idx
            }

            num_arr.push(ch);
        }

        if !num_arr.is_empty() {
            in_num = false;
            let mut crates: Vec<String> = Vec::new();

            for (_, v) in header.iter().rev().skip(1).enumerate() {
                // There are no crates for this stack.
                if crate_index > v.len() {
                    break;
                }
                // There are no more crates for this stack.
                if let Some(c) = v.chars().nth(crate_index) {
                    if c == ' ' {
                        break;
                    };
                }

                let end_create_index = match v[crate_index..].find(']') {
                    Some(i) => i,
                    None => return Err("invalid input"),
                };

                crates.push(v[crate_index..end_create_index + crate_index].to_string());
            }

            stacks.push(crates);
            num_arr.clear();
        }
    }

    return Ok(stacks);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(&String::from("../inputs/day05.sample.txt")).unwrap();
        assert_eq!(result, "CMZ");

        let result = part1(&String::from("../inputs/day05.input.txt")).unwrap();
        assert_eq!(result, "WHTLRMZRC");
    }

    #[test]
    fn test_part2() {
        let result = part2(&String::from("../inputs/day05.sample.txt")).unwrap();
        assert_eq!(result, "MCD");

        let result = part2(&String::from("../inputs/day05.input.txt")).unwrap();
        assert_eq!(result, "GMPMLWNMG");
    }
}
