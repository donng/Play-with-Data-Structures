package main

import (
	"bytes"
	"fmt"
	"github.com/donng/Play-with-Data-Structures/05-Recursion/06-Debug-Recursive-Solution/listnode"
)

func RemoveElements(head *listnode.ListNode, val int, depth int) *listnode.ListNode {
	// depthString 代表递归深度
	depthString := generateDepthString(depth)
	// 递归调用前打印
	fmt.Print(depthString)
	fmt.Println("Call: remove", val, " in ", head)

	if head == nil {
		fmt.Print(depthString)
		fmt.Println("Return: ", head)
		return nil
	}

	res := RemoveElements(head.Next, val, depth+1)
	// 递归调用后打印
	fmt.Print(depthString)
	fmt.Println("After remove ", val, ": ", res)

	ret := &listnode.ListNode{}
	if head.Val == val {
		ret = res
	} else {
		head.Next = res
		ret = head
	}
	// 节点处理后打印
	fmt.Print(depthString)
	fmt.Println("Return: ", ret)

	return ret
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := listnode.GetListNode(nums)
	fmt.Println(head)

	res := RemoveElements(head, 6, 0)
	fmt.Println(res)
}
