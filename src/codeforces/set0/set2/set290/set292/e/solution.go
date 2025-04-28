package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	queries := make([][]int, m)
	for i := range m {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			queries[i] = make([]int, 4)
			queries[i][0] = 1
		} else {
			queries[i] = make([]int, 2)
			queries[i][0] = 2
		}
		pos := 2
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	return solve(a, b, queries)
}

func solve(a []int, b []int, queries [][]int) []int {
	n := len(a)
	tr := NewSegTree(n)

	var ans []int

	for i, cur := range queries {
		if cur[0] == 1 {
			y, k := cur[2], cur[3]
			y--
			// 必须range update
			tr.Update(y, y+k, i)
		} else {
			p := cur[1] - 1
			j := tr.Get(p)
			if j < 0 {
				ans = append(ans, b[p])
			} else {
				x, y := queries[j][1], queries[j][2]
				x--
				y--
				// j - y < k and j - y >= 0
				offset := p - y
				ans = append(ans, a[x+offset])
			}
		}
	}
	return ans
}

const inf = 1 << 50

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		arr[i] = -1
	}
	return SegTree(arr)
}

func (tr SegTree) Update(l int, r int, v int) {
	n := len(tr) / 2
	l += n
	r += n
	for l < r {
		if l&1 == 1 {
			tr[l] = max(tr[l], v)
			l++
		}
		if r&1 == 1 {
			r--
			tr[r] = max(tr[r], v)
		}
		l >>= 1
		r >>= 1
	}
}

func (tr SegTree) Get(p int) int {
	n := len(tr) / 2
	p += n

	res := -1

	for p > 0 {
		res = max(res, tr[p])
		p >>= 1
	}
	return res
}
