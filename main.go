package main

import (
	"fmt"
	"os"
	pushswap "pushswap/push_swap"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		fmt.Println("USAGE: go run . \"<list of number>\"")
		return
	}
	elements := arguments[0]
	stack := parseArgs(strings.Split(elements, " "))
	fmt.Println(stack)
	if !sort.IntsAreSorted(stack) {
		if len(stack) <= 5 {
			pushswap.SmallSort(&stack)
		} else {
			pushswap.BigSort(&stack)
		}
	}
	fmt.Println(pushswap.Instructions)
	fmt.Println("We have ",pushswap.NumberOfInstruction, "instructions.")
	fmt.Println(stack)
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
