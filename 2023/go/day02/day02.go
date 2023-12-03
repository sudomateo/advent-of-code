package day02

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// cubes represents the color and count of cubes either drawn from a bag or
// within a bag.
type cubes map[string]int

func Part01(r io.Reader) (string, error) {
	var sum int

	bag := cubes{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(r)
LOOP: // Labeled to short-circuit when a drawing is not possible.
	for scanner.Scan() {
		line := scanner.Text()

		game, err := parseGame(line)
		if err != nil {
			return "", fmt.Errorf("failed parsing game number: %w", err)
		}

		drawings, err := parseDrawings(line)
		if err != nil {
			return "", fmt.Errorf("failed parsing drawings: %w", err)
		}

		for _, drawing := range drawings {
			if !drawingIsPossible(bag, drawing) {
				goto LOOP
			}
		}

		sum += game
	}

	return fmt.Sprintf("%d", sum), nil
}

func Part02(r io.Reader) (string, error) {
	var sum int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		drawings, err := parseDrawings(line)
		if err != nil {
			return "", fmt.Errorf("failed parsing drawings: %w", err)
		}

		power := 1
		for _, count := range minimumBag(drawings) {
			power *= count
		}

		sum += power
	}

	return fmt.Sprintf("%d", sum), nil
}

// drawingIsPossible determines whether it's possible to make the drawing from
// the given bag.
func drawingIsPossible(bag map[string]int, drawing map[string]int) bool {
	for color, count := range drawing {
		if _, ok := bag[color]; !ok {
			return false
		}

		if count > bag[color] {
			return false
		}
	}

	return true
}

// parseGame calculates the game number from the given input.
func parseGame(input string) (int, error) {
	fields := strings.Split(input, ":")
	gameFields := strings.Split(strings.TrimSpace(fields[0]), " ")
	return strconv.Atoi(gameFields[1])
}

// parseDrawings calculates the drawings from the given input.
func parseDrawings(input string) ([]cubes, error) {
	drawings := make([]cubes, 0)

	fields := strings.Split(input, ":")

	drawingFields := strings.Split(strings.TrimSpace(fields[1]), ";")
	for _, field := range drawingFields {
		d := make(cubes)

		cubes := strings.Split(field, ",")
		for _, cube := range cubes {
			details := strings.Split(strings.TrimSpace(cube), " ")

			quantity, err := strconv.Atoi(details[0])
			if err != nil {
				return nil, err
			}

			color := details[1]

			if _, ok := d[color]; !ok {
				d[color] = 0
			}
			d[color] += quantity
		}

		if len(d) > 0 {
			drawings = append(drawings, d)
		}
	}

	return drawings, nil
}

// minimumBag calculates the minimum set of cubes needed in the bag to support
// the drawings.
func minimumBag(drawings []cubes) cubes {
	minBag := make(cubes)

	for _, drawing := range drawings {
		for color, count := range drawing {
			// This is the first time we've seen this cube. Add it to the bag.
			if _, ok := minBag[color]; !ok {
				minBag[color] = count
				continue
			}

			// Determine whether this bag needs of a certain cube.
			minBag[color] = max(minBag[color], count)
		}
	}

	return minBag
}
