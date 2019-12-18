package main

import "fmt"

/// Leetcode 350. Intersection of Two Arrays II
/// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	nums1Map := make(map[int]int)

	for _, num := range nums1 {
		nums1Map[num]++
	}

	for _, num := range nums2 {
		if nums1Map[num] > 0 {
			res = append(res, num)
			nums1Map[num]--
		}
	}

	return res
}

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	fmt.Println(intersect(num1, num2))
}
