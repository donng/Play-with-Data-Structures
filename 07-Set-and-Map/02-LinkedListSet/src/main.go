package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/02-LinkedListSet/src/LinkedListSet"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("pride-and-prejudice")

	filename, _ := filepath.Abs("07-Set-and-Map/02-LinkedListSet/pride-and-prejudice.txt")
	words1 := FileOperation.ReadFile(filename)
	fmt.Println(len(words1))
	set1 := LinkedListSet.Constructor()
	for _, word := range words1 {
		set1.Add(word)
	}
	fmt.Println(set1.GetSize())

	fmt.Println("a-tale-of-two-cities")

	filename, _ = filepath.Abs("07-Set-and-Map/02-LinkedListSet/a-tale-of-two-cities.txt")
	words2 := FileOperation.ReadFile(filename)
	fmt.Println(len(words2))
	set2 := LinkedListSet.Constructor()
	for _, word := range words2 {
		set2.Add(word)
	}
	fmt.Println(set2.GetSize())
}
