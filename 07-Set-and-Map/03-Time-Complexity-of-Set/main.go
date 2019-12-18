package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/03-Time-Complexity-of-Set/BSTSet"
	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/03-Time-Complexity-of-Set/linkedlistset"
	"github.com/donng/Play-with-Data-Structures/07-Set-and-Map/03-Time-Complexity-of-Set/set"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func testSet(set set.Set, filename string) time.Duration {
	startTime := time.Now()

	words := utils.ReadFile(filename)
	fmt.Println("Total words:", len(words))
	for _, word := range words {
		set.Add(word)
	}
	fmt.Println("Total different words:", set.GetSize())

	return time.Now().Sub(startTime)
}

func main() {
	filename, _ := filepath.Abs("pride-and-prejudice.txt")

	bstSet := BSTSet.New()
	time1 := testSet(bstSet, filename)
	fmt.Println("BST set :", time1)

	linkedListSet := linkedlistset.New()
	time2 := testSet(linkedListSet, filename)
	fmt.Println("linkedList set:", time2)
}
