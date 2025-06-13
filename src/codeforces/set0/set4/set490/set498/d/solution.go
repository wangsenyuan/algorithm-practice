package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
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
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	queries := make([]string, m)
	for i := range m {
		queries[i] = readString(reader)
	}
	return solve(a, queries)
}

func solve(a []int, queries []string) []int {
	tr := NewTree(a)

	var ans []int

	for _, cur := range queries {
		if cur[0] == 'A' {
			var l, r int
			pos := readInt([]byte(cur), 2, &l) + 1
			readInt([]byte(cur), pos, &r)
			l--
			r -= 2
			ans = append(ans, tr.Get(l, r))
		} else {
			var i, v int
			pos := readInt([]byte(cur), 2, &i) + 1
			readInt([]byte(cur), pos, &v)
			i--
			tr.Update(i, v)
		}
	}
	return ans
}

// va[i][j]表示当进入区间i时，离开这个区间时的时间是多少
type Tree struct {
	val [][]uint32
	sz  int
}

const LL = 60

func merge(c []uint32, a []uint32, b []uint32) {
	// a[i]表示左区间，开始时的时刻t % 6 = i 时，左区间要花费的时间
	// 如果只是6似乎不对，假设左区间花费了时间12
	// 那么求余6 = 0, 也就是说，如果a[i] = 2, 或者a[i] = 3 的时候，就不会+1了, 所以是不对的
	//
	for i := range LL {
		c[i] = a[i] + b[(uint32(i)+a[i])%LL]
	}
}

func NewTree(a []int) *Tree {
	n := len(a)
	val := make([][]uint32, 4*n)
	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		val[i] = make([]uint32, LL)

		if l == r {
			// 至少要花1秒到下一个位置
			for j := range LL {
				val[i][j] = 1
				if j%a[l] == 0 {
					val[i][j]++
				}
			}
			return
		}
		mid := (l + r) >> 1
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		merge(val[i], val[2*i+1], val[2*i+2])
	}
	build(0, 0, n-1)
	return &Tree{val, n}
}

func (tr *Tree) Update(pos int, val int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			for j := range LL {
				tr.val[i][j] = 1
				if j%val == 0 {
					tr.val[i][j]++
				}
			}
			return
		}
		mid := (l + r) >> 1
		if pos <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		merge(tr.val[i], tr.val[2*i+1], tr.val[2*i+2])
	}
	loop(0, 0, tr.sz-1)
}

func (tr *Tree) Get(L int, R int) int {
	var loop func(i int, l int, r int, L int, R int) []uint32
	loop = func(i int, l int, r int, L int, R int) []uint32 {
		if L == l && r == R {
			return tr.val[i]
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
		c := make([]uint32, LL)
		merge(c, a, b)
		return c
	}
	res := loop(0, 0, tr.sz-1, L, R)
	return int(res[0])
}
