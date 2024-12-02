package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 1")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

func part1(lines []string) int {
	col1 := make([]int, 0, len(lines))
	col2 := make([]int, 0, len(lines))

	// Parse input into two columns
	for _, line := range lines {
		nums := strings.Fields(line)
		if len(nums) != 2 {
			continue
		}
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		col1 = append(col1, n1)
		col2 = append(col2, n2)
	}

	// Sort both columns
	sort.Ints(col1)
	sort.Ints(col2)

	// Calculate sum of absolute differences
	sum := 0
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}

func part2(lines []string) int {
	col1 := make([]int, 0, len(lines))
	col2 := make([]int, 0, len(lines))

	// Parse input into two columns
	for _, line := range lines {
		nums := strings.Fields(line)
		if len(nums) != 2 {
			continue
		}
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		col1 = append(col1, n1)
		col2 = append(col2, n2)
	}

	// Count frequencies of numbers in right list
	freqMap := make(map[int]int)
	for _, num := range col2 {
		freqMap[num]++
	}

	// Calculate similarity score
	score := 0
	for _, num := range col1 {
		score += num * freqMap[num]
	}

	return score
}
