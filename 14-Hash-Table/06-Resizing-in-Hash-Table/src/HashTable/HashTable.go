package HashTable

import (
	"Play-with-Data-Structures/14-Hash-Table/06-Resizing-in-Hash-Table/src/RBTree"
	"strconv"
	"hash/fnv"
	"fmt"
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

func (this *HashTable) GetSize() int {
	return this.size
}

func (this *HashTable) Add(key interface{}, value interface{}) {
	m := this.hashtable[this.hash(key)]
	if m.Contains(key) {
		m.Add(key, value)
	} else {
		m.Add(key, value)
		this.size++

		if this.size >= upperTol*this.M {
			this.resize(2 * this.M)
		}
	}
}

func (this *HashTable) Remove(key interface{}) interface{} {
	m := this.hashtable[this.hash(key)]
	if m.Contains(key) {
		ret := m.Remove(key)
		this.size--
		if this.size < lowerTol*this.M && this.M/2 > initCapacity {
			this.resize(this.M / 2)
		}
		return ret
	} else {
		return nil
	}
}

func (this *HashTable) Set(key interface{}, value interface{}) {
	m := this.hashtable[this.hash(key)]
	if !m.Contains(key) {
		panic(fmt.Sprintf("%s doesn't exist!", key))
	}
	m.Set(key, value)
}

func (this *HashTable) Contains(key interface{}) bool {
	return this.hashtable[this.hash(key)].Contains(key)
}

func (this *HashTable) Get(key interface{}) interface{} {
	return this.hashtable[this.hash(key)].Get(key)
}

func (this *HashTable) resize(newM int) {
	var newHashTable []*RBTree.RBTree
	for i := 0; i < newM; i++ {
		newHashTable = append(newHashTable, RBTree.Constructor())
	}
	oldM := this.M
	this.M = newM
	for i := 0; i < oldM; i++ {
		m := this.hashtable[i]
		for _, key := range m.KeySet() {
			newHashTable[this.hash(key)].Add(key, m.Get(key))
		}
	}
	this.hashtable = newHashTable
}

func (this *HashTable) hash(key interface{}) int {
	return (int(HashCode(key)) & 0x7fffffff) % this.M
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
