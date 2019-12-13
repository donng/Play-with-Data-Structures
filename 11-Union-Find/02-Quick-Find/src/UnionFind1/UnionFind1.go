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

func (u *UnionFind) GetSize() int {
	return len(u.id)
}

// 查找元素p所对应的集合编号
// O(1)复杂度
func (u *UnionFind) find(p int) int {
	if p < 0 || p >= len(u.id) {
		panic("element out of range")
	}

	return u.id[p]
}

// 查看元素p和元素q是否所属一个集合
// O(1)复杂度
func (u *UnionFind) IsConnected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

// 合并元素p和元素q所属的集合
// O(n) 复杂度
func (u *UnionFind) UnionElements(p int, q int) {
	pID := u.find(p)
	qID := u.find(q)

	if pID == qID {
		return
	}

	// 合并过程需要遍历一遍所有元素, 将两个元素的所属集合编号合并
	for i := 0; i < len(u.id); i++ {
		if u.id[i] == pID {
			u.id[i] = qID
		}
	}
}
