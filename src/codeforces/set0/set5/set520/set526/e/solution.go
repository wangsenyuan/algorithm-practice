package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	for _, x := range res {
		fmt.Println(x)
	}
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
	n, q := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([]int, q)
	for i := range q {
		queries[i] = readNum(reader)
	}
	return solve(a, queries)
}

func solve(a []int, queries []int) []int {
	n := len(a)
	sum := make([]int, n+1)
	for i, x := range a {
		sum[i+1] = sum[i] + x
	}

	ans := make([]int, len(queries))

	f := make([]int, n+1)
	begin := make([]int, n+1)
	for j, b := range queries {
		if sum[n] <= b {
			ans[j] = 1
			continue
		}
		var l int
		f[0] = 0
		for i := 1; i <= n; i++ {
			for sum[i]-sum[l] > b {
				l++
			}
			f[i] = f[l] + 1
			if f[i] == 2 {
				begin[i] = l
			} else {
				begin[i] = begin[l]
			}
			if sum[n]-(sum[i]-sum[begin[i]]) <= b {
				ans[j] = max(f[i], 2)
				break
			}
		}
	}

	return ans
}
