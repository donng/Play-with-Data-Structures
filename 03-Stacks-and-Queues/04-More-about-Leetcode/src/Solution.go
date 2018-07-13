package main

import "fmt"

// go 的字符串实际是 byte 类型组成的切片
func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := getArrayStack(20)

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			// 入栈
			stack.push(char)
		} else if stack.getSize() > 0 && brackets[char] == stack.peek() {
			// 栈中有数据，且此元素与栈尾元素相同
			stack.pop()
		} else {
			return false
		}
	}

	// 循环结束，栈中还有数据则 false
	return stack.getSize() == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
}
