package bsf

// 130. 被围绕的区域
// 从边缘开始，广度优先搜索
// 越界返回
// 不等于制定字符，继续广度
// 将搜索到的修改为特殊字符，并将四周添加到队列中
// 遍历，修改特殊字符

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	n, m := len(board), len(board[0])
	var q [][]int
	var bfs func()
	bfs = func() {
		l := len(q)
		for i := 0; i < l; i++ {
			p := q[i]
			if p[0] < 0 || p[0] >= n || p[1] < 0 || p[1] >= m || board[p[0]][p[1]] != 'O' {
				continue
			}
			board[p[0]][p[1]] = 'A'
			q = append(q, []int{p[0] + 1, p[1]})
			q = append(q, []int{p[0] - 1, p[1]})
			q = append(q, []int{p[0], p[1] + 1})
			q = append(q, []int{p[0], p[1] - 1})
			l = len(q)
		}
	}

	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			q = append(q, []int{i, 0})
		}

		if board[i][m-1] == 'O' {
			q = append(q, []int{i, m - 1})
		}
	}

	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			q = append(q, []int{0, i})
		}

		if board[n-1][i] == 'O' {
			q = append(q, []int{n - 1, i})
		}
	}

	bfs()

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
