package checker

import (
	"reflect"
	"testing"
)

func TestPushStack(t *testing.T) {
	stackA := []int{1, 2, 3}
	stackB := []int{4, 5}

	pushStack(&stackB, &stackA)

	expectedStackA := []int{4, 1, 2, 3}
	expectedStackB := []int{5}

	if !reflect.DeepEqual(stackA, expectedStackA) {
		t.Errorf("PushA: Expected stackA %v, got %v", expectedStackA, stackA)
	}

	if !reflect.DeepEqual(stackB, expectedStackB) {
		t.Errorf("PushA: Expected stackB %v, got %v", expectedStackB, stackB)
	}
}
func TestPushStackWithEmptyStack(t *testing.T) {
	stackA := []int{1, 2, 3}
	stackB := []int{}

	pushStack(&stackB, &stackA)

	expectedStackA := []int{1, 2, 3}
	expectedStackB := []int{}

	if !reflect.DeepEqual(stackA, expectedStackA) {
		t.Errorf("PushA: Expected stackA %v, got %v", expectedStackA, stackA)
	}

	if !reflect.DeepEqual(stackB, expectedStackB) {
		t.Errorf("PushA: Expected stackB %v, got %v", expectedStackB, stackB)
	}
}

func TestRotateStack(t *testing.T) {
	stackA := []int{1, 2, 3, 4}

	rotateStack(&stackA)

	expectedStackA := []int{2, 3, 4, 1}

	if !reflect.DeepEqual(stackA, expectedStackA) {
		t.Errorf("RotateA: Expected stackA %v, got %v", expectedStackA, stackA)
	}
}
func TestRotateStackWithOneElement(t *testing.T) {
	stackA := []int{1}

	rotateStack(&stackA)

	expectedStackA := []int{1}

	if !reflect.DeepEqual(stackA, expectedStackA) {
		t.Errorf("RotateA: Expected stackA %v, got %v", expectedStackA, stackA)
	}
}
func TestReverseRotateStack(t *testing.T) {
	stackA := []int{1, 2, 3, 4}

	reverseRotateStack(&stackA)

	expectedStackA := []int{4, 1, 2, 3}

	if !reflect.DeepEqual(stackA, expectedStackA) {
		t.Errorf("ReverseRotateA: Expected stackA %v, got %v", expectedStackA, stackA)
	}
}
func TestReverseRotateStackWithOneElement(t *testing.T) {
	stackA := []int{1}

	reverseRotateStack(&stackA)

	expectedStackA := []int{1}

	if !reflect.DeepEqual(stackA, expectedStackA) {
		t.Errorf("ReverseRotateA: Expected stackA %v, got %v", expectedStackA, stackA)
	}
}
func TestIsValidInstruction(t *testing.T) {
	tests := map[string]bool{
		"sa": true, "sb": true, "ss": true,
		"pa": true, "pb": true,
		"ra": true, "rb": true, "rr": true,
		"rra": true, "rrb": true, "rrr": true,
		"invalid": false, "sss": false,
	}

	for instr, expected := range tests {
		result := isValidInstruction(instr)
		if result != expected {
			t.Errorf("isValidInstruction(%s) - Expected: %v, Got: %v", instr, expected, result)
		}
	}
}
func TestSwapStack(t *testing.T) {
	stack := []int{1, 2, 3}
	expectedStack := []int{2, 1, 3}

	swapStack(&stack)

	if !reflect.DeepEqual(stack, expectedStack) {
		t.Errorf("SwapStack - Expected stack %v, Got stack %v", expectedStack, stack)
	}
}
func TestSwapStackWithOneElement(t *testing.T) {
	stack := []int{1}
	expectedStack := []int{1}

	swapStack(&stack)

	if !reflect.DeepEqual(stack, expectedStack) {
		t.Errorf("SwapStack - Expected stack %v, Got stack %v", expectedStack, stack)
	}
}
func TestIsSorted(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{}, true},               // Empty slice is considered sorted
		{[]int{1}, true},              // Single element is considered sorted
		{[]int{1, 2, 3, 4, 5}, true},  // Sorted in ascending order
		{[]int{5, 4, 3, 2, 1}, false}, // Not sorted in ascending order
		{[]int{1, 2, 3, 2, 5}, false}, // Not sorted in ascending order
		{[]int{3, 3, 3, 3}, true},     // All equal elements are considered sorted
	}

	for _, test := range tests {
		result := isSorted(test.input)
		if result != test.expected {
			t.Errorf("isSorted(%v) - Expected: %v, Got: %v", test.input, test.expected, result)
		}
	}
}
func TestCheckInstructions(t *testing.T) {
	tests := []struct {
		stack        []int
		instructions []string
		expected     bool
	}{
		{[]int{}, []string{}, true},
		{[]int{3, 2, 1, 0}, []string{"rra", "pb", "sa", "rra", "pa"}, true},
		{[]int{0, 9, 1, 8, 2}, []string{"pb", "ra", "pb", "ra", "sa", "ra", "pa", "pa"}, true},
		{[]int{5, 4, 3, 2, 1}, []string{"sa", "sb", "ss"}, false},
		{[]int{3, 2, 1}, []string{"ra", "pa", "pb", "rra"}, false},
		{[]int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}, []string{"sa", "pb", "rrr"}, false},
	}

	for _, test := range tests {
		result := checkInstructions(test.stack, test.instructions)
		if result != test.expected {
			t.Errorf("checkInstructions(%v, %v) - Expected: %v, Got: %v", test.stack, test.instructions, test.expected, result)
		}
	}
}
