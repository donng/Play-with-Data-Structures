package main

import (
	"Play-with-Data-Structures/11-Union-Find/02-Quick-Find/src/LinkedListSet"
	"fmt"
)

/// Leetcode 547. Friend Circles
/// https://leetcode.com/problems/friend-circles/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的UF类
// 我们的第一版Union-Find
// 我们的第二版Union-Find, 使用一个数组构建一棵指向父节点的树
// parent[i]表示第一个元素所指向的父节点
type UnionFind2 struct {
	parent []int
}

// 构造函数
func Constructor(size int) *UnionFind2 {
	parent := make([]int, size)
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	return &UnionFind2{parent}
}

func (u2 *UnionFind2) GetSize() int {
	return len(u2.parent)
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (u2 *UnionFind2) find(p int) int {
	if p < 0 || p > len(u2.parent) {
		panic("p is out of range.")
	}

	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != u2.parent[p] {
		p = u2.parent[p]
	}

	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度, h为树的高度
func (u2 *UnionFind2) IsConnected(p int, q int) bool {
	return u2.find(p) == u2.find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func (u2 *UnionFind2) UnionElements(p int, q int) {
	pRoot := u2.find(p)
	qRoot := u2.find(q)

	if pRoot == qRoot {
		return
	}

	u2.parent[pRoot] = qRoot
}

func findCircleNum(M [][]int) int {
	length := len(M)
	uf := Constructor(length)

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if M[i][j] == 1 {
				uf.UnionElements(i, j)
			}
		}
	}

	linkedListSet := LinkedListSet.Constructor()
	for i := 0; i < length; i++ {
		linkedListSet.Add(uf.find(i))
	}

	return linkedListSet.GetSize()
}

func main() {
	M := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	fmt.Println(findCircleNum(M))
}
