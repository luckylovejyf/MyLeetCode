package hot100

// 46. 全排列 回溯
// 递归入参：第几个数字
// 返回条件：第n+1个数字
// 单层遍历：没有添加过的话，添加并记录，递归结束后，删除添加记录
func permute(nums []int) [][]int {
	var ans [][]int
	var path []int
	onPath := make(map[int]struct{}, len(nums))
	var permuteBacktrace func(n int)
	permuteBacktrace = func(n int) {
		if n == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if _, ok := onPath[nums[i]]; !ok {
				path = append(path, nums[i])
				onPath[nums[i]] = struct{}{}
				permuteBacktrace(n + 1)
				delete(onPath, nums[i])
				path = path[:len(path)-1]
			}
		}
	}

	permuteBacktrace(0)
	return ans
}
