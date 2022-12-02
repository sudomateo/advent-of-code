package day02

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Solution represents the solution to the Advent of Code challenge.
type Solution struct {
	Part1 int
	Part2 int
}

type Choice string

const (
	Rock     Choice = "rock"
	Paper    Choice = "paper"
	Scissors Choice = "scissors"
)

// choiceMap maps a Choice an array of Choice values where index 0 is the
// Choice that the map key wins against and index 1 is the Choice that the map
// key loses against.
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

// Solve solves the Advent of Code challenge.
func Solve(r io.Reader) (Solution, error) {
	var score1 int
	var score2 int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, " ")

		theirChoice, err := getChoice(fields[0])
		if err != nil {
			return Solution{}, err
		}

		myChoice, err := getChoice(fields[1])
		if err != nil {
			return Solution{}, err
		}

		score1 += scoreOne(myChoice, theirChoice)
		score2 += scoreTwo(theirChoice, fields[1])
	}

	s := Solution{
		Part1: score1,
		Part2: score2,
	}

	return s, nil
}

// getChoice returns a Choice from the string s, or an error if the string s is
// an invalid choice.
func getChoice(s string) (Choice, error) {
	switch s {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	default:
		return Choice(""), errors.New("invalid choice")
	}
}

// scoreOne calculates the score as required in part 1 of the problem.
func scoreOne(myChoice Choice, theirChoice Choice) int {
	// Draw
	if myChoice == theirChoice {
		return 3 + scoreMap[myChoice]
	}

	// Win
	if choiceMap[myChoice][0] == theirChoice {
		return 6 + scoreMap[myChoice]
	}

	// Lose
	return scoreMap[myChoice]
}

// scoreTwo calculates the score as required in part 2 of the problem.
func scoreTwo(theirChoice Choice, s string) int {
	switch s {
	// Lose
	case "X":
		return scoreMap[choiceMap[theirChoice][0]]
	// Draw
	case "Y":
		return 3 + scoreMap[theirChoice]
	// Win
	case "Z":
		return 6 + scoreMap[choiceMap[theirChoice][1]]
	default:
		return 0
	}
}
