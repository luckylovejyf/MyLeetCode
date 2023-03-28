package hot100

// 146. LRU 缓存
// 成员变量：size、capacity、map、list
// map存储key与对应节点的映射
// list是双向链表 越往前表示最近访问
// get：判断map中是否存在，不存在返回-1；存在，挪到开头
// put：判断map中是否存在，不存在，新增添加到表头，如果超过长度，删除最末尾；存在，修改值，挪到开头

type LRUNode struct {
	Val  int
	Key  int
	Pre  *LRUNode
	Next *LRUNode
}

type LRUCache struct {
	Size     int
	Capacity int
	M        map[int]*LRUNode
	Head     *LRUNode
	Tail     *LRUNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		Size:     0,
		Capacity: capacity,
		M:        nil,
		Head:     &LRUNode{0, 0, nil, nil},
		Tail:     &LRUNode{0, 0, nil, nil},
	}
	lruCache.Head.Next = lruCache.Tail
	lruCache.Tail.Pre = lruCache.Head
	lruCache.M = make(map[int]*LRUNode, capacity)
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.M[key]; !ok {
		return -1
	} else {
		this.MoveToHead(n)
		return n.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.M[key]; !ok {
		newN := &LRUNode{value, key, nil, nil}
		this.AddHead(newN)
		if this.Size == this.Capacity {
			this.RemoveTail()
		} else {
			this.Size++
		}
		this.M[key] = newN
	} else {
		n.Val = value
		this.MoveToHead(n)
	}
}

func (this *LRUCache) MoveToHead(n *LRUNode) {
	this.Remove(n)
	this.AddHead(n)
}

func (this *LRUCache) Remove(n *LRUNode) {
	n.Next.Pre = n.Pre
	n.Pre.Next = n.Next
}

func (this *LRUCache) AddHead(n *LRUNode) {
	n.Pre = this.Head
	n.Next = this.Head.Next
	this.Head.Next.Pre = n
	this.Head.Next = n
}

func (this *LRUCache) RemoveTail() {
	n := this.Tail.Pre
	delete(this.M, n.Key)
	this.Tail.Pre = this.Tail.Pre.Pre
	this.Tail.Pre.Next = this.Tail
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
