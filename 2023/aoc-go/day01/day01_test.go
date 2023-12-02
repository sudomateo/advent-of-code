package day01_test

import (
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sudomateo/advent-of-code/2023/aoc-go/day01"
)

type testCase struct {
	testFn         func(io.Reader) (string, error)
	inputFile      string
	expectedOutput string
}

func TestDay01(t *testing.T) {
	tt := map[string]testCase{
		"Part01/Sample": {
			testFn:         day01.Part01,
			inputFile:      "../../inputs/day01/part01-sample.txt",
			expectedOutput: "142",
		},
		"Part01/Final": {
			testFn:         day01.Part01,
			inputFile:      "../../inputs/day01/part01.txt",
			expectedOutput: "54940",
		},
		"Part02/Sample": {
			testFn:         day01.Part02,
			inputFile:      "../../inputs/day01/part02-sample.txt",
			expectedOutput: "281",
		},
		"Part02/Final": {
			testFn:         day01.Part02,
			inputFile:      "../../inputs/day01/part02.txt",
			expectedOutput: "54208",
		},
	}

	for testName, testCase := range tt {
		t.Run(testName, func(t *testing.T) {
			f, err := os.Open(testCase.inputFile)
			if err != nil {
				t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
			}

			solution, err := testCase.testFn(f)
			if err != nil {
				t.Fatalf("unexpected error solving challenge: %v", err)
			}

			if diff := cmp.Diff(solution, testCase.expectedOutput); diff != "" {
				t.Fatalf("unexpected diff in solution:\n%s", diff)
			}
		})
	}
}
