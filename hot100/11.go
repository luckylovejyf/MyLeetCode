package hot100

// 11. 盛最多水的容器 双指针
// 双指针从两侧向中间移动，直到指针相遇
// 每次计算当前最大容量，然后移动短的指针
func maxArea(height []int) int {
	l, r := 0, len(height)-1

	ans := 0
	for l < r {
		cur := min(height[l], height[r]) * (r - l)
		ans = max(cur, ans)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
