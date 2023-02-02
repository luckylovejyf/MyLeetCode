package dp

import "sort"

// 368. 最大整除子集
// 长度如果小于等于1，直接返回
// dp表示包含当前数字的最长长度，初始化为1
// 从小到大排序原数组
// 遍历前序数组，如果某个能整除，并且+1长度大于当前，则状态转换
// 找到最大长度
// 后序遍历，如果最大长度等于当前，并且当前数字可以被上一个整除，说明属于当下结果序列，添加

func LargestDivisibleSubset(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}

	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	sort.Ints(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	maxDp := 0
	for i := 0; i < l; i++ {
		maxDp = max(maxDp, dp[i])
	}

	res := make([]int, 0)
	for i := l - 1; i >= 0; i-- {
		if dp[i] == maxDp {
			if len(res) == 0 || res[len(res)-1]%nums[i] == 0 {
				res = append(res, nums[i])
				maxDp--
			}
		}
	}

	return res
}
