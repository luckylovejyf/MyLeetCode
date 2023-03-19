package hot100

// 49. 字母异位词分组
// 异位词的字符出现次数相同，借助golang数组可以比较的特性，将异位词通过相同的字符出现数组串联起来
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string, len(strs))
	var ans [][]string
	for _, str := range strs {
		cur := [26]int{}
		for i := 0; i < len(str); i++ {
			cur[str[i]-'a']++
		}
		m[cur] = append(m[cur], str)
	}
	for _, strings := range m {
		ans = append(ans, strings)
	}
	return ans
}
