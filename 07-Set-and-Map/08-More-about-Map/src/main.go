package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/08-More-about-Map/src/Map"
	"time"
	"fmt"
	"os"
	"path/filepath"
	"Play-with-Data-Structures/07-Set-and-Map/08-More-about-Map/src/FileOperation"
	"Play-with-Data-Structures/07-Set-and-Map/08-More-about-Map/src/BSTMap"
	"Play-with-Data-Structures/07-Set-and-Map/08-More-about-Map/src/LinkedList"
)

func testSet(p Map.Map, filename string) time.Duration {
	startTime := time.Now()

	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words:", len(words))

	for _, word := range words {
		if p.Contains(word) {
			p.Set(word, p.Get(word).(int) + 1)
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
	projectPath, _ := os.Getwd()
	currentPath := filepath.Join(projectPath, "07-Set-and-Map", "08-More-about-Map")

	filename := filepath.Join(currentPath, "pride-and-prejudice.txt")

	bstMap := BSTMap.GetBSTMap()
	time1 := testSet(bstMap, filename)
	fmt.Println("BST map :", time1)


	linkedListMap := LinkedListMap.GetLinkedListMap()
	time2 := testSet(linkedListMap, filename)
	fmt.Println("linkedListMap set:", time2)
}