package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
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
		if x == len(bs) {
			return res[:i]
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	ops := make([][]int, m)
	for i := 0; i < m; i++ {
		ops[i] = readNNums(reader, 3)
	}
	return solve(a, ops)
}

func solve(a []int, ops [][]int) []int {
	nums := slices.Clone(a)
	arr := slices.Clone(nums)

	for _, op := range ops {
		if op[0] == 1 {
			p, d := op[1]-1, op[2]
			nums[p] += d
			arr = append(arr, nums[p])
		}
	}
	sort.Ints(arr)
	arr = slices.Compact(arr)
	m := len(arr)
	tr := NewTree(m)
	for _, v := range a {
		i := sort.SearchInts(arr, v)
		tr.Update(i, v, true)
	}
	copy(nums, a)
	var ans []int
	for _, op := range ops {
		if op[0] == 1 {
			p, d := op[1]-1, op[2]
			i := sort.SearchInts(arr, nums[p])
			tr.Update(i, nums[p], false)
			nums[p] += d
			i = sort.SearchInts(arr, nums[p])
			tr.Update(i, nums[p], true)
		} else {
			l, r := op[1], op[2]
			// arr[i] >= l
			i := sort.SearchInts(arr, l)
			j := sort.SearchInts(arr, r)
			if j == m || r < arr[j] {
				j--
			}
			if i <= j {
				ans = append(ans, tr.Query(i, j))
			} else {
				ans = append(ans, 0)
			}
		}
	}

	return ans
}

type data struct {
	cnt int
	sum int
	val int
}

func merge(l data, r data) data {
	// 主要是计算val
	val := l.cnt*r.sum - l.sum*r.cnt
	val += r.val + l.val

	// 假设左边是a, b, 右边是 x, y
	// x - a + x - b + y - a + y - b
	// 2 * x - (a + b) + 2 * y - (a + b) = 2 * (x + y) - 2 * (a + b)
	// (a, b) x => 2 * x - (a + b)
	// a, (x, y) = x + y - 2 * a
	return data{l.cnt + r.cnt, l.sum + r.sum, val}
}

type Tree struct {
	arr []data
	sz  int
}

func NewTree(n int) *Tree {
	arr := make([]data, 4*n)
	return &Tree{arr, n}
}

func (t *Tree) Update(p int, v int, on bool) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.arr[i].val = 0
			if on {
				t.arr[i].cnt = 1
				t.arr[i].sum = v
			} else {
				t.arr[i].cnt = 0
				t.arr[i].sum = 0
			}
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		t.arr[i] = merge(t.arr[2*i+1], t.arr[2*i+2])
	}

	loop(0, 0, t.sz-1)
}

func (t *Tree) Query(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) data
	loop = func(i int, l int, r int, L int, R int) data {
		if L == l && R == r {
			return t.arr[i]
		}
		mid := (l + r) >> 1
		if R <= mid {
			return loop(2*i+1, l, mid, L, R)
		}
		if mid < L {
			return loop(2*i+2, mid+1, r, L, R)
		}
		a := loop(2*i+1, l, mid, L, mid)
		b := loop(2*i+2, mid+1, r, mid+1, R)
		return merge(a, b)
	}
	return loop(0, 0, t.sz-1, L, R).val
}
