package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/06-LinkedListMap/src/LinkedListMap"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("pride-and-prejudice")

	filename, _ := filepath.Abs("07-Set-and-Map/06-LinkedListMap/pride-and-prejudice.txt")
	words := FileOperation.ReadFile(filename)

	fmt.Println("Total words: ", len(words))

	customMap := LinkedListMap.Constructor()
	for _, word := range words {
		if customMap.Contains(word) {
			customMap.Set(word, customMap.Get(word).(int)+1)
		} else {
			customMap.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", customMap.GetSize())
	fmt.Println("Frequency of PRIDE: ", customMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", customMap.Get("prejudice"))
}
