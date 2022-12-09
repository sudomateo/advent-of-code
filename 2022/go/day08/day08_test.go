package day08_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"aoc2022/day08"
)

func TestDay08(t *testing.T) {
	testCases := map[string]struct {
		inputFile string
		part1     int
		part2     int
	}{
		"SampleInput": {
			inputFile: "../../inputs/day08.sample.txt",
			part1:     21,
			part2:     8,
		},
		"RealInput": {
			inputFile: "../../inputs/day08.input.txt",
			part1:     1796,
			part2:     288120,
		},
	}

	t.Run("Part1", func(t *testing.T) {
		for testName, testCase := range testCases {
			t.Run(testName, func(t *testing.T) {
				f, err := os.Open(testCase.inputFile)
				if err != nil {
					t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
				}

				solution, err := day08.Part1(f)
				if err != nil {
					t.Fatalf("unexpected error solving challenge: %v", err)
				}

				if diff := cmp.Diff(solution, testCase.part1); diff != "" {
					t.Fatalf("unexpected diff in solution:\n%s", diff)
				}
			})
		}
	})

	t.Run("Part2", func(t *testing.T) {
		for testName, testCase := range testCases {
			t.Run(testName, func(t *testing.T) {
				f, err := os.Open(testCase.inputFile)
				if err != nil {
					t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
				}

				solution, err := day08.Part2(f)
				if err != nil {
					t.Fatalf("unexpected error solving challenge: %v", err)
				}

				if diff := cmp.Diff(solution, testCase.part2); diff != "" {
					t.Fatalf("unexpected diff in solution:\n%s", diff)
				}
			})
		}
	})
}
