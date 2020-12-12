package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func main() {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(h)
	fmt.Println(*h)
	fmt.Println(heap.Pop(h)) // 调用pop
	heap.Push(h, 7)
	fmt.Println("new", *h)
	for len(*h) > 0 { //持续推出顶部最小元素
		fmt.Printf("%d\n", heap.Pop(h))
	}
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	fmt.Printf("old:", old)

	x := old[n-1]
	*h = old[0 : n-1]
	fmt.Println(*h)
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
