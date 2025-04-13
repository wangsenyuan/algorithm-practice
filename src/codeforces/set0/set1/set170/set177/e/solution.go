package main

import (
	"bufio"
	"fmt"
	"math/big"
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
	n, c := readTwoNums(reader)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		a[i], b[i] = readTwoNums(reader)
	}
	return solve(n, c, a, b)
}

const inf = 1e18

func solve(n int, c int, a []int, b []int) int {
	if c < n {
		return 0
	}

	if slices.Max(a) == 0 {
		// 只能待一天
		if c == n {
			return -1
		}
	}

	bigC := big.NewInt(int64(c))
	count := func(x int) int {
		var sum int
		tmp := big.NewInt(int64(x))
		for i, v := range a {
			y := big.NewInt(int64(v))
			y = y.Mul(y, tmp)
			y = y.Div(y, big.NewInt(int64(b[i])))
			y = y.Add(y, big.NewInt(1))
			if y.Cmp(bigC) > 0 {
				return c + 1
			}
			sum += int(y.Int64())
			if sum > c {
				break
			}
		}
		return sum
	}
	r := sort.Search(inf, func(x int) bool {
		return count(x) > c
	})

	l := sort.Search(r+1, func(x int) bool {
		return count(x) >= c
	})

	if l == 0 {
		l++
	}

	return max(0, r-l)
}
