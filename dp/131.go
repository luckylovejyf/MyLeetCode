package dp

// 错误做法，状态转移方程错误
func Partition(s string) [][]string {
	res := make([][]string, 0)
	if len(s) == 0 {
		return res
	}
	res = append(res, []string{string(s[0])})

	for i := 1; i < len(s); i++ {
		cur := make([][]string, 0)
		for _, r := range res {
			cur = append(cur, append(r, string(s[i])))
		}

		l := len(res)
		for j := 0; j < l; j++ {
			lr := len(res[j])
			str := string(s[i])
			for k := lr - 1; k >= 0; k-- {
				str = res[j][k] + str
				if isHui(str) {
					cur = append(cur, append(res[j][:k], str))
				}
			}
		}

		res = purge(cur)
	}

	return res
}

func purge(cur [][]string) [][]string {
	res := make([][]string, 0)
	m := make(map[string]struct{}, len(cur))
	for _, strs := range cur {
		key := ""
		for _, str := range strs {
			key += str
		}
		_, ok := m[key]
		if !ok {
			res = append(res, strs)
			m[key] = struct{}{}
		}
	}
	return res
}

func isHui(str string) bool {
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func partitionAns(s string) [][]string {
	n := len(s)
	var ans [][]string
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	var splits []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return ans
}
