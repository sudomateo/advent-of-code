package day01

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func Part01(r io.Reader) (string, error) {
	var sum int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		numsInLine, err := parseLinePart01(line)
		if err != nil {
			return "", fmt.Errorf("failed parsing numbers in line: %w", err)
		}
		sum += (numsInLine[0] * 10) + numsInLine[len(numsInLine)-1]
	}

	return fmt.Sprintf("%d", sum), nil
}

// parseLinePart01 parses all of the numbers in a given line. Numbers are
// single digits (e.g., 0, 1).
func parseLinePart01(line string) ([]int, error) {
	nums := make([]int, 0)

	for _, r := range line {
		if unicode.IsDigit(r) {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
	}

	return nums, nil
}

func Part02(r io.Reader) (string, error) {
	var sum int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		numsInLine, err := parseLinePart02(line)
		if err != nil {
			return "", fmt.Errorf("failed parsing numbers in line: %w", err)
		}
		sum += (numsInLine[0] * 10) + numsInLine[len(numsInLine)-1]
	}

	return fmt.Sprintf("%d", sum), nil
}

// parseLinePart02 uses a sliding window to parse all of the numbers in a given
// line. Numbers can be single digits (e.g., 0, 1) or the lowercase word of a
// number (e.g., zero, one).
func parseLinePart02(line string) ([]int, error) {
	mapNums := map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	nums := make([]int, 0)

	l := 0
	r := 0

	for r < len(line) {
		// This window is invalid, pull the left pointer in but not if it'll
		// intersect with the right pointer.
		for invalidWindow(string(line[l:r+1]), mapNums) && l < r {
			l++
		}

		// We have a valid window. Parse its number.
		window := string(line[l : r+1])
		if num, err := parseNumber(window, mapNums); err == nil {
			nums = append(nums, num)
		}

		// Expand the window to the right.
		r++
	}

	return nums, nil
}

// invalidWindow determines if a given window is invalid.
func invalidWindow(window string, mapNums map[string]int) bool {
	// If the window is a 1 digit character it it can't be invalid.
	if len(window) == 1 {
		if unicode.IsDigit(rune(window[0])) {
			return false
		}
	}

	// If the window is a prefix of a valid number it can't be invalid.
	for key := range mapNums {
		if strings.HasPrefix(key, window) {
			return false
		}
	}

	// This must be an invalid window.
	return true
}

// parseNumber parses a number out of the window.
func parseNumber(window string, mapNums map[string]int) (int, error) {
	if len(window) == 1 {
		if unicode.IsDigit(rune(window[0])) {
			num, err := strconv.Atoi(window)
			if err != nil {
				return 0, err
			}
			return num, nil
		}
	}

	if num, ok := mapNums[window]; ok {
		return num, nil
	}

	return 0, errors.New("no number in window")
}
