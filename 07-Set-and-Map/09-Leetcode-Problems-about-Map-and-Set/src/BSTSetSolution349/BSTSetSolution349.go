package main

import (
	"fmt"
	"Play-with-Data-Structures/07-Set-and-Map/02-LinkedListSet/src/LinkedListSet"
)

/// Leetcode 349. Intersection of Two Arrays
/// https://leetcode.com/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	BSTSet := LinkedListSet.GetLinkedListSet()

	for _, num := range nums1 {
		BSTSet.Add(num)
	}

	for _, num := range nums2 {
		if BSTSet.Contains(num) {
			res = append(res, num)
			BSTSet.Remove(num)
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}
