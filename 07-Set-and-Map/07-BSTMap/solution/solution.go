package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/07-BSTMap/BSTMap"
)

/// Leetcode 350. Intersection of Two Arrays II
/// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的LinkedListMap类
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	customMap := BSTMap.New()

	for _, num := range nums1 {
		if customMap.Contains(num) {
			customMap.Set(num, customMap.Get(num).(int)+1)
		} else {
			customMap.Add(num, 1)
		}
	}

	for _, num := range nums2 {
		value := customMap.Get(num)

		if value != nil && value.(int) > 0 {
			res = append(res, num)
			customMap.Set(num, value.(int)-1)
		}
	}

	return res
}

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	fmt.Println(intersect(num1, num2))
}
