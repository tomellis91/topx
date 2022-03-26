package cmd

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func scan(path string, n int) {
	fmt.Printf("Reading file from the path: '%s'\n", path)
	fmt.Printf("Checking for the top '%d' numbers\n", n)

	h := &IntHeap{}
	heap.Init(h)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num := scanner.Text()

		i, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("could not convert string %s to int\n", num)
			continue
		}

		// have not hit N keep adding
		if len(*h) < n {
			heap.Push(h, i)
			continue
		}

		cm := (*h)[0]
		// add if greater than currentMin
		if i > cm {
			// remove cm and add new cm value
			heap.Pop(h)
			heap.Push(h, i)
		}

	}
	sort.Sort(sort.Reverse(*h))
	fmt.Printf("The top %d values are %v\n", n, *h)
}
