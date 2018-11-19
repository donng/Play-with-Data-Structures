package main

import (
	"Play-with-Data-Structures/09-Segment-Tree/06-Update-Single-Element-in-Segment-Tree/src/SegmentTree"
	"fmt"
)

/// Leetcode 307. Range Sum Query - Mutable
/// https://leetcode.com/problems/range-sum-query-mutable/description/
///
/// 使用sum数组的思路：TLE
type NumArray struct {
	segmentTree *SegmentTree.SegmentTree
}

func Constructor(nums []int) NumArray {
	data := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}

	segTree := SegmentTree.GetSegmentTree(data, func(i interface{}, j interface{}) interface{} {
		return i.(int) + j.(int)
	})

	return NumArray{segTree}
}

func (this *NumArray) Update(i int, val int) {
	this.segmentTree.Set(i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.segmentTree.Query(i, j).(int)
}

func main() {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)

	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))
}
