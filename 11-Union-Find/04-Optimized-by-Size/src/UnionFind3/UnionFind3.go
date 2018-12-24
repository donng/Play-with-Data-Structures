package UnionFind3

// 我们的第三版Union-Find
type UnionFind3 struct {
	parent []int // parent[i]表示第一个元素所指向的父节点
	sz     []int // sz[i]表示以i为根的集合中元素个数
}

// 构造函数
func Constructor(size int) *UnionFind3 {
	sz := make([]int, size)
	parent := make([]int, size)
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < len(parent); i++ {
		sz[i] = 1
		parent[i] = i
	}

	return &UnionFind3{
		sz:     sz,
		parent: parent,
	}
}

func (this *UnionFind3) GetSize() int {
	return len(this.parent)
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (this *UnionFind3) find(p int) int {
	if p < 0 || p > len(this.parent) {
		panic("p is out of range.")
	}

	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != this.parent[p] {
		p = this.parent[p]
	}

	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度, h为树的高度
func (this *UnionFind3) IsConnected(p int, q int) bool {
	return this.find(p) == this.find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func (this *UnionFind3) UnionElements(p int, q int) {
	pRoot := this.find(p)
	qRoot := this.find(q)

	if pRoot == qRoot {
		return
	}

	// 根据两个元素所在树的元素个数不同判断合并方向
	// 将元素个数少的集合合并到元素个数多的集合上
	if this.sz[pRoot] < this.sz[qRoot] {
		this.parent[pRoot] = qRoot
		this.sz[qRoot] += this.sz[pRoot]
	} else { // sz[qRoot] <= sz[pRoot]
		this.parent[qRoot] = pRoot
		this.sz[pRoot] += this.sz[qRoot]
	}
}
