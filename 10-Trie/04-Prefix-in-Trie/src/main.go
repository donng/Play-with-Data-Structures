package main

import (
	"Play-with-Data-Structures/10-Trie/04-Prefix-in-Trie/src/BSTSet"
	"Play-with-Data-Structures/10-Trie/04-Prefix-in-Trie/src/Trie"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Pride and Prejudice")

	filename, _ := filepath.Abs("10-Trie/04-Prefix-in-Trie/pride-and-prejudice.txt")
	words := FileOperation.ReadFile(filename)

	startTime := time.Now()

	bstSet := BSTSet.Constructor()
	for _, word := range words {
		bstSet.Add(word)
	}
	for _, word := range words {
		bstSet.Contains(word)
	}

	diffTime := time.Now().Sub(startTime)

	fmt.Println("Total different words:", bstSet.GetSize())
	fmt.Println("BSTSet:", diffTime)

	// ---

	startTime = time.Now()

	trie := Trie.Constructor()
	for _, word := range words {
		trie.Add(word)
	}
	for _, word := range words {
		trie.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", trie.GetSize())
	fmt.Println("BSTSet:", diffTime)
}
