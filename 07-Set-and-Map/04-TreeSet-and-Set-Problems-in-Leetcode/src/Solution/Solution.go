package main

import (
	"bytes"
	"fmt"
)

// Go 中没有 set 类型，这里使用 map 实现
func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	buffer := bytes.Buffer{}
	uniqueWord := make(map[string]bool)
	for _, word := range words {
		buffer.Reset()
		for _, letter := range word {
			buffer.WriteString(morseCodes[letter-'a'])
		}
		uniqueWord[buffer.String()] = true
	}

	return len(uniqueWord)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}

	fmt.Println(uniqueMorseRepresentations(words))
}
