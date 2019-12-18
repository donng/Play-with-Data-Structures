package trie3

type node struct {
	isWord bool
	next   [26]*node
}

type Trie struct {
	root *node
	size int
}

func generateNode() *node {
	return &node{next: [26]*node{}}
}

func New() *Trie {
	return &Trie{root: generateNode()}
}

// 获得Trie中存储的单词数量
func (t *Trie) GetSize() int {
	return t.size
}

// 向Trie中添加一个新的单词word
func (t *Trie) Add(word string) {
	cur := t.root

	for _, w := range word {
		offset := w - 'a'
		if offset > 26 {
			offset -= 26
		}

		if cur.next[offset] == nil {
			cur.next[offset] = generateNode()
		}
		cur = cur.next[offset]
	}

	if cur.isWord == false {
		cur.isWord = true
		t.size++
	}
}

// 查询单词word是否在Trie中
func (t *Trie) Contains(word string) bool {
	cur := t.root

	for _, w := range []rune(word) {
		offset := w - 'a'
		if cur.next[offset] == nil {
			return false
		}
		cur = cur.next[offset]
	}

	return cur.isWord
}
