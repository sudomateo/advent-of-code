package day02_test

import (
	"aoc2022/day02"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolve(t *testing.T) {
	testCases := map[string]struct {
		inputFile        string
		expectedSolution day02.Solution
	}{
		"sample input": {
			inputFile: "../../inputs/day02.sample.txt",
			expectedSolution: day02.Solution{
				Part1: 15,
				Part2: 12,
			},
		},
		"real input": {
			inputFile: "../../inputs/day02.input.txt",
			expectedSolution: day02.Solution{
				Part1: 12458,
				Part2: 12683,
			},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			f, err := os.Open(testCase.inputFile)
			if err != nil {
				t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
			}

			solution, err := day02.Solve(f)
			if err != nil {
				t.Fatalf("unexpected error solving challenge: %v", err)
			}

			if diff := cmp.Diff(solution, testCase.expectedSolution); diff != "" {
				t.Fatalf("unexpected diff in solution:\n%s", diff)
			}
		})
	}
}
