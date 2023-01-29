package dp

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}

	dp := make([][2]int, l)

	dp[0][1] = -prices[0]

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return max(dp[l-1][0], dp[l-1][1])
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
