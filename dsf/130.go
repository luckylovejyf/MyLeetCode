package dsf

// 130. 被围绕的区域
// 从边缘节点开始，深度搜索
// 越界返回
// 不需要替换，返回
// 替换为其他标志
// 深度搜索四个方向
// 遍历替换特殊标识

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	n, m := len(board), len(board[0])
	var dsf func(x int, y int)
	dsf = func(x int, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'

		dsf(x-1, y)
		dsf(x+1, y)
		dsf(x, y-1)
		dsf(x, y+1)
	}

	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			dsf(i, 0)
		}

		if board[i][m-1] == 'O' {
			dsf(i, m-1)
		}
	}

	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			dsf(0, i)
		}

		if board[n-1][i] == 'O' {
			dsf(n-1, i)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
