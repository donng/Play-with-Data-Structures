package main

import (
	"fmt"
	"path/filepath"

	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/06-LinkedListMap/linkedlistmap"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func main() {
	fmt.Println("pride-and-prejudice")

	filename, _ := filepath.Abs("pride-and-prejudice.txt")
	words := utils.ReadFile(filename)

	fmt.Println("Total words: ", len(words))

	customMap := linkedlistmap.New()
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
