package hot100

// LengthOfLongestSubstring 3. 无重复字符的最长子串
// 暴力解法：直接两层遍历O(N^2)
// 滑动窗口：固定左指针，右指针向右遍历，直到出现越界或者出现重复字符
// 计算当前是否为最大长度
// 清理左指针字符在map中的记录，左指针右移一位
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	r, l := 0, 0
	m := map[byte]int{}
	ans := 0

	for l = 0; l < n; l++ {
		for r < n && m[s[r]] == 0 {
			m[s[r]]++
			r++
		}
		if r-l > ans {
			ans = r - l
		}
		delete(m, s[l])
	}

	return ans
}
