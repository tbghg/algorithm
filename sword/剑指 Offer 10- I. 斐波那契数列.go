package sword

func Fib(n int64) int64 {
	if n < 2 {
		return n
	}
	dp := [2]int64{0, 1}
	for i := 2; i <= int(n); i++ {
		dp[0], dp[1] = dp[1], dp[0]+dp[1]
	}
	return dp[1] % 1000000007
}

func Fib2(n int) int {
	if n < 2 {
		return n
	}
	dp := [2]int{0, 1}
	for i := 2; i < n+1; i++ {
		sum := dp[0] + dp[1]
		dp[0], dp[1] = dp[1], sum
	}
	return dp[1]
}
