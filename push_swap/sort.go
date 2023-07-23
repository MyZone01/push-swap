package pushswap

import (
	"sort"
)

func PushSwap3Number(stack *[]int) {
	if (*stack)[2] > (*stack)[0] && (*stack)[2] > (*stack)[1] && (*stack)[1] < (*stack)[0] {
		swapStack(stack, "a")
	} else if (*stack)[2] < (*stack)[0] && (*stack)[2] < (*stack)[1] && (*stack)[1] < (*stack)[0] {
		swapStack(stack, "a")
		reverseRotateStack(stack, "a")
	} else if (*stack)[2] < (*stack)[0] && (*stack)[2] > (*stack)[1] && (*stack)[1] < (*stack)[0] {
		rotateStack(stack, "a")
	} else if (*stack)[2] > (*stack)[0] && (*stack)[2] < (*stack)[1] && (*stack)[1] > (*stack)[0] {
		swapStack(stack, "a")
		rotateStack(stack, "a")
	} else if (*stack)[2] < (*stack)[0] && (*stack)[2] < (*stack)[1] && (*stack)[1] > (*stack)[0] {
		reverseRotateStack(stack, "a")
	}
}

func SmallSort(stack *[]int) {
	if len(*stack) == 2 {
		if !sort.IntsAreSorted(*stack) {
			swapStack(stack, "a")
		}
	} else {
		stackA := stack
		stackB := make([]int, len(*stack))
		for len(*stackA) != 3 {
			pushStack(stackA, &stackB, "b")
		}
		PushSwap3Number(stackA)
		for !isVoid(stackB) || !sort.IntsAreSorted(*stackA) {
			pushStack(&stackB, stackA, "a")
			sortSmallStack(stackA)
		}
		stack = stackA
	}
}

func sortSmallStack(stackA *[]int) {
	if (*stackA)[0] <= (*stackA)[1] {
		return
	} else if (*stackA)[0] >= (*stackA)[len((*stackA))-1] {
		rotateStack(stackA, "a")
	} else {
		numberOfRA := 0
		for !sort.IntsAreSorted(*stackA) {
			if (*stackA)[0] >= (*stackA)[1] {
				swapStack(stackA, "a")
				if (*stackA)[1] >= (*stackA)[2] {
					rotateStack(stackA, "a")
					numberOfRA++
				} else {
					break
				}
			} else {
				break
			}
		}
		for i := 0; i < numberOfRA; i++ {
			reverseRotateStack(stackA, "a")
		}
	}
}
