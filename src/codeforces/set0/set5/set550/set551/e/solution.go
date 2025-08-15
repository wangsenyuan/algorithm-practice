package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var l, r, x int
			fmt.Fscan(reader, &l, &r, &x)
			queries[i] = []int{1, l, r, x}
		} else {
			var y int
			fmt.Fscan(reader, &y)
			queries[i] = []int{2, y}
		}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	block_size := int(math.Sqrt(float64(n)))

	var blocks []*Segment
	for i := 0; i < n; i += block_size {
		j := min(n, i+block_size)
		st := NewSegment(a[i:j])
		blocks = append(blocks, st)
	}

	var res []int

	for _, cur := range queries {
		if cur[0] == 1 {
			l, r, x := cur[1], cur[2], cur[3]
			l--
			r--
			bl := l / block_size
			br := r / block_size

			if bl == br {
				blocks[bl].updateRange(l-bl*block_size, r-bl*block_size, x)
				continue
			}

			for i := bl + 1; i < br; i++ {
				blocks[i].updateOffset(x)
			}
			if l%block_size == 0 {
				blocks[bl].updateOffset(x)
			} else {
				blocks[bl].updateRange(l-bl*block_size, block_size-1, x)
			}

			if (r+1)%block_size == 0 {
				blocks[br].updateOffset(x)
			} else {
				blocks[br].updateRange(0, r-br*block_size, x)
			}
		} else {
			y := cur[1]
			l, r := n, -1
			for i := range blocks {
				pos := blocks[i].find(y)
				if pos != nil {
					l = min(l, pos[0]+i*block_size)
					r = max(r, pos[1]+i*block_size)
				}
			}
			res = append(res, max(-1, r-l))
		}
	}

	return res
}

type pair struct {
	first  int
	second int
}

func cmp_pair(a, b pair) int {
	return cmp.Or(a.first-b.first, a.second-b.second)
}

type Segment struct {
	offset     int
	arr        []int
	sorted_arr []pair
}

func NewSegment(arr []int) *Segment {
	sorted_arr := make([]pair, len(arr))
	for i, v := range arr {
		sorted_arr[i] = pair{first: v, second: i}
	}
	slices.SortFunc(sorted_arr, cmp_pair)
	return &Segment{
		offset:     0,
		arr:        arr,
		sorted_arr: sorted_arr,
	}
}

func (this *Segment) updateOffset(offset int) {
	this.offset += offset
}

func (this *Segment) updateRange(l int, r int, x int) {
	if l == 0 && r == len(this.arr)-1 {
		this.updateOffset(x)
		return
	}

	if this.offset != 0 {
		for i := range this.arr {
			this.arr[i] += this.offset
		}
		this.offset = 0
	}
	for i := l; i <= min(r, len(this.arr)-1); i++ {
		this.arr[i] += x
	}

	for i, v := range this.arr {
		this.sorted_arr[i] = pair{first: v, second: i}
	}
	slices.SortFunc(this.sorted_arr, cmp_pair)
}

func (this *Segment) find(y int) []int {
	y -= this.offset
	r := sort.Search(len(this.sorted_arr), func(i int) bool {
		return this.sorted_arr[i].first > y
	})
	l := sort.Search(len(this.sorted_arr), func(i int) bool {
		return this.sorted_arr[i].first >= y
	})
	if l < r {
		return []int{this.sorted_arr[l].second, this.sorted_arr[r-1].second}
	}
	return nil
}
