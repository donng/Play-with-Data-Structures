package main

import (
	"Play-with-Data-Structures/09-Segment-Tree/05-Segment-Tree-Problems-in-Leetcode/src/SegmentTree"
	"fmt"
)

type NumArray struct {
	segmentTree *SegmentTree.SegmentTree
}

func Constructor(nums []int) NumArray {
	data := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}

	setTree := SegmentTree.Constructor(data, func(i interface{}, j interface{}) interface{} {
		return i.(int) + j.(int)
	})

	return NumArray{setTree}
}

func (this *NumArray) SumRange(i int, j int) int {
	if this.segmentTree == nil {
		panic("Segment Tree is null")
	}
	return this.segmentTree.Query(i, j).(int)
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)

	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
