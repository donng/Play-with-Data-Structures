package solution

import (
	"github.com/donng/Play-with-Data-Structures/04-Linked-List/06-Implement-Stack-in-LinkedList/linkedliststack"
)

/// Leetcode 20. Valid Parentheses
/// https://leetcode.com/problems/valid-parentheses/description/
/// 括号匹配问题
///
/// 使用LinkedListStack测试20号问题的代码在视频中没有提及
/// 该代码主要用于使用Leetcode上的问题测试我们的LinkedListStack：）
func IsValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := linkedliststack.New()

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			// 括号左半部直接入栈
			stack.Push(char)
		} else if stack.GetSize() > 0 && brackets[char] == stack.Peek() {
			// 括号右半部：栈中有数据，且此元素与栈尾元素相同，出栈
			stack.Pop()
		} else {
			return false
		}
	}

	// 循环结束，栈中还有数据返回 false
	return stack.GetSize() == 0
}
