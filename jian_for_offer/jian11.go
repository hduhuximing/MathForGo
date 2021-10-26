package jian_for_offer

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if numbers[0] < numbers[len(numbers)-1] {
		return numbers[0]
	}
	start, end := 0, len(numbers)-1
	for start < end {
		mid := (end-start)/2 + start
		if numbers[mid] > numbers[end] {
			start = mid + 1
		} else if numbers[mid] < numbers[end] {
			end = mid
		} else {
			end--
		}
	}
	return numbers[start]
}
