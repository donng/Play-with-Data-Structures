package main

import (
	"bytes"
	"fmt"
	"github.com/donng/Play-with-Data-Structures/06-Binary-Search-Tree/04-Improved-Add-Elements-in-BST/BST"
)

/// Leetcode 804. Unique Morse Code Words
/// https://leetcode.com/problems/unique-morse-code-words/description/
///
/// 课程中在这里暂时没有介绍这个问题
/// 该代码主要用于使用Leetcode上的问题测试我们的BST类
func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	buffer := bytes.Buffer{}
	uniqueTree := BST.New()

	for _, word := range words {
		buffer.Reset()
		for _, letter := range word {
			buffer.WriteString(morseCode[letter-'a'])
		}
		uniqueTree.Add(buffer.String())
	}

	return uniqueTree.GetSize()
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}

	fmt.Println(uniqueMorseRepresentations(words))
}
