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
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	ma := slices.Max(a)
	var primes []int
	lpf := make([]int, ma+1)

	for i := 2; i <= ma; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p > ma {
				break
			}
			lpf[i*p] = p
			if i%p == 0 {
				break
			}
		}
	}

	m := len(primes)
	sum := make([]int, m+1)

	pos := make([]int, ma+1)
	for i, p := range primes {
		pos[p] = i
	}

	marked := make([]bool, ma+1)

	for _, num := range a {
		for tmp := num; tmp > 1; tmp /= lpf[tmp] {
			// x是质数
			x := lpf[tmp]
			if !marked[x] {
				sum[pos[x]]++
				marked[x] = true
			}
		}
		for tmp := num; tmp > 1; tmp /= lpf[tmp] {
			marked[lpf[tmp]] = false
		}
	}

	for i := 1; i <= m; i++ {
		sum[i] += sum[i-1]
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		j := sort.SearchInts(primes, l)
		// primes[j] >= l
		k := sort.SearchInts(primes, r)
		if k < m && primes[k] > r {
			k--
		}
		ans[i] = sum[k]
		if j > 0 {
			ans[i] -= sum[j-1]
		}
	}
	return ans
}
