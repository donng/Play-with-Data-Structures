package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/donng/Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/BSTSet"
	"github.com/donng/Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/trie"
	"github.com/donng/Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/trie2"
	"github.com/donng/Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/trie3"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func main() {
	fmt.Println("Pride and Prejudice")

	filename1, _ := filepath.Abs("pride-and-prejudice.txt")
	filename2, _ := filepath.Abs("a-tale-of-two-cities.txt")
	words1 := utils.ReadFile(filename1)
	words2 := utils.ReadFile(filename2)
	words := append(words1, words2...)

	for i := 0; i < 200; i++ {
		fmt.Println(words[i])
	}

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

	t1 := trie.New()
	for _, word := range words {
		t1.Add(word)
	}
	for _, word := range words {
		t1.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", t1.GetSize())
	fmt.Println("TreeMap Trie:", diffTime)

	// ---

	startTime = time.Now()

	t2 := trie2.New()
	for _, word := range words {
		t2.Add(word)
	}
	for _, word := range words {
		t2.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", t2.GetSize())
	fmt.Println("HashMap Trie:", diffTime)

	// ---

	startTime = time.Now()

	t3 := trie3.New()
	for _, word := range words {
		t3.Add(word)
	}
	for _, word := range words {
		t3.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", t3.GetSize())
	fmt.Println("Array(Map) Trie:", diffTime)
}
