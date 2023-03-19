package hot100

// 48. 旋转图像
// 顺时针旋转90度，可以发现四个点刚好围成一个圈进行数值交换
// 要判断遍历多少范围
// 如果n是偶数，只需要遍历n/2
// 如果n是奇数，遍历(n-1)/2和(n+1)/2
// 综合来看，遍历n/2和(n+1)/2
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
