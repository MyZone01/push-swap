package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func generatePermutations(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}

	if len(arr) == 1 {
		return [][]int{arr}
	}

	var result [][]int
	for i, num := range arr {
		rest := make([]int, len(arr)-1)
		copy(rest, arr[:i])
		copy(rest[i:], arr[i+1:])

		subPermutations := generatePermutations(rest)
		for _, subPerm := range subPermutations {
			perm := append([]int{num}, subPerm...)
			result = append(result, perm)
		}
	}

	return result
}

func executePushSwap(input string) string {
	cmd := exec.Command("../push-swap", input)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "Error"
	}
	return string(output)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	permutations := generatePermutations(arr)

	for _, perm := range permutations {
		input := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(perm)), " "), "[]")
		output := executePushSwap(input)

		fmt.Printf("Input: %v\n", input)
		fmt.Printf("Output: %v\n", output)
		fmt.Println("===============================")
	}
}
