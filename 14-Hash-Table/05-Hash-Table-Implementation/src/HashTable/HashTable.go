package HashTable

import (
	"Play-with-Data-Structures/14-Hash-Table/05-Hash-Table-Implementation/src/RBTree"
	"fmt"
	"hash/fnv"
	"strconv"
)

type HashTable struct {
	hashtable []*RBTree.RBTree
	M         int
	size      int
}

func Constructor(M int) *HashTable {
	var hashtable []*RBTree.RBTree
	for i := 0; i < M; i++ {
		hashtable = append(hashtable, RBTree.Constructor())
	}

	return &HashTable{hashtable, M, 0}
}

func (h *HashTable) GetSize() int {
	return h.size
}

func (h *HashTable) Add(key interface{}, value interface{}) {
	m := h.hashtable[h.hash(key)]
	if m.Contains(key) {
		m.Add(key, value)
	} else {
		m.Add(key, value)
		h.size++
	}
}

func (h *HashTable) Remove(key interface{}) interface{} {
	m := h.hashtable[h.hash(key)]
	if m.Contains(key) {
		ret := m.Remove(key)
		h.size--
		return ret
	} else {
		return nil
	}
}

func (h *HashTable) Set(key interface{}, value interface{}) {
	m := h.hashtable[h.hash(key)]
	if !m.Contains(key) {
		panic(fmt.Sprintf("%s doesn't exist!", key))
	}
	m.Set(key, value)
}

func (h *HashTable) Contains(key interface{}) bool {
	return h.hashtable[h.hash(key)].Contains(key)
}

func (h *HashTable) Get(key interface{}) interface{} {
	return h.hashtable[h.hash(key)].Get(key)
}

func (h *HashTable) hash(key interface{}) int {
	return (int(HashCode(key)) & 0x7fffffff) % h.M
}

func HashCode(source interface{}) uint64 {
	switch source.(type) {
	case int:
		source = strconv.Itoa(source.(int))
	case float64:
		source = strconv.FormatFloat(source.(float64), 'f', 6, 64)
	}

	h := fnv.New64a()
	h.Write([]byte(source.(string)))
	return h.Sum64()
}
