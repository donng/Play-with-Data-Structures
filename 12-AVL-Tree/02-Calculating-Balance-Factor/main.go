package main

import (
	"fmt"
	"path/filepath"

	"github.com/donng/Play-with-Data-Structures/12-AVL-Tree/02-Calculating-Balance-Factor/BSTMap"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func main() {
	filename, _ := filepath.Abs("pride-and-prejudice.txt")

	bstMap := BSTMap.New()

	words := utils.ReadFile(filename)
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
