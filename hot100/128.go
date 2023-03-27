package hot100

// 128. 最长连续序列
// 遍历每一个元素，连续向下寻找是否存在下一个数字，每次结束后更新最大长度
// 在寻找是否存在下一个数字时，可以通过hash表减少时间复杂度
// 为了避免最长序列的子序列被遍历，可以在遍历之前判断当前元素的前一个连续元素是否存在，存在则说明当前元素为起点的连续序列为子序列，不计算
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}

	ans := 1

	for _, num := range nums {
		if _, ok := m[num-1]; ok {
			continue
		}
		curNum := num
		curLen := 0
		has := true
		for has {
			curLen++
			curNum++
			_, has = m[curNum]
		}

		if curLen > ans {
			ans = curLen
		}
	}
	return ans
}
