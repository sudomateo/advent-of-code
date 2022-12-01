package day01_test

import (
	"aoc2022/day01"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolve(t *testing.T) {
	testCases := map[string]struct {
		inputFile        string
		expectedSolution day01.Solution
	}{
		"example input": {
			inputFile: "test/example.txt",
			expectedSolution: day01.Solution{
				Part1: 24000,
				Part2: 45000,
			},
		},
		"real input": {
			inputFile: "test/input.txt",
			expectedSolution: day01.Solution{
				Part1: 72240,
				Part2: 210957,
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			f, err := os.Open(testCase.inputFile)
			if err != nil {
				t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
			}

			solution, err := day01.Solve(f, 3)
			if err != nil {
				t.Fatalf("unexpected error solving challenge: %v", err)
			}

			if diff := cmp.Diff(solution, testCase.expectedSolution); diff != "" {
				t.Fatalf("unexpected diff in solution:\n%s", diff)
			}
		})
	}
}
