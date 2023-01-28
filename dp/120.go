package dp

import "math"

func MinimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = triangle[i][j]
			if i > 0 {
				if j > 0 {
					if j != i {
						dp[i][j] = triangle[i][j] + min(dp[i-1][j-1], dp[i-1][j])
					} else {
						dp[i][j] = triangle[i][j] + dp[i-1][j-1]
					}
				} else {
					dp[i][j] = triangle[i][j] + dp[i-1][j]
				}
			}
		}
	}

	res := math.MaxInt
	for i := 0; i < m; i++ {
		res = min(res, dp[m-1][i])
	}

	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
