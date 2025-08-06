package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	r, _ := os.Open("input.txt")
	reader := bufio.NewReader(r)

	w, _ := os.Create("output.txt")

	writer := bufio.NewWriter(w)

	defer writer.Flush()

	cnt, best := process(reader)

	writer.WriteString(fmt.Sprintf("%d\n%s\n", cnt, best))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (int, string) {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) (int, string) {
	freq := make([]int, 26)
	n := len(s)

	pos := make([][]int, 26)
	for i := range n {
		x := int(t[i] - 'A')
		freq[x]++

		y := int(s[i] - 'A')
		pos[y] = append(pos[y], i)
	}

	t2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})

	add := func(i int) {
		t2.Update(i, i)
	}

	rem := func(i int) {
		t2.Update(i, n)
	}

	var cnt int
	stop := make([]int, 26)
	for i := range 26 {
		stop[i] = n
	}
	// 先把头部多余的加入
	for x := range 26 {
		if len(pos[x]) > freq[x] {
			cnt += len(pos[x]) - freq[x]
			j := len(pos[x]) - freq[x]
			if j < len(pos[x]) {
				stop[x] = pos[x][j]
			}
			for i := range j {
				add(pos[x][i])
			}
		}
	}

	buf := []byte(s)

	for x := range 26 {
		if len(pos[x]) < freq[x] {
			need := freq[x] - len(pos[x])
			for range need {
				it := t2.Get(0, n)
				y := int(s[it] - 'A')
				if x < y {
					// 如果是 x > y， 是从后边开始处理的
					pos[y] = pos[y][1:]
				}
				buf[it] = byte('A' + x)
				rem(it)
			}
		} else if len(pos[x]) > freq[x] {
			// x多出来了，把它后面的部分给加入进去
			// 先把它之前记录的部分去掉
			for i := 0; i < len(pos[x]) && pos[x][i] < stop[x]; i++ {
				rem(pos[x][i])
			}
			for len(pos[x]) > freq[x] {
				ln := len(pos[x])
				it := pos[x][ln-1]
				add(it)
				pos[x] = pos[x][:ln-1]
			}
		}
	}

	return cnt, string(buf)
}

type SegTree struct {
	size       int
	arr        []int
	init_value int
	op         func(int, int) int
}

func NewSegTree(n int, v int, op func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = v
	}
	return &SegTree{n, arr, v, op}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = seg.op(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := seg.init_value
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = seg.op(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = seg.op(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
