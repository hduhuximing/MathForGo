package jian_for_offer

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	pre, after := 1, 1
	for i := 2; i <= n; i++ {
		temp := after
		after = (pre + after) % 1000000007
		pre = temp
	}
	return after
}
