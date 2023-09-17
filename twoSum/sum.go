package main

import "fmt"

func main() {
	var target = 9
	arr := []int{6, 5, 7, 2}
	result := twoSum(arr, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
    	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				res[0] = i
				res[1] = j
			}
		}
	}
	return res
}