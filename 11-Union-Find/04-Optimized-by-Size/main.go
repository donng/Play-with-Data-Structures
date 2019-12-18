package _4_Optimized_by_Size

import (
	"Play-with-Data-Structures/11-Union-Find/04-Optimized-by-Size/src/UF"
	"Play-with-Data-Structures/11-Union-Find/04-Optimized-by-Size/src/UnionFind1"
	"Play-with-Data-Structures/11-Union-Find/04-Optimized-by-Size/src/UnionFind2"
	"Play-with-Data-Structures/11-Union-Find/04-Optimized-by-Size/src/UnionFind3"
	"fmt"
	"math/rand"
	"time"
)

func testUF(uf UF.UF, m int) time.Duration {
	size := uf.GetSize()
	rand.Seed(time.Now().Unix())

	startTime := time.Now()

	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.UnionElements(a, b)
	}

	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.IsConnected(a, b)
	}

	return time.Now().Sub(startTime)
}

func main() {
	// UnionFind1 慢于 UnionFind2
	//size := 100000
	//m := 10000

	// UnionFind2 慢于 UnionFind1, 但UnionFind3最快
	size := 100000
	m := 100000

	uf1 := UnionFind1.Constructor(size)
	fmt.Println(testUF(uf1, m))

	uf2 := UnionFind2.Constructor(size)
	fmt.Println(testUF(uf2, m))

	uf3 := UnionFind3.Constructor(size)
	fmt.Println(testUF(uf3, m))
}
