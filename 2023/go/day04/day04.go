package day04

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Part01 implements the solution for part 1.
func Part01(r io.Reader) (string, error) {
	var sum int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		matched, err := matchedNumbers(line)
		if err != nil {
			return "", fmt.Errorf("failed to find matches: %w", err)
		}

		if matched > 0 {
			// Matches to the power of 2.
			sum += 1 << (matched - 1)
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

// matchedNumbers parses the winning numbers and our numbers and calculates how
// many of our numbers matched the winning numbers.
func matchedNumbers(line string) (int, error) {
	// Ignore `Card X:` prefix.
	_, data, found := strings.Cut(line, ":")
	if !found {
		return 0, errors.New("invalid input")
	}

	// Index 0 is space-delimited winning numbers.
	// Index 1 is space-delimited our numbers.
	fields := strings.Split(data, "|")

	// Build a set of winning numbers for O(1) lookups below.
	winningNumbers := make(map[int]struct{})
	for _, number := range strings.Fields(strings.TrimSpace(fields[0])) {
		num, err := strconv.Atoi(number)
		if err != nil {
			return 0, fmt.Errorf("failed to parse winning number: %w", err)
		}
		winningNumbers[num] = struct{}{}
	}

	// Calculate how many of our numbers match the winning numbers.
	var numMatched int
	for _, number := range strings.Fields(strings.TrimSpace(fields[1])) {
		num, err := strconv.Atoi(number)
		if err != nil {
			return 0, fmt.Errorf("failed to parse our number: %w", err)
		}
		if _, ok := winningNumbers[num]; ok {
			numMatched++
		}
	}

	return numMatched, nil
}

// Part02 implements the solution for part 2.
func Part02(r io.Reader) (string, error) {
	var totalCopies int

	// Card number => number of copies won.
	copies := make(map[int]int)

	// To hold cards we need to visit in our breadth-first search.
	queue := make([]int, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		cardNumber, err := parseCardNumber(line)
		if err != nil {
			return "", err
		}

		numMatched, err := matchedNumbers(line)
		if err != nil {
			return "", err
		}

		copies[cardNumber] = numMatched
		queue = append(queue, cardNumber)
	}

	// Breadth-first traversal over the cards to calculate the total number of
	// copies.
	for len(queue) > 0 {
		card := queue[0]
		queue = queue[1:]

		// Each time we visit a card, that's a copy!
		totalCopies++

		numCopies, ok := copies[card]
		if !ok || numCopies == 0 {
			continue
		}

		for i := 1; i <= numCopies; i++ {
			queue = append(queue, card+i)
		}
	}

	return fmt.Sprintf("%d", totalCopies), nil
}

// parseCardNumber parses the number of the given card.
func parseCardNumber(line string) (int, error) {
	data, _, found := strings.Cut(line, ":")
	if !found {
		return 0, errors.New("invalid input")
	}

	fields := strings.Fields(data)
	if len(fields) < 2 {
		return 0, errors.New("invalid input")
	}

	return strconv.Atoi(fields[1])
}
