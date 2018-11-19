package main

import (
	"Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/src/BSTSet"
	"Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/src/FileOperation"
	"Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/src/Trie"
	"Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/src/Trie2"
	"Play-with-Data-Structures/10-Trie/08-Trie-Using-HashMap-and-Array/src/Trie3"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Pride and Prejudice")

	projectPath, _ := os.Getwd()
	currentPath := filepath.Join(projectPath, "10-Trie", "08-Trie-Using-HashMap-and-Array")

	filename1 := filepath.Join(currentPath, "pride-and-prejudice.txt")
	filename2 := filepath.Join(currentPath, "a-tale-of-two-cities.txt")
	words1 := FileOperation.ReadFile(filename1)
	words2 := FileOperation.ReadFile(filename2)
	words := append(words1, words2...)

	for i := 0; i < 200; i++ {
		fmt.Println(words[i])
	}

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
	fmt.Println("TreeMap Trie:", diffTime)

	// ---

	startTime = time.Now()

	trie2 := Trie2.Constructor()
	for _, word := range words {
		trie2.Add(word)
	}
	for _, word := range words {
		trie2.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", trie2.GetSize())
	fmt.Println("HashMap Trie:", diffTime)

	// ---

	startTime = time.Now()

	trie3 := Trie3.Constructor()
	for _, word := range words {
		trie3.Add(word)
	}
	for _, word := range words {
		trie3.Contains(word)
	}

	diffTime = time.Now().Sub(startTime)

	fmt.Println("Total different words:", trie3.GetSize())
	fmt.Println("Array(Map) Trie:", diffTime)
}
