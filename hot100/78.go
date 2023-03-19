package hot100

// 78. 子集 回溯
// 递归入参：当前的index
// 终止条件：index == n
// 每层遍历：加入当前数字、不加入当前数字
func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	var ans [][]int
	var path []int
	var subsetsBacktrace func(index int)
	subsetsBacktrace = func(index int) {
		if index == n {
			ans = append(ans, append([]int{}, path...))
			return
		}

		path = append(path, nums[index])
		subsetsBacktrace(index + 1)
		path = path[:len(path)-1]
		subsetsBacktrace(index + 1)
	}

	subsetsBacktrace(0)
	return ans
}
