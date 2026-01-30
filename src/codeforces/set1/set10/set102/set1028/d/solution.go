package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	n := readNum(reader)
	commands := make([]string, n)
	for i := range n {
		commands[i] = readString(reader)
	}
	return solve(commands)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

const inf = 1 << 60

func solve(commands []string) int {
	var prices []int
	var op []int
	for _, command := range commands {
		var p int
		if command[1] == 'D' {
			readInt([]byte(command), 4, &p)
			prices = append(prices, p)
		} else {
			readInt([]byte(command), 7, &p)
			p *= -1
		}
		op = append(op, p)
	}

	slices.Sort(prices)
	n := len(prices)

	set := NewSet(n)

	res := 1
	best := []int{-inf, inf}
	var rem int
	for _, p := range op {
		if p > 0 {
			if p > best[0] && p < best[1] {
				rem++
			}
			set.Add(sort.SearchInts(prices, p))
		} else {
			p *= -1

			if p < best[0] || best[1] < p {
				return 0
			}

			if best[0] < p && p < best[1] && rem > 0 {
				res = mul(res, 2)
			}
			j := sort.SearchInts(prices, p)
			set.Remove(j)
			rem = 0
			j1 := set.UpperBound(j)
			if j1 < n {
				best[1] = prices[j1]
			} else {
				best[1] = inf
			}
			j2 := set.LowerBound(j)
			if j2 >= 0 {
				best[0] = prices[j2]
			} else {
				best[0] = -inf
			}
		}
	}

	return mul(res, rem+1)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

type Set struct {
	t1 *SegTree
	t2 *SegTree
	n  int
}

func NewSet(n int) *Set {
	t1 := NewSegTree(n, inf, func(a, b int) int {
		return min(a, b)
	})
	t2 := NewSegTree(n, -inf, func(a, b int) int {
		return max(a, b)
	})
	return &Set{t1, t2, n}
}

func (set *Set) Add(p int) {
	set.t1.Update(p, p)
	set.t2.Update(p, p)
}

func (set *Set) Remove(p int) {
	set.t1.Update(p, inf)
	set.t2.Update(p, -inf)
}

func (set *Set) UpperBound(p int) int {
	return set.t1.Get(p, set.n)
}

func (set *Set) LowerBound(p int) int {
	return set.t2.Get(0, p+1)
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
