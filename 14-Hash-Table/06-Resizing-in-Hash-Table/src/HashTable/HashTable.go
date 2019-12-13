package HashTable

import (
	"Play-with-Data-Structures/14-Hash-Table/06-Resizing-in-Hash-Table/src/RBTree"
	"fmt"
	"hash/fnv"
	"strconv"
)

const upperTol = 10
const lowerTol = 2
const initCapacity = 7

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

		if h.size >= upperTol*h.M {
			h.resize(2 * h.M)
		}
	}
}

func (h *HashTable) Remove(key interface{}) interface{} {
	m := h.hashtable[h.hash(key)]
	if m.Contains(key) {
		ret := m.Remove(key)
		h.size--
		if h.size < lowerTol*h.M && h.M/2 > initCapacity {
			h.resize(h.M / 2)
		}
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

func (h *HashTable) resize(newM int) {
	var newHashTable []*RBTree.RBTree
	for i := 0; i < newM; i++ {
		newHashTable = append(newHashTable, RBTree.Constructor())
	}
	oldM := h.M
	h.M = newM
	for i := 0; i < oldM; i++ {
		m := h.hashtable[i]
		for _, key := range m.KeySet() {
			newHashTable[h.hash(key)].Add(key, m.Get(key))
		}
	}
	h.hashtable = newHashTable
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
