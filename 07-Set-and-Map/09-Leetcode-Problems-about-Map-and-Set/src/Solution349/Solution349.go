package main

import "fmt"

/// Leetcode 349. Intersection of Two Arrays
/// https://leetcode.com/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	uniqueNums1 := make(map[int]bool)

	for _, num := range nums1 {
		uniqueNums1[num] = true
	}

	for _, num := range nums2 {
		if uniqueNums1[num] == true {
			res = append(res, num)
			uniqueNums1[num] = false
		}
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}
