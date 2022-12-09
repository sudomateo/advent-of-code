package day08

import (
	"bufio"
	"io"
	"strconv"
)

type tree struct {
	height int
}

func Part1(r io.Reader) (int, error) {
	graph, err := buildForest(r)
	if err != nil {
		return 0, err
	}

	var numVisible int

	for y := 0; y < len(graph); y++ {
		for x := 0; x < len(graph[0]); x++ {
			if _, visible := visibilityScore(y, x, graph); visible {
				numVisible++
			}
		}
	}

	return numVisible, nil
}

func Part2(r io.Reader) (int, error) {
	graph, err := buildForest(r)
	if err != nil {
		return 0, err
	}

	var maxScore int

	for y := 0; y < len(graph); y++ {
		for x := 0; x < len(graph[0]); x++ {
			score, _ := visibilityScore(y, x, graph)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore, nil
}

func buildForest(r io.Reader) ([][]tree, error) {
	forest := make([][]tree, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]tree, 0, len(line))

		for _, r := range line {
			height, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}

			t := tree{
				height: height,
			}

			row = append(row, t)
		}

		forest = append(forest, row)
	}

	return forest, nil
}

func visibilityScore(y int, x int, graph [][]tree) (int, bool) {
	height := graph[y][x].height

	var topScore int
	topVisible := true
	for newY := y - 1; newY >= 0; newY-- {
		newHeight := graph[newY][x].height
		if height <= newHeight {
			topScore++
			topVisible = false
			break
		}
		topScore++
	}

	var rightScore int
	rightVisible := true
	for newX := x + 1; newX < len(graph[0]); newX++ {
		newHeight := graph[y][newX].height
		if height <= newHeight {
			rightScore++
			rightVisible = false
			break
		}
		rightScore++
	}

	var bottomScore int
	bottomVisible := true
	for newY := y + 1; newY < len(graph); newY++ {
		newHeight := graph[newY][x].height
		if height <= newHeight {
			bottomScore++
			bottomVisible = false
			break
		}
		bottomScore++
	}

	var leftScore int
	leftVisible := true
	for newX := x - 1; newX >= 0; newX-- {
		newHeight := graph[y][newX].height
		if height <= newHeight {
			leftScore++
			leftVisible = false
			break
		}
		leftScore++
	}

	score := topScore * rightScore * bottomScore * leftScore
	visible := topVisible || rightVisible || leftVisible || bottomVisible

	return score, visible
}
