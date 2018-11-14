package main

import "fmt"

/// Leetcode 211. Add and Search Word - Data structure design
/// https://leetcode.com/problems/add-and-search-word-data-structure-design/description/
type node struct {
	isWord bool
	next   map[string]*node
}

type WordDictionary struct {
	root *node
}

func getNode() *node {
	return &node{next: make(map[string]*node)}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{getNode()}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	cur := this.root

	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = getNode()
		}
		cur = cur.next[c]
	}

	cur.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.match(this.root, word, 0)
}

func (this *WordDictionary) match(n *node, word string, index int) bool {
	if index == len(word) {
		return n.isWord
	}

	c := string([]rune(word)[index])

	if c != "." {
		if n.next[c] == nil {
			return false
		}
		return this.match(n.next[c], word, index+1)
	} else {
		for w := range n.next {
			if this.match(n.next[w], word, index+1) {
				return true
			}
		}
		return false
	}
}

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")

	fmt.Println(obj.Search("pad"))
	fmt.Println(obj.Search("bad"))
	fmt.Println(obj.Search(".ad"))
	fmt.Println(obj.Search("b.."))
}
