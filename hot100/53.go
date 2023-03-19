package hot100

// 53. 最大子数组和 动态规划
// 记录以每一个数字结尾的最大数组和
// 当前最大数组和取决于当前数字与当前数字+前一个数字结尾的最大数组和
// 状态只与前一个有关，可以滚动窗口降低空间复杂度
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := nums[0]

	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
