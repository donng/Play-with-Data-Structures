package main

import "fmt"

/// Leetcode 307. Range Sum Query - Mutable
/// https://leetcode.com/problems/range-sum-query-mutable/description/
///
/// 使用sum数组的思路：TLE
type NumArray struct {
	sum  []int
	data []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	sum[0] = 0

	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	return NumArray{
		sum:  sum,
		data: nums,
	}
}

func (this *NumArray) Update(i int, val int) {
	if i < 0 || i > len(this.data)-1 {
		panic("Index out of range.")
	}

	this.data[i] = val
	for i := i + 1; i <= len(this.data); i++ {
		this.sum[i] = this.data[i-1] + this.sum[i-1]
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}

func main() {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)

	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))
}
