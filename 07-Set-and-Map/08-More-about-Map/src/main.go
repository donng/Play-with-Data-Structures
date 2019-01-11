package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/08-More-about-Map/src/BSTMap"
	"Play-with-Data-Structures/07-Set-and-Map/08-More-about-Map/src/LinkedListMap"
	"Play-with-Data-Structures/07-Set-and-Map/08-More-about-Map/src/Map"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
	"time"
)

func testSet(p Map.Map, filename string) time.Duration {
	startTime := time.Now()

	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words:", len(words))

	for _, word := range words {
		if p.Contains(word) {
			p.Set(word, p.Get(word).(int)+1)
		} else {
			p.Add(word, 1)
		}
	}
	fmt.Println("Total different words: ", p.GetSize())
	fmt.Println("Frequency of PRIDE:", p.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", p.Get("prejudice"))

	return time.Now().Sub(startTime)
}

func main() {
	filename, _ := filepath.Abs("07-Set-and-Map/08-More-about-Map/pride-and-prejudice.txt")

	bstMap := BSTMap.Constructor()
	time1 := testSet(bstMap, filename)
	fmt.Println("BST map :", time1)

	linkedListMap := LinkedListMap.Constructor()
	time2 := testSet(linkedListMap, filename)
	fmt.Println("linkedListMap set:", time2)
}
