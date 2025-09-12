package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	ops := make([]string, m)
	for i := range m {
		ops[i] = readString(reader)
	}
	return solve(a, ops)
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

const X int32 = 10000

var lucky = []int{4, 7, 44, 47, 74, 77, 444, 447, 474, 477, 744, 747, 774, 777, 4444, 4447, 4474, 4477, 4744, 4747, 4774, 4777, 7444, 7447, 7474, 7477, 7744, 7747, 7774, 7777, 1e9}

type seg []struct{ cnt, minD, todo int }

func (t seg) maintain(i int) {
	t[i].cnt = t[2*i+1].cnt + t[2*i+2].cnt
	t[i].minD = min(t[2*i+1].minD, t[2*i+2].minD)
}

func Build(a []int) seg {
	n := len(a)
	tr := make(seg, 4*n)
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr.set(i, a, l)
			return
		}
		mid := (l + r) >> 1
		f(i*2+1, l, mid)
		f(i*2+2, mid+1, r)
		tr.maintain(i)
	}
	f(0, 0, n-1)
	return tr
}

func (tr seg) set(i int, a []int, pos int) {
	j := sort.SearchInts(lucky, a[pos])
	tr[i].minD = lucky[j] - a[pos]
	if tr[i].minD == 0 {
		tr[i].cnt = 1
	} else {
		tr[i].cnt = 0
	}
}

func (tr seg) apply(i int, d int) {
	tr[i].minD -= d
	tr[i].todo += d
}

func (tr seg) spread(i int) {
	if tr[i].todo != 0 {
		tr.apply(2*i+1, tr[i].todo)
		tr.apply(2*i+2, tr[i].todo)
		tr[i].todo = 0
	}
}

func (tr seg) Update(L int, R int, d int, a []int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if L <= l && r <= R && tr[i].minD > d {
			tr.apply(i, d)
			return
		}
		if l == r {
			a[l] += tr[i].todo + d
			tr[i].todo = 0
			tr.set(i, a, l)
			return
		}

		tr.spread(i)
		mid := (l + r) >> 1
		if L <= mid {
			f(2*i+1, l, mid)
		}
		if mid < R {
			f(2*i+2, mid+1, r)
		}
		tr.maintain(i)
	}

	f(0, 0, len(tr)/4-1)
}

func (tr seg) Query(L int, R int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if L <= l && r <= R {
			return tr[i].cnt
		}
		tr.spread(i)
		mid := (l + r) >> 1
		if R <= mid {
			return f(2*i+1, l, mid)
		}
		if mid < L {
			return f(2*i+2, mid+1, r)
		}
		return f(2*i+1, l, mid) + f(2*i+2, mid+1, r)
	}
	n := len(tr) / 4
	return f(0, 0, n-1)
}

func solve(a []int, ops []string) []int {
	tr := Build(a)
	var ans []int

	for _, cur := range ops {
		if strings.HasPrefix(cur, "add") {
			var l, r, d int
			pos := readInt([]byte(cur), 4, &l) + 1
			pos = readInt([]byte(cur), pos, &r) + 1
			readInt([]byte(cur), pos, &d)
			tr.Update(l-1, r-1, d, a)
		} else {
			// count
			var l, r int
			pos := readInt([]byte(cur), 6, &l) + 1
			readInt([]byte(cur), pos, &r)

			ans = append(ans, tr.Query(l-1, r-1))
		}
	}

	return ans
}
