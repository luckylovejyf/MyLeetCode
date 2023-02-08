package dsf

// 211. 添加与搜索单词 - 数据结构设计
// 字典树
// AddWord 依次添加，最后一个node要标记为结尾
// Search 当前下标 == 字符长度 返回节点是否为结尾
// Search 当前字符不为. 判断是否有当前字符对应节点，并且深度优先遍历下一个字符
// Search 当前字符为. 依次判断是否有当前字符对应节点，并且深度优先遍历下一个字符

type WordDictionary struct {
	w *wordNode
}

type wordNode struct {
	IsFinal   bool
	NextNodes [26]*wordNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		w: new(wordNode),
	}
}

func (this *WordDictionary) AddWord(word string) {
	curN := this.w
	for _, c := range word {
		node := curN.NextNodes[c-'a']
		if node == nil {
			curN.NextNodes[c-'a'] = new(wordNode)
		}
		curN = curN.NextNodes[c-'a']
	}
	curN.IsFinal = true
}

func (this *WordDictionary) Search(word string) bool {
	curN := this.w
	var dfs func(index int, node *wordNode) bool
	dfs = func(index int, node *wordNode) bool {
		if index == len(word) {
			return node.IsFinal
		}

		if word[index] != '.' {
			if node.NextNodes[word[index]-'a'] != nil && dfs(index+1, node.NextNodes[word[index]-'a']) {
				return true
			}
		} else {
			for _, nextNode := range node.NextNodes {
				if nextNode != nil && dfs(index+1, nextNode) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, curN)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
