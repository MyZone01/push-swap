package main

import (
	"fmt"
	"os"
	"pushswap/checker"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 1 {
		return
	}
	elements := arguments[0]
	stack := parseArgs(strings.Split(elements, " "))
	checker.Run(stack)
}

// Parse input arguments and store them in the stack
func parseArgs(args []string) []int {
	stack := make([]int, len(args))
	for i, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error")
			os.Exit(0)
		}
		stack[i] = num
	}
	if hasDuplicates(stack) {
		fmt.Println("Error")
		os.Exit(0)
	}
	return stack
}

func hasDuplicates(stack []int) bool {
	seen := make(map[int]bool)
	for _, num := range stack {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

// ...
func pushSwap(stack []int) []string {
	instructions := []string{}
	stackB := []int{}

	for len(stack) > 0 {
		minIndex := findMinIndex(&stack)
		if minIndex <= len(stack)/2 {
			for i := 0; i < minIndex; i++ {
				instructions = append(instructions, "ra") // Rotate stack a
			}
		} else {
			for i := 0; i < len(stack)-minIndex; i++ {
				instructions = append(instructions, "rra") // Reverse rotate stack a
			}
		}
		instructions = append(instructions, "pb") // Push to stack b
		stackB = append(stackB, stack[0])
		stack = stack[1:]
	}

	for len(stackB) > 0 {
		maxIndex := findMaxIndex(&stackB)
		if maxIndex <= len(stackB)/2 {
			for i := 0; i < maxIndex; i++ {
				instructions = append(instructions, "rb") // Rotate stack b
			}
		} else {
			for i := 0; i < len(stackB)-maxIndex; i++ {
				instructions = append(instructions, "rrb") // Reverse rotate stack b
			}
		}
		instructions = append(instructions, "pa") // Push to stack a
		stack = append([]int{stackB[0]}, stack...)
		stackB = stackB[1:]
	}

	return instructions
}

func findMinIndex(stack *[]int) int {
	minIndex := 0
	for i, num := range *stack {
		if num < (*stack)[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func findMaxIndex(stack *[]int) int {
	maxIndex := 0
	for i, num := range *stack {
		if num > (*stack)[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// ...
