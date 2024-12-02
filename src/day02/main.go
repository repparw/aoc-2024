package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 2")
	lines := strings.Split(strings.TrimSpace(input), "\n")

	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

func parseInput(line string) []int {
	fields := strings.Fields(line)
	numbers := make([]int, 0, len(fields))
	
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func part1(lines []string) int {
	safeReports := 0
	for _, line := range lines {
		numbers := parseInput(line)
		if len(numbers) < 2 {
			continue
		}
		if isSafeSequence(numbers) {
			safeReports++
		}
	}
	return safeReports
}

func isSafeSequence(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	isAscending := true
	isDescending := true
	
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff > 0 {
			isDescending = false
			if diff > 3 {
				isAscending = false
				break
			}
		} else if diff < 0 {
			isAscending = false
			if -diff > 3 {
				isDescending = false
				break
			}
		} else {
			isAscending = false
			isDescending = false
			break
		}
	}

	return isAscending || isDescending
}

func part2(lines []string) int {
	safeReports := 0
	for _, line := range lines {
		numbers := parseInput(line)
		if len(numbers) < 3 { // Need at least 3 numbers to remove one
			if isSafeSequence(numbers) {
				safeReports++
			}
			continue
		}

		// Check if original sequence is safe
		if isSafeSequence(numbers) {
			safeReports++
			continue
		}

		// Try removing each number to see if it makes the sequence safe
		for i := 0; i < len(numbers); i++ {
			// Create a new slice without the current number
			testSequence := make([]int, 0, len(numbers)-1)
			testSequence = append(testSequence, numbers[:i]...)
			testSequence = append(testSequence, numbers[i+1:]...)

			if isSafeSequence(testSequence) {
				safeReports++
				break // Found a valid sequence by removing one number, no need to check further
			}
		}
	}
	return safeReports
}
