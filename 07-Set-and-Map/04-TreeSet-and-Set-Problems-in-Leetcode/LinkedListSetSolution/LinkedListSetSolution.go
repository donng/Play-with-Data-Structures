package main

import (
	"bytes"
	"fmt"
	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/02-LinkedListSet/linkedlistset"
)

// Go 中没有 set 类型，这里使用 map 实现
func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	buffer := bytes.Buffer{}
	uniqueWordSet := linkedlistset.New()
	for _, word := range words {
		buffer.Reset()
		for _, letter := range word {
			buffer.WriteString(morseCodes[letter-'a'])
		}
		uniqueWordSet.Add(buffer.String())
	}

	return uniqueWordSet.GetSize()
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}

	fmt.Println(uniqueMorseRepresentations(words))
}
