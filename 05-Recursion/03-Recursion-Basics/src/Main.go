package main

import "fmt"

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := GetListNode(nums)
	fmt.Println(head)

	res := removeElements(head, 6)
	fmt.Println(res)
}
