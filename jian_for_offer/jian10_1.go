package jian_for_offer

func fib(n int) int {
	if n < 2 {
		return n
	}
	pre, after := 0, 1
	for i := 2; i <= n; i++ {
		temp := after
		after = (pre + after) % 1000000007
		pre = temp
	}
	return after
}
