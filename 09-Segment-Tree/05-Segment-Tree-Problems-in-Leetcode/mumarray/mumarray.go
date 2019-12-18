package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/09-Segment-Tree/05-Segment-Tree-Problems-in-Leetcode/segmenttree"
)

type NumArray struct {
	segmentTree *segmenttree.SegmentTree
}

func New(nums []int) NumArray {
	data := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}

	setTree := segmenttree.New(data, func(i interface{}, j interface{}) interface{} {
		return i.(int) + j.(int)
	})

	return NumArray{setTree}
}

func (na *NumArray) SumRange(i int, j int) int {
	if na.segmentTree == nil {
		panic("segment tree is null")
	}
	return na.segmentTree.Query(i, j).(int)
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := New(nums)

	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
