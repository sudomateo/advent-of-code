package day03

import (
	"bufio"
	"io"
)

func Part1(r io.Reader) (int, error) {
	var sum int

	leftPouch := make(map[rune]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		// Take note of what's in the left pouch.
		for _, r := range line[:len(line)/2] {
			if _, ok := leftPouch[r]; !ok {
				leftPouch[r] = struct{}{}
			}
		}

		// Check the right pouch for the duplicate.
		for _, r := range line[len(line)/2:] {
			if _, ok := leftPouch[r]; ok {
				sum += priority(r)
				for key := range leftPouch {
					delete(leftPouch, key)
				}
				break
			}
		}
	}

	return sum, nil
}

func Part2(r io.Reader) (int, error) {
	var sum int

	rucksacks := [2]map[rune]struct{}{
		make(map[rune]struct{}),
		make(map[rune]struct{}),
	}

	group := 1

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		for _, r := range line {

			// We're on the 3rd group, check for the duplicate in groups 1 and 2.
			if group%3 == 0 {
				_, inGroupOne := rucksacks[0][r]
				_, inGroupTwo := rucksacks[1][r]

				if inGroupOne && inGroupTwo {
					sum += priority(r)

					for key := range rucksacks[0] {
						delete(rucksacks[0], key)
					}
					for key := range rucksacks[1] {
						delete(rucksacks[1], key)
					}
				}

				continue
			}

			// Keep track of what's in the current group.
			if _, ok := rucksacks[(group%3)-1][r]; !ok {
				rucksacks[(group%3)-1][r] = struct{}{}
			}
		}

		group++
		if group > 3 {
			group = 1
		}
	}

	return sum, nil
}

// priority calculates the priority of the given rune r.
func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	}
	if r >= 'A' && r <= 'Z' {
		return int(r - 38)
	}
	return 0
}
