package ds

import "fmt"

type HashMap struct {
	table [][]int
}

func NewHashMap() *HashMap {
	return &HashMap{
		table: make([][]int, 10),
	}
}

func (h *HashMap) Set(key int, value int) {
	index := key % len(h.table)
	h.table[index] = append(h.table[index], value)
}

func (h *HashMap) Get(key int) int {
	index := key % len(h.table)
	return h.table[index][0]
}

func (h *HashMap) Remove(key int) {
	index := key % len(h.table)
	h.table[index] = h.table[index][1:]
}

func (h *HashMap) Print() {
	for i, v := range h.table {
		fmt.Printf("%d: %v\n", i, v)
	}
}
