package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/01-Set-Basics-and-BSTSet/BSTSet"
)

/// Leetcode 349. Intersection of Two Arrays
/// https://leetcode.com/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	bstSet := BSTSet.New()

	for _, num := range nums1 {
		bstSet.Add(num)
	}

	for _, num := range nums2 {
		if bstSet.Contains(num) {
			res = append(res, num)
			bstSet.Remove(num)
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}
