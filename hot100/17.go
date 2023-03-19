package hot100

// 17. 电话号码的字母组合 回溯
// 记录当前为第几层遍历，以及当前遍历的结果
// 如果遍历到最后一层，塞进结果集，返回
// 否则，遍历下一层所有的可能，层数+1，当前结果处理
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{}
	}

	var ans []string
	var backtrack func(index int, combination string)
	backtrack = func(index int, combination string) {
		if index == n {
			ans = append(ans, combination)
			return
		}

		digit := string(digits[index])
		chars := phoneMap[digit]
		for i := 0; i < len(chars); i++ {
			backtrack(index+1, combination+string(chars[i]))
		}
	}

	backtrack(0, "")
	return ans
}
