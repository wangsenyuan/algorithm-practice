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
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
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

func process(reader *bufio.Reader) string {
	_, m := readTwoNums(reader)
	s := readString(reader)
	ops := make([]string, m)
	for i := range m {
		ops[i] = readString(reader)
	}
	return solve(s, ops)
}

func solve(s string, ops []string) string {
	m := len(ops)

	pos := make([][]SegTree, 3)
	for i := range pos {
		pos[i] = make([]SegTree, 3)
		for j := range pos[i] {
			pos[i][j] = NewSegTree(m)
		}
	}

	for i := m - 1; i >= 0; i-- {
		x, y := int(ops[i][0]-'a'), int(ops[i][2]-'a')
		pos[x][y].Update(i, i)
	}

	buf := []byte(s)

	for i := range len(buf) {
		if buf[i] == 'a' {
			continue
		}
		x := int(buf[i] - 'a')
		if x == 1 {
			// b -> a
			j := pos[1][0].Query(0, m)
			if j < inf {
				buf[i] = 'a'
				pos[1][0].Update(j, inf)
				continue
			}
			// j == inf
			// b -> c -> a
			j = pos[1][2].Query(0, m)
			if j < inf {
				k := pos[2][0].Query(j, m)
				if k < inf {
					buf[i] = 'a'
					pos[2][0].Update(k, inf)
					pos[1][2].Update(j, inf)
					continue
				}
			}
			// 保留b是最好的选择
			continue
		}
		// x == 2
		j := pos[2][0].Query(0, m)
		if j < inf {
			buf[i] = 'a'
			pos[2][0].Update(j, inf)
			continue
		}
		// c -> b -> a
		j = pos[2][1].Query(0, m)
		if j < inf {
			k := pos[1][0].Query(j, m)
			if k < inf {
				buf[i] = 'a'
				pos[1][0].Update(k, inf)
				pos[2][1].Update(j, inf)
				continue
			}
			// 变成b是最好的选择
			buf[i] = 'b'
			pos[2][1].Update(j, inf)
		}
	}

	return string(buf)
}

type SegTree []int

const inf = 1 << 30

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = inf
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	p += len(tr) / 2
	tr[p] = v
	for p > 1 {
		tr[p>>1] = min(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Query(l, r int) int {
	l += len(tr) / 2
	r += len(tr) / 2
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
