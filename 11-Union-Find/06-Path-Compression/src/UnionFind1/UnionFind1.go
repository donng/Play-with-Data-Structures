package UnionFind1

// 我们的第一版Union-Find
type UnionFind struct {
	id []int // 我们的第一版Union-Find本质就是一个数组
}

func Constructor(size int) *UnionFind {
	id := make([]int, size)
	// 初始化, 每一个id[i]指向自己, 没有合并的元素
	for i := 0; i < len(id); i++ {
		id[i] = i
	}
	return &UnionFind{id}
}

func (this *UnionFind) GetSize() int {
	return len(this.id)
}

// 查找元素p所对应的集合编号
// O(1)复杂度
func (this *UnionFind) find(p int) int {
	if p < 0 || p >= len(this.id) {
		panic("element out of range")
	}

	return this.id[p]
}

// 查看元素p和元素q是否所属一个集合
// O(1)复杂度
func (this *UnionFind) IsConnected(p int, q int) bool {
	return this.find(p) == this.find(q)
}

// 合并元素p和元素q所属的集合
// O(n) 复杂度
func (this *UnionFind) UnionElements(p int, q int) {
	pID := this.find(p)
	qID := this.find(q)

	if pID == qID {
		return
	}

	// 合并过程需要遍历一遍所有元素, 将两个元素的所属集合编号合并
	for i := 0; i < len(this.id); i++ {
		if this.id[i] == pID {
			this.id[i] = qID
		}
	}
}
