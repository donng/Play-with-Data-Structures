package Array

type Array struct {
	data []interface{}
	size int
}

func Constructor(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
	}
}

func (this *Array) GetCapacity() int {
	return len(this.data)
}

func (this *Array) GetSize() int {
	return this.size
}

func (this *Array) IsEmpty() bool {
	return this.size == 0
}

func (this *Array) Add(index int, e interface{}) {
	if index < 0 || index > this.size {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	if this.size == len(this.data) {
		this.resize(2 * this.size)
	}

	for i := this.size - 1; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}

	this.data[index] = e
	this.size++
}

func (this *Array) AddLast(e interface{}) {
	this.Add(this.size, e)
}

func (this *Array) AddFirst(e interface{}) {
	this.Add(0, e)
}

func (this *Array) Get(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Get failed. Index is illegal.")
	}
	return this.data[index]
}

func (this *Array) Set(index int, e interface{}) {
	if index < 0 || index >= this.size {
		panic("Set failed. Index is illegal.")
	}
	this.data[index] = e
}

func (this *Array) Contains(e interface{}) bool {
	for i := 0; i < this.size; i++ {
		if Interfaces.Compare(this.data[i], e) == 0 {
			return true
		}
	}
	return false
}

func (this *Array) Find(e interface{}) int {
	for i := 0; i < this.size; i++ {
		if Interfaces.Compare(this.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

func (this *Array) FindAll(e interface{}) (indexes []int) {
	for i := 0; i < this.size; i++ {
		if Interfaces.Compare(this.data[i], e) == 0 {
			indexes = append(indexes, i)
		}
	}
	return
}

func (this *Array) Remove(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Set failed. Index is illegal.")
	}

	e := this.data[index]
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	this.size--
	this.data[this.size] = nil

	if this.size == len(this.data)/2 {
		this.resize(len(this.data) / 2)
	}
	return e
}

func (this *Array) RemoveFirst() interface{} {
	return this.Remove(0)
}

func (this *Array) RemoveLast() interface{} {
	return this.Remove(this.size - 1)
}

func (this *Array) RemoveElement(e interface{}) bool {
	index := this.Find(e)
	if index == -1 {
		return false
	}

	this.Remove(index)
	return true
}

func (this *Array) RemoveAllElement(e interface{}) bool {
	if this.Find(e) == -1 {
		return false
	}

	for i := 0; i < this.size; i++ {
		if Interfaces.Compare(this.data[i], e) == 0 {
			this.Remove(i)
		}
	}
	return true
}

func (this *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[i]
	}

	this.data = newData
}
