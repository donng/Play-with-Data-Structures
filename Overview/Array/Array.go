package Array

import "Play-with-Data-Structures/Utils/Interfaces"

type Array struct {
	data []interface{}
	size int
}

func Constructor(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
	}
}

func (a *Array) GetCapacity() int {
	return len(a.data)
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = e
	a.size++
}

func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}
	return a.data[index]
}

func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed. Index is illegal.")
	}
	a.data[index] = e
}

func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			return true
		}
	}
	return false
}

func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

func (a *Array) FindAll(e interface{}) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			indexes = append(indexes, i)
		}
	}
	return
}

func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Set failed. Index is illegal.")
	}

	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil

	if a.size == len(a.data)/2 {
		a.resize(len(a.data) / 2)
	}
	return e
}

func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveElement(e interface{}) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}

	a.Remove(index)
	return true
}

func (a *Array) RemoveAllElement(e interface{}) bool {
	if a.Find(e) == -1 {
		return false
	}

	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			a.Remove(i)
		}
	}
	return true
}

func (a *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}

	a.data = newData
}
