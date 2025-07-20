package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		readNum(reader)
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) int {
	n := len(s)
	pref := make([]int, n+1)
	var arr []int
	arr = append(arr, 0)
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		pref[i+1] = pref[i] + x
		arr = append(arr, 2*pref[i+1]-(i+1))
	}
	sort.Ints(arr)
	arr = slices.Compact(arr)
	m := len(arr)
	cnt := make(BIT, m+2)
	cnt.Add(sort.SearchInts(arr, 0), 1)
	sum1 := make(BIT, m+2)
	sum0 := make(BIT, m+2)

	var res int
	for i := 1; i <= n; i++ {
		j := sort.SearchInts(arr, 2*pref[i]-i)
		// 寻找[l...r]区间，且1的个数 >= 0的部分
		c1 := cnt.QueryRange(0, j)
		s1 := sum1.QueryRange(0, j)
		res += c1*pref[i] - s1

		sum1.Add(j, pref[i])

		c0 := cnt.QueryRange(j+1, m)
		s0 := sum0.QueryRange(j+1, m)

		res += c0*(i-pref[i]) - s0

		sum0.Add(j, i-pref[i])
		cnt.Add(j, 1)
	}

	return res
}

type BIT []int

func (bit BIT) Add(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) Query(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) QueryRange(l, r int) int {
	return bit.Query(r) - bit.Query(l-1)
}
