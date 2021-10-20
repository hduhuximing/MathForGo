package jian_for_offer

import "sort"

func minMoves(nums []int) int {
	sort.Ints(nums)
	var count int
	for i := 0; i < len(nums); i++ {
		count += nums[i] - nums[0]
	}
	return count
}
