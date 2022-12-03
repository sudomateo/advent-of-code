package day01

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"strconv"
)

func Part1(r io.Reader) (int, error) {
	var caloriesSum int
	var maxCalories int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if caloriesSum > maxCalories {
				maxCalories = caloriesSum
			}
			caloriesSum = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		caloriesSum += n
	}

	return maxCalories, nil
}

func Part2(r io.Reader) (int, error) {
	calorieHeap := &IntHeap{}
	heap.Init(calorieHeap)

	var calorieSum int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			heap.Push(calorieHeap, calorieSum)
			calorieSum = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}

		calorieSum += n
	}
	heap.Push(calorieHeap, calorieSum)

	if calorieHeap.Len() < 3 {
		return 0, fmt.Errorf("cannot count 3 elves with input size %d", calorieHeap.Len())
	}

	var totalCalories int
	for i := 0; i < 3; i++ {
		totalCalories += heap.Pop(calorieHeap).(int)
	}

	return totalCalories, nil
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
