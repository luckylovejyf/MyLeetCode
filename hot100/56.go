package hot100

import "sort"

// 56. 合并区间 贪心
// 先按照左边界进行从小到大排序
// 比较当前区间与前一个区间是否有重叠，有的话，合并，否则当前区间成为前一个区间，并记录到结果集中
func merge(intervals [][]int) [][]int {
	var ans [][]int
	n := len(intervals)
	if n == 0 {
		return ans
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	pre := intervals[0]
	ans = append(ans, pre)
	for i := 1; i < n; i++ {
		cur := intervals[i]
		if pre[1] >= cur[0] {
			pre[1] = max(pre[1], cur[1])
		} else {
			ans = append(ans, cur)
			pre = cur
		}
	}
	return ans
}
