package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

const B = 300

func solve(a []int, queries [][]int) []int {
	type Q struct {
		l  int
		r  int
		id int
	}

	qs := make([]Q, len(queries))

	for i, cur := range queries {
		l, r := cur[0], cur[1]
		qs[i] = Q{l - 1, r - 1, i}
	}

	cmp := func(a Q, b Q) int {
		if a.r/B != b.r/B {
			return a.r - b.r
		}
		return b.l - a.l
	}

	slices.SortFunc(qs, cmp)

	x := slices.Max(a)

	freq := make([]int, x+1)

	var score int

	update := func(x int, v int) {
		y := freq[x]
		score -= y * y * x
		freq[x] += v
		y = freq[x]
		score += y * y * x
	}

	var L, R int
	ans := make([]int, len(queries))

	for _, cur := range qs {
		for R <= cur.r {
			update(a[R], 1)
			R++
		}
		for R-1 > cur.r {
			R--
			update(a[R], -1)
		}

		for L > cur.l {
			L--
			update(a[L], 1)
		}
		for L < cur.l {
			update(a[L], -1)
			L++
		}
		ans[cur.id] = score
	}
	return ans
}
