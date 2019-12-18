package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/09-Segment-Tree/03-Building-Segment-Tree/segmenttree"
)

func main() {
	nums := []interface{}{-2, 0, 3, -5, 2, -1}

	setTree := segmenttree.New(nums, func(a interface{}, b interface{}) interface{} {
		return a.(int) + b.(int)
	})
	fmt.Println(setTree)
}
