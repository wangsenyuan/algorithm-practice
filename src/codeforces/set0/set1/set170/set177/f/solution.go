package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

type edge struct{ h, w, r int }

func drive(reader *bufio.Reader) int {
	var n, k, t int
	fmt.Fscan(reader, &n, &k, &t)
	_ = n
	edges := make([]edge, k)
	for i := range edges {
		fmt.Fscan(reader, &edges[i].h, &edges[i].w, &edges[i].r)
		edges[i].h--
		edges[i].w--
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].r < edges[j].r
	})
	return solve(t, edges)
}

// state represents a completed matching:
//   value    = total weight
//   lastI    = index of the last edge added (edges are sorted by r)
//   menMask  = bitmask of men used
//   womenMask= bitmask of women used
type state struct {
	value, lastI, menMask, womenMask int
}

type minHeap []state

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].value < h[j].value }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)         { *h = append(*h, x.(state)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// solve enumerates matchings in non-decreasing order of value using a min-heap.
// Each matching {e_{i1} < e_{i2} < ... < e_{im}} is generated exactly once
// from the matching {e_{i1}, ..., e_{i(m-1)}} by extension with e_{im}.
// The empty matching (value=0, lastI=-1) is the seed.
func solve(t int, edges []edge) int {
	k := len(edges)
	h := &minHeap{{value: 0, lastI: -1, menMask: 0, womenMask: 0}}
	heap.Init(h)

	for count := 1; ; count++ {
		s := heap.Pop(h).(state)
		if count == t {
			return s.value
		}
		// Extend by adding any compatible edge with index > lastI.
		for j := s.lastI + 1; j < k; j++ {
			e := edges[j]
			if (s.menMask>>e.h)&1 == 0 && (s.womenMask>>e.w)&1 == 0 {
				heap.Push(h, state{
					value:     s.value + e.r,
					lastI:     j,
					menMask:   s.menMask | (1 << e.h),
					womenMask: s.womenMask | (1 << e.w),
				})
			}
		}
	}
}
