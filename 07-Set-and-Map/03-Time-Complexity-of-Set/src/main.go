package main

import (
	"Play-with-Data-Structures/07-Set-and-Map/03-Time-Complexity-of-Set/src/BSTSet"
	"Play-with-Data-Structures/07-Set-and-Map/03-Time-Complexity-of-Set/src/LinkedListSet"
	"Play-with-Data-Structures/07-Set-and-Map/03-Time-Complexity-of-Set/src/Set"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
	"time"
)

func testSet(set Set.Set, filename string) time.Duration {
	startTime := time.Now()

	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words:", len(words))
	for _, word := range words {
		set.Add(word)
	}
	fmt.Println("Total different words:", set.GetSize())

	return time.Now().Sub(startTime)
}

func main() {
	filename, _ := filepath.Abs("07-Set-and-Map/03-Time-Complexity-of-Set/pride-and-prejudice.txt")

	bstSet := BSTSet.Constructor()
	time1 := testSet(bstSet, filename)
	fmt.Println("BST set :", time1)

	linkedListSet := LinkedListSet.Constructor()
	time2 := testSet(linkedListSet, filename)
	fmt.Println("linkedList set:", time2)
}
