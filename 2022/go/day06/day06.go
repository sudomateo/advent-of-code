package day06

import (
	"io"
)

func Part1(r io.Reader) (int, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	datastream := string(b)

	return decode(datastream, 4), nil
}

func Part2(r io.Reader) (int, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	datastream := string(b)

	return decode(datastream, 14), nil
}

func decode(datastream string, window int) int {
	freq := make(map[string]int)

	left := 0
	right := 0

	for right < len(datastream) {
		// Build the frequency table for the window.
		if right-left < window {
			if _, ok := freq[string(datastream[right])]; !ok {
				freq[string(datastream[right])] = 1
			} else {
				freq[string(datastream[right])]++
			}
			right++
			continue
		}

		if isWindowUnique(freq) {
			break
		}

		// Move the right side of the window forward.
		if _, ok := freq[string(datastream[right])]; !ok {
			freq[string(datastream[right])] = 1
		} else {
			freq[string(datastream[right])]++
		}

		// Move the left side of the window forward.
		freq[string(datastream[left])]--
		if freq[string(datastream[left])] == 0 {
			delete(freq, string(datastream[left]))
		}

		left++
		right++
	}

	return right
}

func isWindowUnique(freq map[string]int) bool {
	for _, count := range freq {
		if count > 1 {
			return false
		}
	}

	return true
}
