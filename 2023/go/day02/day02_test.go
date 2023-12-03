package day02_test

import (
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sudomateo/advent-of-code/2023/go/day02"
)

type testCase struct {
	testFn         func(io.Reader) (string, error)
	inputFile      string
	expectedOutput string
}

func TestDay02(t *testing.T) {
	tt := map[string]testCase{
		"Part01/Sample": {
			testFn:         day02.Part01,
			inputFile:      "../../inputs/day02/part01-sample.txt",
			expectedOutput: "8",
		},
		"Part01/Final": {
			testFn:         day02.Part01,
			inputFile:      "../../inputs/day02/part01.txt",
			expectedOutput: "2348",
		},
		"Part02/Sample": {
			testFn:         day02.Part02,
			inputFile:      "../../inputs/day02/part02-sample.txt",
			expectedOutput: "2286",
		},
		"Part02/Final": {
			testFn:         day02.Part02,
			inputFile:      "../../inputs/day02/part02.txt",
			expectedOutput: "76008",
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
