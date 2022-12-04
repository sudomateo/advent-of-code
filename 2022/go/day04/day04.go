package day04

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part1(r io.Reader) (int, error) {
	var sum int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, ",")

		elfOne, err := helper(fields[0])
		if err != nil {
			return 0, err
		}
		elfTwo, err := helper(fields[1])
		if err != nil {
			return 0, err
		}

		var elfOneCount int
		for _, v := range elfOne {
			if v == 1 {
				elfOneCount++
			}
		}

		var elfTwoCount int
		for _, v := range elfTwo {
			if v == 1 {
				elfTwoCount++
			}
		}

		small := elfOne
		large := elfTwo
		if elfTwoCount < elfOneCount {
			small = elfTwo
			large = elfOne
		}

		totalOverlap := true 
		for i := range small {
			if small[i] == 0 {
				continue
			}

			if large[i] != 1 {
				totalOverlap = false
				break
			}
		}

		if totalOverlap {
			sum++
		}
	}

	return sum, nil
}

func Part2(r io.Reader) (int, error) {
	var sum int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, ",")

		elfOne, err := helper(fields[0])
		if err != nil {
			return 0, err
		}
		elfTwo, err := helper(fields[1])
		if err != nil {
			return 0, err
		}

		for i := 0; i < 100; i++ {
			if (elfOne[i] == 1) && (elfTwo[i] == 1) {
				sum++
				break
			}
		}
	}

	return sum, nil
}

func helper(s string) ([]int, error) {
	fields := strings.Split(s, "-")

	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid input %s", s)
	}

	num1, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, err
	}

	num2, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}

	res := make([]int, 100, 100)

	for num1 <= num2 {
		res[num1] = 1
		num1++
	}

	return res, nil
}
