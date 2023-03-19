package hot100

// 55. 跳跃游戏 贪心
// 记录当前可以跳跃的最远距离，如果到达最后一个，返回true
// 遍历，如果遍历到超过最远距离，返回false
func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	farest := 0
	for i := 0; i < n; i++ {
		if i > farest {
			return false
		}
		farest = max(i+nums[i], farest)
		if farest >= n-1 {
			return true
		}
	}
	return false
}
