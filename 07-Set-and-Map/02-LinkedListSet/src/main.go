package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/01-Set-Basics-and-BSTSet/src/FileOperation"
	"path/filepath"
	"os"
	"fmt"
	"Play-with-Data-Structures/07-Set-and-Map/02-LinkedListSet/src/LinkedListSet"
)

func main() {
	// 获得当前目录绝对路径
	projectPath, _ := os.Getwd()
	currentPath := filepath.Join(projectPath, "07-Set-and-Map", "01-Set-Basics-and-BSTSet")

	fmt.Println("pride-and-prejudice")
	words1 := FileOperation.ReadFile(filepath.Join(currentPath, "pride-and-prejudice.txt"))
	fmt.Println(len(words1))
	set1 := LinkedListSet.GetLinkedListSet()
	for _, word := range words1 {
		set1.Add(word)
	}
	fmt.Println(set1.GetSize())

	fmt.Println("a-tale-of-two-cities")
	words2 := FileOperation.ReadFile(filepath.Join(currentPath, "a-tale-of-two-cities.txt"))
	fmt.Println(len(words2))
	set2 := LinkedListSet.GetLinkedListSet()
	for _, word := range words2 {
		set2.Add(word)
	}
	fmt.Println(set2.GetSize())
}
