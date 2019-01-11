package main

import (
	"Play-with-Data-Structures/12-AVL-Tree/02-Calculating-Balance-Factor/src/BSTMap"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
)

func main() {
	filename, _ := filepath.Abs("12-AVL-Tree/02-Calculating-Balance-Factor/pride-and-prejudice.txt")

	bstMap := BSTMap.Constructor()

	words := FileOperation.ReadFile(filename)
	for _, word := range words {
		if bstMap.Contains(word) {
			bstMap.Set(word, bstMap.Get(word).(int)+1)
		} else {
			bstMap.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", bstMap.GetSize())
	fmt.Println("Frequency of PRIDE:", bstMap.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", bstMap.Get("prejudice"))
}
