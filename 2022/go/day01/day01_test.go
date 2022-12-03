package day01_test

import (
	"aoc2022/day01"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay1(t *testing.T) {
	testCases := map[string]struct {
		inputFile string
		part1     int
		part2     int
	}{
		"SampleInput": {
			inputFile: "../../inputs/day01.sample.txt",
			part1:     24000,
			part2:     45000,
		},
		"RealInput": {
			inputFile: "../../inputs/day01.input.txt",
			part1:     72240,
			part2:     210957,
		},
	}

	t.Run("Part1", func(t *testing.T) {
		for testName, testCase := range testCases {
			t.Run(testName, func(t *testing.T) {
				f, err := os.Open(testCase.inputFile)
				if err != nil {
					t.Fatalf("failed opening input file %s: %v", testCase.inputFile, err)
				}

				solution, err := day01.Part1(f)
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

				solution, err := day01.Part2(f)
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
