package main

import "fmt"

// sum[i]存储前i个元素和, sum[0] = 0
// 即sum[i]存储nums[0...i-1]的和
// sum(i, j) = sum[j + 1] - sum[i]
type NumArray struct {
	sum []int
}

func New(nums []int) NumArray {
	sum := make([]int, len(nums)+1)
	sum[0] = 0

	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	return NumArray{sum}
}

func (na *NumArray) SumRange(i int, j int) int {
	return na.sum[j+1] - na.sum[i]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := New(nums)

	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
