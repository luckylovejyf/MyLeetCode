package dp

import (
	"strconv"
	"strings"
)

func NumDecodings(s string) int {
	// 1 长度为0，答案为0
	if len(s) == 0 {
		return 0
	}

	// 2 以0开头，没有此组合
	if strings.HasPrefix(s, "0") {
		return 0
	}

	// 3 长度为1，且不为0，只有一种可能
	if len(s) == 1 {
		return 1
	}

	// 动态规划
	dp := make([]int, len(s))
	dp[0] = 1
	n, _ := strconv.Atoi(s[:2])
	if n > 0 && n <= 26 {
		if n%10 != 0 {
			dp[1] = 2
		} else {
			dp[1] = 1
		}
	} else {
		if n%10 != 0 {
			dp[1] = 1
		} else {
			dp[1] = 0
		}
	}

	for i := 2; i < len(s); i++ {
		tmpNum := 0
		curS := string(s[i])
		if curS != "0" {
			tmpNum = 1
		}
		dp[i] = dp[i] + dp[i-1]*tmpNum

		tmpNum = 0
		curS = string(s[i-1 : i+1])
		if !strings.HasPrefix(curS, "0") {
			tmpN, _ := strconv.Atoi(curS)
			if tmpN > 0 && tmpN <= 26 {
				tmpNum = 1
			}
		}
		dp[i] = dp[i] + dp[i-2]*tmpNum
	}

	return dp[len(s)-1]
}

// NumDecodingsWith2 优化 只需要三个变量
func NumDecodingsWith2(s string) int {
	// 1 长度为0，答案为0
	if len(s) == 0 {
		return 0
	}

	// 2 以0开头，没有此组合
	if strings.HasPrefix(s, "0") {
		return 0
	}

	// 3 长度为1，且不为0，只有一种可能
	if len(s) == 1 {
		return 1
	}

	// 动态规划
	i, j, k := 0, 0, 0
	i = 1
	n, _ := strconv.Atoi(s[:2])
	if n > 0 && n <= 26 {
		if n%10 != 0 {
			j = 2
		} else {
			j = 1
		}
	} else {
		if n%10 != 0 {
			j = 1
		} else {
			j = 0
		}
	}

	for index := 2; index < len(s); index++ {
		tmpNum := 0
		curS := string(s[index])
		if curS != "0" {
			tmpNum = 1
		}
		k = k + j*tmpNum

		tmpNum = 0
		curS = string(s[index-1 : index+1])
		if !strings.HasPrefix(curS, "0") {
			tmpN, _ := strconv.Atoi(curS)
			if tmpN > 0 && tmpN <= 26 {
				tmpNum = 1
			}
		}
		k = k + i*tmpNum

		i, j, k = j, k, 0
	}

	return j
}
