package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ele := []int{1, 1, 1, 2, 2, 2, 4, 4, 4, 4, 4, 5, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7}
	num := 2
	a := topKFrequent(ele, num)
	fmt.Println(a)
}

type Element struct {
	value     int
	frequency int
}

type Heap []Element

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].frequency < h[j].frequency
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// push and pop work on interface{}
func (h *Heap) Push(x interface{}) {
	// specify the interface type
	// append to dereferenced pointer
	*h = append(*h, x.(Element))
}

func (h *Heap) Pop() interface{} {
	var x interface{}
	// set x to the last ele (get it)
	// set the dereferenced pointer to h, to everything before the last element (pop it off)
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	hs := map[int]int{}
	// Heap{} is just []Element
	minHeap := Heap{}

	// initialize answer
	res := make([]int, 0, k)

	// collect hash map (idx => frequency )
	for _, v := range nums {
		hs[v]++
	}

	// fill heap
	heap.Init(&minHeap)
	for i, v := range hs {
		heap.Push(&minHeap, Element{i, v})

		if len(minHeap) > k {
			heap.Pop(&minHeap)
		}
	}

	for _, elem := range minHeap {
		res = append(res, elem.value)
	}

	// reverse a slice using index
	// loop backwards
	for i := len(res)/2 - 1; i >= 0; i-- {
		// calculate forwrds index
		opp := len(res) - 1 - i
		//switch
		res[i], res[opp] = res[opp], res[i]
	}

	return res
}
