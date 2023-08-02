package pushswap

import (
	"fmt"
	"sort"
)

var (
	NumberOfInstruction int = 0
	Instructions        string
	_stackA             []int
	_stackB             []int
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
		for len(*stackA) != 3 && !sort.IntsAreSorted(*stackA) {
			pushStack(stackA, &stackB, "b")
		}
		if !IsDecreaseSorted(&stackB) {
			swapStack(&stackB, "b")
		}
		PushSwap3Number(stackA)
		MergeSmallStacks(stackA, &stackB)
	}
}

func BigSort(stackA *[]int) {
	stackB := []int{}
	RadixSort(stackA, &stackB)
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

// MergeSmallStacks merges the elements of stackB into stackA in sorted order using a combination of push, rotate, and reverse rotate operations.
// The function assumes that stackA and stackB are already initialized.
// The function uses a max and min value to identify the elements that need special handling when merging.
func MergeSmallStacks(stackA, stackB *[]int) {
	// Define the max and min values for special handling
	max := 5
	min := 1

	// Loop until stackB is empty
	for len(*stackB) != 0 {
		var counter int

		// Handle the max value separately
		if (*stackB)[0] == max {
			pushStack(stackB, stackA, "a")
			rotateStack(stackA, "a")
			continue
			// Handle the min value separately
		} else if (*stackB)[0] == min {
			pushStack(stackB, stackA, "a")
			continue
		}

		// If the top element of stackB is less than the top element of stackA,
		// push it to stackA and continue with the next element in stackB.
		if (*stackB)[0] < (*stackA)[0] {
			pushStack(stackB, stackA, "a")
		} else {
			// Find the correct position to insert the top element of stackB in stackA.
			// This is done by finding the right counter value that indicates the position of insertion.
			for i := 0; i < len(*stackA); i++ {
				if i == 0 {
					if (*stackA)[0] > (*stackB)[0] && (*stackA)[len(*stackA)-1] < (*stackB)[0] {
						counter = i + 1
						break
					}
				} else if i == len(*stackA)-1 {
					if (*stackA)[0] < (*stackB)[0] && (*stackA)[len(*stackA)-1] > (*stackB)[0] {
						counter = i + 1
						break
					}
				} else if (*stackA)[i] < (*stackB)[0] && (*stackA)[i+1] > (*stackB)[0] {
					counter = i + 1
					break
				} else if (*stackA)[i-1] < (*stackB)[0] && (*stackA)[i] > (*stackB)[0] {
					counter = i - 1
					break
				}
			}

			// The loop aims to find the correct position to insert the top element of stackB into stackA.
			// It uses the counter value obtained earlier to determine the position of insertion.
			// The loop iterates over the elements of stackA to find the appropriate position.
			// There are two cases to consider: when the counter value minus the loop index is zero or when it is equal to zero.
			// These two cases correspond to two possible insertion positions: at the beginning or end of stackA.
			for i := 1; i < len(*stackA); i++ {
				if (counter)-i == 0 || counter == 0 {
					// If the counter value minus the loop index is zero or the counter is zero, it means the insertion position is at the beginning of stackA.
					// To achieve this, the loop rotates the stackA elements 'i' times to bring the top element to the bottom.
					for j := 0; j < i; j++ {
						rotateStack(stackA, "a")
					}

					// Then, it pushes the top element of stackB to stackA using the pushStack function.
					pushStack(stackB, stackA, "a")

					// If stackB is not empty, it checks if the top element of stackB can be pushed to stackA again based on certain conditions.
					// Specifically, if the top element of stackB is less than the top element of stackA and greater than the last element of stackA,
					// it can be pushed again to achieve a more optimized merge.
					for len(*stackB) != 0 && (*stackB)[0] < (*stackA)[0] && (*stackB)[0] > (*stackA)[len(*stackA)-1] {
						pushStack(stackB, stackA, "a")
					}

					// Finally, it reverse rotates the stackA elements 'i' times to bring the top element back to its original position.
					for j := 0; j < i; j++ {
						reverseRotateStack(stackA, "a")
					}
					break
				} else if counter+i == len(*stackA) || counter == len(*stackA) {
					// If the counter value plus the loop index is equal to the length of stackA or the counter is equal to the length of stackA,
					// it means the insertion position is at the end of stackA.
					// To achieve this, the loop reverse rotates the stackA elements 'i' times to bring the top element to the top again.
					for j := 0; j < i; j++ {
						reverseRotateStack(stackA, "a")
					}

					// Then, it pushes the top element of stackB to stackA using the pushStack function.
					pushStack(stackB, stackA, "a")

					// If stackB is not empty, it checks if the top element of stackB can be pushed to stackA again based on certain conditions.
					// Specifically, if the top element of stackB is less than the top element of stackA and greater than the last element of stackA,
					// it can be pushed again to achieve a more optimized merge.
					if len(*stackB) != 0 {
						if (*stackB)[0] < (*stackA)[0] && (*stackB)[0] > (*stackA)[len(*stackA)-1] {
							pushStack(stackB, stackA, "a")
							i++
						}
					}

					// Finally, it rotates the stackA elements 'i+1' times to bring the top element back to its original position.
					for j := 0; j < i+1; j++ {
						rotateStack(stackA, "a")
					}
					break
				}
			}

		}
	}
}

// RadixSort returns a string of instructions.
// Its result is correct if there's no negative number.
func RadixSort(stackA, stackB *[]int) {
	l := len(*stackA)
	maxNum := len(*stackA) - 1
	maxBits := 0
	for maxNum>>maxBits != 0 {
		maxBits++
	}
	for i := 0; i < maxBits; i++ {
		for j := 0; j < l; j++ {
			num := (*stackA)[len(*stackA)-1]
			if (num>>i)&1 == 1 {
				rotateStack(stackA, "a")
			} else {
				pushStack(stackA, stackB, "b")
			}
		}
		for len(*stackB) != 0 {
			pushStack(stackB, stackA, "a")
		}
	}
}
