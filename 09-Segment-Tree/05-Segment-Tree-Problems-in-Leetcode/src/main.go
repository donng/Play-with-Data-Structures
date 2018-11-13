package main

import (
	"Play-with-Data-Structures/09-Segment-Tree/04-Query-in-Segment-Tree/src/SegmentTree"
	"fmt"
)

/// 该测试用例来源：Leetcode 303. Range Sum Query - Immutable
/// https://leetcode.com/problems/range-sum-query-immutable/description/
func main() {
	nums := []interface{}{-2, 0, 3, -5, 2, -1}

	setTree := SegmentTree.GetSegmentTree(nums, func(a interface{}, b interface{}) interface{} {
		return a.(int) + b.(int)
	})
	fmt.Println(setTree)

	fmt.Println(setTree.Query(0, 2))
	fmt.Println(setTree.Query(2, 5))
	fmt.Println(setTree.Query(0, 5))
}
