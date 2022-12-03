package day02

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Choice int

const (
	Rock Choice = iota
	Paper
	Scissors
)

// choiceMap maps a Choice to an array of Choice values where index 0 is the
// Choice that the map key wins against and index 1 is the Choice that the
// map key loses against.
var choiceMap = map[Choice][2]Choice{
	Rock:     {Scissors, Paper},
	Paper:    {Rock, Scissors},
	Scissors: {Paper, Rock},
}

// scoreMap maps a Choice to its score value.
var scoreMap = map[Choice]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func Part1(r io.Reader) (int, error) {
	var score int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		choices := strings.Split(line, " ")

		theirChoice, err := getChoice(choices[0])
		if err != nil {
			return 0, err
		}

		myChoice, err := getChoice(choices[1])
		if err != nil {
			return 0, err
		}

		// Draw
		if myChoice == theirChoice {
			score += 3 + scoreMap[myChoice]
			continue
		}

		// Win
		if choiceMap[myChoice][0] == theirChoice {
			score += 6 + scoreMap[myChoice]
			continue
		}

		// Lose
		score += scoreMap[myChoice]
	}

	return score, nil
}

func Part2(r io.Reader) (int, error) {
	var score int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		choices := strings.Split(line, " ")

		theirChoice, err := getChoice(choices[0])
		if err != nil {
			return 0, err
		}

		switch choices[1] {
		// Lose
		case "X":
			score += scoreMap[choiceMap[theirChoice][0]]
		// Draw
		case "Y":
			score += 3 + scoreMap[theirChoice]
		// Win
		case "Z":
			score += 6 + scoreMap[choiceMap[theirChoice][1]]
		default:
		}
	}

	return score, nil
}

// getChoice returns a Choice from the string s, or an error if the string s
// is an invalid choice.
func getChoice(s string) (Choice, error) {
	switch s {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return Choice(-1), errors.New("invalid choice")
	}
}
