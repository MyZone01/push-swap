package checker

import (
	"bufio"
	"fmt"
	"os"
)

func Run(stack []int) {

	//Parse input arguments and store them in the stack
	scanner := bufio.NewScanner(os.Stdin)
	instructions := []string{}

	// Read instructions from stdin
	for scanner.Scan() {
		instruction := scanner.Text()
		if !isValidInstruction(instruction) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
		instructions = append(instructions, instruction)
	}
	fmt.Printf("%#v\n", instructions)

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

	return isSorted(stackA)
}

func swapStack(stack *[]int) {
	if len(*stack) < 2 {
		return // Implement the reverse rotate operation for both stacks
	}
	(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
}

// Take the top first of src and push in to dst
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

// Implement the other stack operations (push, rotate, reverse rotate) similarly

func isSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[i-1] {
			return false
		}
	}
	return true
}