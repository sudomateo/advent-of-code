package day05_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"aoc2022/day05"
)

func TestDay05(t *testing.T) {
	testCases := map[string]struct {
		inputFile string
		part1     string
		part2     string
	}{
		"SampleInput": {
			inputFile: "../../inputs/day05.sample.txt",
			part1:     "CMZ",
			part2:     "MCD",
		},
		"RealInput": {
			inputFile: "../../inputs/day05.input.txt",
			part1:     "WHTLRMZRC",
			part2:     "GMPMLWNMG",
		},
	}

	t.Run("Part1", func(t *testing.T) {
		for testName, testCase := range testCases {
			t.Run(testName, func(t *testing.T) {
				f, err := os.Open(testCase.inputFile)
				if err != nil {
					t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
				}

				solution, err := day05.Part1(f)
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

				solution, err := day05.Part2(f)
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
