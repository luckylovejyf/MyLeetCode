package hot100

import "sort"

// 39. 组合总和 回溯法
// 递归入参：index本次要遍历的下标，target当前剩余的target值
// 返回条件：当前target为0
// 搜索逻辑：target减去当前数值，从当前数据向后遍历，注意剪枝
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var path []int
	var combinationSumBacktrack func(index int, target int)
	combinationSumBacktrack = func(index int, target int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
		}

		for i := index; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}
			path = append(path, candidates[i])
			combinationSumBacktrack(i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}

	combinationSumBacktrack(0, target)
	return ans
}
