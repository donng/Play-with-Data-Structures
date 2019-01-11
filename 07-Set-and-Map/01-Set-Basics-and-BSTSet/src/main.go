package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/01-Set-Basics-and-BSTSet/src/BSTSet"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("pride-and-prejudice")

	filename, _ := filepath.Abs("07-Set-and-Map/01-Set-Basics-and-BSTSet/pride-and-prejudice.txt")
	words1 := FileOperation.ReadFile(filename)
	fmt.Println(len(words1))

	set1 := BSTSet.Constructor()
	for _, word := range words1 {
		set1.Add(word)
	}
	fmt.Println(set1.GetSize())

	fmt.Println("a Tale of Two Cities")

	filename, _ = filepath.Abs("07-Set-and-Map/01-Set-Basics-and-BSTSet/a-tale-of-two-cities.txt")
	words2 := FileOperation.ReadFile(filename)
	fmt.Println(len(words2))
	set2 := BSTSet.Constructor()
	for _, word := range words2 {
		set2.Add(word)
	}
	fmt.Println(set2.GetSize())
}
