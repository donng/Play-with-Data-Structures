package unionfind4

// 我们的第四版Union-Find
type UnionFind4 struct {
	parent []int // parent[i]表示第一个元素所指向的父节点
	rank   []int // rank[i]表示以i为根的集合中元素个数
}

// 构造函数
func New(size int) *UnionFind4 {
	rank := make([]int, size)
	parent := make([]int, size)
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < len(parent); i++ {
		rank[i] = 1
		parent[i] = i
	}

	return &UnionFind4{
		rank:   rank,
		parent: parent,
	}
}

func (u4 *UnionFind4) GetSize() int {
	return len(u4.parent)
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (u4 *UnionFind4) find(p int) int {
	if p < 0 || p > len(u4.parent) {
		panic("p is out of range.")
	}

	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != u4.parent[p] {
		p = u4.parent[p]
	}

	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度, h为树的高度
func (u4 *UnionFind4) IsConnected(p int, q int) bool {
	return u4.find(p) == u4.find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func (u4 *UnionFind4) UnionElements(p int, q int) {
	pRoot := u4.find(p)
	qRoot := u4.find(q)

	if pRoot == qRoot {
		return
	}

	// 根据两个元素所在树的rank不同判断合并方向
	// 将rank低的集合合并到rank高的集合上
	if u4.rank[pRoot] < u4.rank[qRoot] {
		u4.parent[pRoot] = u4.parent[qRoot]
	} else if u4.rank[pRoot] > u4.rank[qRoot] {
		u4.parent[qRoot] = u4.parent[pRoot]
	} else { // rank[pRoot] == rank[qRoot]
		u4.parent[pRoot] = qRoot
		u4.rank[qRoot] += 1 // 此时, 我维护rank的值
	}
}
