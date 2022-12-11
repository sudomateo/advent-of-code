package day09

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type knot struct {
	x int
	y int
}

func Part1(r io.Reader) (int, error) {
	return simulate(r, 2)
}

func Part2(r io.Reader) (int, error) {
	return simulate(r, 10)
}

func simulate(r io.Reader, n int) (int, error) {
	// Build the rope.
	rope := make([]knot, n)
	for i := 0; i < len(rope); i++ {
		rope[i] = knot{}
	}

	visited := make(map[knot]struct{})
	visited[knot{}] = struct{}{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, " ")
		direction := fields[0]
		moves, err := strconv.Atoi(fields[1])
		if err != nil {
			return 0, err
		}

		for move := 0; move < moves; move++ {
			// Move the lead knot.
			switch direction {
			case "U":
				rope[0].y++
			case "R":
				rope[0].x++
			case "D":
				rope[0].y--
			case "L":
				rope[0].x--
			default:
				return 0, fmt.Errorf("invalid direction %s", direction)
			}

			// Move subsequent knots.
			for i := 1; i < len(rope); i++ {
				xDistance := rope[i-1].x - rope[i].x
				yDistance := rope[i-1].y - rope[i].y

				if abs(xDistance) > 1 || abs(yDistance) > 1 {
					rope[i] = knot{
						x: rope[i].x + sign(xDistance),
						y: rope[i].y + sign(yDistance),
					}
				}
			}

			visited[rope[len(rope)-1]] = struct{}{}
		}
	}

	return len(visited), nil
}

func sign(i int) int {
	if i > 0 {
		return 1
	}
	if i < 0 {
		return -1
	}
	return 0
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
