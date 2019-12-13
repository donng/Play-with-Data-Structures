package HashTable

import (
	"Play-with-Data-Structures/14-Hash-Table/07-More-About-Resizing-in-Hash-Table/src/RBTree"
	"fmt"
	"hash/fnv"
	"strconv"
)

const upperTol = 10
const lowerTol = 2

var capacityIndex = 0
var capacity = []int{53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593,
	49157, 98317, 196613, 393241, 786433, 1572869, 3145739, 6291469,
	12582917, 25165843, 50331653, 100663319, 201326611, 402653189, 805306457, 1610612741}

type HashTable struct {
	hashtable []*RBTree.RBTree
	M         int
	size      int
}

func Constructor() *HashTable {
	M := capacity[capacityIndex]
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

		if h.size >= upperTol*h.M && capacityIndex+1 < len(capacity) {
			capacityIndex++
			h.resize(capacity[capacityIndex])
		}
	}
}

func (h *HashTable) Remove(key interface{}) interface{} {
	m := h.hashtable[h.hash(key)]
	if m.Contains(key) {
		ret := m.Remove(key)
		h.size--
		if h.size < lowerTol*h.M && capacityIndex-1 >= 0 {
			capacityIndex--
			h.resize(capacity[capacityIndex])
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
