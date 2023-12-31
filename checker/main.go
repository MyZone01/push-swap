package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	Run(stack)
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

func Run(stack []int) {
	//Parse input arguments and store them in the stack
	scanner := bufio.NewScanner(os.Stdin)
	instructions := []string{}

	// Read instructions from stdin
	for scanner.Scan() {
		instruction := scanner.Text()
		if len(instruction) == 0 {
			continue
		}
		if !isValidInstruction(instruction) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(0)
		}
		instructions = append(instructions, instruction)
	}

	// Check if the instructions sort the stack correctly
	if checkInstructions(stack, instructions) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func isValidInstruction(instruction string) bool {
	validInstructions := map[string]bool{
		"sa": true, "sb": true, "ss": true,
		"pa": true, "pb": true,
		"ra": true, "rb": true, "rr": true,
		"rra": true, "rrb": true, "rrr": true,
	}
	_, valid := validInstructions[instruction]
	return valid
}

func checkInstructions(stack []int, instructions []string) bool {
	stackA := make([]int, len(stack))
	stackB := make([]int, len(stack))
	copy(stackA, stack)

	for _, instr := range instructions {
		switch instr {
		case "sa":
			swapStack(&stackA)
		case "sb":
			swapStack(&stackB)
		case "ss":
			swapStack(&stackA)
			swapStack(&stackB)
		case "pa":
			pushStack(&stackB, &stackA)
		case "pb":
			pushStack(&stackA, &stackB)
		case "ra":
			rotateStack(&stackA)
		case "rb":
			rotateStack(&stackB)
		case "rr":
			rotateStack(&stackA)
			rotateStack(&stackB)
		case "rra":
			reverseRotateStack(&stackA)
		case "rrb":
			reverseRotateStack(&stackB)
		case "rrr":
			reverseRotateStack(&stackA)
			reverseRotateStack(&stackB)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	sort.Ints(stack)
	return isSorted(stackA, stack)
}

func swapStack(stack *[]int) {
	if len(*stack) < 2 {
		return
	}
	(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
}

func pushStack(src *[]int, dst *[]int) {
	if len(*src) == 0 {
		return
	}
	num := (*src)[0]
	*src = (*src)[1:]
	*dst = append([]int{num}, *dst...)
}

func rotateStack(stack *[]int) {
	if len(*stack) < 2 {
		return
	}
	num := (*stack)[0]
	*stack = (*stack)[1:]
	*stack = append(*stack, num)
}

func reverseRotateStack(stack *[]int) {
	if len(*stack) < 2 {
		return
	}
	num := (*stack)[len(*stack)-1]
	*stack = append([]int{num}, (*stack)[:len(*stack)-1]...)
}

func isSorted(stack, sort []int) bool {
	for i := 0; i < len(stack); i++ {
		if stack[i] != sort[i] {
			return false
		}
	}
	return true
}
