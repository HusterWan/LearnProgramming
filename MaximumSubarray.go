package main

import "fmt"

// Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

// For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
// the contiguous subarray [4,-1,2,1] has the largest sum = 6.

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	tempSum := 0
	for _, v := range nums {
		tempSum += v
		if tempSum > maxSum {
			maxSum = tempSum
		}
		if tempSum < 0 {
			tempSum = 0
		}
	}

	return maxSum
}
