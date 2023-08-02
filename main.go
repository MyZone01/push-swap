package main

import (
	"fmt"
	"os"
	pushswap "pushswap/push_swap"
	"sort"
	"strconv"
	"strings"
)

// ANSI escape code for setting text color to red
const redColor = "\033[31m"

// ANSI escape code for resetting text color to default
const resetColor = "\033[0m"

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		fmt.Println("USAGE: go run . \"<list of number>\"")
		return
	}
	elements := arguments[0]
	stack := parseArgs(strings.Split(elements, " "))
	if !sort.IntsAreSorted(stack) {
		if len(stack) <= 6 {
			pushswap.SmallSort(&stack)
		} else {
			pushswap.BigSort(&stack)
		}
	}
	fmt.Println(pushswap.Instructions)
	if pushswap.NumberOfInstruction > 12 {
		fmt.Printf("%sWe have %d instructions %s%v\n", redColor, pushswap.NumberOfInstruction, resetColor, stack)
	} else {
		fmt.Println("We have", pushswap.NumberOfInstruction, "instructions.", stack)
	}
}

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
