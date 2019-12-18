package trie

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
	size int
}

func generateNode() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

func New() *Trie {
	return &Trie{
		root: generateNode(),
	}
}

// 获得Trie中存储的单词数量
func (t *Trie) GetSize() int {
	return t.size
}

// 向Trie中添加一个新的单词word
func (t *Trie) Add(word string) {
	cur := t.root

	for _, w := range []rune(word) {
		c := string(w)

		if cur.next[c] == nil {
			cur.next[c] = generateNode()
		}
		cur = cur.next[c]
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
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}

	return cur.isWord
}

// 查询是否在Trie中有单词以prefix为前缀
func (t *Trie) IsPrefix(prefix string) bool {
	cur := t.root

	for _, s := range []rune(prefix) {
		c := string(s)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}

	return true
}
