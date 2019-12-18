package main

import (
	"fmt"
	"path/filepath"

	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/01-Set-Basics-and-BSTSet/BSTSet"
	"github.com/donng/Play-with-Data-Structures/utils"
	"github.com/donng/Play-with-Data-Structures/utils/FileOperation"
)

func main() {
	fmt.Println("pride-and-prejudice")

	filename, _ := filepath.Abs("pride-and-prejudice.txt")
	fmt.Println(filename)
	words1 := utils.ReadFile(filename)
	fmt.Println(len(words1))

	set1 := BSTSet.New()
	for _, word := range words1 {
		set1.Add(word)
	}
	fmt.Println(set1.GetSize())

	fmt.Println("a Tale of Two Cities")

	filename, _ = filepath.Abs("a-tale-of-two-cities.txt")
	words2 := FileOperation.ReadFile(filename)
	fmt.Println(len(words2))
	set2 := BSTSet.New()
	for _, word := range words2 {
		set2.Add(word)
	}
	fmt.Println(set2.GetSize())
}
