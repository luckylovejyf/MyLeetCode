package hot100

// 139. 单词拆分 动态规划
// dp[i]表示s[:i]能否被wordDict拼接成功
// dp[0]=true
// dp[i]=dp[j]&&has(s[j:i)
func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	dp := make([]bool, l+1)
	dp[0] = true

	m := make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		m[w] = struct{}{}
	}

	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			_, has := m[s[j:i]]
			if dp[j] && has {
				dp[i] = true
				break
			}
		}
	}
	return dp[l]
}
