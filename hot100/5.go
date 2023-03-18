package hot100

// LongestPalindrome 5. 最长回文子串
// 动态规划
// 每一个子串是回文子串，必然是在开头结尾字符相同的情况下，去头去尾也是回文子串，dp问题
// dp[i][j]表示是否为回文子串
// 一个字符为回文子串
// 外层遍历长度，内层遍历起点位置
// 计算右索引位置，如果越界，结束
// 左右节点不相等，直接false
// 左右节点相等，长度为1、2，为true
// 左右节点相等，长度>=3，看子字符串是否为回文
// 每结束一次，判断是否超过之前最大长度
func LongestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	maxLen := 1
	maxStart := 0

	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}

	for L := 2; L <= length; L++ {
		for i := 0; i < length; i++ {
			j := i + L - 1
			if j >= length {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				maxStart = i
			}
		}
	}

	return s[maxStart : maxStart+maxLen]
}
