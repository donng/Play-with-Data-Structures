package main

import (
	"fmt"
	"path/filepath"

	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/02-LinkedListSet/linkedlistset"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func main() {
	fmt.Println("pride-and-prejudice")

	filename, _ := filepath.Abs("pride-and-prejudice.txt")
	words1 := utils.ReadFile(filename)
	fmt.Println(len(words1))
	set1 := linkedlistset.New()
	for _, word := range words1 {
		set1.Add(word)
	}
	fmt.Println(set1.GetSize())

	fmt.Println("a-tale-of-two-cities")

	filename, _ = filepath.Abs("a-tale-of-two-cities.txt")
	words2 := utils.ReadFile(filename)
	fmt.Println(len(words2))
	set2 := linkedlistset.New()
	for _, word := range words2 {
		set2.Add(word)
	}
	fmt.Println(set2.GetSize())
}
