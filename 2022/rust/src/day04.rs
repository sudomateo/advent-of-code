use std::{
    fs::File,
    io::{BufRead, BufReader},
};

fn helper(s: &str) -> Result<[isize; 100], &'static str> {
    let fields: Vec<&str> = s.split("-").collect();

    if fields.len() < 2 {
        return Err("invalid input");
    }

    let mut num1 = fields[0].parse::<isize>().unwrap();
    let num2 = fields[1].parse::<isize>().unwrap();

    let mut res: [isize; 100] = [0; 100];

    while num1 <= num2 {
        res[num1 as usize] = 1;
        num1 = num1 + 1;
    }

    Ok(res)
}

#[allow(dead_code)]
fn part1(filename: String) -> Result<isize, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut sum: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        let fields: Vec<&str> = line.split(",").collect();

        let elf_one = helper(fields[0]).unwrap();
        let elf_two = helper(fields[1]).unwrap();

        let mut elf_one_count: isize = 0;
        for v in elf_one.iter() {
            if *v == 1 {
                elf_one_count = elf_one_count + 1;
            }
        }

        let mut elf_two_count: isize = 0;
        for v in elf_two.iter() {
            if *v == 1 {
                elf_two_count = elf_two_count + 1;
            }
        }

        let mut small = elf_one;
        let mut large = elf_two;
        if elf_two_count < elf_one_count {
            small = elf_two;
            large = elf_one;
        }

        let mut total_overlap = true;
        for (i, _) in small.iter().enumerate() {
            if small[i] == 0 {
                continue
            }

            if large[i] != 1 {
                total_overlap = false;
                break
            }
        }

        if total_overlap {
            sum = sum + 1;
        }
    }

    return Ok(sum);
}

#[allow(dead_code)]
fn part2(filename: String) -> Result<isize, &'static str> {
    let f = File::open(filename).unwrap();
    let r = BufReader::new(f);

    let mut sum: isize = 0;

    for line in r.lines() {
        let line = line.unwrap();

        let fields: Vec<&str> = line.split(",").collect();

        let elf_one = helper(fields[0]).unwrap();
        let elf_two = helper(fields[1]).unwrap();

        for i in 0..100 {
            if elf_one[i] == 1 && elf_two[i] == 1 {
                sum = sum + 1;
                break;
            }
        }
    }

    return Ok(sum);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let result = part1(String::from("../inputs/day04.sample.txt")).unwrap();
        assert_eq!(result, 2);

        let result = part1(String::from("../inputs/day04.input.txt")).unwrap();
        assert_eq!(result, 496);
    }

    #[test]
    fn test_part2() {
        let result = part2(String::from("../inputs/day04.sample.txt")).unwrap();
        assert_eq!(result, 4);

        let result = part2(String::from("../inputs/day04.input.txt")).unwrap();
        assert_eq!(result, 847);
    }
}
