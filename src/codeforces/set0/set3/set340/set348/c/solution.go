package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func drive(reader *bufio.Reader) []int {
	n, m, q := readThreeNums(reader)
	a := readNNums(reader, n)
	sets := make([][]int, m)
	for i := range m {
		var k int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &k) + 1
		sets[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &sets[i][j]) + 1
		}
	}
	queries := make([]string, q)
	for i := range q {
		queries[i] = readString(reader)
	}
	return solve(a, sets, queries)
}

func solve(a []int, sets [][]int, queries []string) []int {
	n := len(a)

	blockSize := int(math.Sqrt(float64(n * 2)))

	var big []int

	for i := range sets {
		if len(sets[i]) > blockSize {
			big = append(big, i)
		}
	}

	m := len(big)
	// m < blockSize

	overlap := make([][]int, len(sets))
	for i := range sets {
		overlap[i] = make([]int, m)
	}

	flag := make([]int, n+1)
	sum := make([]int, m)

	for j, x := range big {
		for _, i := range sets[x] {
			flag[i] = 1
			sum[j] += a[i-1]
		}

		for i, cur := range sets {
			var cnt int
			for _, y := range cur {
				cnt += flag[y]
			}
			overlap[i][j] = cnt
		}

		for _, i := range sets[x] {
			flag[i] = 0
		}
	}

	tag := make([]int, m)

	updateSmall := func(k int, x int) {
		for _, i := range sets[k] {
			a[i-1] += x
		}
		for j := range m {
			sum[j] += x * overlap[k][j]
		}
	}

	updateLarge := func(k int, x int) {
		tag[sort.SearchInts(big, k)] += x
	}

	querySmall := func(k int) int {
		var res int
		for _, i := range sets[k] {
			res += a[i-1]
		}
		for j := range m {
			res += tag[j] * overlap[k][j]
		}
		return res
	}

	queryLarge := func(k int) int {
		res := sum[sort.SearchInts(big, k)]
		for j := range m {
			res += tag[j] * overlap[k][j]
		}
		return res
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == '+' {
			var k, x int
			pos := readInt([]byte(cur), 2, &k) + 1
			readInt([]byte(cur), pos, &x)
			k--
			if len(sets[k]) <= blockSize {
				updateSmall(k, x)
			} else {
				updateLarge(k, x)
			}
		} else {
			var k int
			readInt([]byte(cur), 2, &k)
			k--
			if len(sets[k]) <= blockSize {
				ans = append(ans, querySmall(k))
			} else {
				ans = append(ans, queryLarge(k))
			}
		}
	}

	return ans
}
