package main

import (
	"fmt"
	"Play-with-Data-Structures/07-Set-and-Map/06-LinkedListMap/src/FileOperation"
	"os"
	"path/filepath"
	"Play-with-Data-Structures/07-Set-and-Map/06-LinkedListMap/src/LinkedListMap"
)

func main() {
	fmt.Println("pride-and-prejudice")

	projectPath, _ := os.Getwd()
	currentPath := filepath.Join(projectPath, "07-Set-and-Map", "06-LinkedListMap")
	filename := filepath.Join(currentPath, "pride-and-prejudice.txt")

	words := FileOperation.ReadFile(filename)

	fmt.Println("Total words: ", len(words))

	customMap := LinkedListMap.GetLinkedListMap()
	for _, word := range words {
		if customMap.Contains(word) {
			customMap.Set(word, customMap.Get(word).(int) + 1)
		} else {
			customMap.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", customMap.GetSize())
	fmt.Println("Frequency of PRIDE: ", customMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", customMap.Get("prejudice"))
}
