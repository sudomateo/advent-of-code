package day03

import (
	"bufio"
	"io"
)

// Solution represents the solution to the Advent of Code challenge.
type Solution struct {
	Part1 int
	Part2 int
}

// Solve solves the Advent of Code challenge.
func Solve(r io.Reader) (Solution, error) {
	var sum1 int
	var sum2 int

	leftPouch := make(map[rune]struct{})

	rucksacks := [2]map[rune]struct{}{
		make(map[rune]struct{}),
		make(map[rune]struct{}),
	}
	group := 1

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		for i, r := range line {
			// Part 1.
			if i < len(line)/2 {
				if _, ok := leftPouch[r]; !ok {
					leftPouch[r] = struct{}{}
				}
			} else {
				if _, ok := leftPouch[r]; ok {
					sum1 += score(r)
					for key := range leftPouch {
						delete(leftPouch, key)
					}
				}
			}

			// Part 2.

			// We're on the last group, check for the common type in all groups.
			if group%3 == 0 {
				_, inGroupOne := rucksacks[0][r]
				_, inGroupTwo := rucksacks[1][r]

				if inGroupOne && inGroupTwo {
					sum2 += score(r)

					for key := range rucksacks[0] {
						delete(rucksacks[0], key)
					}
					for key := range rucksacks[1] {
						delete(rucksacks[1], key)
					}
				}

				continue
			}

			if _, ok := rucksacks[(group%3)-1][r]; !ok {
				rucksacks[(group%3)-1][r] = struct{}{}
			}
		}

		group++
		if group > 3 {
			group = 1
		}
	}

	s := Solution{
		Part1: sum1,
		Part2: sum2,
	}

	return s, nil
}

func score(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	}
	if r >= 'A' && r <= 'Z' {
		return int(r - 38)
	}
	return 0
}
