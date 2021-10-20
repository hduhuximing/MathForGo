package jian_for_offer

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[nums[i]]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}
func findRepeatNumber1(nums []int) int {
	mp := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if mp[nums[i]] {
			return nums[i]
		}
		mp[nums[i]] = true
	}
	return -1
}
