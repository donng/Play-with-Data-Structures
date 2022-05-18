package solution

// IsValid go 的字符串实际是 byte 类型组成的切片
func IsValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			// 括号左半部直接入栈
			stack = append(stack, char)
		} else if len(stack) > 0 && brackets[char] == stack[len(stack)-1] {
			// 括号右半部：栈中有数据，且此元素与栈尾元素相同，出栈
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	// 循环结束，栈中还有数据返回 false
	return len(stack) == 0
}
