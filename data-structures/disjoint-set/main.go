package main

import "fmt"

type UnionFind struct {
	root []int
}

func (uf *UnionFind) make(size int) {
	uf.root = make([]int, size)
	for i := 0; i < size; i++ {
		uf.root[i] = i
	}
}

func (uF *UnionFind) find(x int) int {
	return uF.root[x]
}

func (uf *UnionFind) connected(x int, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *UnionFind) union(x int, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		for i := 0; i < len(uf.root); i++ {
			if uf.root[i] == rootY {
				uf.root[i] = rootX
			}
		}
	}
}

func main() {
	var uf UnionFind = UnionFind{}
	uf.make(10)
	uf.union(1, 2)
	uf.union(2, 5)
	uf.union(5, 6)
	uf.union(6, 7)
	uf.union(3, 8)
	uf.union(8, 9)
	fmt.Println(uf.connected(1, 5)) //true
	fmt.Println(uf.connected(5, 7)) //true
	fmt.Println(uf.connected(4, 9)) //false
}
