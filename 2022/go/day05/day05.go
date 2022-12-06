package day05

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type move struct {
	numCrates int
	fromStack int
	toStack   int
}

func Part1(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	// Read the header from the input.
	header := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		header = append(header, line)
	}

	stacks, err := buildStacks(header)
	if err != nil {
		return "", err
	}

	// Start moving crates around.
	for scanner.Scan() {
		line := scanner.Text()

		move, err := parseMove(line)
		if err != nil {
			return "", err
		}

		for i := 0; i < move.numCrates; i++ {
			crate := stacks[move.fromStack][len(stacks[move.fromStack])-1]
			stacks[move.fromStack] = stacks[move.fromStack][:len(stacks[move.fromStack])-1]
			stacks[move.toStack] = append(stacks[move.toStack], crate)
		}
	}

	return output(stacks), nil
}

func Part2(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	// Read the header from the input.
	header := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		header = append(header, line)
	}

	stacks, err := buildStacks(header)
	if err != nil {
		return "", err
	}

	for scanner.Scan() {
		line := scanner.Text()

		move, err := parseMove(line)
		if err != nil {
			return "", err
		}

		tmp := make([]string, 0)

		for i := 0; i < move.numCrates; i++ {
			crate := stacks[move.fromStack][len(stacks[move.fromStack])-1]
			stacks[move.fromStack] = stacks[move.fromStack][:len(stacks[move.fromStack])-1]
			tmp = append(tmp, crate)
		}

		for i := len(tmp) - 1; i >= 0; i-- {
			stacks[move.toStack] = append(stacks[move.toStack], tmp[i])
		}
	}

	return output(stacks), nil
}

// buildStacks builds the stack of crates from header.
func buildStacks(header []string) ([][]string, error) {
	stacks := make([][]string, 0)
	stacks = append(stacks, make([]string, 0))

	var crateIndex int
	var inNum bool
	numArr := make([]rune, 0)

	// Build the stacks of crates.
	for i, r := range header[len(header)-1] {
		if unicode.IsDigit(r) {
			if !inNum {
				inNum = true
				crateIndex = i
			}
			numArr = append(numArr, r)
		}

		// We are no longer on a number, but we've found one previously.
		if len(numArr) > 0 {
			inNum = false
			crates := make([]string, 0)

			// Start reading up the column to construct the stack.
			for i := len(header) - 2; i >= 0; i-- {
				// There are no crates for this stack.
				if crateIndex > len(header[i]) {
					break
				}
				// There are no more crates for this stack.
				if string(header[i][crateIndex]) == " " {
					break
				}

				// Crates must be in a box ending in ']'.
				endCrateIdx := strings.IndexRune(header[i][crateIndex:], ']')
				if endCrateIdx == -1 {
					return nil, errors.New("invalid input")
				}

				crates = append(crates, header[i][crateIndex:endCrateIdx+crateIndex])
			}

			// We're done building this stack of crates. Add it to the crane yard.
			stacks = append(stacks, crates)
			numArr = make([]rune, 0)
		}
	}

	return stacks, nil
}

// parseMove parses a move for the crane.
func parseMove(line string) (move, error) {
	moves := strings.Split(line, " ")
	numCrates, err := strconv.Atoi(moves[1])
	if err != nil {
		return move{}, err
	}
	fromStack, err := strconv.Atoi(moves[3])
	if err != nil {
		return move{}, err
	}
	toStack, err := strconv.Atoi(moves[5])
	if err != nil {
		return move{}, err
	}

	m := move{
		numCrates: numCrates,
		fromStack: fromStack,
		toStack:   toStack,
	}

	return m, nil
}

// Output prints the top crates for each stack.
func output(stacks [][]string) string {
	buf := new(bytes.Buffer)

	for i := 0; i < len(stacks); i++ {
		if len(stacks[i]) > 0 {
			str := stacks[i][len(stacks[i])-1]
			buf.WriteString(str)
		}
	}

	return buf.String()
}
