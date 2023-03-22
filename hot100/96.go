package hot100

// 96. 不同的二叉搜索树 动态规划
// 固定根节点，左右子树数量乘积之和，就是当前所有的可能
// 当前问题，变成左右子树的问题——动态规划
func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
