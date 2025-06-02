package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
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
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

type query struct {
	l  int
	r  int
	id int
}

const B = 300

func cmp(a query, b query) int {
	if a.r/B != b.r/B {
		return a.r/B - b.r/B
	}
	// 第几个区间
	i := a.r / B
	if i&1 == 0 {
		return a.l - b.l
	}
	return b.l - a.l
}

func solve(a []int, queries [][]int) []int {

	qs := make([]query, len(queries))
	for i, cur := range queries {
		qs[i] = query{l: cur[0] - 1, r: cur[1] - 1, id: i}
	}

	slices.SortFunc(qs, cmp)

	n := len(a)
	freq := make([]int, n+1)
	var cnt int

	add := func(i int) {
		if a[i] <= n {
			if freq[a[i]] == a[i] {
				cnt--
			}
			freq[a[i]]++
			if freq[a[i]] == a[i] {
				cnt++
			}
		}
	}

	rem := func(i int) {
		if a[i] <= n {
			if freq[a[i]] == a[i] {
				cnt--
			}
			freq[a[i]]--
			if freq[a[i]] == a[i] {
				cnt++
			}
		}
	}

	ans := make([]int, len(queries))
	var l, r int
	for _, cur := range qs {
		for r <= cur.r {
			add(r)
			r++
		}
		for r-1 > cur.r {
			r--
			rem(r)
		}
		for l < cur.l {
			rem(l)
			l++
		}
		for l > cur.l {
			l--
			add(l)
		}
		ans[cur.id] = cnt
	}

	return ans
}
