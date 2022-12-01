package day01

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"strconv"
)

// Solution represents the solution to the Advent of Code challenge for day 1.
type Solution struct {
	Part1 int
	Part2 int
}

// Solve solves the Advent of Code challenge for day 1.
func Solve(r io.Reader, k int) (Solution, error) {
	calories := &IntHeap{}
	heap.Init(calories)

	var sum int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			heap.Push(calories, sum)
			sum = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return Solution{}, err
		}

		sum += n
	}
	heap.Push(calories, sum)

	if calories.Len() < k {
		return Solution{}, fmt.Errorf("cannot count %d elves with input size %d", k, calories.Len())
	}

	part1 := heap.Pop(calories).(int)
	part2 := part1 + int(heap.Pop(calories).(int)) + int(heap.Pop(calories).(int))

	s := Solution{
		Part1: part1,
		Part2: part2,
	}

	return s, nil
}

// IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
