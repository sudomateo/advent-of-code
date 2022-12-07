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
	if len(datastream) <= window {
		return 0
	}

	seen := make(map[rune]struct{})

	for l, r := 0, window; r < len(datastream); l, r = l+1, r+1 {
		for _, num := range datastream[l:r] {
			seen[num] = struct{}{}
		}

		if len(seen) == window {
			return l + window
		}

		for k := range seen {
			delete(seen, k)
		}
	}

	return 0
}
