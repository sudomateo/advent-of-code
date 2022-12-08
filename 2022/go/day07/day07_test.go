package day07_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"aoc2022/day07"
)

func TestDay07(t *testing.T) {
	testCases := map[string]struct {
		inputFile string
		part1     int
		part2     int
	}{
		"SampleInput": {
			inputFile: "../../inputs/day07.sample.txt",
			part1:     95437,
			part2:     24933642,
		},
		"RealInput": {
			inputFile: "../../inputs/day07.input.txt",
			part1:     2031851,
			part2:     2568781,
		},
	}

	t.Run("Part1", func(t *testing.T) {
		for testName, testCase := range testCases {
			t.Run(testName, func(t *testing.T) {
				f, err := os.Open(testCase.inputFile)
				if err != nil {
					t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
				}

				solution, err := day07.Part1(f)
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

				solution, err := day07.Part2(f)
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
