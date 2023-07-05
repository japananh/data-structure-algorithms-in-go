package main

import "fmt"

func backtrack(nums []int, used []bool, current []int, result *[][]int) {
	// Return if the goal reached
	if len(current) == len(nums) {
		// Found a permutation, add it to the result
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// Iterate through the number of choices
	for i := 0; i < len(nums); i++ {
		// Run iif choice[i] is valid
		if !used[i] {
			// Make choices[i]
			used[i] = true
			current = append(current, nums[i])

			// Call backtrack
			backtrack(nums, used, current, result)

			// Undo choices
			// Backtrack by removing the last added number
			used[i] = false
			current = current[:len(current)-1]
		}
	}
}

func main() {
	nums := []int{1, 2, 3}

	permutations := [][]int{}
	used := make([]bool, len(nums))
	current := []int{}

	backtrack(nums, used, current, &permutations)

	fmt.Println("Permutations:")
	for _, perm := range permutations {
		fmt.Println(perm)
	}
}
