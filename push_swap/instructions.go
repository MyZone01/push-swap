package pushswap

import "fmt"

func swapStack(stack *[]int, name string) {
	if name != "" {
		Instructions += fmt.Sprintln("s" + name)
	}
	if len(*stack) < 2 {
		return
	}
	(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
	NumberOfInstruction++
}

func pushStack(src *[]int, dst *[]int, name string) {
	if name != "" {
		Instructions += fmt.Sprintln("p" + name)
	}
	if len(*src) == 0 {
		return
	}
	num := (*src)[0]
	*src = (*src)[1:]
	*dst = append([]int{num}, *dst...)
	NumberOfInstruction++
}

func rotateStack(stack *[]int, name string) {
	if name != "" {
		Instructions += fmt.Sprintln("r" + name)
	}
	if len(*stack) < 2 {
		return
	}
	num := (*stack)[0]
	*stack = (*stack)[1:]
	*stack = append(*stack, num)
	NumberOfInstruction++
}

func reverseRotateStack(stack *[]int, name string) {
	if name != "" {
		Instructions += fmt.Sprintln("rr" + name)
	}
	if len(*stack) < 2 {
		return
	}
	num := (*stack)[len(*stack)-1]
	*stack = append([]int{num}, (*stack)[:len(*stack)-1]...)
	NumberOfInstruction++
}
