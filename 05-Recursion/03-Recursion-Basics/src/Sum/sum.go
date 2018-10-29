package main

import "fmt"

func Sum(arr []int) int {
	return sum(arr, 0)
}

// 计算 arr[l...n) 这个区间内所有数字的和
func sum(arr []int, l int) int {
	// 求解最基本问题
	if l == len(arr) {
		return 0
	}
	// 把原问题转换成更小的问题
	return arr[l] + sum(arr, l+1)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Sum(nums))
}
