package hot100

// 22. 括号生成 回溯法
// 回溯关注两个点：1. 什么时候收集结果返回；2. 每一层做什么处理，以及遍历下一层
// 长度为2*n时收集结果返回
// 每一层可能时左括号，可能是右括号；左括号需要当前层左括号<n，右括号需要当前层的左括号>右括号
func generateParenthesis(n int) []string {
	m := 2 * n
	var ans []string
	path := make([]byte, m)

	var generateParenthesisBacktrack func(index int, open int)
	generateParenthesisBacktrack = func(index int, open int) {
		if index == m {
			ans = append(ans, string(path))
			return
		}

		if open < n {
			path[index] = '('
			generateParenthesisBacktrack(index+1, open+1)
		}
		if index-open < open {
			path[index] = ')'
			generateParenthesisBacktrack(index+1, open)
		}
	}
	generateParenthesisBacktrack(0, 0)
	return ans
}
