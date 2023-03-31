package main

// Given an array, rotate the array to the right by k steps, where k is non-negative.
// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = k % len(nums) // k can be larger than len(nums)
	if k == 0 {
		return
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
