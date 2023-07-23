package pushswap

import (
	"fmt"
	"sort"
)

var (
	NumberOfInstruction int = 0
	Instructions string
	stackA *[]int
	stackB *[]int
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

func SmallSort(stackA *[]int) {
	if len(*stackA) == 2 {
		swapStack(stackA, "a")
	} else {
		stackB := []int{}
		for len(*stackA) != 3 {
			pushStack(stackA, &stackB, "b")
		}
		fmt.Println(&stackB)
		if !IsDecreaseSorted(&stackB) {
			swapStack(&stackB, "b")
		}
		PushSwap3Number(stackA)
		for !isVoid(stackB) || !sort.IntsAreSorted(*stackA) {
			pushStack(&stackB, stackA, "a")
			if len(*stackA) > 1 {
				sortStackA(stackA)
			}
		}
	}
}

func BigSort(stackA *[]int) {
	stackB := []int{}
	GoToStackB(stackA, &stackB)
	GoToStackA(stackA, &stackB)
	fmt.Println(stackA)
	fmt.Println(stackB)
}

func GoToStackB(stackA, stackB *[]int) {
	for !isVoid(*stackA) || !IsDecreaseSorted(stackB) {
		pushStack(stackA, stackB, "b")
		if len(*stackB) > 1 {
			sortStackB(stackB)
		}
	}
}

func IsDecreaseSorted(stack *[]int) bool {
	for i := 0; i < len(*stack)-1; i++ {
		if (*stack)[i] < (*stack)[i+1] {
			return false
		}
	}
	return true
}

func GoToStackA(stackA, stackB *[]int) {

}

func sortStackB(stackB *[]int) {
	if (*stackB)[0] >= (*stackB)[1] {
		return
	} else if (*stackB)[0] <= (*stackB)[len((*stackB))-1] {
		rotateStack(stackB, "a")
	} else {
		numberOfRA := 0
		for !sort.IntsAreSorted(*stackB) {
			if len(*stackB) >= 2 && (*stackB)[0] <= (*stackB)[1] {
				swapStack(stackB, "a")
				if len(*stackB) >= 3 && (*stackB)[1] <= (*stackB)[2] {
					rotateStack(stackB, "a")
					numberOfRA++
				} else {
					break
				}
			} else {
				break
			}
		}
		for i := 0; i < numberOfRA; i++ {
			reverseRotateStack(stackB, "a")
		}
	}
}

func sortStackA(stackA *[]int) {
	if (*stackA)[0] <= (*stackA)[1] {
		return
	} else if (*stackA)[0] >= (*stackA)[len((*stackA))-1] {
		rotateStack(stackA, "a")
	} else {
		numberOfRA := 0
		for !sort.IntsAreSorted(*stackA) {
			if len(*stackA) >= 2 && (*stackA)[0] >= (*stackA)[1] {
				swapStack(stackA, "a")
				if len(*stackA) >= 3 && (*stackA)[1] >= (*stackA)[2] {
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
