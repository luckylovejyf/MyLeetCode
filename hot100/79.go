package hot100

// 79. 单词搜索 深度优先
// 遍历每一个节点，检查从该节点出发，能不能完成word
// 递归入参：当前节点坐标，word索引
// 结束条件：当前字符不等于当前word字符，返回false；当前word索引==word长度-1，返回true（先后顺序注意）
// 每层遍历：记录当前坐标已访问，四个方向未访问遍历，如果任意一个方向返回true，则返回true，否则清理访问记录，返回false

var dicts = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var existBFS func(i int, j int, k int) bool
	existBFS = func(i int, j int, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}

		visited[i][j] = true
		defer func() { visited[i][j] = false }()

		for _, dict := range dicts {
			if i+dict[0] >= 0 && i+dict[0] < m && j+dict[1] >= 0 && j+dict[1] < n && !visited[i+dict[0]][j+dict[1]] {
				if existBFS(i+dict[0], j+dict[1], k+1) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if existBFS(i, j, 0) {
				return true
			}
		}
	}
	return false

}
