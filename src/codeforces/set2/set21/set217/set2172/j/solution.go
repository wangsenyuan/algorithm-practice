package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	for i, num := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}

func drive(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	h := readNNums(reader, n-1)
	return solve(a, h)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(a []int, h []int) []int {
	n := len(a)
	rem := make([][]int, n+1)
	for i, x := range a {
		rem[x] = append(rem[x], i)
	}
	open := make([][]int, n)
	for i, x := range h {
		open[x] = append(open[x], i)
	}

	set := NewDSU(n, a)
	for _, i := range open[0] {
		set.Union(i, i+1)
	}

	diff := make([]int, n+1)
	add := func(l int, r int, v int) {
		if l > r || v == 0 {
			return
		}
		diff[l] += v
		diff[r+1] -= v
	}
	flush := func(root int, y int) {
		root = set.Find(root)
		duration := y - set.last[root]
		if duration <= 0 {
			return
		}
		cnt := set.cnt[root]
		if cnt > 0 {
			add(set.r[root]-cnt+1, set.r[root], duration)
		}
		set.last[root] = y
	}

	for y := 2; y <= n; y++ {
		for _, i := range rem[y-1] {
			root := set.Find(i)
			flush(root, y)
			set.cnt[root]--
		}
		for _, i := range open[y-1] {
			a := set.Find(i)
			b := set.Find(i + 1)
			if a != b {
				flush(a, y)
				flush(b, y)
				set.Union(a, b)
			}
		}
	}

	for i := range n {
		root := set.Find(i)
		if i == root {
			flush(root, n+1)
		}
	}

	res := make([]int, n)
	var cur int
	for i := range n {
		cur += diff[i]
		res[i] = cur
	}
	return res
}

type DSU struct {
	arr  []int
	l    []int
	r    []int
	cnt  []int
	last []int
}

func NewDSU(n int, a []int) *DSU {
	set := new(DSU)
	set.arr = make([]int, n)
	set.l = make([]int, n)
	set.r = make([]int, n)
	set.cnt = make([]int, n)
	set.last = make([]int, n)
	for i := range n {
		set.arr[i] = i
		set.l[i] = i
		set.r[i] = i
		if a[i] > 0 {
			set.cnt[i] = 1
		}
		set.last[i] = 1
	}
	return set
}

func (set *DSU) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *DSU) Union(a int, b int) int {
	a = set.Find(a)
	b = set.Find(b)
	if a == b {
		return a
	}
	if set.r[a] > set.r[b] {
		a, b = b, a
	}
	set.arr[b] = a
	set.l[a] = min(set.l[a], set.l[b])
	set.r[a] = max(set.r[a], set.r[b])
	set.cnt[a] += set.cnt[b]
	return a
}
