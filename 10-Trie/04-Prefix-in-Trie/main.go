package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/donng/Play-with-Data-Structures/10-Trie/04-Prefix-in-Trie/BSTSet"
	"github.com/donng/Play-with-Data-Structures/10-Trie/04-Prefix-in-Trie/trie"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func main() {
	fmt.Println("Pride and Prejudice")

	filename, _ := filepath.Abs("pride-and-prejudice.txt")
	words := utils.ReadFile(filename)

	startTime := time.Now()

	bstSet := BSTSet.New()
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

	t := trie.New()
	for _, word := range words {
		t.Add(word)
	}
	for _, word := range words {
		t.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", t.GetSize())
	fmt.Println("BSTSet:", diffTime)
}
