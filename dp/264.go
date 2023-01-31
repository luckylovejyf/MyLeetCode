package dp

// 丑数
// 第一个为1
// 后面三个指针，依次乘上2、3、5
// 选取最小的，同时对应指针++

func NthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 1

	p1, p2, p3 := 0, 0, 0

	for i := 1; i < n; i++ {
		dp[i] = min(dp[p1]*2, min(dp[p2]*3, dp[p3]*5))
		if dp[i] == dp[p1]*2 {
			p1++
		}
		if dp[i] == dp[p2]*3 {
			p2++
		}
		if dp[i] == dp[p3]*5 {
			p3++
		}
	}
	return dp[n-1]
}
