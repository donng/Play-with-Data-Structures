package main

import (
	"Play-with-Data-Structures/08-Heap-and-Priority-Queue/09-Segment-Tree/03-Building-Segment-Tree/src/SegmentTree"
	"fmt"
)

func main()  {
	nums := []interface{}{-2, 0, 3, -5, 2, -1}

	setTree := SegmentTree.GetSegmentTree(nums, func(a interface{}, b interface{}) interface{} {
		return a.(int) + b.(int)
	})
	fmt.Println(setTree)
}