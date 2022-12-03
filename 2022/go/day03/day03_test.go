package day03_test

import (
	"aoc2022/day03"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolve(t *testing.T) {
	testCases := map[string]struct {
		inputFile        string
		expectedSolution day03.Solution
	}{
		"sample input": {
			inputFile: "../../inputs/day03.sample.txt",
			expectedSolution: day03.Solution{
				Part1: 157,
				Part2: 70,
			},
		},
		"real input": {
			inputFile: "../../inputs/day03.input.txt",
			expectedSolution: day03.Solution{
				Part1: 8202,
				Part2: 2864,
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			f, err := os.Open(testCase.inputFile)
			if err != nil {
				t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
			}

			solution, err := day03.Solve(f)
			if err != nil {
				t.Fatalf("unexpected error solving challenge: %v", err)
			}

			if diff := cmp.Diff(solution, testCase.expectedSolution); diff != "" {
				t.Fatalf("unexpected diff in solution:\n%s", diff)
			}
		})
	}
}
