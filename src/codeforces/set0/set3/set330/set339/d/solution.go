package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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
	a := readNNums(reader, 1<<n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)

	vals := make([]int, 2*n)

	h := bits.Len(uint(n))

	op := make([]int, h+1)
	op[1] = 0
	for i := 2; i <= h; i++ {
		op[i] = op[i-1] ^ 1
	}

	merge := func(i int, dist int) {
		level := bits.Len(uint(dist)) - 1
		if op[level] == 0 {
			vals[i] = vals[2*i] | vals[2*i+1]
		} else {
			vals[i] = vals[2*i] ^ vals[2*i+1]
		}
	}

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			vals[i] = a[l]
			return
		}
		mid := (l + r) >> 1
		build(2*i, l, mid)
		build(2*i+1, mid+1, r)
		merge(i, r-l+1)
	}
	build(1, 0, n-1)

	var update func(i int, l int, r int, pos int, v int)

	update = func(i int, l int, r int, pos int, v int) {
		if l == r {
			vals[i] = v
			return
		}
		mid := (l + r) >> 1
		if pos <= mid {
			update(2*i, l, mid, pos, v)
		} else {
			update(2*i+1, mid+1, r, pos, v)
		}
		merge(i, r-l+1)
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		p, b := cur[0]-1, cur[1]
		update(1, 0, n-1, p, b)
		ans[i] = vals[1]
	}

	return ans
}
