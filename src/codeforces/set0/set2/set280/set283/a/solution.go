package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%.7f\n", x))
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

func process(reader *bufio.Reader) []float64 {
	n := readNum(reader)
	ops := make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			ops[i] = make([]int, 3)
		} else if s[0] == '2' {
			ops[i] = make([]int, 2)
		} else {
			ops[i] = make([]int, 1)
		}
		var pos int
		for j := range len(ops[i]) {
			pos = readInt(s, pos, &ops[i][j]) + 1
		}
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []float64 {
	tr := make(SegTree, 2*(n+1))
	sum := 0

	var res []float64

	stack := make([]int, n+1)
	top := 1

	for _, cur := range ops {
		if cur[0] == 1 {
			a, x := cur[1], cur[2]
			a = min(a, top)
			tr.Update(0, a, x)
			sum += a * x
		} else if cur[0] == 2 {
			stack[top] = cur[1]
			top++
			sum += cur[1]
		} else {
			// cur[0] == 3
			if top > 1 {
				v := tr.Get(top - 1)
				// 把这里的增量给取消掉
				tr.Update(top-1, top, -v)
				sum -= v
				sum -= stack[top-1]
				top--
			}
		}
		res = append(res, float64(sum)/float64(top))
	}
	return res
}

type SegTree []int

func (tr SegTree) Update(l int, r int, v int) {
	n := len(tr) / 2
	l += n
	r += n
	for l < r {
		if l&1 == 1 {
			tr[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			tr[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (tr SegTree) Get(p int) int {
	n := len(tr) / 2
	p += n
	var res int
	for p > 0 {
		res += tr[p]
		p >>= 1
	}
	return res
}
