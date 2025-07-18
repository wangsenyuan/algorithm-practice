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
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	qs := make([][]int, m)
	for i := range m {
		qs[i] = readNNums(reader, 2)
	}
	return solve(a, qs)
}

func solve(a []int, queries [][]int) []int {
	n := len(a)
	at := make([][]int, n)
	for i, cur := range queries {
		r := cur[1] - 1
		at[r] = append(at[r], i)
	}

	nums := slices.Clone(a)
	sort.Ints(nums)
	nums = slices.Compact(nums)

	k := len(nums)
	pos := make([]int, k)
	for i := range k {
		pos[i] = -1
	}

	set := make(BIT, n+2)

	sums := make([]int, n+1)

	ans := make([]int, len(queries))
	for r := range n {
		sums[r+1] = sums[r] ^ a[r]
		i := sort.SearchInts(nums, a[r])
		if pos[i] >= 0 {
			set.update(pos[i], a[r])
		}
		pos[i] = r
		set.update(r, a[r])

		for _, j := range at[r] {
			l := queries[j][0] - 1
			tmp := set.query(l, r)
			ans[j] = tmp ^ sums[r+1] ^ sums[l]
		}
	}

	return ans
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] ^= v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int
	for i > 0 {
		res ^= bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) query(l int, r int) int {
	res := bit.get(r)
	if l > 0 {
		res ^= bit.get(l - 1)
	}
	return res
}
