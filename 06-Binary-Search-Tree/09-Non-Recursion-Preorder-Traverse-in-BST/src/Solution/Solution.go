package main

import (
	"Play-with-Data-Structures/06-Binary-Search-Tree/09-Non-Recursion-Preorder-Traverse-in-BST/src/ArrayStack"
	"fmt"
)

/// Leetcode 144. Binary Tree Preorder Traversal
/// https://leetcode.com/problems/binary-tree-preorder-traversal/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的BST类
/// 该测试主要测试前序遍历的非递归写法

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return preorderNoRec(root)
}

func preorderNoRec(root *TreeNode) []int {
	var nodes []int

	if root == nil {
		return nodes
	}

	stack := ArrayStack.GetArrayStack(20)
	stack.Push(root)
	for !stack.IsEmpty() {
		cur := stack.Pop().(*TreeNode)
		nodes = append(nodes, cur.Val)

		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}

	return nodes
}

func main() {
	root := &TreeNode{Val: 1}

	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println(preorderTraversal(root))
}
