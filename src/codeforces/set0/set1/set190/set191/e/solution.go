package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(k, a)
}

const inf = 1 << 50

func solve(k int, a []int) int {
	n := len(a)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + a[i]
	}
	s := sortAndUnique(sum)
	m := len(s)

	pos := make([]int, n+1)

	for i := range n + 1 {
		pos[i] = sort.SearchInts(s, sum[i])
	}

	bit := make(BIT, m+3)

	// x越小，越满足条件

	check := func(x int) bool {
		clear(bit)

		bit.Add(pos[n], 1)

		var cnt int

		for i := n - 1; i >= 0; i-- {
			j := sort.SearchInts(s, sum[i]+x)
			tmp := bit.Query(j, m-1)
			cnt += tmp
			bit.Add(pos[i], 1)
		}

		return cnt >= k
	}

	r := inf
	l := -inf

	offset := inf

	l += offset
	r += offset

	for l < r {
		mid := (l + r) / 2
		if check(mid - offset) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return r - 1 - offset
}

func sortAndUnique(a []int) []int {
	b := slices.Clone(a)
	sort.Ints(b)
	return slices.Compact(b)
}

type BIT []int

func (bit BIT) Add(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) Get(i int) int {
	var res int
	i++
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) Query(l int, r int) int {
	res := bit.Get(r)
	if l > 0 {
		res -= bit.Get(l - 1)
	}
	return res
}
