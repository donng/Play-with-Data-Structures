package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/donng/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/AVLTree"
	"github.com/donng/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/BSTMap"
	"github.com/donng/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/Map"
	"github.com/donng/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/linkedlistmap"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func testSet(p Map.Map, filename string) time.Duration {
	startTime := time.Now()

	words := utils.ReadFile(filename)
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
	filename, _ := filepath.Abs("pride-and-prejudice.txt")

	bstMap := BSTMap.New()
	time1 := testSet(bstMap, filename)
	fmt.Println("BST Map:", time1)

	linkedListMap := linkedlistmap.New()
	time2 := testSet(linkedListMap, filename)
	fmt.Println("linked List Map:", time2)

	avlMap := AVLTree.New()
	time3 := testSet(avlMap, filename)
	fmt.Println("AVL Map:", time3)
}
