package jian_for_offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return make([]int, 0)
	}
	temp := head
	index := 1
	for temp.Next != nil {
		index++
		temp = temp.Next
	}
	res := make([]int, index)
	for head != nil {
		index--
		res[index] = head.Val
		head = head.Next
	}
	return res
}
