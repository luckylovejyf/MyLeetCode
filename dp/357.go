package dp

// 357. 统计各位数字都不同的数字个数
// 边界是0和1
// 每+1，相当于在前一个的基础上添加一位，为保证不相等，可选择的数字递减
// 归纳法可以看到最后的转换方程

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 10
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 10

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + (dp[i-1]-dp[i-2])*(10-(i-1))
	}
	return dp[n]
}
